#!bin/bash

goctl rpc protoc seq.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.