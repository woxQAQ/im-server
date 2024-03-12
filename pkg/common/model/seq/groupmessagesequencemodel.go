package seq

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupMessageSequenceModel = (*customGroupMessageSequenceModel)(nil)

type (
	// GroupMessageSequenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupMessageSequenceModel.
	GroupMessageSequenceModel interface {
		groupMessageSequenceModel
	}

	customGroupMessageSequenceModel struct {
		*defaultGroupMessageSequenceModel
	}
)

// NewGroupMessageSequenceModel returns a model for the database table.
func NewGroupMessageSequenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupMessageSequenceModel {
	return &customGroupMessageSequenceModel{
		defaultGroupMessageSequenceModel: newGroupMessageSequenceModel(conn, c, opts...),
	}
}
