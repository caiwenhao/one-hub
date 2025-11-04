## 1. Implementation
- [ ] 1.1 配置：新增 `domains` 品牌映射结构与加载（含 OAuth 凭据、CORS 允许源、前端地址）
- [ ] 1.2 中间件：BrandContext（从 `Host`/`X-Forwarded-Host`/`X-Forwarded-Proto` 推导 BaseURL 与品牌）
- [ ] 1.3 CORS：改为 `AllowOriginFunc` 按品牌校验并回显 Origin，`AllowCredentials=true`
- [ ] 1.4 Session Cookie：`Secure` 随 scheme 动态、`SameSite=Lax`、Host-only Domain
- [ ] 1.5 前端路由：router/main.go 使用品牌前端地址进行 NoRoute 重定向
- [ ] 1.6 OIDC：
  - [ ] Provider/Verifier 单例保留
  - [ ] 新增 `BuildOAuth2Config(c)` 动态 RedirectURL
  - [ ] OAuth 凭据按品牌读取（client_id/secret）
  - [ ] 回调处理兼容多域 state 校验
- [ ] 1.7 WebAuthn：按 RPID/Origin 懒加载实例池 `GetWebAuthn(c)`，控制器改用该接口
- [ ] 1.8 邮件/机器人：支持从请求上下文透传 BaseURL；无上下文时使用默认品牌或显式参数
- [ ] 1.9 文档：更新 `config.example.yaml` 与部署文档（Nginx 头、如何新增品牌域）

## 2. Tests
- [ ] 2.1 双域端到端：登录（密码、GitHub、OIDC）、回跳保持会话
- [ ] 2.2 WebAuthn 多域注册/登录（RPID/Origin 正确）
- [ ] 2.3 CORS 预检与带凭证请求通过（不同品牌）
- [ ] 2.4 邮件中的链接按触发品牌域名生成可访问
- [ ] 2.5 代理 https/http 组合下 RedirectURL 与 Cookie Secure 正常

## 3. Migration
- [ ] 3.1 在 IDP/OAuth 平台注册各品牌回调 URL
- [ ] 3.2 Nginx 透传 `Host`/`X-Forwarded-Proto`/`X-Real-IP`
- [ ] 3.3 老节点保持默认全局行为（未命中品牌时退化为单域逻辑）

