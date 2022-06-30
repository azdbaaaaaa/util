login:
	buf register login

lint:
	buf lint

format:
	buf format -w

breaking:format
	# "proto" a nested directory in util
	buf breaking proto --against "https://github.com/azdbaaaaaa/util.git#branch=master"

generate:breaking
	cd proto && buf mod update && cd ..
	buf generate
	sed -i.bak "s/google.golang.org\/protobuf\/types\/known\/anypb/github.com\/gogo\/protobuf\/types/g" proto/common/response.pb.go && rm proto/common/response.pb.go.bak

push:generate
	cd proto && buf push && cd ..