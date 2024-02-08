package utils

import (
	"crypto/md5"
	"time"
)

func Timestamp() int {
	return int(time.Now().UnixNano() / 1e6)
}

func Md5(origin string) string {
	return string(md5.New().Sum([]byte(origin)))
}
