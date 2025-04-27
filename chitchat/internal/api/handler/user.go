package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/service"
	"github.com/lllllan02/chitchat/internal/utils"
	"github.com/lllllan02/chitchat/pkg/response"
)

// 初始化帖子服务
var postService = service.NewPostService()

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 查询用户信息
	user, err := userService.GetUserByID(userID.(uint))
	if err != nil {
		response.ServerError(c, "获取用户信息失败")
		return
	}

	response.Success(c, user)
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 绑定请求参数
	var req model.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 查询用户信息
	user, err := userService.GetUserByID(userID.(uint))
	if err != nil {
		response.ServerError(c, "获取用户信息失败")
		return
	}

	// 更新用户信息
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}

	if err := userService.UpdateUser(user); err != nil {
		response.ServerError(c, "更新用户信息失败")
		return
	}

	response.Success(c, user)
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 绑定请求参数
	var req model.PasswordChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 查询用户信息
	user, err := userService.GetUserByID(userID.(uint))
	if err != nil {
		response.ServerError(c, "获取用户信息失败")
		return
	}

	// 验证旧密码
	if !utils.CheckPasswordHash(req.OldPassword, user.PasswordHash) {
		response.BadRequest(c, "旧密码错误")
		return
	}

	// 加密新密码
	newPasswordHash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		response.ServerError(c, "密码加密失败")
		return
	}

	// 更新密码
	if err := userService.ChangePassword(user.ID, newPasswordHash); err != nil {
		response.ServerError(c, "修改密码失败")
		return
	}

	response.Success(c, "密码修改成功")
}

// GetUser 获取用户信息
func GetUser(c *gin.Context) {
	// 获取用户ID
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID")
		return
	}

	// 查询用户信息
	user, err := userService.GetUserByID(uint(userID))
	if err != nil {
		response.NotFound(c, "用户不存在")
		return
	}

	response.Success(c, user)
}

// ListUsers 获取用户列表
func ListUsers(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 查询用户列表
	users, total, err := userService.ListUsers(page, pageSize)
	if err != nil {
		response.ServerError(c, "获取用户列表失败")
		return
	}

	response.Success(c, gin.H{
		"users": users,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

// UpdateUserRole 更新用户角色
func UpdateUserRole(c *gin.Context) {
	// 获取用户ID
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID")
		return
	}

	// 绑定请求参数
	var req struct {
		Role string `json:"role" binding:"required,oneof=user moderator admin"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 查询用户信息
	user, err := userService.GetUserByID(uint(userID))
	if err != nil {
		response.NotFound(c, "用户不存在")
		return
	}

	// 更新用户角色
	if err := userService.UpdateUserRole(user.ID, req.Role); err != nil {
		response.ServerError(c, "更新用户角色失败")
		return
	}

	response.Success(c, "更新用户角色成功")
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	// 获取用户ID
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID")
		return
	}

	// 删除用户
	if err := userService.DeleteUser(uint(userID)); err != nil {
		response.ServerError(c, "删除用户失败")
		return
	}

	response.Success(c, "删除用户成功")
}

// ListUserPosts 获取用户的帖子列表
func ListUserPosts(c *gin.Context) {
	// 获取用户ID
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的用户ID")
		return
	}

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 调用帖子服务
	posts, total, err := postService.GetPostsByUserID(uint(userID), page, pageSize)
	if err != nil {
		response.ServerError(c, "获取用户帖子失败")
		return
	}

	response.Success(c, gin.H{
		"posts": posts,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}
