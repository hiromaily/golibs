#!/bin/sh

###########################################################
# Variable
###########################################################
#export GOTRACEBACK=single
#export GOTRACEBACK=all
#CURRENTDIR=`pwd`

JSONPATH=${GOPATH}/src/github.com/hiromaily/booking-teacher/settings.json
TOMLPATH=${GOPATH}/src/github.com/hiromaily/golibs/settings.toml
BOLTPATH=${GOPATH}/src/github.com/hiromaily/golibs/boltdb

TEST_MODE=0  #0:off, 1:run all test, 2:test for specific one
BENCH=0
COVERAGRE=0
PROFILE=0

###########################################################
# Update all package
###########################################################
#go get -u -v
#go get -u -v ./...
#go get -u -f -v ./...


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
# it's too strict
#golint ./...


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
    echo '============== test =============='
    go test -v cipher/encryption/encryption_test.go
    go test -v cipher/hash/hash_test.go
    go test -v compress/compress_test.go
    go test -v config/config_test.go -fp ${TOMLPATH}

    go test -v db/boltdb/boltdb_test.go -fp ${BOLTPATH}
    go test -v db/cassandra/cassandra_test.go
    go test -v db/gorm/gorm_test.go
    go test -v db/gorp/gorp_test.go
    go test -v db/mongodb/mongodb_test.go -fp ${JSONPATH}
    go test -v db/mysql/mysql_test.go
    go test -v db/redis/redis_test.go

    go test -v defaultdata/defaultdata_test.go
    go test -v draw/draw_test.go
    go test -v exec/exec_test.go
    go test -v -race files/files_test.go
    go test -v flag/flag_test.go -iv 1 -sv abcde
    go test -v -race goroutine/goroutine_test.go
    go test -v heroku/heroku_test.go
    go test -v http/http_test.go
    go test -v html/html_test.go
    go test -v json/json_test.go -fp ${JSONPATH}
    go test -v mails/mails_test.go -fp ${TOMLPATH}
    go test -v os/os_test.go
    go test -v reflects/reflects_test.go
    go test -v regexp/regexp_test.go
    go test -v runtimes/runtimes_test.go
    go test -v serial/serial_test.go
    go test -v times/times_test.go
    go test -v tmpl/tmpl_test.go
    go test -v validator/validator_test.go
elif [ $TEST_MODE -eq 2 ]; then
    go test -v html/html_test.go
fi

###########################################################
# go test benchmark
###########################################################
if [ $BENCH -eq 1 ]; then
    echo '============== benchmark =============='

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
