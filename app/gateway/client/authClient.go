package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jichenssg/tikmall/gateway/config"
	"github.com/jichenssg/tikmall/gen/kitex_gen/auth/authservice"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

var (
	authClient     authservice.Client
	authClientOnce sync.Once
)

func GetAuthClient() (authservice.Client, error) {
	authClientOnce.Do(func() {
		resolver, err := consul.NewConsulResolver(fmt.Sprintf("%s:%d", config.GetConf().Consul.Host, config.GetConf().Consul.Port))
		if err != nil {
			klog.CtxFatalf(context.Background(), "fail to create consul resolver %v", err)
		}

		authClient, err = authservice.NewClient(
			config.GetConf().Endpoint.Auth,
			client.WithResolver(resolver),
			client.WithSuite(tracing.NewClientSuite()),
		)
		if err != nil {
			klog.CtxFatalf(context.Background(), "fail to create auth client %v", err)
		}
	})

	return authClient, nil
}
