package gateway

import (
	"github.com/woxQAQ/im-service/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")

		ctx.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar",
		)
		ctx.Header("Access-Control-Max-Age", "172800")
		ctx.Header("Access-Control-Allow-Credentials", "false")
		ctx.Header("Content-Type", "application/json")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatusJSON(http.StatusOK, "Options Request!")
			return
		}
		ctx.Next()
	}
}

func domainIntercept() gin.HandlerFunc {
	localIp, _ := utils.GetLocalAddr()
	return func(ctx *gin.Context) {
		if ctx.Request.Host == "localhost" || ctx.Request.Host == "127.0.0.1" || ctx.Request.Host == localIp.String() {
			token := ctx.GetHeader("Authentication")

		}
	}
}
