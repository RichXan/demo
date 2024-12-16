package handler

import (
	"x-micro-blog/internal/http/handler/dto"

	"github.com/gin-gonic/gin"
	"github.com/richxan/xcommon/xerror"
	"github.com/richxan/xcommon/xhttp"
)

// HandleListSocialAccounts 获取社交账号列表
func HandleListSocialAccounts(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	accounts, err := userService.ListSocialAccounts(c.Request.Context(), userID)
	if err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, accounts)
}

// HandleBindSocialAccount 绑定社交账号
func HandleBindSocialAccount(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	provider := c.Param("provider")

	var req dto.SocialBind
	if err := c.ShouldBindJSON(&req); err != nil {
		xhttp.Error(c, xerror.ParamError)
		return
	}

	if err := userService.BindSocialAccount(c.Request.Context(), userID, provider, req.Code, req.State); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}

// HandleUnbindSocialAccount 解绑社交账号
func HandleUnbindSocialAccount(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		xhttp.Error(c, xerror.Unauthorized)
		return
	}

	provider := c.Param("provider")
	if err := userService.UnbindSocialAccount(c.Request.Context(), userID, provider); err != nil {
		xhttp.Error(c, err)
		return
	}

	xhttp.Success(c, nil)
}
