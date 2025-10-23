# Vidu 集成与升级说明

## 能力总览
- 覆盖官方 `img2video`、`text2video`、`reference2video`、`start-end2video`、`tasks` 查询及取消接口，新增 `reference2image` 支持。
- 支持最新模型：`viduq2-pro`、`viduq2-turbo`、`viduq1`、`viduq1-classic`、`vidu2.0`、`vidu1.5`，风格、比例、水印等参数与官方保持一致。
- 统一渠道常量、任务平台、价格与模型归属，保持控制台列表、计费、任务日志一致性。
## 接口矩阵
| 能力 | 路径 | 方法 | 备注 |
| --- | --- | --- | --- |
| 图片转视频 | `/vidu/ent/v2/img2video` | POST | 支持多模型/分辨率/运动幅度 |
| 文本转视频 | `/vidu/ent/v2/text2video` | POST | 支持 `style`、`aspect_ratio` |
| 参考生视频 | `/vidu/ent/v2/reference2video` | POST | 支持 1-7 张图片、多分辨率 |
| 首尾帧转视频 | `/vidu/ent/v2/start-end2video` | POST | 支持 2-8 秒 |
| 参考生图 | `/vidu/ent/v2/reference2image` | POST | 新增接口，最多 7 张图片 |
| 查询任务 | `/vidu/ent/v2/tasks`、`/tasks/{task_id}`、`/tasks/{id}/creations` | GET | 官方查询路径 |
| 取消任务 | `/vidu/ent/v2/tasks/{task_id}/cancel` | POST | 官方取消接口 |
## 代码结构
```
providers/vidu/
├── base.go              // Provider 核心逻辑
├── type.go              // 请求/响应结构体
├── submit.go            // 任务提交（含模型名规范化、默认值）
├── fetch.go             // 查询与取消任务实现
relay/task/vidu/         // 任务封装与调度
model/price.go           // 默认价格映射
router/relay-router.go   // 官方路由注册
```
## 核心实现要点
- **模型名规范化**：`normalizeModelName` 同时兼容 `vidu-1.5`/`vidu1.5` 等旧格式，确保老客户端可用。
- **价格生成**：`buildDetailedModelName` 按 action/模型/时长/分辨率/风格组合，默认价格与官方积分表对齐。
- **参数校验**：涵盖图片数量、分辨率、运动幅度、水印、错峰等参数，并在任务提交阶段返回统一错误。
- **任务管理**：支持 created/queueing/processing/success/failed 状态，缓存查询结果并支持取消任务。
## 示例调用
```bash
curl -X POST "https://your-domain.com/vidu/ent/v2/img2video" \
  -H "Authorization: Token <api_key>" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "viduq2-pro",
    "images": ["https://example.com/image.png"],
    "duration": 5,
    "resolution": "1080p",
    "movement_amplitude": "auto"
  }'
```
## 验证建议
1. **功能测试**：按 action 覆盖四类视频生成功能＋参考生图，确认参数透传与状态查询无误。
2. **计费校验**：在后台 "模型价格" 页面触发价格同步，确认生成键值与 `DefaultViduPrice` 一致。
3. **错误场景**：验证非法模型、图片数量超限、错峰与回调配置、取消任务失败等边界情况。
4. **回归检查**：历史路由与任务平台标识（`TaskPlatformVidu`）保持不变，确保现有业务无回归。
