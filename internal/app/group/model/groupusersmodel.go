package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupUsersModel = (*customGroupUsersModel)(nil)

type (
	// GroupUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupUsersModel.
	GroupUsersModel interface {
		groupUsersModel
	}

	customGroupUsersModel struct {
		*defaultGroupUsersModel
	}
)

// NewGroupUsersModel returns a model for the database table.
func NewGroupUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupUsersModel {
	return &customGroupUsersModel{
		defaultGroupUsersModel: newGroupUsersModel(conn, c, opts...),
	}
}
