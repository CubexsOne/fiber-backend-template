#!/bin/sh

ENV=$1

cd /app
env GO111MODULE=on go install github.com/cortesi/modd/cmd/modd

go mod download && modd --file=modd-$ENV.conf