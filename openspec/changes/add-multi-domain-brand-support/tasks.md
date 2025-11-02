# 实现任务清单

## 1. 数据库表和数据模型
- [ ] 1.1 创建品牌数据库表
  - 创建数据库迁移文件（根据项目使用的迁移工具）
  - 定义 `brands` 表结构（id, name, domains, system_name, logo, favicon, description, keywords, author, is_default, enabled, created_at, updated_at）
  - 设置 name 字段为唯一索引
  - 设置合适的字段类型和约束
  - 执行迁移创建表
  - 初始化 Kapon AI 默认品牌记录（name: kapon, domains: ["models.kapon.cloud"], is_default: true）
  - _Requirements: Brand Management - Requirement 1.1_

- [ ] 1.2 创建品牌数据模型
  - 创建 `model/brand.go` 文件
  - 定义 `Brand` 结构体，映射数据库表字段
  - 实现 CRUD 方法：Create, GetByID, GetAll, Update, Delete
  - 实现 GetByDomain(domain string) 方法（查询 domains 字段包含指定域名的品牌）
  - 实现 GetDefault() 方法（查询 is_default = true 的品牌）
  - 实现 GetEnabled() 方法（查询 enabled = true 的品牌）
  - 实现 Toggle(id int) 方法（切换启用状态）
  - 实现 SetDefault(id int) 方法（设置默认品牌，同时取消其他品牌的默认状态）
  - domains 字段使用 JSON 存储，提供序列化/反序列化方法
  - _Requirements: Brand Management - Requirement 1.2_

- [ ] 1.3 实现品牌管理器（缓存层）
  - 创建 `common/brand_manager.go` 文件
  - 定义 `BrandManager` 结构体（brands map, domainMap, defaultBrand, mutex）
  - 实现 LoadFromDatabase() 方法（从数据库加载所有启用的品牌到内存）
  - 实现 GetBrandByDomain(domain string) 方法（从内存缓存查询）
  - 实现 GetDefaultBrand() 方法（返回默认品牌或全局配置）
  - 实现 RefreshCache() 方法（重新从数据库加载品牌）
  - 实现域名到品牌的映射逻辑（支持端口）
  - 在 `main.go` 启动时调用 LoadFromDatabase()
  - 添加日志记录（启动时输出品牌配置信息）
  - _Requirements: Brand Management - Requirement 1, 1.3_

- [ ]* 1.4 编写数据模型和管理器单元测试
  - 创建 `model/brand_test.go` 文件
  - 测试 CRUD 操作
  - 测试 GetByDomain 方法
  - 测试 SetDefault 方法（确保唯一性）
  - 创建 `common/brand_manager_test.go` 文件
  - 测试 GetBrandByDomain 方法（精确匹配、端口匹配、未匹配）
  - 测试 RefreshCache 方法
  - 测试默认品牌逻辑
  - _Requirements: Brand Management - Requirement 1.1, 1.2, 1.3_

## 2. 品牌识别中间件
- [ ] 2.1 创建品牌识别中间件
  - 创建 `middleware/brand.go` 文件
  - 实现 `BrandMiddleware()` 函数
  - 从 `c.Request.Host` 读取域名
  - 调用 `BrandManager.GetBrandByDomain()` 获取品牌
  - 将品牌对象注入到 `c.Set("brand", brand)`
  - 处理未匹配品牌的情况（使用默认品牌）
  - _Requirements: Brand Management - Requirement 2.1_

- [ ] 2.2 注册品牌中间件到路由
  - 修改 `router/main.go` 的 `SetRouter` 函数
  - 在日志中间件之后、认证中间件之前注册品牌中间件
  - 确保所有路由都经过品牌中间件
  - _Requirements: Brand Management - Requirement 2.1_

- [ ]* 2.3 编写中间件单元测试
  - 创建 `middleware/brand_test.go` 文件
  - 测试正常品牌识别流程
  - 测试未知域名处理
  - 测试 Context 注入
  - _Requirements: Brand Management - Requirement 2.1_

