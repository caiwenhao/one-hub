---
title: Pollo · Vidu 兼容上游（渠道对接规范）
description: Pollo.ai 上游在本系统中需完全兼容 Vidu 官方接口路由、请求体与响应体，保证对用户一致的调用体验。
---

使用场景

- 当渠道上游选择 Pollo.ai 时，网关对外仍提供与 Vidu 官方完全一致的接口：路径、请求体、响应体、状态语义均保持一致。
- 本文为“渠道上游供应商（Pollo）对接规范”，用于明确兼容要求与差异点（仅鉴权与基础域名不同）。

路由与方法（保持与 Vidu 官方一致）

- 提交任务：`POST /vidu/ent/v2/{action}`
  - `action ∈ { img2video | reference2video | start-end2video | text2video | reference2image }`
- 查询结果（creations）：`GET /vidu/ent/v2/tasks/{task_id}/creations`
- 取消任务：`POST /vidu/ent/v2/tasks/{task_id}/cancel`

请求头（差异点仅在鉴权）

- 官方 Vidu：`Authorization: Token <key>`
- Pollo 上游：`x-api-key: <key>`
- 其他头保持一致：`Content-Type: application/json`

请求体与响应体（完全与 Vidu 官方一致）

- 字段、命名与含义：请参考 docs/api/vidu.md；上游必须按 Vidu 官方规范解析请求体并返回相同结构。
- 典型请求字段：`model, images, prompt, duration, seed, resolution, movement_amplitude, bgm, payload, off_peak, watermark, wm_position, wm_url, style, aspect_ratio, callback_url, meta_data`。
- 典型响应字段：提交 `{ task_id, state, ... }`；查询 `{ id, state, credits, payload, bgm, off_peak, creations: [{ id, url, cover_url, watermarked_url }] }`；取消 `{}`。
- 状态枚举：`created | queueing | processing | success | failed`。

模型支持（Pollo 覆盖范围）

- 基础模型名遵循 Vidu 官方：
  - 已要求支持：`viduq2-pro`、`viduq2-turbo`、`viduq1`
  - 可逐步扩展：`viduq1-classic`、`vidu2.0`、`vidu1.5`（若上游完成适配后，可在后台开启）
- 注意：客户端仅填写基础模型名；网关会按 action/时长/分辨率/风格/错峰自动映射精确计费组合，Pollo 上游需据此生成一致的结果。

合法性约束（与 Vidu 官方一致）

- `viduq1 / viduq1-classic`：固定 5s 且 1080p。
- `vidu2.0 / vidu1.5`：8s 强制 720p；4s 可 360p/720p/1080p。
- `viduq2-pro / viduq2-turbo`：1–8s，540p/720p/1080p。
- `text2video` 的 Q2 家族：上游可接受 `viduq2` 或等价归一名（网关会按需归一以保证兼容）。

错峰与计费（对齐 Vidu 官方）

- `off_peak=true` 表示错峰，价格为正常价的一半（若出现小数向上取整）；Pollo 上游需遵循与 Vidu 相同的计费/状态语义。
- 失败不计费；回调 `callback_url` 的返回体与“查询任务”一致。

示例（Pollo 基础域名 + Vidu 标准路由与体）

提交（text2video）：

```bash
curl -X POST "https://pollo.ai/api/platform/vidu/ent/v2/text2video" \
  -H "x-api-key: <POLLO_KEY>" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "viduq2-pro",
    "prompt": "A cinematic sunset over the ocean, slow dolly shot.",
    "style": "general",
    "duration": 4,
    "resolution": "720p",
    "off_peak": false,
    "callback_url": "https://example.com/callback"
  }'
```

查询（creations）：

```bash
curl "https://pollo.ai/api/platform/vidu/ent/v2/tasks/$TASK_ID/creations" \
  -H "x-api-key: <POLLO_KEY>"
```

取消：

```bash
curl -X POST "https://pollo.ai/api/platform/vidu/ent/v2/tasks/$TASK_ID/cancel" \
  -H "x-api-key: <POLLO_KEY>"
```

