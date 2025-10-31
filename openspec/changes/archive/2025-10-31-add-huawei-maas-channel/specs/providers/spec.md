## ADDED Requirements
### Requirement: Huawei MaaS Channel
系统 MUST 支持接入华为 ModelArts MaaS 作为新的渠道类型，ID 固定为 59，使用 OpenAI 兼容接口对接（自定义 `base_url` + Bearer 鉴权）。

#### Scenario: 渠道创建成功
- WHEN 管理员在控制台选择渠道类型 `Huawei MaaS` 并填写 `base_url=https://api.modelarts-maas.com/v1` 与有效的 API Key
- THEN 渠道保存成功并可在列表中查看到新渠道

#### Scenario: Chat 非流式应答
- WHEN 通过该渠道调用 `/v1/chat/completions`（非流式）发送文本消息
- THEN 返回 `200` 且 `choices[0].message.content` 非空，`model` 为配置的华为模型（如 `deepseek-v3`）

#### Scenario: Chat 流式应答含用量
- WHEN 通过该渠道以 `stream=true` 并设置 `stream_options={"include_usage":true}` 调用
- THEN 分片返回中包含 `usage` 字段（累计 tokens 统计）

#### Scenario: 前端类型可见
- WHEN 进入“新建渠道”页面
- THEN 渠道类型下拉中可选择 `Huawei MaaS`，并展示默认 `base_url` 与示例模型

#### Scenario: 文档可用
- WHEN 查看 `docs/api/huawei.md`
- THEN 能获得 Chat/图片生成/参数说明与多语言示例，指导用户正确接入
