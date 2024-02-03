package gateway

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	// ClientId is the Identify of a Client
	ClientId string

	// UserId is the Id of the user of Client
	// UserId is different from ClientId
	// UserId identify a real user of the Im service
	// ClientId, on the other hand, can be called "sessionId"
	UserId string

	// Conn is the connection that Client connect to server
	Conn *websocket.Conn

	ConnectTime uint64

	// Mgr is the ClientManager that Client register to.
	// the Mgr is used to manager the client,
	// including sending and reading message...
	Mgr *ClientMgr

	// MessageChan is used to send message to the ClientManager
	MessageChan chan []byte
}

func NewClient(clientId string, conn *websocket.Conn, mgr *ClientMgr) *Client {
	return &Client{
		ClientId:    clientId,
		Conn:        conn,
		Mgr:         mgr,
		ConnectTime: uint64(time.Now().Unix()),
	}
}

func (c *Client) Read() {
	defer func() {
		c.Mgr.unregisterChan <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(1024)
	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zap.S().Infof("error: %v", err)
			}
			break
		}
		// remove '\n' to be space, and remove the message's edge's space
		// is used to process multiline text to be one line
		data = bytes.TrimSpace(bytes.Replace(data, []byte("\n"), []byte(" "), -1))
		c.Mgr.receivedChan <- data
	}
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
