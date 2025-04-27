.PHONY: help backend frontend dev dev-simple build clean

# 默认目标，显示帮助信息
help:
	@echo "可用命令:"
	@echo "  make backend     - 启动后端服务"
	@echo "  make frontend    - 启动前端服务"
	@echo "  make dev         - 同时启动前后端服务 (需要tmux)"
	@echo "  make dev-simple  - 在后台启动后端服务，前台启动前端服务"
	@echo "  make build       - 构建前端项目"
	@echo "  make clean       - 清理构建文件"

# 启动后端服务
backend:
	@echo "启动后端服务..."
	cd chitchat && go run cmd/server/main.go

# 启动前端服务
frontend:
	@echo "启动前端服务..."
	cd chitchat-web && npm run dev

# 同时启动前后端服务
dev:
	@echo "启动后端服务 (后台运行)..."
	cd chitchat && go run cmd/server/main.go > /tmp/chitchat-backend.log 2>&1 &
	@echo "后端服务已在后台启动，日志保存在 /tmp/chitchat-backend.log"
	@echo "启动前端服务 (前台运行)..."
	cd chitchat-web && npm run dev

# 构建前端项目
build:
	@echo "构建前端项目..."
	cd chitchat-web && npm run build

# 清理构建文件
clean:
	@echo "清理构建文件..."
	rm -rf chitchat-web/.next
	rm -rf chitchat-web/out 