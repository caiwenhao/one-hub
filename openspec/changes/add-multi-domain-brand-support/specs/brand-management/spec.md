# Brand Management Specification

## ADDED Requirements

### Requirement 1: 品牌数据库存储

系统 SHALL 使用数据库存储品牌配置信息，支持动态管理，并初始化 Kapon AI 作为默认品牌。

#### Scenario: 初始化默认品牌
- **GIVEN** 数据库迁移执行完成
- **WHEN** 系统首次启动
- **THEN** 系统 SHALL 创建 Kapon AI 默认品牌记录
- **AND** 默认品牌 SHALL 包含 name: "kapon"
- **AND** 默认品牌 SHALL 包含 domains: ["models.kapon.cloud"]
- **AND** 默认品牌 SHALL 设置 is_default: true
- **AND** 默认品牌 SHALL 设置 enabled: true

#### Scenario: 从数据库加载品牌配置
- **GIVEN** 数据库中存储了多个品牌配置（包括默认品牌）
- **WHEN** 系统启动时
- **THEN** 系统 SHALL 从数据库读取所有启用的品牌
- **AND** 系统 SHALL 加载品牌配置到内存缓存
- **AND** 系统 SHALL 构建域名到品牌的映射关系
- **AND** 系统 SHALL 识别 Kapon AI 为默认品牌
- **AND** 系统 SHALL 记录品牌配置加载成功的日志

#### Scenario: 数据库中无品牌配置
- **GIVEN** 数据库中没有任何品牌配置
- **WHEN** 系统启动时
- **THEN** 系统 SHALL 使用单品牌模式
- **AND** 系统 SHALL 使用全局配置（SystemName, Logo）
- **AND** 系统 SHALL 正常运行

#### Scenario: 刷新品牌配置缓存
- **GIVEN** 管理员在管理后台更新了品牌配置
- **WHEN** 调用刷新缓存 API
- **THEN** 系统 SHALL 重新从数据库加载品牌配置
- **AND** 系统 SHALL 更新内存缓存
- **AND** 系统 SHALL 重新构建域名映射
- **AND** 新的配置 SHALL 立即生效

### Requirement 1.1: 品牌数据库表结构

系统 SHALL 定义品牌数据库表，存储品牌的所有配置信息。

#### Scenario: 品牌表包含完整字段
- **GIVEN** 品牌数据库表已创建
- **WHEN** 查询表结构
- **THEN** 表 SHALL 包含以下字段：
  - id（主键，自增）
  - name（品牌唯一标识，唯一索引）
  - domains（关联域名列表，JSON 格式）
  - system_name（系统显示名称）
  - logo（Logo URL 或路径）
  - favicon（Favicon URL 或路径）
  - description（品牌描述）
  - keywords（SEO 关键词）
  - author（作者信息）
  - is_default（是否为默认品牌）
  - enabled（是否启用）
  - created_at（创建时间）
  - updated_at（更新时间）

#### Scenario: 品牌名称唯一性约束
- **GIVEN** 数据库中已存在名为 "kapon" 的品牌
- **WHEN** 尝试创建另一个名为 "kapon" 的品牌
- **THEN** 系统 SHALL 返回错误
- **AND** 错误信息 SHALL 提示品牌名称已存在

### Requirement 1.2: 品牌数据模型

系统 SHALL 定义品牌数据模型，用于内存缓存和业务逻辑。

#### Scenario: 品牌模型包含完整信息
- **GIVEN** 品牌数据模型已定义
- **WHEN** 从数据库加载品牌时
- **THEN** 品牌实例 SHALL 包含所有数据库字段
- **AND** domains 字段 SHALL 从 JSON 解析为字符串数组
- **AND** 品牌实例 SHALL 提供便捷的访问方法

### Requirement 1.3: 品牌管理器

系统 SHALL 提供品牌管理器，负责品牌配置的缓存、查询和刷新。

#### Scenario: 根据域名查询品牌
- **GIVEN** 品牌管理器已从数据库加载多个品牌配置
- **AND** 域名 "models.kapon.cloud" 关联到 "kapon" 品牌
- **WHEN** 调用 GetBrandByDomain("models.kapon.cloud")
- **THEN** 系统 SHALL 从内存缓存返回 "kapon" 品牌对象
- **AND** 品牌对象 SHALL 包含完整的品牌信息

