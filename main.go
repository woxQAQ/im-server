package main

import (
	"flag"
)

var gatewayConfigPath = flag.String("f", "config/gateway.yaml", "gateway config file")

func main() {
	flag.Parse()

}
