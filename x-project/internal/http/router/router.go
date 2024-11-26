package router

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB) error {

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	pprof.Register(r, "pprofyioeanltqw") // 性能
	baseRouter := r.Group("")

	InitAccountRouter(baseRouter, db)
	return nil
}
