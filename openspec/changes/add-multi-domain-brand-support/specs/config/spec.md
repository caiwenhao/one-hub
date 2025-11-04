## ADDED Requirements

### Requirement: Domain Mapping Configuration
系统必须支持在配置文件中声明多品牌域名映射，并可为不同品牌配置前端地址、CORS 允许源以及 OAuth 凭据。

#### Scenario: Example structure
- WHEN 管理员在配置中添加如下片段
- THEN 系统加载并在请求时按 host 命中对应品牌

```
domains:
  - host: a.example.com
    server_address: https://a.example.com
    frontend_base_url: https://app-a.example.com
    cors:
      allow_origins:
        - https://app-a.example.com
        - https://a.example.com
    oauth:
      oidc:
        client_id: A_CLIENT_ID
        client_secret: A_CLIENT_SECRET
        issuer: https://idp.example.com
        scopes: "openid,email,profile"
      github:
        client_id: GH_A
        client_secret: GH_A_SECRET
  - host: b.example.com
    server_address: https://b.example.com
    frontend_base_url: https://b.example.com
    cors:
      allow_origins:
        - https://b.example.com
    oauth:
      oidc:
        client_id: B_CLIENT_ID
        client_secret: B_CLIENT_SECRET
        issuer: https://idp.example.com
        scopes: "openid,email,profile"
```

#### Scenario: Fallback defaults
- WHEN `domains` 未配置或 host 未命中
- THEN 使用全局 `ServerAddress`、`frontend_base_url` 与默认 CORS 策略

