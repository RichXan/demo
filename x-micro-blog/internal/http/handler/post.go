package handler

import (
	"strconv"

	"x-micro-blog/internal/http/handler/dto"

	"github.com/gin-gonic/gin"
	"github.com/richxan/xcommon/xerror"
	"github.com/richxan/xcommon/xhttp"
)

// HandleCreatePost 创建文章
func HandleCreatePost(c *gin.Context) {
	var req dto.PostCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	// 获取当前用户ID
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	// 调用服务创建文章
	post, err := postService.CreatePost(c.Request.Context(), userID, req.Title, req.Content)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, post)
}

// HandleGetPost 获取文章详情
func HandleGetPost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	post, err := postService.GetPost(c.Request.Context(), postID)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, post)
}

// HandleListPosts 获取文章列表
func HandleListPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	posts, total, err := postService.ListPosts(c.Request.Context(), page, pageSize)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, xhttp.NewPage(page, pageSize, total, posts))
}

// HandleUpdatePost 更新文章
func HandleUpdatePost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	var req dto.PostUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	// 获取当前用户ID
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	// 检查权限
	post, err := postService.GetPost(c.Request.Context(), postID)
	if err != nil {
		xhttp.Error(c, err)
		return
	}
	if post.UserID != uint64(userID) {
		xhttp.Error(c, xerror.PostForbidden)
		return
	}

	// 更新文章
	post, err = postService.UpdatePost(c.Request.Context(), postID, req.Title, req.Content)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, post)
}

// HandleDeletePost 删除文章
func HandleDeletePost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	// 获取当前用户ID
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	// 检查权限
	post, err := postService.GetPost(c.Request.Context(), postID)
	if err != nil {
		xhttp.Error(c, err)
		return
	}
	if post.UserID != uint64(userID) {
		xhttp.Error(c, xerror.PostForbidden)
		return
	}

	// 删除文章
	if err := postService.DeletePost(c.Request.Context(), postID); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleCreateComment 创建评论
func HandleCreateComment(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	var req struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	// 获取当前用户ID
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	// 创建评论
	comment, err := postService.CreateComment(c.Request.Context(), postID, userID, req.Content)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, comment)
}

// HandleListComments 获取评论列表
func HandleListComments(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	comments, total, err := postService.ListComments(c.Request.Context(), postID, page, pageSize)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, xhttp.NewPage(page, pageSize, total, comments))
}

// HandleDeleteComment 删除评论
func HandleDeleteComment(c *gin.Context) {
	commentID, err := strconv.ParseInt(c.Param("comment_id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	// 获取当前用户ID
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	// 检查权限
	comment, err := postService.GetComment(c.Request.Context(), commentID)
	if err != nil {
		xhttp.Error(c, err)
		return
	}
	if comment.UserID != uint64(userID) {
		xhttp.Error(c, xerror.CommentForbidden)
		return
	}

	// 删除评论
	if err := postService.DeleteComment(c.Request.Context(), commentID); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleLikePost 点赞文章
func HandleLikePost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	// 获取当前用户ID
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	// 点赞文章
	if err := postService.LikePost(c.Request.Context(), postID, userID); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleUnlikePost 取消点赞
func HandleUnlikePost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	// 获取当前用户ID
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	// 取消点赞
	if err := postService.UnlikePost(c.Request.Context(), postID, userID); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleListPostLikes 获取文章点赞列表
func HandleListPostLikes(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := postService.ListPostLikes(c.Request.Context(), postID, page, pageSize)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, xhttp.NewPage(page, pageSize, total, users))
}
