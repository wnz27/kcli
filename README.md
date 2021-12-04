# kcli  [[en]](README_en.md)
命令行工具，希望覆盖大部分工作场景

借鉴 [cli](https://github.com/urfave/cli) 现成的实现，来植入日常工作所需的功能。

## use
```shell
git clone https://github.com/wnz27/kcli.git
cd kcli
go build
./kcli
```
(windows 应该是 .exe 文件)
显示类似这样:
```shell
./kcli                                                              
NAME:
   kcli command tool - kcli 通用命令行工具

USAGE:
   kcli [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   install   安装 kcli 命令行工具所有依赖 （Todo）
   run       启动微服务
   build     编译微服务
   gen-grpc  根据 proto 文件 自动生成代码文件
             use: protoc --proto_path xxx --go_out=plugins=grpc: xxxx xxxx.proto

   lint-fmt  修复这类 lint 报错: File is not `gofmt`-ed with `-s` (gofmt)
             use: gofmt -s -w {{ lint-file }}

   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --service_name value, --sn value      服务名称可以为: 
                                         1、service1 
                                         2、service2 
                                         3、service3
   --config value, -c value              当前仅支持json文件, 默认为当前目录下的config.json文件
                                         没有指定的话会 默认为 在 internal/服务名称/ 的路径下拼接config.json 文件
   --env value                           app运行环境, 当前仅支持: dev、test、prod (default: "dev")
   --mode value, -m value                app运行模式, 仅支持debug, release (default: "debug")
   --grpc-gen-path value, --gpath value  grpc 自动生成的代码文件的输出位置(建议绝对路径, 如果相对路径则根据本命令行执行位置来判断即可)
                                         （暂默认 proto 文件也在这个位置）默认是: ./common/meta
 (default: "./common/meta")
   --proto-file-name value, --pfile value  目标 proto 文件名称, "service 和 client 在同一个文件
   --lint-file value, --lfmt value         目标 需要执行 gofmt 的文件 路径
   --help, -h                              show help (default: false)
   --version, -v                           print the version (default: false)
```

## TODO
### 开发迭代点
- [ ] 默认设置抽出去同一放到default中去, 可读性更好
- [ ] 如何 mac 和 windows 使用都友好，（目前功能俩都能用）
- [ ] color string 优化
### 功能
- 编译
  - [X] 基础
  - [ ] 增加编译参数
- 运行
  - [X] 基础
- 生成项目模板，可以同时支持多种协议
  - [X] grpc
  - [ ] thrift
  - [ ] http
  - [ ] mqtt
- 根据 MySQL 的 information schema 库自动生成基本的表增、删、改、查功能的 dao 模块(参考 [beego](https://github.com/beego/beego))
  - [ ] sqlx
  - [ ] gorm
  - [ ] xorm
- lint 某些报错修复
  - [ ] lint安装相关
  - [X] golang-lint: "File is not `gofmt`-ed with `-s` (gofmt)"
- 文档
  - [ ] 生成 swagger 注释 (方向利用 ast 模块解析 func 的注释)
- 依赖相关
  - 安装 kcli 依赖
    - [ ] protoc
    - [ ] protoc-gen-go 
  - [ ] 集成 wire 依赖注入
- 基础接口模板
  - [ ] MQ
  - [ ] DB
  - [ ] 注册中心
  - [ ] 配置中心
  - [ ] Cache
- 参考 go-micro 接口设计，支持一些主流框架的一键按配置切换任意的外部实现
  - 模块
      - [ ] MQ
      - [ ] DB
      - [ ] 注册中心
      - [ ] 配置中心 
      - [ ] Cache
  - [ ] gin (先用gin 趟趟水)
  - [ ] kratos
  - [ ] go-zero
  - [ ] beego
  - [ ] ego
