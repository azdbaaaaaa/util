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
proto_path=${proto_path}:${GOPATH}/src/github.com/azdbaaaaaa/util/


cd "${home}"/proto/common
protoc --gofast_out=plugins=grpc:. --proto_path=${proto_path} --validate_out="lang=go:." common.proto
#mv "${home}"/proto/common/github.com/azdbaaaaaa/util/proto/common/* "${home}"/proto/common
#rm -r "${home}"/proto/common/github.com/

cd "${home}"/proto/pagination
protoc --gofast_out=plugins=grpc:. --proto_path=${proto_path} --validate_out="lang=go:." pagination.proto
#mv "${home}"/proto/pagination/github.com/azdbaaaaaa/util/proto/pagination/* "${home}"/proto/pagination
#rm -r "${home}"/proto/pagination/github.com/
