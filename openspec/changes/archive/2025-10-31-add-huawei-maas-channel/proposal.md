## Why
为平台引入华为 ModelArts MaaS 官方 OpenAI 兼容能力，补齐国内主流供应商渠道，支持统一的网关转发与运营管理。

## What Changes
- 新增渠道类型常量：`ChannelTypeHuawei = 59`
- 默认品牌归属映射新增：`Huawei MaaS`
- 前端渠道类型选择中新增 `Huawei MaaS`，提供默认 `base_url`、模型示例与表单提示
- 文档：新增/完善 `docs/api/huawei.md`（Chat / 图片生成、参数说明、示例）
- 采用 OpenAI 兼容回退 Provider，对接 `https://api.modelarts-maas.com/v1`

## Impact
- Backend：`common/config/constants.go`、`model/model_ownedby.go`
- Frontend：`web/src/views/Channel/type/Config.js`
- Docs：`docs/api/huawei.md`

## Notes
- 暂不实现独立 Provider 工厂，保持最小可用改动；若后续需要华为特有能力/错误码映射，可追加专用 Provider
- 需在控制台通过 Bearer Token（区域对应的 API Key）鉴权
