package gateway

import (
	"bytes"
	"context"
	"log"
	"runtime/debug"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/woxQAQ/im-service/pkg/utils"
	"go.uber.org/zap"
)

var (
	ErrConnClosed                = errors.New("conn has closed")
	ErrNotSupportMessageProtocol = errors.New("not support message protocol")
	ErrClientClosed              = errors.New("client actively close the connection")
	ErrPanic                     = errors.New("panic error")
)

type Client struct {
	w *sync.Mutex
	// ClientId is the Identify of a Client
	ClientId string

	// UserId is the Id of the user of Client
	// UserId is different from ClientId
	// UserId identify a real user of the Im service
	// ClientId, on the other hand, can be called "sessionId"
	UserId string

	// PlatformId is the system of the user
	PlatformId int

	// MessageChan is used to send message to the ClientManager
	MessageChan chan []byte

	// closed used to mark the client is closed or not
	closed atomic.Bool

	// CloseErr is the error why the client close
	CloseErr error

	// token is the token of the user
	token string

	ConnectTime uint64

	// Context is the
	Context *gin.Context

	RemoteIp string

	// Conn is the connection that Client connect to server
	Conn *websocket.Conn

	// Mgr is the ClientManager that Client register to.
	// the Mgr is used to manager the client,
	// including sending and reading message...
	Mgr *ClientMgr

	Server *WsServer
}

func NewClient(
	ctx *gin.Context,
	conn *websocket.Conn,
	mgr *ClientMgr,
	token string,
	server *WsServer,
) *Client {
	platformId, _ := strconv.Atoi(ctx.Query("platformId"))
	return &Client{
		w:           new(sync.Mutex),
		ClientId:    utils.Md5(ctx.RemoteIP() + "_" + strconv.Itoa(utils.Timestamp())),
		Conn:        conn,
		Mgr:         mgr,
		Server:      server,
		ConnectTime: uint64(time.Now().Unix()),
		PlatformId:  platformId,
		RemoteIp:    ctx.RemoteIP(),
		Context:     ctx,
		token:       token,
	}
}

func (c *Client) resetClient(
	conn *websocket.Conn,
	mgr *ClientMgr,
	userId, token string,
	platformId int,
) {
	c.ClientId = uuid.New().String()
	c.w = new(sync.Mutex)
	c.PlatformId = platformId
	c.token = token
	c.UserId = userId
	c.Conn = conn
	c.Mgr = mgr
	c.ConnectTime = uint64(time.Now().Unix())
	c.CloseErr = nil
	c.closed.Store(false)
}

// close used to close a client
func (c *Client) close() {
	if c.closed.Load() {
		return
	}
	c.w.Lock()
	defer c.w.Unlock()

	c.closed.Store(true)
	c.Conn.Close()
	c.Server.clientManager.unregisterChan <- c
}

// Read used to read message from peer of client
func (c *Client) Read() {
	defer func() {
		// program exit may due to panic, here capture panic
		if r := recover(); r != nil {
			c.CloseErr = ErrPanic
			zap.S().Panic("conn have panic err:", r, string(debug.Stack()))
		}
		c.close()
	}()

	// make prepare for reading: setting maxMessageSize and deadline
	c.Conn.SetReadLimit(maxMessageSize)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongwait))

	for {
		messageType, data, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zap.S().Info("error due to peers close normal: ", err)
				break
			}
			zap.S().Error("Unexpected close error:", err)
			break
		}
		switch messageType {

		// Json data, which we dont use
		case websocket.TextMessage:
			c.CloseErr = ErrNotSupportMessageProtocol
			return

		// protobuf data
		case websocket.BinaryMessage:
			_ = c.Conn.SetReadDeadline(time.Now().Add(pongwait))
			// TODO: handler message: send the message to the transfer server

			// remove '\n' to be space, and remove the message's edge's space
			// is used to process multiline text to be one line
			data = bytes.TrimSpace(bytes.Replace(data, []byte("\n"), []byte(" "), -1))
			rmqMsg := golang.Message{
				Topic: c.Server.rmqTopic,
				Body:  data,
			}

			rmqMsg.SetTag("chatmsg")
			c.Server.rmqProducer.SendAsync(context.TODO(), &rmqMsg, func(ctx context.Context, sr []*golang.SendReceipt, err error) {
				if err != nil {
					log.Fatal(err)
				}
			})

		// PingMessage is used to validate a conn is alive or not
		case websocket.PingMessage:
			err = c.pingHandler()
			zap.S().Error(err)

		case websocket.CloseMessage:
			c.CloseErr = ErrClientClosed
			return
		default:
		}
	}
}

// func (c *Client) handlerRequest(Request []byte) error {
// 	// TODO: handle message
// 	var req = getReq()
// 	defer freeReq(req)

// 	err := json.Unmarshal(Request, req)
// 	if err != nil {
// 		return err
// 	}
// 	if req.SenderId != c.user_id {
// 		return ErrSenderIdNotMatch
// 	}

// 	ctx := context.WithValue(context.Background(), "serviceId", req.SessionType)
// 	_, err = c.Server.RpcRouterHandler.SendMessage(ctx, req)

//		return err
//	}
func (c *Client) pingHandler() error {
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongwait))
	return c.writePongMessage()
}

func (c *Client) writePongMessage() error {
	if c.closed.Load() {
		return nil
	}

	// make sure for security
	c.w.Lock()
	defer c.w.Unlock()

	err := c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}
	return c.Conn.WriteMessage(websocket.PongMessage, nil)
}

func (c *Client) Write() {
	ticker := time.NewTicker(54)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.MessageChan:
			c.Conn.SetWriteDeadline(time.Now().Add(10))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)
			n := len(c.MessageChan)
			for i := 0; i < n; i++ {
				w.Write([]byte("\n"))
				w.Write(<-c.MessageChan)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
