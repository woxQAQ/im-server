package gateway

import (
	"errors"
	"time"
)

const (
	maxMessageSize = 51200
	pongwait       = 30 * time.Second
)

var (
	ErrConnOverMaxNumLimit = errors.New("ConnOverMaxNumLimit")
	ErrWebsockerUpgrade    = errors.New("WebsocketUpgrade")
	ErrArgumentErr         = errors.New("ArgumentError")
	ErrSenderIdNotMatch    = errors.New("SenderId not match the userId it comes from")
)
