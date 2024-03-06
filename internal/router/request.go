package router

import (
	"github.com/goccy/go-json"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/msg"
)

type MsgData struct {
	Content string `json:"content"`
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

func (r *SendRequest) toMsgData() (*msg.MsgData, error) {
	contentData, err := json.Marshal(r.Content)
	if err != nil {
		return nil, err
	}
	return &msg.MsgData{
		Content: string(contentData),
	}, nil
}