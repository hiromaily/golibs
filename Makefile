# Note: tabs by space can't not used for Makefile!
PROJECT_ROOT=${GOPATH}/src/github.com/hiromaily/golibs

JSONPATH=${PROJECT_ROOT}/testdata/json/teachers.json
TOMLPATH=${PROJECT_ROOT}/config/travis.toml
XMLPATH=${PROJECT_ROOT}/example/xml/rssfeeds/

KAFKA_IP=`docker ps -f name=lib-kafka1 --format "{{.Ports}}" | sed -e 's/0.0.0.0://g' | sed -e 's/->9092\/tcp//g'`

LOGLEVEL=1 #1:Debug, 2:Info, 3:Error, 4:Fatal, 5:No Log
LOG_ARG='-v'


###############################################################################
# Docker
###############################################################################
docker:
	sh ./docker-create.sh

up:
	docker-compose up

up_product:
	docker-compose -f docker-compose.yml up

dcbld:
	docker-compose build --no-cache

mysql:
	docker-compose up mysql

pg:
	docker-compose up pg


###############################################################################
# PKG Dependencies
###############################################################################
update:
	go get -u github.com/golang/dep/...
	go get -u github.com/rakyll/hey
	go get -u github.com/davecheney/httpstat
	go get -u github.com/client9/misspell/cmd/misspell
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/pilu/fresh
	go get -u github.com/tools/godep

	go get -u github.com/alecthomas/gometalinter
	#gometalinter --install

	go get -u -v ./...


###############################################################################
# Managing Dependencies
###############################################################################
godep:
	rm -rf Godeps
	rm -rf ./vendor
	godep save ./...


###############################################################################
# Golang formatter and detection
###############################################################################
fmt:
	go fmt `go list ./... | grep -v '/vendor/'`

vet:
	go vet `go list ./... | grep -v '/vendor/'`

fix:
	go fix `go list ./... | grep -v '/vendor/'`

lint:
	golint ./... | grep -v '^vendor\/' || true
	misspell `find . -name "*.go" | grep -v '/vendor/'`
	ineffassign .

chk:
	go fmt `go list ./... | grep -v '/vendor/'`
	go vet `go list ./... | grep -v '/vendor/'`
	go fix `go list ./... | grep -v '/vendor/'`
	golint ./... | grep -v '^vendor\/' || true
	misspell `find . -name "*.go" | grep -v '/vendor/'`
	ineffassign .


###############################################################################
# Install
###############################################################################
install:
	go install -v ./...


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
	#cd cast/;go test -bench=. -benchmem;cd ../;

	#cd db/mysql/;go test -bench=. -benchmem;cd ../;
	#cd db/redis/;go test -bench=. -benchmem;cd ../;
	#cd db/boltdb/;go test -bench=. -benchmem -fp ${BOLTPATH};cd ../;

	#cd example/flag/;go test -bench=. -benchmem -iv 1 -sv abcde;cd ../;
	#cd example/join/;go test -bench . -benchmem;cd ../;
	#cd example/join/;go test -bench=. -benchmem;cd ../;

	#cd files/;go test -bench=. -benchmem;
	#cd serial/;go test -bench . -benchmem;cd ../;


###############################################################################
# Clean
###############################################################################
cln:
	go clean -n

clnok:
	go clean

