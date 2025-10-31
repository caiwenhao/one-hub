## Why
在 Gemini 渠道中，当前 Veo 视频能力仅支持官方 Google 原生接口（`:predictLongRunning` + `/operations/*`）。为降低统一视频能力的接入与迁移成本，希望在 Gemini 渠道内新增“速推（sutui/st-ai）”作为上游供应商，实现：继续通过 Gemini 风格路由（`/gemini/{version}/models/veo-*:predictLongRunning`、`/gemini/{version}/operations/{name}`）调用，但底层转接到速推的 OpenAI 风格视频接口（参考 `docs/api/st-ai-veo.md`）。

## What Changes
- 在 Gemini 渠道增加“速推（sutui）”上游选择与自动识别：
  - 优先读取渠道插件 `plugin.gemini.video.vendor=sutui`；
  - 备用：若 BaseURL 包含 `sutui`/`st-ai` 关键字则判定为 sutui。
- Veo 生成初始化：将 `POST /gemini/{version}/models/veo-*:predictLongRunning` 映射为 sutui 的 `POST /v1/videos`（支持 JSON 与 multipart/form-data 直透/映射），并返回 operation name（使用 sutui 返回的 `id` 封装）。
- 任务轮询：将 `GET /gemini/{version}/operations/{name}` 映射为 sutui 的 `GET /v1/videos/{id}` 并转换为 Gemini `operations` 响应结构（包含 `done`、`response.generateVideoResponse.generatedSamples[]`）。
- 下载：如需直链，保留 sutui 返回的 `video_url`（或在 `generatedSamples[].uri` 中给出）；不改变现有 `/v1/videos/{id}/content` 路由。
- 计费：Veo 依旧按“秒”计费，沿用 `seconds × 1000` 的内部换算。

## Impact
- 影响能力：Gemini 视频（Veo long-running）
- 影响代码：
  - `providers/gemini/*`（新增 sutui 检测与路由映射）
  - `relay/relay_util/quota.go`（计费透传，保持现状）
  - 可选：`docs/gemini/video.mdx`（补充 sutui 作为上游的使用说明）

