# One Hub 本地开发环境配置指南

## 🎯 开发环境架构

```
┌─────────────────────────────────────────────────────────────┐
│                    本地开发环境                              │
├─────────────────────────────────────────────────────────────┤
│  前端 (React + Vite)     │  后端 (Go + Gin)                │
│  端口: 3010              │  端口: 3000                      │
│  URL: localhost:3010     │  API: localhost:3000/api        │
├─────────────────────────────────────────────────────────────┤
│                Docker Compose 服务                          │
│  MySQL: 3306            │  Redis: 6379                     │
└─────────────────────────────────────────────────────────────┘
```

## 📋 环境要求

### 必需软件
- **Docker** & **Docker Compose** - 数据库服务
- **Go 1.19+** - 后端开发
- **Node.js 16+** - 前端开发  
- **Yarn** - 前端包管理

### 检查安装
```bash
docker --version
docker compose version  # 或 docker-compose --version
go version
node --version
yarn --version
```

## 🚀 快速启动

### 一键启动 (推荐)
```bash
# Linux/macOS
./dev-start.sh start

# Windows
dev-start.bat start
```

### 手动启动
```bash
# 1. 启动数据库服务
docker compose -f docker-compose-dev.yml up -d

# 2. 启动后端 (新终端)
go build -o tmp/one-hub-dev main.go
./tmp/one-hub-dev --config config-dev.yaml

# 3. 启动前端 (新终端)
cd web
yarn install  # 首次运行
yarn dev
```

## 🛠️ 开发环境管理

### 启动脚本命令
```bash
./dev-start.sh start    # 启动完整开发环境
./dev-start.sh stop     # 停止所有服务
./dev-start.sh restart  # 重启所有服务
./dev-start.sh status   # 查看服务状态
```

### 服务访问地址
- **开发环境**: http://localhost:3010 (前端 + API 代理)
- **后端 API**: http://localhost:3000 (直接访问)
- **MySQL**: localhost:3306 (用户名: oneapi, 密码: 123456)
- **Redis**: localhost:6379

## 📁 项目结构

```
one-hub/
├── docker-compose-dev.yml    # 开发环境数据库配置
├── config-dev.yaml          # 开发环境后端配置
├── dev-start.sh            # Linux/macOS 启动脚本
├── dev-start.bat           # Windows 启动脚本
├── DEV-SETUP.md           # 本文档
├── web/                   # 前端代码
│   ├── vite.config.mjs   # Vite 配置 (包含 API 代理)
│   └── package.json      # 前端依赖
├── logs/                 # 开发日志
├── tmp/                  # 临时文件和 PID
└── cache/               # 缓存目录
```

## 🔧 配置说明

### 数据库配置 (docker-compose-dev.yml)
- **MySQL 8.2.0**: 开发专用实例
- **Redis**: 缓存服务
- **数据持久化**: `./data/mysql-dev/`
- **网络隔离**: 独立的 `one-hub-dev` 网络

### 后端配置 (config-dev.yaml)
- **调试模式**: `gin_mode: debug`, `log_level: debug`
- **数据库连接**: 连接本地 Docker MySQL
- **Redis 连接**: 连接本地 Docker Redis
- **开发优化**: 放宽速率限制，启用内存缓存

### 前端配置 (web/vite.config.mjs)
- **开发端口**: 3010
- **API 代理**: `/api` 请求代理到 `http://127.0.0.1:3000`
- **热重载**: 自动刷新和模块热替换

## 🐛 常见问题

### 1. 端口冲突
```bash
# 检查端口占用
lsof -i :3000  # 后端端口
lsof -i :3010  # 前端端口
lsof -i :3306  # MySQL 端口
lsof -i :6379  # Redis 端口

# 停止占用进程
kill -9 <PID>
```

### 2. 数据库连接失败
```bash
# 检查 MySQL 容器状态
docker ps | grep mysql-dev

# 查看 MySQL 日志
docker logs mysql-dev

# 重启数据库服务
docker compose -f docker-compose-dev.yml restart mysql-dev
```

### 3. 前端依赖问题
```bash
cd web
rm -rf node_modules yarn.lock
yarn install
```

### 4. 后端构建失败
```bash
# 清理 Go 模块缓存
go clean -modcache
go mod download
go mod tidy
```

## 📊 开发工作流

### 1. 日常开发
```bash
# 启动开发环境
./dev-start.sh start

# 访问应用
open http://localhost:3010

# 查看日志
tail -f logs/backend-dev.log   # 后端日志
tail -f logs/frontend-dev.log  # 前端日志
```

### 2. 代码修改
- **前端**: 保存后自动热重载
- **后端**: 需要重启后端服务
  ```bash
  # 仅重启后端
  pkill -f one-hub-dev
  go build -o tmp/one-hub-dev main.go
  ./tmp/one-hub-dev --config config-dev.yaml &
  ```

### 3. 数据库操作
```bash
# 连接 MySQL
docker exec -it mysql-dev mysql -u oneapi -p one-api

# 连接 Redis
docker exec -it redis-dev redis-cli
```

## 🔄 与生产环境的差异

| 项目 | 开发环境 | 生产环境 |
|------|----------|----------|
| 前端服务 | Vite 开发服务器 (3010) | 嵌入到 Go 二进制 |
| 后端配置 | debug 模式，详细日志 | release 模式 |
| 数据库 | 本地 Docker MySQL | 外部 MySQL 服务 |
| 访问方式 | localhost:3010 | 统一端口 3000 |
| 跨域处理 | Vite 代理 | 无需处理 |

## 🧪 测试建议

### API 测试
```bash
# 健康检查
curl http://localhost:3000/api/status

# 通过前端代理访问
curl http://localhost:3010/api/status
```

### 前端测试
- 访问 http://localhost:3010
- 检查浏览器开发者工具网络面板
- 确认 API 请求正确代理到后端

## 📝 开发注意事项

1. **配置文件**: 不要提交 `config-dev.yaml` 中的敏感信息
2. **数据持久化**: 开发数据存储在 `./data/mysql-dev/`
3. **日志文件**: 定期清理 `logs/` 目录下的日志文件
4. **缓存目录**: `cache/` 目录可以安全删除重建
5. **跨域问题**: 开发环境通过 Vite 代理解决，生产环境无此问题

## 🆘 获取帮助

如果遇到问题：
1. 检查 `logs/` 目录下的日志文件
2. 运行 `./dev-start.sh status` 查看服务状态
3. 查看 Docker 容器日志: `docker logs <container_name>`
4. 重启开发环境: `./dev-start.sh restart`
