package router

import (
	"x-project/internal/http/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAccountRouter(r *gin.RouterGroup, db *gorm.DB) {
	accountRouter := r.Group("account")
	accountController := controller.NewAccountController(db)

	// curd
	accountRouter.POST("", accountController.Create)
	// accountRouter.PUT("", accountController.Update)
	// accountRouter.DELETE("/:id", accountController.Delete)
	// accountRouter.GET("/:id", accountController.Get)
	// accountRouter.GET("/list", accountController.List)
}
