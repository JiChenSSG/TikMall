package utils

import (
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/jichenssg/tikmall/gateway/biz/model/frontend/common"
)

func ParseRpcError(err error) (code int, response *common.Response) {
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if isBizErr {
		return int(bizErr.BizStatusCode()), &common.Response{
			Message: bizErr.BizMessage(),
		}
	}

	return consts.StatusInternalServerError, &common.Response{
		Message: err.Error(),
	}
}
