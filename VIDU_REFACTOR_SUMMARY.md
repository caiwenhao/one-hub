# Vidu 渠道重构总结文档

## 概述

本次重构完全按照 Vidu 官方文档对 One Hub 平台的 Vidu 渠道进行了全面升级，实现了完整的 API 对齐、新模型支持和精确的计费系统。

## 主要变更

### 1. 核心配置更新

#### 1.1 baseURL 调整
- **变更前**: `https://api.vidu.com`
- **变更后**: `https://api.vidu.cn`
- **影响文件**: `providers/vidu/base.go`

#### 1.2 路由配置
- **保持**: `/vidu/ent/v2/` 路由前缀
- **新增**: 官方查询接口 `GET /vidu/ent/v2/tasks/{id}/creations`
- **新增**: 取消任务接口 `POST /vidu/ent/v2/tasks/{id}/cancel`
- **影响文件**: `router/relay-router.go`

### 2. 类型定义重构

#### 2.1 请求体结构完全对齐
**影响文件**: `providers/vidu/type.go`

**新增模型支持**:
```go
const (
    ViduModelQ2Pro     = "viduq2-pro"     // 新模型，效果好，细节丰富
    ViduModelQ2Turbo   = "viduq2-turbo"   // 新模型，效果好，生成快
    ViduModelQ1        = "viduq1"         // 画面清晰，平滑转场，运镜稳定
    ViduModelQ1Classic = "viduq1-classic" // 画面清晰，转场、运镜更丰富
    ViduModel20        = "vidu2.0"        // 生成速度快
    ViduModel15        = "vidu1.5"        // 动态幅度大
)
```

**新增参数支持**:
- `aspect_ratio`: 比例参数（16:9、9:16、1:1、auto）
- `style`: 风格参数（general、anime）- 仅 text2video 支持
- `watermark`: 水印参数

#### 2.2 响应体结构对齐
```go
type ViduResponse struct {
    TaskID            string   `json:"task_id"`
    State             string   `json:"state"`
    Model             string   `json:"model,omitempty"`
    Images            []string `json:"images,omitempty"`
    Prompt            string   `json:"prompt,omitempty"`
    Duration          *int     `json:"duration,omitempty"`
    Seed              *int     `json:"seed,omitempty"`
    Resolution        string   `json:"resolution,omitempty"`
    MovementAmplitude string   `json:"movement_amplitude,omitempty"`
    BGM               *bool    `json:"bgm,omitempty"`
    Payload           string   `json:"payload,omitempty"`
    OffPeak           *bool    `json:"off_peak,omitempty"`
    Credits           *int     `json:"credits,omitempty"`
    Watermark         *bool    `json:"watermark,omitempty"`
    CreatedAt         string   `json:"created_at,omitempty"`
    Style             string   `json:"style,omitempty"`
    AspectRatio       string   `json:"aspect_ratio,omitempty"`
}
```

### 3. 新接口实现

#### 3.1 参考生图接口
```go
type ViduReference2ImageRequest struct {
    Model       string   `json:"model"`                    // 模型名：viduq1
    Images      []string `json:"images"`                  // 图像参考，支持1-7张图片
    Prompt      string   `json:"prompt"`                  // 文本提示词
    Seed        *int     `json:"seed,omitempty"`          // 随机种子
    AspectRatio string   `json:"aspect_ratio,omitempty"` // 比例：16:9、9:16、1:1、auto
    Payload     string   `json:"payload,omitempty"`       // 透传参数
    CallbackURL string   `json:"callback_url,omitempty"`  // 回调URL
}
```

#### 3.2 取消任务接口
```go
type ViduCancelRequest struct {
    ID string `json:"id"` // 任务ID
}
```

**新增路由处理器**: `relay/task/vidu/fetch.go`
```go
func RelayTaskCancel(c *gin.Context) {
    // 取消任务逻辑
    // 更新本地任务状态为已取消
    // 返回空响应
}
```

### 4. Provider 实现重构

#### 4.1 统一接口设计
**影响文件**: `providers/vidu/base.go`

```go
func (p *ViduProvider) Submit(action string, request interface{}) (*ViduResponse, *types.OpenAIError) {
    // 根据不同action类型处理请求
    // 支持 ViduTaskRequest 和 ViduReference2ImageRequest
}
```

#### 4.2 模型名称规范化
```go
func normalizeModelName(model string) string {
    // 支持新模型
    case "viduq2-pro":
        return ViduModelQ2Pro
    case "viduq2-turbo":
        return ViduModelQ2Turbo
    // 向后兼容
    case "vidu-q2-pro":
        return ViduModelQ2Pro
    // ...
}
```

### 5. 价格配置完全对齐

#### 5.1 基于官方积分消耗表的精确定价
**影响文件**: `model/price.go`

