
.PHONY:update
update:
	go get -u google.golang.org/grpc

.PHONY:build
build:
	go build -i -v -o ${GOPATH}/bin/grpc-server ./server/
	go build -i -v -o ${GOPATH}/bin/grpc-client ./client/

.PHONY:server
server:
	grpc-server

.PHONY:client
client:
	grpc-client -mode 1 -name Mike -q 1
	grpc-client -mode 1 -name Mike -q 10
	grpc-client -mode 1 -name Mike -q 100

	grpc-client -mode 2
	grpc-client -mode 3 -name Mike -q 1
	grpc-client -mode 4 -name Mike -q 1

# Evans
.PHONY:setup-evans
setup-evans:
	brew tap ktr0731/evans
	brew install evans

# TLS mode should be disabled
.PHONY:connect-server
connect-server:
    evans ../protobuf/proto/sample/sample.proto
