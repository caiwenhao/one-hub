# 多域名多品牌支持 - 技术设计

## Context

One Hub 当前仅支持单一品牌配置，系统名称、Logo 等信息是全局的。为了支持多品牌运营（如 Kapon AI 和 Grouplay AI），需要实现基于域名的品牌识别和动态配置加载机制。

### 背景
- 现有系统：单品牌，配置存储在 `options` 表
- 目标场景：同一套系统支持多个品牌，不同域名显示不同品牌
- 技术栈：Go + Gin（后端）、React + MUI（前端）

### 约束条件
- 必须向后兼容现有单品牌配置
- 不能影响现有 API 功能
- 性能开销最小化
- 配置管理简单易用

## Goals / Non-Goals

### Goals
1. 支持配置多个品牌及其域名映射
2. 根据访问域名自动识别并返回对应品牌配置
3. 前端根据品牌配置动态加载 Logo、Favicon 等资源
4. 支持每个品牌独立的 meta 信息（title、description）
5. 保持向后兼容，未配置多品牌时使用原有逻辑

### Non-Goals
1. 不支持品牌级别的数据隔离（用户、订单等共享）
2. 不支持品牌级别的权限控制
3. 不支持运行时动态添加品牌（需重启服务）
4. 不支持品牌级别的独立配置（除品牌标识外）


## Decisions

### 1. 品牌配置存储方式

**决策**：使用数据库存储品牌配置，启动时加载到内存缓存

**理由**：
- 支持通过管理后台动态管理品牌
- 配置变更无需重启服务（刷新缓存即可）
- 便于扩展和维护
- 性能优化：内存缓存 + 数据库持久化

**备选方案**：
- 方案 A：配置文件（YAML）
  - 优点：简单，便于版本控制
  - 缺点：需要重启服务，不支持动态管理
- 方案 B：环境变量
  - 优点：部署灵活
  - 缺点：配置复杂，不支持多品牌

### 2. 域名匹配策略

**决策**：使用精确匹配 + 默认品牌的策略

**匹配规则**：
1. 精确匹配：`models.kapon.cloud` → `kapon` 品牌
2. 端口匹配：`localhost:3000` → 支持开发环境
3. 默认品牌：未匹配到任何品牌时使用第一个品牌或全局配置

**理由**：
- 简单明确，易于理解和维护
- 满足当前需求（2-3 个品牌）
- 性能最优（map 查找 O(1)）

**备选方案**：
- 方案 A：支持通配符匹配（如 `*.kapon.cloud`）
  - 优点：更灵活
  - 缺点：增加复杂度，当前不需要
- 方案 B：支持正则表达式
  - 优点：最灵活
  - 缺点：性能开销大，过度设计

### 3. 中间件实现方式

**决策**：创建独立的品牌识别中间件，在路由层注册

**实现位置**：`middleware/brand.go`

**执行时机**：在认证中间件之前，在日志中间件之后

**理由**：
- 职责单一，易于测试和维护
- 可复用，所有路由自动生效
- 不侵入业务逻辑

### 4. 前端部署方式

**决策**：Kapon AI 保持原有部署，Grouplay AI 独立部署并反向代理

**理由**：
- 最大程度保持 Kapon AI 原有代码和部署不动
- Grouplay AI 可以独立开发和部署
- 通过反向代理实现统一后端和管理后台访问
- 部署灵活，易于扩展

**架构**：

**Kapon AI（默认品牌）**：
```
用户访问 models.kapon.cloud
  ↓
Kapon AI 服务器（原有部署）
  ├─ 前端静态资源
  ├─ /api/* → 后端服务
  └─ /panel → 管理后台
```

**Grouplay AI（新品牌）**：
```
用户访问 model.grouplay.cn
  ↓
Grouplay AI 前端服务器
  ├─ 前端静态资源（基于 Kapon AI 克隆，移除 panel）
  ├─ /api/* → 反向代理到 Kapon AI 后端（传递 Host 头）
  └─ /panel → 反向代理到 Kapon AI 管理后台
```

**备选方案**：
- 方案 A：统一部署所有前端到后端服务器
  - 优点：部署简单
  - 缺点：需要修改后端路由逻辑，增加复杂度
