# Vidu AI 视频生成支持集成文档

## 概述

本次开发成功为 one-hub 项目添加了对 Vidu AI 视频生成服务的完整支持。Vidu 是一款先进的AI视频生成平台，支持图片转视频、文本转视频、参考视频生成和起止图转视频等多种功能。

## 实现功能

### 1. 核心Provider实现
- **文件位置**: `providers/vidu/`
- **类型定义**: `type.go` - 完整的API请求响应结构体
- **基础实现**: `base.go` - Vidu Provider核心逻辑
- **支持接口**:
  - `img2video` - 图片转视频
  - `text2video` - 文本转视频  
  - `reference2video` - 参考视频生成
  - `start-end2video` - 起止图转视频

### 2. 异步任务处理
- **文件位置**: `relay/task/vidu/`
- **任务提交**: `submit.go` - 处理任务提交和验证
- **任务查询**: `fetch.go` - 支持单个任务和批量任务查询
- **任务状态管理**: 支持created/queueing/processing/success/failed状态跟踪

### 3. 系统集成
- **常量定义**: 添加 `ChannelTypeVidu = 57` 和 `RelayModeVidu`
- **路由配置**: 支持 `/vidu/ent/v2/` 路由组
- **工厂注册**: 在 providers.go 中注册 ViduProviderFactory
- **前端支持**: 添加渠道选项和图标配置

### 4. 数据库和模型
- **平台支持**: 添加 `TaskPlatformVidu = "vidu"`
- **价格配置**: 为不同模型和时长配置默认价格
- **模型归属**: 添加Vidu品牌标识和图标

## API端点

### 任务提交
```
POST /vidu/ent/v2/{action}
```
支持的action:
- `img2video` - 图片转视频
- `text2video` - 文本转视频
- `reference2video` - 参考视频生成  
- `start-end2video` - 起止图转视频

### 任务查询
```
GET /vidu/ent/v2/task/{task_id}    # 单个任务查询
GET /vidu/ent/v2/tasks             # 批量任务查询
```

## 请求参数

### 通用参数
- `model`: 模型名称，根据官方文档支持 (viduq1, vidu1.5, vidu2.0, viduq1-classic)
- `prompt`: 文本描述 (最长2000字符)
- `duration`: 视频时长 (4s/5s/8s)
- `seed`: 随机种子
- `resolution`: 分辨率 (1080p/720p/360p)
- `movement_amplitude`: 运动幅度 (auto/small/medium/large)
- `bgm`: 是否添加背景音乐
- `off_peak`: 离峰模式
- `callback_url`: 回调地址

### 特定参数
- **img2video**: `images` (图片URL数组)
- **text2video**: `prompt` (必填)
- **reference2video**: `reference_videos` (参考视频URL数组)
- **start-end2video**: `start_image`, `end_image` (起始和终止图片)

## 认证机制

使用Token认证方式:
```
Authorization: Token {your_api_key}
```

## 价格配置

按照不同模型和时长设置了差异化价格:
- viduq1-classic: 30-100积分
- viduq1: 40-120积分  
- vidu1.5: 50-180积分
- vidu2.0: 60-220积分

## 状态管理

基于内存中的任务状态管理机制，支持：
- 任务状态轮询查询
- 进度跟踪 (10%-100%)
- 错误处理和重试机制
- 结果数据缓存

## 错误处理

完整的错误处理机制:
- 参数验证错误
- 认证错误
- 服务不可用错误
- 任务状态查询错误
- 网络超时错误

## 扩展性

架构设计支持：
- 新模型快速接入
- 新接口类型扩展
- 自定义参数配置
- 灵活的价格策略

## 使用示例

### 图片转视频
```bash
curl -X POST "https://your-domain.com/vidu/ent/v2/img2video" \
  -H "Authorization: Token your_api_key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu1.5",
    "images": ["https://example.com/image.jpg"],
    "prompt": "让这张图片动起来",
    "duration": 5
  }'
```

### 文本转视频
```bash
curl -X POST "https://your-domain.com/vidu/ent/v2/text2video" \
  -H "Authorization: Token your_api_key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-2.0", 
    "prompt": "一只可爱的小猫在花园里玩耍",
    "duration": 8,
    "resolution": "1080p"
  }'
```

### 查询任务状态
```bash
curl -X GET "https://your-domain.com/vidu/ent/v2/task/{task_id}" \
  -H "Authorization: Token your_api_key"
```

## 技术特点

1. **异步处理**: 基于任务队列的异步视频生成
2. **状态轮询**: 支持实时状态查询和进度跟踪
3. **多模型支持**: 支持Vidu全系列模型
4. **灵活配置**: 支持多种视频参数配置
5. **错误恢复**: 完善的重试和错误处理机制
6. **性能优化**: 支持批量查询和缓存机制

## 总结

此次集成为 one-hub 平台新增了强大的AI视频生成能力，完全兼容Vidu官方API，支持四种主要的视频生成模式，提供了完整的任务管理和状态跟踪功能。用户可以通过统一的API接口访问Vidu的所有视频生成功能，享受一站式的AI服务体验。

## 补充：模型名规范化兼容

为便于与官方 API 接口对齐，系统在向上游 Vidu 提交请求时，会保持官方文档规定的模型名称格式：

- `viduq1` → `viduq1`（官方文档确认格式）
- `vidu1.5` → `vidu1.5`（官方文档确认格式）
- `vidu2.0` → `vidu2.0`（保持原格式）
- `viduq1-classic` → `viduq1-classic`（保持原格式）

向后兼容连字符格式：
- `vidu-1.5` → `vidu1.5`
- `vidu-2.0` → `vidu2.0`
- `vidu-q1` → `viduq1`
- `vidu-q1-classic` → `viduq1-classic`

说明：
- 计费所用的 `OriginalModel` 采用新命名规则 `vidu-<action>-<model>-<resolution>-<duration>s[-style]`（例如 `vidu-text2video-vidu1.5-720p-8s`），与默认定价模型键保持一致。
- 若传入其他未知模型名，系统将原样透传，避免阻断新模型接入。
