package client

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/jichenssg/tikmall/app/common/clientsuite"
	"github.com/jichenssg/tikmall/gen/kitex_gen/auth/authservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/product/productcatalogservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/user/userservice"
)

var (
	AuthClient    authservice.Client
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
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
