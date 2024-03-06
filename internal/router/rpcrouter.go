package router

import (
	"context"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/msg"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/pb"
	"github.com/woxQAQ/im-service/pkg/common/errs"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/protobuf/proto"
)

var reqPool = sync.Pool{
	New: func() any {
		return new(SendRequest)
	},
}

func resetReq(r *SendRequest) *SendRequest {
	*r = SendRequest{}
	r.Content = MsgData{}
	return r
}

func getReq() *SendRequest {
	return resetReq(reqPool.Get().(*SendRequest))
}

func freeReq(req *SendRequest) {
	reqPool.Put(req)
}

//type handler func(context context.Context, data *Request) ([]byte,error)

type RpcRouterHandler interface {
	SendMessage(context context.Context, data *SendRequest) ([]byte, error)
}

var _ RpcRouterHandler = (*Handler)(nil)

type Handler struct {
	msg      msg.Msg
	validate *validator.Validate
	reqPool  sync.Pool
}

func newHandler(zrpcConf zrpc.RpcClientConf) *Handler {
	return &Handler{
		validate: validator.New(),
		msg:      msg.NewMsg(zrpc.MustNewClient(zrpcConf)),
	}
}

func (h *Handler) SendMessage(ctx context.Context, data *SendRequest) ([]byte, error) {
	if err := h.validate.Struct(data); err != nil {
		return nil, err
	}

	msgData, err := data.toMsgData()
	if err != nil {
		return nil, err
	}
	req := msg.SendMessageReq{}
	switch ctx.Value("serviceId") {
	case ServiceGroupChat:
		if data.GroupId == "" {
			return nil, errs.ErrGroupIdNotFound
		}
		req = msg.SendMessageReq{
			SenderId:    data.SenderId,
			GroupId:     data.GroupId,
			Content:     msgData,
			ContentType: pb.MessageType(data.ContentType),
		}
	case ServiceSingleChat:
		if data.ReceiverId == "" {
			return nil, errs.ErrRecvIdNotFound
		}
		req = msg.SendMessageReq{
			SenderId:    data.SenderId,
			ReceiverId:  data.ReceiverId,
			Content:     msgData,
			ContentType: pb.MessageType(data.ContentType),
		}
	default:
		return nil, errs.ErrArgumentErr
	}

	resp, err := h.msg.SendMsg(ctx, &req)
	if err != nil {
		return nil, err
	}
	res, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return res, nil
}
