package main

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/jichenssg/tikmall/app/common/client"
	"github.com/jichenssg/tikmall/app/common/dal/mysql"
	"github.com/jichenssg/tikmall/app/common/obs"
	"github.com/kanhai-syd/hailog/logging"
	"github.com/natefinch/lumberjack"

	"github.com/jichenssg/tikmall/app/cart/config"
	"github.com/jichenssg/tikmall/app/cart/dal/model"
	serversuite "github.com/jichenssg/tikmall/app/common/serversuite"
	cart "github.com/jichenssg/tikmall/gen/kitex_gen/cart/cartservice"
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

	mysql.Init(
		config.GetConf().Mysql.Dsn,
		config.GetConf().Mysql.User,
		config.GetConf().Mysql.Password,
		config.GetConf().Mysql.Host,
		config.GetConf().Mysql.Port,
		config.GetConf().Mysql.Database,
	)
	mysql.GetDB().AutoMigrate(model.Cart{})

	client.Init(
		config.GetConf().Server.Name,
		fmt.Sprintf("%v:%v", config.GetConf().Consul.Host, config.GetConf().Consul.Port),
		fmt.Sprintf("%v:%v", config.GetConf().Telemetry.Host, config.GetConf().Telemetry.Port),
	)

	klog.Infof("Starting service %v", config.GetConf().Server.Name)

	kitexInit()
}

func kitexInit() {
	svr := cart.NewServer(
		new(CartServiceImpl),
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
		logging.Fatal(err)
	}
}