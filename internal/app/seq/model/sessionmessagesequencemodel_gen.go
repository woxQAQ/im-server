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
	sessionMessageSequenceFieldNames          = builder.RawFieldNames(&SessionMessageSequence{})
	sessionMessageSequenceRows                = strings.Join(sessionMessageSequenceFieldNames, ",")
	sessionMessageSequenceRowsExpectAutoSet   = strings.Join(stringx.Remove(sessionMessageSequenceFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sessionMessageSequenceRowsWithPlaceHolder = strings.Join(stringx.Remove(sessionMessageSequenceFieldNames, "`session_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSessionMessageSequenceSessionIdPrefix = "cache:sessionMessageSequence:sessionId:"
)

type (
	sessionMessageSequenceModel interface {
		Insert(ctx context.Context, data *SessionMessageSequence) (sql.Result, error)
		FindOne(ctx context.Context, sessionId int64) (*SessionMessageSequence, error)
		Update(ctx context.Context, data *SessionMessageSequence) error
		Delete(ctx context.Context, sessionId int64) error
	}

	defaultSessionMessageSequenceModel struct {
		sqlc.CachedConn
		table string
	}

	SessionMessageSequence struct {
		SessionId     int64 `db:"session_id"`
		MsgId         int64 `db:"msg_id"`
		MsgSeqSession int64 `db:"msg_seq_session"`
	}
)

func newSessionMessageSequenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSessionMessageSequenceModel {
	return &defaultSessionMessageSequenceModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`session_message_sequence`",
	}
}

func (m *defaultSessionMessageSequenceModel) Delete(ctx context.Context, sessionId int64) error {
	sessionMessageSequenceSessionIdKey := fmt.Sprintf("%s%v", cacheSessionMessageSequenceSessionIdPrefix, sessionId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `session_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, sessionId)
	}, sessionMessageSequenceSessionIdKey)
	return err
}

func (m *defaultSessionMessageSequenceModel) FindOne(ctx context.Context, sessionId int64) (*SessionMessageSequence, error) {
	sessionMessageSequenceSessionIdKey := fmt.Sprintf("%s%v", cacheSessionMessageSequenceSessionIdPrefix, sessionId)
	var resp SessionMessageSequence
	err := m.QueryRowCtx(ctx, &resp, sessionMessageSequenceSessionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `session_id` = ? limit 1", sessionMessageSequenceRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, sessionId)
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

func (m *defaultSessionMessageSequenceModel) Insert(ctx context.Context, data *SessionMessageSequence) (sql.Result, error) {
	sessionMessageSequenceSessionIdKey := fmt.Sprintf("%s%v", cacheSessionMessageSequenceSessionIdPrefix, data.SessionId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sessionMessageSequenceRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.SessionId, data.MsgId, data.MsgSeqSession)
	}, sessionMessageSequenceSessionIdKey)
	return ret, err
}

func (m *defaultSessionMessageSequenceModel) Update(ctx context.Context, data *SessionMessageSequence) error {
	sessionMessageSequenceSessionIdKey := fmt.Sprintf("%s%v", cacheSessionMessageSequenceSessionIdPrefix, data.SessionId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `session_id` = ?", m.table, sessionMessageSequenceRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.MsgId, data.MsgSeqSession, data.SessionId)
	}, sessionMessageSequenceSessionIdKey)
	return err
}

func (m *defaultSessionMessageSequenceModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSessionMessageSequenceSessionIdPrefix, primary)
}

func (m *defaultSessionMessageSequenceModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `session_id` = ? limit 1", sessionMessageSequenceRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSessionMessageSequenceModel) tableName() string {
	return m.table
}