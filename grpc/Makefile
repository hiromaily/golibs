
.PHONY:update
update:
	go get -u google.golang.org/grpc

.PHONY:build
build:
	go build -i -v -o ${GOPATH}/bin/grpc-server ./server/
	go build -i -v -o ${GOPATH}/bin/grpc-client ./client/

.PHONY:run
run:
	grpc-server
	grpc-client -mode 1 -name Mike -q 1
