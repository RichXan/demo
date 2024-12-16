package main

import (
	"fmt"
	"log"
	"os"

	"x-micro-blog/cmd"
	"x-micro-blog/global"
	"x-micro-blog/internal/http"
	"x-micro-blog/internal/micro/service/impl"

	"github.com/richxan/xcommon/xauth"
	"github.com/richxan/xcommon/xcache"
	"github.com/richxan/xcommon/xdatabase"
	"github.com/richxan/xcommon/xlog"

	"github.com/redis/go-redis/v9"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

func main() {
	app := &cli.App{
		Name:    "x-micro-blog",
		Usage:   "x-micro-blog service",
		Version: cmd.VERSION,
		Authors: cmd.Authors,
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
	}

	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "start http server",
			Action: startHttpServer,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startHttpServer(c *cli.Context) error {
	// 加载配置
	if err := global.LoadConfig(c.String("config")); err != nil {
		return fmt.Errorf("load config error: %v", err)
	}

	// 根据环境设置覆盖配置
	global.Config.System.Env = c.String("env")
	if c.Bool("debug") {
		global.Config.Log.Level = "debug"
	}

	// 初始化日志
	logger := xlog.NewLogger(global.Config.Log)

	// 初始化数据库
	db, err := xdatabase.NewMySQLGormDb(&global.Config.MySQL)
	if err != nil {
		panic(fmt.Errorf("init database error: %v", err))
	}

	// 初始化Redis
	redisClient, err := xcache.NewRedisClient(global.Config.Redis.MasterName, global.Config.Redis.Addresses, global.Config.Redis.Password, logger)
	if err != nil {
		panic(fmt.Errorf("init redis error: %v", err))
	}
	redisSimpleClient, ok := redisClient.Client().(*redis.Client)
	if !ok {
		panic(fmt.Errorf("redis client is not a simple client"))
	}

	// 初始化oauthConfig
	oauthConfig := &xauth.OAuthConfig{
		Providers: map[string]*xauth.OAuthProvider{
			"github": {
				Config: &oauth2.Config{
					ClientID:     global.Config.Social.OAuth.Providers.Github.ClientID,
					ClientSecret: global.Config.Social.OAuth.Providers.Github.ClientSecret,
					Scopes:       global.Config.Social.OAuth.Providers.Github.Scopes,
				},
			},
		},
	}

	// 初始化服务
	userService := impl.NewUserService(db, logger)
	postService := impl.NewPostService(db, logger)
	authService := impl.NewAuthService(db, logger, redisSimpleClient, oauthConfig)

	// 打印启动信息
	logger.Info().
		Str("version", cmd.VERSION).
		Str("env", global.Config.System.Env).
		Bool("debug", c.Bool("debug")).
		Msg("Starting service...")

	// 启动HTTP服务
	http.Start(logger, userService, postService, authService)
	return nil
}
