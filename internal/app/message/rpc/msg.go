package main

import (
	"flag"
	"fmt"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/config"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/server"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/msg.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMsgServer(grpcServer, server.NewMsgServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
