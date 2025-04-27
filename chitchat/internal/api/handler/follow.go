package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/pkg/response"
)

// FollowUser 关注用户
func FollowUser(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取目标用户ID
	targetIDStr := c.Param("id")
	_, err := strconv.ParseUint(targetIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID")
		return
	}

	// TODO: 实现关注用户功能
	response.NotImplemented(c, "功能未实现")
}

// UnfollowUser 取消关注用户
func UnfollowUser(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取目标用户ID
	targetIDStr := c.Param("id")
	_, err := strconv.ParseUint(targetIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID")
		return
	}

	// TODO: 实现取消关注用户功能
	response.NotImplemented(c, "功能未实现")
}

// ListFollowers 获取粉丝列表
func ListFollowers(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取分页参数
	_, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	_, _ = strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// TODO: 实现获取粉丝列表功能
	response.NotImplemented(c, "功能未实现")
}

// ListFollowing 获取关注列表
func ListFollowing(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取分页参数
	_, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	_, _ = strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// TODO: 实现获取关注列表功能
	response.NotImplemented(c, "功能未实现")
}
