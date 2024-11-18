// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	order "github.com/jichenssg/tikmall/app/gateway/biz/router/order"
	user "github.com/jichenssg/tikmall/app/gateway/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	order.Register(r)

	user.Register(r)
}
