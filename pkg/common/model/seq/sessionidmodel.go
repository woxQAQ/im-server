package seq

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SessionIdModel = (*customSessionIdModel)(nil)

type (
	// SessionIdModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSessionIdModel.
	SessionIdModel interface {
		sessionIdModel
	}

	customSessionIdModel struct {
		*defaultSessionIdModel
	}
)

// NewSessionIdModel returns a model for the database table.
func NewSessionIdModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SessionIdModel {
	return &customSessionIdModel{
		defaultSessionIdModel: newSessionIdModel(conn, c, opts...),
	}
}
