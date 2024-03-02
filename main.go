package main

import (
	"flag"
	"github.com/woxQAQ/im-service/internal/gateway"
)

var gatewayConfigPath = flag.String("f", "internal/gateway/etc/gateway.yaml", "the config file")

func main() {
	srv, err := gateway.NewWsServer(*gatewayConfigPath, gateway.WithPort(8089))
	if err != nil {
		panic(err)
	}
	srv.Bootstrap()
}
