package model

import (
	"context"

	"gorm.io/gorm"
)

type Cart struct {
	ID        int64 `gorm:"primary_key"`
	UserId    int64 `gorm:"index"`
	ProductId int64
	Quantity  int
	Version   int
	gorm.Model
}

func CreateCart(db *gorm.DB, ctx context.Context, cart *Cart) (err error) {
	return db.WithContext(ctx).Create(cart).Error
}

func GetCartByUserID(db *gorm.DB, ctx context.Context, userID int64) (cart []Cart, err error) {
	return cart, db.WithContext(ctx).Where("user_id = ?", userID).Find(&cart).Error
}

func GetCartByUserIDAndProductID(db *gorm.DB, ctx context.Context, userID int64, productID int64) (cart *Cart, err error) {
	cart = &Cart{}
	return cart, db.WithContext(ctx).Where("user_id = ? AND product_id = ?", userID, productID).First(cart).Error
}

func UpdateCart(db *gorm.DB, ctx context.Context, cart *Cart) (err error) {
	return db.WithContext(ctx).Model(cart).Updates(cart).Error
}

func DeleteCartByUserID(db *gorm.DB, ctx context.Context, userID int64) (err error) {
	return db.WithContext(ctx).Where("user_id = ?", userID).Delete(&Cart{}).Error
}
