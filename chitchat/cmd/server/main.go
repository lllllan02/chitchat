package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lllllan02/chitchat/internal/api/router"
	"github.com/lllllan02/chitchat/internal/model"
	"github.com/lllllan02/chitchat/internal/utils"
	"github.com/lllllan02/chitchat/pkg/logger"
)

func main() {
	// 初始化日志
	if err := logger.Init("INFO", ""); err != nil {
		fmt.Printf("初始化日志失败: %v\n", err)
		return
	}
	defer logger.Close()

	// 加载配置
	if err := utils.LoadConfig(""); err != nil {
		logger.Fatal("加载配置失败: %v", err)
	}

	// 初始化数据库
	if err := utils.InitDB(); err != nil {
		logger.Fatal("初始化数据库失败: %v", err)
	}

	// 自动迁移表
	if err := model.AutoMigrate(); err != nil {
		logger.Fatal("数据库迁移失败: %v", err)
	}

	// 加载初始数据
	if err := model.SeedData(); err != nil {
		logger.Fatal("填充初始数据失败: %v", err)
	}

	// 初始化路由
	r := router.InitRouter()

	// 启动服务器
	port := utils.AppConfig.Server.Port
	if port == 0 {
		port = 8080
	}

	// 非阻塞方式启动
	go func() {
		addr := fmt.Sprintf(":%d", port)
		logger.Info("服务器启动成功，监听端口: %d", port)
		if err := r.Run(addr); err != nil {
			logger.Fatal("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("正在关闭服务器...")

	logger.Info("服务器已关闭")
}
