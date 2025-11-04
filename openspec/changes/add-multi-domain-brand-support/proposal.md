## Why
现有后端默认单域假设：固定 `ServerAddress`、静态 OIDC RedirectURL、单实例 WebAuthn、全局前端重定向等，无法满足“多品牌（多域名）并行、各自独立前端与 OAuth 凭据、邮件/机器人回链按品牌域名生成”的需求。

## What Changes
- 引入按请求 Host 解析的 BrandContext 中间件与配置（domains 映射）
- CORS 调整为按品牌 Origin 校验与回显，允许带凭证跨域
- Session Cookie 策略按请求 scheme 动态 Secure，SameSite=Lax（避免 OAuth 回跳丢会话）
- OIDC：保留 Provider/Verifier 单例，RedirectURL 改为按请求动态构造；OAuth 凭据按品牌配置
- WebAuthn：按 RPID/Origin 懒加载实例池，支持多域
- 前端重定向：基于品牌前端地址路由，老品牌继续可使用内置前端
- 邮件/机器人链接：按品牌域名生成（从请求上下文或任务入参透传）
- 配置文件新增 `domains[].host/server_address/frontend_base_url/cors.allow_origins/oauth.*`

## Impact
- 受影响模块：router、middleware（cors、auth）、common（oidc、webauthn、stmp、telegram）、controller（webauthn、oauth）、配置加载
- 安全：CORS 限定 Origin；Cookie Secure/ SameSite 调整
- 运维：Nginx 需透传 `Host`/`X-Forwarded-Proto`

