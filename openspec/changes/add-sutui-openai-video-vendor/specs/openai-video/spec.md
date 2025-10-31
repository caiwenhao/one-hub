## ADDED Requirements

### Requirement: OpenAI 视频支持“速推”上游供应商
系统 SHALL 在 OpenAI 视频能力（/v1/videos）中支持新增上游供应商“速推（sutui）”，并保持对外 OpenAI 兼容结构与路由不变。

#### Scenario: 供应商选择（插件配置优先）
- **WHEN** 渠道插件 `channel.plugin.sora.vendor` 配置为 `sutui`
- **THEN** Provider 选择“速推”作为上游供应商

#### Scenario: 供应商选择（BaseURL 自动识别）
- **WHEN** BaseURL 包含 `sutui` 或 `st-ai` 关键字
- **THEN** Provider 选择“速推”作为上游供应商

#### Scenario: 创建视频（JSON 请求）
- **WHEN** 客户端以 JSON 调用 `POST /v1/videos`
- **THEN** 系统 SHALL 将请求映射至“速推”创建接口，并将响应映射为 `types.VideoJob`，其中：
  - `object=video`；
  - 缺失的 `created_at`、`quality` SHALL 被补齐（如 `quality=standard`）；
  - 若上游响应顶层包含 `video_url`，则映射为 `result.video_url`；
  - `status/progress/seconds/size` SHALL 与上游语义一致；
  - 错误响应 SHALL 统一为 OpenAI 错误结构 `{ error: { message, type, code, param } }`。

#### Scenario: 创建视频（multipart 请求直透）
- **WHEN** 客户端以 `multipart/form-data` 调用 `POST /v1/videos`
- **THEN** 系统 SHALL 直透 multipart 请求体至“速推”，并按 JSON 请求的映射规则构造 `types.VideoJob` 响应

#### Scenario: 查询视频
- **WHEN** 客户端调用 `GET /v1/videos/{id}`
- **THEN** 系统 SHALL 查询“速推”并返回 `types.VideoJob`，保证：
  - `status` 为 `queued|in_progress|completed|failed` 之一（必要时进行映射）；
  - `progress` 为 0–100；
  - 若上游返回视频地址，映射为 `result.video_url`；
  - 保留/补齐 `model/seconds/size/created_at`。

#### Scenario: 下载视频内容
- **WHEN** 客户端调用 `GET /v1/videos/{id}/content`
- **THEN** 系统 SHALL 获取 `result.video_url` 并代理/透传视频流；
- **AND** 若 `variant` 非空且非 `video`，返回 501（不支持）。

#### Scenario: Remix（回退策略）
- **WHEN** 客户端调用 `POST /v1/videos/{id}/remix` 且“速推”无原生 remix 接口
- **THEN** 系统 SHALL 回退为通过 create 接口实现：`remix_video_id={id}`，并传递 `prompt`
- **AND** 返回的新任务 `remixed_from_video_id` SHALL 标注来源 ID

#### Scenario: 列表/删除（不强制）
- **WHEN** 客户端调用 `GET /v1/videos` 或 `DELETE /v1/videos/{id}`
- **THEN** 若“速推”不支持，上述接口 SHALL 返回 501 Not Implemented（保持其它供应商不变）

#### Scenario: 计费与配额
- **WHEN** 使用“速推”创建/Remix 视频
- **THEN** 系统 SHALL 按现有 Sora 策略进行配额扣减（以 `seconds` 为基数，换算与现有实现对齐），并在失败时回滚预扣配额

