package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 标准响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// Fail 返回失败响应
func Fail(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// BadRequest 返回400错误
func BadRequest(c *gin.Context, message string) {
	if message == "" {
		message = "请求参数错误"
	}
	Fail(c, http.StatusBadRequest, message)
}

// Unauthorized 返回401错误
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "未授权访问"
	}
	Fail(c, http.StatusUnauthorized, message)
}

// Forbidden 返回403错误
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "禁止访问"
	}
	Fail(c, http.StatusForbidden, message)
}

// NotFound 返回404错误
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "资源不存在"
	}
	Fail(c, http.StatusNotFound, message)
}

// ServerError 返回500错误
func ServerError(c *gin.Context, message string) {
	if message == "" {
		message = "服务器内部错误"
	}
	Fail(c, http.StatusInternalServerError, message)
}

// NotImplemented 返回501错误
func NotImplemented(c *gin.Context, message string) {
	if message == "" {
		message = "功能未实现"
	}
	Fail(c, http.StatusNotImplemented, message)
}
