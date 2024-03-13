package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IdGeneratorModel = (*customIdGeneratorModel)(nil)

type (
	// IdGeneratorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIdGeneratorModel.
	IdGeneratorModel interface {
		idGeneratorModel
	}

	customIdGeneratorModel struct {
		*defaultIdGeneratorModel
	}
)

// NewIdGeneratorModel returns a model for the database table.
func NewIdGeneratorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IdGeneratorModel {
	return &customIdGeneratorModel{
		defaultIdGeneratorModel: newIdGeneratorModel(conn, c, opts...),
	}
}
