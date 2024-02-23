package main

import "github.com/woxQAQ/im-service/internal/gateway"

func main() {
	srv, err := gateway.NewWsServer(gateway.WithPort(8089))
	if err != nil {
		panic(err)
	}
	srv.Bootstrap()
}
