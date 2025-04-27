package repository

import (
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/utils"
	"gorm.io/gorm"
)

// PostRepository 帖子仓库接口
type PostRepository interface {
	Create(post *model.Post) error
	GetByID(id uint, includeUser bool) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
	List(page, pageSize int, categoryID, userID uint, keyword, orderBy string) ([]*model.Post, int64, error)
	IncrementViewCount(id uint) error
	UpdateLikeCount(id uint, count int) error
	SetPinned(id uint, isPinned bool) error
	SetFeatured(id uint, isFeatured bool) error
	GetPinnedPosts(categoryID uint, limit int) ([]*model.Post, error)
	GetFeaturedPosts(limit int) ([]*model.Post, error)
}

// postRepository 帖子仓库实现
type postRepository struct {
	db *gorm.DB
}

// NewPostRepository 创建帖子仓库
func NewPostRepository() PostRepository {
	return &postRepository{db: utils.DB}
}

// Create 创建帖子
func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

// GetByID 根据ID获取帖子
func (r *postRepository) GetByID(id uint, includeUser bool) (*model.Post, error) {
	var post model.Post
	query := r.db

	if includeUser {
		query = query.Preload("User").Preload("Category")
	}

	err := query.First(&post, id).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// Update 更新帖子
func (r *postRepository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

// Delete 删除帖子
func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&model.Post{}, id).Error
}

// List 获取帖子列表
func (r *postRepository) List(page, pageSize int, categoryID, userID uint, keyword, orderBy string) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	query := r.db.Model(&model.Post{}).Preload("User").Preload("Category")

	// 筛选条件
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}

	if keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 排序
	switch orderBy {
	case "popular":
		query = query.Order("like_count DESC, created_at DESC")
	default:
		query = query.Order("created_at DESC")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// IncrementViewCount 增加浏览次数
func (r *postRepository) IncrementViewCount(id uint) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}

// UpdateLikeCount 更新点赞数
func (r *postRepository) UpdateLikeCount(id uint, count int) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).Update("like_count", count).Error
}

// SetPinned 设置置顶状态
func (r *postRepository) SetPinned(id uint, isPinned bool) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).Update("is_pinned", isPinned).Error
}

// SetFeatured 设置精华状态
func (r *postRepository) SetFeatured(id uint, isFeatured bool) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).Update("is_featured", isFeatured).Error
}

// GetPinnedPosts 获取置顶帖子
func (r *postRepository) GetPinnedPosts(categoryID uint, limit int) ([]*model.Post, error) {
	var posts []*model.Post
	query := r.db.Where("is_pinned = ?", true).Preload("User").Preload("Category").Order("updated_at DESC")

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&posts).Error
	return posts, err
}

// GetFeaturedPosts 获取精华帖子
func (r *postRepository) GetFeaturedPosts(limit int) ([]*model.Post, error) {
	var posts []*model.Post
	query := r.db.Where("is_featured = ?", true).Preload("User").Preload("Category").Order("updated_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&posts).Error
	return posts, err
}
