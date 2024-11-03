#!/bin/bash

svcName=${1}

mkdir -p app/${svcName}
cd app/${svcName}
protoc client -I ../idl --type RPC --service ${svcName} --module github.com/jichenssg/tik-mall/app/${svcName} --idl ../../idl/${svcName}.proto
protoc server -I ../idl --type RPC --service ${svcName} --module github.com/jichenssg/tik-mall/app/${svcName} --idl ../../idl/${svcName}.proto