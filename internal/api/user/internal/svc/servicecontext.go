package svc

import (
	"github.com/woxQAQ/im-service/internal/api/user/internal/config"
	"github.com/woxQAQ/im-service/internal/rpc/user/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
