# 视频 URL 代理（统一文档）

本功能用于隐藏上游视频供应商域名，返回看起来像直链的视频地址，形式为：

- https://your-worker.workers.dev/v/<token>.mp4

当前仅支持 Proxy（流式转发）模式，不再支持 Redirect 模式；仅支持“面板配置”，不再支持 YAML/环境变量配置。
支持两种外观：
- 标准：`/v/<short>.mp4?token=<长加密串>`（默认，零依赖）
- KV 短链：`/v/<slug>.mp4`（可选，需绑定 KV 并开启后端开关）

## 功能概览

- 完全隐藏真实上游域名
- AES‑256‑GCM 加密 + SHA‑256 密钥派生
- Token 自动过期（默认 24 小时）
- 链接外观为 .mp4 直链（/v/<token>.mp4）
- Cloudflare Workers 边缘加速

## 快速开始（5 分钟）

1) 部署 Cloudflare Worker

- Dash → Workers & Pages → Create Worker → 命名 `video-proxy`
- 将仓库中的 `cloudflare-workers/video-proxy.js` 复制到在线编辑器并部署
- 在 Worker 的 Settings → Variables 添加环境变量：
  - 名称：`API_KEY`（Secret 类型），值为随机密钥

2) 后端面板配置（仅面板）

- 后台 → 运营设置 → 其他设置 → 填写：
  - `CFWorkerVideoUrl`：例如 `https://video-proxy.your-account.workers.dev`
  - `CFWorkerVideoKey`：与 Worker 的 `API_KEY` 一致
- 保存后新产生/查询到的视频 URL 将自动替换为代理 URL

3) （可选）开启 KV 短链模式

- 在 Worker 绑定 KV 命名空间，变量名：`VIDEO_KV`
- 无需改代码，保持同一 Worker；绑定完成后可直接使用
- 后端面板 → 运营设置 → 其他设置：
  - `CFWorkerVideoKVEnabled`：开启
  - `CFWorkerVideoShortLength`：短码长度（默认 12）

4) 验证

```bash
# Worker 健康检查
curl https://video-proxy.your-account.workers.dev/health
# 预期：{"status":"ok","service":"video-proxy","mode":"proxy"}

# 业务侧：创建视频任务并查询状态，检查返回的 result.video_url
# 预期形如：https://video-proxy.your-account.workers.dev/v/<token>.mp4
```

## 部署与配置细节

### Cloudflare Worker 端

- 仅 Proxy 模式：Worker 解密 token 后发起上游请求并流式返回
- 路由：
  - `/v/<token>.mp4` 或 `/v/<token>`（token 可为“长加密串”或“短码 slug”）
  - 兼容旧路由：`/video?token=<token>`（后端默认不再生成）
  - `/register`（POST）：服务端注册短码 → KV 映射，需 `X-API-Key`
- 变量：
  - `API_KEY`（Secret）与后端 `CFWorkerVideoKey` 必须一致
  - `VIDEO_KV`（KV Namespace，可选，短链模式需要）

### 后端端

- 面板选项：`CFWorkerVideoUrl`、`CFWorkerVideoKey`
- KV 短链（可选）：`CFWorkerVideoKVEnabled`、`CFWorkerVideoShortLength`
- 只要配置了 `CFWorkerVideoUrl`，服务端会在返回前把 `result.video_url`、`thumbnail_url`、`download_url`、`spritesheet_url` 等通过 `common/video` 的加密方法改写为 Worker 代理 URL

## 安全与最佳实践

- 加密失败时建议通过日志告警（本仓库默认静默回退，必要时可加强）
- 可在 Worker 端加上 scheme/host 白名单校验，降低密钥泄露后的滥用风险
- 建议使用自定义域名与就近路由，提升下载体验
- 若上游返回有限 TTL，可在服务端取 `min(上游 TTL, 默认上限)` 作为 token 过期（当前默认 24h）

## 成本与性能

- 免费版配额参考：10万次请求/天、10GB 出站/天（以 Cloudflare 官方为准）
- 代理为流式转发，速度略慢于直连；可通过缓存/CDN/自定义域名优化

## 故障排查（FAQ）

- 400 Invalid token：前后端密钥不一致或 token 过期 → 核对 `API_KEY` 与面板密钥
- 链接仍为上游域名：后端未读取到面板配置或未保存成功 → 检查“其他设置”并保存
- 下载慢：使用就近域名，或升级套餐、配合 CDN 缓存
- 404/未命中短码：检查后端是否开启 KV 短链、Worker 是否绑定 `VIDEO_KV`、注册是否成功（/register）

## 迁移与回滚

- 迁移：只需按上文新增 Worker 与面板两项配置即可，旧任务链接仍按原样可用
- 回滚：
  - 仅关闭短链：关闭 `CFWorkerVideoKVEnabled`，将恢复 `?token=` 样式；已生成的短码不会再被返回
  - 完全回滚：清空 `CFWorkerVideoUrl` 并保存，后端将恢复返回上游真实 URL

## 变更记录

- 2024-11-09：移除 Redirect 模式，仅保留 Proxy；移除 YAML/环境变量配置，改为仅面板配置；链接格式改为 `/v/<token>.mp4`
