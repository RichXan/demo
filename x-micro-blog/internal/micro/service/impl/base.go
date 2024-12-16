package impl

import (
	"github.com/richxan/xcommon/xlog"

	"gorm.io/gorm"
)

// BaseService 基础服务
type BaseService struct {
	db     *gorm.DB
	logger *xlog.Logger
}

// NewBaseService 创建基础服务
func NewBaseService(db *gorm.DB, logger *xlog.Logger) BaseService {
	return BaseService{
		db:     db,
		logger: logger,
	}
}
