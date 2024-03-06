package gateway

import "time"

type configs struct {
	// port is the port of websocket server
	port int

	// maxConnNum is the max websocket connection number
	maxConnNum int64

	// handshakeTimeout is the websocket handshake timeout
	handshakeTimeout time.Duration

	// maxMsgLength is the max message length
	maxMsgLength int

	// writeBufSize is the write buffer size
	writeBufSize int
}



func WithPort(port int) Option {
	return func(opt *configs) {
		opt.port = port
	}
}

func WithMaxConnNum(maxConnNum int64) Option {
	return func(opt *configs) {
		opt.maxConnNum = maxConnNum
	}
}

func WithHandshakeTimeout(handshakeTimeout time.Duration) Option {
	return func(opt *configs) {
		opt.handshakeTimeout = handshakeTimeout
	}
}

func WithMaxMessageLength(maxMsgLength int) Option {
	return func(opt *configs) {
		opt.maxMsgLength = maxMsgLength
	}
}

func WithWriteBufSize(writeBufSize int) Option {
	return func(opt *configs) {
		opt.writeBufSize = writeBufSize
	}
}

type Option func(opt *configs)