- 方案 B：使用 CDN 存储前端资源
  - 优点：性能更好
  - 缺点：增加部署复杂度


### 5. API 响应格式

**决策**：在 `/api/status` 接口中新增品牌相关字段

**新增字段**：
```json
{
  "brand_name": "kapon",
  "brand_logo": "/brands/kapon/logo.png",
  "brand_favicon": "/brands/kapon/favicon.ico",
  "system_name": "Kapon AI",
  "description": "Kapon AI 为中小企业..."
}
```

**理由**：
- 最小化 API 变更
- 前端无需额外请求
- 向后兼容（新增字段不影响现有逻辑）

## Architecture

### 系统架构图

```
┌─────────────────────────────────────────────────────────────┐
│                         用户请求                              │
│                  (models.kapon.cloud)                        │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                    Nginx / 反向代理                           │
│              (传递 Host 头到后端服务)                          │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                     Gin HTTP Server                          │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  1. RequestId Middleware                             │   │
│  │  2. Logger Middleware                                │   │
│  │  3. Brand Middleware ← 新增                          │   │
│  │     - 读取 Host 头                                    │   │
│  │     - 匹配品牌配置                                    │   │
│  │     - 注入到 Context                                 │   │
│  │  4. Auth Middleware                                  │   │
│  │  5. RateLimit Middleware                             │   │
│  └──────────────────────────────────────────────────────┘   │
│                         │                                    │
│                         ▼                                    │
│  ┌──────────────────────────────────────────────────────┐   │
│  │              Router / Controller                      │   │
│  │  - /api/status → GetStatus()                         │   │
│  │    从 Context 读取品牌配置并返回                       │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                   品牌配置管理器                              │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  BrandManager (内存)                                 │   │
│  │  - brands: map[string]*Brand                         │   │
│  │  - domainMap: map[string]string                      │   │
│  │  - defaultBrand: *Brand                              │   │
│  │                                                      │   │
│  │  方法:                                                │   │
│  │  - GetBrandByDomain(domain) *Brand                   │   │
│  │  - LoadFromConfig(config)                            │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                   配置文件 (config.yaml)                      │
│  brands:                                                     │
│    - name: kapon                                             │
│      domains: [models.kapon.cloud, localhost:3000]          │
│      system_name: "Kapon AI"                                 │
│      logo: "/brands/kapon/logo.png"                          │
│      ...                                                     │
│    - name: grouplay                                          │
│      domains: [model.grouplay.cn]                            │
│      system_name: "Grouplay AI"                              │
│      ...                                                     │
└─────────────────────────────────────────────────────────────┘
```

### 数据流

#### 1. 启动时初始化
```
main.go
  └─> config.InitConf()
       └─> brand.InitBrandManager()
            └─> 读取 config.yaml 中的 brands 配置
            └─> 构建 domainMap (domain → brand_name)
            └─> 设置 defaultBrand
```

#### 2. 请求处理流程
```
用户请求 (Host: models.kapon.cloud)
  └─> Nginx 转发 (保留 Host 头)
       └─> Gin Server
            └─> Brand Middleware
                 ├─> 读取 c.Request.Host
                 ├─> 调用 BrandManager.GetBrandByDomain()
                 ├─> 将 Brand 对象注入 c.Set("brand", brand)
                 └─> next()
            └─> Controller (GetStatus)
                 ├─> 从 Context 读取 brand: c.Get("brand")
                 ├─> 构建响应 JSON (包含品牌信息)
                 └─> 返回给前端
```

#### 3. 前端渲染流程
```
前端启动
  └─> StatusContext.loadStatus()
       └─> 调用 /api/status
            └─> 接收品牌配置 (brand_logo, system_name, etc.)
            └─> dispatch(SET_SITE_INFO, data)
                 └─> 更新 Redux store
  └─> Logo 组件
       └─> 从 Redux 读取 siteInfo.brand_logo
       └─> 渲染 <img src={brand_logo} />
```


## Components and Interfaces

### 后端组件

#### 1. Brand 数据模型 (`model/brand.go`)

