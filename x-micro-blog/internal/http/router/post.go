package router

import (
	"x-micro-blog/internal/http/handler"
	"x-micro-blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

// setupPostRoutes 设置文章相关路由
func setupPostRoutes(r *gin.RouterGroup) {
	post := r.Group("")
	post.Use(middleware.Auth())
	{
		// 文章相关
		posts := post.Group("/posts")
		{
			posts.POST("", handler.HandleCreatePost)
			posts.GET("", handler.HandleListPosts)
			posts.GET("/:id", handler.HandleGetPost)
			posts.PUT("/:id", handler.HandleUpdatePost)
			posts.DELETE("/:id", handler.HandleDeletePost)

			// 评论相关
			posts.POST("/:id/comments", handler.HandleCreateComment)
			posts.GET("/:id/comments", handler.HandleListComments)
			posts.DELETE("/:id/comments/:comment_id", handler.HandleDeleteComment)

			// 点赞相关
			posts.POST("/:id/like", handler.HandleLikePost)
			posts.DELETE("/:id/like", handler.HandleUnlikePost)
			posts.GET("/:id/likes", handler.HandleListPostLikes)
		}
	}
}
