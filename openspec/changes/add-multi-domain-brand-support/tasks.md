# 实现任务清单

## 1. 数据库表和数据模型
- [ ] 1.1 创建品牌数据库表
  - 创建数据库迁移文件（根据项目使用的迁移工具）
  - 定义 `brands` 表结构（id, name, domains, system_name, logo, favicon, description, keywords, author, frontend_type, frontend_path, frontend_url, is_default, enabled, created_at, updated_at）
  - 设置 name 字段为唯一索引
  - 设置合适的字段类型和约束
  - 执行迁移创建表
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
  - 验证必填字段（name, domains, system_name）
  - 验证品牌名称格式（小写字母、数字、连字符，2-50 字符）
  - 验证域名格式（支持带端口）
  - 验证域名唯一性（跨品牌检查）
  - 验证前端配置完整性（embedded 需要 frontend_path，external 需要 frontend_url）
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

## 5. 前端资源路由系统
- [ ] 5.1 创建前端资源路由处理器
  - 创建 `router/frontend.go` 文件
  - 实现 `FrontendRouter()` 函数，根据品牌配置路由到对应前端资源
  - 支持 SPA 路由（非文件路径返回 index.html）
  - 处理 `/panel` 路径路由到管理后台
  - 处理静态资源文件（js, css, images 等）
  - 实现文件存在性检查和 404 处理
  - _Requirements: Brand Management - Requirement 5.1_

- [ ] 5.2 注册前端路由到主路由
  - 修改 `router/main.go` 的 `SetRouter` 函数
  - 在品牌识别中间件之后注册前端路由
  - 确保 API 路由优先级高于前端路由
  - 配置正确的路由顺序：API -> /panel -> 品牌前端
  - _Requirements: Brand Management - Requirement 5.1_

- [ ]* 5.3 编写前端路由单元测试
  - 创建 `router/frontend_test.go` 文件
  - 测试品牌前端资源路由
  - 测试管理后台路由
  - 测试 SPA 路由回退
  - 测试 404 处理
  - _Requirements: Brand Management - Requirement 5.1_

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
    - Logo 路径（logo，文本输入，必填，带提示）
    - Favicon 路径（favicon，文本输入，必填，带提示）
    - 描述（description，文本域）
    - 关键词（keywords，文本输入）
    - 作者（author，文本输入）
    - 前端类型（frontend_type，下拉选择：embedded/external）
    - 前端资源路径（frontend_path，文本输入，embedded 时显示）
    - 前端 URL（frontend_url，文本输入，external 时显示）
    - 是否默认品牌（is_default，复选框）
    - 是否启用（enabled，复选框，默认选中）
  - 实现表单验证（前端验证 + 后端验证）
  - 实现前端类型切换逻辑（显示/隐藏对应字段）
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

## 7. 多前端项目架构
- [ ] 7.1 创建 Kapon 品牌前端项目
  - 创建 `web/kapon-portal/` 目录
  - 初始化前端项目（npm init, 选择框架如 React/Vite）
  - 配置构建输出到 `../public/brands/kapon/`
  - 创建基础页面结构和路由
  - 实现 API 调用逻辑（调用 /api/status 获取品牌配置）
  - 实现品牌特定的 UI 设计
  - _Requirements: Brand Management - Requirement 5.4_

- [ ] 7.2 创建 Grouplay 品牌前端项目
  - 创建 `web/grouplay-portal/` 目录
  - 初始化前端项目（可以使用不同框架如 Vue）
  - 配置构建输出到 `../public/brands/grouplay/`
  - 创建基础页面结构和路由
  - 实现 API 调用逻辑
  - 实现品牌特定的 UI 设计
  - _Requirements: Brand Management - Requirement 5.4_

- [ ] 6.3 创建或改造管理后台项目
  - 确保管理后台位于 `web/admin/` 目录
  - 配置构建输出到 `../public/admin/`
  - 修改 Logo 组件支持动态品牌 logo
  - 从 /api/status 获取品牌信息并显示对应 logo
  - 确保其他 UI 元素保持一致（不受品牌影响）
  - _Requirements: Brand Management - Requirement 5.2_

- [ ] 6.4 配置多前端构建脚本
  - 在根目录 `package.json` 中添加构建脚本
  - 实现 `build:kapon` 脚本（构建 Kapon 前端）
  - 实现 `build:grouplay` 脚本（构建 Grouplay 前端）
  - 实现 `build:admin` 脚本（构建管理后台）
  - 实现 `build:all` 脚本（构建所有前端）
  - 实现开发模式脚本（dev:kapon, dev:grouplay, dev:admin）
  - _Requirements: Brand Management - Requirement 5.6_

## 8. 管理后台品牌标识支持
- [ ] 8.1 修改管理后台 Logo 组件
  - 修改 `web/admin/src/components/Logo.jsx`（或对应路径）
  - 从 Redux store 或 Context 读取品牌信息
  - 动态显示品牌 logo 和系统名称
  - 保持其他 UI 元素不变
  - _Requirements: Brand Management - Requirement 5.2_

