package model

import (
	"time"

	"gorm.io/gorm"
)

// Like 点赞模型
type Like struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"index;not null" json:"user_id"`
	PostID    *uint          `gorm:"index" json:"post_id"`
	CommentID *uint          `gorm:"index" json:"comment_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	User    User     `gorm:"foreignKey:UserID" json:"user"`
	Post    *Post    `gorm:"foreignKey:PostID" json:"post,omitempty"`
	Comment *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`
}

// TableName 设置表名
func (Like) TableName() string {
	return "likes"
}

// LikeRequest 点赞请求
type LikeRequest struct {
	PostID    *uint `json:"post_id"`
	CommentID *uint `json:"comment_id"`
}
