#!bin/bash

goctl rpc protoc msg.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.