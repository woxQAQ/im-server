package main

import (
	"flag"

	"github.com/woxQAQ/im-service/config"
	"github.com/woxQAQ/im-service/internal/gateway"
	"github.com/zeromicro/go-zero/core/conf"
)

var gatewayConfigPath = flag.String("f", "config/gateway.yaml", "gateway config file")

func main() {
	flag.Parse()

	var gwConfig config.GatewayConfig
	conf.MustLoad(*gatewayConfigPath, &gwConfig)

	srv, err := gateway.NewWsServer(gateway.WithPort(9111), gateway.WithTopic(gwConfig.Rmq.Topic), gateway.WithRmqEndpoint(gwConfig.Rmq.NamesrvAddress))
	if err != nil {
		panic(err)
	}
	srv.Bootstrap()
}
