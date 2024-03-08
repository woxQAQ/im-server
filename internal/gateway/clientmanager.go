package gateway

import (
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

type ClientMgr struct {
	// ClientMap stores the Client register to the manager
	ClientMap *UserMap

	// registerChan receive clients which going to
	// register to this manager
	registerChan chan *Client

	// unregisterChan same as the registerChan
	// receive clients which is going to unregister
	unregisterChan chan *Client

	// receivedChan is used to receive message from registered client
	receivedChan chan []byte

	// goroutinePool is a goroutine goroutinePool which is used to reuse goroutine
	goroutinePool *ants.Pool

	// clientPool is used to reuse clients objects
	clientPool sync.Pool

	// onlineNum is the number of users of server
	onlineNum atomic.Int64

	// onlineUserConnNum is the number of connection of server
	// NOTE: onlineNum different from onlineNum because
	// one user may login several terminal at a same time
	// different terminal should be handlerred as different clients
	onlineUserConnNum atomic.Int64
}

func newClientManager() *ClientMgr {
	pool, err := ants.NewPool(
		1024,
		ants.WithPreAlloc(true),
		ants.WithNonblocking(true),
	)
	if err != nil {
		panic(err)
	}
	return &ClientMgr{
		ClientMap:      newUserMap(),
		receivedChan:   make(chan []byte, 1024),
		registerChan:   make(chan *Client),
		unregisterChan: make(chan *Client),
		goroutinePool:  pool,
		clientPool: sync.Pool{
			New: func() any {
				return new(Client)
			},
		},
	}
}

func (m *ClientMgr) Run(done chan struct{}) error {
	for {
		select {
		case <-done:
			return nil
		case client := <-m.registerChan:
			m.RegisterClient(client)
		case client := <-m.unregisterChan:
			m.unregisterClient(client)
		}
	}
}

func (m *ClientMgr) RegisterClient(client *Client) {
	_, userOk, clientOk := m.ClientMap.Get(client.UserId, client.PlatformId)
	// There is No key "user_id" in the ClientMap.
	// It indicates that the user is login to the server for the first time
	if !userOk {
		m.ClientMap.Set(client.UserId, client)
		m.onlineNum.Add(1)
		m.onlineUserConnNum.Add(1)
	} else {
		if clientOk {
			// this terminal has been logined
			// TODO: return a REPEAT ERROR
			m.ClientMap.Set(client.ClientId, client)
			m.onlineUserConnNum.Add(1)
		} else {
			// this terminal has not been logined
			m.ClientMap.Set(client.UserId, client)
			m.onlineUserConnNum.Add(1)
		}
	}

	zap.S().Info("user: ", client.UserId,
		" arrived\nonline user number: ",
		m.onlineNum.Load(),
		"\nonline user conn number: ",
		m.onlineUserConnNum.Load(),
	)
}

func (m *ClientMgr) unregisterClient(client *Client) {
	defer m.clientPool.Put(client)
	if isDel := m.ClientMap.Delete(client.ClientId, client.Conn.RemoteAddr().String()); isDel {
		m.onlineNum.Add(-1)
	}
	m.onlineUserConnNum.Add(-1)
	zap.S().Info("user offline! ", "close Error:", client.CloseErr)
}
