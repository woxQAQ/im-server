package gateway

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ws *WsServer) registerRouter(r *gin.Engine) {
	r.Any("/", func(ctx *gin.Context) {
		wsHandler(ctx, ws)
	})
}

func wsHandler(ctx *gin.Context, ws *WsServer) {
	client := ws.clientManager.clientPool.Get().(*Client)
	if ws.clientManager.onlineUserConnNum.Load() >= ws.wsMaxConnNum {
		ctx.AbortWithError(http.StatusBadRequest, ErrConnOverMaxNumLimit)
		return
	}

	platformIdStr, exited := ctx.GetQuery("platformId")
	if !exited {
		ctx.AbortWithError(http.StatusBadRequest, ErrArgumentErr)
		return
	}
	platformId, err := strconv.Atoi(platformIdStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrArgumentErr)
	}

	userId, exited := ctx.GetQuery("userId")
	if !exited {
		ctx.AbortWithError(http.StatusBadRequest, ErrArgumentErr)
		return
	}

	token := ctx.GetHeader("Authentication")

	//token, exited := ctx.GetQuery("token")
	//if !exited {
	//	ctx.AbortWithError(http.StatusBadRequest, ErrArgumentErr)
	//	return
	//}

	conn, err := ws.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	client.resetClient(conn, ws.clientManager, userId, token, platformId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrWebsockerUpgrade)
	}
	ws.clientManager.registerChan <- client
	ws.clientManager.goroutinePool.Submit(client.Read)
}
