package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	Path        string `yaml:"path"` // host + port
	Database    string `yaml:"database"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Config      string `yaml:"config"`
	MaxIdleConn int    `yaml:"max_idle_conn"`
	MaxOpenConn int    `yaml:"max_open_conn"`
	IsConsole   bool   `yaml:"is_console"`
}

func NewGormDb(mysqlConfig *MysqlConfig) (e *gorm.DB, err error) {
	var (
		userName    = mysqlConfig.User
		password    = mysqlConfig.Password
		path        = mysqlConfig.Path
		database    = mysqlConfig.Database
		config      = mysqlConfig.Config
		maxIdleConn = mysqlConfig.MaxIdleConn
		maxOpenConn = mysqlConfig.MaxOpenConn
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", userName, password, path, database, config)

	cfg := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	defer func() {
		if err != nil {
			return
		}
		db, _ := e.DB()
		db.SetMaxIdleConns(maxIdleConn)
		db.SetMaxOpenConns(maxOpenConn)
	}()

	loggerLevel := logger.Silent
	// 20220307 添加对应的 console 配置
	if mysqlConfig.IsConsole {
		loggerLevel = logger.Info
	}
	return gorm.Open(mysql.New(cfg), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      loggerLevel, // Log level
				Colorful:      false,       // 禁用彩色打印
			},
		),
	})
}
