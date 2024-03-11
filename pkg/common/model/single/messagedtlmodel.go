package single

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageDtlModel = (*customMessageDtlModel)(nil)

type (
	// MessageDtlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageDtlModel.
	MessageDtlModel interface {
		messageDtlModel
	}

	customMessageDtlModel struct {
		*defaultMessageDtlModel
	}
)

// NewMessageDtlModel returns a model for the database table.
func NewMessageDtlModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MessageDtlModel {
	return &customMessageDtlModel{
		defaultMessageDtlModel: newMessageDtlModel(conn, c, opts...),
	}
}
