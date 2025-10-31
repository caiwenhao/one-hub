# Project Context

## Purpose
One Hub 是一个基于 one-api 二次开发的 AI 能力聚合与中转平台，目标是：
- 统一对接多家大模型/多媒体 AI 供应商，提供 OpenAI/Gemini/Claude 等兼容 API 网关与代理能力
- 提供完善的运营后台（用户仪表盘、管理员分析面板）、计费与套餐、模型价格管理与自动同步、日志与监控
- 提供前后端一体的交付形态（Go 后端嵌入打包 Web 前端），便于单体部署与运维
- 在保持与上游协议兼容的同时，支持更丰富的供应商能力与平台化特性（支付、分组限速、告警通知、状态页等）

## Tech Stack
- Backend
  - Go 1.24（Gin、GORM、Viper、Zap、Sessions、Gzip、embed）
  - 数据库：MySQL（默认），兼容 PostgreSQL、SQLite（见 GORM 驱动）
  - 缓存：Redis，用于会话/限流/缓存等
  - 文档与指标：Swaggo/Swagger、Prometheus 指标暴露（BasicAuth 保护）
  - 其它：Telegram Bot、OIDC/GitHub/WeChat 登录、支付网关（Epay/Stripe/Alipay/WeChat）
- Frontend
  - React 18 + Vite 5，MUI 组件库，Redux，i18next，多语言
  - 构建与工具：ESLint（react-app preset）、Prettier、Playwright（可用于 e2e）
- Dev/Infra
  - Docker Compose（MySQL/Redis 开发编排）、Taskfile 任务、Yarn 1 包管理
  - Go embed 打包前端静态资源到单一二进制

## Project Conventions

### Code Style
- Go
  - 统一使用 `gofmt`/`goimports` 格式化，`golangci-lint` 保障基础静态检查
  - 已开启的 linters：goimports、gofmt、govet、misspell、ineffassign、typecheck、whitespace、gocyclo、revive、unused（见 `.golangci.yml`）
  - 日志：统一使用 `zap`，输出到控制台与滚动文件
- Frontend
  - 2 空格缩进，`.editorconfig` 统一换行与尾行规则
  - ESLint + Prettier 统一格式与校验，React 函数组件优先
- 命名
  - Go 采用惯用驼峰与包级职责划分（controller/middleware/model/provider 等）
  - 前端以 MUI 设计体系与路由视图划分模块

### Architecture Patterns
- 单体服务：Go + Gin 提供 API 与静态资源服务，`embed` 内嵌前端产物
- 分层与职责
  - `router/` 路由聚合（`/api` 管理与业务接口、`/v1` 等 OpenAI 兼容转发、`/swagger` 文档、`/api/metrics` 指标）
  - `controller/` 业务控制器（用户、支付、认证、通知、文件等）
  - `middleware/` 认证、限流、日志、CORS、会话等横切能力
  - `model/` GORM 实体与仓储、迁移、价格/账单/分组等领域模型
  - `relay/` OpenAI/Gemini/Claude 等协议兼容层与供应商分发
  - `providers/` 各供应商适配实现
- 配置优先：`viper` 读取 YAML/环境变量，支持本地与容器化部署
- 指标与文档：`/api/metrics` 暴露 Prometheus 指标（BasicAuth）；`/swagger` 在 debug 或显式开启时提供 API 文档

### Testing Strategy
- Go 单元测试：覆盖存储上传、通知通道、图片处理、任务参数校验、价格换算等模块（`go test ./...`）
- 前端：仓库包含若干 `.test.js` 单测与 Playwright 依赖，建议引入 Vitest/Jest 运行单测，E2E 可用 Playwright 配置
- 建议在 PR 中至少补齐核心路径的单测或手测截图（见 PR 模板要求）

### Git Workflow
- 分支：建议使用 `feature/*`、`fix/*`、`chore/*`、`docs/*` 等前缀进行主题分支开发
- 提交流程：通过 PR 合入默认分支；PR 需自测并附带截图与说明（见 `pull_request_template.md`）
- 提交信息：未强制启用 commitlint，推荐遵循 Conventional Commits（feat/fix/docs/chore/refactor/test）
- 质量门禁：通过基础 Lint 与构建；后端建议 `golangci-lint run`，前端 `eslint`/`prettier`

## Domain Context
- 平台定位：面向多模型供应商聚合的 API 网关/中转与运营平台，兼容 OpenAI 风格接口，并扩展 Gemini/Claude/图像/音频/视频等多模态
- 核心域模型：用户、用户组（限速/RPM）、通道（供应商接入）、模型价格（自动同步/倍率/币种）、用量/账单、订单/支付
- 鉴权与会话：支持 Token、会话 Cookie、OIDC/GitHub/WeChat/WebAuthn 等多种登录方式
- 监控与可观测：Prometheus 指标、请求日志（含耗时）、可选 Uptime Kuma 状态面板读取
- 前端：管理后台 + 用户仪表盘，支持多语言与主题，API 代理到后端

## Important Constraints
- 与上游 one-api 的数据库不兼容，请勿与原版混用同一数据库实例
- 法规合规：遵循各供应商使用条款与当地法规（README 中对中国大陆生成式 AI 服务合规的提醒）；禁止非法用途
- 安全：指标端点需 BasicAuth 防护；开放的管理接口必须鉴权；敏感速率接口采用更严格限流
- 运行时要求：Go 1.24、Node.js 16+、Yarn 1；开发环境依赖 Docker（MySQL/Redis）
- 配置：优先使用 `config-*.yaml` 或环境变量；部分功能需显式开启（如 Swagger 文档）

## External Dependencies
- 数据与中间件：MySQL、Redis、（可选）对象存储/图床（Aliyun OSS/SMMS/Imgur）
- 指标与监控：Prometheus（/api/metrics）、Uptime Kuma（读取状态页）
- 身份与登录：OIDC、GitHub OAuth、WeChat、WebAuthn（Passkey）
- 支付：Epay、Stripe、Alipay、WeChat
- 供应商 API：OpenAI、Azure OpenAI/Speech、Anthropic、Gemini、百度文心、通义千问、以及 README 中列举的其它厂商
- 其他：Telegram Bot Webhook
