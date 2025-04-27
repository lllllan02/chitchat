package repository

import (
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/utils"
	"gorm.io/gorm"
)

// CategoryRepository 分类仓库接口
type CategoryRepository interface {
	Create(category *model.Category) error
	GetByID(id uint) (*model.Category, error)
	GetByName(name string) (*model.Category, error)
	Update(category *model.Category) error
	Delete(id uint) error
	List() ([]*model.Category, error)
	UpdatePostCount(id uint, count int) error
	IncrementPostCount(id uint) error
	DecrementPostCount(id uint) error
}

// categoryRepository 分类仓库实现
type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository 创建分类仓库
func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{db: utils.DB}
}

// Create 创建分类
func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

// GetByID 根据ID获取分类
func (r *categoryRepository) GetByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetByName 根据名称获取分类
func (r *categoryRepository) GetByName(name string) (*model.Category, error) {
	var category model.Category
	err := r.db.Where("name = ?", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// Update 更新分类
func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

// Delete 删除分类
func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}

// List 获取所有分类
func (r *categoryRepository) List() ([]*model.Category, error) {
	var categories []*model.Category
	err := r.db.Order("id ASC").Find(&categories).Error
	return categories, err
}

// UpdatePostCount 更新分类帖子数量
func (r *categoryRepository) UpdatePostCount(id uint, count int) error {
	return r.db.Model(&model.Category{}).Where("id = ?", id).Update("post_count", count).Error
}

// IncrementPostCount 增加分类帖子数量
func (r *categoryRepository) IncrementPostCount(id uint) error {
	return r.db.Model(&model.Category{}).Where("id = ?", id).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// DecrementPostCount 减少分类帖子数量
func (r *categoryRepository) DecrementPostCount(id uint) error {
	return r.db.Model(&model.Category{}).Where("id = ?", id).UpdateColumn("post_count", gorm.Expr("CASE WHEN post_count > 0 THEN post_count - 1 ELSE 0 END")).Error
}
