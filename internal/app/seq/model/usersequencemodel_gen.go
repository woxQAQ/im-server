// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userSequenceFieldNames          = builder.RawFieldNames(&UserSequence{})
	userSequenceRows                = strings.Join(userSequenceFieldNames, ",")
	userSequenceRowsExpectAutoSet   = strings.Join(stringx.Remove(userSequenceFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userSequenceRowsWithPlaceHolder = strings.Join(stringx.Remove(userSequenceFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserSequenceUserIdPrefix = "cache:userSequence:userId:"
)

type (
	userSequenceModel interface {
		Insert(ctx context.Context, data *UserSequence) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserSequence, error)
		Update(ctx context.Context, data *UserSequence) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserSequenceModel struct {
		sqlc.CachedConn
		table string
	}

	UserSequence struct {
		UserId int64 `db:"user_id"`
		CurSeq int64 `db:"cur_seq"`
		MaxSeq int64 `db:"max_seq"`
	}
)

func newUserSequenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserSequenceModel {
	return &defaultUserSequenceModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_sequence`",
	}
}

func (m *defaultUserSequenceModel) Delete(ctx context.Context, userId int64) error {
	userSequenceUserIdKey := fmt.Sprintf("%s%v", cacheUserSequenceUserIdPrefix, userId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, userSequenceUserIdKey)
	return err
}

func (m *defaultUserSequenceModel) FindOne(ctx context.Context, userId int64) (*UserSequence, error) {
	userSequenceUserIdKey := fmt.Sprintf("%s%v", cacheUserSequenceUserIdPrefix, userId)
	var resp UserSequence
	err := m.QueryRowCtx(ctx, &resp, userSequenceUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userSequenceRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserSequenceModel) Insert(ctx context.Context, data *UserSequence) (sql.Result, error) {
	userSequenceUserIdKey := fmt.Sprintf("%s%v", cacheUserSequenceUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userSequenceRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CurSeq, data.MaxSeq)
	}, userSequenceUserIdKey)
	return ret, err
}

func (m *defaultUserSequenceModel) Update(ctx context.Context, data *UserSequence) error {
	userSequenceUserIdKey := fmt.Sprintf("%s%v", cacheUserSequenceUserIdPrefix, data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userSequenceRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CurSeq, data.MaxSeq, data.UserId)
	}, userSequenceUserIdKey)
	return err
}

func (m *defaultUserSequenceModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserSequenceUserIdPrefix, primary)
}

func (m *defaultUserSequenceModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userSequenceRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserSequenceModel) tableName() string {
	return m.table
}