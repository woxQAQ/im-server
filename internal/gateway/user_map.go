package gateway

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

// UserMap is to store a goroutine-safe key-value bind
// bind structure like [ClientId,[]*Clients]
type UserMap struct {
	m sync.Map
}

// newUserMap create a new UserMap
func newUserMap() *UserMap {
	return &UserMap{}
}

// Get fetch k-v pair, with clients with clientId and platformId
// the main return value is the []*Clients correspinding to the key and platformId
// the bool values, first is which the clients exited
// and second is which the clients correspinding to the platformId exited or not
func (u *UserMap) Get(userId string, platformId int) ([]*Client, bool, bool) {
	allClients, userExited := u.m.Load(userId)
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

// Set store client into UserMap
// it may insert to a exited slice or create a new k-v pair
// NOTE: when shoule Set to be called? only when a client
// wants to be registered to server
func (u *UserMap) Set(userId string, client *Client) {
	allClients, userExited := u.m.Load(userId)
	if userExited {
		zap.S().Debug(context.Background(), "user exited", "userid", userId, "client", client)
		oldClients := allClients.([]*Client)
		oldClients = append(oldClients, client)
		u.m.Store(userId, oldClients)
	} else {
		zap.S().Debug(context.Background(), " user not exited, ", "userid: ", userId, " client ", client)
		var clients []*Client
		clients = append(clients, client)
		u.m.Store(userId, clients)
	}
}

func (u *UserMap) Delete(key string, remoteAddr string) bool {
	allClients, userExited := u.m.Load(key)
	if userExited {
		oldClients := allClients.([]*Client)
		var newClients []*Client
		for _, client := range oldClients {
			if client.Context.RemoteIP() != remoteAddr {
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

func (u *UserMap) deleteAll(key string) {
	u.m.Delete(key)
}

func (u *UserMap) GetAll(key string) ([]*Client, bool) {
	allClients, ok := u.m.Load(key)
	if ok {
		return allClients.([]*Client), ok
	}
	return nil, ok
}
