# 1. grpc
## 什么是grpc
grpc是google开源的一款RPC框架，基于HTTP/2协议设计，支持多语言，支持流式处理，支持双向流式处理。

> 这个时候就有聪明的同学要问到了，那么什么是RPC框架呢？

## 什么是RPC框架
RPC框架是远程过程调用（Remote Procedure Call）的缩写，它允许一个程序调用另一个程序的函数或方法，就像本地调用一样。RPC框架通常用于分布式系统中，使得不同的服务可以相互通信和协作。

简单来说，RPC框架就是一种用于实现分布式系统中服务间通信的机制。在A服务中调用B服务中的方法，就像调用本地的函数一样简单。

> 这个时候又有聪明的同学要问了，啊，那两个服务之间的通信是使用什么格式来通信交流的呢？

很聪明啊这个同学。他们是使用protobuf来通信交流的。

> 这个时候又有聪明的同学要问了，啊，protobuf是什么？

## 什么是protobuf
protobuf也叫Protocol Buffers，是google开源的一款数据序列化工具，可以将结构化数据序列化为二进制格式，也可以将二进制格式反序列化为结构化数据。它独立于语言，独立于平台。google 提供了多种语言的实现：java、c#、c++、go 和 python，每一种实现都包含了相应语言的编译器以及库文件。

由于它是一种二进制的格式，比使用 xml 、json进行数据交换快许多。可以把它用于分布式应用之间的数据通信或者异构环境下的数据交换。作为一种效率和兼容性都很优秀的二进制数据传输格式，可以用于诸如网络传输、配置文件、数据存储等诸多领域。

> 那在golang中，我该如何使用grpc呢？

# 2. 环境搭建

## 1. 安装protoc

从 [Protobuf Releases](https://github.com/protocolbuffers/protobuf/releases) 下载最先版本的发布包安装。如果是 Ubuntu，可以按照如下步骤操作

**下载安装包**
```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
```
**解压到 /usr/local 目录下**
```bash
sudo 7z x protoc-3.11.2-linux-x86_64.zip -o/usr/local
```
如果不想安装在 /usr/local 目录下，可以解压到其他的其他，并把解压路径下的 bin 目录 加入到环境变量即可。

如果能正常显示版本，则表示安装成功。

```bash
$ protoc --version
libprotoc 3.11.2
```


## 2. 安装protoc-gen-go & protoc-gen-go-grpc
**protoc-gen-go**： 这个工具用来将 .proto 文件转换为 Golang 代码。

**protoc-gen-go-grpc**： 是一个用于生成 gRPC 代码的插件，它与 Protocol Buffers 编译器 (protoc) 一起使用。它可以根据 .proto 文件生成 Go 语言的 gRPC **服务代码**。

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
protoc-gen-go 将自动安装到 $GOPATH/bin 目录下，也需要将这个目录加入到环境变量中。若使用时失败，请将 protoc-gen-go 和 protoc-gen-go-grpc 的 bin 目录加入到环境变量中。

`go env GOPATH` ： 查看 GOPATH 路径，将 GOPATH 路径下的 bin 目录加入到环境变量中。

`export PATH=$PATH:/root/go/bin` ： 将 GOPATH 路径下的 bin 目录加入到环境变量中。

`source ~/.bashrc` ： 使环境变量生效。


## 3. 编译proto
```bash
protoc --go_out=.  --go-grpc_out=. ./*.proto
```

# 3. 进阶了解ProtoBuf
> tips: 写proto和写go最大的区别是需要在结尾添加分号的;，在开发过程中给自己提个醒：如果是写proto需要加分号，如果是写go不需要加分号。

```proto
syntax = "proto3"; // 指定proto版本

package pay; // 指定包名

option go_package = "pay/proto"; // 指定go包名

message PaymentRequest { // 定义消息体
    int64 amount = 1;
    string out_trade_no = 2;
    repeated string goods = 3; // 定义一个重复字段，相当于go中的数组
}
```
## 3.1. 关键字
- **syntax**：是必须写的，而且要定义在第一行；目前proto3是主流，不写默认使用proto2
- **package**：定义我们proto文件的包名
- **option go_package**：定义生成的pb.go的包名，我们通常在proto文件中定义。如果不在proto文件中定义，也可以在使用protoc生成代码时指定pb.go文件的包名
- **message**：非常重要，用于定义消息结构体，不用着急，下文会重点讲解

参考资料：
[grpc-go](https://grpc.io/docs/languages/go/quickstart/)、
[极客兔兔 Go Protobuf 简明教程](https://geektutu.com/post/quick-go-protobuf.html)

## 3.2. protobuf 数据类型
- 基本数据类型：int32、int64、float、double、bool、string
- 复合数据类型：message
- 特殊数据类型：repeated，用于定义一个重复字段，相当于go中的数组
- 枚举数据类型：enum
- 其他数据类型：Any, map<>, oneof

```proto
// 定义一个枚举
enum PaymentStatus {
    SUCCESS = 0;
    FAILED = 1;
}

// 定义一个Any
message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}

// 定义一个oneof
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}

// 定义一个map
message MapRequest {
  map<string, int32> points = 1;
}

```

## 3.3. rpc 状态码
#### 成功（OK） （代码：0）
> 描述：RPC 成功完成。这是理想的结果，表明服务器无问题地处理了请求。
#### 错误代码（用户生成）
> 这些代码通常由服务器端的应用逻辑生成，并指示在 RPC 过程中遇到的特定问题。
####  CANCELLED（代码：1）
> 描述：操作被取消，通常是应客户的请求。这可能是由于超时、用户交互或其他原因。
#### UNKNOWN（代码：2）
> 描述：服务器上发生了意外错误，且没有更具体的问题细节。这是一个未预见问题的集合。
#### INVALID_ARGUMENT（代码：3）
> 描述：客户在请求中提供了无效参数。这可能是因为缺少必需字段、数据类型错误或超出预期范围的值。
#### DEADLINE_EXCEEDED（代码：4）
> 描述：请求完成耗时过长，超过了设定的截止时间。这可能是因为服务器处理慢、网络问题或传输的数据量过大。
#### NOT_FOUND（代码：5）
> 描述：服务器上未找到请求的资源（例如文件、数据库条目）。
#### ALREADY_EXISTS（代码：6）
> 描述：尝试创建已经存在的资源。这可能是在尝试插入重复数据或创建具有冲突名称的内容时发生的。
#### PERMISSION_DENIED（代码：7）
> 描述：客户端缺乏执行请求操作的必要权限。这可能是由于访问控制不足或安全设置问题。
#### RESOURCE_EXHAUSTED（代码：8）
> 描述：服务器耗尽了完成请求所需的资源（例如内存、磁盘空间）。
#### FAILED_PRECONDITION（代码：9）
> 描述：由于服务器处于意外状态，请求无法处理。这可能是由于与参数本身无直接相关的请求中的无效数据，或服务器处于不一致状态。
#### ABORTED（代码：10）
> 描述：服务器端中止了操作。这可能是由于服务器实现特定的各种原因。
#### OUT_OF_RANGE（代码：11）
> 描述：请求包含在预期范围之外的值。这可能是一组有效数字之外的数字或不符合允许时间范围的日期。
#### UNIMPLEMENTED（代码：12）
> 描述：服务器不支持请求的 RPC 方法。这可能是因为服务器缺少实现或过时的客户端尝试使用较新功能。
#### INTERNAL（代码：13）
> 描述：发生内部服务器错误。这是在服务器遇到意外问题且无法更具体分类时使用的通用错误代码。