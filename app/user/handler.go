package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jichenssg/tikmall/app/user/dal/model"
	"github.com/jichenssg/tikmall/app/user/dal/mysql"
	"github.com/jichenssg/tikmall/gen/kitex_gen/user"
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
		Password: req.Password,
		Email:    req.Email,
	}

	user, err = model.CreateUser(mysql.GetDB(), ctx, user)
	if err != nil {
		err = kerrors.NewBizStatusError(500, err.Error())
		return nil, err
	}

	resp = &userrpc.RegisterResp{
		UserId:  user.ID,
		Success: true,
	}

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoReq) (resp *user.InfoResp, err error) {
	// TODO: Your code here...
	return
}

// Delete implements the UserServiceImpl interface.
func (s *UserServiceImpl) Delete(ctx context.Context, req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	// TODO: Your code here...
	return
}

// Update implements the UserServiceImpl interface.
func (s *UserServiceImpl) Update(ctx context.Context, req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	// TODO: Your code here...
	return
}
