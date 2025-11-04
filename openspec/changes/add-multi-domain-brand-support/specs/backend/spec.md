## ADDED Requirements

### Requirement: Brand Context And Routing
后端必须基于每个请求的 Host/Proto 识别品牌并应用对应配置（BaseURL、前端地址、CORS、OAuth 凭据）。

#### Scenario: Resolve brand by host
- WHEN 请求携带 `Host=a.example.com` 且 `domains[].host` 包含该主机
- THEN 在请求上下文中提供 `brand.base_url=https://a.example.com`、`brand.frontend_base_url`、`brand.cors.allowed_origins`、`brand.oauth.*`

#### Scenario: Fallback to global settings
- WHEN 未命中任何品牌
- THEN 使用全局 `ServerAddress`/`frontend_base_url`/CORS 默认

### Requirement: Frontend Redirect Per Brand
当启用前后端分离时，后端对页面路由 `NoRoute` 应重定向至当前品牌前端地址。

#### Scenario: Redirect to brand frontend
- WHEN 请求 `GET /` to `b.example.com`
- THEN 301/302 重定向至 `brand.frontend_base_url + '/'`

### Requirement: Dynamic CORS With Credentials
后端必须按品牌配置校验请求 Origin，并在 `AllowCredentials=true` 情形下回显具体允许的 Origin。

#### Scenario: Preflight allowed
- WHEN 请求 Origin 在 `brand.cors.allowed_origins` 列表
- THEN 预检与实际请求均返回 `Access-Control-Allow-Origin=<该Origin>` 且允许携带 Cookie

#### Scenario: Preflight denied
- WHEN 请求 Origin 不在允许列表
- THEN 拒绝跨域请求，不返回 `Access-Control-Allow-Origin`

### Requirement: Session Cookie Policy For Multi-Domain
Cookie 安全属性必须依据请求 scheme 动态设置，避免 OAuth 回跳丢会话并保障安全。

#### Scenario: Secure over HTTPS
- WHEN 请求为 HTTPS（或 `X-Forwarded-Proto=https`）
- THEN 设置 `Secure=true` 且 `SameSite=Lax`，Domain 不设置（host-only）

#### Scenario: HTTP fallback
- WHEN 请求为 HTTP 测试环境
- THEN `Secure=false`，其余一致

### Requirement: OIDC RedirectURL Per Brand And Credentials Per Brand
OIDC 登录必须使用按品牌构造的 RedirectURL 与对应品牌的 client_id/client_secret。

#### Scenario: Build login URL per brand
- WHEN 请求来自 `a.example.com`
- THEN 使用 `client_id=domains[a].oauth.oidc.client_id`，`RedirectURL=https://a.example.com/api/oauth/oidc`

#### Scenario: Callback validation with session state
- WHEN OIDC 回调在 `b.example.com` 到达且 state 合法
- THEN 使用 b 品牌凭据成功验证并建立会话

### Requirement: WebAuthn Per-RPID Instance Pool
WebAuthn 必须按 RPID/Origin 懒加载实例，支持多域 RP 配置。

#### Scenario: Register on brand A
- WHEN WebAuthn 注册请求在 `a.example.com`
- THEN 使用 RPID=`a.example.com`、RPOrigins 包含 `https://a.example.com`

#### Scenario: Login on brand B
- WHEN WebAuthn 登录请求在 `b.example.com`
- THEN 使用 RPID=`b.example.com`、RPOrigins 包含 `https://b.example.com`

### Requirement: Brand-Aware Absolute Links (Email/Robot)
系统在生成邮件/机器人内链接时应优先使用触发请求的品牌 BaseURL；无上下文时允许指定默认品牌。

#### Scenario: Email link from brand A action
- WHEN 用户在 `a.example.com` 触发 "重置密码"
- THEN 邮件中的链接以 `https://a.example.com` 开头

#### Scenario: Background task without request context
- WHEN 定时任务发送用量提醒
- THEN 链接以默认品牌 BaseURL 或显式配置的 BaseURL 开头

