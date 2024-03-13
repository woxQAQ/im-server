package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SessionMessageSequenceModel = (*customSessionMessageSequenceModel)(nil)

type (
	// SessionMessageSequenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSessionMessageSequenceModel.
	SessionMessageSequenceModel interface {
		sessionMessageSequenceModel
	}

	customSessionMessageSequenceModel struct {
		*defaultSessionMessageSequenceModel
	}
)

// NewSessionMessageSequenceModel returns a model for the database table.
func NewSessionMessageSequenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SessionMessageSequenceModel {
	return &customSessionMessageSequenceModel{
		defaultSessionMessageSequenceModel: newSessionMessageSequenceModel(conn, c, opts...),
	}
}