```go
package model

// Brand 品牌配置
type Brand struct {
    Name        string   `yaml:"name" json:"name"`               // 品牌标识
    Domains     []string `yaml:"domains" json:"domains"`         // 关联域名列表
    SystemName  string   `yaml:"system_name" json:"system_name"` // 系统名称
    Logo        string   `yaml:"logo" json:"logo"`               // Logo 路径
    Favicon     string   `yaml:"favicon" json:"favicon"`         // Favicon 路径
    Description string   `yaml:"description" json:"description"` // 描述
    Keywords    string   `yaml:"keywords" json:"keywords"`       // 关键词
    Author      string   `yaml:"author" json:"author"`           // 作者
}

// BrandManager 品牌管理器
type BrandManager struct {
    brands       map[string]*Brand // brand_name -> Brand
    domainMap    map[string]string // domain -> brand_name
    defaultBrand *Brand            // 默认品牌
}

// GetBrandByDomain 根据域名获取品牌
func (bm *BrandManager) GetBrandByDomain(domain string) *Brand

// LoadFromConfig 从配置加载品牌
func (bm *BrandManager) LoadFromConfig(brands []Brand)
```

#### 2. 品牌中间件 (`middleware/brand.go`)

```go
package middleware

import (
    "github.com/gin-gonic/gin"
    "one-api/model"
)

// BrandMiddleware 品牌识别中间件
func BrandMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 获取请求的 Host
        host := c.Request.Host
        
        // 2. 从 BrandManager 获取品牌
        brand := model.GlobalBrandManager.GetBrandByDomain(host)
        
        // 3. 注入到 Context
        if brand != nil {
            c.Set("brand", brand)
        }
        
        c.Next()
    }
}
```

#### 3. Controller 修改 (`controller/misc.go`)

```go
func GetStatus(c *gin.Context) {
    // 获取品牌配置
    var brandInfo gin.H
    if brand, exists := c.Get("brand"); exists {
        b := brand.(*model.Brand)
        brandInfo = gin.H{
            "brand_name":   b.Name,
            "brand_logo":   b.Logo,
            "brand_favicon": b.Favicon,
            "system_name":  b.SystemName,
            "description":  b.Description,
            "keywords":     b.Keywords,
            "author":       b.Author,
        }
    } else {
        // 向后兼容：使用全局配置
        brandInfo = gin.H{
            "system_name": config.SystemName,
            "logo":        config.Logo,
        }
    }
    
    // 合并到响应
    response := gin.H{
        "success": true,
        "data": gin.H{
            "version": config.Version,
            // ... 其他字段
        },
    }
    
    // 合并品牌信息
    for k, v := range brandInfo {
        response["data"].(gin.H)[k] = v
    }
    
    c.JSON(http.StatusOK, response)
}
```

### 前端组件

#### 1. Logo 组件修改 (`web/src/ui-component/Logo.jsx`)

```jsx
import { useSelector } from 'react-redux';

const Logo = () => {
  const siteInfo = useSelector((state) => state.siteInfo);

  if (siteInfo.isLoading) {
    return null;
  }

  // 优先使用品牌 logo，否则使用全局 logo
  const logoSrc = siteInfo.brand_logo || siteInfo.logo || '/logo.png';
  const altText = siteInfo.system_name || 'Logo';

  return <img src={logoSrc} alt={altText} height="50" />;
};

export default Logo;
```

#### 2. Redux Store 扩展 (`web/src/store/siteInfoReducer.js`)

```javascript
export const initialState = {
  ...config.siteInfo,
  brand_name: '',
  brand_logo: '',
  brand_favicon: '',
  description: '',
  keywords: '',
  author: '',
  ownedby: []
};
```

### 配置文件格式

#### `config.yaml` 品牌配置段

```yaml
# 品牌配置（可选，未配置则使用单品牌模式）
brands:
  - name: "kapon"
    domains:
      - "models.kapon.cloud"
      - "localhost:3000"  # 开发环境
    system_name: "Kapon AI"
    logo: "/brands/kapon/logo.png"
    favicon: "/brands/kapon/favicon.ico"
    description: "Kapon AI 为中小企业与开发者提供100%官方正源大模型API"
    keywords: "AI API,大模型API,OpenAI API,企业级AI,Kapon AI"
    author: "Kapon AI"
    
  - name: "grouplay"
    domains:
      - "model.grouplay.cn"
    system_name: "Grouplay AI"
    logo: "/brands/grouplay/logo.png"
    favicon: "/brands/grouplay/favicon.ico"
    description: "Grouplay AI 企业级AI服务平台"
    keywords: "AI API,企业AI,Grouplay AI"
    author: "Grouplay AI"
```

