# 多域名多品牌支持

## Why

当前系统仅支持单一品牌配置（系统名称、Logo、Favicon 等），无法满足以下业务需求：

1. **多品牌运营需求**：需要在同一套系统上运营多个品牌（如 Kapon AI 和 Grouplay AI）
2. **域名隔离需求**：不同品牌使用不同域名访问（models.kapon.cloud 和 model.grouplay.cn）
3. **品牌差异化需求**：不同域名访问时需要显示对应品牌的 Logo、系统名称、描述等信息
4. **SEO 优化需求**：每个域名需要独立的 meta 信息（title、description、keywords）

当前实现的局限性：
- 系统名称、Logo 等配置是全局单一的，存储在 `options` 表中
- 前端 Logo 组件硬编码使用 `/logo.png`
- HTML meta 信息是静态的，无法根据域名动态调整
- 缺少域名到品牌的映射机制

## What Changes

### 核心变更

1. **后端品牌配置系统**
   - 新增品牌配置数据模型（支持多品牌配置）
   - 新增域名识别中间件（根据 Host 识别品牌）
   - 修改 `/api/status` 接口返回品牌相关配置
   - 支持通过配置文件或数据库管理品牌

2. **前端品牌资源管理**
   - 重组静态资源目录结构（按品牌组织）
   - 修改 Logo 组件支持动态加载品牌 Logo
   - 修改 HTML 模板支持动态 meta 信息
   - 前端根据后端返回的品牌配置渲染对应资源

3. **配置文件扩展**
   - 在 `config.yaml` 中新增 `brands` 配置段
   - 支持配置多个品牌及其域名映射关系
   - 每个品牌可配置独立的系统名称、Logo、Favicon、描述等

### 具体变更内容

#### 后端变更
- `model/brand.go`：新增品牌数据模型
- `middleware/brand.go`：新增品牌识别中间件
- `controller/misc.go`：修改 `GetStatus` 方法返回品牌配置
- `common/config/brand.go`：新增品牌配置管理
- `config.example.yaml`：新增品牌配置示例

#### 前端变更
- `web/public/brands/`：新增品牌资源目录
- `web/src/ui-component/Logo.jsx`：修改支持动态品牌 Logo
- `web/index.html`：修改支持动态 meta 信息（可选）
- `web/src/config.js`：新增品牌相关配置字段

#### 配置变更
- 新增 `brands` 配置段，支持配置多个品牌
- 每个品牌包含：name、domains、system_name、logo、favicon、description 等字段

### 向后兼容性

- **完全向后兼容**：未配置多品牌时，系统保持原有单品牌行为
- 默认品牌：如果未匹配到任何品牌，使用原有的全局配置（SystemName、Logo 等）
- 数据库兼容：不需要修改现有数据库结构（品牌配置存储在配置文件或新表中）

## Impact

### 影响的模块

- **后端模块**：
  - `model/`：新增品牌模型
  - `middleware/`：新增品牌识别中间件
  - `controller/misc.go`：修改状态接口
  - `common/config/`：新增品牌配置管理
  - `router/main.go`：注册品牌中间件

- **前端模块**：
  - `web/public/`：重组静态资源
  - `web/src/ui-component/Logo.jsx`：修改 Logo 组件
  - `web/src/store/siteInfoReducer.js`：可能需要新增品牌字段
  - `web/index.html`：可能需要支持动态 meta

- **配置文件**：
  - `config.example.yaml`：新增品牌配置示例
  - 用户需要更新自己的配置文件

### 影响的代码文件

**后端（新增）**：
- `model/brand.go`
- `middleware/brand.go`
- `common/config/brand.go`

**后端（修改）**：
- `controller/misc.go`
- `router/main.go`
- `config.example.yaml`

**前端（新增）**：
- `web/public/brands/kapon/logo.png`
- `web/public/brands/kapon/favicon.ico`
- `web/public/brands/grouplay/logo.png`
- `web/public/brands/grouplay/favicon.ico`

**前端（修改）**：
- `web/src/ui-component/Logo.jsx`
- `web/src/config.js`
- `web/index.html`（可选）

### 部署影响

- **配置更新**：需要在配置文件中添加品牌配置
- **静态资源**：需要准备各品牌的 Logo 和 Favicon
- **DNS 配置**：需要将多个域名指向同一服务器
- **Nginx 配置**：需要配置多域名支持（传递 Host 头）
- **无需数据库迁移**：不涉及数据库结构变更

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