回调（Callback）

- `callback_url` 的 POST 回调体与“查询任务”返回体一致；状态推进：`created → queueing → processing → success/failed`。

备注

- 若 Pollo 实际提供的兼容基础域名不同于 `https://pollo.ai/api/platform`，以实际部署为准；只要路由与体保持 Vidu 标准即可。
- 若上游暂不支持取消，可先返回空体 200 并落库状态为 `failed/cancelled`，或返回标准错误；最终目标仍是对齐 Vidu 官方行为。


Vidu Q2 Turbo API Documentation
Vidu Q2 Turbo API delivers fast motion-heavy videos with stable camera transitions. Ideal for quick short-form content, it balances speed and quality effectively. Integrate it now.

POST
/
generation
/
vidu
/
viduq2-turbo

Try it
Check out Vidu Q2 Turbo API pricing here, cheaper than Fal AI and Replicate.
Authorizations
​
x-api-key
stringheaderrequired
API key to authorize requests

Body
application/json
​
input
objectrequired
Hide child attributes

​
input.image
stringrequired
Only image URLs are accepted (HTTPS preferred); base64 is not allowed.
Supported formats include JPG, PNG, and JPEG.
Image aspect ratio must be less than 1:4 or 4:1.

​
input.prompt
string
The prompt of the generation

Maximum length: 2000
​
input.movementAmplitude
enum<string>default:auto
Available options: auto, small, medium, large 
​
input.length
enum<number>default:5
The length of the generation

Available options: 1, 2, 3, 4, 5, 6, 7, 8 
​
input.resolution
enum<string>default:720p
The resolution of the generation

Available options: 540p, 720p, 1080p 
​
input.seed
number
The seed of the generation

​
input.generateAudio
boolean
Generate natural-sounding, fitting audio for the output video.

​
webhookUrl
string
Response

200

application/json
Successful response

​
taskId
stringrequired
​
status
enum<string>required
Available options: waiting, succeed, failed, processing 

Vidu Q1 API Documentation
Vidu Q1 is a next-gen AI video generator released in 2025 that creates high-quality, realistic 1080p videos from text prompts or images. It features smooth motion, cinematic lighting, and detailed animations. Learn how to integrate it below.

POST
/
generation
/
vidu
/
vidu-q1

Try it
Check out Vidu Q1 API pricing here, cheaper than Fal AI and Replicate.
Authorizations
​
x-api-key
stringheaderrequired
API key to authorize requests

Body
application/json
​
input
objectrequired
Image To Video
Text To Video
Hide child attributes

​
input.image
stringrequired
Only image URLs are accepted (HTTPS preferred); base64 is not allowed.
Supported formats include JPG, PNG, and JPEG.
Image aspect ratio must be less than 1:4 or 4:1.

​
input.prompt
string
The prompt of the generation

Maximum length: 2500
​
input.imageTail
string
URL of the tail image for video generation.
Only image URLs are accepted (HTTPS preferred); base64 is not allowed.
Supported formats include JPG, PNG, and JPEG.
Image aspect ratio must be less than 1:4 or 4:1.

​
input.movementAmplitude
enum<string>default:auto
Available options: auto, small, medium, large 
​
input.length
enum<number>default:5
The length of the generation

Available options: 5 
​
input.resolution
enum<string>default:1080p
The resolution of the generation

Available options: 1080p 
​
input.seed
number
The seed of the generation

​
webhookUrl
string
Response

200

application/json
Successful response

​
taskId
stringrequired
​
status
enum<string>required
Available options: waiting, succeed, failed, processing 

Vidu 2.0 API Documentation
Vidu 2.0 is an AI video generation API that creates videos by combining reference images with text prompts. It uses advanced technology to keep characters, objects, and environments consistent throughout the video, ensuring smooth and natural animations. Learn how to integrate it below.

POST
/
generation
/
vidu
/
vidu-v2-0

Try it
Check out Vidu 2.0 API pricing here, cheaper than Fal AI and Replicate.
Authorizations
​
x-api-key
stringheaderrequired
API key to authorize requests

