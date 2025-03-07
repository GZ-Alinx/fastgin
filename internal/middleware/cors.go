package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			if Logger != nil {
				Logger.Debug("处理OPTIONS请求",
					zap.String("path", c.Request.URL.Path),
					zap.String("origin", c.Request.Header.Get("Origin")))
			}
			c.AbortWithStatus(204)
			return
		}

		if Logger != nil {
			Logger.Debug("CORS中间件",
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("origin", c.Request.Header.Get("Origin")))
		}
		c.Next()
	}
}
