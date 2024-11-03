#!/bin/bash

svcName=${1}

mkdir -p rpc_gen && cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/jichenssg/tikmall/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

cd ..

mkdir -p app/${svc} && cd app/${svc} && go mod init github.com/jichenssg/tikmall/app/${svc} &&cwgo server --type RPC --service ${svc} --module github.com/jichenssg/tikmall/app/${svc} --pass "-use github.com/jichenssg/tikmall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto