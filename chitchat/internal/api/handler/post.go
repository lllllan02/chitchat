package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/internal/utils"
	"github.com/lllllan02/chitchat/pkg/response"
)

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 绑定请求参数
	var req struct {
		Title      string `json:"title" binding:"required"`
		Content    string `json:"content" binding:"required"`
		CategoryID uint   `json:"category_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 创建帖子
	post, err := postService.CreatePost(userID.(uint), req.CategoryID, req.Title, req.Content)
	if err != nil {
		response.ServerError(c, "创建帖子失败: "+err.Error())
		return
	}

	response.Success(c, post)
}

// GetPost 获取帖子详情
func GetPost(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 获取帖子
	post, err := postService.GetPostByID(uint(postID), true)
	if err != nil {
		response.NotFound(c, "帖子不存在")
		return
	}

	// 增加浏览次数
	go func() {
		_ = postService.ViewPost(uint(postID))
	}()

	response.Success(c, post)
}

// UpdatePost 更新帖子
func UpdatePost(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}

	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 绑定请求参数
	var req struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		CategoryID uint   `json:"category_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 更新帖子
	post, err := postService.UpdatePost(uint(postID), userID.(uint), req.Title, req.Content, req.CategoryID)
	if err != nil {
		if err == utils.ErrPermissionDenied {
			response.Forbidden(c, "没有权限更新该帖子")
			return
		}
		response.ServerError(c, "更新帖子失败: "+err.Error())
		return
	}

	response.Success(c, post)
}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	// 从上下文中获取用户ID和角色
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "用户未认证")
		return
	}
	userRole, _ := c.Get("userRole")
	isAdmin := userRole.(string) == "admin"

	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 删除帖子
	err = postService.DeletePost(uint(postID), userID.(uint), isAdmin)
	if err != nil {
		if err == utils.ErrPermissionDenied {
			response.Forbidden(c, "没有权限删除该帖子")
			return
		}
		response.ServerError(c, "删除帖子失败: "+err.Error())
		return
	}

	response.Success(c, "删除帖子成功")
}

// ListPosts 获取帖子列表
func ListPosts(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取筛选参数
	categoryIDStr := c.DefaultQuery("category_id", "0")
	categoryID, _ := strconv.ParseUint(categoryIDStr, 10, 32)
	keyword := c.DefaultQuery("keyword", "")
	orderBy := c.DefaultQuery("order_by", "recent")

	// 查询帖子列表
	posts, total, err := postService.ListPosts(page, pageSize, uint(categoryID), 0, keyword, orderBy)
	if err != nil {
		response.ServerError(c, "获取帖子列表失败")
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

// LikePost 点赞帖子
func LikePost(c *gin.Context) {
	// TODO: 实现点赞帖子功能
	response.NotImplemented(c, "功能未实现")
}

// UnlikePost 取消点赞帖子
func UnlikePost(c *gin.Context) {
	// TODO: 实现取消点赞帖子功能
	response.NotImplemented(c, "功能未实现")
}

// PinPost 置顶帖子
func PinPost(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 设置置顶
	if err := postService.SetPostPinned(uint(postID), true); err != nil {
		response.ServerError(c, "置顶帖子失败")
		return
	}

	response.Success(c, "置顶帖子成功")
}

// UnpinPost 取消置顶帖子
func UnpinPost(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 取消置顶
	if err := postService.SetPostPinned(uint(postID), false); err != nil {
		response.ServerError(c, "取消置顶失败")
		return
	}

	response.Success(c, "取消置顶成功")
}

// FeaturePost 设置精华帖子
func FeaturePost(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 设置精华
	if err := postService.SetPostFeatured(uint(postID), true); err != nil {
		response.ServerError(c, "设置精华失败")
		return
	}

	response.Success(c, "设置精华成功")
}

// UnfeaturePost 取消精华帖子
func UnfeaturePost(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的帖子ID")
		return
	}

	// 取消精华
	if err := postService.SetPostFeatured(uint(postID), false); err != nil {
		response.ServerError(c, "取消精华失败")
		return
	}

	response.Success(c, "取消精华成功")
}
