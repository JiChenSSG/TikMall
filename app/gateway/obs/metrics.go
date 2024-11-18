package obs

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/route"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/jichenssg/tikmall/app/gateway/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func initMetric() route.CtxCallback {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	consulEndpoint := fmt.Sprintf("%s:%d", config.GetConf().Consul.Host, config.GetConf().Consul.Port)
	consulConf := consulapi.DefaultConfig()
	consulConf.Address = consulEndpoint
	consulClient, _ := consulapi.NewClient(consulConf)
	r := consul.NewConsulRegister(consulClient)

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.GetConf().Server.Host, config.GetConf().Metrics.Port))
	if err != nil {
		hlog.Error(err)
	}

	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": "gateway"},
	}

	err = r.Register(registryInfo)
	if err != nil {
		hlog.Error(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(fmt.Sprintf(":%d", config.GetConf().Metrics.Port), nil) //nolint:errcheck
	return func(ctx context.Context) {
		r.Deregister(registryInfo) //nolint:errcheck
	}
}
