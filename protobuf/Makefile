PWD=$(shell pwd)
PROTOC_BIN=docker run --rm -v $(PWD):$(PWD) -w $(PWD) znly/protoc:0.4.0
GOGO_MODULES=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types

.PHONY:vendor
vendor:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

.PHONY:build
build:
	$(PROTOC_BIN) --proto_path=proto -I=. \
		--gogoslick_out=$(GOGO_MODULES):./pb/ proto/**/*.proto
	chown -R $(whoami):$(whoami) ./pb

#.PHONY:java
#java:
#	$(PROTOC_BIN) --proto_path=proto -I=. \
#		--java_out=./pb/java/src/main/java proto/**/*.proto
#	chown -R $(whoami):$(whoami) ./pb/java

all: go java

clean:
	rm -rf pb/go/*