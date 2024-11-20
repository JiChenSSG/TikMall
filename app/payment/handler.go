package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/jichenssg/tikmall/app/common/dal/mysql"
	"github.com/jichenssg/tikmall/app/payment/dal/model"
	payment "github.com/jichenssg/tikmall/gen/kitex_gen/payment"
	"gorm.io/gorm"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	klog.Infof("Charge service")

	card := req.CreditCard
	amount := req.Amount
	orderID := req.OrderId
	userID := req.UserId

	klog.CtxDebugf(ctx, "Charge service card: %v, amount: %v, orderID: %v, userID: %v", card, amount, orderID, userID)

	cardValidator := &creditcard.Card{
		Number: card.CreditCardNumber,
		Month:  fmt.Sprint(card.CreditCardExpirationMonth),
		Year:   fmt.Sprint(card.CreditCardExpirationYear),
		Cvv:    fmt.Sprint(card.CreditCardCvv),
	}

	err = cardValidator.Validate()
	if err != nil {
		klog.Errorf("Charge service card validate failed: %v", err)
		err = kerrors.NewBizStatusError(400, "card validate failed")
		return nil, err
	}

	payLog := &model.PayLog{
		UserID:  userID,
		OrderID: orderID,
		Amount:  amount,
	}

	for retry := 0; retry < 3; retry++ {
		payLog.TranscationID = uuid.NewString()

		err = model.CreatePayLog(mysql.GetDB(), ctx, payLog)

		if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
			klog.Warn("order uuid conflict, retry")
			continue
		}

		break
	}
	if err != nil {
		klog.Errorf("Charge service err: %v", err)
		return nil, err
	}

	resp = &payment.ChargeResp{
		TransactionId: payLog.TranscationID,
	}

	return
}
