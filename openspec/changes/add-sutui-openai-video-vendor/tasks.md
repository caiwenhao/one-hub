## 1. Implementation（后端）
- [x] 1.1 Provider 识别：在 `providers/openai/video.go` 的 `detectSoraVendor()` 中新增 `sutui` 分支（优先读取 `channel.plugin.sora.vendor`，其次 BaseURL 包含 `sutui`/`st-ai` 关键字）。
- [x] 1.2 Create 适配：新增 `createSutuiVideo(request)`，支持 multipart 直透与 JSON 映射；将“速推”响应顶层字段适配为 `types.VideoJob`（含 `result.video_url`）。
- [x] 1.3 Retrieve 适配：新增 `retrieveSutuiVideo(videoID)`，映射 `status/progress/seconds/size/video_url` 至标准 `VideoJob`。
- [x] 1.4 Download 适配：新增 `downloadSutuiVideo(videoID, variant)`，基于 `result.video_url` 发起下载并透传；`variant` 非空且非 `video` 返回 501。
- [x] 1.5 Remix 适配：若无上游 remix 接口，回退为 `create(remix_video_id=videoID, prompt=...)`。
- [x] 1.6 错误处理：统一转换为 OpenAI 错误结构（对非 JSON 响应体进行兜底消息封装）。
- [x] 1.7 计费校准：沿用 Sora 秒级计费（参见 `relay/videos.go` 与 `relay/relay_util/quota.go`），补充 `sutui` 分支覆盖率。
- [ ] 1.8 文档：补充 `docs/openai/video.mdx` 用法说明（如何通过渠道插件启用 `sutui`）。

## 2. Implementation（前端）
- [x] 2.1 在渠道编辑弹窗中，OpenAI 上游供应商下拉新增 `sutui（速推）` 选项（`web/src/views/Channel/component/EditModal.jsx`）。
- [x] 2.2 当选择 `sutui` 时：
  - 默认填充 `base_url` 为占位 `https://api.sutui.ai`（可在提交前按实际端点修正）。
  - 将 `custom_parameter.upstream` 同步为 `sutui`。
  - 将插件字段 `plugin.sora.vendor` 设为 `sutui`（便于后端识别）。
- [x] 2.3 当切换回 `official/openrouter/mountsea` 时，清空 `plugin.sora.vendor`，避免误判。
- [x] 2.4 表单校验：OpenAI 且 `upstream in {mountsea, sutui}` 时 `base_url` 必填。
- [x] 2.5 在 `Plugin.json` 为 OpenAI（type=1）新增 `sora.vendor` 字段定义，展示在“插件”区域，支持手工覆盖。
- [ ] 2.6 i18n 提示补充：`zh_CN/en_US/ja_JP` 增加 sutui 说明（可选）。

## 3. Validation
- [ ] 2.1 `openspec validate add-sutui-openai-video-vendor --strict` 通过
- [ ] 2.2 创建视频（JSON/multipart）均返回 `object=video`、`created_at`、`status`，并含 `result.video_url`（完成后可 `GET /v1/videos/{id}/content` 下载）
- [ ] 2.3 查询视频：状态流转 `queued` → `in_progress` → `completed`（或 `failed`）映射正确、`progress` 数值正确
- [ ] 2.4 下载视频：`/v1/videos/{id}/content` 可获取视频内容流；不支持 `variant` 时返回 501
- [ ] 2.5 Remix：返回新任务 `id`，并在 `remixed_from_video_id` 标明来源 ID
- [ ] 2.6 错误映射：上游错误转为 `{ error: { message, type, code, param } }`，HTTP 状态码符合预期

## 4. Rollout
- [ ] 3.1 在测试环境配置一个渠道：`plugin.sora.vendor=sutui` + BaseURL/Key
- [ ] 3.2 回归测试 `official` 与 `mountsea`，确保未回归
- [ ] 3.3 文档更新发布与通知
- [ ] 3.4 观察日志与配额，验证下载代理稳定性
