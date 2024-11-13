package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	Username  string                `json:"username"`
	Password  string                `json:"password"`
	Email     string                `json:"email" gorm:"uniqueIndex:udx_name"`
	DeletedAt soft_delete.DeletedAt `gorm:"index" gorm:"uniqueIndex:udx_name"`

	gorm.Model
}
