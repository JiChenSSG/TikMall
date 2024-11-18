package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	ID          int64 `gorm:"primary_key"`
	Name        string
	Description string
	Price       float32
	Picture     string

	Categories []Category `gorm:"many2many:category_products;"`

	gorm.Model
}

func CreateProduct(db *gorm.DB, ctx context.Context, product *Product) (err error) {
	return db.WithContext(ctx).Create(product).Error
}

func GetProduct(db *gorm.DB, ctx context.Context, id int64) (product *Product, err error) {
	return product, db.WithContext(ctx).Preload("Categories").First(product, id).Error
}

func UpdateProduct(db *gorm.DB, ctx context.Context, product *Product) (err error) {
	tx := db.WithContext(ctx).Begin()
	if err = tx.Model(product).Association("Categories").Replace(product.Categories); err != nil {
		tx.Rollback()
		return
	}

	product.Categories = nil

	if err = tx.Model(product).Updates(product).Error; err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}

func DeleteProduct(db *gorm.DB, ctx context.Context, id int64) (err error) {
	return db.WithContext(ctx).Delete(&Product{}, id).Error
}

func ListProducts(db *gorm.DB, ctx context.Context, page int, pageSize int, categoryId int64) (products []Product, err error) {
	return products, db.WithContext(ctx).Preload("Categories").Where("category_id = ?", categoryId).Limit(pageSize).Offset((page - 1) * pageSize).Find(&products).Error
}

func SearchProducts(db *gorm.DB, ctx context.Context, keyword string) (products []Product, err error) {
	return products, db.WithContext(ctx).Preload("Categories").Where("name like ? OR descrpition like ?", "%"+keyword+"%", "%"+keyword+"%").Find(&products).Error
}