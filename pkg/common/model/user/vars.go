package user

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	ErrNotFound      = sqlx.ErrNotFound
	ErrGenderInvalid = errors.New("Gender invalid")
)
