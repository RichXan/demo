package router

import (
	"net/http"
	"x-micro-blog/global"
	"x-micro-blog/internal/http/handler"
	"x-micro-blog/internal/http/middleware"
	"x-micro-blog/internal/micro/service"

	"github.com/richxan/xcommon/xlog"
	"github.com/richxan/xcommon/xmiddleware"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Setup 设置路由
func Setup(
	tracer opentracing.Tracer,
	logger *xlog.Logger,
	us service.UserService,
	ps service.PostService,
	as service.AuthService,
) *gin.Engine {
	// 初始化服务依赖
	handler.InitServices(us, ps, as)

	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(xmiddleware.Cors())
	r.Use(xmiddleware.RequestID())
	r.Use(xmiddleware.Logger(logger, global.Config.System.Debug))
	r.Use(xmiddleware.TracingMiddleware(tracer))
	r.Use(middleware.MetricsMiddleware())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API版本
	v1 := r.Group("/api/v1")
	{
		setupUserRoutes(v1)
		setupPostRoutes(v1)
	}

	return r
}
