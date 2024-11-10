package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/jichenssg/tikmall/app/auth/config"
	"github.com/jichenssg/tikmall/app/common/obs"
	auth "github.com/jichenssg/tikmall/gen/kitex_gen/auth/authservice"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/natefinch/lumberjack"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	config.GetConf()

	obs.InitLog(&lumberjack.Logger{
		Filename:   config.GetConf().Kitex.LogFileName,
		MaxSize:    config.GetConf().Kitex.LogMaxSize,
		MaxBackups: config.GetConf().Kitex.LogMaxBackups,
		MaxAge:     config.GetConf().Kitex.LogMaxAge,
	}, config.LogLevel())

	obs.InitTrack(config.GetConf().Server.Name, fmt.Sprintf("%v:%v", config.GetConf().Telemetry.Host, config.GetConf().Telemetry.Port))

	kitexInit()
}

func kitexInit() {
	r, err := consul.NewConsulRegister(fmt.Sprintf("%v:%v", config.GetConf().Consul.Host, config.GetConf().Consul.Port))

	if err != nil {
		log.Println(err.Error())
		return
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", config.GetConf().Server.Host, config.GetConf().Server.Port))
	if err != nil {
		log.Println(err.Error())
		return
	}

	svr := auth.NewServer(
		new(AuthServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "auth",
		}),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)



	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
