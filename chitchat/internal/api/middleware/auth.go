package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/internal/utils"
	"github.com/lllllan02/chitchat/pkg/response"
)

// JWT 认证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization头中获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "请提供令牌")
			c.Abort()
			return
		}

		// Bearer token格式验证
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Unauthorized(c, "令牌格式错误")
			c.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Unauthorized(c, "无效的令牌")
			c.Abort()
			return
		}

		// 将用户信息保存到上下文中
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// 管理员权限中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户角色
		role, exists := c.Get("role")
		if !exists {
			response.Unauthorized(c, "未认证访问")
			c.Abort()
			return
		}

		// 检查是否为管理员
		if role != "admin" {
			response.Forbidden(c, "无权限访问")
			c.Abort()
			return
		}

		c.Next()
	}
}

// 版主权限中间件
func ModeratorAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户角色
		role, exists := c.Get("role")
		if !exists {
			response.Unauthorized(c, "未认证访问")
			c.Abort()
			return
		}

		// 检查是否为管理员或版主
		if role != "admin" && role != "moderator" {
			response.Forbidden(c, "无权限访问")
			c.Abort()
			return
		}

		c.Next()
	}
}
