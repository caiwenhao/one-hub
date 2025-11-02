# Project Context

## Purpose
One Hub 是一个基于 one-api 二次开发的 AI 能力聚合与中转平台，目标是：
- 统一对接多家大模型/多媒体 AI 供应商，提供 OpenAI/Gemini/Claude 等兼容 API 网关与代理能力
- 提供完善的运营后台（用户仪表盘、管理员分析面板）、计费与套餐、模型价格管理与自动同步、日志与监控
- 提供前后端一体的交付形态（Go 后端嵌入打包 Web 前端），便于单体部署与运维
- 在保持与上游协议兼容的同时，支持更丰富的供应商能力与平台化特性（支付、分组限速、告警通知、状态页等）

## Tech Stack
- Backend
  - Go 1.24（工具链 1.24.2）
  - Web 框架：Gin（路由、中间件、会话管理）
  - ORM：GORM v1.30.0（支持 MySQL、PostgreSQL、SQLite）
  - 配置管理：Viper v1.20.1（YAML/环境变量）
  - 日志：Zap v1.27.0 + Lumberjack v2.2.1（滚动日志）
  - 缓存：Redis v9.10.0 + FreeCache（内存缓存）+ eko/gocache v4（缓存抽象层）
  - 认证与安全：
    - JWT（golang-jwt/jwt v5.2.3）
    - OIDC（coreos/go-oidc v3.14.1）
    - WebAuthn（go-webauthn/webauthn v0.13.4）
    - OAuth2（golang.org/x/oauth2）
  - 支付集成：
    - Stripe（stripe-go v80.2.1）
    - 支付宝（smartwalle/alipay v3.2.25）
    - 微信支付（wechatpay-apiv3/wechatpay-go v0.2.20）
  - AI 供应商 SDK：
    - Google AI（google.golang.org/api）
    - AWS SDK（aws-sdk-go v1.55.7）
    - 阿里云 OSS（aliyun-oss-go-sdk v3.0.2）
  - 任务调度：go-co-op/gocron v2.16.2
  - 监控与指标：Prometheus client_golang v1.22.0
  - 通知：
    - Telegram Bot（PaulSonOfLars/gotgbot v2.0.0-rc.32）
    - 邮件（wneessen/go-mail v0.6.2）
  - 其他工具：
    - Snowflake ID（bwmarrin/snowflake v0.3.0）
    - UUID（google/uuid v1.6.0）
    - Decimal（shopspring/decimal v1.4.0）
    - Tiktoken（pkoukk/tiktoken-go v0.1.7）
    - WebSocket（gorilla/websocket v1.5.3）
    - Markdown（gomarkdown/markdown）
  - API 文档：Swaggo/Swag v1.16.3（生成 Swagger 文档）
  - 静态资源：Go embed（将前端打包进二进制）
  
- Frontend
  - 核心框架：React 18.2.0 + React DOM 18.2.0
  - 构建工具：Vite 5.4.18 + @vitejs/plugin-react 4.3.1
  - UI 组件库：
    - Material-UI v5.15.7（@mui/material、@mui/icons-material、@mui/lab）
    - MUI X Data Grid v6.19.4
    - MUI X Date Pickers v6.18.5
    - Tabler Icons v2.46.0
    - Iconify React v5.0.2
  - 状态管理：Redux v5.0.1 + React-Redux v9.1.0
  - 路由：React Router v6.21.3 + React Router DOM v6.21.3
  - 国际化：i18next v23.11.5 + react-i18next v14.1.2 + i18next-browser-languagedetector v8.0.0
  - 表单管理：Formik v2.4.5 + Yup v0.32.11
  - 数据可视化：ApexCharts 3.35.3 + react-apexcharts 1.4.0
  - 动画：Framer Motion v11.0.3
  - 通知：Notistack v3.0.1
  - HTTP 客户端：Axios v0.30.0
  - 其他 UI 组件：
    - react-perfect-scrollbar v1.5.8
    - react-qrcode-logo v3.0.0
    - react-turnstile v1.1.2（Cloudflare Turnstile）
    - material-ui-popup-state v5.0.10
    - country-flag-icons v1.5.19
  - 工具库：
    - dayjs v1.11.10（日期处理）
    - decimal.js v10.4.3（精确计算）
    - marked v4.1.1（Markdown 渲染）
    - highlight.js v11.10.0（代码高亮）
  - 代码质量：
    - ESLint v8.56.0（eslint-config-react-app、eslint-plugin-prettier）
    - Prettier v3.2.4
    - Playwright v1.56.1（E2E 测试）
  - 样式：Sass v1.70.0 + Emotion（@emotion/react、@emotion/styled）
  - 设计系统：Style Dictionary v5.1.1
  
- Dev/Infra
  - 包管理：Yarn 1.22.22（前端）、Go Modules（后端）
  - 任务运行：Taskfile（Go Task）
  - 容器化：Docker + Docker Compose（开发环境 MySQL/Redis 编排）
  - 热重载：Air（.air.toml 配置）
  - 代码检查：golangci-lint（.golangci.yml 配置）
  - 编辑器配置：.editorconfig（统一代码风格）
  - 版本控制：Git（.gitignore、pull_request_template.md）
  - 构建产物：单一二进制文件（Go embed 嵌入前端静态资源）

## Project Conventions

### Code Style
- Go
  - 格式化：统一使用 `gofmt`/`goimports`，自动格式化代码
  - 静态检查：`golangci-lint`（配置文件：`.golangci.yml`）
  - 已启用的 linters：
    - goimports（导入排序）
    - gofmt（代码格式）
    - govet（静态分析）
    - misspell（拼写检查）
    - ineffassign（无效赋值检查）
    - typecheck（类型检查）
    - whitespace（空白符检查）
    - gocyclo（圈复杂度检查）
    - revive（代码风格检查）
    - unused（未使用代码检查）
  - 日志规范：
    - 统一使用 `zap` 结构化日志
    - 日志级别：Debug、Info、Warn、Error、Fatal
    - 输出目标：控制台 + 滚动文件（lumberjack）
    - 日志文件：`logs/one-hub.log`（开发环境：`logs/one-hub-dev.log`）
  - 错误处理：使用 `error` 返回值，避免 panic（除非不可恢复错误）
  - 注释：导出的函数、类型、常量必须有文档注释
  
