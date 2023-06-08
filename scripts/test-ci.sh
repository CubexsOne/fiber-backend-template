#!/bin/sh
apk add build-base

DIR=$(dirname $(readlink -f $0))
cd /app

go install gotest.tools/gotestsum@latest
go mod download

gotestsum --junitfile $DIR/../junit.xml ./src/...
