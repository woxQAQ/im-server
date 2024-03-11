package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

func UseSharding(db *gorm.DB, shardingKey string, numberOfShards uint, PrimaryKeyGenerator int, tblName ...string) (*sqlx.SqlConn, error) {
	err := db.Use(sharding.Register(sharding.Config{
		ShardingKey:         shardingKey,
		NumberOfShards:      numberOfShards,
		PrimaryKeyGenerator: PrimaryKeyGenerator,
	}, tblName))
	if err != nil {
		return nil, err
	}

	DB, err := db.DB()
	if err != nil {
		return nil, err
	}

	conn := sqlx.NewSqlConnFromDB(DB)
	return &conn, nil
}
