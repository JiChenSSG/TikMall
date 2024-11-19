package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	ID   int64 `gorm:"primary_key"`
	Name string

	Products []Product `gorm:"many2many:category_products;"`

	gorm.Model
}

func CreateCategory(db *gorm.DB, ctx context.Context, catrgory *Category) (err error) {
	return db.WithContext(ctx).Create(catrgory).Error
}

func GetCategory(db *gorm.DB, ctx context.Context, id int64) (category *Category, err error) {
	return category, db.WithContext(ctx).First(category, id).Error
}

func UpdateCategory(db *gorm.DB, ctx context.Context, category *Category) (err error) {
	return db.WithContext(ctx).Updates(category).Error
}

func DeleteCategory(db *gorm.DB, ctx context.Context, id int64) (err error) {
	return db.WithContext(ctx).Delete(&Category{}, id).Error
}

func ListCategories(db *gorm.DB, ctx context.Context) (categories []Category, err error) {
	return categories, db.WithContext(ctx).Find(&categories).Error
}