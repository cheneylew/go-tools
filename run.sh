#!/usr/bin/env bash

#export GOPATH=`pwd`/:`pwd`/vendor
export GOPATH=$GOPATH:`pwd`:`pwd`/vendor
echo "GOPATH: "${GOPATH}

./bee run

#go build  -o ./bin/bin github.com/cheneylew/go-tools/filetool
#./bin/bin

# ./bee run