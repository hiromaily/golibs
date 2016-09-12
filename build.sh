#!/bin/sh

###########################################################
# Variable
###########################################################
#export GOTRACEBACK=single
#export GOTRACEBACK=all
#CURRENTDIR=`pwd`

JSONPATH=${GOPATH}/src/github.com/hiromaily/go-book-teacher/json/teachers.json
TOMLPATH=${GOPATH}/src/github.com/hiromaily/golibs/settings.toml
XMLPATH=${GOPATH}/src/github.com/hiromaily/golibs/example/xml/rssfeeds/
BOLTPATH=${GOPATH}/src/github.com/hiromaily/golibs/db/boltdb/boltdb

TEST_MODE=2  #0:off, 1:run all test, 2:test for specific one
BENCH=0
COVERAGRE=0
PROFILE=0

LOGLEVEL=1 #0: don't show t.Log() and log level is over or equal to INFO
           #1: show t.Log() and log level is DEBUG

GO_GET=0
GO_LINT=1

# when using go 1.7 for the first time, delete all inside pkg directory and run go install.
#go install -v ./...

###########################################################
# Update all package
###########################################################
if [ $GO_GET -eq 1 ]; then
    #go get -u -v
    go get -u -v ./...
    #go get -u -f -v ./...
fi

###########################################################
# Adjust version dependency of projects
###########################################################
#cd ${GOPATH}/src/github.com/aws/aws-sdk-go
#git checkout v0.9.17
#git checkout master


###########################################################
# go fmt and go vet
###########################################################
echo '============== go fmt; go vet; =============='
go fmt ./...
go vet ./...
EXIT_STATUS=$?
if [ $EXIT_STATUS -gt 0 ]; then
    exit $EXIT_STATUS
fi

# when there is vendor directory under project for management package dependency
#go vet `go list ./... | grep -v '/vendor/'`


###########################################################
# go lint
###########################################################
#go get -u github.com/golang/lint/golint
if [ $GO_LINT -eq 1 ]; then
    echo '============== golint; =============='
    #golint ./...
    #golint `go list ./... | grep -v '/vendor/'`
    golint ./... | grep -v '^vendor\/' || true
fi


###########################################################
# go install
###########################################################
echo '============== go install; =============='
#go install -a -v ./...
go install -v ./...
EXIT_STATUS=$?

if [ $EXIT_STATUS -gt 0 ]; then
    exit $EXIT_STATUS
fi


###########################################################
# go test
###########################################################
if [ $TEST_MODE -eq 1 ]; then
    #t.LogのON/OFF
    echo '============== test =============='
    #go test -v html/html_test.go

    #Check OK
    #
    go test -v auth/jwt/jwt_test.go -log ${LOGLEVEL}
    go test -v cipher/encryption/encryption_test.go -log ${LOGLEVEL}
    go test -v cipher/hash/hash_test.go -log ${LOGLEVEL}
    go test -v compress/compress_test.go -log ${LOGLEVEL}
    go test -v config/config_test.go -fp ${TOMLPATH} -log ${LOGLEVEL}

    #db
    go test -v db/boltdb/boltdb_test.go -fp ${BOLTPATH} -log ${LOGLEVEL}
    go test -v db/cassandra/cassandra_test.go -log ${LOGLEVEL}
    go test -v db/gorm/gorm_test.go -log ${LOGLEVEL}
    go test -v db/gorp/gorp_test.go -log ${LOGLEVEL}
    go test -v db/mongodb/mongodb_test.go -fp ${JSONPATH} -log ${LOGLEVEL}
    go test -v db/mysql/mysql_test.go -log ${LOGLEVEL}
    go test -v db/redis/redis_test.go -log ${LOGLEVEL}


    #example
    go test -v example/defaultdata/defaultdata_test.go
    go test -v example/exec/exec_test.go
    go test -v example/flag/flag_test.go -log ${LOGLEVEL} -iv 1 -sv abcde
    go test -v example/http/http_test.go -log ${LOGLEVEL}
    go test -v example/json/json_test.go -fp ${JSONPATH} -log ${LOGLEVEL}
    go test -v example/xml/xml_test.go -fp ./rssfeeds/techcrunch.xml -log ${LOGLEVEL}

    #
    go test -v -race files/files_test.go -log ${LOGLEVEL}
    go test -v -race goroutine/goroutine_test.go -log ${LOGLEVEL}
    go test -v heroku/heroku_test.go -log ${LOGLEVEL}
    go test -v log/log_test.go -log ${LOGLEVEL}
    go test -v mail/mail_test.go -log ${LOGLEVEL} -fp ${TOMLPATH}

    # messaging
    go test -v messaging/kafka/kafka_test.go -log ${LOGLEVEL}
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

elif [ $TEST_MODE -eq 2 ]; then
    go test -v utils/utils_test.go -log ${LOGLEVEL}

fi

###########################################################
# go test benchmark
###########################################################
if [ $BENCH -eq 1 ]; then
    echo '============== benchmark =============='
    #OK
    cd serial/;go test -bench . -benchmem;cd ../;


    #cd cast/;go test -bench=. -benchmem;cd ../;

    cd flag/;go test -bench=. -benchmem -iv 1 -sv abcde;cd ../;

    #cd files/;go test -bench=. -benchmem;

    #cd join/;go test -bench . -benchmem;cd ../;
    #cd join/;go test -bench=. -benchmem;cd ../;

    #cd serial/;go test -bench . -benchmem;cd ../;
    #cd serial/;go test -bench=. -benchmem;cd ../;

    #cd db/mysql/;go test -bench=. -benchmem;cd ../;

    #cd db/redis/;go test -bench=. -benchmem;cd ../;

    #cd db/boltdb/;go test -bench=. -benchmem -fp ${BOLTPATH};cd ../;
fi

###########################################################
# go coverage
###########################################################
if [ $COVERAGRE -eq 1 ]; then
    echo '============== coverage =============='

    #how to use it
    #go tool cover

    #it doesn't work below
    #go test -coverprofile=cover.out -v cipher/hash/hash_test.go
    #instead of it, exec below
    #cd cipher/hash/;go test -coverprofile=cover.out;cd ../../;

    #check result on the web
    #go tool cover -html=cipher/hash/cover.out
fi

###########################################################
# go profile
###########################################################
if [ $PROFILE -eq 1 ]; then
    echo '============== profile =============='

    #serial
    #cd serial/;go test -run=NONE -bench=BenchmarkSerializeStruct -cpuprofile=cpu.log .;cd ../;
    #cd serial/;go tool pprof -text -nodecount=10 ./serial.test cpu.log;
fi


###########################################################
# cross-compile for linux
###########################################################
#GOOS=linux go install -v ./...


###########################################################
# godoc
###########################################################
#godoc -http :8000
#http://localhost:8000/pkg/
