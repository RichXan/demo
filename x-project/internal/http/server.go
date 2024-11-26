package http

import (
	"fmt"
	"x-project/global"
	"x-project/internal/http/router"
	"x-project/pkg/db"
	"x-project/pkg/log"

	"github.com/gin-gonic/gin"
)

func Start() {

	// gin 初始化
	r := gin.Default()

	// 测试
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// init logger
	logger := log.NewLogger(log.LoggerConfig{
		Level:            global.AppConfig.Log.Level,
		Directory:        global.AppConfig.Log.Directory,
		ProjectName:      global.AppConfig.Log.ProjectName,
		LoggerName:       global.AppConfig.Log.LoggerName,
		MaxSize:          global.AppConfig.Log.MaxSize,
		MaxBackups:       global.AppConfig.Log.MaxBackup,
		SaveLoggerAsFile: global.AppConfig.Log.SaveLoggerAsFile,
	})

	// init db
	database, err := db.NewPostgresGormDbWithDSN(global.AppConfig.Database.Postgres.GormDSN)
	if err != nil {
		logger.Debug().Msgf("init db  error: %v", err)
		return
	}

	// init redis
	// myRedisClient, err := db.NewRedisClientByConfig(&db.RedisConfig{
	// 	Sentinels:  global.AppConfig.Redis.Sentinels,
	// 	Password:   global.AppConfig.Redis.Password,
	// 	MasterName: global.AppConfig.Redis.MasterName,
	// 	Db:         global.AppConfig.Redis.Db,
	// }, logger)
	// if err != nil {
	// 	panic(fmt.Errorf("init redis error and error is %v", err))
	// }
	// redisClient := myRedisClient.Client()

	// 路由表初始化
	router.InitRouter(r, database)
	// listen and serve on 0.0.0.0:8080
	listen := fmt.Sprintf(":%d", global.AppConfig.System.Port)
	logger.Debug().Msgf("listen: %s", listen)
	r.Run(listen)
}
