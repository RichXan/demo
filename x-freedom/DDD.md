# DDD(Domain Driven Design)

## 1. 前言

DDD 是领域驱动设计，是一种软件设计方法论，它强调将复杂的业务领域分解为多个子域，每个子域都有自己的领域模型和设计原则。DDD 的核心思想是通过将业务领域分解为多个子域，每个子域都有自己的领域模型和设计原则，从而提高系统的可维护性和可扩展性。

## 2. 使用

### 2.1 创建领域服务

使用freedom框架创建领域服务

#### 安装
`$ go install github.com/8treenet/freedom/freedom@latest`
`$ freedom version`

#### 脚手架创建项目

`$ freedom new-project [project-name]`
`$ cd [project-name]`
`$ go mod tidy`
`$ go run server/main.go`

#### 脚手架生成增删查改和持久化对象
##### 查看更多
`$ freedom new-po -h`

##### 进入项目目录
`$ cd [project-name]`

##### 数据库数据源方式
`$ freedom new-po --dsn "root:123123@tcp(127.0.0.1:3306)/freedom?charset=utf8"`

##### JSON 数据源方式
`$ freedom new-po --json ./domain/po/schema.json`
