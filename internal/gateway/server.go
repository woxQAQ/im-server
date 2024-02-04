package gateway

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

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

type WsServer struct {
	// clientManager is used to manager clients
	clientManager *ClientMgr

	port int

	clientsMap *Usermap

	clientPool sync.Pool

	handshakeTimeout time.Duration

	writeBufSize int

	wsMaxConnNum int64

	wsMaxMsgLength int
}

func NewWsServer(opts ...option) (*WsServer, error) {
	var config configs
	for _, o := range opts {
		o(&config)
	}
	return &WsServer{
		port:             config.port,
		writeBufSize:     config.writeBufSize,
		handshakeTimeout: config.handshakeTimeout,
		wsMaxConnNum:     config.maxConnNum,
		wsMaxMsgLength:   config.maxMsgLength,
		clientManager:    newClientManager(),
		clientPool: sync.Pool{
			New: func() any {
				return new(Client)
			},
		},
		clientsMap: newUserMap(),
	}, nil
}

func (ws *WsServer) Bootstrap() error {
	var (
		client *Client
		wg     errgroup.Group

		signs = make(chan os.Signal, 1)
		done  = make(chan struct{}, 1)
	)
	r := gin.Default()
	// TODO: register gin router

	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(int64(ws.port), 10),
		Handler: r,
	}

	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signs

	return nil
}
