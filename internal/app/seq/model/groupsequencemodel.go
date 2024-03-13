package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupSequenceModel = (*customGroupSequenceModel)(nil)

type (
	// GroupSequenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupSequenceModel.
	GroupSequenceModel interface {
		groupSequenceModel
	}

	customGroupSequenceModel struct {
		*defaultGroupSequenceModel
	}
)

// NewGroupSequenceModel returns a model for the database table.
func NewGroupSequenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupSequenceModel {
	return &customGroupSequenceModel{
		defaultGroupSequenceModel: newGroupSequenceModel(conn, c, opts...),
	}
}
