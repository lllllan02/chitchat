package service

import (
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/repository"
)

// CategoryService 分类服务接口
type CategoryService interface {
	CreateCategory(name, description string) (*model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	UpdateCategory(category *model.Category) error
	DeleteCategory(id uint) error
	ListCategories() ([]*model.Category, error)
}

// categoryService 分类服务实现
type categoryService struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryService 创建分类服务
func NewCategoryService() CategoryService {
	return &categoryService{
		categoryRepo: repository.NewCategoryRepository(),
	}
}

// CreateCategory 创建分类
func (s *categoryService) CreateCategory(name, description string) (*model.Category, error) {
	category := &model.Category{
		Name:        name,
		Description: description,
	}

	err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategoryByID 根据ID获取分类
func (s *categoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.categoryRepo.GetByID(id)
}

// UpdateCategory 更新分类
func (s *categoryService) UpdateCategory(category *model.Category) error {
	return s.categoryRepo.Update(category)
}

// DeleteCategory 删除分类
func (s *categoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.Delete(id)
}

// ListCategories 获取分类列表
func (s *categoryService) ListCategories() ([]*model.Category, error) {
	return s.categoryRepo.List()
}
