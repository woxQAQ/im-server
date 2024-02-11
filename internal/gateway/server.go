package gateway

type Server struct {
	wsServer    *WsServer
	rpcPort     int
	monitorPort int
}

func newServer(rpcPort int, monitorPort int, wsServer *WsServer) *Server {
	return &Server{
		wsServer:    wsServer,
		rpcPort:     rpcPort,
		monitorPort: monitorPort,
	}
}
