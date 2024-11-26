package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
	"x-project/config"
	"x-project/pkg/db"
	"x-project/pkg/log"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ExchangeTradeDB *gorm.DB              //全局数据库访问对象
	RedisClient     redis.UniversalClient //全局redis客户端访问对象
	RedSync         *redsync.Redsync      //全局redis客户端访问对象
	mutex           sync.Mutex
	logger          *log.Logger
)

// IntiData 初始化数据配置及对象
func IntiData(appConfig config.Config, log *log.Logger) {
	dataDb, err := db.NewMysqlGormDb(&db.MysqlConfig{
		Path:        appConfig.Database.Mysql.Path,
		Database:    appConfig.Database.Mysql.Database,
		User:        appConfig.Database.Mysql.User,
		Password:    appConfig.Database.Mysql.Password,
		MaxOpenConn: appConfig.Database.Mysql.MaxOpenConn,
		MaxIdleConn: appConfig.Database.Mysql.MaxIdleConn,
		IsConsole:   appConfig.Database.Mysql.IsConsole,
		Config:      appConfig.Database.Mysql.Config,
	})
	logger = log
	if err != nil {
		panic(fmt.Sprintf("mysql conn error %s", err))
	}
	taDb, _ := dataDb.DB()
	if err = taDb.Ping(); err != nil {
		panic(fmt.Errorf("init mysql error and error is %v", err))
	}
	ExchangeTradeDB = dataDb
	addresses := strings.Split(appConfig.Redis.Sentinels, ",")
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           addresses,
		DB:              appConfig.Redis.Db,
		Password:        appConfig.Redis.Password,
		MaxRetries:      3,
		DialTimeout:     time.Second * 5,
		ReadTimeout:     time.Second * 15,
		WriteTimeout:    time.Second * 15,
		PoolSize:        100,
		MinIdleConns:    20,
		ConnMaxIdleTime: time.Second * 120,
		ConnMaxLifetime: time.Hour * 2,
		MasterName:      appConfig.Redis.MasterName,
	})
	if err = redisClient.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("init redis error and error is %v", err))
	}
	RedisClient = redisClient
	newRedisSync(redisClient)
}

// 创建分布式锁对象
func newRedisSync(redisClient redis.UniversalClient) *redsync.Redsync {
	if RedSync == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if redisClient == nil {
			panic("redis client is nil")
		}
		RedSync = redsync.New(goredis.NewPool(redisClient))
	}
	return RedSync
}

// TxFunc 事务控制封装
func TxFunc(db *gorm.DB, fn func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	if err = tx.Error; err != nil {
		logger.Error().Msgf("开启事务异常 %v", err)
		return err
	}
	defer func() {
		r := recover()
		if r != nil {
			// 在发生panic时回滚事务
			tx.Rollback()
			logger.Error().Msgf("事务发生异常 %v", r)
		} else {
			switch err {
			case nil:
				err = tx.Commit().Error
			default:
				logger.Error().Msgf("事务发生异常 %v", err)
				tx.Rollback()
			}
		}
	}()
	err = fn(tx)
	return err
}

// RedLockFunc RedLock分布式锁控制封装
func RedLockFunc(key string, fn func() error, options ...redsync.Option) (err error) {
	s := time.Now().UnixMilli()
	redSyncMutex := RedSync.NewMutex(key, options...)
	if err = redSyncMutex.Lock(); err != nil {
		logger.Error().Msgf("获取分布式锁 %s异常 %v", key, err)
		return errors.New("busy business, please try again later")
	}
	logger.Info().Msgf("获取分布式锁 %s 耗时%d", key, time.Now().UnixMilli()-s)
	defer func() {
		count, existsErr := RedisClient.Exists(context.Background(), key).Result()
		if existsErr == nil && count > 0 { //存在key才释放，防止出现超时key失效异常
			if ok, err1 := redSyncMutex.Unlock(); !ok || err1 != nil {
				logger.Error().Msgf("释放分布式锁 %s异常 %v", key, err1)
				err = err1
			} else {
				logger.Info().Msgf("释放分布式锁 %s 耗时%d", key, time.Now().UnixMilli()-s)
			}
		}
	}()
	err = fn()
	if err != nil {
		logger.Error().Msgf("业务逻辑处理失败 %s异常 %v", key, err)
	}
	return err
}

func Get() {

}
