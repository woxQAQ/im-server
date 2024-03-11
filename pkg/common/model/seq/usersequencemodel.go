package seq

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSequenceModel = (*customUserSequenceModel)(nil)

type (
	// UserSequenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSequenceModel.
	UserSequenceModel interface {
		userSequenceModel
	}

	customUserSequenceModel struct {
		*defaultUserSequenceModel
	}
)

// NewUserSequenceModel returns a model for the database table.
func NewUserSequenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserSequenceModel {
	return &customUserSequenceModel{
		defaultUserSequenceModel: newUserSequenceModel(conn, c, opts...),
	}
}
