# hystrix
## 0. 前言
### 服务雪崩
多个服务之间在调用的时候，假设微服务A调用微服务B和微服务C，微服务B和微服务C又调用其它的微服务，这就是所谓的“扇出”。如果扇出的链路上某个微服务的调用响应时间过长或者不可用，对微服务A的调用就会占用越来越多的系统资源，进而引起系统崩溃，所谓“雪崩效应”。

简单来说就是：一个服务失败，导致整条链路的服务都失败的情形，我们称之为服务雪崩。
### 引起服务雪崩的原因
- 硬件故障
- 程序Bug
- 缓存击穿（用户大量访问缓存中没有的键值，导致大量请求查询数据库，使数据库压力过大）
- 用户大量请求

### 服务雪崩的三个阶段
- 第一阶段：服务端不可用
- 第二阶段：调用端重试加大流量（用户重试/代码逻辑重试）
- 第三阶段：服务调用者不可用（同步等待造成的资源耗尽）

### 服务雪崩的解决方案
- 应用扩容（扩大服务器承受力）
- 流量控制（超出限定流量，返回类似重试页面让用户稍后再试）
- 缓存（使用缓存，减少数据库查询）
- 服务降级（服务接口拒绝服务、页面拒绝服务、延迟持久化、随机拒绝服务）
- 服务熔断（服务端设置熔断机制，当服务端不可用时，快速失败）


## 什么是Hystrix？
Hystrix [hɪst’rɪks]，中文含义是豪猪，因其背上长满棘刺，从而拥有了自我保护的能力。

Hystrix 是在**客户端**使用的熔断器实现。

Hystrix是一个用于处理分布式系统中延迟和容错的库。它通过隔离服务之间的访问、提供回退机制、监控和控制故障传播，从而提高系统的可靠性和容错能力。Hystrix能够保证在一个依赖出问题的情况下，不会导致整个服务失败，避免级联故障，以提高分布式系统的弹性。

> “断路器”本身是一种开关装置，当某个服务单元发生故障之后，通过断路器的故障监控（类似熔断保险丝），向调用方返回一个符合预期的、可处理的预备响应（FallBack），而不是长时间等待或者抛出调用方无法处理的异常，这样就保证了服务调用方的线程不会被长时间、不必要的占用，从而避免了故障在分布式系统中蔓延，乃至雪崩。

```text
客户端 (使用 Hystrix) ---> 调用 ---> 服务端
         |
         |---> 熔断器监控
         |---> 超时控制
         |---> 降级处理
```

## 1. 为什么需要Hystrix？

在分布式系统中，服务之间的调用不可避免地会遇到各种问题，如网络延迟、服务不可用、超时等。这些问题会导致请求失败，甚至引发级联故障，最终导致整个系统瘫痪。Hystrix通过以下机制来解决这些问题：
- 隔离：通过将每个依赖服务隔离在单独的线程池中，防止单个服务的故障影响其他服务。
- 回退机制：提供默认的响应策略，当依赖服务不可用时，可以快速失败并返回默认值，避免请求长时间等待。
- 熔断器：监控服务的健康状况，当某个服务的错误率超过一定阈值时，自动熔断该服务的请求，避免故障扩散。
- 监控和控制：提供详细的监控和统计信息，帮助开发人员及时发现和解决问题，优化系统性能。

### 重要概念
**服务降级(fallback)**

> 1、当某个服务单元发生故障之后，通过断路器的故障监控（类似熔断保险丝），**向调用方返回一个符合预期的、可处理的预备响应（FallBack）**，而不是长时间等待或者抛出调用方无法处理的异常。比如：服务繁忙，请稍后再试，不让客户端等待并立刻返回一个友好提示：fallback。
> 
> 2、哪些情况会触发降级
> 1. 程序运行异常 
> 2. 超时
> 3. 服务熔断触发服务降级
> 4. 线程池/信号量打满也会导致服务降级


**服务熔断(break)**
> 1、系统发到最大服务访问量后，直接拒绝访问，限制后续的服务访问，并调用服务降级方法返回友好提示。
> 
> 2、就是保险丝：服务降级–>进而熔断–>恢复调用链路

