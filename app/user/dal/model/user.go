package model

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID int64 `json:"id" gorm:"primaryKey"`

	Username  string                `json:"username"`
	Password  string                `json:"password"`
	Email     string                `json:"email" gorm:"type:varchar(255);uniqueIndex:udx_email"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"uniqueIndex:udx_email"`

	gorm.Model
}

func (User) TableName() string {
	return "user"
}

func CreateUser(db *gorm.DB, ctx context.Context, user *User) (*User, error) {
	klog.Infof("Create user: %v", user)

	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		klog.Errorf("Create user failed: %v", err)
		return nil, err
	}

	return user, nil
}
