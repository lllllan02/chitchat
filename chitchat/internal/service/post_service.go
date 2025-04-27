package service

import (
	"time"

	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/repository"
	"github.com/lllllan02/chitchat/internal/utils"
)

// PostService 帖子服务接口
type PostService interface {
	CreatePost(userID, categoryID uint, title, content string) (*model.Post, error)
	GetPostByID(id uint, includeUser bool) (*model.Post, error)
	UpdatePost(id, userID uint, title, content string, categoryID uint) (*model.Post, error)
	DeletePost(id, userID uint, isAdmin bool) error
	ListPosts(page, pageSize int, categoryID, userID uint, keyword, orderBy string) ([]*model.Post, int64, error)
	GetPostsByUserID(userID uint, page, pageSize int) ([]*model.Post, int64, error)
	ViewPost(id uint) error
	SetPostPinned(id uint, isPinned bool) error
	SetPostFeatured(id uint, isFeatured bool) error
	GetPinnedPosts(categoryID uint, limit int) ([]*model.Post, error)
	GetFeaturedPosts(limit int) ([]*model.Post, error)
}

// postService 帖子服务实现
type postService struct {
	postRepo     repository.PostRepository
	categoryRepo repository.CategoryRepository
}

// NewPostService 创建帖子服务
func NewPostService() PostService {
	return &postService{
		postRepo:     repository.NewPostRepository(),
		categoryRepo: repository.NewCategoryRepository(),
	}
}

// CreatePost 创建帖子
func (s *postService) CreatePost(userID, categoryID uint, title, content string) (*model.Post, error) {
	// 检查分类是否存在
	if categoryID > 0 {
		_, err := s.categoryRepo.GetByID(categoryID)
		if err != nil {
			return nil, utils.ErrCategoryNotFound
		}
	}

	post := &model.Post{
		UserID:     userID,
		CategoryID: categoryID,
		Title:      title,
		Content:    content,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := s.postRepo.Create(post)
	if err != nil {
		return nil, err
	}

	// 更新分类帖子数量
	if categoryID > 0 {
		if err := s.categoryRepo.IncrementPostCount(categoryID); err != nil {
			return nil, err
		}
	}

	return post, nil
}

// GetPostByID 根据ID获取帖子
func (s *postService) GetPostByID(id uint, includeUser bool) (*model.Post, error) {
	return s.postRepo.GetByID(id, includeUser)
}

// UpdatePost 更新帖子
func (s *postService) UpdatePost(id, userID uint, title, content string, categoryID uint) (*model.Post, error) {
	// 获取原始帖子
	post, err := s.postRepo.GetByID(id, false)
	if err != nil {
		return nil, err
	}

	// 检查权限
	if post.UserID != userID {
		return nil, utils.ErrPermissionDenied
	}

	// 检查分类是否需要更改
	if categoryID > 0 && post.CategoryID != categoryID {
		// 检查新分类是否存在
		_, err := s.categoryRepo.GetByID(categoryID)
		if err != nil {
			return nil, utils.ErrCategoryNotFound
		}

		// 更新旧分类帖子数量
		if post.CategoryID > 0 {
			if err := s.categoryRepo.DecrementPostCount(post.CategoryID); err != nil {
				return nil, err
			}
		}

		// 更新新分类帖子数量
		if err := s.categoryRepo.IncrementPostCount(categoryID); err != nil {
			return nil, err
		}

		post.CategoryID = categoryID
	}

	// 更新帖子信息
	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}
	post.UpdatedAt = time.Now()

	// 保存更新
	err = s.postRepo.Update(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// DeletePost 删除帖子
func (s *postService) DeletePost(id, userID uint, isAdmin bool) error {
	// 获取原始帖子
	post, err := s.postRepo.GetByID(id, false)
	if err != nil {
		return err
	}

	// 检查权限
	if post.UserID != userID && !isAdmin {
		return utils.ErrPermissionDenied
	}

	// 删除帖子
	err = s.postRepo.Delete(id)
	if err != nil {
		return err
	}

	// 更新分类帖子数量
	if post.CategoryID > 0 {
		if err := s.categoryRepo.DecrementPostCount(post.CategoryID); err != nil {
			return err
		}
	}

	return nil
}

// ListPosts 获取帖子列表
func (s *postService) ListPosts(page, pageSize int, categoryID, userID uint, keyword, orderBy string) ([]*model.Post, int64, error) {
	return s.postRepo.List(page, pageSize, categoryID, userID, keyword, orderBy)
}

// GetPostsByUserID 获取用户的帖子列表
func (s *postService) GetPostsByUserID(userID uint, page, pageSize int) ([]*model.Post, int64, error) {
	return s.postRepo.List(page, pageSize, 0, userID, "", "")
}

// ViewPost 浏览帖子
func (s *postService) ViewPost(id uint) error {
	return s.postRepo.IncrementViewCount(id)
}

// SetPostPinned 设置帖子置顶状态
func (s *postService) SetPostPinned(id uint, isPinned bool) error {
	return s.postRepo.SetPinned(id, isPinned)
}

// SetPostFeatured 设置帖子精华状态
func (s *postService) SetPostFeatured(id uint, isFeatured bool) error {
	return s.postRepo.SetFeatured(id, isFeatured)
}

// GetPinnedPosts 获取置顶帖子
func (s *postService) GetPinnedPosts(categoryID uint, limit int) ([]*model.Post, error) {
	return s.postRepo.GetPinnedPosts(categoryID, limit)
}

// GetFeaturedPosts 获取精华帖子
func (s *postService) GetFeaturedPosts(limit int) ([]*model.Post, error) {
	return s.postRepo.GetFeaturedPosts(limit)
}