- [ ] 8.2 管理后台集成品牌配置
  - 在管理后台启动时调用 /api/status
  - 存储品牌信息到状态管理（Redux/Context）
  - 确保品牌信息在整个应用中可访问
  - _Requirements: Brand Management - Requirement 5.2_

- [ ]* 8.3 编写管理后台品牌显示测试
  - 测试不同品牌域名访问管理后台
  - 验证显示正确的品牌 logo
  - 验证其他 UI 元素保持一致
  - _Requirements: Brand Management - Requirement 5.2_

## 9. CORS 配置（可选，用于外部前端部署）
- [ ]* 9.1 实现 CORS 中间件配置
  - 修改或创建 CORS 中间件
  - 从配置文件读取 allowed_origins
  - 支持动态配置允许的域名列表
  - 配置 allowed_methods 和 allowed_headers
  - 支持 credentials
  - _Requirements: Brand Management - Requirement 5.5_

- [ ]* 9.2 测试 CORS 配置
  - 测试跨域 API 调用
  - 验证 preflight 请求处理
  - 验证不同域名的访问权限
  - _Requirements: Brand Management - Requirement 5.5_

## 10. 集成测试与验证
- [ ] 10.1 多前端路由测试
  - 通过 models.kapon.cloud 访问，验证返回 Kapon 前端
  - 通过 model.grouplay.cn 访问，验证返回 Grouplay 前端
  - 通过 models.kapon.cloud/panel 访问，验证返回管理后台
  - 验证管理后台显示 Kapon 品牌 logo
  - 通过 model.grouplay.cn/panel 访问，验证显示 Grouplay 品牌 logo
  - _Requirements: Brand Management - Requirement 5.1, 5.2_

- [ ] 10.2 API 调用测试
  - 从 Kapon 前端调用 API，验证自动识别品牌
  - 从 Grouplay 前端调用 API，验证自动识别品牌
  - 验证不同品牌前端获取相同的数据
  - 验证前端不需要在请求中添加品牌参数
  - _Requirements: Brand Management - Requirement 5.3_

- [ ] 10.3 SPA 路由测试
  - 访问 Kapon 前端的子路径（如 /dashboard）
  - 验证返回 index.html 并由前端路由处理
  - 刷新页面验证路由正常工作
  - _Requirements: Brand Management - Requirement 5.1_

- [ ] 10.4 品牌管理界面测试
  - 测试通过管理后台添加品牌
  - 测试编辑品牌配置
  - 测试删除品牌
  - 测试启用/禁用品牌
  - 测试设置默认品牌
  - 验证配置更改后立即生效（缓存刷新）
  - _Requirements: Brand Management - Requirement 6, 7_

- [ ] 10.5 向后兼容性测试
  - 数据库中无品牌配置时，验证系统正常运行
  - 验证使用全局配置（SystemName, Logo）
  - 验证前端正常显示
  - _Requirements: Brand Management - Requirement 4.1_

- [ ] 10.6 边界情况测试
  - 测试未知域名访问（使用默认品牌）
  - 测试品牌前端资源不存在（返回 404）
  - 测试域名冲突验证（创建品牌时）
  - 测试删除默认品牌（应该被阻止）
  - _Requirements: Brand Management - Requirement 4.2, 8_

## 11. 文档与部署
- [ ] 11.1 编写数据库迁移文档
  - 说明如何执行数据库迁移
  - 提供迁移脚本示例
  - 说明回滚方案
  - _Requirements: Brand Management - Requirement 1.1_

- [ ] 11.2 编写品牌管理使用文档
  - 说明如何通过管理后台添加品牌
  - 说明各字段的含义和填写规范
  - 说明品牌资源文件的放置位置
  - 说明前端项目的构建和部署
  - 提供配置示例和最佳实践
  - _Requirements: Brand Management - Requirement 7_

- [ ] 11.3 编写多前端部署文档
  - 创建多品牌多前端部署指南
  - 说明前端项目构建步骤
  - 说明配置步骤和注意事项
  - 提供 Nginx 配置示例（统一部署模式）
  - 提供 Nginx 配置示例（独立部署模式 + CORS）
  - 提供 DNS 配置说明
  - 说明如何添加新品牌
  - 提供回滚方案
  - _Requirements: Brand Management - Requirement 5.4, 5.5_

- [ ] 11.4 更新 README
  - 在 README 中添加多品牌多前端支持说明
  - 添加架构图和管理界面截图
  - 说明如何通过管理后台配置品牌
  - 说明前端项目结构
  - 说明构建和部署流程
  - 说明向后兼容性
  - _Requirements: Brand Management - Requirement 5.4, 7_

- [ ] 11.5 创建前端开发指南
  - 说明如何创建新的品牌前端项目
  - 说明前端项目的目录结构规范
  - 说明如何调用后端 API
  - 说明如何获取和使用品牌配置
  - 提供前端项目模板或脚手架
  - _Requirements: Brand Management - Requirement 5.4_

- [ ]* 11.6 创建验证脚本（可选）
  - 创建品牌配置验证脚本
  - 创建品牌资源检查脚本
  - 创建前端构建验证脚本
  - 创建数据库迁移验证脚本
  - _Requirements: Brand Management - Requirement 1.1, 8_
