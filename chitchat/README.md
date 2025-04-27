# ChitChat 论坛系统

ChitChat 是一个使用 Go 语言开发的现代化论坛系统，提供了丰富的社交功能，包括发帖、评论、点赞、关注等。

## 技术栈

### 后端

- 语言: Go
- Web框架: Gin
- ORM: GORM
- 数据库: MySQL
- 认证: JWT

### 前端

- 框架: Next.js 14 (React)
- UI组件: Shadcn UI
- 样式: Tailwind CSS
- 状态管理: React Query + Context API
- 类型: TypeScript

## 功能特点

- 用户管理: 注册、登录、个人资料
- 内容管理: 发帖、评论、分类
- 互动功能: 点赞、关注、通知
- 管理功能: 内容审核、权限管理

## 目录结构

```
chitchat/                # 后端
├── cmd/                 # 应用入口
│   └── server/          # 服务器入口
├── configs/             # 配置文件
├── internal/            # 内部包
│   ├── api/             # API处理器
│   │   ├── handler/     # 请求处理
│   │   ├── middleware/  # 中间件
│   │   └── router/      # 路由
│   ├── model/           # 数据模型
│   ├── repository/      # 数据访问层
│   ├── service/         # 业务逻辑层
│   └── utils/           # 工具函数
├── pkg/                 # 可导出的包
├── uploads/             # 文件上传目录
└── tests/               # 测试

chitchat-web/           # 前端
├── public/             # 静态资源
└── src/                # 源代码
    ├── app/            # 页面
    ├── components/     # 组件
    ├── lib/            # 工具
    ├── hooks/          # 钩子
    ├── context/        # 上下文
    ├── types/          # 类型定义
    └── styles/         # 样式
```

## 快速开始

### 后端

1. 配置环境

```bash
# 克隆项目
git clone https://github.com/yourusername/chitchat.git
cd chitchat

# 安装依赖
go mod tidy
```

2. 配置数据库

```bash
# 创建数据库
mysql -u root -p
CREATE DATABASE chitchat CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

3. 配置文件

```bash
# 创建配置文件
cp configs/config.example.yaml configs/config.yaml

# 编辑配置文件，填写数据库信息等
```

4. 启动服务

```bash
# 启动服务
go run cmd/server/main.go
```

### 前端

1. 安装依赖

```bash
cd chitchat-web
npm install
```

2. 配置环境变量

```bash
cp .env.example .env.local
# 编辑 .env.local 文件，设置API地址
```

3. 启动开发服务器

```bash
npm run dev
```

## API 文档

服务启动后，访问以下端点查看功能:

- `GET /api/v1/ping`: 测试API是否可用
- `POST /api/v1/auth/register`: 用户注册
- `POST /api/v1/auth/login`: 用户登录
- `GET /api/v1/categories`: 获取所有分类
- `GET /api/v1/posts`: 获取帖子列表
- `GET /api/v1/posts/:id`: 获取帖子详情
- ...更多API请参考代码或文档

## 许可证

MIT
