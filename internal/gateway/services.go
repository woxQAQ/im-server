package gateway

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/msg"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/protobuf/proto"
	"sync"
)

type Request struct {
	SenderID  string `json:"send_id" validate:"required"`
	RecvID    string `json:"recv_id" validate:"required"`
	Token     string `json:"token"`
	ServiceId string `json:"service_id" validate:"required"`
	Data      []byte
}

func (r *Request) String() string {
	var tReq Request
	tReq.Token = r.Token
	tReq.SenderID = r.SenderID
	tReq.RecvID = r.RecvID
	tReq.ServiceId = r.ServiceId
	jsonData, _ := json.Marshal(tReq)
	return string(jsonData)
}

var reqPool = sync.Pool{
	New: func() any {
		return new(Request)
	},
}

func resetReq(r *Request) *Request {
	r.SenderID = ""
	r.RecvID = ""
	r.Token = ""
	r.ServiceId = ""
	r.Data = nil
	return r
}

func getReq() *Request {
	return resetReq(reqPool.Get().(*Request))
}

func freeReq(req *Request) {
	reqPool.Put(req)
}

//type handler func(context context.Context, data *Request) ([]byte,error)

type ServiceHandler interface {
	SendMessage(context context.Context, data *Request) ([]byte, error)
}

var _ ServiceHandler = (*Handler)(nil)

type Handler struct {
	msgRpcClient pb.MsgClient
	validate     *validator.Validate
}

func newHandler(validate *validator.Validate) *Handler {
	var config struct {
		MsgRpc zrpc.RpcClientConf
	}
	conf.MustLoad("gateway.yaml", config)
	conn := zrpc.MustNewClient(config.MsgRpc)
	return &Handler{
		pb.NewMsgClient(conn.Conn()),
		validate,
	}
}

func (h *Handler) SendMessage(ctx context.Context, data *Request) ([]byte, error) {
	msgData := msg.MsgData{}
	if err := proto.Unmarshal(data.Data, &msgData); err != nil {
		return nil, err
	}
	if err := h.validate.Struct(data); err != nil {
		return nil, err
	}
	req := msg.SendMessageReq{
		Data: &msgData,
	}
	resp, err := h.msgRpcClient.SendMsg(ctx, &req)
	if err != nil {
		return nil, err
	}
	res, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return res, nil
}
