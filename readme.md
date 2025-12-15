# go-webmvc

[![Go](https://img.shields.io/badge/go-1.24.6-blue)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

中文文档

---

## 简介

**go-webmvc** 是一个基于于golang,整合了主流的组件和工具库的MVC开发框架.采用典型的企业应用分层：handler → service → repository → model.

**解决什么问题?** 帮助开发者快速搭建一个结构清晰、易于维护的 Go Web后端项目骨架， 避免每次新建项目都要从零开始配置各种常用组件;省时省力,统一项目结构, 提升开发效率, 对新手入门go开发友好,开箱即用。

项目集成并演示了常见组件的接入方案， 包括：
- Gin（HTTP 框架）
- GORM（MySQL）与 gorm-gen（代码生成示例）
- Viper 配置管理
- Redis 客户端封装
- JWT 鉴权中间件
- NaTS 连接示例
- 结构化日志（zap + lumberjack）

目标受众：想要一套轻量、可改造的 Go 服务骨架，用于学习或作为内部模板。

> 说明： 本项目只是整合了目前使用Go开发web项目中最常用的一些组件,并按MVC模式进行代码的组织,主要目的是方便使用者更快速的搭建项目,快速进入业务代码的开发,项目不对代码做任何限制,使用者可对代码进行任何的修改和使用。

---

## 目录（快速导航）
- [简介](#简介)
- [特性](#特性)
- [快速开始（3 分钟上手）](#快速开始3-分钟上手)
- [项目结构（按实际代码）](#项目结构按实际代码)
- [本地开发入门 (更新中...)](#本地开发入门-更新中)
  - [从最简单的API开始](#从最简单的api开始)
  - [如何新增一个http请求处理器示例](#如何新增一个http请求处理器示例)
  - [配置路由示例](#配置路由示例)
  - [扩展一个service示例](#扩展一个service示例)
  - [新增一个数据库表示例](#新增一个数据库表示例)
  - [使用 配置示例](#使用-配置示例)
  - [生成swagger文档示例](#生成swagger文档示例)
  - [使用 Redis 示例](#使用-redis-示例)
  - [使用 Nats 示例](#使用-nats-示例)
- [贡献指南](#贡献指南)
- [许可证](#许可证)
- [致谢](#致谢)
- [English summary](#english-summary)

---

## 特性

- 清晰的目录分层，便于开发与维护
- Gin 框架，支持中间件扩展
- 封装数据库（GORM）与 Redis 客户端初始化逻辑
- 支持 gorm-gen 代码生成（`internal/repository/query`）
- 集成 Viper 配置管理，支持 YAML 文件与环境变量覆盖
- 日志系统使用 `zap`，支持文件切割（lumberjack）
- 包含登录、验证码、用户接口示例
- Docker镜像构建和运行示例

---

## 快速开始（3 分钟上手）

>环境要求：已装好 Go语言环境（1.24.6 及以上），根据需要安装 MySQL 和 Redis 。

1. 克隆仓库并进入项目目录：

```bash
git clone git@github.com:bitedge-tech/go-webmvc.git go-webmvc
cd go-webmvc
```

2. 准备配置（参考下文 `配置详例`）：编辑 `config/config.dev.yaml` 
> 注意：请确保本地有 MySQL 和 Redis 实例在运行; 如果不需要redis,可以在server/main.go中注释掉redis的初始化代码。

3. 安装依赖:

```cmd
cd go-webmvc #项目根目录下
go mod tidy
``` 

4. 运行程序:

```cmd
go run ./cmd/server/main.go
```
- 看到  Listening and serving HTTP on :8080 表示启动成功。
- 在浏览器访问 http://localhost:8080/   页面显示: "Welcome to Go WebMVC!" 说明服务运行正常。

## 项目结构

仓库主要目录（摘录）：

- `cmd/` ：应用入口
  - `server/main.go`：应用入口，负责配置加载、各组件初始化（日志、DB、Redis）、路由注册与启动。
  - `gen/`：代码生成相关（gorm-gen）或生成脚本。
- `config/`：配置加载器与示例 YAML（`config.go`, `config.dev.yaml`）。
- `internal/` ：核心业务代码都放在这里
  - `handler/`：HTTP 层处理函数（login、users、index 等）。
  - `service/`：业务逻辑层。
  - `repository/`：持久层（`model/` + `query/`）。
  - `router/`：路由注册（`internal/router/router.go`）。
  - `middleware/`：中间件（如 JWTAuth）。
- `pkg/` ：第三方组件封装
  - `pkg/db/`：MySQL 初始化与迁移（`mysql.go`）。
  - `pkg/logger/`：zap 日志封装（`logger.go`）。
  - `pkg/redis/`：Redis 客户端（`redis.go`）。
  - `pkg/natCon/`：NATS 连接示例。
- `build/`：编译生成的二进制文件,可以对应系统中直接运行（`build/app1`）。
- `Dockerfile`：镜像构建脚本,用于生成docker镜像运行在docker下。

---

## 开发入门 (更新中...)

### 1.从最简单的API开始
**先看下项目的 http://127.0.0.1:8080/ 返回 "Welcome to Go WebMVC!" 的实现.**
* 在handler层中,internal/handler/index/index_handler.go文件下新增Index函数:
```golang
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "welcome to go-webmvc!")
}

```
* 在router中,internal/router/router.go文件下注册路由:
```golang
r.GET("/", index.Index)
``` 
* 运行项目,访问 http://127.0.0.1:8080/ 即可看到返回结果.
<br>
<br>

### 2.实现一个查询用户信息的完整API示例.

>1 首先需要有user表, 我们通过使用model结构体来生成数据库表.
(**前提:确保项目已经连接到数据库**);
(1) 在 model 层定义 User 结构体（`internal/repository/model/user.go`）:
```golang
package model
import (
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(50);not null" json:"username"` //用户名
	Password  string    `gorm:"type:varchar(50); " json:"-"`               //密码
	Nickname  string    `gorm:"type:varchar(50); " json:"nickname"`        //昵称
	Salt      string    `gorm:"type:varchar(100);" json:"-"`               //密码盐值
	Phone     string    `gorm:"type:varchar(20); " json:"phone"`           //手机号
	Email     string    `gorm:"type:varchar(100); " json:"email"`          //邮箱
	Avatar    string    `gorm:"type:varchar(255); " json:"avatar"`         //头像URL
	Status    int       `gorm:"type:int ;default:0" json:"status"`         //0:保存, 1:启用, 9:禁用
	RoleID    int64     `gorm:"type:bigint; " json:"role_id"`              //角色ID
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"create_at"`
	CreatedBy int64     `gorm:"type:bigint; " json:"create_by"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:updated_at;" json:"update_at"`
	UpdatedBy int64     `gorm:"type:bigint; " json:"update_by"`
}
func (User) TableName() string { return "user" }

```
(2) 将 mode.User添加到数据库迁移中,在`pkg/db/mysql.go`文件的Migrate函数中添加:
```golang

func Migrate(db *gorm.DB) error {
    // 自动迁移模式
    return db.AutoMigrate(
        &model.User{}, //新增User表
        // 其他模型...
    )
}
```
(3) 运行项目,程序启动时会自动创建user表. (完成数据库表的创建)
```shell
 go run ./cmd/server/main.go
```
<br>

>2 使用 gorm-gen生成数据库的操作结构体(类)-query.
> 
> 
(1) 在 cmd/gen/main.go 文件中添加 User 表的生成配置:
```golang
g.ApplyBasic(
	model.User{},  //添加这一行
	)

```
(2) 运行代码生成命令:
```shell
go run ./cmd/gen/main.go
```
生成的文件在 `internal/repository/query/user.gen.go`，包含对 user 表的增删改查等操作方法。
<br>

>3 实现查询用户信息的业务逻辑.
(1) 在 handler 层添加用户查询处理器, 在internal/handler/users/user_handler.go`文件中添加函数: UserInfo()
```golang
package users

import (
	"go-webmvc/internal/dto"
	"go-webmvc/internal/handler"
	"go-webmvc/internal/service"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {

	// 1. 参数绑定:从客户端请求中获取用户ID, get请求的参数中
	req := dto.UserInfoRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		handler.Failed(c, "参数绑定失败")
		return
	}

	// 2. 调用服务层获取用户信息
	userService := service.Services.User
	userInfo, err := userService.UserInfo(req.UserID)
	if err != nil {
		handler.Failed(c, "获取用户信息失败")
		return
	}

	// 3. 返回成功响应给客户端
	if userInfo == nil {
		handler.Failed(c, "用户不存在")
		return
	}
	handler.Success(c, userInfo)

}
```
(2) 在 service 层添加用户查询逻辑,在`internal/service/user_service.go`文件中添加函数:UserInfo()
```golang
package service

import (
	"errors"
	"go-webmvc/internal/repository/model"
	"go-webmvc/internal/repository/query"

	"gorm.io/gorm"
)

type UserI interface {
	UserInfo(userID int64) (user *model.User, err error)
}

type userService struct {
}

func (*userService) UserInfo(userID int64) (user *model.User, err error) {

	//从数据库查询用户信息
	q := query.User
	user, err = q.WithContext(nil).Where(q.ID.Eq(userID)).First()

	// 如果查询出错且不是记录未找到错误，则返回错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return user, nil

}

```
(3) 在 dto 层定义请求参数结构体,在`internal/dto/user_dto.go`文件中添加:
```golang
type UserInfoRequest struct {
	UserID int64 `form:"user_id" binding:"required"`
}
```
(4) 在 router 层注册路由,在`internal/router/router.go`文件中添加:
```golang
r.GET("/user/info", users.UserInfo)
```
(5) 运行项目,测试接口:
```shell
 go run ./cmd/server/main.go
```
在浏览器中打开 http://localhost:8080/user/userInfo?user_id=1 查看结果. 在数据库中插入一条 user 记录,即可看到返回的用户信息.

### 用户信息查询的POST请求示例
如果想使用POST请求来查询用户信息,只需修改3个地方:
1.在dto中将参数绑定标签改为json:
```golang
type UserInfoRequest struct {
    UserID int64 `json:"user_id" binding:"required"`
}
```

2.在handler中将参数绑定方式改为ShouldBindJSON,其他代码不变.
```golang
// 1. 参数绑定:从客户端请求中获取用户ID, post请求的json体中
req := dto.UserInfoRequest{}
if err := c.ShouldBindJSON(&req); err != nil {
    handler.Failed(c, "参数绑定失败") 
    return
}
// 其他代码不变
```
3.在router中将路由方法改为POST:
```golang
r.POST("/user/info", users.UserInfo)
``` 
通过以上3个修改,即可实现POST请求查询用户信息. 重启服务后,使用POST请求访问 http://localhost:8080/user/info ,请求体为json



### 3.使用配置示例
项目的配置信息存放在 config 目录下,主要文件有:
- `config.go`：配置加载器，使用 Viper 实现。
- `config.dev.yaml`：开发环境的示例配置文件。
- `config.prod.yaml`：生产环境的示例配置文件。

在代码中使用配置示例:

(1) 在 config.yaml 文件中添加需要的配置信息;

(2) 在 config.go 文件中的 Config 结构体中添加对应的字段;

(3) 在代码中通过 config.AppConfig 变量获取需的要配置信息.

例如,获取服务器端口号:
```golang
port := config.AppConfig.App.Port
```


### 4.生成swagger文档示例

### 5.使用 Redis 示例

### 6.使用 Nats 示例


---

## 贡献指南

欢迎贡献！建议流程：

1. Fork 仓库 → 新增分支（feature/xxx 或 fix/xxx）
2. 编写代码，添加/修改单元测试
3. 提交 PR，并在说明中列出更改点、测试方式与影响范围
4. 维护者会在 PR 中 review 并合并

请在变更 API 或结构前先提交 Issue 讨论设计。

---


## 许可证

本项目建议使用 MIT 许可证。若要发布，请在仓库根添加 `LICENSE` 文件并替换成你需要的许可证文本。

---

## 致谢

感谢以下开源项目：Gin, GORM, zap, gorm-gen 等。

---

## English summary

# go-webmvc
A Go web service template for learning and rapid startup, following a typical enterprise application layering: handler → service → repository → model. The project integrates and demonstrates common component integration schemes, including: 
- Gin (HTTP framework)
- GORM (MySQL) and gorm-gen (code generation example)
- Redis client encapsulation  
- JWT authentication middleware
- NaTS connection example
- Structured logging (zap + lumberjack)
- Target audience: Those looking for a lightweight, customizable Go service skeleton for learning or as an internal template.
- Note: This project simply integrates some of the most commonly used packages in Go web development and organizes the code according to the MVC pattern. The main purpose is to help users quickly set up projects and get into business code development. The project does not impose any restrictions on the code; users can modify and use the code as they wish.



