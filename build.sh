#!/usr/bin/env bash

export GOPATH=/mnt/f/workspace/godev

cd proto

# 遍历每个文件夹
for file in `ls`
do
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. "$file/$file.proto"
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. "$file/$file.proto"
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. "$file/$file.proto"
done

cd ..

go build