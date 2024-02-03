package gateway

import "sync"

type Usermap struct {
	m sync.Map
}

func newUserMap() *Usermap {
	return &Usermap{}
}

func (u *Usermap) Get(key string) ([]*Client, bool, bool) {
	allClients, userExited := u.m.Load(key)
	if userExited {
	}
}
