## Context
- 两个品牌独立域名与前端部署，新品牌由 Nginx 反代 API；需要后端基于请求 Host 识别品牌并应用不同配置（前端地址、OAuth 凭据、CORS、生成链接的 BaseURL 等）
- 现状为单域假设：全局 `ServerAddress`、静态 OIDC RedirectURL、单实例 WebAuthn、CORS 允许所有 Origin 等

## Goals / Non-Goals
- Goals：多域品牌并行、按品牌 OAuth、邮件/机器人链接按品牌生成、安全 CORS、会话稳定
- Non-Goals：引入复杂的租户数据隔离（本次仅运行时配置层面的多域支持）

## Decisions
- BrandContext 中间件：
  - 解析 `Host`/`X-Forwarded-Host` 与 `X-Forwarded-Proto` 推导 `scheme://host` 作为 BaseURL
  - 命中 `domains[].host` 则载入品牌配置（优先于全局），注入到 `gin.Context`
- CORS：使用 `AllowOriginFunc(func (origin) bool)` 校验是否在品牌允许列表中，并回显该 Origin；`AllowCredentials=true`
- Session Cookie：`Secure` 按请求 scheme、`SameSite=Lax`、Host-only Domain，兼容 OAuth 回跳
- OIDC：
  - Provider/Verifier 仍可单例（与 ClientID 绑定时可按“默认品牌”或多 provider 映射；若各品牌不同 IdP，支持 brand→provider cache）
  - 登录跳转/回调使用 `BuildOAuth2Config(c)` 动态构造 RedirectURL 与 client_id/secret（按品牌）
- WebAuthn：按 RPID/Origin 懒加载 `*webauthn.WebAuthn` 实例池（sync.Map），`GetWebAuthn(c)` 返回实例
- 生成绝对链接（邮件/机器人）：优先从触发请求的品牌 BaseURL 透传；无上下文场景允许指定默认品牌或显式 BaseURL

## Alternatives considered
- 保持全局单域：不满足多品牌需求
- 仅在 Nginx 层做路径分流：无法解决 OIDC/WebAuthn/邮件中的绝对链接与 CORS/会话问题

## Risks / Trade-offs
- 需要在 IdP/OAuth 平台注册多个回调 URL；不同品牌凭据管理复杂度上升 → 通过配置中心化与校验降低风险
- CORS 配置错误会导致跨域失败或安全风险 → 严格校验 Origin，默认拒绝

## Migration Plan
1) 上线前在各 IdP 注册 `{brand_base}/api/oauth/oidc` 等回调
2) Nginx 透传 `Host`、`X-Forwarded-Proto`、`X-Real-IP`
3) 配置 `domains` 映射并灰度单品牌验证 → 双品牌 → 全量
4) 回退：删除 `domains` 配置即回退单域行为；保留原全局配置

## Open Questions
- 是否存在“不同 IdP/Provider” 的品牌？若是需为每品牌维护 provider/verifier 实例缓存

