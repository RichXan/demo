package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main_hystrix() {
	// 1. 配置 hystrix command
	hystrix.ConfigureCommand("my_service", hystrix.CommandConfig{
		Timeout:                1000, // 超时时间设置为 1000 毫秒
		MaxConcurrentRequests:  100,  // 最大并发请求数
		ErrorPercentThreshold:  25,   // 错误百分比阈值
		RequestVolumeThreshold: 10,   // 请求量阈值
		SleepWindow:            5000, // 熔断器打开后，多久后尝试服务是否恢复
	})

	// 2. 使用 hystrix 包装服务调用
	err := hystrix.Do("my_service", func() error {
		// 这里是正常的业务逻辑
		return callRemoteService()
	}, func(err error) error {
		// 这里是降级逻辑
		return fallback()
	})

	if err != nil {
		fmt.Printf("服务调用失败: %v\n", err)
	}
}

// 模拟远程服务调用
func callRemoteService() error {
	// 模拟一个耗时的 HTTP 请求
	client := &http.Client{}
	_, err := client.Get("http://example.com")
	if err != nil {
		return err
	}
	return nil
}

// 降级处理函数
func fallback() error {
	// 返回降级响应
	return errors.New("服务暂时不可用，请稍后重试")
}
