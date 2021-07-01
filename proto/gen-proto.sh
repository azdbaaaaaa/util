#!/usr/bin/env bash
set -e
# shellcheck disable=SC2006
home=`pwd`
cd "$home"
#protoc --gofast_out=plugins=grpc:. *.proto
#protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogofast_out=. *.proto
#cd $home && cd src/proto
#protoc --gofast_out=plugins=grpc:. *.proto
proto_path=.:${GOPATH}/src/
proto_path=${proto_path}:${GOPATH}/src/github.com/gogo/protobuf/protobuf/
proto_path=${proto_path}:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate/
proto_path=${proto_path}:${GOPATH}/src/util/


cd "${home}"/proto
protoc --gofast_out=plugins=grpc:. --proto_path=${proto_path} --validate_out="lang=go:." common.proto
protoc --gofast_out=plugins=grpc:. --proto_path=${proto_path} --validate_out="lang=go:." pagination.proto