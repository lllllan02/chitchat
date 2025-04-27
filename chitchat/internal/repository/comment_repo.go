package repository

import (
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/utils"
	"gorm.io/gorm"
)

// CommentRepository 评论仓库接口
type CommentRepository interface {
	Create(comment *model.Comment) error
	GetByID(id uint) (*model.Comment, error)
	Update(comment *model.Comment) error
	Delete(id uint) error
	GetByPostID(postID uint, page, pageSize int) ([]*model.Comment, int64, error)
	GetReplies(parentID uint) ([]*model.Comment, error)
	UpdateLikeCount(id uint, count int) error
}

// commentRepository 评论仓库实现
type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository 创建评论仓库
func NewCommentRepository() CommentRepository {
	return &commentRepository{db: utils.DB}
}

// Create 创建评论
func (r *commentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

// GetByID 根据ID获取评论
func (r *commentRepository) GetByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.Preload("User").First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// Update 更新评论
func (r *commentRepository) Update(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

// Delete 删除评论
func (r *commentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Comment{}, id).Error
}

// GetByPostID 获取帖子的评论列表
func (r *commentRepository) GetByPostID(postID uint, page, pageSize int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64

	// 只获取顶级评论（没有父评论的）
	query := r.db.Where("post_id = ? AND parent_id IS NULL", postID).Preload("User")

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	// 为每个顶级评论加载回复
	for i := range comments {
		if err := r.db.Where("parent_id = ?", comments[i].ID).Preload("User").Order("created_at ASC").Find(&comments[i].Replies).Error; err != nil {
			return nil, 0, err
		}
	}

	return comments, total, nil
}

// GetReplies 获取评论的回复列表
func (r *commentRepository) GetReplies(parentID uint) ([]*model.Comment, error) {
	var replies []*model.Comment
	err := r.db.Where("parent_id = ?", parentID).Preload("User").Order("created_at ASC").Find(&replies).Error
	return replies, err
}

// UpdateLikeCount 更新点赞数
func (r *commentRepository) UpdateLikeCount(id uint, count int) error {
	return r.db.Model(&model.Comment{}).Where("id = ?", id).Update("like_count", count).Error
}
