package svc

import (
	"github.com/woxQAQ/im-service/internal/api/ws_gateway/internal/config"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/msg"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	MsgRpc msg.Msg
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MsgRpc: msg.NewMsg(zrpc.MustNewClient(c.MsgRpc)),
	}
}
