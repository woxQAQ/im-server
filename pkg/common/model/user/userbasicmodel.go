package user

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserbasicModel = (*customUserbasicModel)(nil)

type (
	// UserbasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserbasicModel.
	UserbasicModel interface {
		userbasicModel
	}

	customUserbasicModel struct {
		*defaultUserbasicModel
	}
)

// NewUserbasicModel returns a model for the database table.
func NewUserbasicModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserbasicModel {
	return &customUserbasicModel{
		defaultUserbasicModel: newUserbasicModel(conn, c, opts...),
	}
}
