package client

import (
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/jichenssg/tikmall/app/common/clientsuite"
	"github.com/jichenssg/tikmall/gateway/config"
	"github.com/jichenssg/tikmall/gen/kitex_gen/auth/authservice"
	"github.com/jichenssg/tikmall/gen/kitex_gen/user/userservice"
)

var (
	AuthClient authservice.Client
	UserClient userservice.Client
	Once       sync.Once
)

func Init() {
	Once.Do(func() {
		commonClientSuite := client.WithSuite(clientsuite.CommonClientSuite{
			ServiceName:      config.GetConf().Server.Name,
			RegistryEndpoint: fmt.Sprintf("%s:%d", config.GetConf().Consul.Host, config.GetConf().Consul.Port),
			OtelEndpoint:     fmt.Sprintf("%s:%d", config.GetConf().Telemetry.Host, config.GetConf().Telemetry.Port),
		})

		initAuthClient(commonClientSuite)
		initUserClient(commonClientSuite)

	})
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
