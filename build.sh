#!/usr/bin/env bash

#export GOTRACEBACK=single
CURRENTDIR=`pwd`


#git checkout v0.9.17
#cd ${GOPATH}/src/github.com/aws/aws-sdk-go

#go get -u -v ./...
go fmt ./...
go vet ./...
#go vet `go list ./... | grep -v '/vendor/'`


#build
go build -o ./dblab/handledb ./dblab/handledb.go


#run
#./db/handledb -mode 3 -toml ${CURRENTDIR}/settings.toml
