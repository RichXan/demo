package handler

import (
	"strconv"

	"x-micro-blog/internal/http/handler/dto"

	"github.com/RichXan/xcommon/xerror"
	"github.com/RichXan/xcommon/xhttp"
	"github.com/gin-gonic/gin"
)

// 处理函数
func HandleRefreshToken(c *gin.Context)  {}
func HandleOAuthLogin(c *gin.Context)    {}
func HandleOAuthCallback(c *gin.Context) {}

// HandleRegister 用户注册
func HandleRegister(c *gin.Context) {
	var req dto.UserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	user, err := userService.Register(c.Request.Context(), req.Username, req.Password, req.Email)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, user)
}

// HandleLogin 用户登录
func HandleLogin(c *gin.Context) {
	var req dto.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	token, err := userService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, gin.H{"token": token})
}

// HandleGetProfile 获取用户信息
func HandleGetProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	user, err := userService.GetUser(c.Request.Context(), userID)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, user)
}

// HandleUpdateProfile 更新用户信息
func HandleUpdateProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	var req dto.UserUpdateProfile
	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	user, err := userService.UpdateUser(c.Request.Context(), userID, req.Nickname, req.Avatar, req.Bio)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, user)
}

// HandleChangePassword 修改密码
func HandleChangePassword(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	var req dto.UserChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	if err := userService.ChangePassword(c.Request.Context(), userID, req.OldPassword, req.NewPassword); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleGetUser 获取指定用户信息
func HandleGetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	user, err := userService.GetUser(c.Request.Context(), userID)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, user)
}

// HandleListUsers 获取用户列表
func HandleListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := userService.ListUsers(c.Request.Context(), page, pageSize)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, xhttp.NewPage(page, pageSize, total, users))
}

// HandleFollowUser 关注用户
func HandleFollowUser(c *gin.Context) {
	targetID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	if err := userService.FollowUser(c.Request.Context(), userID, targetID); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleUnfollowUser 取消关注
func HandleUnfollowUser(c *gin.Context) {
	targetID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	if err := userService.UnfollowUser(c.Request.Context(), userID, targetID); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleListFollowers 获取粉丝列表
func HandleListFollowers(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := userService.ListFollowers(c.Request.Context(), userID, page, pageSize)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, xhttp.NewPage(page, pageSize, total, users))
}

// HandleListFollowing 获取关注列表
func HandleListFollowing(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := userService.ListFollowing(c.Request.Context(), userID, page, pageSize)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, xhttp.NewPage(page, pageSize, total, users))
}
