package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan02/chitchat/internal/api/handler"
	"github.com/lllllan02/chitchat/internal/api/middleware"
	"github.com/lllllan02/chitchat/internal/utils"
	"github.com/lllllan02/chitchat/pkg/response"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(utils.AppConfig.Server.Mode)

	// 创建默认路由
	r := gin.New()

	// 使用中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	// 静态文件
	r.Static("/uploads", "./uploads")

	// API版本v1
	v1 := r.Group("/api/v1")
	{
		// 无需认证的路由
		v1.GET("/ping", func(c *gin.Context) {
			response.Success(c, "pong")
		})

		// 认证相关路由
		auth := v1.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
		}

		// 分类相关路由
		categories := v1.Group("/categories")
		{
			categories.GET("", handler.ListCategories)
			categories.GET("/:id", handler.GetCategory)
		}

		// 帖子相关路由 - 公开部分
		posts := v1.Group("/posts")
		{
			posts.GET("", handler.ListPosts)
			posts.GET("/:id", handler.GetPost)
			posts.GET("/:id/comments", handler.ListPostComments)
		}

		// 需要认证的路由
		authorized := v1.Group("")
		authorized.Use(middleware.JWT())
		{
			// 用户相关
			users := authorized.Group("/users")
			{
				users.GET("/me", handler.GetCurrentUser)
				users.PUT("/me", handler.UpdateUser)
				users.PUT("/me/password", handler.ChangePassword)
				users.GET("/:id", handler.GetUser)
				users.GET("", handler.ListUsers)
				users.GET("/:id/posts", handler.ListUserPosts)
			}

			// 帖子相关 - 需要认证
			posts := authorized.Group("/posts")
			{
				posts.POST("", handler.CreatePost)
				posts.PUT("/:id", handler.UpdatePost)
				posts.DELETE("/:id", handler.DeletePost)
				posts.POST("/:id/like", handler.LikePost)
				posts.DELETE("/:id/like", handler.UnlikePost)
			}

			// 评论相关
			comments := authorized.Group("/comments")
			{
				comments.POST("", handler.CreateComment)
				comments.PUT("/:id", handler.UpdateComment)
				comments.DELETE("/:id", handler.DeleteComment)
				comments.POST("/:id/like", handler.LikeComment)
				comments.DELETE("/:id/like", handler.UnlikeComment)
			}

			// 通知相关
			notifications := authorized.Group("/notifications")
			{
				notifications.GET("", handler.ListNotifications)
				notifications.PUT("/:id/read", handler.MarkNotificationAsRead)
				notifications.PUT("/read-all", handler.MarkAllNotificationsAsRead)
			}

			// 关注相关
			follows := authorized.Group("/follows")
			{
				follows.POST("/:id", handler.FollowUser)
				follows.DELETE("/:id", handler.UnfollowUser)
				follows.GET("/followers", handler.ListFollowers)
				follows.GET("/following", handler.ListFollowing)
			}
		}

		// 管理员相关路由
		admin := v1.Group("/admin")
		admin.Use(middleware.JWT(), middleware.AdminAuth())
		{
			// 分类管理
			categories := admin.Group("/categories")
			{
				categories.POST("", handler.CreateCategory)
				categories.PUT("/:id", handler.UpdateCategory)
				categories.DELETE("/:id", handler.DeleteCategory)
			}

			// 帖子管理
			posts := admin.Group("/posts")
			{
				posts.PUT("/:id/pin", handler.PinPost)
				posts.PUT("/:id/unpin", handler.UnpinPost)
				posts.PUT("/:id/feature", handler.FeaturePost)
				posts.PUT("/:id/unfeature", handler.UnfeaturePost)
			}

			// 用户管理
			users := admin.Group("/users")
			{
				users.PUT("/:id/role", handler.UpdateUserRole)
				users.DELETE("/:id", handler.DeleteUser)
			}
		}
	}

	// 未找到路由
	r.NoRoute(func(c *gin.Context) {
		response.NotFound(c, "接口不存在")
	})

	return r
}
