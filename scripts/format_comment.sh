#!/bin/bash

gofiles=$(find . -name "*.go" | grep -v "/vendor/")

for gofile in $gofiles; do
    echo $gofile
    sed '/\/\/[^[:space:]]/s/\/\//\/\/ /g' $gofile > tmp
    mv tmp $gofile
done
