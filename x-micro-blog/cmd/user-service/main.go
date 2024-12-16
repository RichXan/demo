package main

import (
	"fmt"
	"log"
	"os"

	"x-micro-blog/cmd"
	"x-micro-blog/global"
	"x-micro-blog/internal/micro/service"

	"github.com/richxan/xcommon/xdatabase"
	"github.com/richxan/xcommon/xlog"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "user-service",
		Usage:   "x-micro-blog user service",
		Authors: cmd.Authors,
		Version: cmd.VERSION,
		Flags:   cmd.Flags,
		Before: func(c *cli.Context) error {
			// 打印版本信息
			if c.Bool("debug") {
				fmt.Printf("Version: %s\n", cmd.VERSION)
				fmt.Printf("BuildTime: %s\n", cmd.BuildTime)
				fmt.Printf("GitCommit: %s\n", cmd.GitCommit)
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			// 加载配置
			if err := global.LoadConfig(c.String("config")); err != nil {
				return fmt.Errorf("load config error: %v", err)
			}

			// 根据环境设置覆盖配置
			global.Config.System.Env = c.String("env")
			if c.Bool("debug") {
				global.Config.Log.Level = "debug"
			}

			// 初始化数据库
			db, err := xdatabase.NewMySQLGormDb(&global.Config.MySQL)
			if err != nil {
				return fmt.Errorf("init database error: %v", err)
			}

			// 初始化日志
			logger := xlog.NewLogger(global.Config.Log)

			// 打印启动信息
			logger.Info().
				Str("version", cmd.VERSION).
				Str("env", global.Config.System.Env).
				Bool("debug", c.Bool("debug")).
				Msg("Starting user service...")

			// 启动用户服务
			service.StartUserService(logger, db)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
