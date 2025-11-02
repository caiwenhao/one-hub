# 多域名多品牌支持（简化版）

## Why

当前系统仅支持单一品牌配置（系统名称、Logo、Favicon 等），无法满足以下业务需求：

1. **多品牌运营需求**：需要在同一套系统上运营多个品牌（Kapon AI 为默认品牌，Grouplay AI 为新增品牌）
2. **域名隔离需求**：
   - Kapon AI (默认品牌): models.kapon.cloud - 保持原有部署模式
   - Grouplay AI (新品牌): model.grouplay.cn - 独立前端部署
3. **品牌标识差异化**：不同品牌需要显示不同的 Logo、系统名称等标识信息
4. **统一管理后台**：所有品牌通过 /panel 路径访问统一的管理后台

当前实现的局限性：
- 系统名称、Logo 等配置是全局单一的，存储在 `options` 表中
- 缺少基于域名的品牌识别机制
- 缺少品牌配置的管理功能
- 无法支持新品牌的独立前端部署

## What Changes

### 核心变更（简化版）

1. **后端品牌配置系统**（最小化修改）
   - 新增品牌数据库表（存储品牌配置）
   - 新增品牌数据模型和管理器（内存缓存 + 数据库持久化）
   - 新增域名识别中间件（根据 Host 识别品牌）
   - 修改 `/api/status` 接口返回品牌相关配置
   - 提供品牌管理 RESTful API（增删改查）
   - 支持动态刷新品牌缓存
   - Kapon AI 作为默认品牌（models.kapon.cloud）

2. **品牌管理界面**
   - 在管理后台新增品牌管理页面
   - 支持可视化添加、编辑、删除品牌
   - 支持配置域名和品牌标识（Logo、系统名称等）
   - 提供表单验证和实时反馈
   - 支持启用/禁用品牌、设置默认品牌

