package model

import (
	"time"

	"gorm.io/gorm"
)

// Post 帖子模型
type Post struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Title      string         `gorm:"type:varchar(255);not null" json:"title"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	UserID     uint           `gorm:"index;not null" json:"user_id"`
	CategoryID uint           `gorm:"index;not null" json:"category_id"`
	ViewCount  int            `gorm:"default:0" json:"view_count"`
	LikeCount  int            `gorm:"default:0" json:"like_count"`
	IsPinned   bool           `gorm:"default:false" json:"is_pinned"`
	IsFeatured bool           `gorm:"default:false" json:"is_featured"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	User     User     `gorm:"foreignKey:UserID" json:"user"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}

// TableName 设置表名
func (Post) TableName() string {
	return "posts"
}

// PostRequest 帖子请求
type PostRequest struct {
	Title      string `json:"title" binding:"required,min=5,max=255"`
	Content    string `json:"content" binding:"required,min=10"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

// PostUpdateRequest 帖子更新请求
type PostUpdateRequest struct {
	Title      string `json:"title" binding:"required,min=5,max=255"`
	Content    string `json:"content" binding:"required,min=10"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

// PostListQuery 帖子列表查询参数
type PostListQuery struct {
	CategoryID uint   `form:"category_id"`
	UserID     uint   `form:"user_id"`
	Keyword    string `form:"keyword"`
	OrderBy    string `form:"order_by" binding:"omitempty,oneof=latest popular"`
	Page       int    `form:"page" binding:"min=1"`
	PageSize   int    `form:"page_size" binding:"min=1,max=100"`
}
