#!/usr/bin/env bash

export GOPATH=$GOPATH:`pwd`:`pwd`/vendor
#export GOPATH=`pwd`/:`pwd`/vendor
echo "GOPATH: "${GOPATH}

go build  -o ./bin/bin github.com/cheneylew/go-tools/filetool
./bin/bin

# ./bee run