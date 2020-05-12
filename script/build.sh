#########################################################################
# File Name: build.sh
# Description: build.sh
# Author: zhangyi
# mail: 450575982@qq.com
# Created Time: 2020-05-08 17:12:40
#########################################################################
#!/bin/bash
export GO111MODULE=on
protoc --go_out=plugins=grpc:. ./pb/*.proto

rm -rf client.x server.x

if [[ $# == 0 ]];
then
    target="all"
else
    target=${1}
fi

if [[ ${target} == "all" || ${target} == "server" ]];
then
    go build -o server.x server/*.go
fi

if [[ ${target} == "all" || ${target} == "client" ]];
then
    go build -o client.x client/*.go
fi

