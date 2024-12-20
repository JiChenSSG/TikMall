package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jichenssg/tikmall/app/auth/config"
	"github.com/jichenssg/tikmall/app/auth/dal/model"
	"github.com/jichenssg/tikmall/app/auth/utils"
	"github.com/jichenssg/tikmall/app/common/dal/redis"
	auth "github.com/jichenssg/tikmall/gen/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverToken(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	klog.Infof("Deliver token for user: %v", req.UserId)

	userID := req.UserId

	access, err := utils.GenToken(userID, time.Duration(config.GetConf().Token.AccessExpire)*time.Second)
	if err != nil {
		klog.CtxErrorf(ctx, "Generate access token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Generate access token failed")
		return
	}

	refresh, err := utils.GenToken(userID, 0)
	if err != nil {
		klog.CtxErrorf(ctx, "Generate refresh token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Generate refresh token failed")
		return
	}

	err = model.CacheToken(redis.GetClient(), ctx, refresh, userID)
	if err != nil {
		klog.CtxErrorf(ctx, "Cache refresh token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Cache refresh token failed")
		return
	}

	resp = &auth.DeliveryResp{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	return resp, nil
}

// RefreshToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *auth.RefreshTokenReq) (resp *auth.RefreshResp, err error) {
	klog.Infof("Refresh token: %v", req.RefreshToken)

	refreshToken := req.RefreshToken

	userID, err := utils.ParseToken(refreshToken)
	if err != nil {
		klog.CtxErrorf(ctx, "Parse refresh token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Parse refresh token failed")
		return
	}

	accessToken, err := utils.GenToken(userID, time.Duration(config.GetConf().Token.AccessExpire)*time.Second)
	if err != nil {
		klog.CtxErrorf(ctx, "Generate access token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Generate access token failed")
		return
	}

	err = model.ExtendToken(redis.GetClient(), ctx, userID)
	if err != nil {
		klog.CtxErrorf(ctx, "Extend refresh token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Extend refresh token failed")
		return
	}

	return &auth.RefreshResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// VerifyToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyToken(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	klog.Infof("Verify token: %v", req.Token)

	accessToken := req.Token
	url := req.Path
	method := req.Method

	var allowed bool
	var userID int64

	if accessToken == "" {
		userID = 0
		allowed, err = utils.EnforceAnonymous(url, method)
		if err != nil {
			klog.CtxErrorf(ctx, "Enforce failed: %v", err)
			err = kerrors.NewBizStatusError(500, "Enforce failed")
			return nil, err
		}

	} else {
		userID, err = utils.ParseToken(accessToken)
		if err != nil {
			klog.CtxErrorf(ctx, "Parse access token failed: %v", err)
			err = kerrors.NewBizStatusError(500, "Parse access token failed")
			return
		}

		allowed, err = utils.Enforce(userID, url, method)
		if err != nil {
			klog.CtxErrorf(ctx, "Enforce failed: %v", err)
			err = kerrors.NewBizStatusError(500, "Enforce failed")
			return
		}
	}

	if !allowed {
		klog.CtxWarnf(ctx, "Permission denied")
		err = kerrors.NewBizStatusError(403, "Permission denied")
		return
	}

	return &auth.VerifyResp{
		UserId: userID,
	}, nil
}

// DeleteToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeleteToken(ctx context.Context, req *auth.DeleteTokenReq) (resp *auth.DeleteTokenResp, err error) {
	klog.Infof("Delete token: %v", req.Token)

	accessToken := req.Token

	userID, err := utils.ParseToken(accessToken)
	if err != nil {
		klog.CtxErrorf(ctx, "Parse access token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Parse access token failed")
		return
	}

	err = model.DeleteToken(redis.GetClient(), ctx, userID)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Delete token failed")
		return
	}

	return &auth.DeleteTokenResp{}, nil
}

// AddRole implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) AddRole(ctx context.Context, req *auth.AddRoleReq) (resp *auth.AddRoleResp, err error) {
	klog.Infof("Add role: %v", req.Role)

	userID := req.UserId
	role := req.Role

	if role == "admin" {
		err = utils.AddAdminRole(userID)
	} else {
		err = utils.AddUserRole(userID)
	}

	if err != nil {
		klog.CtxErrorf(ctx, "Add role failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Add role failed")
		return
	}

	return &auth.AddRoleResp{}, nil
}

// RemoveRole implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RemoveRole(ctx context.Context, req *auth.RemoveRoleReq) (resp *auth.RemoveRoleResp, err error) {
	klog.Infof("Remove role: %v", req.Role)

	userID := req.UserId
	role := req.Role

	if role == "admin" {
		err = utils.RemoveAdminRole(userID)
	} else {
		err = utils.RemoveUserRole(userID)
	}

	if err != nil {
		klog.CtxErrorf(ctx, "Remove role failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Remove role failed")
		return
	}

	return &auth.RemoveRoleResp{}, nil
}

// GetRoles implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GetRoles(ctx context.Context, req *auth.GetRolesReq) (resp *auth.GetRolesResp, err error) {
	klog.Infof("Get roles: %v", req.UserId)

	userID := req.UserId

	roles, err := utils.GetRoles(userID)
	if err != nil {
		klog.CtxErrorf(ctx, "Get roles failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Get roles failed")
		return
	}

	return &auth.GetRolesResp{
		Roles: roles,
	}, nil
}

// RemoveAllRoles implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RemoveAllRoles(ctx context.Context, req *auth.RemoveAllRolesReq) (resp *auth.RemoveAllRolesResp, err error) {
	klog.Infof("Remove all roles: %v", req.UserId)

	userID := req.UserId

	err = utils.RemoveAllRoles(userID)
	if err != nil {
		klog.CtxErrorf(ctx, "Remove all roles failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Remove all roles failed")
		return
	}

	return &auth.RemoveAllRolesResp{}, nil
}
