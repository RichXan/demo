## 项目目录结构

project/
├── cmd/
│   └── server/
│       └── main.go           # 主启动文件，运行服务器
|
├── config/
│   ├── config.yaml           # 配置文件
│   └── config.go             # 配置文件解析代码
|
├── internal/
|   |
│   ├── access/               # 数据库访问层(DAL层：Data Access Layer)
│   │   ├── entity/           # 用于保存数据库表结构 ✅ 
│   │   └── repository/       # 数据访问层          ✅ 直接访问数据库。增删改查数据。封装了数据库操作，提供给业务逻辑层使用。
|   |
│   ├── api/                  # 向外部提供的接口定义 ✅
│   │   ├── biz/              # 外部接口调用方法定义 ✅
│   │   └── model/            # 外部接口调用请求参数和返回参数定义 ✅
|   |
│   ├── busi/                 # business 业务逻辑层 ✅
│   │   ├── dto/              # 数据传输对象 ✅
│   │   └── service/          # 服务层 ✅
|   |
│   ├── handlers/             # HTTP处理器 ✅
│   │   └── user_handler.go   # 用户模块处理器 ✅
│   │   
│   ├── models/               # 数据模型定义
│   │   ├── user.go           # 用户模型
│   │   └── auth.go           # 认证模型
│   ├── repositories/         # 数据访问层
│   │   ├── user_repo.go      # 用户数据访问层
│   │   └── auth_repo.go      # 认证数据访问层
│   ├── services/             # 业务逻辑层
│   │   ├── user_service.go   # 用户服务层
│   │   └── auth_service.go   # 认证服务层|
|   |
│   └── jobs/                 # 定时任务定义 ✅
│       └── product_job.go    # 产品定时任务 ✅
|
├── templates/                 # 模板文件
├── go.mod                     # Go模块依赖文件
├── go.sum                     # Go模块依赖校验文件
└── Makefile                   # Makefile脚本，用于自动化构建、测试等任务

