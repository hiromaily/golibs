#!/bin/bash

gofiles=$(find . -name "*.go" | grep -v "/vendor/")

for gofile in $gofiles; do
    echo $gofile
    sed '/^import/,/^[[:space:]]*)/ { /^[[:space:]]*$/ d; }' $gofile > tmp
    mv tmp $gofile
done

go fmt `go list ./... | grep -v "/vendor/"`
goimports -local github.com/hiromaily/golinbs/ -w `goimports -local github.com/hiromaily/golinbs/ -l ./ | grep -v "/vendor/"`
