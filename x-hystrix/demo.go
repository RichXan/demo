package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func init() {
	// 启动 hystrix metrics stream handler
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "8831"), hystrixStreamHandler)
}

type Service struct {
	timeout    time.Duration
	maxRetries int
}

func NewService() *Service {
	// 配置 hystrix
	hystrix.ConfigureCommand("api_service", hystrix.CommandConfig{
		Timeout:                2000,
		MaxConcurrentRequests:  100,
		ErrorPercentThreshold:  25,
		RequestVolumeThreshold: 10,
		SleepWindow:            10000,
	})

	return &Service{
		timeout:    2 * time.Second,
		maxRetries: 3,
	}
}

func (s *Service) Call() error {
	return hystrix.Do("api_service", func() error {
		// 正常业务逻辑
		return s.doAPICall()
	}, s.fallbackFunction)
}

func (s *Service) doAPICall() error {
	// 实际的 API 调用
	client := &http.Client{
		Timeout: s.timeout,
	}

	resp, err := client.Get("http://api.example.com/data")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	return nil
}

func (s *Service) fallbackFunction(err error) error {
	// 降级处理逻辑
	fmt.Printf("Fallback triggered due to error: %v\n", err)
	return fmt.Errorf("service temporarily unavailable")
}

func main() {
	service := NewService()

	// 模拟多次调用
	for i := 0; i < 5; i++ {
		err := service.Call()
		if err != nil {
			fmt.Printf("Call %d failed: %v\n", i+1, err)
		} else {
			fmt.Printf("Call %d succeeded\n", i+1)
		}
		time.Sleep(500 * time.Millisecond)
	}
	for {
	}
}