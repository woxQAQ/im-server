// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userbasicFieldNames          = builder.RawFieldNames(&Userbasic{})
	userbasicRows                = strings.Join(userbasicFieldNames, ",")
	userbasicRowsExpectAutoSet   = strings.Join(stringx.Remove(userbasicFieldNames, "`Id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userbasicRowsWithPlaceHolder = strings.Join(stringx.Remove(userbasicFieldNames, "`Id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserbasicIdPrefix          = "cache:userbasic:Id:"
	cacheUserbasicEmailPrefix       = "cache:userbasic:email:"
	cacheUserbasicMobilePhonePrefix = "cache:userbasic:mobilePhone:"
)

type (
	userbasicModel interface {
		Insert(ctx context.Context, data *Userbasic) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Userbasic, error)
		FindOneByEmail(ctx context.Context, email string) (*Userbasic, error)
		FindOneByMobilePhone(ctx context.Context, mobilePhone string) (*Userbasic, error)
		Update(ctx context.Context, data *Userbasic) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserbasicModel struct {
		sqlc.CachedConn
		table string
	}

	Userbasic struct {
		Id          int64     `db:"Id"`
		Name        string    `db:"name"`
		Gender      string    `db:"gender"`
		MobilePhone string    `db:"mobile_phone"`
		Email       string    `db:"email"`
		Password    string    `db:"password"`
		CreateTime  time.Time `db:"create_time"`
		UpdateTime  time.Time `db:"update_time"`
	}
)

func newUserbasicModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserbasicModel {
	return &defaultUserbasicModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`userbasic`",
	}
}

func (m *defaultUserbasicModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userbasicEmailKey := fmt.Sprintf("%s%v", cacheUserbasicEmailPrefix, data.Email)
	userbasicIdKey := fmt.Sprintf("%s%v", cacheUserbasicIdPrefix, id)
	userbasicMobilePhoneKey := fmt.Sprintf("%s%v", cacheUserbasicMobilePhonePrefix, data.MobilePhone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `Id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userbasicEmailKey, userbasicIdKey, userbasicMobilePhoneKey)
	return err
}

func (m *defaultUserbasicModel) FindOne(ctx context.Context, id int64) (*Userbasic, error) {
	userbasicIdKey := fmt.Sprintf("%s%v", cacheUserbasicIdPrefix, id)
	var resp Userbasic
	err := m.QueryRowCtx(ctx, &resp, userbasicIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `Id` = ? limit 1", userbasicRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUserbasicModel) FindOneByEmail(ctx context.Context, email string) (*Userbasic, error) {
	userbasicEmailKey := fmt.Sprintf("%s%v", cacheUserbasicEmailPrefix, email)
	var resp Userbasic
	err := m.QueryRowIndexCtx(ctx, &resp, userbasicEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userbasicRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserbasicModel) FindOneByMobilePhone(ctx context.Context, mobilePhone string) (*Userbasic, error) {
	userbasicMobilePhoneKey := fmt.Sprintf("%s%v", cacheUserbasicMobilePhonePrefix, mobilePhone)
	var resp Userbasic
	err := m.QueryRowIndexCtx(ctx, &resp, userbasicMobilePhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile_phone` = ? limit 1", userbasicRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobilePhone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserbasicModel) Insert(ctx context.Context, data *Userbasic) (sql.Result, error) {
	userbasicEmailKey := fmt.Sprintf("%s%v", cacheUserbasicEmailPrefix, data.Email)
	userbasicIdKey := fmt.Sprintf("%s%v", cacheUserbasicIdPrefix, data.Id)
	userbasicMobilePhoneKey := fmt.Sprintf("%s%v", cacheUserbasicMobilePhonePrefix, data.MobilePhone)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userbasicRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Gender, data.MobilePhone, data.Email, data.Password)
	}, userbasicEmailKey, userbasicIdKey, userbasicMobilePhoneKey)
	return ret, err
}

func (m *defaultUserbasicModel) Update(ctx context.Context, newData *Userbasic) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userbasicEmailKey := fmt.Sprintf("%s%v", cacheUserbasicEmailPrefix, data.Email)
	userbasicIdKey := fmt.Sprintf("%s%v", cacheUserbasicIdPrefix, data.Id)
	userbasicMobilePhoneKey := fmt.Sprintf("%s%v", cacheUserbasicMobilePhonePrefix, data.MobilePhone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `Id` = ?", m.table, userbasicRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Gender, newData.MobilePhone, newData.Email, newData.Password, newData.Id)
	}, userbasicEmailKey, userbasicIdKey, userbasicMobilePhoneKey)
	return err
}

func (m *defaultUserbasicModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserbasicIdPrefix, primary)
}

func (m *defaultUserbasicModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `Id` = ? limit 1", userbasicRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserbasicModel) tableName() string {
	return m.table
}
