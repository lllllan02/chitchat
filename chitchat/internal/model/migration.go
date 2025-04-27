package model

import (
	"github.com/lllllan02/chitchat/internal/utils"
	"github.com/lllllan02/chitchat/pkg/logger"
)

// AutoMigrate 自动迁移数据库结构
func AutoMigrate() error {
	logger.Info("开始数据库迁移...")

	// 迁移表结构
	if err := utils.DB.AutoMigrate(
		&User{},
		&Category{},
		&Post{},
		&Comment{},
		&Like{},
		&Notification{},
		&Follow{},
	); err != nil {
		logger.Error("数据库迁移失败: %v", err)
		return err
	}

	logger.Info("数据库迁移成功")
	return nil
}

// SeedData 填充初始数据
func SeedData() error {
	// 检查是否已有管理员账号
	var adminCount int64
	utils.DB.Model(&User{}).Where("role = ?", "admin").Count(&adminCount)

	// 如果没有管理员账号，创建默认管理员
	if adminCount == 0 {
		logger.Info("创建默认管理员账号...")

		passwordHash, err := utils.HashPassword("admin123")
		if err != nil {
			logger.Error("密码哈希失败: %v", err)
			return err
		}

		admin := User{
			Username:     "admin",
			Email:        "admin@example.com",
			PasswordHash: passwordHash,
			Role:         "admin",
			Bio:          "系统管理员",
		}

		if err := utils.DB.Create(&admin).Error; err != nil {
			logger.Error("创建管理员账号失败: %v", err)
			return err
		}

		logger.Info("管理员账号创建成功")
	}

	// 检查是否已有默认分类
	var categoryCount int64
	utils.DB.Model(&Category{}).Count(&categoryCount)

	// 如果没有分类，创建默认分类
	if categoryCount == 0 {
		logger.Info("创建默认分类...")

		categories := []Category{
			{Name: "综合讨论", Description: "各种话题的综合讨论区"},
			{Name: "技术交流", Description: "编程、技术相关的交流讨论"},
			{Name: "生活分享", Description: "日常生活经验分享"},
			{Name: "兴趣爱好", Description: "分享你的兴趣爱好"},
			{Name: "意见反馈", Description: "网站意见与建议"},
		}

		if err := utils.DB.Create(&categories).Error; err != nil {
			logger.Error("创建默认分类失败: %v", err)
			return err
		}

		logger.Info("默认分类创建成功")
	}

	return nil
}
