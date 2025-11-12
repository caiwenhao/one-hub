# OpenAI 风格视频 API（Veo 适配）

最后更新时间：2025-11-12

本文件定义在不改变现有 OpenAI 接口的前提下，通过统一入口 `/v1/videos` 使用 Google Gemini Veo 模型（例如 `veo-3.1-generate-preview`）进行视频生成、查询与下载的网关行为与参数规范。除文生视频外，本文亦补充 Veo 3.1 的参考图、首尾帧与扩展视频能力的统一用法。

## 目标与现状

- 目标：让调用方仅使用 OpenAI 风格接口，即可同时访问 OpenAI Sora 与 Google Gemini Veo 能力。
- 已实现能力：
  - Create（POST /v1/videos）：支持 `veo-*` 与 `sora-*`；Veo 走 Google 官方 `predictLongRunning`；Sora 维持原能力。
  - Retrieve（GET /v1/videos/{id}）：Veo 走 Google 官方 operations，完成后返回顶层 `video_url` 字段。
  - Download（GET /v1/videos/{id}/content）：直连上游 `video.uri` 并携带 `x-goog-api-key` 拉流返回。
- 不提供：列表、删除、Remix 端点（统一不暴露）。

## 鉴权与限流

- 请求头：`Authorization: Bearer <token>`
- 速率：沿用网关动态限流策略（令牌组/渠道维度）。
- 上游访问：网关在访问 Google Gemini 时自动追加 `x-goog-api-key`，客户端无需关心。

## 支持的模型

- Sora：`sora-2`、`sora-2-pro`
- Veo（仅接受官方型号字符串，不做别名归一化）：
  - `veo-3.1-generate-preview`
  - `veo-3.1-fast-generate-preview`

> 说明：Veo 3.1 起支持“参考图（1–3 张）”“首尾帧插值（2 张）”与“扩展视频（+7s 续写）”。本文给出统一入参与约束；具体可用性取决于所选通道是否直连 Google 官方 Gemini API（第三方代理可能不支持）。

## 端点一览

- 创建视频：`POST /v1/videos`
- 查询进度：`GET /v1/videos/{id}`
- 下载内容：`GET /v1/videos/{id}/content`

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
| prompt  | string          | 否   | 文本提示词。文生/图生必填；扩展视频可选（可用于引导续写段落）。|
| seconds | number or string| 否   | 支持 `4`/`6`/`8`；非法值回落 `6`。Veo@1080p 强制 `8` 秒。|
| size    | string          | 否   | 输出分辨率，支持：`1280x720`、`720x1280`、`1920x1080`、`1080x1920`；其他值回落 `1280x720`。|
| input_reference | string or string[] | 否   | 统一参考图字段：
  - 单值或数组；元素允许 `http(s)` 链接或 `data:` Base64（例如 `data:image/png;base64,...`）。
  - 1–3 张视为“参考图”；恰 2 张按“首/尾帧”解释并做插值。|
| extend_from | string       | 否   | 扩展视频：目标历史视频 ID（必须为 Veo 成片）。每次固定续写 +7 秒，`size` 限制为 720p。|

Body（multipart/form-data）

- 同名文本字段 `model/prompt/seconds/size/extend_from`。
- 文件字段仅保留 `input_reference`：可重复 1–3 次；若恰 2 次则解释为首尾帧。


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
  "video_url": "https://.../video.mp4"
}
```

### 3) 下载内容（GET /v1/videos/{id}/content）

返回视频二进制流（`application/octet-stream`）。网关自动携带 `x-goog-api-key` 拉取上游资源。

```bash
curl -L -H "Authorization: Bearer $TOKEN" \
  "$BASE_URL/v1/videos/abCDefGh123/content" -o out.mp4
```

## 高级能力（Veo 3.1）

以下三种能力仅在 `veo-3.1-*` 系列可用；部分通道仅支持其中国行 720p 路线，详见每小节“限制”。

### A. 参考图（Reference Images, 1–3 张）

语义：提供 1–3 张参考图来约束主体外观或风格。参考图越贴合目标主体，保真度越高。

- 入参方式：
  - JSON：`input_reference: ["https://.../a.jpg", "https://.../b.png"]` 或单值字符串
  - Multipart：重复 `-F input_reference=@/abs/a.jpg -F input_reference=@/abs/b.png`
- 约束与建议：
  - 数量：1–3 张；超过忽略
  - 类型/大小：jpeg/png/webp，建议 ≤10MB/张（data:URI 解码后 ≤10MB）
  - 分辨率：与目标输出比例接近（16:9 或 9:16），避免极端长宽比
  - 与 `seconds/size` 可同时出现；1080p 仍强制 8 秒

示例（multipart，2 张参考图）

```bash
curl -X POST "$BASE_URL/v1/videos" \
  -H "Authorization: Bearer $TOKEN" \
  -F "model=veo-3.1-generate-preview" \
  -F "prompt=银色跑车在夜色雨巷疾驰，霓虹反射" \
  -F "seconds=6" \
  -F "size=1280x720" \
  -F "input_reference=@/data/car_front.jpg" \
  -F "input_reference=@/data/car_side.jpg"