- Frontend
  - 缩进：2 空格（`.editorconfig` 统一配置）
  - 换行符：LF（Unix 风格）
  - 文件编码：UTF-8
  - 代码格式化：Prettier v3.2.4（配置文件：`.prettierrc`）
  - 代码检查：ESLint v8.56.0（配置：`eslintConfig` in package.json）
    - 继承：eslint-config-react-app
    - 插件：prettier、react、react-hooks、jsx-a11y、import
  - React 规范：
    - 优先使用函数组件（Hooks）
    - 避免使用类组件（除非必要）
    - Props 验证使用 PropTypes
    - 状态管理优先使用 Redux
  - 文件命名：
    - 组件文件：PascalCase（如 `UserProfile.jsx`）
    - 工具文件：camelCase（如 `formatDate.js`）
    - 常量文件：UPPER_SNAKE_CASE（如 `API_CONSTANTS.js`）
  
- 命名约定
  - Go：
    - 包名：小写单词，简短（如 `model`、`controller`、`middleware`）
    - 导出标识符：PascalCase（如 `UserModel`、`GetUser`）
    - 私有标识符：camelCase（如 `userCache`、`validateToken`）
    - 常量：PascalCase 或 UPPER_SNAKE_CASE（如 `MaxRetries`、`API_VERSION`）
    - 接口：通常以 `-er` 结尾（如 `Reader`、`Writer`）
  - Frontend：
    - 组件：PascalCase（如 `UserDashboard`、`PaymentForm`）
    - Hooks：以 `use` 开头（如 `useAuth`、`useNotification`）
    - 常量：UPPER_SNAKE_CASE（如 `API_BASE_URL`、`MAX_FILE_SIZE`）
    - 函数：camelCase（如 `fetchUserData`、`handleSubmit`）
    - 文件夹：kebab-case（如 `user-profile`、`payment-gateway`）

### Architecture Patterns
- 架构风格：单体应用（Monolithic）
  - Go 后端 + Gin 框架提供 RESTful API
  - 前端静态资源通过 Go embed 嵌入二进制文件
  - 单一进程部署，简化运维
  
- 后端分层架构（按目录职责）：
  - `main.go`：应用入口，初始化各模块（数据库、缓存、定时任务、HTTP 服务器）
  - `router/`：路由聚合与注册
    - `/api/*`：管理后台与业务 API（用户、渠道、订单、日志等）
    - `/v1/*`：OpenAI 兼容 API（chat/completions、embeddings、images 等）
    - `/gemini/*`：Gemini 格式 API
    - `/claude/*`：Claude 格式 API
    - `/swagger/*`：Swagger API 文档（debug 模式或显式开启）
    - `/api/metrics`：Prometheus 指标端点（BasicAuth 保护）
  - `controller/`：业务控制器（处理 HTTP 请求，调用 model 层）
    - 用户管理（user.go）
    - 渠道管理（channel.go、channel-test.go、channel-billing.go）
    - 订单与支付（order.go、payment.go、billing.go）
    - 认证（github.go、oidc.go、wechat.go、lark.go、telegram.go、webauthn.go）
    - 日志与分析（log.go、analytics.go、system_log.go）
    - 其他（token.go、redemption.go、pricing.go、midjourney.go、task.go）
  - `middleware/`：中间件（横切关注点）
    - 认证与授权（auth.go）
    - 限流（rate-limit.go、api-limit.go）
    - 日志记录（logger.go）
    - CORS（cors.go）
    - 会话管理（context-userid.go）
    - 请求分发（distributor.go）
    - 价格计算（prices.go）
    - 监控指标（metrics.go）
    - 错误恢复（recover.go）
    - 其他（request-id.go、cache.go、turnstile-check.go、telegram.go）
  - `model/`：数据模型与数据访问层
    - GORM 实体定义（user.go、channel.go、token.go、order.go、log.go 等）
    - 数据库迁移（migrate.go）
    - 缓存管理（cache.go）
    - 业务逻辑（pricing.go、statistics.go、balancer.go）
  - `relay/`：AI 供应商协议转换与请求中转
    - OpenAI/Gemini/Claude 等协议适配
    - 请求转发与响应处理
    - 流式响应支持
  - `providers/`：各 AI 供应商具体实现
    - 每个供应商一个子目录（如 `openai/`、`azure/`、`gemini/`、`claude/` 等）
    - 实现统一接口，处理供应商特定逻辑
  - `common/`：公共工具与基础设施
    - 配置管理（config/）
    - 缓存抽象（cache/）
    - 数据库连接（database/）
    - Redis 客户端（redis/）
    - 日志工具（logger/）
    - 通知服务（notify/）
    - 存储服务（storage/）
    - 搜索服务（search/）
    - 限流器（limit/）
    - 调度器（scheduler/）
    - OIDC（oidc/）
    - WebAuthn（webauthn/）
    - Telegram（telegram/）
    - 其他工具（utils/、requester/、image/）
  - `cron/`：定时任务（如数据同步、自动测试渠道）
  - `metrics/`：监控指标收集
  - `payment/`：支付网关集成
    - `gateway/`：各支付渠道实现
    - `types/`：支付相关类型定义
  - `safty/`：安全检查工具
  - `cli/`：命令行工具（导出数据、标志解析）
  - `mcp/`：MCP（Model Context Protocol）服务器实现
    - `tools/`：MCP 工具定义
  
- 前端架构（React SPA）：
  - `web/src/`：源代码目录
    - `App.jsx`：应用根组件
    - `index.jsx`：应用入口
    - `routes/`：路由配置（React Router）
    - `views/`：页面组件（按功能模块划分）
    - `layout/`：布局组件（MainLayout、MinimalLayout）
    - `ui-component/`：可复用 UI 组件
    - `store/`：Redux 状态管理
    - `contexts/`：React Context（主题、配置等）
    - `hooks/`：自定义 Hooks
    - `utils/`：工具函数
    - `constants/`：常量定义
    - `i18n/`：国际化配置
    - `locales/`：翻译文件
    - `themes/`：主题配置（MUI）
    - `menu-items/`：菜单配置
    - `assets/`：静态资源（图片、字体等）
  - `web/build/`：构建产物（Vite 打包后的静态文件）
  - `web/public/`：公共静态资源
  
