## ADDED Requirements

### Requirement: Gemini Veo 支持“速推（sutui）”上游
系统 SHALL 在 Gemini Veo 能力（`/gemini/{version}/models/veo-*:predictLongRunning` + `/gemini/{version}/operations/{name}`）中支持“速推（sutui/st-ai）”作为上游供应商；对外保持 Gemini 原生路由与响应风格不变。

#### Scenario: 上游选择（插件配置优先）
- **WHEN** 渠道插件 `plugin.gemini.video.vendor` 配置为 `sutui`
- **THEN** 系统 SHALL 选用“速推”作为上游供应商

#### Scenario: 上游选择（BaseURL 自动识别）
- **WHEN** 渠道 BaseURL 包含 `sutui` 或 `st-ai`
- **THEN** 系统 SHALL 选用“速推”作为上游供应商

#### Scenario: 任务创建（JSON 请求）
- **WHEN** 客户端以 JSON 调用 `POST /gemini/{version}/models/veo-*:predictLongRunning`
- **THEN** 系统 SHALL 将请求参数映射至 `POST /v1/videos`（sutui 上游），字段规则：
  - 文本提示 `prompt` → sutui `prompt`
  - 时长 `durationSeconds`（如存在）→ sutui `seconds`
  - 画幅 `aspectRatio`（如存在）→ sutui `size`（按 16:9/9:16 映射到 `1600x900`/`900x1600` 等常见分辨率）
  - 其他字段保持透传或忽略（不影响核心流程）
- **AND** 系统 SHALL 以 sutui 返回的 `id` 封装 `operations` 资源名（如 `operations/{id}`）并返回 200

#### Scenario: 任务创建（multipart 请求直透）
- **WHEN** 客户端以 `multipart/form-data` 调用 `POST /gemini/{version}/models/veo-*:predictLongRunning`
- **THEN** 系统 SHALL 直透 multipart 请求至 sutui `POST /v1/videos`；允许字段：
  - `model`、`prompt`、`size`、`input_reference`（单/多张）等（参考 `docs/api/st-ai-veo.md`）
- **AND** 同样以 sutui `id` 组装并返回 `operations` 资源名

#### Scenario: 任务轮询
- **WHEN** 客户端调用 `GET /gemini/{version}/operations/{name}`
- **THEN** 系统 SHALL 将 `{name}` 中的任务 ID 映射到 sutui `GET /v1/videos/{id}` 查询，并转换为 Gemini `operations` 响应：
  - `done`: 依据 sutui `status in {completed, failed}` 判定
  - `response.generateVideoResponse.generatedSamples[]`: 
    - `uri` = sutui `result.video_url`（或顶层 `video_url`）
    - `metadata.durationSeconds` = sutui `seconds`
    - 可选 `metadata.size` = sutui `size`

#### Scenario: 错误映射
- **WHEN** sutui 返回错误
- **THEN** 系统 SHALL 转换为 Gemini operations 错误语义（HTTP 非 2xx 时以 JSON 错误体返回），同时保持 `done=true` 与错误状态，或直接返回网关错误（可配置）

#### Scenario: 计费
- **WHEN** 通过 sutui 上游生成/轮询 Veo 任务
- **THEN** 系统 SHALL 以视频时长（秒）计费，内部按 `seconds × 1000` 折算，与现有 Veo/Sora 策略一致

## MODIFIED Requirements

### Requirement: Gemini Veo 任务创建与轮询（原生 Google）
系统 SHALL 支持官方 Google Veo 路由（`:predictLongRunning` + `/operations/*`）。

#### Scenario: 上游选择（默认 Google）
- **WHEN** 未命中 `plugin.gemini.video.vendor=sutui` 且 BaseURL 不包含 `sutui`/`st-ai`
- **THEN** 系统 SHALL 使用 Google 官方端点与密钥

