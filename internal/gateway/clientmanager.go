package gateway

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

type ClientMgr struct {
	// ClientMap stores the Client register to the manager
	ClientMap sync.Map

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
}

func NewClientManager() *ClientMgr {
	pool, err := ants.NewPool(1024, ants.WithPreAlloc(true), ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}
	return &ClientMgr{
		ClientMap:      sync.Map{},
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
			m.ClientMap.Store(client.ClientId, client)

		case client := <-m.unregisterChan:
			if _, ok := m.ClientMap.Load(client.ClientId); ok {
				m.ClientMap.Delete(client.ClientId)
				close(client.MessageChan)
			}

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
}
