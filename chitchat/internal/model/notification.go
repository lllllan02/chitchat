package model

import (
	"time"

	"gorm.io/gorm"
)

// NotificationType 通知类型
type NotificationType string

const (
	NotificationTypeReply   NotificationType = "reply"
	NotificationTypeLike    NotificationType = "like"
	NotificationTypeFollow  NotificationType = "follow"
	NotificationTypeMention NotificationType = "mention"
	NotificationTypeSystem  NotificationType = "system"
)

// Notification 通知模型
type Notification struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	UserID    uint             `gorm:"index;not null" json:"user_id"`
	SenderID  *uint            `gorm:"index" json:"sender_id"`
	Type      NotificationType `gorm:"type:varchar(20);not null" json:"type"`
	Content   string           `gorm:"type:text;not null" json:"content"`
	PostID    *uint            `gorm:"index" json:"post_id"`
	CommentID *uint            `gorm:"index" json:"comment_id"`
	Link      string           `gorm:"type:varchar(255)" json:"link"`
	IsRead    bool             `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`

	// 关联
	User    User     `gorm:"foreignKey:UserID" json:"user"`
	Sender  *User    `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	Post    *Post    `gorm:"foreignKey:PostID" json:"post,omitempty"`
	Comment *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`
}

// TableName 设置表名
func (Notification) TableName() string {
	return "notifications"
}

// NotificationListQuery 通知列表查询参数
type NotificationListQuery struct {
	Type     string `form:"type" binding:"omitempty,oneof=reply like follow mention system"`
	IsRead   *bool  `form:"is_read"`
	Page     int    `form:"page" binding:"min=1"`
	PageSize int    `form:"page_size" binding:"min=1,max=100"`
}
