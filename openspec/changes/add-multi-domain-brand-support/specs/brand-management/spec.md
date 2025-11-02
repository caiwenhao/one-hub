# Brand Management Specification

## ADDED Requirements

### Requirement 1: 品牌配置管理

系统 SHALL 支持配置多个品牌及其关联的域名和品牌标识信息。

#### Scenario: 配置多个品牌
- **GIVEN** 管理员在配置文件中定义了多个品牌
- **AND** 每个品牌包含唯一的名称、关联域名列表、系统名称、Logo 路径等信息
- **WHEN** 系统启动时
- **THEN** 系统 SHALL 成功加载所有品牌配置到内存
- **AND** 系统 SHALL 构建域名到品牌的映射关系
- **AND** 系统 SHALL 记录品牌配置加载成功的日志

#### Scenario: 处理配置错误
- **GIVEN** 配置文件中的品牌配置格式不正确
- **WHEN** 系统启动时
- **THEN** 系统 SHALL 记录错误日志
- **AND** 系统 SHALL 回退到单品牌模式（使用全局配置）
- **AND** 系统 SHALL 继续正常运行

#### Scenario: 处理域名冲突
- **GIVEN** 两个品牌配置了相同的域名
- **WHEN** 系统启动时
- **THEN** 系统 SHALL 记录警告日志
- **AND** 系统 SHALL 使用第一个匹配的品牌
- **AND** 系统 SHALL 继续正常运行

### Requirement 1.1: 品牌数据模型

系统 SHALL 定义品牌数据模型，包含品牌的所有必要信息。

#### Scenario: 品牌模型包含完整信息
- **GIVEN** 品牌数据模型已定义
- **WHEN** 创建品牌实例时
- **THEN** 品牌实例 SHALL 包含以下字段：
  - name（品牌唯一标识）
  - domains（关联域名列表）
  - system_name（系统显示名称）
  - logo（Logo 文件路径）
  - favicon（Favicon 文件路径）
  - description（品牌描述）
  - keywords（SEO 关键词）
  - author（作者信息）

### Requirement 1.2: 品牌管理器

系统 SHALL 提供品牌管理器，负责品牌配置的加载、存储和查询。

#### Scenario: 根据域名查询品牌
- **GIVEN** 品牌管理器已加载多个品牌配置
- **AND** 域名 "models.kapon.cloud" 关联到 "kapon" 品牌
- **WHEN** 调用 GetBrandByDomain("models.kapon.cloud")
- **THEN** 系统 SHALL 返回 "kapon" 品牌对象
- **AND** 品牌对象 SHALL 包含完整的品牌信息

#### Scenario: 查询未知域名
- **GIVEN** 品牌管理器已加载多个品牌配置
- **AND** 域名 "unknown.example.com" 未关联到任何品牌
- **WHEN** 调用 GetBrandByDomain("unknown.example.com")
- **THEN** 系统 SHALL 返回默认品牌对象
- **AND** 默认品牌 SHALL 是配置的第一个品牌或全局配置

#### Scenario: 支持带端口的域名匹配
- **GIVEN** 品牌管理器已加载品牌配置
- **AND** 域名 "localhost:3000" 关联到 "kapon" 品牌
- **WHEN** 调用 GetBrandByDomain("localhost:3000")
- **THEN** 系统 SHALL 返回 "kapon" 品牌对象

### Requirement 1.3: 配置文件支持

系统 SHALL 支持通过 YAML 配置文件定义品牌配置。

#### Scenario: 从配置文件加载品牌
- **GIVEN** config.yaml 文件包含 brands 配置段
- **AND** brands 配置段定义了两个品牌（kapon 和 grouplay）
- **WHEN** 系统启动并加载配置
- **THEN** 系统 SHALL 成功解析 brands 配置
- **AND** 系统 SHALL 创建两个品牌实例
- **AND** 系统 SHALL 将品牌实例加载到品牌管理器

#### Scenario: 配置文件不包含品牌配置
- **GIVEN** config.yaml 文件不包含 brands 配置段
- **WHEN** 系统启动并加载配置
- **THEN** 系统 SHALL 使用单品牌模式
- **AND** 系统 SHALL 使用全局配置（SystemName, Logo 等）
- **AND** 系统 SHALL 正常运行

## ADDED Requirements

### Requirement 2: 品牌识别与注入

系统 SHALL 根据请求的域名自动识别品牌，并将品牌信息注入到请求上下文中。