**服务限流(rate limiting)**
> 限流的目的是为了保护系统不被大量请求冲垮，通过限制请求的速度来保护系统。在电商的秒杀活动中，限流是必不可少的一个环节。限制高并发，请求进行排队，一秒处理N个请求，有序的进行。

### 2. 为什么要在客户端使用Hystrix？
1. 及时熔断
   - 可以直接在服务消费方进行熔断
   - 避免请求继续发送到已经出问题的服务端
   - 减少网络资源的浪费
2. 精细控制
   - 每个客户端可以根据自己的需求设置不同的熔断策略
   - 可以为不同的服务设置不同的超时时间和重试策略
3. 降级处理
   - 在客户端直接进行降级处理
   - 可以立即返回备用响应或缓存数据

## 2. 使用Hystrix

### 1. 使用示例

```go
package main

import (
    "github.com/afex/hystrix-go/hystrix"
)

// UserService 客户端
type UserService struct {
    baseURL string
}

func NewUserService(baseURL string) *UserService {
    // 为用户服务配置熔断器
    hystrix.ConfigureCommand("get_user", hystrix.CommandConfig{
        Timeout:                1000,
        MaxConcurrentRequests: 100,
        ErrorPercentThreshold: 25,
    })

    return &UserService{
        baseURL: baseURL,
    }
}

func (s *UserService) GetUser(id string) (*User, error) {
    var user *User
    output := make(chan bool, 1)
    
    // 同步调用
    err := hystrix.Do("get_user", func() error {
        // 实际的服务调用
        return s.doGetUser(id)
    }, func(err error) error {
        // 降级处理：返回缓存数据或默认用户
        return s.getFallbackUser(id)
    })

    // 异步调用
    errors := hystrix.Go("get_user", func() error {
        output <- true
        return s.doGetUser(id)
    }, func(err error) error {
        return s.getFallbackUser(id)
    })

    select {
    case out := <-output:
        // success
    case err := <-errors:
        // failure
    }
    return user, err
}
```

### 合理的命名、适当的配置
```go
// 为不同的服务使用不同的command名称
hystrix.ConfigureCommand("userService", ...)
hystrix.ConfigureCommand("orderService", ...)
hystrix.ConfigureCommand("paymentService", ...)

// 根据业务场景配置合适的参数
hystrix.ConfigureCommand("critical_service", hystrix.CommandConfig{
    // 关键服务可以设置更长的超时时间
    Timeout: 2000,
    // 更高的错误阈值
    ErrorPercentThreshold: 50,
    // 更长的恢复时间
    SleepWindow: 10000,
})
```

### 监控集成
```go
// 添加监控
hystrixStreamHandler := hystrix.NewStreamHandler()
hystrixStreamHandler.Start()
go http.ListenAndServe(":8081", hystrixStreamHandler)
```
通过在客户端使用 Hystrix，可以有效地保护客户端免受依赖服务故障的影响，提高系统的可用性和稳定性。


## 2.2 常见使用场景

### 1. 微服务调用
```go
// 服务A调用服务B
serviceB := NewServiceBClient("http://service-b")
hystrix.ConfigureCommand("service_b_api", hystrix.CommandConfig{
    Timeout: 1000,
})

// 使用熔断器包装调用
err := hystrix.Do("service_b_api", func() error {
    return serviceB.Call()
}, nil)
```

### 2. 第三方API调用
```go
// 调用外部支付API
paymentAPI := NewPaymentAPI()
hystrix.ConfigureCommand("payment_api", hystrix.CommandConfig{
    Timeout: 3000,
    ErrorPercentThreshold: 20,
})

// 包装支付调用
err := hystrix.Do("payment_api", func() error {
    return paymentAPI.ProcessPayment()
}, func(err error) error {
    // 支付失败时的降级处理
    return useBackupPaymentSystem()
})
```

### 3. 数据库操作
```go
// 数据库操作包装
hystrix.ConfigureCommand("db_operation", hystrix.CommandConfig{
    Timeout: 500,
})

err := hystrix.Do("db_operation", func() error {
    return db.Query()
}, func(err error) error {
    // 使用缓存数据
    return cache.Get()
})
```
参考：
- [Hystrix断路器原理及实现（服务降级、熔断、限流）](https://blog.csdn.net/qq_36763419/article/details/120119872)