## Data Models

### Brand 配置结构

```yaml
Brand:
  name: string          # 品牌唯一标识（kebab-case）
  domains: []string     # 关联域名列表（支持端口）
  system_name: string   # 系统显示名称
  logo: string          # Logo 文件路径（相对于 public/）
  favicon: string       # Favicon 文件路径
  description: string   # 品牌描述（用于 meta）
  keywords: string      # SEO 关键词
  author: string        # 作者信息
```

### API 响应格式

#### `/api/status` 响应（新增字段）

```json
{
  "success": true,
  "data": {
    "version": "v1.0.0",
    "system_name": "Kapon AI",
    "logo": "/brands/kapon/logo.png",
    
    // 新增品牌字段
    "brand_name": "kapon",
    "brand_logo": "/brands/kapon/logo.png",
    "brand_favicon": "/brands/kapon/favicon.ico",
    "description": "Kapon AI 为中小企业...",
    "keywords": "AI API,大模型API...",
    "author": "Kapon AI",
    
    // 其他现有字段...
    "email_verification": false,
    "github_oauth": false
  }
}
```

## Error Handling

### 错误场景及处理

1. **配置文件格式错误**
   - 场景：brands 配置格式不正确
   - 处理：启动时记录错误日志，使用单品牌模式
   - 日志：`logger.SysError("failed to load brands config: " + err.Error())`

2. **域名未匹配到品牌**
   - 场景：访问的域名不在任何品牌的 domains 列表中
   - 处理：使用默认品牌（第一个品牌或全局配置）
   - 日志：`logger.SysLog("domain not matched, using default brand")`

3. **品牌资源文件不存在**
   - 场景：配置的 logo 或 favicon 文件不存在
   - 处理：前端显示默认图片或占位符
   - 用户体验：不影响功能，仅显示问题

4. **多个品牌配置相同域名**
   - 场景：配置错误，两个品牌包含相同域名
   - 处理：使用第一个匹配的品牌，记录警告日志
   - 日志：`logger.SysWarn("duplicate domain in brands config")`

## Testing Strategy

### 单元测试

#### 后端测试

1. **BrandManager 测试** (`model/brand_test.go`)
   ```go
   func TestGetBrandByDomain(t *testing.T)
   func TestLoadFromConfig(t *testing.T)
   func TestDefaultBrand(t *testing.T)
   ```

2. **Brand Middleware 测试** (`middleware/brand_test.go`)
   ```go
   func TestBrandMiddleware(t *testing.T)
   func TestBrandMiddlewareWithUnknownDomain(t *testing.T)
   ```

#### 前端测试

1. **Logo 组件测试** (`web/src/ui-component/Logo.test.jsx`)
   ```javascript
   test('renders brand logo when available')
   test('falls back to default logo')
   test('renders nothing when loading')
   ```

### 集成测试

1. **多域名访问测试**
   - 使用不同 Host 头请求 `/api/status`
   - 验证返回的品牌信息正确

2. **前端渲染测试**
   - 模拟不同品牌配置
   - 验证 Logo 组件渲染正确

### 手动测试清单

- [ ] 配置两个品牌（kapon 和 grouplay）
- [ ] 准备品牌资源文件（logo.png、favicon.ico）
- [ ] 使用 `models.kapon.cloud` 访问，验证显示 Kapon 品牌
- [ ] 使用 `model.grouplay.cn` 访问，验证显示 Grouplay 品牌
- [ ] 使用未配置的域名访问，验证显示默认品牌
- [ ] 删除品牌配置，验证回退到单品牌模式
- [ ] 检查浏览器控制台无错误
- [ ] 检查服务器日志无错误

## Migration Plan

### 升级步骤

#### 1. 准备阶段
- 准备各品牌的 Logo 和 Favicon 文件
- 规划域名映射关系
- 编写品牌配置

