package server

import (
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/jichenssg/tikmall/app/common/obs"
	"github.com/kanhai-syd/hailog/logging"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	registryconsul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	ServiceName      string
	Host             string
	Port             int
	RegistryEndpoint string
	OtelEndpoint     string
}

func (s CommonServerSuite) Options() (opts []server.Option) {
	r, err := registryconsul.NewConsulRegister(s.RegistryEndpoint)
	if err != nil {
		logging.Fatal(err)
	}

	opts = append(opts, server.WithRegistry(r))

	opts = append(opts, server.WithMetaHandler(transmeta.ServerHTTP2Handler))

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", s.Host, s.Port))
	if err != nil {
		logging.Fatal(err)
		return
	}

	_ = provider.NewOpenTelemetryProvider(
		provider.WithSdkTracerProvider(obs.GetTraceProvider()),
		provider.WithEnableMetrics(false),
		provider.WithExportEndpoint(s.OtelEndpoint),
	)

	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.ServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServiceAddr(addr),
		server.WithTracer(
			prometheus.NewServerTracer(
				"", "",
				prometheus.WithDisableServer(true),
				prometheus.WithRegistry(obs.Registry),
			),
		),
	)

	return opts
}
