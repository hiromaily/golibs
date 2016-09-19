#!/bin/sh

###############################################################################
# ENVIRONMENT VARIABLE
###############################################################################
PROJECT_ROOT=${GOPATH}/src/github.com/hiromaily/golibs

TOMLPATH=${PROJECT_ROOT}/config/travis.toml
JSONPATH=${PROJECT_ROOT}/testdata/json/teachers.json


###############################################################################
# DOCKER
###############################################################################
# GET Kafka port
#docker ps -f name=lib-kafka1 --format "{{.Ports}}" | sed -e 's/[0-9]\{5\}//g'
KAFKA_IP=`docker ps -f name=lib-kafka1 --format "{{.Ports}}" | sed -e 's/0.0.0.0://g' | sed -e 's/->9092\/tcp//g'`
#echo $KAFKA_IP



###############################################################################
# FUNC
###############################################################################
set -e
# when breaking, it cleans

cleanup() {
  retval=$?
  if [ $tmpprof != "" ] && [ -f $tmpprof ]; then
    rm -f $tmpprof
  fi
  exit $retval
}

trap cleanup INT QUIT TERM EXIT


###############################################################################
# MAIN
###############################################################################
prof=${1:-"profile.cov"}
echo "mode: count" > $prof
gopath1=$(echo $GOPATH | cut -d: -f1)
#go list ./... | grep -v '/example\|workinprogress\|web/'
for pkg in $(go list ./... | grep -v '/example\|workinprogress\|web\|signal\|kafka'); do
#for pkg in $(go list ./... | grep -v '/example\|workinprogress\|web\|signal'); do
  tmpprof=$gopath1/src/$pkg/profile.tmp
  go test -covermode=count -coverprofile=$tmpprof $pkg -log 5 -fp ${TOMLPATH} -jfp ${JSONPATH} -kip ${KAFKA_IP}
  if [ -f $tmpprof ]; then
    cat $tmpprof | tail -n +2 >> $prof
    rm $tmpprof
  fi
done