- 配置管理：
  - 配置文件：`config-*.yaml`（如 `config-local.yaml`、`config-dev.yaml`）
  - 环境变量：支持通过环境变量覆盖配置
  - Viper 加载顺序：配置文件 → 环境变量
  - 配置项包括：
    - 数据库连接（MySQL/PostgreSQL/SQLite）
    - Redis 连接
    - 服务器端口
    - 日志级别
    - 会话密钥
    - 支付网关配置
    - OIDC/OAuth 配置
    - Telegram Bot Token
    - 等等
  
- 数据流：
  1. 用户请求 → Gin 路由
  2. 中间件处理（认证、限流、日志）
  3. Controller 处理业务逻辑
  4. Model 层访问数据库/缓存
  5. Relay/Provider 转发到 AI 供应商（如果是 AI 请求）
  6. 返回响应
  
- 缓存策略：
  - Redis：会话、限流计数、热点数据
  - FreeCache：内存缓存（可选，通过配置启用）
  - 数据库查询结果缓存（通过 eko/gocache 抽象层）
  
- 监控与可观测性：
  - Prometheus 指标：`/api/metrics` 端点（需 BasicAuth）
  - 结构化日志：Zap 输出到文件和控制台
  - 请求日志：记录每个请求的耗时、状态码、错误信息
  - Uptime Kuma 集成：可选的状态页监控
  
- API 文档：
  - Swagger/OpenAPI：通过 Swaggo 自动生成
  - 访问路径：`/swagger/index.html`
  - 仅在 debug 模式或显式配置时启用

### Testing Strategy
- 后端测试（Go）：
  - 单元测试：使用 Go 标准库 `testing` 包
  - 测试覆盖模块：
    - 存储上传（common/storage）
    - 通知通道（common/notify）
    - 图片处理（common/image）
    - 任务参数校验（model/task）
    - 价格换算（model/price_vidu_test.go）
  - 运行测试：`go test ./...`
  - Mock 工具：golang/mock（如需要）
  - 测试数据库：建议使用 SQLite 内存数据库或 Docker 容器
  
- 前端测试：
  - 单元测试：
    - 框架：建议使用 Vitest 或 Jest
    - 现有测试文件：部分 `.test.js` 文件
    - 测试内容：组件渲染、工具函数、状态管理
  - E2E 测试：
    - 工具：Playwright v1.56.1（已安装）
    - 配置：需要创建 `playwright.config.js`
    - 测试场景：关键用户流程（登录、创建渠道、发起请求等）
  - 视觉回归测试：
    - 脚本：`scripts/baseline-screenshot.mjs`（基线截图）
    - 对比检查：`scripts/contrast-check.mjs`（对比度检查）
  
- 集成测试：
  - 渠道测试：`controller/channel-test.go`（自动测试 AI 供应商连通性）
  - API 测试：可使用 Postman/Insomnia 或编写 Go 集成测试
  
- 测试环境：
  - 开发环境：`docker-compose-dev.yml`（MySQL + Redis）
  - 配置文件：`config-dev.yaml`、`config-local.yaml`
  - 启动脚本：`dev-start.sh`（Linux/macOS）、`dev-start.bat`（Windows）
  
- 质量保证：
  - PR 要求：
    - 至少补齐核心路径的单元测试
    - 提供手动测试截图（见 `pull_request_template.md`）
    - 通过 golangci-lint 检查
    - 通过 ESLint 检查
  - CI/CD：建议配置 GitHub Actions 自动运行测试
  
- 测试数据：
  - SQL 脚本：`init_kling_prices.sql`（初始化价格数据）
  - 测试脚本：`test_kling_api.sh`、`test-price-converter.js`
  - Demo 脚本：`kapon_chat_demo.py`（Python 客户端示例）

### Git Workflow
- 分支策略：
  - 主分支：`main` 或 `master`（生产代码）
  - 开发分支：建议使用语义化前缀
    - `feature/*`：新功能（如 `feature/add-payment-gateway`）
    - `fix/*`：Bug 修复（如 `fix/login-error`）
    - `chore/*`：杂项任务（如 `chore/update-dependencies`）
    - `docs/*`：文档更新（如 `docs/update-readme`）
    - `refactor/*`：代码重构（如 `refactor/optimize-cache`）
    - `test/*`：测试相关（如 `test/add-unit-tests`）
  
- 提交流程：
  1. 从主分支创建功能分支
  2. 本地开发与测试
  3. 提交代码（遵循提交信息规范）
  4. 推送到远程仓库
  5. 创建 Pull Request
  6. 代码审查
  7. 合并到主分支
  
- Pull Request 要求（见 `pull_request_template.md`）：
  - 描述变更内容与原因
  - 附带测试截图或测试结果
  - 确保通过所有检查（Lint、构建、测试）
  - 关联相关 Issue（如有）
  
- 提交信息规范：
  - 推荐遵循 Conventional Commits
  - 格式：`<type>(<scope>): <subject>`
  - 类型（type）：
    - `feat`：新功能
    - `fix`：Bug 修复
    - `docs`：文档更新
    - `style`：代码格式（不影响功能）
    - `refactor`：重构
    - `perf`：性能优化
    - `test`：测试相关
    - `chore`：构建/工具/依赖更新
  - 示例：
    - `feat(payment): add Stripe payment gateway`
    - `fix(auth): resolve JWT token expiration issue`
    - `docs(readme): update installation instructions`
  
- 质量门禁：
  - 后端：
    - 运行 `golangci-lint run` 确保代码质量
    - 运行 `go test ./...` 确保测试通过
    - 运行 `go build` 确保编译成功
  - 前端：
    - 运行 `yarn lint` 或 `npm run lint` 检查代码风格
    - 运行 `yarn prettier` 或 `npm run prettier` 格式化代码
    - 运行 `yarn build` 或 `npm run build` 确保构建成功
  
- 版本管理：
  - 版本号文件：`VERSION`（根目录）、`web/VERSION`（前端）
  - 版本号格式：语义化版本（Semantic Versioning）`MAJOR.MINOR.PATCH`
  - 发布流程：
    1. 更新 `VERSION` 文件
    2. 更新 `CHANGELOG`（如有）
    3. 创建 Git Tag（如 `v1.2.0`）
    4. 推送 Tag 到远程仓库
    5. 创建 GitHub Release
  
- 忽略文件：
  - `.gitignore`：忽略构建产物、依赖、日志、缓存等
  - `.cursorignore`：Cursor 编辑器忽略文件
  
