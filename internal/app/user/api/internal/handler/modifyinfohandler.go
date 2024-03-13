package handler

import (
	"net/http"

	"github.com/woxQAQ/im-service/internal/api/api/internal/logic"
	"github.com/woxQAQ/im-service/internal/api/api/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewModifyInfoLogic(r.Context(), svcCtx)
		err := l.ModifyInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
