#!/bin/sh

chain_path=$(go list -f '{{.Dir}}' "github.com/assetcloud/chain")
protoc --go_out=plugins=grpc:../types ./*.proto --proto_path=. --proto_path="${chain_path}/types/proto/"
