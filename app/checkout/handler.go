package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jichenssg/tikmall/app/common/client"
	checkout "github.com/jichenssg/tikmall/gen/kitex_gen/checkout"
	order "github.com/jichenssg/tikmall/gen/kitex_gen/order"
	payment "github.com/jichenssg/tikmall/gen/kitex_gen/payment"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	klog.Infof("Checkout service")

	userID := req.UserId
	orderID := req.OrderId
	card := req.CreditCard

	klog.CtxDebugf(ctx, "Checkout service userID: %v, orderID: %v, card: %v", userID, orderID, card)

	orderclient := client.OrderClient
	getorderrresp, err := orderclient.GetOrder(ctx, &order.GetOrderReq{
		OrderId: orderID,
	})
	if err != nil {
		klog.Errorf("Checkout service err: %v", err)
		return nil, err
	}

	if getorderrresp.Order.Status != "PLACED" {
		klog.Errorf("Checkout service err: %v", err)
		err = kerrors.NewBizStatusError(500, "order status not PLACED")
		return
	}

	if getorderrresp.Order.UserId != userID {
		klog.Errorf("Checkout service err: %v", err)
		err = kerrors.NewBizStatusError(500, "order not match")
		return nil, err
	}

	paymentclient := client.PaymentClient
	payresp, err := paymentclient.Charge(ctx, &payment.ChargeReq{
		UserId:  userID,
		OrderId: orderID,
		Amount:  func() float32 {
			var amount float32
			for _, item := range getorderrresp.Order.OrderItems {
				amount += item.Cost
			}
			return amount
		}(),
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          card.CreditCardNumber,
			CreditCardExpirationMonth: card.CreditCardExpirationMonth,
			CreditCardExpirationYear:  card.CreditCardExpirationYear,
			CreditCardCvv:             card.CreditCardCvv,
		},
	})

	if err != nil {
		klog.Errorf("Checkout service err: %v", err)
		return nil, err
	}

	_, err = orderclient.MarkOrderPaid(ctx, &order.MarkOrderPaidReq{
		OrderId: orderID,
	})

	if err != nil {
		klog.Errorf("Checkout service err: %v", err)
		return nil, err
	}

	resp = &checkout.CheckoutResp{
		TransactionId: payresp.TransactionId,
	}

	return
}
