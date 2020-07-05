package middleware

import (
	"go-stock/app/common"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recovery捕获所有panic，并且返回错误信息
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := Context{Ctx: c}
		defer func() {
			if err := recover(); err != nil {
				logger.Error("recover error:", err)
				logger.Warn("debug stack warn:", string(debug.Stack()))
				ctx.Response(500, common.SERVER_ERROR, nil)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
