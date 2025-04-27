package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/pkg/response"
)

// ListPostComments 获取帖子评论
func ListPostComments(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Param("id")
	_, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 获取分页参数
	_, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	_, _ = strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// TODO: 实现获取帖子评论功能
	response.NotImplemented(c, "功能未实现")
}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 绑定请求参数
	var req struct {
		PostID   uint   `json:"post_id" binding:"required"`
		Content  string `json:"content" binding:"required"`
		ParentID uint   `json:"parent_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// TODO: 实现创建评论功能
	response.NotImplemented(c, "功能未实现")
}

// UpdateComment 更新评论
func UpdateComment(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取评论ID
	commentIDStr := c.Param("id")
	_, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的评论ID")
		return
	}

	// 绑定请求参数
	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// TODO: 实现更新评论功能
	response.NotImplemented(c, "功能未实现")
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	// 从上下文中获取用户ID和角色
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}
	userRole, _ := c.Get("userRole")
	_ = userRole.(string) == "admin"

	// 获取评论ID
	commentIDStr := c.Param("id")
	_, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的评论ID")
		return
	}

	// TODO: 实现删除评论功能
	response.NotImplemented(c, "功能未实现")
}

// LikeComment 点赞评论
func LikeComment(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取评论ID
	commentIDStr := c.Param("id")
	_, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的评论ID")
		return
	}

	// TODO: 实现点赞评论功能
	response.NotImplemented(c, "功能未实现")
}

// UnlikeComment 取消点赞评论
func UnlikeComment(c *gin.Context) {
	// 从上下文中获取用户ID
	_, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取评论ID
	commentIDStr := c.Param("id")
	_, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的评论ID")
		return
	}

	// TODO: 实现取消点赞评论功能
	response.NotImplemented(c, "功能未实现")
}