## 3. 品牌管理 API
- [ ] 3.1 创建品牌管理控制器
  - 创建 `controller/brand.go` 文件
  - 实现 GetBrands() 方法（GET /api/brands）- 获取品牌列表
  - 实现 GetBrand() 方法（GET /api/brands/:id）- 获取单个品牌
  - 实现 CreateBrand() 方法（POST /api/brands）- 创建品牌
  - 实现 UpdateBrand() 方法（PUT /api/brands/:id）- 更新品牌
  - 实现 DeleteBrand() 方法（DELETE /api/brands/:id）- 删除品牌
  - 实现 ToggleBrand() 方法（PATCH /api/brands/:id/toggle）- 启用/禁用品牌
  - 实现 SetDefaultBrand() 方法（PATCH /api/brands/:id/set-default）- 设置默认品牌
  - 实现 RefreshBrands() 方法（POST /api/brands/refresh）- 刷新缓存
  - _Requirements: Brand Management - Requirement 6_

- [ ] 3.2 实现品牌配置验证
  - 在 `controller/brand.go` 中实现验证逻辑
  - 验证必填字段（name, domains, system_name, logo, favicon）
  - 验证品牌名称格式（小写字母、数字、连字符，2-50 字符）
  - 验证域名格式（支持带端口）
  - 验证域名唯一性（跨品牌检查）
  - 返回具体的验证错误信息
  - _Requirements: Brand Management - Requirement 8_

- [ ] 3.3 注册品牌管理路由
  - 修改 `router/main.go`
  - 在认证中间件保护下注册品牌管理路由
  - 确保只有管理员可以访问品牌管理 API
  - _Requirements: Brand Management - Requirement 6_

- [ ] 3.4 API 调用后自动刷新缓存
  - 在 CreateBrand, UpdateBrand, DeleteBrand, ToggleBrand, SetDefaultBrand 方法中
  - 数据库操作成功后调用 BrandManager.RefreshCache()
  - 确保配置更改立即生效
  - _Requirements: Brand Management - Requirement 1, 6_

- [ ]* 3.5 编写品牌 API 集成测试
  - 创建 `controller/brand_test.go` 文件
  - 测试所有 CRUD 操作
  - 测试验证逻辑（各种错误情况）
  - 测试域名唯一性检查
  - 测试默认品牌设置
  - 测试缓存刷新
  - _Requirements: Brand Management - Requirement 6, 8_

## 4. GetStatus 接口修改
- [ ] 4.1 修改 GetStatus 接口返回品牌信息
  - 修改 `controller/misc.go` 的 `GetStatus` 函数
  - 从 Context 读取品牌：`c.Get("brand")`
  - 构建品牌信息 JSON（brand_name, brand_logo, brand_favicon, system_name, description, keywords, author）
  - 合并到响应的 data 字段中
  - 保持向后兼容：未匹配品牌时使用全局配置（config.SystemName, config.Logo）
  - _Requirements: Brand Management - Requirement 2.2_

- [ ]* 4.2 编写 GetStatus 集成测试
  - 创建测试用例模拟不同 Host 头请求
  - 验证返回的品牌信息正确
  - 验证向后兼容性（无品牌配置时）
  - _Requirements: Brand Management - Requirement 2.2_

## 5. Grouplay AI 前端项目创建
- [ ] 5.1 克隆 Kapon AI 前端代码
  - 复制 Kapon AI 前端项目到 `web/grouplay/` 目录
  - 保留所有核心功能代码
  - 保留所有 API 调用逻辑
  - 保留所有 UI 文案（不做任何修改）
  - _Requirements: Brand Management - Requirement 5.4_

