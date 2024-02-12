package MQ

import (
	"os"
	"strings"
)

func readAddrFromEnv(fallback []string) []string {
	envAddr := os.Getenv("KAFKA_ADDR")
	envPort := os.Getenv("KAFKA_PORT")
	if envAddr != "" && envPort != "" {
		address := strings.Split(envAddr, ",")
		for i, addr := range address {
			address[i] = addr + ":" + envPort
		}
		return address
	}
	return fallback
}
