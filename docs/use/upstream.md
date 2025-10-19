# 上游供应商模式（统一说明）

本节说明 One Hub 中“上游供应商模式”的统一配置方式、优先级与示例。目标是在不增加前端复杂度的前提下，为所有渠道提供可切换的上游供应商能力，并保持接口响应与品牌模型字段对齐。

## 设计与优先级

- 配置位置：渠道的 `custom_parameter`（JSON）。
- 顶层键：`upstream`（默认上游）、`api_key`（如需覆盖渠道主密钥）、`base_url`、`auth_header`、`auth_scheme`、`extra_headers`。
- 通常无需在 JSON 中写 `api_key`，系统默认使用渠道密钥；仅在极端场景（不同能力使用不同密钥）再写入。
- 分能力覆盖：如需更细粒度，可在能力块中声明（例如 `{"video": {"upstream": "ppinfra"}}`）。
- 生效优先级：
  1) 分能力块（如 `video.upstream`）
  2) 顶层 `upstream`
  3) 渠道字段（`base_url` 等）
  4) Provider 默认配置

> 说明：前端“渠道编辑”页在不增加复杂字段的前提下，仅新增一个“上游供应商”下拉（目前先对 MiniMax 启用），其本质是自动读写 `custom_parameter.upstream`。专家用户仍可通过下方“额外参数”JSON 文本框进行分能力覆盖或更细粒度的自定义。

## MiniMax 示例

最简切换（顶层）：

```json
{"upstream": "ppinfra"}
```

分能力覆盖（仅视频走 PPInfra，其他能力默认官方）：

```json
{
  "upstream": "official",
  "video": { "upstream": "ppinfra" }
}
```

如需在顶层补充**不同于渠道密钥**的专用密钥，可同时写入（一般无需配置）：

```json
{
  "upstream": "ppinfra",
  "api_key": "Bearer ppinfra_xxx"  // 留空则仍使用渠道 Key
}
```

生效效果：
- `official`（默认）：`https://api.minimaxi.com`，路径 `/v1/video_generation`、`/v1/query/video_generation`、`/v1/files/retrieve`。
- `ppinfra`：自动切换 `base_url = https://api.ppinfra.com`，使用 `/v3/async/%s`、`/v3/async/task-result` 等路径模板；响应将自动 Normalize 为 MiniMax 品牌字段（`task_id`、`status`、`video_url` 等），调用方无须改动。

## OpenAI 示例（首批：OpenRouter）

最简切换为 OpenRouter：

```json
{"upstream": "openrouter"}
```

生效效果：
- `official`（默认）：`https://api.openai.com`。
- `openrouter`：自动切换 `base_url = https://openrouter.ai/api`，接口路径沿用 OpenAI 兼容路径（如 `/v1/chat/completions`）。
- 可选：若需为 OpenRouter 指定独立密钥，可写 `{"upstream":"openrouter","api_key":"sk-openrouter-xxx"}`；额外请求头（如 `HTTP-Referer`/`X-Title`）可通过“自定义模型请求头（model_headers）”或代理网关配置。

## 其他渠道（逐步开放）

- Zhipu / Deepseek / Groq / VolcArk：第一阶段仅支持通过 `upstream` 切换 `base_url`/常用 `headers`（低成本、风险小）；如存在路径/鉴权/响应差异，将在后续为对应 Provider 增加轻量适配层，继续保持“品牌字段对齐”。

## 前端交互（最简）

- MiniMax 渠道已提供“上游供应商”下拉：默认官方；可选 PPInfra。保存后系统自动写入或移除 `custom_parameter.upstream`。
- 其他渠道会逐步开放该下拉；在开放前，可直接在“额外参数”JSON 文本框中手动写入 `{"upstream":"..."}` 达到相同效果。

## 常见问题

- 与 `base_url` 冲突？按优先级处理：分能力 > 顶层 upstream > 渠道字段 > 默认。若明确要强制覆盖，请使用分能力块（例如 `video.upstream`）。
- 计费/限流差异？Upstream 只影响上游地址/header/路径，系统的用量统计与品牌字段对齐维持不变。若上游策略变化，请在价格表或模型归属页配置。
