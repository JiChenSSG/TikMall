package model

import "gorm.io/gorm"

type Product struct {
	ID          int64 `gorm:"primary_key"`
	Name        string
	Description string
	Price       float64
	Picture     string

	Categories []Category `gorm:"many2many:category_products;"`

	gorm.Model
}
