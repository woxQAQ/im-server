package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
)

var (
	_ SessionIdModel = (*customSessionIdModel)(nil)
)

type (
	// SessionIdModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSessionIdModel.
	SessionIdModel interface {
		sessionIdModel
		ExistSessions(context.Context, int64, int64) (bool, error)
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

func (m *customSessionIdModel) ExistSessions(ctx context.Context, i int64, i2 int64) (bool, error) {
	//TODO implement me
	userIdsKey := fmt.Sprintf("cache:%s:%s", strconv.FormatInt(i, 10), strconv.FormatInt(i2, 10))
	var resp SessionId
	err := m.QueryRowCtx(ctx, resp, userIdsKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query_a := fmt.Sprintf("select %s from %s where `user_id_1` = ? and `user_id_2` = ?", sessionIdRows, m.table)
		query_b := fmt.Sprintf("select %s from %s where `user_id_2` = ? and `user_id_1` = ?", sessionIdRows, m.table)
		err := conn.QueryRowCtx(ctx, v, query_a, i, i2)
		if errors.Is(err, sqlx.ErrNotFound) {
			err = conn.QueryRowCtx(ctx, v, query_b, i, i2)
			if errors.Is(err, sqlx.ErrNotFound) {
				return ErrNotFound
			} else {
				return nil
			}
		}
		return nil
	})
	switch {
	case errors.Is(err, ErrNotFound):
		return false, err
	case err == nil:
		return true, nil
	default:
		return false, err
	}
}
