package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string `json:"email"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       int32  `json:"zip_code"`
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "PLACED"
	OrderStatePaid     OrderState = "PAID"
	OrderStateCanceled OrderState = "CANCALED"
)

type Order struct {
	ID         int64       `gorm:"primaryKey"`
	OrderID    string      `gorm:"uniqueIndex;size:256"`
	UserID     int64       `gorm:"index"`
	Consignee  Consignee   `gorm:"embedded"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:OrderID"`
	OrderState OrderState
	PaidTime   time.Time

	gorm.Model
}

func CreateOrder(db *gorm.DB, ctx context.Context, order *Order) error {
	return db.WithContext(ctx).Create(order).Error
}

func GetOrderByOrderID(db *gorm.DB, ctx context.Context, orderID string) (*Order, error) {
	var order Order
	err := db.WithContext(ctx).Where("order_id = ?", orderID).Preload("OrderItems").First(&order).Error
	return &order, err
}

func GetOrdersByUserID(db *gorm.DB, ctx context.Context, userID int64) ([]Order, error) {
	var orders []Order
	err := db.WithContext(ctx).Where("user_id = ?", userID).Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func PayOrder(db *gorm.DB, ctx context.Context, orderID string) error {
	return db.WithContext(ctx).Model(&Order{}).Where("order_id = ?", orderID).
		Updates(map[string]interface{}{"paid_time": time.Now(), "order_state": OrderStatePaid}).Error
}

func CancelOrder(db *gorm.DB, ctx context.Context, orderID string) error {
	return db.WithContext(ctx).Model(&Order{}).Where("order_id = ?", orderID).Update("order_state", OrderStateCanceled).Error
}
