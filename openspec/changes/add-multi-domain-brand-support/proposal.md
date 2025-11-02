# 多域名多品牌支持

## Why

当前系统仅支持单一品牌配置（系统名称、Logo、Favicon 等），无法满足以下业务需求：

1. **多品牌运营需求**：需要在同一套系统上运营多个品牌（如 Kapon AI 和 Grouplay AI）
2. **域名隔离需求**：不同品牌使用不同域名访问（models.kapon.cloud 和 model.grouplay.cn）
3. **品牌差异化需求**：
   - **前端门户**：不同品牌需要完全不同的 UI 设计、布局、文案和用户体验
   - **管理后台**：只需品牌 logo 和标识差异化，其他 UI 保持一致
4. **SEO 优化需求**：每个域名需要独立的 meta 信息（title、description、keywords）
5. **独立前端需求**：支持每个品牌使用独立的前端项目，甚至不同的技术栈

当前实现的局限性：
- 系统名称、Logo 等配置是全局单一的，存储在 `options` 表中
- 前端是单一项目，无法支持完全不同的 UI 设计
- 缺少基于域名的前端资源路由机制
- 管理后台无法根据品牌动态显示标识
- 缺少域名到品牌的映射机制

## What Changes

### 核心变更

1. **后端品牌配置系统**
   - 新增品牌数据库表（存储品牌配置）
   - 新增品牌数据模型和管理器（内存缓存 + 数据库持久化）
   - 新增域名识别中间件（根据 Host 识别品牌）
   - 修改 `/api/status` 接口返回品牌相关配置
   - 提供品牌管理 RESTful API（增删改查）
   - 支持配置前端类型（embedded 或 external）
   - 支持动态刷新品牌缓存

2. **多前端架构支持**
   - 支持每个品牌使用独立的前端项目（完全不同的 UI）
   - 基于域名的前端资源路由（自动识别并返回对应品牌的前端）
   - 支持前端统一部署或独立部署两种模式
   - 管理后台通过 `/panel` 路径访问，只需品牌标识差异化

3. **前端项目组织**
   - 每个品牌一个独立的前端项目目录
   - 独立的 package.json、构建配置和依赖
   - 构建产物输出到品牌专属目录
   - 支持不同品牌使用不同技术栈

4. **品牌管理界面**
   - 在管理后台新增品牌管理页面
   - 支持可视化添加、编辑、删除品牌
   - 支持配置域名、前端路径、品牌标识等
   - 提供表单验证和实时反馈
   - 支持启用/禁用品牌、设置默认品牌
   - 无需修改配置文件，动态生效

### 具体变更内容

#### 后端变更（新增）
- `model/brand.go`：品牌数据模型和数据库操作
- `middleware/brand.go`：品牌识别中间件
- `controller/brand.go`：品牌管理 API 控制器
- `router/frontend.go`：前端资源路由处理器
- 数据库迁移文件：创建 brands 表

#### 后端变更（修改）
- `controller/misc.go`：修改 GetStatus 返回品牌信息
- `router/main.go`：注册品牌中间件、品牌 API 和前端路由

#### 前端变更（新增）
- `web/kapon-portal/`：Kapon 品牌独立前端项目
- `web/grouplay-portal/`：Grouplay 品牌独立前端项目
- `web/public/brands/kapon/`：Kapon 前端构建产物目录
- `web/public/brands/grouplay/`：Grouplay 前端构建产物目录
- `web/admin/`：管理后台项目（共享，支持品牌标识动态显示）

#### 前端变更（修改）
- `web/admin/src/components/Logo.jsx`：管理后台 Logo 组件支持动态品牌 Logo
- 构建脚本：新增多前端项目构建支持

#### 前端变更（管理后台）
- `web/admin/src/pages/BrandManagement/`：新增品牌管理页面
  - `BrandList.jsx`：品牌列表页面
  - `BrandForm.jsx`：添加/编辑品牌表单
  - `BrandAPI.js`：品牌 API 调用封装
- `web/admin/src/components/Logo.jsx`：修改支持动态品牌 logo
- 路由配置：新增品牌管理路由

### 向后兼容性

- **完全向后兼容**：未配置多品牌时，系统保持原有单品牌行为
- 默认品牌：如果未匹配到任何品牌，使用原有的全局配置（SystemName、Logo 等）
- 数据库兼容：不需要修改现有数据库结构（品牌配置存储在配置文件或新表中）

## Impact

### 影响的模块

- **后端模块**：
  - `model/brand.go`：新增品牌模型和数据库操作
  - `middleware/brand.go`：新增品牌识别中间件
  - `controller/brand.go`：新增品牌管理 API
  - `controller/misc.go`：修改状态接口
  - `router/frontend.go`：新增前端资源路由
  - `router/main.go`：注册中间件和路由
  - 数据库迁移：新增 brands 表

- **前端模块**：
  - `web/kapon-portal/`：新增 Kapon 独立前端项目
  - `web/grouplay-portal/`：新增 Grouplay 独立前端项目
  - `web/admin/src/pages/BrandManagement/`：新增品牌管理页面
  - `web/admin/src/components/Logo.jsx`：修改支持动态品牌
  - `web/public/brands/`：品牌前端构建产物目录

- **数据库**：
  - 新增 `brands` 表
  - 用户需要执行数据库迁移

### 影响的代码文件

**后端（新增）**：
- `model/brand.go`：品牌数据模型和数据库操作
- `middleware/brand.go`：品牌识别中间件
- `controller/brand.go`：品牌管理 API 控制器
- `router/frontend.go`：前端资源路由处理器
- `migrations/xxx_create_brands_table.sql`：数据库迁移文件

**后端（修改）**：
- `controller/misc.go`：修改 GetStatus 返回品牌信息
- `router/main.go`：注册品牌中间件、品牌 API 和前端路由

**前端项目（新增）**：
- `web/kapon-portal/`：Kapon 品牌独立前端项目（完整项目结构）
- `web/grouplay-portal/`：Grouplay 品牌独立前端项目（完整项目结构）
- `web/admin/`：管理后台项目（如果之前不存在）

**前端构建产物（新增）**：
- `web/public/brands/kapon/`：Kapon 前端构建产物
- `web/public/brands/grouplay/`：Grouplay 前端构建产物

**前端（管理后台新增）**：
- `web/admin/src/pages/BrandManagement/BrandList.jsx`：品牌列表页面
- `web/admin/src/pages/BrandManagement/BrandForm.jsx`：品牌表单页面
- `web/admin/src/api/brandAPI.js`：品牌 API 封装

**前端（修改）**：
- `web/admin/src/components/Logo.jsx`：管理后台 Logo 组件支持动态品牌
- `web/admin/src/routes/index.js`：新增品牌管理路由
- `package.json`：新增多前端构建脚本

### 部署影响

- **数据库迁移**：需要执行数据库迁移创建 brands 表
- **品牌配置**：通过管理后台界面配置品牌，无需修改配置文件
- **前端构建**：需要分别构建各品牌的前端项目
- **静态资源**：
  - 需要准备各品牌的前端构建产物
  - 需要将品牌资源（logo、favicon）放置到指定目录（/brands/{brand_name}/）
- **DNS 配置**：需要将多个域名指向同一服务器（或独立部署前端）
- **Nginx 配置**：
  - 需要配置多域名支持（传递 Host 头）
  - 如果前端统一部署，需要配置前端资源路由规则
  - 如果前端独立部署，需要配置 CORS

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
