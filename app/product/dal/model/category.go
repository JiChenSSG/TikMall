package model

import "gorm.io/gorm"

type Category struct {
	ID   int64 `gorm:"primary_key"`
	Name string

	Products []Product `gorm:"many2many:category_products;"`

	gorm.Model
}