#### Scenario: 查询未知域名
- **GIVEN** 品牌管理器已加载多个品牌配置
- **AND** 域名 "unknown.example.com" 未关联到任何品牌
- **WHEN** 调用 GetBrandByDomain("unknown.example.com")
- **THEN** 系统 SHALL 返回默认品牌对象
- **AND** 默认品牌 SHALL 是标记为 is_default 的品牌或全局配置

#### Scenario: 支持带端口的域名匹配
- **GIVEN** 品牌管理器已加载品牌配置
- **AND** 域名 "localhost:3000" 关联到 "kapon" 品牌
- **WHEN** 调用 GetBrandByDomain("localhost:3000")
- **THEN** 系统 SHALL 返回 "kapon" 品牌对象

#### Scenario: 刷新品牌缓存
- **GIVEN** 品牌管理器已加载品牌配置到内存
- **WHEN** 调用 RefreshBrands() 方法
- **THEN** 系统 SHALL 重新从数据库加载所有启用的品牌
- **AND** 系统 SHALL 更新内存缓存
- **AND** 系统 SHALL 重新构建域名映射关系

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

### Requirement 5: 多前端支持

系统 SHALL 支持为不同品牌配置独立的前端项目，每个品牌的前端独立构建和部署。

#### Scenario: 独立前端部署
- **GIVEN** 每个品牌有独立的前端项目
- **WHEN** 前端独立构建和部署
- **THEN** 前端 SHALL 通过 Nginx 反向代理到同一后端
- **AND** 前端 SHALL 通过 `/api/status` 获取品牌配置
- **AND** 前端 SHALL 根据品牌配置显示对应的 Logo 和标识

### Requirement 5.2: 管理后台品牌标识支持（可选）

管理后台 MAY 支持根据访问域名显示对应品牌的 Logo 和标识。

#### Scenario: 管理后台显示品牌标识
- **GIVEN** 管理员通过 "models.kapon.cloud" 访问管理后台
- **AND** 请求上下文包含 kapon 品牌信息
- **WHEN** 管理后台加载
- **THEN** 管理后台 MAY 从 /api/status 获取品牌信息
- **AND** 管理后台 MAY 显示 kapon 品牌的 logo
- **AND** 管理后台的其他 UI 元素 SHALL 保持一致

#### Scenario: 不同品牌管理员看到相同数据
- **GIVEN** 管理员 A 通过 "models.kapon.cloud" 登录
- **AND** 管理员 B 通过 "model.grouplay.cn" 登录
- **WHEN** 两个管理员查看用户列表
- **THEN** 两个管理员 SHALL 看到相同的用户数据
- **AND** 数据 SHALL 不按品牌隔离

### Requirement 5.3: API 路径保持一致

系统 SHALL 确保所有品牌的前端调用相同的 API 路径，不需要在请求中额外标识品牌。

#### Scenario: 前端调用 API 不需要品牌标识
- **GIVEN** kapon 前端通过域名 "models.kapon.cloud" 运行
- **WHEN** 前端调用 "/api/user/info"
- **THEN** 请求 SHALL 自动包含 Host 头 "models.kapon.cloud"
- **AND** 后端品牌识别中间件 SHALL 自动识别品牌为 kapon
- **AND** 前端代码 SHALL 不需要在请求中添加品牌参数

#### Scenario: 不同品牌前端调用相同 API
- **GIVEN** kapon 前端调用 "/api/models/list"
- **AND** grouplay 前端调用 "/api/models/list"
- **WHEN** 两个请求到达服务器
- **THEN** 两个请求 SHALL 返回相同的模型列表数据
- **AND** API 逻辑 SHALL 不区分品牌
- **AND** 数据 SHALL 在所有品牌间共享

### Requirement 5.4: 前端项目独立性

系统 SHALL 支持 Kapon AI 保持原有部署，Grouplay AI 基于 Kapon AI 克隆并独立部署。

