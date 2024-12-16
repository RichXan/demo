package router

import (
	"x-micro-blog/internal/http/handler"
	"x-micro-blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

// setupUserRoutes 设置用户相关路由
func setupUserRoutes(r *gin.RouterGroup) {
	public := r.Group("")
	{
		public.POST("/register", handler.HandleRegister)
		public.POST("/login", handler.HandleLogin)
		public.POST("/refresh", handler.HandleRefreshToken)

		// OAuth2相关路由
		oauth := public.Group("/oauth")
		{
			oauth.GET("/:provider", handler.HandleOAuthLogin)
			oauth.GET("/:provider/callback", handler.HandleOAuthCallback)
		}
	}
	
	auth := r.Group("")
	auth.Use(middleware.Auth())
	{
		users := auth.Group("/user")
		{
			// 用户相关
			users.GET("/profile", handler.HandleGetProfile)
			users.PUT("/profile", handler.HandleUpdateProfile)
			users.GET("/social-accounts", handler.HandleListSocialAccounts)
			users.POST("/social-accounts/:provider/bind", handler.HandleBindSocialAccount)
			users.DELETE("/social-accounts/:provider", handler.HandleUnbindSocialAccount)
			users.PUT("/password", handler.HandleChangePassword)
			users.GET("", handler.HandleListUsers)
			users.GET("/:id", handler.HandleGetUser)
			users.POST("/:id/follow", handler.HandleFollowUser)
			users.DELETE("/:id/follow", handler.HandleUnfollowUser)
			users.GET("/:id/followers", handler.HandleListFollowers)
			users.GET("/:id/following", handler.HandleListFollowing)
		}
	}
}
