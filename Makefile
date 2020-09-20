# Note: tabs by space can't not used for Makefile!
PROJECT_ROOT=${GOPATH}/src/github.com/hiromaily/golibs

JSONPATH=${PROJECT_ROOT}/testdata/json/teachers.json
TOMLPATH=${PROJECT_ROOT}/config/travis.toml
XMLPATH=${PROJECT_ROOT}/example/xml/rssfeeds/

KAFKA_IP=`docker ps -f name=lib-kafka1 --format "{{.Ports}}" | sed -e 's/0.0.0.0://g' | sed -e 's/->9092\/tcp//g'`

# These variables can be overriden `make LOGLEVEL=2 target`
LOGLEVEL ?= 1 #1:Debug, 2:Info, 3:Error, 4:Fatal, 5:No Log

LOG_ARG='-v'
#if [ $LOGLEVEL -eq 5 ]; then
#    LOG_ARG=''
#fi

modVer=$(shell cat go.mod | head -n 3 | tail -n 1 | awk '{print $2}' | cut -d'.' -f2)
currentVer=$(shell go version | awk '{print $3}' | sed -e "s/go//" | cut -d'.' -f2)

###############################################################################
# Managing Dependencies
###############################################################################
.PHONY: check-ver
check-ver:
	#echo $(modVer)
	#echo $(currentVer)
	@if [ ${currentVer} -lt ${modVer} ]; then\
		echo go version ${modVer}++ is required but your go version is ${currentVer};\
	fi


GOLINT = $(GOPATH)/bin/golangci-lint

$(GOLINT):
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: update
update:
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	GO111MODULE=off go get -u github.com/rakyll/hey
	GO111MODULE=off go get -u github.com/davecheney/httpstat
	GO111MODULE=off go get -u github.com/pilu/fresh
	go get -u -d -v ./...


.PHONY: mod-private-git
mod-private-git:
	git config --global --add url."git@github.yourhost.com:".insteadOf "https://github.yourhost.com/"

.PHONY: mod-vendor
mod-vendor:
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor


###############################################################################
# Golang formatter and detection
###############################################################################
.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: imports
imports:
	./scripts/imports.sh

.PHONY: imports-settings
imports-settings:
	GO111MODULE=off go get -u github.com/pwaller/goimports-update-ignore
	goimports-update-ignore -max-depth 5


###############################################################################
# Clean
###############################################################################
.PHONY: clean
clean:
	go clean -n

.PHONY: cleanok
cleanok:
	go clean


###############################################################################
# Docker
###############################################################################
.PHONY: dc-create
dc-create:
	sh ./scripts/create-containers

.PHONY: dc-up
dc-up:
	docker-compose up

.PHONY: dc-up-product
dc-up-product:
	docker-compose -f docker-compose.yml up

.PHONY: dc-bld
dc-bld:
	docker-compose build --no-cache

.PHONY: dc-mysql
dc-mysql:
	docker-compose up mysql

.PHONY: dc-pg
dc-pg:
	docker-compose up pg


###############################################################################
# Test
###############################################################################
solo:
	go test -v web/context/context_test.go -log 1

util:
	go test -v utils/utils_test.go -log 1

strings:
	go test -v example/strings/strings_test.go -log 1

fil:
	go test -v files/files_test.go -log 1

tim:
	go test -v time/time_test.go -log 1

texec:
	go test -v example/exec/exec_test.go -run TestExecParams -log 1

enc:
	go test -v cipher/encryption/encryption_test.go -log 1

runtime:
	go test -v runtimes/runtimes_test.go -log 1

json:
	go test -v example/json/json_test.go -run TestLoadWithDecode -jfp ${JSONPATH} -log 1

regex:
	go test -v regexp/regexp_test.go -run TestReplace -log 1

goflag:
	go run ./example/go-flag/cmd/main.go cmd1 -i -path abc

cli:
	go build -o cli ./example/cli/cmd/...; ./cli add -debug; rm ./cli
	#go build -o cli ./main.go; ./cli nest other -debug; rm ./cli

