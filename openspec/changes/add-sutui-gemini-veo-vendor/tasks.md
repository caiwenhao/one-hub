## 1. Implementation

 - [x] 1.1 上游选择逻辑（Gemini Provider）
  - 在 `providers/gemini` 增加 sutui 检测：优先读 `plugin.gemini.video.vendor`，其次 BaseURL 关键字。
 - [x] 1.2 Veo 初始化映射（predictLongRunning → sutui create）
  - JSON：将 `prompt/durationSeconds/aspectRatio` 映射为 sutui `prompt/seconds/size`；构造 `POST /v1/videos` 请求。
  - multipart：直透请求体与 Content-Type 到 sutui。
  - 返回：封装 `operations/{id}` 为 Gemini 响应。
 - [x] 1.3 Operations 轮询映射
  - 解析 `{name}` 获取任务 ID，调用 sutui `GET /v1/videos/{id}`。
  - 将 `status/seconds/size/video_url` 转换到 Gemini operations：`done` 与 `generatedSamples[]`。
 - [x] 1.4 错误转换与健壮性
  - 统一非 2xx 错误格式；为缺失字段做兼容（如 `progress`）。
 - [x] 1.5 计费
  - 复用现有 Veo 秒级计费：从 sutui 查询结果或请求参数推断 `seconds`，写入 `video_seconds` 元数据。
- [ ] 1.6 文档
  - 在 `docs/gemini/video.mdx` 增补“sutui 作为上游”的使用说明与示例。

## 2. Validation

- [ ] 2.1 `openspec validate add-sutui-gemini-veo-vendor --strict` 通过
- [ ] 2.2 JSON/multipart 两种创建方式均返回有效 `operations/{id}`
- [ ] 2.3 轮询返回 `done` 与 `generatedSamples[].uri` 可下载
- [ ] 2.4 秒级计费：`video_seconds` 元数据正确（8s 示例）
- [ ] 2.5 错误映射：sutui 错误被正确转换为 Gemini 语义

## 3. Rollout

- [ ] 3.1 新建 Gemini 渠道：`plugin.gemini.video.vendor=sutui` + BaseURL/Key
- [ ] 3.2 验证官方 Google 路由未回归
- [ ] 3.3 发布文档与渠道配置指引
- [ ] 3.4 观察日志/用量，验证稳定性

## 4. Frontend

 - [x] F1 在前端渠道“插件”面板新增 Gemini 视频上游配置
  - 目标文件：`web/src/views/Channel/type/Plugin.json`
  - 在键 `"25"`（Gemini 渠道类型）下新增一个插件分组（建议 key：`gemini_video` 或 `video`），示例：
    - `name`: "Gemini 视频上游"
    - `description`: "Veo 视频上游供应商选择（google/sutui）。用于 /gemini/{version}/models/veo-*"
    - `params.vendor`: `string`，可填 `google`（默认）或 `sutui`
- [ ] F2 编辑弹窗适配 `plugin.gemini.video.vendor`
  - 目标文件：`web/src/views/Channel/component/EditModal.jsx`
  - 当渠道类型为 25（Gemini）时，渲染上述插件分组；读取/写回表单到字段 `plugin.gemini.video.vendor`
  - 若选择 `sutui`：在 UI 提示需确保 sutui 的 BaseURL/Key 已配置（或沿用系统默认）
- [ ] F3 i18n（可选）
  - 为“Gemini 视频上游”“视频上游供应商”等增加多语言文案（`web/src/i18n/locales/*.json`）
- [ ] F4 回归
  - 验证切换 `vendor=google/sutui` 后保存成功；再次打开能够正确回显
