package svc

import (
	"github.com/woxQAQ/im-service/internal/rpc/user/internal/config"
	model "github.com/woxQAQ/im-service/pkg/common/sql/user"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserbasicModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
