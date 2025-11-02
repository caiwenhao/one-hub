# 多品牌多前端架构设计

## 架构概览

```
┌─────────────────────────────────────────────────────────────────┐
│                         用户访问层                               │
├─────────────────────────────────────────────────────────────────┤
│  models.kapon.cloud          model.grouplay.cn                  │
│  ├─ / (Kapon 门户)           ├─ / (Grouplay 门户)               │
│  └─ /panel (管理后台)        └─ /panel (管理后台)               │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      后端服务（共享）                            │
├─────────────────────────────────────────────────────────────────┤
│  ┌──────────────────┐  ┌──────────────────┐                    │
│  │ 品牌识别中间件    │  │  前端资源路由     │                    │
│  │ (基于 Host 头)   │  │  (基于品牌)      │                    │
│  └──────────────────┘  └──────────────────┘                    │
│                                                                 │
│  ┌──────────────────┐  ┌──────────────────┐                    │
│  │   统一 API       │  │   共享数据库      │                    │
│  │  (所有品牌共享)  │  │  (所有品牌共享)   │                    │
│  └──────────────────┘  └──────────────────┘                    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      前端资源层                                  │
├─────────────────────────────────────────────────────────────────┤
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────┐ │
│  │ Kapon 前端       │  │ Grouplay 前端     │  │ 管理后台      │ │
│  │ (独立项目)       │  │ (独立项目)        │  │ (共享项目)    │ │
│  │ /brands/kapon/   │  │ /brands/grouplay/ │  │ /admin/      │ │
│  └──────────────────┘  └──────────────────┘  └──────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## 核心组件

### 1. 品牌识别中间件

**职责**：
- 从 HTTP 请求的 Host 头读取域名
- 查询品牌管理器获取对应品牌配置
- 将品牌信息注入到请求上下文

**流程**：
```
请求到达 → 读取 Host 头 → 查询品牌 → 注入上下文 → 继续处理
```

**示例**：
```go
// 伪代码
func BrandMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        host := c.Request.Host
        brand := brandManager.GetBrandByDomain(host)
        c.Set("brand", brand)
        c.Next()
    }
}
```

### 2. 前端资源路由

**职责**：
- 根据品牌配置路由到对应的前端资源
- 支持 SPA 路由（非文件路径返回 index.html）
- 处理管理后台路由（/panel）

**路由规则**：
```
models.kapon.cloud/          → /brands/kapon/index.html
models.kapon.cloud/assets/*  → /brands/kapon/assets/*
models.kapon.cloud/panel     → /admin/index.html
models.kapon.cloud/api/*     → API 处理器

model.grouplay.cn/           → /brands/grouplay/index.html
model.grouplay.cn/assets/*   → /brands/grouplay/assets/*
model.grouplay.cn/panel      → /admin/index.html
model.grouplay.cn/api/*      → API 处理器
```

**示例**：
```go
// 伪代码
func FrontendRouter(c *gin.Context) {
    brand := c.MustGet("brand").(Brand)
    path := c.Request.URL.Path
    
    if strings.HasPrefix(path, "/panel") {
        // 管理后台
        c.File("/admin/index.html")
    } else if strings.HasPrefix(path, "/api") {
        // API 请求，跳过
        c.Next()
    } else {
        // 品牌前端
        filePath := brand.FrontendPath + path
        if fileExists(filePath) {
            c.File(filePath)
        } else {
            // SPA 路由，返回 index.html
            c.File(brand.FrontendPath + "/index.html")
        }
    }
}
```

### 3. 品牌管理器

**职责**：
- 加载和管理品牌配置
- 提供域名到品牌的查询接口
- 处理默认品牌逻辑

**接口**：
```go
type BrandManager interface {
    LoadBrands(config Config) error
    GetBrandByDomain(domain string) *Brand
    GetDefaultBrand() *Brand
    GetAllBrands() []*Brand
}
```

## 前端项目结构

### 目录组织

```
web/
├── kapon-portal/              # Kapon 品牌独立前端
│   ├── package.json
│   ├── vite.config.js         # 或其他构建配置
│   ├── src/
│   │   ├── App.jsx
│   │   ├── pages/
│   │   ├── components/
│   │   └── styles/
│   └── dist/ → ../public/brands/kapon/
│
├── grouplay-portal/           # Grouplay 品牌独立前端
│   ├── package.json
│   ├── vite.config.js
│   ├── src/
│   │   ├── App.vue            # 可以使用不同框架
│   │   ├── pages/
│   │   ├── components/
│   │   └── styles/
│   └── dist/ → ../public/brands/grouplay/
│
├── admin/                     # 管理后台（共享）
│   ├── package.json
│   ├── src/
│   │   ├── components/
│   │   │   └── Logo.jsx       # 支持动态品牌 logo
│   │   ├── pages/
│   │   └── App.jsx
│   └── dist/ → ../public/admin/
│
└── public/
    ├── brands/
    │   ├── kapon/             # Kapon 前端构建产物
    │   │   ├── index.html
    │   │   ├── assets/
    │   │   ├── logo.png
    │   │   └── favicon.ico
    │   └── grouplay/          # Grouplay 前端构建产物
    │       ├── index.html
    │       ├── assets/
    │       ├── logo.png
    │       └── favicon.ico
    └── admin/                 # 管理后台构建产物
        ├── index.html
        └── assets/
```

### 构建脚本

```json
{
  "scripts": {
    "build:kapon": "cd kapon-portal && npm run build",
    "build:grouplay": "cd grouplay-portal && npm run build",
    "build:admin": "cd admin && npm run build",
    "build:all": "npm run build:kapon && npm run build:grouplay && npm run build:admin",
    "dev:kapon": "cd kapon-portal && npm run dev",
    "dev:grouplay": "cd grouplay-portal && npm run dev",
    "dev:admin": "cd admin && npm run dev"
  }
}
```

## 数据流

### 1. 用户访问流程

```
用户访问 models.kapon.cloud
    ↓
Nginx 转发到后端（保留 Host 头）
    ↓
品牌识别中间件识别为 kapon 品牌
    ↓
前端资源路由返回 /brands/kapon/index.html
    ↓
前端加载并调用 /api/status 获取品牌配置
    ↓
前端根据品牌配置渲染 UI
```

### 2. API 调用流程

```
前端调用 /api/user/info
    ↓
请求自动包含 Host 头（models.kapon.cloud）
    ↓
品牌识别中间件识别品牌并注入上下文
    ↓
API 处理器处理请求（可选择性使用品牌信息）
    ↓
返回响应（数据在所有品牌间共享）
```

### 3. 管理后台访问流程

```
管理员访问 models.kapon.cloud/panel
    ↓
品牌识别中间件识别为 kapon 品牌
    ↓
前端资源路由返回 /admin/index.html
    ↓
管理后台加载并调用 /api/status
    ↓
获取 kapon 品牌的 logo 和系统名称
    ↓
管理后台显示 kapon 品牌标识
    ↓
管理员看到的数据与品牌无关（共享数据）
```

## 部署模式

### 模式 1：统一部署（推荐）

**特点**：
- 所有前端和后端部署在同一服务器
- 通过域名区分品牌
- 配置简单，运维方便

**配置**：
```yaml
brands:
  - name: kapon
    frontend_type: embedded
    frontend_path: "/brands/kapon/"
  - name: grouplay
    frontend_type: embedded
    frontend_path: "/brands/grouplay/"
```

**Nginx 配置**：
```nginx
server {
    listen 80;
    server_name models.kapon.cloud model.grouplay.cn;
    
    location / {
        proxy_pass http://backend:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### 模式 2：前端独立部署

**特点**：
- 前端独立部署到 CDN 或静态服务器
- 后端只提供 API 服务
- 需要配置 CORS

**配置**：
```yaml
brands:
  - name: kapon
    frontend_type: external
    frontend_url: "https://models.kapon.cloud"
  - name: grouplay
    frontend_type: external
    frontend_url: "https://model.grouplay.cn"

cors:
  enabled: true
  allowed_origins:
    - "https://models.kapon.cloud"
    - "https://model.grouplay.cn"
```

**部署架构**：
```
CDN (models.kapon.cloud)  ─┐
                           ├─→ 后端 API (api.example.com)
CDN (model.grouplay.cn)   ─┘
```

## 向后兼容

### 单品牌模式

如果配置文件不包含 `brands` 配置段，系统自动回退到单品牌模式：

```yaml
# 传统配置（无 brands 配置）
SystemName: "One API"
Logo: "/logo.png"
Favicon: "/favicon.ico"
```

**行为**：
- 品牌识别中间件返回默认品牌
- 前端资源路由使用默认路径
- API 返回传统的配置字段
- 所有功能正常工作

## 扩展性

### 添加新品牌

只需在配置文件中添加新品牌配置：

```yaml
brands:
  - name: newbrand
    domains:
      - new.example.com
    system_name: "New Brand"
    frontend_type: embedded
    frontend_path: "/brands/newbrand/"
```

然后：
1. 创建新的前端项目 `web/newbrand-portal/`
2. 构建并输出到 `web/public/brands/newbrand/`
3. 重启服务器加载新配置

### 支持更多前端类型

未来可以扩展支持：
- `frontend_type: redirect`：重定向到外部 URL
- `frontend_type: proxy`：反向代理到其他服务
- `frontend_type: dynamic`：动态生成前端内容

## 安全考虑

1. **域名验证**：只允许配置的域名访问
2. **CORS 配置**：严格限制允许的源
3. **路径遍历防护**：防止访问品牌目录外的文件
4. **管理后台认证**：所有品牌的 /panel 都需要认证
5. **API 权限**：基于用户权限，不基于品牌
