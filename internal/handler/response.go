package handler

type ResponseData struct {
	ClientMsgId string `json:"clientMsgId"`
	ServerMsgId string `json:"serverMsgId"`
	SendTime    string `json:"sendTime"`
}

type SendResponse struct {
	ErrCode int          `json:"errCode"`
	ErrMsg  string       `json:"errMsg"`
	ErrDlt  string       `json:"errDlt"`
	Data    ResponseData `json:"data"`
}
