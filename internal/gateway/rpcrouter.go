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

type MsgData struct {
	Content string `json:"content"`
}

type ResponseData struct {
	ClientMsgId string `json:"clientMsgId"`
	ServerMsgId string `json:"serverMsgId"`
	SendTime    string `json:"sendTime"`
}

type SendRequest struct {
	SenderId       string  `json:"senderId" validate:"required"`
	ReceiverId     string  `json:"receiverId"`
	SenderPlatform int32   `json:"senderPlatform"`
	GroupId        string  `json:"groupId"`
	Content        MsgData `json:"content"`
	ContentType    int     `json:"contentType"`
	SessionType    int     `json:"sessionType"`
	SendTime       int64   `json:"sendTime"`
}

type SendResponse struct {
	ErrCode int          `json:"errCode"`
	ErrMsg  string       `json:"errMsg"`
	ErrDlt  string       `json:"errDlt"`
	Data    ResponseData `json:"data"`
}

func (r *SendRequest) String() string {
	var tReq SendRequest
	tReq.ReceiverId = r.ReceiverId
	tReq.SenderId = r.SenderId
	tReq.Content = r.Content
	tReq.ContentType = r.ContentType
	tReq.GroupId = r.GroupId
	tReq.SenderPlatform = r.SenderPlatform
	tReq.SendTime = r.SendTime
	tReq.SessionType = r.SessionType
	jsonData, _ := json.Marshal(tReq)
	return string(jsonData)
}

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
	msgRpcClient pb.MsgClient
	validate     *validator.Validate
}

type Config struct {
	MsgRpc zrpc.RpcClientConf
}

func newHandler(configFile string, validate *validator.Validate) *Handler {
	var config Config
	conf.MustLoad(configFile, &config)
	conn := zrpc.MustNewClient(config.MsgRpc)
	return &Handler{
		pb.NewMsgClient(conn.Conn()),
		validate,
	}
}

func (r *SendRequest) toMsgData() (*msg.MsgData, error) {
	contentData, err := json.Marshal(r.Content)
	if err != nil {
		return nil, err
	}
	return &msg.MsgData{
		Content: string(contentData),
	}, nil
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
			return nil, ErrGroupIdNotFound
		}
		req = msg.SendMessageReq{
			SenderId:    data.SenderId,
			GroupId:     data.GroupId,
			Content:     msgData,
			ContentType: pb.MessageType(data.ContentType),
		}
	case ServiceSingleChat:
		if data.ReceiverId == "" {
			return nil, ErrRecvIdNotFound
		}
		req = msg.SendMessageReq{
			SenderId:    data.SenderId,
			ReceiverId:  data.ReceiverId,
			Content:     msgData,
			ContentType: pb.MessageType(data.ContentType),
		}
	default:
		return nil, ErrArgumentErr
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