3. **前端部署模式**
   - **Kapon AI（默认品牌）**：保持原有部署模式，不做任何改动
   - **Grouplay AI（新品牌）**：
     - 基于 Kapon AI 前端代码克隆
     - 移除 panel 相关代码（管理后台）
     - 独立构建和部署
     - 通过 Nginx 反向代理 /api/* 到后端
     - 通过 Nginx 反向代理 /panel 到 Kapon AI 的管理后台
   - 前端通过 `/api/status` 获取品牌配置
   - 前端调整 Logo 等品牌标识的读取方式

### 具体变更内容

#### 后端变更（新增）
- `model/brand.go`：品牌数据模型和数据库操作
- `middleware/brand.go`：品牌识别中间件
- `controller/brand.go`：品牌管理 API 控制器
- 数据库迁移文件：创建 brands 表
- 初始化数据：创建 Kapon AI 默认品牌记录

#### 后端变更（修改）
- `controller/misc.go`：修改 GetStatus 返回品牌信息
- `router/main.go`：注册品牌中间件和品牌 API

#### 前端变更（Kapon AI - 默认品牌）
- **保持原有代码不变**
- 可选：`web/default/src/components/Logo.jsx`：修改支持从 `/api/status` 读取品牌 Logo

#### 前端变更（Grouplay AI - 新品牌）
- 创建 `web/grouplay/`：基于 Kapon AI 前端代码完整克隆
- 只移除 panel 相关代码和路由（管理后台）
- 保持其他代码和 UI 文案完全一致
- 清理仅用于管理后台的依赖（可选）
- 详细步骤参考：`FRONTEND-CLONE-GUIDE.md`

#### 前端变更（管理后台新增）
- `web/admin/src/pages/BrandManagement/`：新增品牌管理页面
  - `BrandList.jsx`：品牌列表页面
  - `BrandForm.jsx`：添加/编辑品牌表单
  - `BrandAPI.js`：品牌 API 调用封装
- 路由配置：新增品牌管理路由

### 部署架构

```
┌─────────────────────────────────────────────────────────────┐
│                    Kapon AI（默认品牌）                        │
│                  models.kapon.cloud                          │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  前端静态资源（原有部署，保持不变）                      │   │
│  │  - 用户门户                                            │   │
│  │  - 管理后台 (/panel)                                   │   │
│  └──────────────────────────────────────────────────────┘   │
│                         │                                    │
│                         ▼                                    │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  后端服务（品牌识别中间件）                             │   │
│  │  - 根据 Host 头识别品牌                                │   │
│  │  - 返回对应品牌配置                                    │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                   Grouplay AI（新品牌）                        │
│                  model.grouplay.cn                           │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  Nginx 反向代理                                        │   │
│  │  - 前端静态资源（基于 Kapon AI 克隆，移除 panel）       │   │
│  │  - /api/* → 反向代理到 Kapon AI 后端                  │   │
│  │  - /panel → 反向代理到 Kapon AI 管理后台              │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

### 向后兼容性

- **完全向后兼容**：Kapon AI 保持原有部署和代码不变
- **默认品牌**：Kapon AI 作为默认品牌（models.kapon.cloud）
- **数据库兼容**：新增 brands 表，不影响现有表结构
- **渐进式升级**：可以先部署后端品牌管理功能，后续再添加新品牌

## Impact

### 影响的模块

- **后端模块**：
  - `model/brand.go`：新增品牌模型和数据库操作
  - `middleware/brand.go`：新增品牌识别中间件
  - `controller/brand.go`：新增品牌管理 API
  - `controller/misc.go`：修改状态接口
  - `router/main.go`：注册中间件和路由
  - 数据库迁移：新增 brands 表

- **前端模块**（管理后台）：
  - `web/admin/src/pages/BrandManagement/`：新增品牌管理页面
  - `web/admin/src/components/Logo.jsx`：可选修改，支持动态品牌

- **数据库**：
  - 新增 `brands` 表
  - 用户需要执行数据库迁移

### 影响的代码文件

**后端（新增）**：
- `model/brand.go`：品牌数据模型和数据库操作
- `middleware/brand.go`：品牌识别中间件
- `controller/brand.go`：品牌管理 API 控制器
- `migrations/xxx_create_brands_table.sql`：数据库迁移文件

**后端（修改）**：
- `controller/misc.go`：修改 GetStatus 返回品牌信息
- `router/main.go`：注册品牌中间件和品牌 API

**前端（管理后台新增）**：
- `web/admin/src/pages/BrandManagement/BrandList.jsx`：品牌列表页面
- `web/admin/src/pages/BrandManagement/BrandForm.jsx`：品牌表单页面
- `web/admin/src/api/brandAPI.js`：品牌 API 封装
- `web/admin/src/routes/index.js`：新增品牌管理路由

**前端（可选修改）**：
- `web/admin/src/components/Logo.jsx`：可选修改，支持动态品牌 Logo

### 部署影响

- **数据库迁移**：
  - 执行数据库迁移创建 brands 表
  - 初始化 Kapon AI 默认品牌记录
  
- **品牌配置**：通过管理后台界面配置品牌，无需修改配置文件

- **前端部署**：
  - **Kapon AI（默认品牌）**：保持原有部署方式，无需改动
  - **Grouplay AI（新品牌）**：
    - 克隆 Kapon AI 前端代码
    - 移除 panel 相关代码
    - 独立构建和部署到新服务器
    - 配置 Nginx 反向代理
    
- **DNS 配置**：
  - models.kapon.cloud → Kapon AI 服务器（保持不变）
  - model.grouplay.cn → Grouplay AI 前端服务器（新增）
  
- **Nginx 配置**：
  - **Kapon AI**：保持原有配置
  - **Grouplay AI**：
    - 配置静态资源服务
    - 配置 /api/* 反向代理到后端（传递 Host 头）
    - 配置 /panel 反向代理到 Kapon AI 管理后台

### 性能影响

- **最小性能影响**：品牌配置在启动时加载到内存，运行时仅需一次 map 查找
- **缓存策略**：品牌配置可缓存，避免重复查询
- **无额外数据库查询**：品牌配置从内存读取

### 用户影响

- **管理员**：需要配置品牌信息和域名映射
- **最终用户**：透明无感知，访问不同域名看到不同品牌
- **API 用户**：无影响，API 功能不变

### 风险评估

- **低风险**：功能是增量的，不影响现有功能
- **向后兼容**：未配置多品牌时保持原有行为
- **测试覆盖**：需要测试多域名访问场景
- **回滚方案**：删除品牌配置即可回退到单品牌模式
