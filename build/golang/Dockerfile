FROM golang:1.21
LABEL maintainer="ShirakamiYuki<woxqaq@gmail.com>"

USER root
RUN apt-get update && \
        apt-get upgrade -y && \
  apt-get install -y --no-install-recommends unzip curl zip git vim screen

USER root
RUN go install github.com/tal-tech/go-zero/tools/goctl@cli

USER root
RUN curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip && \
    unzip -d /tmp/protoc /tmp/protoc.zip && \
    mv /tmp/protoc/bin/protoc $GOPATH/bin

USER root
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

ENV PATH $GOPATH/bin:$PATH

WORKDIR /usr/src/code

EXPOSE 8000
EXPOSE 8001
EXPOSE 8002
EXPOSE 8003
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003