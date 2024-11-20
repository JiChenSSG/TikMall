package main

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/jichenssg/tikmall/app/common/dal/mysql"
	"github.com/jichenssg/tikmall/app/order/config"
	"github.com/jichenssg/tikmall/app/order/dal/model"
	"github.com/jichenssg/tikmall/gen/kitex_gen/cart"
	order "github.com/jichenssg/tikmall/gen/kitex_gen/order"
	"gorm.io/gorm"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	klog.Infof("PlaceOrder service")

	userID := req.UserId
	address := req.Address
	items := req.OrderItems
	email := req.Email

	klog.CtxDebugf(ctx, "PlaceOrder service userID: %v, address: %v, items: %v, email: %v", userID, address, items, email)

	o := &model.Order{
		UserID: userID,
		Consignee: model.Consignee{
			Email:         email,
			StreetAddress: address.StreetAddress,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
		},
		OrderItems: func() []model.OrderItem {
			itemsModel := make([]model.OrderItem, 0, len(items))
			for _, item := range items {
				itemsModel = append(itemsModel, model.OrderItem{
					ProductID: item.Item.ProductId,
					Quantity:  item.Item.Quantity,
					Cost:      item.Cost,
				})
			}
			return itemsModel
		}(),
		OrderID: uuid.NewString(),
		OrderState: model.OrderStatePlaced,
	}

	for retry := 0; retry < 3; retry++ {
		err = model.CreateOrder(mysql.GetDB(), ctx, o)
		if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
			klog.Warn("order uuid conflict, retry")
			o.OrderID = uuid.NewString()
			continue
		}

		break
	}

	if err != nil {
		klog.Errorf("PlaceOrder service err: %v", err)
		err = kerrors.NewBizStatusError(500, "place order failed")
		return nil, err
	}

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: o.OrderID,
		},
	}

	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	klog.Infof("ListOrder service")

	userID := req.UserId
	klog.CtxDebugf(ctx, "ListOrder service userID: %v", userID)

	orders, err := model.GetOrdersByUserID(mysql.GetDB(), ctx, userID)
	if err != nil {
		klog.Errorf("ListOrder service err: %v", err)
		err = kerrors.NewBizStatusError(500, "list order failed")
		return nil, err
	}

	// check if order is timeout
	for _, o := range orders {
		if ok := checkTimeout(&o); !ok {
			klog.Info("ListOrder service order timeout")

			err = model.CancelOrder(mysql.GetDB(), ctx, o.OrderID)
			if err != nil {
				klog.Errorf("ListOrder service err: %v", err)
				err = kerrors.NewBizStatusError(500, "fail to update order status")

				return nil, err
			}

			o.OrderState = model.OrderStateCanceled
		}

	}

	resp = &order.ListOrderResp{
		Orders: func() []*order.Order {
			ordersResp := make([]*order.Order, 0, len(orders))
			for _, o := range orders {
				ordersResp = append(ordersResp, order2resp(&o))
			}
			return ordersResp
		}(),
	}

	return
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	klog.Infof("MarkOrderPaid service")

	orderID := req.OrderId
	klog.CtxDebugf(ctx, "MarkOrderPaid service orderID: %v", orderID)

	err = model.PayOrder(mysql.GetDB(), ctx, orderID)
	if err != nil {
		klog.Errorf("MarkOrderPaid service err: %v", err)
		err = kerrors.NewBizStatusError(500, "mark order paid failed")
		return nil, err
	}

	resp = &order.MarkOrderPaidResp{}

	return
}

func order2resp(o *model.Order) *order.Order {
	return &order.Order{
		OrderId:   o.OrderID,
		UserId:    o.UserID,
		Status:    string(o.OrderState),
		CreatedAt: int32(o.CreatedAt.Unix()),
		Email:     o.Consignee.Email,
		Address: &order.Address{
			StreetAddress: o.Consignee.StreetAddress,
			City:          o.Consignee.City,
			State:         o.Consignee.State,
			Country:       o.Consignee.Country,
			ZipCode:       o.Consignee.ZipCode,
		},
		OrderItems: func() []*order.OrderItem {
			var items []*order.OrderItem
			for _, item := range o.OrderItems {
				items = append(items, &order.OrderItem{
					Item: &cart.CartItem{
						ProductId: item.ProductID,
						Quantity:  int32(item.Quantity),
					},

					Cost: item.Cost,
				})
			}
			return items
		}(),
	}
}

// MarkOrderCancelled implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderCancelled(ctx context.Context, req *order.CancelOrderReq) (resp *order.CancelOrderResp, err error) {
	klog.Infof("MarkOrderCancelled service")

	orderID := req.OrderId
	klog.CtxDebugf(ctx, "MarkOrderCancelled service orderID: %v", orderID)

	err = model.CancelOrder(mysql.GetDB(), ctx, orderID)
	if err != nil {
		klog.Errorf("MarkOrderCancelled service err: %v", err)
		err = kerrors.NewBizStatusError(500, "mark order cancelled failed")
		return nil, err
	}

	resp = &order.CancelOrderResp{}

	return
}

// GetOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	klog.Infof("GetOrder service")

	orderID := req.OrderId
	klog.CtxDebugf(ctx, "GetOrder service orderID: %v", orderID)

	o, err := model.GetOrderByOrderID(mysql.GetDB(), ctx, orderID)
	if err != nil {
		klog.Errorf("GetOrder service err: %v", err)
		err = kerrors.NewBizStatusError(500, "get order failed")
		return nil, err
	}

	// check if order is timeout
	if ok := checkTimeout(o); !ok {
		klog.Info("GetOrder service order timeout")

		err = model.CancelOrder(mysql.GetDB(), ctx, o.OrderID)
		if err != nil {
			klog.Errorf("GetOrder service err: %v", err)
			err = kerrors.NewBizStatusError(500, "fail to update order status")
			return nil, err
		}

		o.OrderState = model.OrderStateCanceled
	}

	resp = &order.GetOrderResp{
		Order: order2resp(o),
	}

	return
}

// UpdateOrderInfo implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderInfo(ctx context.Context, req *order.UpdateOrderInfoReq) (resp *order.UpdateOrderInfoResp, err error) {
	klog.Infof("UpdateOrderInfo service")

	orderID := req.OrderId
	address := req.Address
	email := req.Email
	klog.CtxDebugf(ctx, "UpdateOrderInfo service orderID: %v, address: %v, email: %v", orderID, address, email)

	err = model.UpdateOrder(mysql.GetDB(), ctx, &model.Order{
		Consignee: model.Consignee{
			Email:         email,
			StreetAddress: address.StreetAddress,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
		},

		OrderID: orderID,
	})

	if err != nil {
		klog.Errorf("UpdateOrderInfo service err: %v", err)
		err = kerrors.NewBizStatusError(500, "update order info failed")
		return nil, err
	}

	return
}

func checkTimeout(o *model.Order) (ok bool) {
	return time.Since(o.CreatedAt) < time.Duration(config.GetConf().Order.CancelTimeout)*time.Minute || o.Consignee.State != string(model.OrderStatePlaced)
}
