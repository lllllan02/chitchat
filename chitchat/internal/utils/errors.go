package utils

import "errors"

// 定义错误常量
var (
	ErrPermissionDenied  = errors.New("permission denied")
	ErrCategoryNotFound  = errors.New("category not found")
	ErrUserNotFound      = errors.New("user not found")
	ErrPostNotFound      = errors.New("post not found")
	ErrInvalidParameters = errors.New("invalid parameters")
	ErrInternalServer    = errors.New("internal server error")
	ErrRecordNotFound    = errors.New("record not found")
	ErrDuplicateRecord   = errors.New("duplicate record")
)
