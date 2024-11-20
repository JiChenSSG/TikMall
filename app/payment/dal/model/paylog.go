package model

import (
	"context"

	"gorm.io/gorm"
)

type PayLog struct {
	ID            int64 `gorm:"primary_key"`
	OrderID       string
	UserID        int64
	Amount        float32
	TranscationID string

	PayAt string `gorm:"autoCreateTime"`
}

func CreatePayLog(db *gorm.DB, ctx context.Context, payLog *PayLog) error {
	return db.WithContext(ctx).Create(payLog).Error
}

func GetPayLogByOrderID(db *gorm.DB, ctx context.Context, orderID string) (*PayLog, error) {
	var payLog PayLog
	err := db.WithContext(ctx).Where("order_id = ?", orderID).First(&payLog).Error
	return &payLog, err
}

func GetPayLogByTranscationID(db *gorm.DB, ctx context.Context, transcationID string) (*PayLog, error) {
	var payLog PayLog
	err := db.WithContext(ctx).Where("transcation_id = ?", transcationID).First(&payLog).Error
	return &payLog, err
}

func GetPayLogsByUserID(db *gorm.DB, ctx context.Context, userID int64) ([]PayLog, error) {
	var payLogs []PayLog
	err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&payLogs).Error
	return payLogs, err
}