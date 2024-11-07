package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	auth "github.com/jichenssg/tikmall/gen/kitex_gen/auth/authservice"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulRegister("127.0.0.1:8500")

	if err != nil {
		log.Println(err.Error())
		return
	}

	addr, err := net.ResolveTCPAddr("tcp", "172.16.53.189:8888")
	if err != nil {
		log.Println(err.Error())
		return
	}

	svr := auth.NewServer(new(AuthServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "auth",
	}), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
