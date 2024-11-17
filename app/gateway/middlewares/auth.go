package middlewares

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jichenssg/tikmall/gateway/client"
	"github.com/jichenssg/tikmall/gateway/utils"
	"github.com/jichenssg/tikmall/gen/kitex_gen/auth"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := string(c.Request.URI().RequestURI())
		method := string(c.Request.Method())
		token := c.Request.Header.Get("Authorization")

		hlog.CtxDebugf(ctx, "Auth middleware: %v %v %v", path, method, token)

		// Check if the user is authorized to access the path
		authclient := client.AuthClient
		resp, err := authclient.VerifyToken(ctx, &auth.VerifyTokenReq{
			Token:  token,
			Path:   path,
			Method: method,
		})

		if err != nil {
			c.JSON(utils.ParseRpcError(err))
			c.Abort()
		}

		// put the user id into the context
		c.Set("user_id", resp.UserId)
	}
}