#### Scenario: Kapon AI 保持原有部署
- **GIVEN** Kapon AI 是默认品牌
- **WHEN** 系统部署时
- **THEN** Kapon AI 前端 SHALL 保持原有部署方式
- **AND** Kapon AI SHALL 不需要任何代码修改
- **AND** Kapon AI SHALL 继续提供完整功能（包括管理后台）

#### Scenario: Grouplay AI 基于 Kapon AI 克隆
- **GIVEN** 需要创建 Grouplay AI 前端
- **WHEN** 克隆 Kapon AI 前端代码
- **THEN** 系统 SHALL 复制 Kapon AI 前端代码到新目录
- **AND** 系统 SHALL 移除 panel 相关代码和路由
- **AND** 系统 SHALL 移除管理功能相关代码
- **AND** 系统 SHALL 保留用户门户核心功能
- **AND** 系统 SHALL 精简不需要的依赖和组件

#### Scenario: Grouplay AI 独立部署
- **GIVEN** Grouplay AI 前端已构建
- **WHEN** 部署到服务器
- **THEN** Nginx SHALL 提供 Grouplay AI 静态资源
- **AND** Nginx SHALL 将 `/api/*` 反向代理到 Kapon AI 后端
- **AND** Nginx SHALL 将 `/panel` 反向代理到 Kapon AI 管理后台
- **AND** Nginx SHALL 传递 Host 头到后端
- **AND** 后端 SHALL 根据 Host 头识别为 Grouplay 品牌

#### Scenario: 统一管理后台访问
- **GIVEN** Grouplay AI 用户需要访问管理后台
- **WHEN** 用户访问 model.grouplay.cn/panel
- **THEN** Nginx SHALL 反向代理到 Kapon AI 管理后台
- **AND** 管理后台 SHALL 正常显示和工作
- **AND** 管理后台 MAY 显示 Grouplay 品牌标识（如果实现了动态品牌）

### Requirement 6: 品牌管理 API

系统 SHALL 提供 RESTful API 用于品牌的增删改查操作。

#### Scenario: 获取品牌列表
- **GIVEN** 数据库中存储了多个品牌
- **WHEN** 管理员调用 GET /api/brands
- **THEN** 系统 SHALL 返回所有品牌列表
- **AND** 每个品牌 SHALL 包含完整信息
- **AND** 响应 SHALL 按创建时间倒序排列

#### Scenario: 获取单个品牌详情
- **GIVEN** 数据库中存在 ID 为 1 的品牌
- **WHEN** 管理员调用 GET /api/brands/1
- **THEN** 系统 SHALL 返回该品牌的完整信息
- **AND** 响应 SHALL 包含所有字段

#### Scenario: 创建新品牌
- **GIVEN** 管理员提供了完整的品牌信息
- **AND** 品牌名称 "newbrand" 不存在
- **WHEN** 管理员调用 POST /api/brands
- **THEN** 系统 SHALL 验证必填字段（name, domains, system_name）
- **AND** 系统 SHALL 验证品牌名称唯一性
- **AND** 系统 SHALL 验证域名格式
- **AND** 系统 SHALL 将品牌信息保存到数据库
- **AND** 系统 SHALL 自动刷新品牌缓存
- **AND** 系统 SHALL 返回创建成功的品牌信息

#### Scenario: 创建品牌时名称冲突
- **GIVEN** 数据库中已存在名为 "kapon" 的品牌
- **WHEN** 管理员尝试创建另一个名为 "kapon" 的品牌
- **THEN** 系统 SHALL 返回 400 错误
- **AND** 错误信息 SHALL 提示品牌名称已存在

#### Scenario: 更新品牌信息
- **GIVEN** 数据库中存在 ID 为 1 的品牌
- **WHEN** 管理员调用 PUT /api/brands/1 更新信息
- **THEN** 系统 SHALL 验证更新的字段
- **AND** 系统 SHALL 更新数据库记录
- **AND** 系统 SHALL 自动刷新品牌缓存
- **AND** 系统 SHALL 返回更新后的品牌信息

#### Scenario: 删除品牌
- **GIVEN** 数据库中存在 ID 为 1 的品牌
- **AND** 该品牌不是默认品牌
- **WHEN** 管理员调用 DELETE /api/brands/1
- **THEN** 系统 SHALL 从数据库删除该品牌
- **AND** 系统 SHALL 自动刷新品牌缓存
- **AND** 系统 SHALL 返回删除成功信息

