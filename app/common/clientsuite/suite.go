package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct {
	ServiceName      string
	RegistryEndpoint string
	OtelEndpoint     string
}

func (s CommonClientSuite) Options() (opts []client.Option) {
	r, err := consul.NewConsulResolver(s.RegistryEndpoint)
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))

	opts = append(opts, 
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	)

	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.ServiceName,
		}),
		client.WithSuite(tracing.NewClientSuite()),
	)

	return opts
}
