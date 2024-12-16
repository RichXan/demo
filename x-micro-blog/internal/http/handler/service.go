package handler

import (
	"x-micro-blog/internal/micro/service"
)

var (
	userService service.UserService
	postService service.PostService
	authService service.AuthService
)

// InitServices 初始化服务依赖
func InitServices(us service.UserService, ps service.PostService, as service.AuthService) {
	userService = us
	postService = ps
	authService = as
}
