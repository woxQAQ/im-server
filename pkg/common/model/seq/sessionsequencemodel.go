package seq

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SessionSequenceModel = (*customSessionSequenceModel)(nil)

type (
	// SessionSequenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSessionSequenceModel.
	SessionSequenceModel interface {
		sessionSequenceModel
	}

	customSessionSequenceModel struct {
		*defaultSessionSequenceModel
	}
)

// NewSessionSequenceModel returns a model for the database table.
func NewSessionSequenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SessionSequenceModel {
	return &customSessionSequenceModel{
		defaultSessionSequenceModel: newSessionSequenceModel(conn, c, opts...),
	}
}