**新模型价格配置示例**:
```go
var DefaultViduPrice = map[string]float64{
    // viduq2-pro 模型 - 2-8秒支持
    "vidu-img2video-viduq2-pro-2s-720p":   8,
    "vidu-img2video-viduq2-pro-2s-1080p":  16,
    "vidu-img2video-viduq2-pro-5s-720p":   20,
    "vidu-img2video-viduq2-pro-5s-1080p":  40,
    
    // viduq2-turbo 模型 - 更快生成
    "vidu-img2video-viduq2-turbo-5s-720p":  15,
    "vidu-img2video-viduq2-turbo-5s-1080p": 30,
    
    // 参考生图新接口
    "vidu-reference2image-viduq1": 2,
    
    // 向后兼容的简化格式
    "vidu-img2video-viduq2-pro-5s": 20,
}
```

#### 5.2 智能计费系统
**影响文件**: `relay/task/vidu/submit.go`

**动态模型名称构建**:
```go
func buildDetailedModelName(action, model string, duration int, resolution, style string) string {
    baseName := fmt.Sprintf("vidu-%s-%s-%ds-%s", action, model, duration, resolution)
    
    // 文生视频支持风格参数
    if action == ViduActionText2Video && style != "" && style != "general" {
        baseName = fmt.Sprintf("%s-%s", baseName, style)
    }
    
    return baseName
}
```

**智能默认值设置**:
```go
func getDefaultDuration(model string) int {
    switch model {
    case ViduModelQ2Pro, ViduModelQ2Turbo:
        return 5  // 新模型默认5秒
    case ViduModelQ1, ViduModelQ1Classic:
        return 5  // 经典模型5秒
    case ViduModel20, ViduModel15:
        return 4  // 快速模型4秒
    default:
        return 5
    }
}
```

### 6. 路由和中间件更新

#### 6.1 新增路由
**影响文件**: `router/relay-router.go`

```go
func setViduRouter(router *gin.Engine) {
    relayViduRouter := router.Group("/vidu")
    relayViduRouter.Use(middleware.RelayPanicRecover(), middleware.OpenaiAuth(), middleware.Distribute())
    
    // 查询接口
    relayViduRouter.GET("/ent/v2/task/:task_id", vidu.RelayTaskFetch)
    relayViduRouter.GET("/ent/v2/tasks", vidu.RelayTaskFetchs)
    relayViduRouter.GET("/ent/v2/tasks/:task_id/creations", vidu.RelayTaskFetch) // 新增官方查询接口
    
    // 取消任务接口
    relayViduRouter.POST("/ent/v2/tasks/:task_id/cancel", vidu.RelayTaskCancel) // 新增取消接口
    
    relayViduRouter.Use(middleware.DynamicRedisRateLimiter())
    {
        // 任务提交接口
        relayViduRouter.POST("/ent/v2/:action", task.RelayTaskSubmit)
    }
}
```

### 7. 任务处理逻辑优化

#### 7.1 多类型请求支持
**影响文件**: `relay/task/vidu/submit.go`

```go
type ViduTask struct {
    base.TaskBase
    Action   string
    Request  interface{} // 支持不同类型的请求
    Provider *ViduProvider
}

func (t *ViduTask) Init() *base.TaskError {
    // 根据不同action解析不同的请求体
    switch t.Action {
    case ViduActionReference2Image:
        var req ViduReference2ImageRequest
        // 解析参考生图请求
    default:
        var req ViduTaskRequest
        // 解析通用视频任务请求
    }
}
```

#### 7.2 查询接口更新
**影响文件**: `relay/task/vidu/fetch.go`

```go
func RelayTaskFetch(c *gin.Context) {
    // 使用新的官方查询接口
    resp, openaiErr := viduProvider.QueryCreations(task.TaskID)
    if openaiErr != nil {
        StringError(c, http.StatusInternalServerError, "query_failed", openaiErr.Message)
        return
    }
    
    c.JSON(http.StatusOK, resp)
}
```

## 支持的接口一览

| 接口名称 | 路径 | 方法 | 支持模型 | 新增功能 |
|---------|------|------|----------|----------|
| 图生视频 | `/vidu/ent/v2/img2video` | POST | 所有模型 | 新模型支持 |
| 参考生视频 | `/vidu/ent/v2/reference2video` | POST | viduq1, vidu2.0, vidu1.5 | aspect_ratio 参数 |
| 首尾帧 | `/vidu/ent/v2/start-end2video` | POST | 所有模型 | 新模型支持 |
| 文生视频 | `/vidu/ent/v2/text2video` | POST | viduq1, vidu1.5 | style 参数支持 |
| **参考生图** | `/vidu/ent/v2/reference2image` | POST | viduq1 | **全新接口** |
| 查询任务 | `/vidu/ent/v2/tasks/{id}/creations` | GET | - | 官方查询接口 |
| **取消任务** | `/vidu/ent/v2/tasks/{id}/cancel` | POST | - | **全新接口** |

## 模型支持矩阵

