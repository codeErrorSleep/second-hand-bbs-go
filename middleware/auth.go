package middleware

import (
	"second-hand-bbs-go/logging"
	"second-hand-bbs-go/pkg/token"
	"second-hand-bbs-go/utils/app"
	"second-hand-bbs-go/utils/e"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	logging.Info("验证token")
	return func(c *gin.Context) {
		appG := app.Gin{c}
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		if err != nil {
			appG.Response(e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		if ctx.ExpirationTime < time.Now().Unix() {
			appG.Response(e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, nil)
			return
		}
		c.Next()
	}
}