- 代码审查：
  - 关注点：
    - 代码质量与可读性
    - 安全性（敏感信息、SQL 注入、XSS 等）
    - 性能（N+1 查询、内存泄漏等）
    - 测试覆盖率
    - 文档完整性

## Domain Context

### 平台定位
One Hub 是一个 AI 能力聚合与中转平台，主要功能包括：
- API 网关：统一对接多家 AI 供应商，提供标准化 API 接口
- 协议兼容：支持 OpenAI、Gemini、Claude 等主流 API 格式
- 多模态支持：文本生成、图像生成、语音合成、语音识别、Embeddings 等
- 运营平台：用户管理、渠道管理、计费系统、订单支付、数据分析
- 中转代理：请求转发、负载均衡、失败重试、流式响应

### 核心领域模型

1. **用户（User）**
   - 用户信息：用户名、邮箱、密码、角色（管理员/普通用户）
   - 额度管理：余额、已用额度、配额
   - 认证方式：密码、GitHub OAuth、OIDC、微信、WebAuthn
   - 用户组：归属用户组，继承组级别的限流配置

2. **用户组（User Group）**
   - 分组管理：将用户分组，统一管理权限和限流
   - RPM 限制：每分钟请求数限制（Rate Per Minute）
   - 自动升级：根据消费金额自动升级用户组

3. **渠道（Channel）**
   - 供应商接入：配置 AI 供应商的 API Key、Base URL、代理等
   - 模型映射：将供应商模型映射到标准模型名称
   - 优先级与权重：用于负载均衡和故障转移
   - 状态管理：启用/禁用、自动测试、健康检查
   - 计费配置：单独的价格倍率、按次收费

4. **模型（Model）**
   - 模型信息：模型名称、所属供应商、模型类型（文本/图像/音频等）
   - 价格管理：输入价格、输出价格、币种、倍率
   - 自动同步：从供应商自动获取最新模型列表和价格
   - 模型通配符：支持模糊匹配（如 `gpt-4*`）

5. **Token**
   - API Token：用户调用 API 的凭证
   - 权限控制：可访问的模型、额度限制、过期时间
   - 使用统计：请求次数、消费金额

6. **日志（Log）**
   - 请求日志：记录每次 API 调用的详细信息
   - 字段：用户、模型、Token 数、耗时、状态码、错误信息
   - 用途：计费、审计、分析、调试

7. **订单与支付（Order & Payment）**
   - 充值订单：用户充值记录
   - 支付方式：支付宝、微信支付、Stripe
   - 订单状态：待支付、已支付、已取消、已退款
   - 支付回调：处理支付网关的异步通知

8. **账单（Billing）**
   - 用量统计：按用户、渠道、模型统计消费
   - 月度账单：自动生成用户月度账单（可选功能）
   - 分析报表：管理员查看平台整体数据

9. **兑换码（Redemption）**
   - 兑换码生成：批量生成兑换码
   - 额度充值：用户使用兑换码充值额度
   - 使用记录：兑换码使用状态和历史

10. **任务（Task）**
    - 异步任务：Midjourney、Suno 等需要异步处理的任务
    - 任务状态：排队、处理中、完成、失败
    - 结果存储：任务结果（图片、音频等）

11. **价格（Pricing）**
    - 价格配置：每个模型的输入/输出价格
    - 自动更新：定期从供应商同步最新价格
    - 更新策略：覆盖更新、只更新现有、只新增

12. **通知（Notification）**
    - 通知渠道：Telegram Bot、邮件
    - 通知场景：额度不足、渠道异常、系统告警

### 业务流程

1. **用户注册与登录**
   - 注册方式：邮箱注册、GitHub OAuth、OIDC、微信登录
   - 认证方式：密码、WebAuthn（Passkey）
   - 会话管理：基于 Cookie 的会话（30 天有效期）

2. **API 调用流程**
   - 用户发起请求 → 认证（Token 验证）→ 限流检查 → 模型路由 → 渠道选择 → 转发到供应商 → 返回响应 → 记录日志 → 扣费

3. **渠道选择策略**
   - 负载均衡：根据权重分配请求
   - 故障转移：失败后自动切换到其他渠道
   - 优先级：优先使用高优先级渠道

4. **计费流程**
   - 请求前：检查用户余额是否足够
   - 请求后：根据 Token 数计算费用，扣除余额
   - 日志记录：记录消费明细

5. **充值流程**
   - 用户创建充值订单 → 选择支付方式 → 跳转支付网关 → 支付成功 → 回调通知 → 增加余额

6. **渠道测试**
   - 自动测试：定期自动测试所有渠道的可用性
   - 手动测试：管理员手动触发测试
   - 测试结果：更新渠道状态（正常/异常）

### 权限模型
- 角色：
  - 管理员（Admin）：完全权限，可管理所有资源
  - 普通用户（User）：只能管理自己的资源（Token、订单等）
- 权限控制：
  - 基于角色的访问控制（RBAC）
  - 中间件层面进行权限检查
  - API Token 级别的权限控制（可访问的模型）

### 监控与可观测性
- Prometheus 指标：
  - 请求数、错误率、响应时间
  - 渠道状态、用户数、订单数
  - 访问端点：`/api/metrics`（需 BasicAuth）
- 日志：
  - 结构化日志（Zap）
  - 请求日志：记录每个请求的详细信息（耗时、状态码、错误）
  - 系统日志：记录系统事件（启动、关闭、错误）
- Uptime Kuma：
  - 可选集成，提供状态页监控
  - 通过环境变量或配置文件启用

### 前端功能
- 用户仪表盘：
  - 查看余额、使用统计
  - 管理 API Token
  - 充值、查看订单
  - 查看请求日志
- 管理后台：
  - 用户管理：查看、编辑、删除用户
  - 渠道管理：添加、编辑、测试渠道
  - 模型管理：配置模型价格、更新模型列表
  - 订单管理：查看、处理订单
  - 日志查看：查看系统日志、请求日志
  - 数据分析：查看平台统计数据（用户数、消费、渠道状态等）
- 多语言支持：中文、英文（通过 i18next）
- 主题支持：亮色、暗色主题

## Important Constraints

