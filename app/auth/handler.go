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
		klog.Errorf("Generate access token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Generate access token failed")
		return
	}

	refresh, err := utils.GenToken(userID, 0)
	if err != nil {
		klog.Errorf("Generate refresh token failed: %v", err)
		err = kerrors.NewBizStatusError(500, "Generate refresh token failed")
		return
	}

	err = model.CacheToken(redis.GetClient(), ctx, refresh, userID)
	if err != nil {
		klog.Errorf("Cache refresh token failed: %v", err)
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
	// TODO: Your code here...
	return
}

// VerifyToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyToken(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	return
}
