package model

type OrderItem struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	OrderID   string `json:"order_id" gorm:"index"`
	ProductID int64  `json:"product_id"`
	Quantity  int32  `json:"quantity"`
	Cost      float32  `json:"cost"`
}
