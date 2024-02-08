package gateway

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/woxQAQ/im-service/pkg/utils"
	"golang.org/x/sync/errgroup"
)

type WsServer struct {
	// clientManager is used to manager clients
	clientManager *ClientMgr

	// clientPool is used to reuse clients objects
	clientPool sync.Pool

	// upgrader is used to upgrade a http conn to be a websocket conn
	upgrader *websocket.Upgrader

	// options
	// port is port number of WsServer
	port int

	// wsMaxConnNum limits the max conn of server
	wsMaxConnNum int64

	// wsMaxMsgLength limits the max length of message
	wsMaxMsgLength int
}

func NewWsServer(opts ...option) (*WsServer, error) {
	var config configs
	for _, o := range opts {
		o(&config)
	}
	return &WsServer{
		port:           config.port,
		wsMaxConnNum:   config.maxConnNum,
		wsMaxMsgLength: config.maxMsgLength,
		clientManager:  newClientManager(),
		clientPool: sync.Pool{
			New: func() any {
				return new(Client)
			},
		},
		upgrader: &websocket.Upgrader{
			HandshakeTimeout:  config.handshakeTimeout,
			WriteBufferSize:   config.writeBufSize,
			EnableCompression: true,
		},
	}, nil
}

func (ws *WsServer) Bootstrap() error {
	var (
		wg errgroup.Group

		signs = make(chan os.Signal, 1)
		done  = make(chan struct{}, 1)
	)
	r := gin.Default()
	// TODO: register gin router
	ws.registerRouter(r)

	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(int64(ws.port), 10),
		Handler: r,
	}

	// Bootstrap clientManager
	wg.Go(func() error {
		return ws.clientManager.Run(done)
	})

	// Bootstrap http server
	wg.Go(func() error {
		return srv.ListenAndServe()
	})

	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signs
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
		_ = wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-time.After(15 * time.Second):
		return utils.WrapWithCallerInfo(errors.New("Timeout Exit"))
	}
}
