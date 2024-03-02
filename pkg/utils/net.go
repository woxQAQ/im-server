package utils

import (
	"github.com/pkg/errors"
	net2 "k8s.io/utils/net"
	"net"
)

func GetLocalAddr() (net.Addr, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range interfaces {
		if iface.Name == "en0" {
			addrs, err := iface.Addrs()
			if err != nil {
				return nil, err
			}
			for _, addr := range addrs {
				if net2.IsIPv4(net.ParseIP(addr.String())) {
					return addr, nil
				}
			}
		}
	}
	return nil, errors.New("No available ip address")
}
