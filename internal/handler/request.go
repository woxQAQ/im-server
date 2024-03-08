package handler

import (
	"github.com/goccy/go-json"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/msg"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/pb"
)

type MsgData struct {
	CurMsgId    int64          `json:"cur_msg_id"`
	PreMsgId    int64          `json:"pre_msg_id"`
	SendTime    int64          `json:"send_time"`
	Content     string         `json:"content"`
	ContentType pb.ContentType `json:"contentType"`
}

type SendRequest struct {
	SenderId   string  `json:"senderId" validate:"required"`
	ReceiverId string  `json:"receiverId"`
	GroupId    string  `json:"groupId"`
	Content    MsgData `json:"content"`
	MsgType    string  `json:"MsgType"`
}

func (r *SendRequest) String() string {
	var tReq SendRequest
	tReq.ReceiverId = r.ReceiverId
	tReq.SenderId = r.SenderId
	tReq.Content = r.Content
	tReq.GroupId = r.GroupId
	tReq.MsgType = r.MsgType
	jsonData, _ := json.Marshal(tReq)
	return string(jsonData)
}

func getRPCReq(data *SendRequest, req *msg.SendMessageReq) {
	req.SenderId = data.SenderId
	req.ReceiverId = data.ReceiverId
	req.GroupId = data.GroupId
	req.Content.ContentType = data.Content.ContentType
	req.Content.Content = data.Content.Content
	req.Content.SendTime = data.Content.SendTime
	req.Content.CurMsgId = data.Content.CurMsgId
	req.Content.PreMsgId = data.Content.PreMsgId
	req.MsgType = data.MsgType
}
