package client

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/jichenssg/tikmall/app/common/clientsuite"
	"github.com/jichenssg/tikmall/gen/kitex_gen/auth/authservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/cart/cartservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/order/orderservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/product/productcatalogservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/user/userservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/payment/paymentservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/checkout/checkoutservice"
)

var (
	AuthClient    authservice.Client
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	OrderClient   orderservice.Client
	PaymentClient paymentservice.Client
	CheckoutClient checkoutservice.Client
	Once          sync.Once
)

func Init(serviceName, registryEndpoint, otelEndpoint string) {
	commonClientSuite := client.WithSuite(clientsuite.CommonClientSuite{
		ServiceName:      serviceName,
		RegistryEndpoint: registryEndpoint,
		OtelEndpoint:     otelEndpoint,
	})

	initAuthClient(commonClientSuite)
	initUserClient(commonClientSuite)
	initProductClient(commonClientSuite)
	initCartClient(commonClientSuite)
	initOrderClient(commonClientSuite)
	initPaymentClient(commonClientSuite)
	initCheckoutClient(commonClientSuite)
}

func initAuthClient(suite client.Option) {
	var err error
	AuthClient, err = authservice.NewClient(
		"auth",
		suite,
	)

	if err != nil {
		panic(err)
	}
}

func initUserClient(suite client.Option) {
	var err error
	UserClient, err = userservice.NewClient(
		"user",
		suite,
	)

	if err != nil {
		panic(err)
	}
}

func initProductClient(suite client.Option) {
	var err error
	ProductClient, err = productcatalogservice.NewClient(
		"product",
		suite,
	)

	if err != nil {
		panic(err)
	}
}

func initCartClient(suite client.Option) {
	var err error
	CartClient, err = cartservice.NewClient(
		"cart",
		suite,
	)

	if err != nil {
		panic(err)
	}
}

func initOrderClient(suite client.Option) {
	var err error
	OrderClient, err = orderservice.NewClient(
		"order",
		suite,
	)

	if err != nil {
		panic(err)
	}
}

func initPaymentClient(suite client.Option) {
	var err error
	PaymentClient, err = paymentservice.NewClient(
		"payment",
		suite,
	)

	if err != nil {
		panic(err)
	}
}

func initCheckoutClient(suite client.Option) {
	var err error
	CheckoutClient, err = checkoutservice.NewClient(
		"checkout",
		suite,
	)

	if err != nil {
		panic(err)
	}
}