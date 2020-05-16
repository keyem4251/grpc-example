#!/bin/bash

rm -rfv ./src/pb

mkdir -p ./src/pb

/usr/local/bin/protoc  --proto_path protocol api.proto --go_out=plugins=grpc:./src/pb
