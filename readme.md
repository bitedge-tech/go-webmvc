# go-webmvc

[![Go](https://img.shields.io/badge/go-1.24.6-blue)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

中文为主文档（English summary at the end）

---

## 简介

`go-webmvc` 是一个用于学习与快速启动的 Go Web 服务模板，采用典型的企业应用分层：handler → service → repository → model。项目集成并演示了常见组件的接入方案，包括：

- Gin（HTTP 框架）
- GORM（MySQL）与 gorm-gen（代码生成示例）
- Redis 客户端封装
- 结构化日志（zap + lumberjack）
- WebSocket 示例、验证码（captcha）示例
- 多阶段 Docker 构建示例

目标受众：想要一套轻量、可改造的 Go 服务骨架，用于学习或作为内部模板。

> 说明： 本项目只是整合了目前使用Go开发web项目中最常用的一些工具包,并按MVC的模式进行了代码的组织,主要目的是方便使用者更快速的搭建项目,快速进入业务代码的开发,不对代码做任务限制,使用者可对代码进行任务的修改和扩展。

---

## 目录（快速导航）

- [特性](#特性)
- [快速开始（3 分钟上手）](#快速开始3-分钟上手)
- [项目结构（按实际代码）](#项目结构按实际代码)
- [配置详例（YAML & 环境变量）](#配置详例yaml--环境变量)
- [本地开发与构建](#本地开发与构建)
- [Docker 与部署](#docker-与部署)
- [排错与常见问题](#排错与常见问题)
- [测试](#测试)
- [CI / CD 建议](#ci--cd-建议)
- [贡献指南](#贡献指南)
- [维护者与联系方式](#维护者与联系方式)
- [许可证](#许可证)
- [English summary](#english-summary)

---

## 特性

- 清晰的目录分层，便于开发与维护
- 封装数据库（GORM）与 Redis 客户端初始化逻辑
- 日志系统使用 `zap`，支持文件切割（lumberjack）
- 包含登录、验证码、用户接口示例
- 支持 gorm-gen 代码生成（`internal/repository/query`）
- Docker镜像构建和运行示例

---

## 快速开始（3 分钟上手）

1. 克隆仓库并进入项目目录：

```bash
git clone <your-repo-url> go-webmvc
cd go-webmvc
```

2. 准备配置（参考下文 `配置详例`）：编辑 `config/config.dev.yaml` 或通过环境变量注入。

3. 开发模式运行：

```cmd
cd %~dp0
go run ./cmd/server
```

4. 在容器中运行（推荐用于与生产环境一致的测试）：

```cmd
set "GOOS=linux" && set "GOARCH=amd64" && go build -ldflags "-s -w" -o build/adminServer ./cmd/server
docker build -t go-webmvc-admin:latest .
```

---

## 项目结构（按实际代码）

仓库主要目录（摘录）：

- `cmd/`
  - `server/main.go`：应用入口，负责配置加载、各组件初始化（日志、DB、Redis）、路由注册与启动。
  - `ws.go`：独立 WebSocket 示例程序。
  - `gen/`：代码生成相关（gorm-gen）或生成脚本。
- `config/`：配置加载器与示例 YAML（`config.go`, `config.dev.yaml`）。
- `internal/`
  - `handler/`：HTTP 层处理函数（login、users、index 等）。
  - `service/`：业务逻辑层。
  - `repository/`：持久层（`model/` + `query/`）。
  - `router/`：路由注册（`internal/router/router.go`）。
  - `middleware/`：中间件（如 JWTAuth）。
- `pkg/`
  - `pkg/db/`：MySQL 初始化与迁移（`mysql.go`）。
  - `pkg/logger/`：zap 日志封装（`logger.go`）。
  - `pkg/redis/`：Redis 客户端（`redis.go`）。
  - `pkg/natCon/`：NATS 连接示例（可选）。
- `build/`：二进制产物目录（`build/adminServer`）。
- `Dockerfile`：镜像构建脚本（建议按 README 示例修改为多阶段构建）。

---

## 配置详例（YAML & 环境变量）

示例 `config/config.dev.yaml`（请按实际 `config.go` 字段调整）：

```yaml
app:
  env: development
  port: "8080"

database:
  host: 127.0.0.1
  port: 3306
  user: root
  password: password
  name: go_webmvc

redis:
  host: 127.0.0.1
  port: 6379
  password: ""

log:
  output: stdout # stdout | file | both
  level: info
  file:
    filename: logs/app.log
    max_size: 100
    max_backups: 7
    max_age: 30
    compress: true
```

推荐的环境变量覆盖（示例）：

- APP_ENV
- APP_PORT
- DB_HOST / DB_PORT / DB_USER / DB_PASSWORD / DB_NAME
- REDIS_HOST / REDIS_PORT / REDIS_PASSWORD
- LOG_OUTPUT / LOG_LEVEL

（项目中的 `config.LoadConfig()` 会读取 YAML 并支持环境变量覆盖，具体键名请参考 `config/config.go`）

---

## 本地开发与构建

安装依赖并构建：

```cmd
cd %~dp0
go mod download
set "GOOS=linux" && set "GOARCH=amd64" && go build -ldflags "-s -w" -o build/adminServer ./cmd/server
```

- 若希望调试（拥有符号），可去掉 `-ldflags` 并将 `GOOS`/`GOARCH` 设为当前平台。
- 推荐在开发时使用本地 MySQL/Redis 或 Docker Compose（下节示例）模拟运行环境。

---

## Docker 与部署

示例 `Dockerfile`（multi-stage, 推荐）已在本仓库 README 中提供。下面给出 `docker-compose.yml` 示例（便于本地联调）：

```yaml
version: '3.8'
services:
  db:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: go_webmvc
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  app:
    build: .
    image: go-webmvc-admin:local
    depends_on:
      - db
      - redis
    environment:
      - DB_HOST=db
      - DB_PORT=3306
      - DB_USER=root
      - DB_PASSWORD=password
      - DB_NAME=go_webmvc
      - REDIS_HOST=redis
      - REDIS_PORT=6379
    ports:
      - "8080:8080"

volumes:
  db_data:
```

运行：

```bash
docker compose up --build
```

生产环境建议：
- 使用 CI 构建并推送镜像到私有 registry
- 使用 Kubernetes / Nomad 等编排工具管理部署
- 使用 Secret 管理敏感信息（K8s Secrets / Vault / AWS Parameter Store 等）

---

## 排错与常见问题

1. `docker build` 报错无法拉取基础镜像：通常是宿主机网络/代理或 Docker Desktop 未配置代理，参考下列排查：
   - 在宿主机运行 `docker pull alpine:3.18` 或 `curl https://registry-1.docker.io/v2/` 检查网络。
   - 在 Docker Desktop 中配置代理或镜像加速器（国内环境可使用阿里云镜像加速器）。
   - 使用 `docker login` 避免匿名拉取限制。

2. MySQL 连接失败：确认 `config` 中 DB 主机/端口/用户名/密码，若使用 Docker Compose，服务名作为 host（例如 `db`）。

3. 日志无法写入：确认运行用户对配置的日志路径有写权限（容器中建议写到 `/var/log/app` 或 stdout）。

---

## 测试

仓库当前未包含广泛单元测试（欢迎补充）。推荐测试策略：

- service 层：使用接口 + mock 依赖进行单元测试
- repository 层：使用测试数据库或 in-memory 模式进行集成测试
- 集成测试：使用 Docker Compose 在 CI 中启动依赖（MySQL / Redis），运行端到端测试

运行测试：

```cmd
go test ./... -v
```

---

## CI / CD 建议

建议使用 GitHub Actions 或其它 CI：
- 执行 `go test`、`go vet`、`golangci-lint`（如启用）
- 构建多平台镜像（建议使用 `docker/build-push-action` + Buildx）
- 将镜像推送到私有 registry，并在部署环境拉取

示例工作流步骤（简述）：
- Checkout
- Set up Go
- Cache Go modules
- Run tests
- Docker build & push

---

## 贡献指南

欢迎贡献！建议流程：

1. Fork 仓库 → 新增分支（feature/xxx 或 fix/xxx）
2. 编写代码，添加/修改单元测试
3. 提交 PR，并在说明中列出更改点、测试方式与影响范围
4. 维护者会在 PR 中 review 并合并

请在变更 API 或结构前先提交 Issue 讨论设计。

---

## 维护者与联系方式

- 维护者：项目作者（请在仓库中补充实际者信息）
- 联系方式：在仓库 Issue 中提出或通过 PR 讨论

---

## 许可证

本项目建议使用 MIT 许可证。若要发布，请在仓库根添加 `LICENSE` 文件并替换成你需要的许可证文本。

---

## 致谢

感谢以下开源项目：Gin, GORM, zap, gorilla/websocket, gorm-gen 等。

---

## English summary

`go-webmvc` is a minimal Go web service template using Gin + GORM + Redis + zap. It includes example Docker multi-stage builds and a WebSocket demo. Build with `go build -o build/adminServer ./cmd/server` (set GOOS/GOARCH for linux builds). Recommended deployment pattern: CI builds image → push to registry → orchestrate with Kubernetes / Docker Compose.

---

如果你同意，我可以：
- 在仓库根创建 `LICENSE`（MIT）文件；
- 添加一个基本的 GitHub Actions workflow（测试 + 构建镜像并不推送）；
- 修改仓库 `Dockerfile` 为 README 中的 multi-stage 示例并提交更改。

请选择想要我继续的项（可多选）。
