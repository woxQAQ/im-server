package svc

import "github.com/woxQAQ/im-service/internal/rpc/imrpc_session/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