#### 2. 代码部署
```bash
# 1. 拉取最新代码
git pull origin main

# 2. 构建前端
cd web
yarn build

# 3. 构建后端
cd ..
go build -o one-api

# 4. 停止服务
systemctl stop one-api

# 5. 备份配置文件
cp config.yaml config.yaml.bak

# 6. 更新配置文件（添加 brands 配置）
vim config.yaml

# 7. 上传品牌资源文件
mkdir -p web/build/brands/kapon
mkdir -p web/build/brands/grouplay
cp kapon-logo.png web/build/brands/kapon/logo.png
cp kapon-favicon.ico web/build/brands/kapon/favicon.ico
cp grouplay-logo.png web/build/brands/grouplay/logo.png
cp grouplay-favicon.ico web/build/brands/grouplay/favicon.ico

# 8. 启动服务
systemctl start one-api

# 9. 验证服务
curl -H "Host: models.kapon.cloud" http://localhost:3000/api/status
curl -H "Host: model.grouplay.cn" http://localhost:3000/api/status
```

#### 3. DNS 配置
```bash
# 确保两个域名都指向服务器
# models.kapon.cloud -> 服务器 IP
# model.grouplay.cn -> 服务器 IP
```

#### 4. Nginx 配置

**Kapon AI（保持原有配置）**：
```nginx
server {
    listen 80;
    server_name models.kapon.cloud;
    
    # 原有配置保持不变
    location / {
        # 前端静态资源或反向代理配置
    }
}
```

**Grouplay AI（新增配置）**：
```nginx
server {
    listen 80;
    server_name model.grouplay.cn;
    
    # 前端静态资源
    location / {
        root /var/www/grouplay;
        try_files $uri $uri/ /index.html;
    }
    
    # API 反向代理到 Kapon AI 后端
    location /api/ {
        proxy_pass http://kapon-backend-server/api/;
        proxy_set_header Host $host;  # 重要：传递 Host 头
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
    
    # 管理后台反向代理到 Kapon AI
    location /panel {
        proxy_pass http://kapon-server/panel;
        proxy_set_header Host models.kapon.cloud;  # 使用 Kapon AI 域名
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

### 回滚方案

如果出现问题，可以快速回滚：

```bash
# 1. 停止服务
systemctl stop one-api

# 2. 恢复配置文件
cp config.yaml.bak config.yaml

# 3. 使用旧版本二进制
cp one-api.old one-api

# 4. 启动服务
systemctl start one-api
```

### 数据迁移

**无需数据迁移**：此功能不涉及数据库结构变更

## Risks / Trade-offs

### 风险

1. **配置错误风险**
   - 风险：管理员配置品牌信息错误
   - 缓解：提供配置示例和验证工具
   - 影响：中等

2. **域名冲突风险**
   - 风险：多个品牌配置相同域名
   - 缓解：启动时检查并记录警告
   - 影响：低

3. **资源文件缺失风险**
   - 风险：配置的 logo 文件不存在
   - 缓解：前端使用默认图片兜底
   - 影响：低

### Trade-offs

1. **灵活性 vs 复杂度**
   - 选择：配置文件 + 内存缓存
   - 放弃：运行时动态修改品牌
   - 理由：当前需求不需要动态修改，简单优先

2. **性能 vs 功能**
   - 选择：启动时加载，运行时内存查找
   - 放弃：数据库存储，支持热更新
   - 理由：性能优先，品牌配置变更频率极低

3. **向后兼容 vs 代码简洁**
   - 选择：保持向后兼容
   - 放弃：移除旧的全局配置逻辑
   - 理由：平滑升级，降低用户迁移成本

## Open Questions

1. **是否需要支持品牌级别的数据隔离？**
   - 当前方案：不支持，所有品牌共享用户、订单等数据
   - 如果需要：需要在数据库表中添加 brand_id 字段
   - 决策：暂不支持，等待实际需求

2. **是否需要支持子域名通配符？**
   - 当前方案：仅支持精确匹配
   - 如果需要：可以支持 `*.kapon.cloud` 匹配
   - 决策：暂不支持，当前需求明确

3. **是否需要支持品牌级别的独立配置？**
   - 当前方案：仅品牌标识不同，其他配置共享
   - 如果需要：可以为每个品牌配置独立的支付、通知等
   - 决策：暂不支持，保持简单

4. **是否需要管理后台支持品牌管理？**
   - 当前方案：通过配置文件管理
   - 如果需要：可以开发品牌管理界面
   - 决策：暂不需要，配置文件足够
