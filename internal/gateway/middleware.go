package gateway

import (
	"github.com/golang-jwt/jwt/v4"
	jwtTools "github.com/woxQAQ/im-service/pkg/common/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var secretkey = "bZ87g@fcW93W8Y!uuK^nSHPAhgJeWKUc"

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


func jwtHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authentication")
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"errCode": 13,
				"errMsg":  "you have no Authentication field in your request header",
			})
			context.Abort()
			return
		}

		parts := strings.Split(tokenString, ".")
		if len(parts) != 3 {
			context.JSON(http.StatusUnauthorized, gin.H{
				//todo errcode define
				"errCode": 13,
				"errMsg":  "your token format is error",
			})
			context.Abort()
			return
		}

		j := jwtTools.Jwt{
			SignedKeys: []byte(secretkey),
		}

		claims, err := j.ParseToken(tokenString)
		
		if err != nil {
			if err == jwt.ErrTokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"errCode": 13,
					"errMsg": "your token has been expired",
				})
				context.Abort()
				return
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"errCode": 13,
				"errMsg": err.Error(),
			})
			context.Abort()
			return 
		}

		context.Set("Claims", claims)

		context.Next()
	}
}