#### Scenario: 识别请求域名对应的品牌
- **GIVEN** 用户通过域名 "models.kapon.cloud" 访问系统
- **AND** 该域名关联到 "kapon" 品牌
- **WHEN** 请求到达服务器
- **THEN** 品牌识别中间件 SHALL 从 Host 头读取域名
- **AND** 中间件 SHALL 调用品牌管理器查询品牌
- **AND** 中间件 SHALL 将 "kapon" 品牌对象注入到请求上下文
- **AND** 后续处理器 SHALL 能够从上下文读取品牌信息

#### Scenario: 处理未知域名请求
- **GIVEN** 用户通过未配置的域名 "unknown.example.com" 访问系统
- **WHEN** 请求到达服务器
- **THEN** 品牌识别中间件 SHALL 从 Host 头读取域名
- **AND** 中间件 SHALL 调用品牌管理器查询品牌
- **AND** 中间件 SHALL 将默认品牌对象注入到请求上下文
- **AND** 系统 SHALL 正常处理请求

### Requirement 2.1: 品牌识别中间件

系统 SHALL 提供品牌识别中间件，在请求处理流程中自动识别品牌。

#### Scenario: 中间件注册到路由
- **GIVEN** 系统启动时
- **WHEN** 初始化路由
- **THEN** 品牌识别中间件 SHALL 被注册到所有路由
- **AND** 中间件 SHALL 在日志中间件之后执行
- **AND** 中间件 SHALL 在认证中间件之前执行

#### Scenario: 中间件读取 Host 头
- **GIVEN** 请求包含 Host 头 "models.kapon.cloud"
- **WHEN** 品牌识别中间件执行
- **THEN** 中间件 SHALL 从 c.Request.Host 读取域名
- **AND** 中间件 SHALL 将域名传递给品牌管理器

#### Scenario: 中间件注入品牌到上下文
- **GIVEN** 品牌管理器返回 "kapon" 品牌对象
- **WHEN** 品牌识别中间件执行
- **THEN** 中间件 SHALL 调用 c.Set("brand", brand) 注入品牌
- **AND** 后续处理器 SHALL 能够通过 c.Get("brand") 读取品牌

### Requirement 2.2: API 接口返回品牌信息

系统 SHALL 在 /api/status 接口中返回当前请求对应的品牌信息。

#### Scenario: 返回品牌信息
- **GIVEN** 请求通过域名 "models.kapon.cloud" 访问 /api/status
- **AND** 请求上下文中包含 "kapon" 品牌对象
- **WHEN** GetStatus 处理器执行
- **THEN** 处理器 SHALL 从上下文读取品牌对象
- **AND** 处理器 SHALL 构建品牌信息 JSON
- **AND** 响应 SHALL 包含以下品牌字段：
  - brand_name: "kapon"
  - brand_logo: "/brands/kapon/logo.png"
  - brand_favicon: "/brands/kapon/favicon.ico"
  - system_name: "Kapon AI"
  - description: "Kapon AI 为中小企业..."
  - keywords: "AI API,大模型API..."
  - author: "Kapon AI"

#### Scenario: 向后兼容单品牌模式
- **GIVEN** 系统未配置多品牌
- **AND** 请求上下文中不包含品牌对象
- **WHEN** GetStatus 处理器执行
- **THEN** 处理器 SHALL 使用全局配置
- **AND** 响应 SHALL 包含 system_name 和 logo 字段
- **AND** 响应 SHALL 不包含 brand_name 等新字段

## ADDED Requirements

### Requirement 3: 前端品牌资源管理

系统 SHALL 支持按品牌组织静态资源，并根据品牌配置动态加载对应资源。

#### Scenario: 按品牌组织静态资源
- **GIVEN** 系统有两个品牌（kapon 和 grouplay）
- **WHEN** 组织静态资源目录
- **THEN** 系统 SHALL 创建 web/public/brands/ 目录
- **AND** 系统 SHALL 为每个品牌创建子目录（brands/kapon/, brands/grouplay/）
- **AND** 每个品牌目录 SHALL 包含 logo.png 和 favicon.ico 文件

### Requirement 3.1: 品牌资源目录结构

系统 SHALL 定义清晰的品牌资源目录结构。

#### Scenario: 品牌资源目录结构
- **GIVEN** 系统需要存储品牌资源
- **WHEN** 组织目录结构
- **THEN** 目录结构 SHALL 如下：
  ```
  web/public/
  ├── brands/
  │   ├── kapon/
  │   │   ├── logo.png
  │   │   └── favicon.ico
  │   └── grouplay/
  │       ├── logo.png
  │       └── favicon.ico
  └── logo.png  # 默认 logo（向后兼容）
  ```

