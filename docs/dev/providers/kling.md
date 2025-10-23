# 可灵（Kling）集成说明

## 能力概览
- 已对接官方文档的全部视频、图像与多模态编辑接口，支持文本/图像生成视频、多图参考生成、图像生成、多模态编辑全链路。
- 兼容官方与现有 `/kling/v1/:class/:action` 双路由，第三方可零改动切换到 `yourdomain.com/kling/v1`。
- 引入 `external_task_id` 双 ID 查询、分页检索、枚举校验等增强逻辑，保证平台级稳定性。

## 接口一览
| 能力 | 路径（POST/GET） | 说明 |
| --- | --- | --- |
| 文生视频 | `/kling/v1/videos/text2video` | 创建/查询/分页检索 |
| 图生视频 | `/kling/v1/videos/image2video` | 支持静态/动态蒙版 |
| 多图参考生视频 | `/kling/v1/videos/multi-image2video` | 故事化串联 |
| 图像生成 | `/kling/v1/images/generations` | 支持参考、分辨率枚举 |
| 多图参考生图 | `/kling/v1/images/multi-image2image` | 多参考图融合 |
| 多模态编辑 | `/kling/v1/videos/multi-elements/*` | 选区管理、任务提交、查询 |

## 核心代码结构
```
providers/kling/
├── base.go                  // Provider 基础逻辑
├── official_handlers.go     // 视频接口处理
├── official_image_handlers.go // 图像接口处理
├── official_multi_handlers.go // 多模态接口处理
├── official_types.go        // 官方请求/响应结构
├── submit.go                // 任务提交
└── fetch.go                 // 任务查询

model/price.go               // 41 个模型价格配置
model/task.go                // external_task_id 字段
router/relay-router.go       // 官方路由注册
```

## 关键实现要点
- **参数校验**：模型名称、时长、相机控制、笔刷轨迹、互斥字段均进行范围与类型校验，避免上游返回 4xx。
- **双路由兼容**：保留旧式路径同时开放官方路径，方便平滑迁移。
- **分页与任务查询**：统一支持 `pageNum` 1-1000、`pageSize` 1-500，按官方规则透传。
- **价格初始化**：`DefaultKlingPrice` 覆盖 41 个模型组合，启动时可通过价格同步脚本导入。

## 示例调用
```bash
curl -X POST "https://yourdomain.com/kling/v1/videos/text2video" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "model_name": "kling-v1",
    "prompt": "一只可爱的猫在花园里追蝴蝶",
    "mode": "std",
    "duration": "5"
  }'
```

## 验证建议
1. **功能覆盖**：文/图生视频、多模态编辑、图像生成分别验证创建与查询流程。
2. **参数边界**：测试相机控制、动态笔刷、分页极值、非法模型名称等场景。
3. **计费准确性**：检查 `model/price.go` 中关键组合价格是否与官方积分表一致。
4. **回归检查**：旧路由 `/kling/v1/:class/:action` 仍可用，确保历史客户端无感升级。
