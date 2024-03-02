package handler

import (
	"net/http"

	"github.com/woxQAQ/im-service/internal/api/ws_gateway/internal/logic"
	"github.com/woxQAQ/im-service/internal/api/ws_gateway/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/ws_gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendMsgLogic(r.Context(), svcCtx)
		resp, err := l.SendMsg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
