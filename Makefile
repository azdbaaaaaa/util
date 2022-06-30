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
	sh proto/gen-proto.sh

push:generate
	cd proto && buf push && cd ..