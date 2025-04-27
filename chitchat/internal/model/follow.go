package model

import (
	"time"

	"gorm.io/gorm"
)

// Follow 用户关注模型
type Follow struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	FollowerID uint           `gorm:"index;not null" json:"follower_id"`
	FollowedID uint           `gorm:"index;not null" json:"followed_id"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Follower User `gorm:"foreignKey:FollowerID" json:"follower"`
	Followed User `gorm:"foreignKey:FollowedID" json:"followed"`
}

// TableName 设置表名
func (Follow) TableName() string {
	return "follows"
}