Body
application/json
​
input
objectrequired
Hide child attributes

​
input.image
stringrequired
Only image URLs are accepted (HTTPS preferred); base64 is not allowed.
Supported formats include JPG, PNG, and JPEG.
Image aspect ratio must be less than 1:4 or 4:1.

​
input.prompt
string
The prompt of the generation

Maximum length: 2500
​
input.length
enum<number>default:4
The length of the generation

Available options: 4, 8 
​
input.resolution

Option 1 · enum<string>
default:720p
The resolution of the generation

Available options: 720p 
​
input.imageTail
string
URL of the tail image for video generation.
Only image URLs are accepted (HTTPS preferred); base64 is not allowed.
Supported formats include JPG, PNG, and JPEG.
Image aspect ratio must be less than 1:4 or 4:1.

​
input.movementAmplitude
enum<string>default:auto
Available options: auto, small, medium, large 
​
input.seed
number
The seed of the generation

​
webhookUrl
string
Response

200

application/json
Successful response

​
taskId
stringrequired
​
status
enum<string>required
Available options: waiting, succeed, failed, processing 

Vidu 1.5 API Documentation
Vidu 1.5 is a powerful AI video generation model with high prompt adherence, featuring smooth transitions and creative effects. Learn how to integrate it below.

POST
/
generation
/
vidu
/
vidu-v1-5

Try it
Check out Vidu 1.5 API pricing here, cheaper than Fal AI and Replicate.
Authorizations
​
x-api-key
stringheaderrequired
API key to authorize requests

Body
application/json
​
input
objectrequired
Image To Video
Text To Video
Hide child attributes

​
input.image
stringrequired
Only image URLs are accepted (HTTPS preferred); base64 is not allowed.
Supported formats include JPG, PNG, and JPEG.
Image aspect ratio must be less than 1:4 or 4:1.

​
input.prompt
string
The prompt of the generation

Maximum length: 2500
​
input.imageTail
string
URL of the tail image for video generation.
Only image URLs are accepted (HTTPS preferred); base64 is not allowed.
Supported formats include JPG, PNG, and JPEG.
Image aspect ratio must be less than 1:4 or 4:1.

​
input.length
enum<number>default:4
The length of the generation

Available options: 4, 8 
​
input.resolution
enum<string>default:360p
The resolution of the generation

Available options: 360p, 720p, 1080p 
​
input.movementAmplitude
enum<string>default:auto
Available options: auto, small, medium, large 
​
input.seed
number
The seed of the generation

​
webhookUrl
string
Response

200

application/json
Successful response

​
taskId
stringrequired
​
status
enum<string>required
Available options: waiting, succeed, failed, processing 

Vidu 1.5 API Documentation
Vidu 1.5 is a powerful AI video generation model with high prompt adherence, featuring smooth transitions and creative effects. Learn how to integrate it below.

POST
/
generation
/
vidu
/
vidu-v1-5

Try it
Check out Vidu 1.5 API pricing here, cheaper than Fal AI and Replicate.
Authorizations
​
x-api-key
stringheaderrequired
API key to authorize requests

Body
application/json
​
input
objectrequired
Image To Video
Text To Video
Hide child attributes

​
input.prompt
stringrequired
The prompt of the generation

Required string length: 1 - 2500
​
input.style
enum<string>default:general
Available options: general, anime 
​
input.length
enum<number>default:4
The length of the generation

Available options: 4, 8 
​
input.resolution
enum<string>default:360p
The resolution of the generation

Available options: 360p, 720p, 1080p 
​
input.movementAmplitude
enum<string>default:auto
Available options: auto, small, medium, large 
​
input.aspectRatio
enum<string>default:16:9
Available options: 16:9, 9:16, 1:1 
​
input.seed
number
The seed of the generation

​
webhookUrl
string
Response

200

application/json
Successful response

​
taskId
stringrequired
​
status
enum<string>required
Available options: waiting, succeed, failed, processing 
