// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	cart "github.com/jichenssg/tikmall/app/gateway/biz/router/cart"
	checkout "github.com/jichenssg/tikmall/app/gateway/biz/router/checkout"
	order "github.com/jichenssg/tikmall/app/gateway/biz/router/order"
	product "github.com/jichenssg/tikmall/app/gateway/biz/router/product"
	user "github.com/jichenssg/tikmall/app/gateway/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	checkout.Register(r)

	cart.Register(r)

	product.Register(r)

	order.Register(r)

	user.Register(r)
}
