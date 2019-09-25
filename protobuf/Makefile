PWD=$(shell pwd)
UID=$(shell id -u)
#UNAME=$(shell id -un)
UNAME=$(shell whoami)
GNAME=$(shell id -gn)

# All inclusive Protocol Buffer and gRPC suite, but it's not maintained much
# https://github.com/znly/docker-protobuf
#PROTOC_BIN=docker run --rm -v $(PWD):$(PWD) -w $(PWD) znly/protoc:0.4.0
PROTOC_BIN=protoc

GOGO_MODULES=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types

# it doesn't work as expected
.PHONY:sleep-tech
sleep-tech:
	echo start
	@$(shell sleep 3)
	echo end

.PHONY:init
init: update
	brew install protobuf

.PHONY:update
update:
    go get -u github.com/gogo/protobuf
	# refer to https://developers.google.com/protocol-buffers/docs/reference/go-generated
	# provides a protoc-gen-go binary which protoc uses when invoked with the --go_out command-line flag.
	# The --go_out flag
	go get -u github.com/golang/protobuf/protoc-gen-go
	# The --gofast_out flag
	go get -u github.com/gogo/protobuf/protoc-gen-gofast
	# The --gogoslick_out flag
	go get -u github.com/gogo/protobuf/protoc-gen-gogoslick

.PHONY:vendor
vendor:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

.PHONY:build
build:
	# e.g. protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
	$(PROTOC_BIN) --proto_path=proto -I=. \
		--gogoslick_out=$(GOGO_MODULES):./pb/ proto/**/*.proto
	chown -R $(UNAME):$(GNAME) ./pb

.PHONY:grpc
grpc:
	# e.g. protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
	$(PROTOC_BIN) --proto_path=proto -I=. \
		--gogoslick_out=plugins=grpc,$(GOGO_MODULES):./pb/ proto/**/*.proto
	chown -R $(UNAME):$(GNAME) ./pb

#.PHONY:java
#java:
#	$(PROTOC_BIN) --proto_path=proto -I=. \
#		--java_out=./pb/java/src/main/java proto/**/*.proto
#	chown -R $(whoami):$(whoami) ./pb/java

clean:
	rm -rf pb/go/*