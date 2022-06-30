generate:
	cd proto && buf mod update && cd ..
	buf generate
	sed -i.bak "s/google.golang.org\/protobuf\/types\/known\/anypb/github.com\/gogo\/protobuf\/types/g" proto/common/response.pb.go && rm proto/common/response.pb.go.bak

push:
	cd proto && buf push && cd ..