#### Scenario: 禁止删除默认品牌
- **GIVEN** 数据库中存在 ID 为 1 的品牌
- **AND** 该品牌是默认品牌（is_default = true）
- **WHEN** 管理员尝试调用 DELETE /api/brands/1
- **THEN** 系统 SHALL 返回 400 错误
- **AND** 错误信息 SHALL 提示不能删除默认品牌

#### Scenario: 启用/禁用品牌
- **GIVEN** 数据库中存在 ID 为 1 的品牌
- **WHEN** 管理员调用 PATCH /api/brands/1/toggle
- **THEN** 系统 SHALL 切换品牌的 enabled 状态
- **AND** 系统 SHALL 自动刷新品牌缓存
- **AND** 系统 SHALL 返回更新后的状态

#### Scenario: 设置默认品牌
- **GIVEN** 数据库中存在 ID 为 2 的品牌
- **WHEN** 管理员调用 PATCH /api/brands/2/set-default
- **THEN** 系统 SHALL 将所有品牌的 is_default 设置为 false
- **AND** 系统 SHALL 将 ID 为 2 的品牌的 is_default 设置为 true
- **AND** 系统 SHALL 自动刷新品牌缓存
- **AND** 系统 SHALL 返回成功信息

#### Scenario: 刷新品牌缓存 API
- **GIVEN** 管理员更新了品牌配置
- **WHEN** 管理员调用 POST /api/brands/refresh
- **THEN** 系统 SHALL 重新从数据库加载所有品牌
- **AND** 系统 SHALL 更新内存缓存
- **AND** 系统 SHALL 返回刷新成功信息

### Requirement 7: 品牌管理界面

系统 SHALL 在管理后台提供品牌管理界面，支持可视化操作。

#### Scenario: 品牌列表页面
- **GIVEN** 管理员登录管理后台
- **WHEN** 管理员访问品牌管理页面
- **THEN** 页面 SHALL 显示所有品牌列表
- **AND** 每个品牌 SHALL 显示：名称、系统名称、域名数量、状态、操作按钮
- **AND** 页面 SHALL 提供"添加品牌"按钮
- **AND** 页面 SHALL 提供搜索和筛选功能

#### Scenario: 添加品牌表单
- **GIVEN** 管理员点击"添加品牌"按钮
- **WHEN** 表单页面加载
- **THEN** 表单 SHALL 包含以下字段：
  - 品牌标识（name，必填，唯一）
  - 系统名称（system_name，必填）
  - 关联域名（domains，必填，支持多个）
  - Logo URL（logo，必填）
  - Favicon URL（favicon，必填）
  - 描述（description）
  - 关键词（keywords）
  - 作者（author）
  - 是否默认品牌（is_default，复选框）
  - 是否启用（enabled，复选框，默认选中）
- **AND** 表单 SHALL 提供实时验证
- **AND** 表单 SHALL 提供"保存"和"取消"按钮

#### Scenario: 域名输入支持多个值
- **GIVEN** 管理员在添加品牌表单中
- **WHEN** 管理员输入域名字段
- **THEN** 字段 SHALL 支持输入多个域名
- **AND** 每个域名 SHALL 可以单独删除
- **AND** 字段 SHALL 提供"添加域名"按钮
- **AND** 字段 SHALL 验证域名格式



#### Scenario: 提交品牌表单
- **GIVEN** 管理员填写完整的品牌信息
- **WHEN** 管理员点击"保存"按钮
- **THEN** 系统 SHALL 验证所有必填字段
- **AND** 系统 SHALL 调用 POST /api/brands API
- **AND** 成功后 SHALL 显示成功提示
- **AND** 成功后 SHALL 跳转到品牌列表页面
- **AND** 失败时 SHALL 显示错误信息

#### Scenario: 编辑品牌
- **GIVEN** 管理员在品牌列表页面
- **WHEN** 管理员点击某个品牌的"编辑"按钮
- **THEN** 系统 SHALL 打开编辑表单
- **AND** 表单 SHALL 预填充该品牌的现有信息
- **AND** 管理员 SHALL 能够修改任何字段
- **AND** 保存时 SHALL 调用 PUT /api/brands/{id} API

