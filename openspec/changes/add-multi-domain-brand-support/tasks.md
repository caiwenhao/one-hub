# 实现任务清单

## 1. 后端品牌配置系统
- [ ] 1.1 创建品牌数据模型
  - 创建 `model/brand.go` 文件
  - 定义 `Brand` 结构体（name, domains, system_name, logo, favicon, description, keywords, author）
  - 定义 `BrandManager` 结构体（brands map, domainMap, defaultBrand）
  - 实现 `GetBrandByDomain(domain string) *Brand` 方法
  - 实现 `LoadFromConfig(brands []Brand)` 方法
  - 实现域名到品牌的映射逻辑（支持端口）
  - _Requirements: Brand Management - Requirement 1.1, 1.2_

- [ ] 1.2 实现品牌配置加载
  - 在 `common/config/brand.go` 中实现配置加载逻辑
  - 从 Viper 读取 `brands` 配置段
  - 初始化全局 `BrandManager` 实例
  - 在 `main.go` 的 `InitConf()` 后调用品牌初始化
  - 处理配置错误（格式错误、域名冲突等）
  - 添加日志记录（启动时输出品牌配置信息）
  - _Requirements: Brand Management - Requirement 1.3_

- [ ]* 1.3 编写品牌管理器单元测试
  - 创建 `model/brand_test.go` 文件
  - 测试 `GetBrandByDomain` 方法（精确匹配、端口匹配、未匹配）
  - 测试 `LoadFromConfig` 方法（正常加载、空配置、错误配置）
  - 测试默认品牌逻辑
  - 测试域名冲突处理
  - _Requirements: Brand Management - Requirement 1.1, 1.2_

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

## 3. API 接口修改
- [ ] 3.1 修改 GetStatus 接口返回品牌信息
  - 修改 `controller/misc.go` 的 `GetStatus` 函数
  - 从 Context 读取品牌：`c.Get("brand")`
  - 构建品牌信息 JSON（brand_name, brand_logo, brand_favicon, system_name, description, keywords, author）
  - 合并到响应的 data 字段中
  - 保持向后兼容：未匹配品牌时使用全局配置（config.SystemName, config.Logo）
  - _Requirements: Brand Management - Requirement 2.2_

- [ ]* 3.2 编写 API 集成测试
  - 创建测试用例模拟不同 Host 头请求
  - 验证返回的品牌信息正确
  - 验证向后兼容性（无品牌配置时）
  - _Requirements: Brand Management - Requirement 2.2_

## 4. 配置文件扩展
- [ ] 4.1 更新配置文件示例
  - 修改 `config.example.yaml`
  - 添加 `brands` 配置段示例
  - 包含两个品牌示例（kapon 和 grouplay）
  - 添加详细的配置说明注释
  - _Requirements: Brand Management - Requirement 1.3_

- [ ] 4.2 更新配置文档
  - 在项目文档中添加多品牌配置说明
  - 说明配置格式和字段含义
  - 提供配置示例和最佳实践
  - _Requirements: Brand Management - Requirement 1.3_

## 5. 前端品牌资源管理
- [ ] 5.1 重组静态资源目录结构
  - 在 `web/public/` 下创建 `brands/` 目录
  - 创建 `brands/kapon/` 子目录
  - 创建 `brands/grouplay/` 子目录
  - 准备各品牌的 logo.png 和 favicon.ico 文件
  - 保留根目录的 logo.png（向后兼容）
  - _Requirements: Brand Management - Requirement 3.1_

- [ ] 5.2 修改 Logo 组件支持动态品牌
  - 修改 `web/src/ui-component/Logo.jsx`
  - 从 Redux store 读取 `siteInfo.brand_logo`
  - 优先使用 brand_logo，否则使用 logo，最后使用默认 /logo.png
  - 使用 system_name 作为 alt 文本
  - 处理加载状态（isLoading 时不显示）
  - _Requirements: Brand Management - Requirement 3.2_

- [ ] 5.3 扩展 Redux Store 支持品牌字段
  - 修改 `web/src/store/siteInfoReducer.js`
  - 在 initialState 中添加品牌相关字段（brand_name, brand_logo, brand_favicon, description, keywords, author）
  - 确保 SET_SITE_INFO action 正确更新这些字段
  - _Requirements: Brand Management - Requirement 3.2_

- [ ]* 5.4 编写 Logo 组件测试
  - 创建 `web/src/ui-component/Logo.test.jsx`
  - 测试品牌 logo 渲染
  - 测试默认 logo 回退
  - 测试加载状态
  - _Requirements: Brand Management - Requirement 3.2_

## 6. HTML Meta 信息动态化（可选）
- [ ]* 6.1 修改 HTML 模板支持动态 meta
  - 修改 `web/index.html`
  - 将 title、description、keywords 改为占位符或通过 JS 动态设置
  - 在前端启动时根据 siteInfo 更新 document.title 和 meta 标签
  - _Requirements: Brand Management - Requirement 3.3_

- [ ]* 6.2 实现动态 meta 更新逻辑
  - 在 `web/src/contexts/StatusContext.jsx` 中添加 meta 更新逻辑
  - 在 loadStatus 成功后更新 document.title
  - 更新 meta description 和 keywords
  - _Requirements: Brand Management - Requirement 3.3_

## 7. 集成测试与验证
- [ ] 7.1 本地开发环境测试
  - 配置两个品牌（kapon 和 grouplay）
  - 使用不同 Host 头测试（curl -H "Host: ..."）
  - 验证 /api/status 返回正确的品牌信息
  - 验证前端 Logo 显示正确
  - 验证浏览器控制台无错误
  - _Requirements: Brand Management - All Requirements_

- [ ] 7.2 向后兼容性测试
  - 删除 brands 配置，验证系统正常运行
  - 验证使用全局配置（SystemName, Logo）
  - 验证前端正常显示
  - _Requirements: Brand Management - Requirement 4.1_

- [ ] 7.3 边界情况测试
  - 测试未知域名访问（使用默认品牌）
  - 测试品牌资源文件不存在（前端显示默认图片）
  - 测试配置格式错误（启动时记录错误，使用单品牌模式）
  - 测试域名冲突（使用第一个匹配的品牌）
  - _Requirements: Brand Management - Requirement 4.2_

## 8. 文档与部署
- [ ] 8.1 编写部署文档
  - 创建多品牌部署指南
  - 说明配置步骤和注意事项
  - 提供 Nginx 配置示例
  - 提供 DNS 配置说明
  - 提供回滚方案
  - _Requirements: Brand Management - Requirement 1.3_

- [ ] 8.2 更新 README
  - 在 README 中添加多品牌支持说明
  - 添加配置示例链接
  - 说明向后兼容性
  - _Requirements: Brand Management - Requirement 1.3_

- [ ]* 8.3 创建迁移脚本（可选）
  - 创建配置验证脚本
  - 创建品牌资源检查脚本
  - 创建配置迁移辅助工具
  - _Requirements: Brand Management - Requirement 1.3_