### 兼容性约束
- **数据库不兼容**：与上游 one-api 的数据库结构不兼容，请勿混用同一数据库实例
- **版本要求**：
  - Go：1.24+（工具链 1.24.2）
  - Node.js：16+（推荐 18 或 20）
  - Yarn：1.22+
  - MySQL：5.7+ 或 8.0+（推荐）
  - PostgreSQL：12+（可选）
  - Redis：6.0+（推荐 7.0+）

### 法规与合规
- **使用条款**：必须遵循各 AI 供应商的使用条款和服务协议
- **法律法规**：遵守当地法律法规，禁止用于非法用途
- **中国大陆特别提醒**：
  - 根据《生成式人工智能服务管理暂行办法》
  - 请勿对中国地区公众提供一切未经备案的生成式人工智能服务
  - 仅供个人学习使用，不保证稳定性，不提供技术支持
- **免责声明**：使用者需自行承担使用风险

### 安全约束
- **认证与授权**：
  - 所有管理接口必须进行身份认证
  - API Token 必须验证有效性和权限
  - 敏感操作需要管理员权限
- **指标端点保护**：
  - `/api/metrics` 端点必须使用 BasicAuth 保护
  - 用户名和密码通过配置文件或环境变量设置
- **限流保护**：
  - 用户级别限流：防止单个用户滥用
  - 用户组级别限流：RPM（每分钟请求数）限制
  - IP 级别限流：防止恶意攻击
  - 敏感接口采用更严格的限流策略
- **数据安全**：
  - 密码使用 bcrypt 加密存储
  - API Key 加密存储
  - 敏感配置（如支付密钥）不得提交到版本控制
  - 日志中不得记录敏感信息（密码、完整 API Key 等）
- **CORS 配置**：
  - 生产环境需配置允许的域名
  - 避免使用 `*` 通配符
- **会话安全**：
  - 使用 HttpOnly Cookie
  - 设置合理的会话过期时间（默认 30 天）
  - 生产环境建议启用 Secure 标志（HTTPS）

### 配置约束
- **配置优先级**：环境变量 > 配置文件 > 默认值
- **配置文件**：
  - 开发环境：`config-dev.yaml`、`config-local.yaml`
  - 生产环境：`config.yaml` 或通过环境变量配置
  - 示例配置：`config.example.yaml`
- **必需配置**：
  - 数据库连接信息
  - Redis 连接信息
  - 会话密钥（`SESSION_SECRET`）
- **可选配置**：
  - Swagger 文档（默认关闭，需显式启用）
  - 内存缓存（默认关闭）
  - 用户月度账单（默认关闭）
  - 模型价格自动更新（默认关闭）
  - Uptime Kuma 集成（默认关闭）
  - Telegram Bot（需配置 Token）
  - 支付网关（需配置各支付渠道的密钥）

### 运行时约束
- **开发环境**：
  - 依赖 Docker 和 Docker Compose（用于 MySQL 和 Redis）
  - 启动脚本：`dev-start.sh`（Linux/macOS）或 `dev-start.bat`（Windows）
  - 热重载：使用 Air（配置文件：`.air.toml`）
- **生产环境**：
  - 单一二进制文件部署（Go embed 嵌入前端）
  - 需要外部 MySQL/PostgreSQL 和 Redis
  - 建议使用 Docker 容器部署（`docker-compose.yml`）
  - 建议使用反向代理（Nginx/Caddy）处理 HTTPS
- **资源要求**：
  - 最小配置：1 CPU、1GB RAM
  - 推荐配置：2 CPU、2GB RAM
  - 存储：取决于日志和数据量，建议至少 10GB

### 性能约束
- **并发限制**：
  - 单个用户的并发请求数受限于用户组 RPM 配置
  - 渠道并发请求数受限于供应商限制
- **缓存策略**：
  - 启用内存缓存可提升性能，但会增加内存占用
  - Redis 缓存用于会话和限流，必须启用
- **数据库连接池**：
  - 根据并发量调整连接池大小
  - 避免连接泄漏
- **日志轮转**：
  - 使用 Lumberjack 自动轮转日志文件
  - 避免日志文件过大影响性能

### 功能约束
- **模型列表更新**：
  - 程序自带的模型列表不再频繁更新（除新增供应商外）
  - 如发现缺少新模型，请在后台手动更新：`后台 → 模型价格 → 更新价格`
- **供应商支持**：
  - 支持的供应商列表见 README.md
  - 部分供应商功能有限制（如 Bedrock 仅支持 Anthropic 模型）
- **协议兼容性**：
  - OpenAI 格式：完全兼容
  - Gemini 格式：支持（需使用 `/gemini/*` 路径）
  - Claude 格式：支持（需使用 `/claude/*` 路径）
- **流式响应**：
  - 支持 Server-Sent Events (SSE)
  - 部分供应商可能不支持流式响应

### 部署约束
- **单体部署**：
  - 当前架构为单体应用，不支持水平扩展
  - 如需高可用，建议使用负载均衡 + 多实例部署
  - 注意：多实例部署时需共享 Redis 和数据库
- **数据库迁移**：
  - 启动时自动执行数据库迁移（GORM AutoMigrate）
  - 建议在升级前备份数据库
- **静态资源**：
  - 前端静态资源嵌入到二进制文件中
  - 如需自定义前端，需重新构建

## Development Workflow

### 本地开发环境搭建
1. **安装依赖**：
   - Go 1.24+
   - Node.js 16+ 和 Yarn 1.22+
   - Docker 和 Docker Compose
   - Air（可选，用于热重载）：`go install github.com/cosmtrek/air@latest`
   - golangci-lint（可选，用于代码检查）

2. **启动数据库和 Redis**：
   ```bash
   docker-compose -f docker-compose-dev.yml up -d
   ```

3. **配置文件**：
   - 复制 `config.example.yaml` 为 `config-local.yaml`
   - 修改数据库和 Redis 连接信息

4. **安装前端依赖**：
   ```bash
   cd web
   yarn install
   ```

5. **启动开发服务器**：
   - 使用脚本：`./dev-start.sh`（Linux/macOS）或 `dev-start.bat`（Windows）
   - 或手动启动：
     ```bash
     # 后端（使用 Air 热重载）
     air
     
     # 前端（另一个终端）
     cd web
     yarn dev
     ```

6. **访问应用**：
   - 前端开发服务器：http://localhost:5173
   - 后端 API：http://localhost:3000
   - Swagger 文档：http://localhost:3000/swagger/index.html（需在配置中启用）

