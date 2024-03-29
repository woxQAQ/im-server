package gateway

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/woxQAQ/im-service/pkg/common/mq"
	"github.com/woxQAQ/im-service/pkg/utils"
	"golang.org/x/sync/errgroup"
)

type WsServer struct {
	// clientManager is used to manager clients
	clientManager *ClientMgr

	// upgrader is used to upgrade a http conn to be a websocket conn
	upgrader *websocket.Upgrader

	// options
	// port is port number of WsServer
	port int

	// wsMaxConnNum limits the max conn of server
	wsMaxConnNum int64

	// wsMaxMsgLength limits the max length of message
	wsMaxMsgLength int

	rmqProducer golang.Producer

	rmqTopic string

	*validator.Validate
}

// NewWsServer gets a new websocket server
func NewWsServer(rmqConfig *golang.Config, rmqTopic string, opts ...Option) (*WsServer, error) {
	var config configs
	for _, o := range opts {
		o(&config)
	}
	if config.port == 0 {
		// use default port
		config.port = 8080
	}
	if config.maxConnNum == 0 {
		config.maxConnNum = 1024
	}
	if config.writeBufSize == 0 {
		config.writeBufSize = 1024
	}
	if config.handshakeTimeout == 0 {
		config.handshakeTimeout = 10 * time.Second
	}
	if config.maxMsgLength == 0 {
		config.maxMsgLength = maxMessageSize
	}

	validate := validator.New()
	producer, err := mq.NewProducer(rmqConfig, rmqTopic)
	if err != nil {
		return nil, err
	}

	return &WsServer{
		port:           config.port,
		wsMaxConnNum:   config.maxConnNum,
		wsMaxMsgLength: config.maxMsgLength,
		clientManager:  newClientManager(),
		Validate:       validate,
		rmqProducer:    producer,
		rmqTopic:       rmqTopic,
		upgrader: &websocket.Upgrader{
			HandshakeTimeout:  config.handshakeTimeout,
			WriteBufferSize:   config.writeBufSize,
			EnableCompression: true,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}, nil
}

// ReplaceLogger replace global log
func ReplaceLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("replace logger failed", err.Error())
	}
	zap.ReplaceGlobals(logger)
}

// Bootstrap begin wsServer
// 1. start http server -> websocket server
// 2. start clientManager
// 3. start rocketMq producer
func (ws *WsServer) Bootstrap() error {
	var (
		wg errgroup.Group

		signs = make(chan os.Signal, 1)
		done  = make(chan struct{}, 1)
	)
	//gin.SetMode(gin.DebugMode)
	r := gin.Default()
	ReplaceLogger()
	r.Use(CorsHandler())
	r.Use(jwtHandler())
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

	wg.Go(func() error {
		return ws.rmqProducer.Start()
	})

	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signs
	go func() {
		fmt.Println("Shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
		_ = wg.Wait()
		ws.rmqProducer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-time.After(15 * time.Second):
		return utils.WrapWithCallerInfo(errors.New("Timeout Exit"))
	}
}
