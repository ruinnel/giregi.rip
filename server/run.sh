#!/bin/bash

PLATFORM=$1
VERSION=$2

run() {
    go run -ldflags "-X main.platform=${PLATFORM}" app/main.go
}

case $1 in
    run) run;;
    *) run;;
esac