### 常用开发命令
- **后端**：
  ```bash
  # 运行测试
  go test ./...
  
  # 代码检查
  golangci-lint run
  
  # 构建
  go build -o one-api
  
  # 生成 Swagger 文档
  go generate
  
  # 使用 Taskfile
  task build    # 构建
  task test     # 测试
  task lint     # 检查
  ```

- **前端**：
  ```bash
  cd web
  
  # 开发服务器
  yarn dev
  
  # 构建
  yarn build
  
  # 代码检查
  yarn lint
  
  # 代码格式化
  yarn prettier
  
  # 国际化
  yarn i18n
  ```

### 调试技巧
- **后端调试**：
  - 使用 VS Code 的 Go 调试器
  - 配置 `.vscode/launch.json`
  - 设置断点，启动调试
- **前端调试**：
  - 使用浏览器开发者工具
  - React DevTools 扩展
  - Redux DevTools 扩展
- **日志调试**：
  - 设置日志级别为 `debug`（配置文件或环境变量）
  - 查看日志文件：`logs/one-hub-dev.log`

### 数据库管理
- **迁移**：
  - 自动迁移：启动时自动执行（GORM AutoMigrate）
  - 手动迁移：修改 `model/migrate.go`
- **备份**：
  ```bash
  # MySQL
  mysqldump -u root -p one_hub > backup.sql
  
  # 恢复
  mysql -u root -p one_hub < backup.sql
  ```
- **重置数据库**：
  ```bash
  # 删除数据库
  docker-compose -f docker-compose-dev.yml down -v
  
  # 重新创建
  docker-compose -f docker-compose-dev.yml up -d
  ```

### 常见问题排查
- **端口冲突**：
  - 修改配置文件中的 `port` 配置
  - 或使用环境变量 `PORT`
- **数据库连接失败**：
  - 检查 Docker 容器是否运行：`docker ps`
  - 检查配置文件中的连接信息
- **前端代理失败**：
  - 检查 `web/package.json` 中的 `proxy` 配置
  - 确保后端服务已启动
- **构建失败**：
  - 清理缓存：`go clean -cache`、`yarn cache clean`
  - 重新安装依赖：`go mod download`、`yarn install`

## External Dependencies

### 必需依赖
- **数据库**：
  - MySQL 5.7+ / 8.0+（推荐，生产环境）
  - PostgreSQL 12+（可选）
  - SQLite 3（可选，仅用于开发/测试）
- **缓存**：
  - Redis 6.0+（必需，用于会话、限流、缓存）

### 可选依赖

#### 存储服务
- **对象存储**（用于存储任务结果、用户上传文件等）：
  - 阿里云 OSS（Aliyun OSS）
  - SMMS 图床
  - Imgur 图床
  - 本地文件系统（默认）

#### 监控与可观测性
- **Prometheus**：
  - 用途：收集和查询指标数据
  - 端点：`/api/metrics`（需 BasicAuth）
  - 指标类型：请求数、错误率、响应时间、渠道状态等
- **Uptime Kuma**：
  - 用途：状态页监控，展示服务可用性
  - 集成方式：读取 Uptime Kuma 的状态页数据
  - 配置：通过环境变量或配置文件启用

#### 身份认证与登录
- **OIDC（OpenID Connect）**：
  - 支持任何符合 OIDC 标准的身份提供商
  - 配置：Issuer URL、Client ID、Client Secret
- **GitHub OAuth**：
  - 用途：使用 GitHub 账号登录
  - 配置：GitHub OAuth App（Client ID、Client Secret）
- **微信登录**：
  - 用途：使用微信账号登录
  - 配置：微信开放平台应用（AppID、AppSecret）
- **飞书登录（Lark）**：
  - 用途：使用飞书账号登录
  - 配置：飞书开放平台应用
- **WebAuthn（Passkey）**：
  - 用途：无密码登录（生物识别、安全密钥）
  - 支持：FIDO2 标准

#### 支付网关
- **Epay**：
  - 用途：易支付聚合支付
  - 配置：商户 ID、密钥、API 地址
- **Stripe**：
  - 用途：国际信用卡支付
  - 配置：Publishable Key、Secret Key、Webhook Secret
- **支付宝（Alipay）**：
  - 用途：支付宝支付
  - 配置：App ID、私钥、公钥
- **微信支付（WeChat Pay）**：
  - 用途：微信支付
  - 配置：商户号、API Key、证书

#### 通知服务
- **Telegram Bot**：
  - 用途：发送通知消息（额度不足、渠道异常等）
  - 配置：Bot Token、Chat ID
  - Webhook：支持 Telegram Bot Webhook
- **邮件（SMTP）**：
  - 用途：发送邮件通知
  - 配置：SMTP 服务器、端口、用户名、密码

#### AI 供应商 API
- **OpenAI**：
  - API：https://api.openai.com/v1
  - 模型：GPT-3.5、GPT-4、DALL-E、Whisper、TTS 等
- **Azure OpenAI**：
  - API：https://{resource}.openai.azure.com
  - 模型：与 OpenAI 相同，但部署在 Azure
- **Azure Speech**：
  - 用途：模拟 TTS 功能
  - API：Azure Cognitive Services Speech API
- **Anthropic（Claude）**：
  - API：https://api.anthropic.com
  - 模型：Claude 3 系列
- **Google Gemini**：
  - API：https://generativelanguage.googleapis.com
  - 模型：Gemini Pro、Gemini Pro Vision 等
- **百度文心（Baidu）**：
  - API：https://aip.baidubce.com
  - 模型：ERNIE 系列
- **通义千问（Alibaba）**：
  - API：https://dashscope.aliyuncs.com
  - 模型：Qwen 系列
- **讯飞星火（iFlytek）**：
  - API：https://spark-api.xf-yun.com
  - 模型：Spark 系列
- **智谱（Zhipu）**：
  - API：https://open.bigmodel.cn
  - 模型：ChatGLM 系列
- **腾讯混元（Tencent）**：
  - API：https://hunyuan.tencentcloudapi.com
  - 模型：Hunyuan 系列
- **百川（Baichuan）**：
  - API：https://api.baichuan-ai.com
  - 模型：Baichuan 系列
- **MiniMax**：
  - API：https://api.minimax.chat
  - 模型：MiniMax 系列
