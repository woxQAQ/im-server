package logic

import (
	"bytes"
	"context"
	"encoding"
	"encoding/json"
	"errors"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	pb "github.com/woxQAQ/im-service/internal/rpc/rpc/pb"
	"github.com/woxQAQ/im-service/pkg/common/convert"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSessionSeqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSessionSeqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionSeqLogic {
	return &GetSessionSeqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type seq struct {
	CurSeq int64 `json:"cur_seq"`

	MaxSeq int64 `json:"max_seq"`
}

var _ encoding.BinaryMarshaler = new(seq)
var _ encoding.BinaryUnmarshaler = new(seq)

func (s *seq) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(byte(s.CurSeq))
	buf.WriteByte('\t')
	buf.WriteByte(byte(s.MaxSeq))
	return buf.Bytes(), nil
}

func (s *seq) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, s)
	if err == nil {
		return nil
	}

	res := bytes.Split(data, []byte{'\t'})
	if len(res) != 2 {
		return errors.New("Not valid format")
	}

	s.CurSeq, err = convert.BytesToInt64(res[0])
	if err != nil {
		return err
	}
	s.MaxSeq, err = convert.BytesToInt64(res[1])
	if err != nil {
		return err
	}
	return nil
}

// GetSessionSeq 通常用于当消息服务器内存内没有对应会话的数据时调用
func (l *GetSessionSeqLogic) GetSessionSeq(in *pb.GetSeqRequest) (*pb.GetSeqResponse, error) {

	switch in.Type {
	case pb.OperationType_OPERATION_TYPE_SESSION:
		res, err := l.svcCtx.SessionSeqModel.FindOne(context.Background(), in.SessionId)
		if err != nil {
			return nil, err
		}
		return &pb.GetSeqResponse{
			Base: &pb.RespBase{
				ErrCode: 0,
			},
			MaxSeq: res.MaxSeq,
		}, nil
	case pb.OperationType_OPERATION_TYPE_GROUP:
		res, err := l.svcCtx.GroupSequenceModel.FindOne(context.Background(), in.GroupId)
		if err != nil {
			return nil, err
		}
		return &pb.GetSeqResponse{
			Base: &pb.RespBase{
				ErrCode: 0,
			},
			MaxSeq: res.MaxSeq,
		}, nil
	default:
		return nil, nil
	}
}