- [ ] 5.2 移除 panel 相关代码
  - 识别并删除管理后台相关页面和组件（通常在 `/panel` 路径下）
  - 删除管理后台路由配置（如 `/panel/*` 路由）
  - 删除管理员权限检查相关代码
  - 删除以下管理功能模块：
    - 用户管理
    - 渠道管理
    - 令牌管理
    - 兑换码管理
    - 系统设置
    - 数据统计
    - 日志查看
  - 保留以下用户功能：
    - 用户登录和注册
    - 个人资料设置
    - 密码修改
    - API Key 管理（用户自己的）
    - 使用记录查看（用户自己的）
  - **重要**：不修改任何 UI 文案，保持与 Kapon AI 完全一致
  - _Requirements: Brand Management - Requirement 5.4_

- [ ] 5.3 清理未使用的依赖
  - 检查 package.json 中的依赖
  - 移除仅用于管理后台的依赖包（如果有）
  - 保留所有用户功能相关的依赖
  - 更新 package.json
  - _Requirements: Brand Management - Requirement 5.4_

- [ ] 5.4 调整品牌标识读取（可选）
  - 如果需要动态显示品牌 Logo，修改 Logo 组件从 `/api/status` 读取品牌配置
  - 如果需要动态页面标题，修改 meta 信息读取方式
  - 如果不需要，保持原有代码不变
  - _Requirements: Brand Management - Requirement 5.4_

- [ ] 5.5 配置构建输出
  - 配置构建输出目录（如 `dist/`）
  - 测试构建流程：`npm run build`
  - 验证构建产物完整性
  - 确保构建产物可以正常运行
  - _Requirements: Brand Management - Requirement 5.4_

## 6. 品牌管理界面
- [ ] 6.1 创建品牌列表页面
  - 创建 `web/admin/src/pages/BrandManagement/BrandList.jsx`
  - 实现品牌列表展示（表格形式）
  - 显示字段：品牌名称、系统名称、域名数量、前端类型、状态（启用/禁用）、是否默认、操作按钮
  - 实现"添加品牌"按钮
  - 实现搜索和筛选功能
  - 实现启用/禁用开关（调用 toggle API）
  - 实现"设为默认"按钮（调用 set-default API）
  - 实现"编辑"和"删除"按钮
  - _Requirements: Brand Management - Requirement 7_

- [ ] 6.2 创建品牌表单页面
  - 创建 `web/admin/src/pages/BrandManagement/BrandForm.jsx`
  - 实现添加/编辑品牌表单
  - 表单字段：
    - 品牌标识（name，文本输入，必填）
    - 系统名称（system_name，文本输入，必填）
    - 关联域名（domains，标签输入，支持多个，必填）
    - Logo URL（logo，文本输入，必填，带提示）
    - Favicon URL（favicon，文本输入，必填，带提示）
    - 描述（description，文本域）
    - 关键词（keywords，文本输入）
    - 作者（author，文本输入）
    - 是否默认品牌（is_default，复选框）
    - 是否启用（enabled，复选框，默认选中）
  - 实现表单验证（前端验证 + 后端验证）
  - 实现"保存"和"取消"按钮
  - _Requirements: Brand Management - Requirement 7_

- [ ] 6.3 实现域名输入组件
  - 创建域名标签输入组件（支持添加、删除多个域名）
  - 实现域名格式验证
  - 实现域名重复检查
  - 提供友好的用户体验（Enter 键添加，点击删除）
  - _Requirements: Brand Management - Requirement 7_

- [ ] 6.4 创建品牌 API 封装
  - 创建 `web/admin/src/api/brandAPI.js`
  - 封装所有品牌管理 API 调用：
    - getBrands() - 获取品牌列表
    - getBrand(id) - 获取单个品牌
    - createBrand(data) - 创建品牌
    - updateBrand(id, data) - 更新品牌
    - deleteBrand(id) - 删除品牌
    - toggleBrand(id) - 启用/禁用品牌
    - setDefaultBrand(id) - 设置默认品牌
    - refreshBrands() - 刷新缓存
  - 统一错误处理
  - _Requirements: Brand Management - Requirement 7_

- [ ] 6.5 添加品牌管理路由
  - 修改 `web/admin/src/routes/index.js`（或对应的路由配置文件）
  - 添加品牌管理路由：
    - /brands - 品牌列表页
    - /brands/new - 添加品牌页
    - /brands/:id/edit - 编辑品牌页
  - 确保路由需要管理员权限
  - _Requirements: Brand Management - Requirement 7_

