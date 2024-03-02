package single

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SingleMessageModel = (*customSingleMessageModel)(nil)

type (
	// SingleMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSingleMessageModel.
	SingleMessageModel interface {
		singleMessageModel
	}

	customSingleMessageModel struct {
		*defaultSingleMessageModel
	}
)

// NewSingleMessageModel returns a model for the database table.
func NewSingleMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SingleMessageModel {
	return &customSingleMessageModel{
		defaultSingleMessageModel: newSingleMessageModel(conn, c, opts...),
	}
}
