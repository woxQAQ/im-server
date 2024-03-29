package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/pb"
	pb2 "github.com/woxQAQ/im-service/internal/rpc/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var (
	ErrReqDataEmpty = errors.New("Data in the Request is empty")
)

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgLogic) SendMsg(in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	// todo: add your logic here and delete this line
	//1. 根据会话id获取 preid
	var (
		resp pb.SendMessageResp
	)

	sessionId := svc.Session{
		Sender:   in.SenderId,
		Receiver: in.ReceiverId,
	}

	// 确保session不存在

	storedPid, ok := l.svcCtx.SinglePreMap[sessionId]
	if !ok {
		// map 内不存在
		// 可能存在以下几种情况：
		// 会话刚刚初始化，此时 preid = 0
		// 服务器重启丢失数据，从数据库中获取最新一条消息的 preid
		if in.Content.PreMsgId == 0 {
			// 会话初始化
			// 直接赋值
			l.svcCtx.SinglePreMap[sessionId] = 0
			resp.Base.ErrCode = 0
			resp.CurMsgId = 1
			seq, err := l.svcCtx.SeqServer.GetSessionSeq(context.Background(), &pb2.GetSeqRequest{
				Type: func(messageType pb.MessageType) pb2.OperationType {
					switch messageType {
					case pb.MessageType_MESSAGE_TYPE_SINGLE:
						return pb2.OperationType_OPERATION_TYPE_SESSION
					case pb.MessageType_MESSAGE_TYPE_GROUP:
						return pb2.OperationType_OPERATION_TYPE_GID
					default:
						return -1
					}
				}(in.MsgType),
				UserId_1: in.SenderId,
				UserId_2: in.ReceiverId,
				GroupId:  in.GroupId,
			})

			if err != nil {
				return nil, err
			}
			return &resp, nil
		} else {
			presistPid, err := l.svcCtx.SenderDtl.FindBySenderAndSession(l.ctx, in.SenderId, sessionId)
			if err != nil {
				resp.Base.ErrCode = 11
				resp.Base.ErrMsg = err.Error()
				return &resp, nil
			}

			if presistPid != in.Content.PreMsgId {
				// preid 不一致，返回错误
				resp.Base.ErrCode = 1011
				err = errors.New(fmt.Sprintf("Your preid %v is error, expected preid is contained in this message", in.Content.PreMsgId))
				resp.Base.ErrMsg = err.Error()
				resp.CurMsgId = presistPid
				return &resp, nil
			}
		}
	}

	// 校验preid
	diff := storedPid - in.Content.PreMsgId
	if diff == 0 {
		// OK
		// 说明发送方时序正确
		// 入库
		// todo: 接收方的时序呢？
		//  接收方如果同时也在发送消息，这完全取决于服务器的接收顺序，是否会产生竞争？
		//  preid 只与发送方强相关，preid不会竞争

		//  考虑用会话自增的序列号，每个会话拥有唯一的版本空间，会话也是一个获取seqid 的实体
		//  每一条发送到服务器的消息，都赋予会话唯一的seqid，会话的各端各自维护自己的seqid
		//  推送服务器定时向用户推送拉取通知，客户端收到后就进行拉取
		//  拉取的内容取决于拉取请求到达服务器后，服务器缓存的会话全局seqid和客户端发来的seqid差值所包括的消息
		//  每一份需要被拉取的消息都携带会话的序列号

		//  在线用户通过websocket/长连接便可拉取，而离线用户则考虑使用离线推送，此处不表

		//  因此对于每个客户端，他们消息的排序仅取决于seqid，对于前端而言也就只需要按照seqid进行排序显示即可

		//  用户是否使用一个全局的序列号？使用，但是不在发送阶段使用，而是接收阶段

		// todo 为消息分配一个msg_id
		l.svcCtx.SeqServer.GetSessionSeq(context.Background(), &pb2.GetSeqRequest{})
		resp = pb.SendMessageResp{
			Base: &pb.ResponseBase{
				ErrCode: 0,
			},
			CurMsgId: in.Content.CurMsgId,
		}
	}
	return &resp, nil
}
