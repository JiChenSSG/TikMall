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

func GetUserByID(db *gorm.DB, ctx context.Context, id int64) (*User, error) {
	klog.Infof("Get user by id: %v", id)

	user := new(User)
	if err := db.WithContext(ctx).First(user, id).Error; err != nil {
		klog.Errorf("Get user failed: %v", err)
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(db *gorm.DB, ctx context.Context, email string) (*User, error) {
	klog.Infof("Get user by email: %v", email)

	user := new(User)
	if err := db.WithContext(ctx).Where("email = ?", email).First(user).Error; err != nil {
		klog.Errorf("Get user failed: %v", err)
		return nil, err
	}

	return user, nil
}

func DeleteUser(db *gorm.DB, ctx context.Context, id int64) error {
	klog.Infof("Delete user by id: %v", id)

	if err := db.WithContext(ctx).Delete(&User{}, id).Error; err != nil {
		klog.Errorf("Delete user failed: %v", err)
		return err
	}

	return nil
}

func UpdateUser(db *gorm.DB, ctx context.Context, user *User) (*User, error) {
	klog.Infof("Update user: %v", user)

	if err := db.WithContext(ctx).Updates(user).Error; err != nil {
		klog.Errorf("Update user failed: %v", err)
		return nil, err
	}

	return user, nil
}
