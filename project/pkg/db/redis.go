package db

import (
	"context"
	"fmt"
	"strings"
	"time"
	"xproject/pkg/log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb    redis.UniversalClient
	logger log.Logger
}

func (r *RedisClient) Shadow(logger *log.Logger) *RedisClient {
	return &RedisClient{r.rdb, *logger}
}

func (r *RedisClient) Set(k, v string, expiration, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	st := time.Now()
	err := r.rdb.Set(ctx, k, v, expiration).Err()
	r.logger.Info().Str("key", k).Str("value", v).Any("error", err).Int("cost(ms)", int(time.Now().Sub(st).Milliseconds())).Msg("set redis finish")
	return err
}

func (r *RedisClient) Get(k string, timeout time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	st := time.Now()
	v, e := r.rdb.Get(ctx, k).Bytes()
	vs := ""
	if v != nil {
		vs = string(v)
	}
	r.logger.Info().Str("key", k).Any("value", vs).Any("error", e).Int("cost(ms)", int(time.Now().Sub(st).Milliseconds())).Msg("get redis finish")
	return v, e
}

func (r *RedisClient) Exists(k string, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	st := time.Now()
	v, e := r.rdb.Exists(ctx, k).Result()
	r.logger.Info().Str("key", k).Int64("Exists", v).Any("error", e).Int("cost(ms)", int(time.Now().Sub(st).Milliseconds())).Msg("key exists finish")
	return v > 0
}

func (r *RedisClient) Delete(k string, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	st := time.Now()
	v, e := r.rdb.Del(ctx, k).Result()
	r.logger.Info().Str("key", k).Int64("Deleted", v).Any("error", e).Int("cost(ms)", int(time.Now().Sub(st).Milliseconds())).Msg("key exists finish")
	return v > 0, e
}

func NewRedisClient(masterName, sentinels, password string, logger *log.Logger) (*RedisClient, error) {
	addresses := strings.Split(sentinels, ",")
	gRedisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        addresses,
		DB:           0,
		Password:     password,
		MaxRetries:   3,
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		PoolSize:     20,
		MinIdleConns: 10,
		MasterName:   masterName,
	})
	if err := gRedisClient.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("init redis error and error is %v", err)
	}
	c := &RedisClient{
		rdb:    gRedisClient,
		logger: *logger,
	}
	return c, nil
}
