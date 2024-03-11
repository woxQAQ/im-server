package logic

import (
	"bytes"
	"context"
	"encoding"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb/pb"
	"github.com/woxQAQ/im-service/pkg/common/convert"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
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

func (l *GetSessionSeqLogic) GetSessionSeq(in *pb.GetSeqRequest) (*pb.GetSeqResponse, error) {
	// todo: add your logic here and delete this line
	switch in.Type {
	case pb.OperationType_OPERATION_TYPE_SESSION:
		var res seq
		seqStr, err := l.svcCtx.Rds.Hget("SessionSeq", string(rune(in.SessionId)))
		if err != nil {
			return nil, err
		}
		if seqStr == "" {
			// redis 中不存在对应 session id 的键值对，查库
			// seqid 是需要入库的
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			sessionSeq, err := l.svcCtx.SessionSeqModel.FindOne(ctx, in.SessionId)

			if err == sqlc.ErrNotFound {
				// mysql 中也没有对应的session id的 seqid
				res.CurSeq = 1
				res.MaxSeq = int64(l.svcCtx.Config.Seq.Step)
				resBuf, err := res.MarshalBinary()
				if err != nil {
					return nil, err
				}

				l.svcCtx.Rds.Hset("SessionSeq", strconv.FormatInt(in.SessionId, 10), string(resBuf))
			} else {
				// 存在记录，将 MaxSeq写入Curseq
				res.CurSeq = sessionSeq.MaxSeq + 1
				res.MaxSeq = sessionSeq.MaxSeq + int64(l.svcCtx.Config.Seq.Step)
				resBuf, err := res.MarshalBinary()
				if err != nil {
					return nil, err
				}

				l.svcCtx.Rds.Hset("SessionSeq", strconv.FormatInt(in.SessionId, 10), string(resBuf))
			}

			return &pb.GetSeqResponse{
				Base: &pb.RespBase{
					ErrCode: 0,
				},
				CurSeq: res.CurSeq,
			}, nil
		}

		res.UnmarshalBinary([]byte(seqStr))
		if res.CurSeq+1 == res.MaxSeq {
			res.CurSeq += 1
			res.MaxSeq += int64(l.svcCtx.Config.Seq.Step)
		}
		writeback, _ := res.MarshalBinary()
		l.svcCtx.Rds.Hset("SessionSeq", strconv.FormatInt(in.SessionId, 10), string(writeback))
		return &pb.GetSeqResponse{
			Base: &pb.RespBase{
				ErrCode: 0,
			},
			CurSeq: res.CurSeq,
		}, nil
	case pb.OperationType_OPERATION_TYPE_GETUID:
		seqStr, err := l.svcCtx.Rds.Get("UserId")
		if err != nil {
			return nil, err
		}

		if seqStr == "" {
			// user id 生成器未初始化，或者未加载
			// 从mysql加载
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			userId, err := l.svcCtx.SessionSeqModel.FindOne(ctx, in.SessionId)

		}
	default:
		return &pb.GetSeqResponse{
			Base: &pb.RespBase{
				ErrCode: 11,
				ErrMsg:  "your session type field is error",
			},
		}, errors.New("UnKnownSessionType")
	}

}
