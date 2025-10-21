# MiniMax 渠道语音接口审查报告

## 范围说明
- 本轮审查**仅聚焦 `/minimaxi` 官方兼容路由**（例如 `/minimaxi/v1/t2a_v2`、`/minimaxi/v1/t2a_async_v2`），暂不处理 OpenAI 兼容入口及其降级策略。
- 目标：核对路由完备性、请求/响应结构是否对齐官方 OpenAPI，并给出完善 `/minimaxi` 侧能力的实施方案。

## 项目结构与技术栈概览
- 后端为 Go 单体服务，基于 Gin 构建 REST API（`router/`、`middleware/`），`relay/` 负责统一调度，`providers/` 根据渠道封装业务逻辑。
- MiniMax 渠道位于 `providers/minimaxi/`，`MiniMaxRelay` 在 `relay/` 中处理 `/minimaxi` 前缀转发，长任务轮询逻辑在 `relay/task/minimax/`。
- 公用组件包括 `common/requester`（HTTP 封装）、`common/config`（路由常量与渠道信息）、`types/`（跨渠道的结构体定义）。

## 审查结论速览
- ✅ `/minimaxi/v1/t2a_v2`（同步语音）已实现，支持基本音频生成，但响应结构与官方规范存在字段缺失/命名差异，subtitle/url 等高级能力尚未覆盖。
- ✅ `/minimaxi/v1/t2a_async_v2`（创建异步任务）可用，能返回 `task_id` 等信息，但类型定义与上游返回仍存在潜在兼容性问题。
- ❌ `/minimaxi/v1/query/t2a_async_query_v2`（异步查询）**未落地**：路由缺失、Provider 无对应方法，导致异步任务无法闭环。

## 详解

### `/minimaxi/v1/t2a_v2` 同步语音合成
- 路径：`router/relay-router.go:166-181` 使用 `relay.MiniMaxRelay` 直接转发至 Provider。
- Provider：`MiniMaxProvider.CreateSpeechOfficial`（`providers/minimaxi/speech.go:291-318`）读取原始 JSON 后调用上游。

#### 差异点
- 响应结构 `SpeechResponse` 定义在 `providers/minimaxi/type.go:39-58`，缺少官方文档中的 `subtitle_file`、`audio_channel`；`bitrate` 字段被命名为 `audio_bitrate`。
- `CreateSpeechOfficial` 在非流场景复用 `CreateSpeech` 的逻辑，但 `CreateSpeech`（`providers/minimaxi/speech.go:158-205`）默认对 `data.audio` 进行 hex 解码：当官方返回 URL 或字幕信息时会触发 `hex.DecodeString` 错误，或导致附加信息丢失。
- 目前未读取 `subtitle_enable`、`output_format` 等开关，即使请求体携带也不会影响处理流程。

#### 影响
- 官方场景（字幕/URL 输出）无法成功落地 `/minimaxi` 路径；响应统计数据不完整，影响监控和调试。

### `/minimaxi/v1/t2a_async_v2` 异步任务创建
- 路径：`router/relay-router.go:176-181`。
- Relay：`relay/minimaxi_async.go`，自动解析 `MiniMaxAsyncSpeechRequest`（`types/audio.go:70-85`）并透传 JSON。
- Provider：`MiniMaxProvider.CreateSpeechAsync`（`providers/minimaxi/speech.go:207-255`），在发送前会做语音别名替换。

#### 差异点
- `MiniMaxAsyncSpeechResponse`（`types/audio.go:103-116`）的 `task_id` 定义为 string，而官方示例使用纯数字；Go JSON 解析时若返回整数将报错。
- 请求层未显式校验 `voice_setting` 等必填字段，错误信息将由上游兜底，缺乏前置校验提示。

#### 影响
- 当上游返回数字 task_id/file_id 时可能触发解析失败；缺乏输入校验会影响错误定位。

### `/minimaxi/v1/query/t2a_async_query_v2` 异步任务查询
- 当前代码无路由注册，也无 Provider 调用实现。
- `relay/task/minimax/fetch.go` 仅实现视频任务相关的查询，与语音无直接共享逻辑。

#### 影响
- 平台无法在 `/minimaxi` 路径下查询语音异步任务状态，业务需要自行调用官方接口，破坏统一入口设计。

