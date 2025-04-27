package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/internal/service"
	"github.com/lllllan02/chitchat/pkg/response"
)

// 初始化分类服务
var categoryService = service.NewCategoryService()

// ListCategories 获取分类列表
func ListCategories(c *gin.Context) {
	categories, err := categoryService.ListCategories()
	if err != nil {
		response.ServerError(c, "获取分类列表失败")
		return
	}

	response.Success(c, categories)
}

// GetCategory 获取分类详情
func GetCategory(c *gin.Context) {
	// 获取分类ID
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的分类ID")
		return
	}

	// 获取分类
	category, err := categoryService.GetCategoryByID(uint(categoryID))
	if err != nil {
		response.NotFound(c, "分类不存在")
		return
	}

	response.Success(c, category)
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	// 绑定请求参数
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 创建分类
	category, err := categoryService.CreateCategory(req.Name, req.Description)
	if err != nil {
		response.ServerError(c, "创建分类失败")
		return
	}

	response.Success(c, category)
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	// 获取分类ID
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的分类ID")
		return
	}

	// 绑定请求参数
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 获取原始分类
	category, err := categoryService.GetCategoryByID(uint(categoryID))
	if err != nil {
		response.NotFound(c, "分类不存在")
		return
	}

	// 更新分类
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}

	if err := categoryService.UpdateCategory(category); err != nil {
		response.ServerError(c, "更新分类失败")
		return
	}

	response.Success(c, category)
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	// 获取分类ID
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的分类ID")
		return
	}

	// 删除分类
	if err := categoryService.DeleteCategory(uint(categoryID)); err != nil {
		response.ServerError(c, "删除分类失败")
		return
	}

	response.Success(c, "删除分类成功")
}
