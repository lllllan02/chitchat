package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/pkg/response"
)

// ListNotifications 获取通知列表
func ListNotifications(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// TODO: 实现获取通知列表功能
	response.NotImplemented(c, "功能未实现")
}

// MarkNotificationAsRead 标记通知为已读
func MarkNotificationAsRead(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取通知ID
	notificationIDStr := c.Param("id")
	_, err := strconv.ParseUint(notificationIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的通知ID")
		return
	}

	// TODO: 实现标记通知为已读功能
	response.NotImplemented(c, "功能未实现")
}

// MarkAllNotificationsAsRead 标记所有通知为已读
func MarkAllNotificationsAsRead(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// TODO: 实现标记所有通知为已读功能
	response.NotImplemented(c, "功能未实现")
}
