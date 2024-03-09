package main

import (
	"flag"
	"github.com/woxQAQ/im-service/config"
	"github.com/woxQAQ/im-service/internal/handler"
	"github.com/zeromicro/go-zero/core/conf"
)

var handlerConfigPath = flag.String("f", "cmd/handler/handler.yaml", "gateway config file")

func main() {
	flag.Parse()

	var Config config.HandlerConfig
	conf.MustLoad(*handlerConfigPath, &Config)

	h, err := handler.New(&Config)
	if err != nil {
		panic(err)
	}

	err = h.Run()
	if err != nil {
		panic(err)
	}
}
