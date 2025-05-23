package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/pkg/logger"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Info("[GIN] %s | %3d | %13v | %15s | %s",
			reqMethod,
			statusCode,
			latencyTime,
			clientIP,
			reqUri,
		)
	}
}
