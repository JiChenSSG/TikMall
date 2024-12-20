package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jichenssg/tikmall/app/common/client"
	"github.com/jichenssg/tikmall/app/common/dal/mysql"
	"github.com/jichenssg/tikmall/app/user/dal/model"
	"github.com/jichenssg/tikmall/app/user/utils"
	"github.com/jichenssg/tikmall/gen/kitex_gen/auth"
	userrpc "github.com/jichenssg/tikmall/gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userrpc.RegisterReq) (resp *userrpc.RegisterResp, err error) {
	klog.Info("User Register")

	if req.Password != req.ConfirmPassword {
		klog.Warnf("password not match")
		err = kerrors.NewBizStatusError(500, "password not match")
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Password: utils.String2MD5(req.Password),
		Email:    req.Email,
	}

	// use transaction to ensure use and role are created successfully
	tx := mysql.GetDB().Begin()

	err = tx.WithContext(ctx).Create(user).Error
	if err != nil {
		err = kerrors.NewBizStatusError(500, err.Error())
		return nil, err
	}

	authclient := client.AuthClient
	_, err = authclient.AddRole(ctx, &auth.AddRoleReq{
		UserId: user.ID,
		Role:   "user",
	})

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	resp = &userrpc.RegisterResp{
		UserId:  user.ID,
		Success: true,
	}

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userrpc.LoginReq) (resp *userrpc.LoginResp, err error) {
	klog.Info("User Login")

	user, err := model.GetUserByEmail(mysql.GetDB(), ctx, req.Email)
	if err != nil {
		err = kerrors.NewBizStatusError(500, "user not found")
		return nil, err
	}

	if user.Password != utils.String2MD5(req.Password) {
		err = kerrors.NewBizStatusError(500, "password incorrect")
		return nil, err
	}

	resp = &userrpc.LoginResp{
		UserId:  user.ID,
		Success: true,
	}

	return resp, nil
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *userrpc.InfoReq) (resp *userrpc.InfoResp, err error) {
	klog.Info("User Info")

	user, err := model.GetUserByID(mysql.GetDB(), ctx, req.UserId)
	if err != nil {
		err = kerrors.NewBizStatusError(500, err.Error())
		return nil, err
	}

	resp = &userrpc.InfoResp{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return resp, nil
}

// Delete implements the UserServiceImpl interface.
func (s *UserServiceImpl) Delete(ctx context.Context, req *userrpc.DeleteReq) (resp *userrpc.DeleteResp, err error) {
	klog.Info("User Delete")

	err = model.DeleteUser(mysql.GetDB(), ctx, req.UserId)
	if err != nil {
		err = kerrors.NewBizStatusError(500, err.Error())
		return nil, err
	}

	resp = &userrpc.DeleteResp{
		Success: true,
	}

	return resp, nil
}

// Update implements the UserServiceImpl interface.
func (s *UserServiceImpl) Update(ctx context.Context, req *userrpc.UpdateReq) (resp *userrpc.UpdateResp, err error) {
	klog.Info("User Update")

	user := &model.User{
		ID:       req.UserId,
		Username: req.Username,
		Email:    req.Email,
	}

	_, err = model.UpdateUser(mysql.GetDB(), ctx, user)
	if err != nil {
		err = kerrors.NewBizStatusError(500, err.Error())
		return nil, err
	}

	resp = &userrpc.UpdateResp{
		Success: true,
	}

	return resp, nil
}