- [ ] 6.6 添加品牌管理菜单项
  - 修改管理后台侧边栏菜单配置
  - 添加"品牌管理"菜单项
  - 配置图标和链接
  - 确保只有管理员可见
  - _Requirements: Brand Management - Requirement 7_

- [ ]* 6.7 编写品牌管理界面测试
  - 测试品牌列表展示
  - 测试添加品牌流程
  - 测试编辑品牌流程
  - 测试删除品牌流程
  - 测试表单验证
  - _Requirements: Brand Management - Requirement 7_

## 7. Kapon AI 前端调整（可选）
- [ ]* 7.1 修改 Kapon AI Logo 组件
  - 修改 `web/default/src/components/Logo.jsx`（或对应路径）
  - 从 Redux store 或 Context 读取品牌信息
  - 动态显示品牌 logo 和系统名称
  - 保持其他 UI 元素不变
  - _Requirements: Brand Management - Requirement 5.2_

- [ ]* 7.2 Kapon AI 集成品牌配置
  - 在前端启动时调用 /api/status
  - 存储品牌信息到状态管理（Redux/Context）
  - 确保品牌信息在整个应用中可访问
  - _Requirements: Brand Management - Requirement 5.2_

## 8. 集成测试与验证
- [ ] 8.1 API 调用测试
  - 通过不同域名调用 /api/status，验证返回对应品牌信息
  - 验证品牌识别中间件正确识别域名
  - 验证不同品牌前端获取相同的数据
  - _Requirements: Brand Management - Requirement 2, 5.3_

- [ ] 8.2 品牌管理界面测试
  - 测试通过管理后台添加品牌
  - 测试编辑品牌配置
  - 测试删除品牌
  - 测试启用/禁用品牌
  - 测试设置默认品牌
  - 验证配置更改后立即生效（缓存刷新）
  - _Requirements: Brand Management - Requirement 6, 7_

- [ ] 8.3 向后兼容性测试
  - 数据库中无品牌配置时，验证系统正常运行
  - 验证使用全局配置（SystemName, Logo）
  - 验证前端正常显示
  - _Requirements: Brand Management - Requirement 4.1_

- [ ] 8.4 边界情况测试
  - 测试未知域名访问（使用默认品牌）
  - 测试域名冲突验证（创建品牌时）
  - 测试删除默认品牌（应该被阻止）
  - _Requirements: Brand Management - Requirement 4.2, 8_

## 9. 文档与部署
- [ ] 9.1 编写数据库迁移文档
  - 说明如何执行数据库迁移
  - 提供迁移脚本示例
  - 说明 Kapon AI 默认品牌初始化
  - 说明回滚方案
  - _Requirements: Brand Management - Requirement 1.1_

- [ ] 9.2 编写品牌管理使用文档
  - 说明如何通过管理后台添加品牌
  - 说明各字段的含义和填写规范
  - 提供配置示例和最佳实践
  - 说明默认品牌（Kapon AI）的特殊性
  - _Requirements: Brand Management - Requirement 7_

- [ ] 9.3 编写 Grouplay AI 前端部署文档
  - 说明如何克隆和精简前端代码
  - 说明如何构建 Grouplay AI 前端
  - 提供 Nginx 配置示例（静态资源 + API 反向代理 + /panel 反向代理）
  - 提供 DNS 配置说明
  - 说明前端如何从 `/api/status` 获取品牌配置
  - _Requirements: Brand Management - Requirement 5.4_

- [ ] 9.4 更新 README
  - 在 README 中添加多品牌支持说明
  - 添加架构图和管理界面截图
  - 说明 Kapon AI 作为默认品牌
  - 说明如何通过管理后台配置新品牌
  - 说明 Grouplay AI 部署流程
  - 说明向后兼容性
  - _Requirements: Brand Management - Requirement 7_
