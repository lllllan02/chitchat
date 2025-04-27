package model

import (
	"time"

	"gorm.io/gorm"
)

// Comment 评论模型
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	UserID    uint           `gorm:"index;not null" json:"user_id"`
	PostID    uint           `gorm:"index;not null" json:"post_id"`
	ParentID  *uint          `gorm:"index" json:"parent_id"`
	LikeCount int            `gorm:"default:0" json:"like_count"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	User    User      `gorm:"foreignKey:UserID" json:"user"`
	Post    Post      `gorm:"foreignKey:PostID" json:"post"`
	Parent  *Comment  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
}

// TableName 设置表名
func (Comment) TableName() string {
	return "comments"
}

// CommentRequest 评论请求
type CommentRequest struct {
	Content  string `json:"content" binding:"required,min=1"`
	PostID   uint   `json:"post_id" binding:"required"`
	ParentID *uint  `json:"parent_id"`
}
