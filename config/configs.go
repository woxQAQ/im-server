package config

import (
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/zeromicro/go-zero/zrpc"
)

type Rmq struct {
	Topic  string
	Config golang.Config
}

type GatewayConfig struct {
	Rmq
}

type RouterConfig struct {
	RmqConfig golang.Config      `validate:"required"`
	RpcConfig zrpc.RpcClientConf `validate:"required"`
	Topic     struct {
		ConsumerTopic string `validate:"required"`
		ProducerTopic string `validate:"required"`
	}
}

type SeqConfig struct {
	Mysql struct{
		DataSource string
	}
}