- **Deepseek**：
  - API：https://api.deepseek.com
  - 模型：Deepseek 系列
- **Moonshot（月之暗面）**：
  - API：https://api.moonshot.cn
  - 模型：Moonshot 系列
- **Mistral AI**：
  - API：https://api.mistral.ai
  - 模型：Mistral 系列
- **Groq**：
  - API：https://api.groq.com
  - 模型：Llama、Mixtral 等
- **Amazon Bedrock**：
  - API：AWS Bedrock API
  - 模型：仅支持 Anthropic Claude 模型
- **零一万物（01.AI）**：
  - API：https://api.lingyiwanwu.com
  - 模型：Yi 系列
- **Cloudflare AI**：
  - API：https://api.cloudflare.com/client/v4/accounts/{account_id}/ai
  - 模型：Llama、Mistral、图像生成、STT 等
- **Cohere**：
  - API：https://api.cohere.ai
  - 模型：Command、Embed 等
- **Stability AI**：
  - API：https://api.stability.ai
  - 模型：Stable Diffusion（图像生成）
- **Coze**：
  - API：https://api.coze.com
  - 模型：Coze 系列
- **Ollama**：
  - API：本地部署的 Ollama 服务
  - 模型：支持 Ollama 支持的所有模型
- **OpenRouter**：
  - API：https://openrouter.ai/api/v1
  - 模型：聚合多个供应商的模型
- **xAI**：
  - API：https://api.x.ai
  - 模型：Grok 系列
- **Vertex AI（Google Cloud）**：
  - API：Google Cloud Vertex AI API
  - 模型：Gemini、PaLM 等
- **Replicate**：
  - API：https://api.replicate.com
  - 模型：各种开源模型
- **SiliconFlow**：
  - API：https://api.siliconflow.cn
  - 模型：SiliconFlow 系列
- **Volcano Engine（火山引擎）**：
  - API：火山引擎 API
  - 模型：火山引擎模型
- **Kling（快手）**：
  - API：快手 Kling API
  - 模型：Kling 视频生成
- **Vidu**：
  - API：Vidu API
  - 模型：Vidu 视频生成
- **Recraft AI**：
  - API：Recraft AI API
  - 模型：Recraft 图像生成
- **Jina AI**：
  - API：https://api.jina.ai
  - 模型：Jina Embeddings

