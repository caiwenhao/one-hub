# OpenAI 风格视频 API（Veo 适配）

最后更新时间：2025-11-12

本文件定义在不改变现有 OpenAI 接口的前提下，通过统一入口 `/v1/videos` 使用 Google Gemini Veo 模型（例如 `veo-3.1-generate-preview`）进行视频生成、查询与下载的网关行为与参数规范。

## 目标与现状

- 目标：让调用方仅使用 OpenAI 风格接口，即可同时访问 OpenAI Sora 与 Google Gemini Veo 能力。
- 已实现能力：
  - Create（POST /v1/videos）：支持 `veo-*` 与 `sora-*`；Veo 走 Google 官方 `predictLongRunning`；Sora 维持原能力。
  - Retrieve（GET /v1/videos/{id}）：Veo 走 Google 官方 operations，完成后返回 `result.video_url`。
  - Download（GET /v1/videos/{id}/content）：直连上游 `video.uri` 并携带 `x-goog-api-key` 拉流返回。
- 未对接能力（Veo）：Remix/List/Delete 暂返回 501（Sora 保持可用）。

## 鉴权与限流

- 请求头：`Authorization: Bearer <token>`
- 速率：沿用网关动态限流策略（令牌组/渠道维度）。
- 上游访问：网关在访问 Google Gemini 时自动追加 `x-goog-api-key`，客户端无需关心。

## 支持的模型

- Sora（原支持）：`sora-2`、`sora-2-pro`
- Veo（新增适配，建议使用官方价目键）：
  - `veo-3.1-generate-preview`（高质量，含音频）
  - `veo-3.1-fast-generate-preview`（加速版）
  - 兼容别名（自动归一化）：`veo-3.1`、`veo3.1`、`veo3.1-fast`、`veo-3.0(-fast)`、`veo-2.0` 等

> 说明：Veo 首版适配仅对接“文本转视频（Text→Video）”。参考图/首尾帧/扩展视频等高级能力后续通过扩展字段开放。

## 端点一览

- 创建视频：`POST /v1/videos`
- 查询进度：`GET /v1/videos/{id}`
- 下载内容：`GET /v1/videos/{id}/content`
- 列表任务：`GET /v1/videos`（仅 Sora；Veo 返回 501）
- Remix：`POST /v1/videos/{id}/remix`（仅 Sora；Veo 返回 501）
- 删除任务：`DELETE /v1/videos/{id}`（仅 Sora；Veo 返回 501）

## 请求与响应规范

### 1) 创建视频（POST /v1/videos）

Headers

```
Authorization: Bearer <token>
Content-Type: application/json | multipart/form-data
```

Body（application/json）

| 字段    | 类型            | 必填 | 说明 |
|---------|-----------------|------|------|
| model   | string          | 是   | `veo-*` 或 `sora-*`，见“支持的模型”。为空时默认 `sora-2`（兼容旧行为）。|
| prompt  | string          | 是   | 文本提示词。|
| seconds | number or string| 否   | 支持 `4`/`6`/`8`；非法值回落 `6`。Veo@1080p 强制 `8` 秒。|
| size    | string          | 否   | 输出分辨率，支持：`1280x720`、`720x1280`、`1920x1080`、`1080x1920`；其他值回落 `1280x720`。|

Body（multipart/form-data）

- 同名文本字段 `model/prompt/seconds/size`。首版不处理文件字段（Veo 图/视频入参后续扩展）。

示例（Veo 3.1）

```bash
curl -X POST "$BASE_URL/v1/videos" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-3.1-generate-preview",
    "prompt": "A cinematic shot of a majestic lion in the savannah.",
    "seconds": 6,
    "size": "1280x720"
  }'
```

成功响应（排队中）

```json
{
  "id": "abCDefGh123",
  "object": "video",
  "status": "queued",
  "model": "veo-3.1-generate-preview",
  "prompt": "A cinematic shot of a majestic lion in the savannah.",
  "seconds": "6",
  "size": "1280x720",
  "progress": 0
}
```

