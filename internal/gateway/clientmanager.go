package gateway

import (
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

type ClientMgr struct {
	// ClientMap stores the Client register to the manager
	ClientMap Usermap

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

	onlineNum atomic.Int64

	onlineUserConnNum atomic.Int64
}

func newClientManager() *ClientMgr {
	pool, err := ants.NewPool(1024, ants.WithPreAlloc(true), ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}
	return &ClientMgr{
		ClientMap:      *newUserMap(),
		receivedChan:   make(chan []byte, 1024),
		registerChan:   make(chan *Client),
		unregisterChan: make(chan *Client),
		goroutinePool:  pool,
	}
}

func (m *ClientMgr) Run() {
	for {
		select {
		case client := <-m.registerChan:
			m.RegisterClient(client)

		case client := <-m.unregisterChan:
			m.unregisterClient(client)

		case message := <-m.receivedChan:
			m.ClientMap.Range(func(connId, session any) bool {
				client := session.(*Client)
				select {
				case client.MessageChan <- message:
				default:
					close(client.MessageChan)
					m.ClientMap.Delete(connId)
				}
				return true
			})
		}
	}
}

func (m *ClientMgr) RegisterClient(client *Client) {
	// oldClients map中获取对应的client切片
	// userOk 标识map中是否有clientID对应的切片
	// clientOK 标识map中是否有 clientID 和 platformId 所对应的client切片
	_, userOk, clientOk := m.ClientMap.Get(client.UserId, client.PlatformId)
	// userMap中不存在client切片
	if !userOk {
		m.ClientMap.Set(client.UserId, client)
		m.onlineNum.Add(1)
		m.onlineUserConnNum.Add(1)
	} else {
		if clientOk {
			// 已有同平台的连接存在
			m.ClientMap.Set(client.ClientId, client)
			m.onlineUserConnNum.Add(1)
		} else {
			m.ClientMap.Set(client.UserId, client)
			m.onlineUserConnNum.Add(1)
		}
	}

	zap.S().Info("user: ", client.UserId,
		"arrived\nonline user number",
		m.onlineNum.Load(),
		"online user conn number: ",
		m.onlineUserConnNum.Load(),
	)
}

func (m *ClientMgr) unregisterClient(client *Client) {
	// TODO:
}
