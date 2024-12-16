package main

import (
	"log"
	"os"

	"x-micro-blog/cmd"
	"x-micro-blog/global"
	"x-micro-blog/internal/micro/service"

	"github.com/richxan/xcommon/xlog"

	"go-micro.dev/v4"
)

func main() {
	// 创建服务
	srv := micro.NewService(
		micro.Name("go.micro.srv.post"),
		micro.Version(cmd.VERSION),
	)

	// 初始化服务
	srv.Init()

	// 加载配置
	configFile := os.Getenv("MICRO_CONFIG_FILE")
	if err := global.LoadConfig(configFile); err != nil {
		log.Fatalf("load config error: %v", err)
	}

	// 根据环境设置覆盖配置
	global.Config.System.Env = os.Getenv("MICRO_ENV")
	if os.Getenv("MICRO_DEBUG") == "true" {
		global.Config.Log.Level = "debug"
	}

	// 初始化日志
	logger := xlog.NewLogger(global.Config.Log)

	// 打印启动信息
	logger.Info().
		Str("version", cmd.VERSION).
		Str("env", global.Config.System.Env).
		Bool("debug", os.Getenv("MICRO_DEBUG") == "true").
		Msg("Starting post service...")

	// 启动文章服务
	if err := service.StartPostService(srv, logger); err != nil {
		log.Fatal(err)
	}

	// 运行服务
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