#### Scenario: 删除品牌确认
- **GIVEN** 管理员在品牌列表页面
- **WHEN** 管理员点击某个品牌的"删除"按钮
- **THEN** 系统 SHALL 显示确认对话框
- **AND** 对话框 SHALL 提示删除操作不可恢复
- **WHEN** 管理员确认删除
- **THEN** 系统 SHALL 调用 DELETE /api/brands/{id} API
- **AND** 成功后 SHALL 刷新品牌列表
- **AND** 成功后 SHALL 显示删除成功提示

#### Scenario: 快速启用/禁用品牌
- **GIVEN** 管理员在品牌列表页面
- **WHEN** 管理员点击某个品牌的启用/禁用开关
- **THEN** 系统 SHALL 调用 PATCH /api/brands/{id}/toggle API
- **AND** 成功后 SHALL 更新列表中的状态显示
- **AND** 成功后 SHALL 显示操作成功提示

#### Scenario: 设置默认品牌
- **GIVEN** 管理员在品牌列表页面
- **WHEN** 管理员点击某个品牌的"设为默认"按钮
- **THEN** 系统 SHALL 调用 PATCH /api/brands/{id}/set-default API
- **AND** 成功后 SHALL 更新列表中的默认品牌标识
- **AND** 成功后 SHALL 显示操作成功提示

#### Scenario: 品牌资源 URL 提示
- **GIVEN** 管理员在添加/编辑品牌表单中
- **WHEN** 管理员查看 Logo URL 字段
- **THEN** 字段 SHALL 显示提示信息："请输入 Logo 的完整 URL 或相对路径"
- **AND** 字段 SHALL 显示示例："https://cdn.example.com/logo.png 或 /static/logo.png"
- **WHEN** 管理员查看 Favicon URL 字段
- **THEN** 字段 SHALL 显示提示信息："请输入 Favicon 的完整 URL 或相对路径"
- **AND** 字段 SHALL 显示示例："https://cdn.example.com/favicon.ico 或 /static/favicon.ico"

### Requirement 8: 品牌配置验证

系统 SHALL 在保存品牌配置时进行完整性验证。

#### Scenario: 验证必填字段
- **GIVEN** 管理员提交品牌表单
- **WHEN** 系统验证表单数据
- **THEN** 系统 SHALL 验证 name 字段不为空
- **AND** 系统 SHALL 验证 domains 字段至少包含一个域名
- **AND** 系统 SHALL 验证 system_name 字段不为空
- **AND** 失败时 SHALL 返回具体的错误信息

#### Scenario: 验证品牌名称格式
- **GIVEN** 管理员输入品牌名称
- **WHEN** 系统验证品牌名称
- **THEN** 系统 SHALL 验证名称只包含小写字母、数字和连字符
- **AND** 系统 SHALL 验证名称长度在 2-50 个字符之间
- **AND** 系统 SHALL 验证名称不以连字符开头或结尾
- **AND** 失败时 SHALL 返回格式错误提示

#### Scenario: 验证域名格式
- **GIVEN** 管理员输入域名列表
- **WHEN** 系统验证域名
- **THEN** 系统 SHALL 验证每个域名的格式正确
- **AND** 系统 SHALL 支持带端口的域名（如 localhost:3000）
- **AND** 系统 SHALL 验证域名不重复
- **AND** 失败时 SHALL 返回具体的域名错误信息

#### Scenario: 验证域名唯一性（跨品牌）
- **GIVEN** 数据库中品牌 A 已使用域名 "models.kapon.cloud"
- **WHEN** 管理员尝试为品牌 B 添加相同域名
- **THEN** 系统 SHALL 返回错误
- **AND** 错误信息 SHALL 提示该域名已被品牌 A 使用



#### Scenario: 验证默认品牌唯一性
- **GIVEN** 数据库中已存在一个默认品牌
- **WHEN** 管理员尝试将另一个品牌设置为默认
- **THEN** 系统 SHALL 自动将原默认品牌的 is_default 设置为 false
- **AND** 系统 SHALL 将新品牌的 is_default 设置为 true
- **AND** 系统 SHALL 确保始终只有一个默认品牌
