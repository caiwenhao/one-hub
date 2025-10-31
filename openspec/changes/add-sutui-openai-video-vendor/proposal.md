## Why
为 OpenAI 视频接口（/v1/videos）新增“速推”上游供应商，以对接现有 st-ai-sora2 能力，统一走 OpenAI 兼容路由，减少调用方改造成本。

## What Changes
- 在 OpenAI 渠道新增供应商标识 `sutui`（除现有 `official`、`mountsea` 之外）。
- Provider 侧实现 `sutui` 的 Create/Retrieve/Download/Remix 适配：
  - Create：支持 JSON 与 multipart/form-data 直透/转换，映射“速推”响应到 `types.VideoJob`（将顶层 `video_url` 映射为 `result.video_url`）。
  - Retrieve：将“速推”的查询响应映射到 `types.VideoJob` 结构。
  - Download：若仅提供 `video_url`，通过代理下载/透传内容；不支持的 `variant` 返回 501。
  - Remix：若“速推”无官方 remix 路由，回退为 `create(remix_video_id=...)` 语义。
- 渠道配置：支持通过 `channel.plugin.sora.vendor=sutui` 指定；或根据 BaseURL 命中 `sutui`/`st-ai` 关键字自动识别。
- 错误映射：统一转换为 OpenAI 错误结构 `{ error: { message, type, code, param } }`。
- 计费：沿用现有 Sora 秒级换算策略，按 `seconds × 1000`（或与现有实现一致）。

## Impact
- 影响能力：OpenAI 视频（/v1/videos）
- 影响代码：
-  - `providers/openai/video.go`（新增 sutui 分支与映射）
-  - 文档：新增/更新“OpenAI 视频 + 速推渠道”的调用说明
-  - 可选：`router/swagger_openai.go`（补充说明 sutui 适配，不改路由）