test:
	#
	go test -v auth/jwt/jwt_test.go -log ${LOGLEVEL}
	go test -v cipher/encryption/encryption_test.go -log ${LOGLEVEL}
	go test -v cipher/hash/hash_test.go -log ${LOGLEVEL}
	go test -v compress/compress_test.go -log ${LOGLEVEL}
	go test -v config/config_test.go -fp ${TOMLPATH} -log ${LOGLEVEL}

	#db
	go test -v db/boltdb/boltdb_test.go -log ${LOGLEVEL}
	go test -v db/cassandra/cassandra_test.go -log ${LOGLEVEL}
	go test -v db/gorm/gorm_test.go -log ${LOGLEVEL}
	go test -v db/gorp/gorp_test.go -log ${LOGLEVEL}
	go test -v db/mongodb/mongodb_test.go -jfp ${JSONPATH} -log ${LOGLEVEL}
	go test -v db/mysql/mysql_test.go -log ${LOGLEVEL}
	go test -v db/redis/redis_test.go -log ${LOGLEVEL}


	#example
	go test -v example/defaultdata/defaultdata_test.go
	go test -v example/exec/exec_test.go
	go test -v example/flag/flag_test.go -log ${LOGLEVEL} -iv 1 -sv abcde
	go test -v example/http/http_test.go -log ${LOGLEVEL}
	go test -v example/json/json_test.go -jfp ${JSONPATH} -log ${LOGLEVEL}
	go test -v example/xml/xml_test.go -fp ./rssfeeds/techcrunch.xml -log ${LOGLEVEL}

	#
	go test -v -race files/files_test.go -log ${LOGLEVEL}
	go test -v -race goroutine/goroutine_test.go -log ${LOGLEVEL}
	go test -v heroku/heroku_test.go -log ${LOGLEVEL}
	go test -v log/log_test.go -log ${LOGLEVEL}
	go test -v mail/mail_test.go -log ${LOGLEVEL} -fp ${TOMLPATH}

	# messaging
	go test -v messaging/kafka/kafka_test.go -kip ${KAFKA_IP} -log ${LOGLEVEL}
	go test -v messaging/nats/nats_test.go -log ${LOGLEVEL}
	go test -v messaging/rabbitmq/rmq_test.go -log ${LOGLEVEL}

	#
	go test -v os/os_test.go -log ${LOGLEVEL}
	go test -v reflects/reflects_test.go -log ${LOGLEVEL}
	go test -v regexp/regexp_test.go -log ${LOGLEVEL}
	go test -v runtimes/runtimes_test.go -log ${LOGLEVEL}
	go test -v serial/serial_test.go -log ${LOGLEVEL}
	#GOTRACEBACK=all go test -v signal/signal_test.go -log ${LOGLEVEL}
	go test -v testutil/testutil_test.go -log ${LOGLEVEL}
	go test -v time/time_test.go -log ${LOGLEVEL}
	go test -v tmpl/tmpl_test.go -log ${LOGLEVEL}
	go test -v utils/utils_test.go -log ${LOGLEVEL}
	go test -v validator/validator_test.go -log ${LOGLEVEL}
	go test -v web/context/context_test.go -log ${LOGLEVEL}
	go test -v web/session/session_test.go -log ${LOGLEVEL}

	go test -v yaml/yaml_test.go -run TestYAMLTable -log 1
	go test -v i18n/i18n_test.go -log 1


bench:
	# for all
	#go test ./... -bench=. -benchmem

	#cd cast/;go test -bench=. -benchmem;cd ../;

	#cd db/mysql/;go test -bench=. -benchmem;cd ../;
	#cd db/redis/;go test -bench=. -benchmem;cd ../;
	#cd db/boltdb/;go test -bench=. -benchmem -fp ${BOLTPATH};cd ../;

	#cd example/flag/;go test -bench=. -benchmem -iv 1 -sv abcde;cd ../;
	#cd example/join/;go test -bench . -benchmem;cd ../;
	#cd example/join/;go test -bench=. -benchmem;cd ../;

	#cd files/;go test -bench=. -benchmem;
	#cd serial/;go test -bench . -benchmem;cd ../;


