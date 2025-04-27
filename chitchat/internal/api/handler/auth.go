package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/service"
	"github.com/lllllan02/chitchat/internal/utils"
	"github.com/lllllan02/chitchat/pkg/response"
	"gorm.io/gorm"
)

// 初始化服务
var userService = service.NewUserService()

// Register 用户注册
func Register(c *gin.Context) {
	// 绑定请求参数
	var req model.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 检查用户名是否已存在
	_, err := userService.GetUserByUsername(req.Username)
	if err == nil {
		response.BadRequest(c, "用户名已存在")
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.ServerError(c, "服务器错误")
		return
	}

	// 检查邮箱是否已存在
	_, err = userService.GetUserByEmail(req.Email)
	if err == nil {
		response.BadRequest(c, "邮箱已存在")
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.ServerError(c, "服务器错误")
		return
	}

	// 密码加密
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		response.ServerError(c, "密码加密失败")
		return
	}

	// 创建用户
	user := &model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash,
		Role:         "user", // 默认为普通用户
	}

	if err := userService.CreateUser(user); err != nil {
		response.ServerError(c, "创建用户失败: "+err.Error())
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		response.ServerError(c, "生成令牌失败")
		return
	}

	// 返回用户信息和token
	response.Success(c, model.UserWithToken{
		User:  user,
		Token: token,
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	// 绑定请求参数
	var req model.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 根据用户名获取用户
	user, err := userService.GetUserByUsername(req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.BadRequest(c, "用户名或密码错误")
		return
	} else if err != nil {
		response.ServerError(c, "服务器错误")
		return
	}

	// 验证密码
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		response.BadRequest(c, "用户名或密码错误")
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		response.ServerError(c, "生成令牌失败")
		return
	}

	// 返回用户信息和token
	response.Success(c, model.UserWithToken{
		User:  user,
		Token: token,
	})
}
