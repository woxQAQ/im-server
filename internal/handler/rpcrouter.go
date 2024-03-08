package handler

//var reqPool = sync.Pool{
//	New: func() any {
//		return new(SendRequest)
//	},
//}
//
//func resetReq(r *SendRequest) *SendRequest {
//	*r = SendRequest{}
//	r.Content = MsgData{}
//	return r
//}
//
//func getReq() *SendRequest {
//	return resetReq(reqPool.Get().(*SendRequest))
//}
//
//func freeReq(req *SendRequest) {
//	reqPool.Put(req)
//}

//type handler func(context context.Context, data *Request) ([]byte,error)

//var _ RpcRouterHandler = (*DefaultRpcHandler)(nil)
//
//type DefaultRpcHandler struct {
//}
//
//func newHandler(zrpcConf zrpc.RpcClientConf) *DefaultRpcHandler {
//	return &DefaultRpcHandler{}
//}
