package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
)

var _ TSenderDtlModel = (*customTSenderDtlModel)(nil)

var cacheTSenderSenderIdPrefix = "cache:tSender:Sender:"

type (
	// TSenderDtlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSenderDtlModel.
	TSenderDtlModel interface {
		tSenderDtlModel
		FindBySenderAndSession(ctx context.Context, senderId int64, sessionId string) (int64, error)
	}

	customTSenderDtlModel struct {
		*defaultTSenderDtlModel
	}
)

// NewTSenderDtlModel returns a model for the database table.
func NewTSenderDtlModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TSenderDtlModel {
	return &customTSenderDtlModel{
		defaultTSenderDtlModel: newTSenderDtlModel(conn, c, opts...),
	}
}

func (m *customTSenderDtlModel) FindBySenderAndSession(ctx context.Context, senderId int64, sessionId string) (int64, error) {
	var res TSenderDtl
	tSenderDtlKey := fmt.Sprintf(cacheTSenderDtlIdPrefix + strconv.FormatInt(senderId, 10))
	err := m.QueryRowCtx(ctx, &res, tSenderDtlKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, "select preid from ? where sender_id = ? and `session` = ? LIMIT 1", m.tableName(), strconv.FormatInt(senderId, 10), sessionId)
	})

	switch err {
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	case nil:
		return res.PreId, nil
	default:
		return 0, err
	}
}
