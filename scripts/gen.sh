#!/bin/bash

isFrontend=$2

# protoc -I ./idl \
#     --go_out ./gen/${svc} --go_opt paths=source_relative \
#     --go-grpc_out ./gen/${svc} --go-grpc_opt paths=source_relative \
# 	--grpc-gateway_out=./gen/${svc} --grpc-gateway_opt=paths=source_relative \
#     ./idl/${svc}.proto

if [ "$isFrontend" = 0 ]; then
	cd gen
	kitex -I ./../idl -module "github.com/jichenssg/tikmall/gen" ./../idl/${svc}.proto
	cd ..

	mkdir -p app/${svc}
	cd app/${svc}
	kitex -I ./../../idl -module "github.com/jichenssg/tikmall/app/auth" -service ${svc} -use "github.com/jichenssg/tikmall/gen/kitex_gen" ./../../idl/${svc}.proto
else
	cd gateway
	# check if frontend folder exists
	if [ ! -f ".hz" ]; then
		hz new -I ./../idl -module "github.com/jichenssg/tikmall/gateway" -idl ./../idl/frontend/${svc}_f.proto
	fi
		hz update -I ./../idl -idl ./../idl/frontend/${svc}_f.proto

fi



