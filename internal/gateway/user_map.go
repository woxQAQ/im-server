package gateway

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

// Usermap is to store a goroutine-safe key-value bind
// bind structure like [ClientId,[]*Clients]
type Usermap struct {
	m sync.Map
}

// newUserMap create a new UserMap
func newUserMap() *Usermap {
	return &Usermap{}
}

// Get fetch k-v pair, with clients with clientId and platformId
// the main return value is the []*Clients correspinding to the key and platformId
// the bool values, first is which the clients exited
// and second is which the clients correspinding to the platformId exited or not
func (u *Usermap) Get(key string, platformId string) ([]*Client, bool, bool) {
	allClients, userExited := u.m.Load(key)
	if userExited {
		var clients []*Client
		for _, client := range allClients.([]*Client) {
			if client.PlatformId == platformId {
				clients = append(clients, client)
			}
		}
		if len(clients) > 0 {
			return clients, userExited, true
		}
		return clients, userExited, false
	}
	return nil, userExited, false
}

// Set store client into Usermap
// it may insert to a exited slice or create a new k-v pair
// NOTE: when shoule Set to be called? only when a client
// wants to be registered to server
func (u *Usermap) Set(key string, v *Client) {
	allClients, userExited := u.m.Load(key)
	if userExited {
		zap.S().Debug(context.Background(), "user exited", "userid", key, "client", v)
		oldClients := allClients.([]*Client)
		oldClients = append(oldClients, v)
		u.m.Store(key, oldClients)
	} else {
		zap.S().Debug(context.Background(), "user not exited", "userid", key, "client", v)
		var clients []*Client
		clients = append(clients, v)
		u.m.Store(key, clients)
	}
}

func (u *Usermap) Delete(key string, remoteAddr string) bool {
	allClients, userExited := u.m.Load(key)
	if userExited {
		oldClients := allClients.([]*Client)
		var newClients []*Client
		for _, client := range oldClients {
			if client.Conn.RemoteAddr().String() != remoteAddr {
				newClients = append(newClients, client)
			}
		}
		if len(newClients) == 0 {
			u.m.Delete(key)
			return true
		} else {
			u.m.Store(key, newClients)
			return false
		}
	}
	return userExited
}

func (u *Usermap) deleteAll(key string) {
	u.m.Delete(key)
}

func (u *Usermap) GetAll(key string) ([]*Client, bool) {
	allClients, ok := u.m.Load(key)
	if ok {
		return allClients.([]*Client), ok
	}
	return nil, ok
}
