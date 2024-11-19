package model

type Cart struct {
	ID        int64 `gorm:"primary_key"`
	UserId    int64 `gorm:"index"`
	ProductId int64
	Quantity  int
}

