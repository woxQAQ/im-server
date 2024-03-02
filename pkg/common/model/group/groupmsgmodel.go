package group

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupMsgModel = (*customGroupMsgModel)(nil)

type (
	// GroupMsgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupMsgModel.
	GroupMsgModel interface {
		groupMsgModel
	}

	customGroupMsgModel struct {
		*defaultGroupMsgModel
	}
)

// NewGroupMsgModel returns a model for the database table.
func NewGroupMsgModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupMsgModel {
	return &customGroupMsgModel{
		defaultGroupMsgModel: newGroupMsgModel(conn, c, opts...),
	}
}
