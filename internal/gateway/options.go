package gateway

import "time"

type configs struct {
	port int

	maxConnNum int64

	handshakeTimeout time.Duration

	maxMsgLength int

	writeBufSize int
}

func withPort(port int) option {
	return func(opt *configs) {
		opt.port = port
	}
}

func withMaxConnNum(maxConnNum int64) option {
	return func(opt *configs) {
		opt.maxConnNum = maxConnNum
	}
}

func withHandshakeTimeout(handshakeTimeout time.Duration) option {
	return func(opt *configs) {
		opt.handshakeTimeout = handshakeTimeout
	}
}

func withMaxMessageLength(maxMsgLength int) option {
	return func(opt *configs) {
		opt.maxMsgLength = maxMsgLength
	}
}

func withWriteBufSize(writeBufSize int) option {
	return func(opt *configs) {
		opt.writeBufSize = writeBufSize
	}
}

type option func(opt *configs)
