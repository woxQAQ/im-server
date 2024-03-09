package main

import (
	"flag"
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/woxQAQ/im-service/config"
	"github.com/woxQAQ/im-service/internal/gateway"
	"github.com/zeromicro/go-zero/core/conf"
)

var gatewayConfigPath = flag.String("f", "cmd/gateway/gateway.yaml", "gateway config file")

func main() {
	flag.Parse()

	var gwConfig config.GatewayConfig
	conf.MustLoad(*gatewayConfigPath, &gwConfig)

	srv, err := gateway.NewWsServer(&golang.Config{
		Endpoint: gwConfig.Config.Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    gwConfig.Config.Credentials.AccessKey,
			AccessSecret: gwConfig.Config.Credentials.AccessSecret,
		},
	}, gwConfig.Topic, gateway.WithPort(10880))
	if err != nil {
		panic(err)
	}
	srv.Bootstrap()
}
