package main

import (
	"fmt"
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/jichenssg/tikmall/app/auth/config"
	"github.com/jichenssg/tikmall/app/common/obs"
	auth "github.com/jichenssg/tikmall/gen/kitex_gen/auth/authservice"
	"github.com/natefinch/lumberjack"

	serversuite "github.com/jichenssg/tikmall/app/common/serversuite"
)

func main() {
	config.GetConf()

	obs.InitTracer(config.GetConf().Server.Name)

	obs.InitMetrics(
		config.GetConf().Server.Name,
		config.GetConf().Server.Host,
		config.GetConf().Metrics.Port,
		fmt.Sprintf("%v:%v", config.GetConf().Consul.Host, config.GetConf().Consul.Port),
	)

	obs.InitLog(&lumberjack.Logger{
		Filename:   config.GetConf().Kitex.LogFileName,
		MaxSize:    config.GetConf().Kitex.LogMaxSize,
		MaxBackups: config.GetConf().Kitex.LogMaxBackups,
		MaxAge:     config.GetConf().Kitex.LogMaxAge,
	}, config.LogLevel())

	klog.Infof("Starting service %v", config.GetConf().Server.Name)

	kitexInit()
}

func kitexInit() {
	svr := auth.NewServer(
		new(AuthServiceImpl),
		server.WithSuite(
			serversuite.CommonServerSuite{
				ServiceName:      config.GetConf().Server.Name,
				Host:             config.GetConf().Server.Host,
				Port:             config.GetConf().Server.Port,
				RegistryEndpoint: fmt.Sprintf("%v:%v", config.GetConf().Consul.Host, config.GetConf().Consul.Port),
				OtelEndpoint:     fmt.Sprintf("%v:%v", config.GetConf().Telemetry.Host, config.GetConf().Telemetry.Port),
			},
		),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
