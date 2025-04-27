package service

import (
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/repository"
)

// UserService 用户服务接口
type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	ListUsers(page, pageSize int) ([]*model.User, int64, error)
	SearchUsers(keyword string, page, pageSize int) ([]*model.User, int64, error)
	ChangePassword(id uint, newPasswordHash string) error
	UpdateUserRole(id uint, role string) error
}

// userService 用户服务实现
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务
func NewUserService() UserService {
	return &userService{
		userRepo: repository.NewUserRepository(),
	}
}

// CreateUser 创建用户
func (s *userService) CreateUser(user *model.User) error {
	return s.userRepo.Create(user)
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// GetUserByUsername 根据用户名获取用户
func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.GetByUsername(username)
}

// GetUserByEmail 根据邮箱获取用户
func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetByEmail(email)
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(user *model.User) error {
	return s.userRepo.Update(user)
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

// ListUsers 获取用户列表
func (s *userService) ListUsers(page, pageSize int) ([]*model.User, int64, error) {
	return s.userRepo.List(page, pageSize)
}

// SearchUsers 搜索用户
func (s *userService) SearchUsers(keyword string, page, pageSize int) ([]*model.User, int64, error) {
	return s.userRepo.FindByKeyword(keyword, page, pageSize)
}

// ChangePassword 修改密码
func (s *userService) ChangePassword(id uint, newPasswordHash string) error {
	return s.userRepo.UpdatePassword(id, newPasswordHash)
}

// UpdateUserRole 更新用户角色
func (s *userService) UpdateUserRole(id uint, role string) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	user.Role = role
	return s.userRepo.Update(user)
}
