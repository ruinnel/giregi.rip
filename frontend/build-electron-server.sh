#!/bin/bash

if [ -z "$1" ]; then
  OS=$(go env GOOS);
else
  OS=$1;
fi

if [ -z "$2" ]; then
  ARCH=$(go env GOARCH);
else
  ARCH=$2;
fi

BINARY_DIR="$(pwd)/bin";
cd ../server;
export CGO_ENABLED=0;
export GOOS="${OS}";
export GOARCH="${ARCH}";

EXT="";

if [ "${GOOS}" == "windows" ]; then
    EXT=".exe";
fi

echo "goos - ${GOOS}, arch - ${GOARCH}";
echo "";

go build -ldflags "-X main.platform=desktop" -o ${BINARY_DIR}/${GOOS}/server${EXT} app/main.go;

cd ../frontend;
cp ${BINARY_DIR}/${GOOS}/server${EXT} ./dist;
cp server-config.yaml ./dist/config.yaml;
