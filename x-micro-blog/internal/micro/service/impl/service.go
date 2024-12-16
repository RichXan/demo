package impl

import (
	"x-micro-blog/internal/micro/service"

	"github.com/richxan/xcommon/xlog"

	"gorm.io/gorm"
)

// NewPostService 创建文章服务
func NewPostService(db *gorm.DB, logger *xlog.Logger) service.PostService {
	return &postService{
		BaseService: NewBaseService(db, logger),
	}
}

// postService 文章服务实现
type postService struct {
	BaseService
}
