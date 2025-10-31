## 1. Implementation
- [x] 1.1 新增渠道类型常量（59）
- [x] 1.2 默认品牌归属映射新增 Huawei MaaS
- [x] 1.3 前端渠道类型配置补充（默认 base_url、模型示例、提示文案）
- [x] 1.4 文档补充 `docs/api/huawei.md`
- [x] 1.5 Chat 扩展参数透传（chat_template_kwargs / use_beam_search / length_penalty / best_of / prompt_logprobs）
- [x] 1.6 stream_options 透传策略（Huawei 默认开启 usage，客户端可显式控制）
- [ ] 1.7 控制台手测（非流式/流式、图片生成）并截图

## 2. Validation
- [ ] 2.1 `openspec validate add-huawei-maas-channel --strict` 通过
- [ ] 2.2 代码 Lint 通过（后端 golangci-lint / 前端 eslint+prettier）
- [ ] 2.3 PR 模板项勾选与自测截图

## 3. Rollout
- [ ] 3.1 Release Note 增加“新增 Huawei MaaS 渠道”
- [ ] 3.2 文档站点导航（可选）：增加 Huawei 专区入口
