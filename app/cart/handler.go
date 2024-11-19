package main

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jichenssg/tikmall/app/cart/dal/model"
	"github.com/jichenssg/tikmall/app/common/dal/mysql"
	cart "github.com/jichenssg/tikmall/gen/kitex_gen/cart"
	"gorm.io/gorm"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	klog.Infof("AddItem service")

	userID := req.UserId
	productID := req.Item.ProductId
	quantity := int(req.Item.Quantity)

	klog.Debugf("AddItem service userID: %v, productID: %v, quantity: %v", userID, productID, quantity)

	for retry := 0; retry < 3; retry++ {
		if retry > 0 {
			klog.Warn("AddItem service retry")
		}

		var cart *model.Cart
		cart, err = model.GetCartByUserIDAndProductID(mysql.GetDB(), ctx, userID, productID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				klog.Errorf("AddItem service err: %v", err)
				err = kerrors.NewBizStatusError(500, "get cart failed")
				return nil, err
			}

			err := model.CreateCart(mysql.GetDB(), ctx, &model.Cart{
				UserId:    userID,
				ProductId: productID,
				Quantity:  quantity,
			})

			if err != nil {
				klog.Errorf("AddItem service err: %v", err)
				err = kerrors.NewBizStatusError(500, "add item failed")
				continue
			}
		}

		cart.Quantity += quantity
		err = mysql.GetDB().WithContext(ctx).Model(cart).Model(&model.Cart{}).
			Where("id = ? AND version = ?", cart.ID, cart.Version).
			Updates(map[string]interface{}{
				"quantity": cart.Quantity,
				"version":  cart.Version + 1,
			}).Error

		if err != nil {
			klog.Errorf("AddItem service err: %v", err)
			err = kerrors.NewBizStatusError(500, "add item failed")
			continue
		}

		break
	}

	if err != nil {
		klog.Errorf("AddItem service err: %v", err)
		err = kerrors.NewBizStatusError(500, "add item failed")
	}

	resp = &cart.AddItemResp{}

	return
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	klog.Infof("GetCart service")

	userID := req.UserId
	klog.Debugf("GetCart service userID: %v", userID)

	carts, err := model.GetCartByUserID(mysql.GetDB(), ctx, userID)
	if err != nil {
		klog.Errorf("GetCart service err: %v", err)
		err = kerrors.NewBizStatusError(500, "get cart failed")
	}

	rescart := &cart.Cart{
		UserId: userID,
		Items:  make([]*cart.CartItem, 0, len(carts)),
	}

	for _, c := range carts {
		cart2resp(rescart, &c)
	}

	resp = &cart.GetCartResp{
		Cart: rescart,
	}

	return
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	klog.Infof("EmptyCart service")

	userID := req.UserId
	klog.Debugf("EmptyCart service userID: %v", userID)

	err = model.DeleteCartByUserID(mysql.GetDB(), ctx, userID)

	if err != nil {
		klog.Errorf("EmptyCart service err: %v", err)
		err = kerrors.NewBizStatusError(500, "empty cart failed")
	}

	resp = &cart.EmptyCartResp{}

	return
}

func cart2resp(res *cart.Cart, c *model.Cart) {
	res.Items = append(res.Items, &cart.CartItem{
		ProductId: c.ProductId,
		Quantity:  int32(c.Quantity),
	})
}