## 主要风险
- **功能缺口**：缺失查询接口导致异步任务不可闭环。
- **稳定性**：同步接口在字幕/URL 场景会直接 500，影响渠道可用性与用户体验。
- **数据一致性**：响应字段缺失导致监控、账务统计信息不完整。

## 设计方案对比
1. **方案 A：全面对齐官方 JSON**  
   - `/minimaxi` 路由始终透传官方 JSON，不再尝试自动解码音频；完善请求/响应结构；新增查询接口。  
   - 优点：逻辑简单，对应官方文档；无解码失败风险。  
   - 缺点：调用方需自行处理 hex/URL 音频，若平台需要统一输出裸音频需额外适配。
2. **方案 B：按配置选择处理模式（推荐）**  
   - 默认按照官方 JSON 透传；若渠道配置要求输出裸音频，可在 `MiniMaxProvider` 中打开解码模式（仅支持 hex）。  
   - 同时补齐响应结构、查询接口，避免字段缺失。  
   - 优点：兼顾官方一致性与平台已有音频分发逻辑，可渐进式迁移。  
   - 缺点：需要新增配置项和分支处理。

## 推荐方案
- 采用 **方案 B**：以官方 JSON 为基线，增加渠道配置 `audio_mode`（`json`|`hex`），默认 `json`。仅在 `hex` 模式下执行音频解码，并显式限制 `output_format` 与 `subtitle_enable` 等参数。
- 同步完善响应结构与异步查询能力，确保 `/minimaxi` 路由覆盖官方列出的全部接口。

## 实施任务清单
1. **补齐路由与 Query 能力（已完成）**
   - 在 `router/relay-router.go` 注册 `/minimaxi/v1/query/t2a_async_query_v2`（GET）。
   - 新增 `relay/minimaxi_async_query.go`，负责 `task_id` 解析、按分组筛选 MiniMax 渠道，并调用 Provider 查询。
   - 在 `providers/minimaxi/` 实现 `QuerySpeechAsync`，并为 `AudioSpeechAsyncQuery` 赋默认路径 ` /v1/query/t2a_async_query_v2`。
2. **调整同步接口结构（已完成）**
   - 更新 `providers/minimaxi/type.go`，补充 `subtitle_file` / `audio_channel`，并将 `audio_bitrate` 更名为 `bitrate`，所有统计字段使用指针以容错。
   - `CreateSpeech` / `CreateSpeechOfficial` 根据数据动态设置头部，默认透传 JSON，若渠道配置 `audio_mode=hex` 则解码音频并输出二进制流。
3. **增强异步请求/响应兼容（已完成）**
   - `MiniMaxAsyncSpeechResponse`、`MiniMaxAsyncSpeechQueryResponse`、`MiniMaxFileObject` 采用 `types.StringOrNumber`，兼容数值型 ID。
   - 异步提交入口新增 `voice_setting.voice_id` 校验，避免缺失必填参数。
4. **文档与测试（进行中）**
   - 需同步更新对外文档与控制台文案，说明 `audio_mode` 配置、字幕/声道头部等变化。
   - 建议补充端到端回归用例：创建→查询异步任务、同步字幕输出、`audio_mode` 切换等。

## 实现说明
- `/minimaxi/v1/query/t2a_async_query_v2` 已支持按 token 分组自动遍历可用渠道，若上游返回错误则透传最后一次错误信息。
- 渠道 `custom_parameter` 可通过 `{"audio_mode":"hex"}` 或 `{"audio":{"mode":"hex"}}` 控制输出模式，默认 `json`；`hex` 模式下自动校验并记录解码失败日志。
- 无论官方还是兼容入口，返回流中附带 `X-Minimax-Subtitle-URL`、`X-Minimax-Audio-Channel` 等头部，方便调用方获取字幕和声道信息。
- 异步响应与文件对象在 ID 解析上具备更强的数值兼容性，避免因上游返回整型导致的反序列化失败。

## 验证建议
- 使用官方示例请求：分别触发 `hex`、`url`、开启字幕三种返回类型，确认 `/minimaxi` 路径能正确返回或降级。
- 提交异步任务后，通过新路由轮询直至 `success/failed/expired`，核对 `status` 与 `base_resp`。
- 切换渠道 `audio_mode` 配置，验证裸音频解码模式下响应头 (`Content-Type`, `Content-Length`) 是否准确，且错误日志清晰。
- 回归现有 MiniMax 视频任务逻辑，确保新增语音查询能力不会影响视频相关接口。
