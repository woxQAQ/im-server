package config

import (
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/zrpc"
)

var ErrConfigFieldNotCompleted = errors.New("ConfigFieldNotComplete")

type RmqConfig struct {
	Endpoint      string
	Namespace     string `json:",optional"`
	ConsumerGroup string `json:",optional"`
	Credentials   struct {
		AccessKey     string
		AccessSecret  string
		SecurityToken string `json:",optional"`
	}
}

func (c *RmqConfig) GetConf() (*golang.Config, error) {
	if c.Endpoint == "" {
		return nil, ErrConfigFieldNotCompleted
	}
	return &golang.Config{
		Endpoint:      c.Endpoint,
		NameSpace:     c.Namespace,
		ConsumerGroup: c.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:     c.Credentials.AccessKey,
			AccessSecret:  c.Credentials.AccessSecret,
			SecurityToken: c.Credentials.SecurityToken,
		},
	}, nil
}

type GatewayConfig struct {
	Topic  string
	Config RmqConfig
}

type HandlerConfig struct {
	RmqConfig RmqConfig
	RpcConfig zrpc.RpcClientConf
	Topic     struct {
		ConsumerTopic string
		ProducerTopic string
	}
}

type SeqConfig struct {
	Mysql struct {
		DataSource string
	}
}