#### 第三方服务集成
- **Midjourney**：
  - 集成方式：通过 [midjourney-proxy](https://github.com/novicezk/midjourney-proxy)
  - 用途：Midjourney 图像生成
- **Suno**：
  - 集成方式：通过 [Suno-API](https://github.com/Suno-API/Suno-API)
  - 用途：Suno 音乐生成

### 开发工具依赖
- **Docker**：用于本地开发环境（MySQL、Redis）
- **Docker Compose**：编排开发环境容器
- **Air**：Go 热重载工具
- **golangci-lint**：Go 代码检查工具
- **Swaggo**：Swagger 文档生成工具
- **Vite**：前端构建工具
- **ESLint**：前端代码检查工具
- **Prettier**：前端代码格式化工具
- **Playwright**：E2E 测试工具

### 网络依赖
- **HTTP/HTTPS 代理**（可选）：
  - 用途：访问被墙的 AI 供应商 API
  - 配置：渠道级别配置 HTTP/SOCKS5 代理
- **DNS**：
  - 需要能够解析各 AI 供应商的域名
- **防火墙**：
  - 需要允许出站 HTTPS 请求（443 端口）
  - 需要允许入站 HTTP/HTTPS 请求（配置的端口，默认 3000）


## Build and Deployment

### 构建生产版本
1. **构建前端**：
   ```bash
   cd web
   yarn build
   ```
   构建产物位于 `web/build/` 目录

2. **构建后端**（嵌入前端）：
   ```bash
   go build -o one-api
   ```
   或使用 Taskfile：
   ```bash
   task build
   ```

3. **构建 Docker 镜像**：
   ```bash
   docker build -t one-hub:latest .
   ```

### 部署方式

#### 1. 二进制部署
- 将构建好的 `one-api` 二进制文件上传到服务器
- 创建配置文件 `config.yaml`
- 启动服务：`./one-api`
- 使用 systemd 管理服务（参考 `one-api.service`）

#### 2. Docker 部署
```bash
docker run -d \
  --name one-hub \
  -p 3000:3000 \
  -v /path/to/config.yaml:/app/config.yaml \
  -v /path/to/logs:/app/logs \
  one-hub:latest
```

#### 3. Docker Compose 部署
```bash
docker-compose up -d
```
配置文件：`docker-compose.yml`

### 环境变量配置
常用环境变量：
- `PORT`：服务端口（默认 3000）
- `GIN_MODE`：Gin 模式（debug/release）
- `LOG_LEVEL`：日志级别（debug/info/warn/error）
- `SESSION_SECRET`：会话密钥
- `SQL_DSN`：数据库连接字符串
- `REDIS_CONN_STRING`：Redis 连接字符串
- `MEMORY_CACHE_ENABLED`：是否启用内存缓存
- `SYNC_FREQUENCY`：同步频率（秒）
- `CHANNEL_UPDATE_FREQUENCY`：渠道更新频率（分钟）
- `CHANNEL_TEST_FREQUENCY`：渠道测试频率（分钟）

### 反向代理配置

#### Nginx 示例
```nginx
server {
    listen 80;
    server_name your-domain.com;
    
    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

#### Caddy 示例
```
your-domain.com {
    reverse_proxy localhost:3000
}
```

## Performance Optimization

### 后端优化
- **启用内存缓存**：设置 `MEMORY_CACHE_ENABLED=true`
- **调整同步频率**：根据实际需求调整 `SYNC_FREQUENCY`
- **数据库连接池**：调整 GORM 连接池配置
- **Redis 连接池**：调整 Redis 客户端配置
- **日志级别**：生产环境使用 `info` 或 `warn` 级别

### 前端优化
- **代码分割**：Vite 自动进行代码分割
- **懒加载**：路由级别的懒加载
- **资源压缩**：Vite 自动压缩 JS/CSS
- **CDN**：将静态资源部署到 CDN

### 数据库优化
- **索引**：为常用查询字段添加索引
- **查询优化**：避免 N+1 查询
- **连接池**：合理配置连接池大小
- **定期清理**：定期清理过期日志和数据

### 缓存优化
- **Redis 持久化**：根据需求选择 RDB 或 AOF
- **缓存过期策略**：合理设置缓存过期时间
- **缓存预热**：启动时预加载热点数据

## Monitoring and Maintenance

### 监控指标
- **系统指标**：CPU、内存、磁盘、网络
- **应用指标**：请求数、错误率、响应时间
- **业务指标**：用户数、订单数、消费金额、渠道状态

### 日志管理
- **日志轮转**：自动轮转日志文件（Lumberjack）
- **日志聚合**：使用 ELK/Loki 等日志聚合工具
- **日志分析**：定期分析错误日志和慢查询

### 备份策略
- **数据库备份**：
  - 每日全量备份
  - 每小时增量备份（可选）
  - 保留最近 7 天的备份
- **配置备份**：
  - 配置文件纳入版本控制
  - 敏感信息使用环境变量或密钥管理服务

### 故障恢复
- **数据库恢复**：从备份恢复数据库
- **服务重启**：使用 systemd 或 Docker 自动重启
- **回滚**：保留多个版本的二进制文件或 Docker 镜像

## Security Best Practices

### 应用安全
- **定期更新依赖**：及时更新 Go 模块和 npm 包
- **安全扫描**：使用 `go mod verify`、`npm audit`
- **输入验证**：严格验证用户输入
- **输出编码**：防止 XSS 攻击
- **SQL 注入防护**：使用 GORM 参数化查询
- **CSRF 防护**：使用 CSRF Token

### 网络安全
- **HTTPS**：生产环境必须使用 HTTPS
- **防火墙**：限制不必要的端口访问
- **DDoS 防护**：使用 Cloudflare 等 CDN 服务
- **限流**：防止恶意请求

### 数据安全
- **加密存储**：敏感数据加密存储
- **传输加密**：使用 HTTPS/TLS
- **访问控制**：最小权限原则
- **审计日志**：记录敏感操作

## Troubleshooting Guide

### 常见错误及解决方案

#### 1. 数据库连接失败
- **错误信息**：`failed to connect to database`
- **可能原因**：
  - 数据库服务未启动
  - 连接信息错误
  - 网络不通
- **解决方案**：
  - 检查数据库服务状态
  - 验证连接字符串
  - 测试网络连通性

#### 2. Redis 连接失败
- **错误信息**：`failed to connect to redis`
- **可能原因**：
  - Redis 服务未启动
  - 连接信息错误
  - 密码错误
- **解决方案**：
  - 检查 Redis 服务状态
  - 验证连接字符串和密码
  - 使用 `redis-cli` 测试连接

#### 3. 渠道测试失败
- **错误信息**：`channel test failed`
- **可能原因**：
  - API Key 无效
  - 网络问题
  - 供应商服务异常
  - 代理配置错误
- **解决方案**：
  - 验证 API Key
  - 检查网络连接
  - 查看供应商状态页
  - 检查代理配置

#### 4. 前端无法访问
- **错误信息**：`Cannot GET /`
- **可能原因**：
  - 后端服务未启动
  - 端口被占用
  - 前端未构建
- **解决方案**：
  - 启动后端服务
  - 更换端口
  - 重新构建前端

#### 5. 支付回调失败
- **错误信息**：`payment callback failed`
- **可能原因**：
  - Webhook URL 配置错误
  - 签名验证失败
  - 网络问题
- **解决方案**：
  - 检查 Webhook URL 配置
  - 验证签名密钥
  - 查看支付网关日志

## Resources and References

### 官方文档
- [演示网站](https://one-hub.xiao5.info/)
- [项目文档](https://one-hub-doc.vercel.app/)
- [GitHub 仓库](https://github.com/MartialBE/one-hub)

### 相关项目
- [one-api](https://github.com/songquanpeng/one-api)（上游项目）
- [new-api](https://github.com/Calcium-Ion/new-api)（Midjourney/Suno 模块来源）
- [midjourney-proxy](https://github.com/novicezk/midjourney-proxy)
- [Suno-API](https://github.com/Suno-API/Suno-API)

### 技术栈文档
- [Go](https://go.dev/doc/)
- [Gin](https://gin-gonic.com/docs/)
- [GORM](https://gorm.io/docs/)
- [React](https://react.dev/)
- [Material-UI](https://mui.com/)
- [Vite](https://vitejs.dev/)
- [Redis](https://redis.io/docs/)

### 社区支持
- 交流群：见 README.md 中的二维码
- Issue 反馈：[GitHub Issues](https://github.com/MartialBE/one-hub/issues)

## Changelog and Versioning

### 版本号规则
- 遵循语义化版本（Semantic Versioning）
- 格式：`MAJOR.MINOR.PATCH`
  - MAJOR：不兼容的 API 变更
  - MINOR：向后兼容的功能新增
  - PATCH：向后兼容的问题修复

### 版本文件
- 根目录：`VERSION`
- 前端：`web/VERSION`
- 代码中：`common/config/version.go`

### 发布流程
1. 更新版本号
2. 更新 CHANGELOG（如有）
3. 提交代码
4. 创建 Git Tag
5. 推送到远程仓库
6. 创建 GitHub Release
7. 构建并发布 Docker 镜像

## Contributing Guidelines

### 贡献方式
- 报告 Bug：提交 Issue
- 功能建议：提交 Issue 讨论
- 代码贡献：提交 Pull Request
- 文档改进：提交 Pull Request

### PR 提交规范
- 遵循代码风格（golangci-lint、ESLint）
- 补充单元测试
- 更新相关文档
- 提供测试截图
- 关联相关 Issue

### 代码审查标准
- 代码质量：可读性、可维护性
- 功能完整性：是否满足需求
- 测试覆盖：是否有足够的测试
- 文档完整：是否更新了文档
- 安全性：是否存在安全隐患

## License

本项目基于 [MIT License](LICENSE) 开源。

## Acknowledgments

感谢以下开源项目：
- [one-api](https://github.com/songquanpeng/one-api)
- [Berry Free React Admin Template](https://github.com/codedthemes/berry-free-react-admin-template)
- [minimal-ui-kit](https://github.com/minimal-ui-kit/material-kit-react)
- [new api](https://github.com/Calcium-Ion/new-api)
- [go-zero](https://github.com/zeromicro/go-zero)