### Requirement 3.2: Logo 组件动态加载

前端 Logo 组件 SHALL 根据品牌配置动态加载对应的 Logo 图片。

#### Scenario: 加载品牌 Logo
- **GIVEN** 前端从 /api/status 接收到品牌配置
- **AND** 品牌配置包含 brand_logo: "/brands/kapon/logo.png"
- **WHEN** Logo 组件渲染
- **THEN** 组件 SHALL 从 Redux store 读取 siteInfo.brand_logo
- **AND** 组件 SHALL 渲染 <img src="/brands/kapon/logo.png" />
- **AND** 组件 SHALL 使用 system_name 作为 alt 文本

#### Scenario: 回退到默认 Logo
- **GIVEN** 前端从 /api/status 接收到配置
- **AND** 配置不包含 brand_logo 字段（单品牌模式）
- **WHEN** Logo 组件渲染
- **THEN** 组件 SHALL 尝试使用 siteInfo.logo
- **AND** 如果 logo 也不存在，组件 SHALL 使用默认路径 "/logo.png"

#### Scenario: 处理加载状态
- **GIVEN** 前端正在加载配置
- **AND** siteInfo.isLoading 为 true
- **WHEN** Logo 组件渲染
- **THEN** 组件 SHALL 返回 null（不显示任何内容）

### Requirement 3.3: 动态 Meta 信息（可选）

前端 SHALL 支持根据品牌配置动态更新 HTML meta 信息。

#### Scenario: 更新页面标题
- **GIVEN** 前端从 /api/status 接收到品牌配置
- **AND** 品牌配置包含 system_name: "Kapon AI"
- **WHEN** 配置加载完成
- **THEN** 前端 SHALL 更新 document.title 为 "Kapon AI"

#### Scenario: 更新 meta 描述和关键词
- **GIVEN** 前端从 /api/status 接收到品牌配置
- **AND** 品牌配置包含 description 和 keywords
- **WHEN** 配置加载完成
- **THEN** 前端 SHALL 更新 meta description 标签
- **AND** 前端 SHALL 更新 meta keywords 标签

## ADDED Requirements

### Requirement 4: 向后兼容性

系统 SHALL 保持向后兼容，未配置多品牌时使用原有单品牌逻辑。

#### Scenario: 单品牌模式运行
- **GIVEN** 配置文件不包含 brands 配置段
- **WHEN** 系统启动并运行
- **THEN** 系统 SHALL 使用全局配置（SystemName, Logo）
- **AND** /api/status 接口 SHALL 返回 system_name 和 logo 字段
- **AND** 前端 SHALL 正常显示 Logo 和系统名称
- **AND** 所有功能 SHALL 正常工作

#### Scenario: 从多品牌回退到单品牌
- **GIVEN** 系统之前配置了多品牌
- **AND** 管理员删除了 brands 配置段
- **WHEN** 系统重启
- **THEN** 系统 SHALL 自动回退到单品牌模式
- **AND** 系统 SHALL 使用全局配置
- **AND** 系统 SHALL 正常运行

### Requirement 4.1: 配置兼容性

系统 SHALL 兼容现有配置，不强制要求配置多品牌。

#### Scenario: 现有配置继续有效
- **GIVEN** 用户使用现有配置文件（不包含 brands 配置）
- **WHEN** 升级到支持多品牌的版本
- **THEN** 系统 SHALL 正常启动
- **AND** 系统 SHALL 使用现有的 SystemName 和 Logo 配置
- **AND** 用户 SHALL 无需修改配置文件

### Requirement 4.2: 错误处理与降级

系统 SHALL 在品牌配置错误时优雅降级，不影响核心功能。

#### Scenario: 品牌资源文件不存在
- **GIVEN** 品牌配置指定 logo: "/brands/kapon/logo.png"
- **AND** 该文件不存在于服务器
- **WHEN** 前端尝试加载 Logo
- **THEN** 浏览器 SHALL 显示 404 错误（不影响页面功能）
- **AND** 前端 SHALL 显示默认图片或占位符
- **AND** 系统其他功能 SHALL 正常工作

#### Scenario: 品牌配置格式错误
- **GIVEN** config.yaml 中的 brands 配置格式不正确
- **WHEN** 系统启动
- **THEN** 系统 SHALL 记录错误日志
- **AND** 系统 SHALL 回退到单品牌模式
- **AND** 系统 SHALL 继续正常运行
- **AND** 管理员 SHALL 能够通过日志发现配置错误