| 模型 | 图生视频 | 参考生视频 | 首尾帧 | 文生视频 | 参考生图 | 时长支持 | 分辨率支持 |
|------|----------|-----------|--------|----------|----------|----------|-----------|
| viduq2-pro | ✅ | ❌ | ✅ | ❌ | ❌ | 2-8s | 720p, 1080p |
| viduq2-turbo | ✅ | ❌ | ✅ | ❌ | ❌ | 2-8s | 720p, 1080p |
| viduq1 | ✅ | ✅ | ✅ | ✅ | ✅ | 5s | 1080p |
| viduq1-classic | ✅ | ❌ | ✅ | ❌ | ❌ | 5s | 1080p |
| vidu2.0 | ✅ | ✅ | ✅ | ❌ | ❌ | 4s, 8s | 360p, 720p, 1080p |
| vidu1.5 | ✅ | ✅ | ✅ | ✅ | ❌ | 4s, 8s | 360p, 720p, 1080p |

## 价格配置示例

### 新模型定价（按积分）

| 模型 | 功能 | 时长 | 分辨率 | 积分消耗 | 适用场景 |
|------|------|------|--------|----------|----------|
| viduq2-pro | 图生视频 | 5s | 720p | 20 | 高质量视频生成 |
| viduq2-pro | 图生视频 | 5s | 1080p | 40 | 超高清视频生成 |
| viduq2-turbo | 图生视频 | 5s | 720p | 15 | 快速视频生成 |
| viduq2-turbo | 图生视频 | 5s | 1080p | 30 | 快速高清生成 |
| viduq1 | 参考生图 | - | - | 2 | 图像生成 |

### 向后兼容定价

为保持向后兼容，同时提供简化版本的价格配置：
```
vidu-img2video-viduq2-pro-5s: 20
vidu-img2video-viduq2-turbo-5s: 15
vidu-reference2image-viduq1: 2
```

## 技术优势

### 1. 完全兼容性
- ✅ 请求体和响应体 100% 对齐官方 API 文档
- ✅ 支持所有官方模型和参数
- ✅ 完整的错误处理机制

### 2. 智能计费
- ✅ 基于官方积分消耗表的精确定价
- ✅ 动态模型名称构建
- ✅ 智能默认参数设置
- ✅ 向后兼容的价格配置

### 3. 扩展性
- ✅ 支持新模型的快速接入
- ✅ 灵活的参数配置系统
- ✅ 模块化的接口设计

### 4. 用户体验
- ✅ 统一的 API 接口
- ✅ 详细的错误信息
- ✅ 实时任务状态查询
- ✅ 任务取消功能

## 测试建议

### 1. 基础功能测试
```bash
# 图生视频 - 新模型
curl -X POST "https://your-domain.com/vidu/ent/v2/img2video" \
  -H "Authorization: Token your_api_key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "viduq2-pro",
    "images": ["https://example.com/image.jpg"],
    "duration": 5,
    "resolution": "1080p",
    "watermark": false
  }'

# 参考生图 - 新接口
curl -X POST "https://your-domain.com/vidu/ent/v2/reference2image" \
  -H "Authorization: Token your_api_key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "viduq1",
    "images": ["https://example.com/ref1.jpg", "https://example.com/ref2.jpg"],
    "prompt": "生成与参考图相似的图像",
    "aspect_ratio": "16:9"
  }'

# 取消任务 - 新接口
curl -X POST "https://your-domain.com/vidu/ent/v2/tasks/{task_id}/cancel" \
  -H "Authorization: Token your_api_key" \
  -H "Content-Type: application/json"
```

### 2. 计费系统测试
- 验证新模型的价格计算
- 测试分辨率和时长组合的计费
- 检查向后兼容的价格配置

### 3. 错误处理测试
- 测试无效模型名称的处理
- 验证参数验证逻辑
- 检查网络错误的处理

## 文件变更清单

### 核心文件
- ✅ `providers/vidu/type.go` - 类型定义重构
- ✅ `providers/vidu/base.go` - Provider 实现更新
- ✅ `model/price.go` - 价格配置对齐
- ✅ `relay/task/vidu/submit.go` - 任务提交逻辑优化
- ✅ `relay/task/vidu/fetch.go` - 查询和取消接口
- ✅ `router/relay-router.go` - 路由配置更新

### 配置文件
- ✅ `model/model_ownedby.go` - 模型归属已配置

## 部署注意事项

### 1. 数据库迁移
- 新的价格配置会在系统启动时自动加载
- 建议在后台"模型价格-更新价格"中同步新模型

### 2. 配置检查
- 确认 Vidu 渠道的 baseURL 已更新为 `https://api.vidu.cn`
- 验证 API Key 的有效性

### 3. 监控建议
- 监控新接口的调用情况
- 跟踪计费准确性
- 观察错误率变化

## 总结

本次重构成功实现了：
- 🎯 **100% 官方文档对齐** - 所有请求和响应格式完全一致
- 🚀 **新模型全面支持** - viduq2-pro、viduq2-turbo 等新模型
- 💰 **精确计费系统** - 基于官方积分表的精准定价
- 🔧 **新接口实现** - reference2image 和 cancel 功能
- 📈 **向后兼容** - 保持现有用户的使用体验
- 🛡️ **强化错误处理** - 完善的异常处理机制

Vidu 渠道现已完全对齐官方 API，为用户提供最新、最完整的 AI 视频生成服务！