### 2) 查询进度（GET /v1/videos/{id}）

语义：轮询任务状态；完成后返回视频 URL。

成功响应（进行中）

```json
{
  "id": "abCDefGh123",
  "object": "video",
  "status": "in_progress",
  "progress": 0,
  "model": "veo-3.1-generate-preview"
}
```

成功响应（已完成）

```json
{
  "id": "abCDefGh123",
  "object": "video",
  "status": "completed",
  "progress": 100,
  "model": "veo-3.1-generate-preview",
  "seconds": "6",
  "size": "1280x720",
  "result": {
    "video_url": "https://.../video.mp4"
  }
}
```

### 3) 下载内容（GET /v1/videos/{id}/content）

返回视频二进制流（`application/octet-stream`）。网关自动携带 `x-goog-api-key` 拉取上游资源。

```bash
curl -L -H "Authorization: Bearer $TOKEN" \
  "$BASE_URL/v1/videos/abCDefGh123/content" -o out.mp4
```

## 参数映射（Veo 专用）

尺寸到 Veo 参数：

| size       | aspectRatio | resolution |
|------------|-------------|------------|
| 1280x720   | 16:9        | 720p       |
| 720x1280   | 9:16        | 720p       |
| 1920x1080  | 16:9        | 1080p（强制 8s） |
| 1080x1920  | 9:16        | 1080p（强制 8s） |

秒数：仅允许 4/6/8，非法值回落 6；1080p 场景强制 8。

## 计费与用量

- 单位：秒（sec）。
- 规则：Veo 按“秒 × 1000 tokens”等价换算计费；价格键与模型名一致（如 `veo-3.1-generate-preview`）。
- 落账元数据：`unit: "sec"`、`video_seconds: <秒数>`。
- 解析优先级：若上游返回 `metadata.durationSeconds` 则以其为准；否则使用创建请求中的 `seconds`。

## 错误码

| HTTP | code               | 说明 |
|------|--------------------|------|
| 400  | invalid_model      | 仅允许 `sora-2`、`sora-2-pro`、`veo-*`。|
| 400  | video_not_ready    | 下载接口在未完成阶段被调用。|
| 404  | not_found          | 任务不存在或无权限。|
| 429  | rate_limited       | 触发令牌或渠道限流。|
| 5xx  | upstream_error     | 上游服务异常或解析失败。|
| 501  | not_implemented    | 当前模型不支持该操作（如 Veo 的 Remix/List/Delete）。|

## 兼容性与限制

- 兼容性：`model` 为空仍默认 `sora-2`；Sora 全量接口不变。
- Veo 限制：首版仅 Text→Video；参考图/首尾帧/扩展视频、negativePrompt、personGeneration 等暂未通过 `/v1/videos` 暴露。
- 上游变体：若 Gemini 渠道 BaseURL 指向第三方（如 sutui），本适配器不保证可用。

## 示例（端到端）

```bash
# 1) 创建 Veo 3.1 视频
curl -X POST "$BASE_URL/v1/videos" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-3.1-generate-preview",
    "prompt": "A lion on the savannah at sunset.",
    "seconds": 6,
    "size": "1280x720"
  }'

# 2) 轮询直至完成
while true; do
  r=$(curl -s -H "Authorization: Bearer $TOKEN" "$BASE_URL/v1/videos/abCDefGh123")
  st=$(echo "$r" | jq -r .status)
  if [ "$st" = "completed" ]; then echo "$r" | jq; break; fi
  sleep 10
done

# 3) 下载视频
curl -L -H "Authorization: Bearer $TOKEN" \
  "$BASE_URL/v1/videos/abCDefGh123/content" -o out.mp4
```

## 变更记录

- 2025-11-12：首版发布。打通 `veo-*` 经 `/v1/videos` 的 Create/Retrieve/Download；Remix/List/Delete 暂未对接。

