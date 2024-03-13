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
	idGeneratorFieldNames          = builder.RawFieldNames(&IdGenerator{})
	idGeneratorRows                = strings.Join(idGeneratorFieldNames, ",")
	idGeneratorRowsExpectAutoSet   = strings.Join(stringx.Remove(idGeneratorFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	idGeneratorRowsWithPlaceHolder = strings.Join(stringx.Remove(idGeneratorFieldNames, "`caller_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheIdGeneratorCallerIdPrefix = "cache:idGenerator:callerId:"
)

type (
	idGeneratorModel interface {
		Insert(ctx context.Context, data *IdGenerator) (sql.Result, error)
		FindOne(ctx context.Context, callerId int64) (*IdGenerator, error)
		Update(ctx context.Context, data *IdGenerator) error
		Delete(ctx context.Context, callerId int64) error
	}

	defaultIdGeneratorModel struct {
		sqlc.CachedConn
		table string
	}

	IdGenerator struct {
		CallerId   int64 `db:"caller_id"`
		CallerType int64 `db:"caller_type"`
		CurSeq     int64 `db:"cur_seq"`
		MaxSeq     int64 `db:"max_seq"`
	}
)

func newIdGeneratorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultIdGeneratorModel {
	return &defaultIdGeneratorModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`id_generator`",
	}
}

func (m *defaultIdGeneratorModel) Delete(ctx context.Context, callerId int64) error {
	idGeneratorCallerIdKey := fmt.Sprintf("%s%v", cacheIdGeneratorCallerIdPrefix, callerId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `caller_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, callerId)
	}, idGeneratorCallerIdKey)
	return err
}

func (m *defaultIdGeneratorModel) FindOne(ctx context.Context, callerId int64) (*IdGenerator, error) {
	idGeneratorCallerIdKey := fmt.Sprintf("%s%v", cacheIdGeneratorCallerIdPrefix, callerId)
	var resp IdGenerator
	err := m.QueryRowCtx(ctx, &resp, idGeneratorCallerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `caller_id` = ? limit 1", idGeneratorRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, callerId)
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

func (m *defaultIdGeneratorModel) Insert(ctx context.Context, data *IdGenerator) (sql.Result, error) {
	idGeneratorCallerIdKey := fmt.Sprintf("%s%v", cacheIdGeneratorCallerIdPrefix, data.CallerId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, idGeneratorRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CallerId, data.CallerType, data.CurSeq, data.MaxSeq)
	}, idGeneratorCallerIdKey)
	return ret, err
}

func (m *defaultIdGeneratorModel) Update(ctx context.Context, data *IdGenerator) error {
	idGeneratorCallerIdKey := fmt.Sprintf("%s%v", cacheIdGeneratorCallerIdPrefix, data.CallerId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `caller_id` = ?", m.table, idGeneratorRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CallerType, data.CurSeq, data.MaxSeq, data.CallerId)
	}, idGeneratorCallerIdKey)
	return err
}

func (m *defaultIdGeneratorModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheIdGeneratorCallerIdPrefix, primary)
}

func (m *defaultIdGeneratorModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `caller_id` = ? limit 1", idGeneratorRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultIdGeneratorModel) tableName() string {
	return m.table
}