```

### B. 首尾帧插值（First/Last Frame Interpolation）

语义：同时提供“首帧”和“尾帧”两张图，Veo 在两帧之间生成具有电影感的过渡镜头。

- 入参方式：
  - JSON：`input_reference: [first_frame_url, last_frame_url]`
  - Multipart：两次 `-F input_reference=@/first.jpg -F input_reference=@/last.jpg`
- 约束与建议：
  - 仅 Veo 3.1；建议 720p 以提高成功率
  - 两张图片分辨率需相近，分辨率比控制在 0.8～1.25；长宽比应与目标 `size` 一致（16:9 或 9:16）
  - 文件类型同上；过大文件将被拒绝
  - 搭配 `prompt` 提供叙事与镜头语言将显著提升质感

示例（JSON）

```bash
curl -X POST "$BASE_URL/v1/videos" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-3.1-generate-preview",
    "prompt": "雾气森林中少女从秋千上消失的电影镜头，低饱和配乐氛围",
    "seconds": 6,
    "size": "1280x720",
    "input_reference": [
      "https://cdn.example.com/first.jpg",
      "https://cdn.example.com/last.jpg"
    ]
  }'
```

### C. 扩展视频（Extend +7s）

语义：对“已由 Veo 生成”的历史成片进行续写，每次固定延长 7 秒，可连续调用，最多累计 ~141 秒（官方限制）。

- 入参方式：
  - 统一使用创建端点：`POST /v1/videos`
  - 指定 `model=veo-3.1-*`，并传入 `extend_from=<历史视频ID>`
  - 可同时提供新的 `prompt` 引导续写段落；`size` 限制为 720p（16:9 或 9:16）
- 行为与限制：
  - 每次调用固定 +7 秒；忽略 `seconds` 入参（允许传但不会影响结果）
  - 仅支持基于 Veo 历史成片续写（Sora/第三方生成的视频不可用）
  - 最长约 141 秒；超过将返回 400/422

示例（续写 1 次）

```bash
curl -X POST "$BASE_URL/v1/videos" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-3.1-generate-preview",
    "extend_from": "abCDefGh123",  // 历史 Veo 成片ID
    "prompt": "镜头推进至花园深处，小狗追逐落在橙色纸花上的纸蝴蝶",
    "size": "1280x720"
  }'
```

示例（脚本连续续写 3 次）

```bash
VID=abCDefGh123
for i in 1 2 3; do
  VID=$(curl -s -X POST "$BASE_URL/v1/videos" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d "{\"model\":\"veo-3.1-generate-preview\",\"extend_from\":\"$VID\",\"size\":\"1280x720\"}" | jq -r .id)
  echo "extended -> $VID"
done
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

参考图/首尾帧/扩展 映射到 Gemini：

- 参考图（1–3 张）：
  - OpenAI：`input_reference`（单值或数组）
  - Gemini：`parameters.referenceImages[]`（或 `config.reference_images`），每个元素为 `asset` 引用；单图时也可走 `image`
- 首尾帧（2 张）：
  - OpenAI：`input_reference[0]=first` + `input_reference[1]=last`
  - Gemini：`image=first` + `parameters.lastFrame=last`
- 扩展视频（+7s）：
  - OpenAI：`extend_from=<历史ID>`
  - Gemini：`video=<历史视频 asset>` + 新 `prompt` + `resolution=720p`

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
 | 501  | not_implemented    | 当前模型或通道不支持该操作（例如第三方上游不支持扩展）。|
| 422  | invalid_arguments  | 参数组合非法，例如两张首尾帧分辨率差异过大、`size` 与图片比例不一致，或扩展目标非 Veo 成片。|

## 兼容性与限制

- 兼容性：`model` 为空仍默认 `sora-2`；本网关仅暴露 Create/Retrieve/Download 三端点（Sora 与 Veo 同步）。
- Veo 限制：
  - 高级能力仅在 `veo-3.1-*` 可用；1080p 仅支持 8 秒且不支持扩展（扩展仅 720p）。
  - 扩展视频仅接受 Veo 历史成片作为输入；每次固定 +7s，累计不超过 ~141s。
  - 第三方上游（如 sutui）可能不支持参考图/首尾帧/扩展能力。

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

- 2025-11-12：首版发布。打通 `veo-*` 经 `/v1/videos` 的 Create/Retrieve/Download；列表/删除/Remix 未暴露。
