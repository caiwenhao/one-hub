# 实施任务清单

## 1. 数据库设计和迁移
- [ ] 1.1 创建 Supplier 表结构
  - 定义字段：id, name, code, description, status, priority, created_time, updated_time
  - 创建唯一索引：name, code
  - _Requirements: 供应商管理_

- [ ] 1.2 创建 SupplierModel 表结构
  - 定义字段：id, supplier_id, model, quality_tier, purchase_discount, channel_type, status, remark, created_time, updated_time
  - 创建复合索引：(supplier_id, model)
  - 创建单独索引：model
  - _Requirements: 供应商模型配置_

- [ ] 1.3 创建 CustomerPricing 表结构
  - 定义字段：id, user_group_id, model, quality_tier, pricing_type, input_value, output_value, priority, status, created_time, updated_time
  - 创建复合索引：(user_group_id, model)
  - _Requirements: 客户定价策略_

- [ ] 1.4 扩展 Channel 表
  - 增加字段：supplier_id (nullable), quality_tier (nullable)
  - 创建索引：supplier_id
  - 确保向后兼容（字段可为 NULL）
  - _Requirements: 渠道关联供应商, 数据迁移兼容_

- [ ] 1.5 编写数据库迁移脚本
  - 在 model/migrate.go 中添加迁移函数
  - 实现 Migrate 和 Rollback 逻辑
  - 添加数据一致性检查
  - _Requirements: 数据迁移兼容_

## 2. 数据模型实现
- [ ] 2.1 实现 Supplier 模型
  - 创建 model/supplier.go 文件
  - 定义 Supplier 结构体和 TableName
  - 实现 CRUD 方法：Insert, Update, Delete
  - 实现查询方法：GetSupplierById, GetAllSuppliers
  - _Requirements: 供应商管理_

- [ ] 2.2 实现 SupplierModel 模型
  - 在 model/supplier.go 中定义 SupplierModel 结构体
  - 实现 CRUD 方法：Insert, Update, Delete
  - 实现查询方法：GetSupplierModelById, GetSupplierModelsByModel, GetSupplierModelsBySupplierId
  - 实现关联查询（Preload Supplier）
  - _Requirements: 供应商模型配置_

- [ ] 2.3 实现 CustomerPricing 模型
  - 在 model/supplier.go 中定义 CustomerPricing 结构体
  - 实现 CRUD 方法：Insert, Update, Delete
  - 实现优先级匹配查询：GetCustomerPricing
  - _Requirements: 客户定价策略_

- [ ] 2.4 扩展 Channel 模型
  - 在 model/channel.go 中增加 SupplierId 和 QualityTier 字段
  - 更新相关查询方法支持供应商筛选
  - 实现 GetChannelsBySupplierId 方法
  - 实现 GetSupplierByChannelId 方法
  - _Requirements: 渠道关联供应商_

## 3. 价格计算逻辑
- [ ] 3.1 实现采购成本计算
  - 在 model/pricing.go 中添加 CalculatePurchaseCost 函数
  - 查询品牌官方价格（Price 表）
  - 查询供应商采购折扣（SupplierModel 表）
  - 计算：采购成本 = 官方价格 × 采购折扣
  - 处理无供应商配置的情况（使用官方价格）
  - _Requirements: 采购成本计算_

- [ ] 3.2 实现客户售价计算
  - 在 model/pricing.go 中添加 CalculateCustomerPrice 函数
  - 实现定价策略优先级匹配逻辑
  - 支持 CustomerPricing 精确匹配、模型匹配、质量档位匹配
  - 回退到 UserGroup.ratio 作为默认
  - 计算：客户售价 = 官方价格 × 销售折扣
  - _Requirements: 客户定价策略_

- [ ] 3.3 实现利润计算
  - 在 model/pricing.go 中添加 CalculateProfit 函数
  - 计算：利润 = 客户售价 - 采购成本
  - 计算：利润率 = 利润 / 客户售价 × 100%
  - 返回详细的计费信息结构体
  - _Requirements: 利润计算_

- [ ] 3.4 集成到现有计费流程
  - 修改现有的 GetPrice 方法，支持供应商折扣
  - 在计费时记录采购成本和利润
  - 更新 Log 表或创建独立的利润记录表
  - 确保向后兼容（未关联供应商时使用原逻辑）
  - _Requirements: 采购成本计算, 利润计算, 数据迁移兼容_

## 4. 路由逻辑增强
- [ ] 4.1 实现质量档位筛选
  - 在路由模块中增加质量档位参数解析
  - 支持从 HTTP Header（X-Quality-Tier）读取
  - 支持从请求参数（quality_tier）读取
  - 支持从用户组默认配置读取
  - 根据质量档位筛选 Channel
  - _Requirements: 质量档位路由_

- [ ] 4.2 实现供应商优先级路由
  - 查询模型的所有 SupplierModel 配置
  - 按 Supplier.priority 排序
  - 筛选启用状态的供应商
  - 获取对应的 Channel 列表
  - 应用现有的负载均衡策略
  - _Requirements: 质量档位路由_

- [ ] 4.3 实现质量档位降级
  - 当请求的质量档位无可用 Channel 时
  - 尝试降级到 standard 档位
  - 记录降级日志
  - 或返回错误提示质量档位不可用
  - _Requirements: 质量档位路由_

## 5. 管理接口和 API
- [ ] 5.1 实现供应商管理 API
  - POST /api/suppliers - 创建供应商
  - GET /api/suppliers - 查询供应商列表
  - GET /api/suppliers/:id - 查询供应商详情
  - PUT /api/suppliers/:id - 更新供应商
  - DELETE /api/suppliers/:id - 删除供应商
  - PUT /api/suppliers/:id/status - 启用/禁用供应商
  - _Requirements: 供应商管理_

- [ ] 5.2 实现供应商模型配置 API
  - POST /api/supplier-models - 创建供应商模型配置
  - GET /api/supplier-models - 查询配置列表
  - GET /api/supplier-models/:id - 查询配置详情
  - PUT /api/supplier-models/:id - 更新配置
  - DELETE /api/supplier-models/:id - 删除配置
  - GET /api/models/:model/suppliers - 查询模型的所有供应商
  - _Requirements: 供应商模型配置_

- [ ] 5.3 实现客户定价策略 API
  - POST /api/customer-pricings - 创建定价策略
  - GET /api/customer-pricings - 查询策略列表
  - GET /api/customer-pricings/:id - 查询策略详情
  - PUT /api/customer-pricings/:id - 更新策略
  - DELETE /api/customer-pricings/:id - 删除策略
  - _Requirements: 客户定价策略_

- [ ] 5.4 扩展 Channel API
  - 在创建/更新 Channel 时支持 supplier_id 和 quality_tier
  - GET /api/suppliers/:id/channels - 查询供应商的所有渠道
  - 在 Channel 列表中显示供应商信息
  - _Requirements: 渠道关联供应商_

- [ ] 5.5 实现利润统计 API
  - GET /api/statistics/profit - 利润统计报表
  - 支持按时间段、供应商、模型、用户组筛选
  - 返回利润汇总、利润率、成本分析等数据
  - _Requirements: 利润计算_

## 6. 前端管理界面
- [ ] 6.1 供应商管理页面
  - 供应商列表页（表格展示，支持搜索、排序）
  - 创建供应商表单（名称、代码、描述、优先级）
  - 编辑供应商表单
  - 删除供应商确认对话框
  - 启用/禁用供应商开关
  - _Requirements: 供应商管理_

- [ ] 6.2 供应商模型配置页面
  - 配置列表页（按供应商分组展示）
  - 创建配置表单（选择供应商、模型、质量档位、采购折扣、渠道类型）
  - 编辑配置表单
  - 删除配置确认对话框
  - 批量导入配置功能
  - _Requirements: 供应商模型配置_

- [ ] 6.3 客户定价策略页面
  - 策略列表页（按用户组分组展示）
  - 创建策略表单（选择用户组、模型、质量档位、定价类型、折扣值）
  - 编辑策略表单
  - 删除策略确认对话框
  - 优先级调整功能
  - _Requirements: 客户定价策略_

- [ ] 6.4 Channel 关联供应商功能
  - 在 Channel 创建/编辑表单中增加供应商选择
  - 在 Channel 创建/编辑表单中增加质量档位选择
  - 在 Channel 列表中显示供应商和质量档位
  - 批量关联供应商功能
  - _Requirements: 渠道关联供应商_

- [ ] 6.5 利润统计报表页面
  - 利润概览仪表板（总利润、利润率、成本占比）
  - 按时间段的利润趋势图
  - 按供应商的利润分析表
  - 按模型的利润分析表
  - 按用户组的利润分析表
  - 导出报表功能
  - _Requirements: 利润计算_

## 7. 测试
- [ ] 7.1 单元测试
  - 测试 Supplier CRUD 操作
  - 测试 SupplierModel CRUD 操作
  - 测试 CustomerPricing CRUD 操作
  - 测试价格计算逻辑（采购成本、客户售价、利润）
  - 测试定价策略优先级匹配
  - 测试质量档位路由逻辑
  - _Requirements: 所有需求_

- [ ] 7.2 集成测试
  - 测试完整的计费流程（从请求到计费记录）
  - 测试供应商禁用后的影响
  - 测试质量档位降级逻辑
  - 测试数据迁移兼容性（新旧逻辑共存）
  - 测试并发场景下的数据一致性
  - _Requirements: 所有需求_

- [ ] 7.3 性能测试
  - 测试价格计算性能（目标 < 10ms）
  - 测试路由选择性能
  - 测试大量供应商配置下的查询性能
  - 测试缓存效果
  - _Requirements: 所有需求_

## 8. 文档和部署
- [ ] 8.1 编写 API 文档
  - 供应商管理 API 文档
  - 供应商模型配置 API 文档
  - 客户定价策略 API 文档
  - 更新 Channel API 文档
  - 利润统计 API 文档
  - _Requirements: 所有需求_

- [ ] 8.2 编写用户手册
  - 供应商管理操作指南
  - 定价策略配置指南
  - 质量档位使用说明
  - 利润分析报表使用说明
  - _Requirements: 所有需求_

- [ ] 8.3 数据迁移准备
  - 编写数据迁移脚本
  - 准备回滚方案
  - 编写迁移操作手册
  - 准备测试数据
  - _Requirements: 数据迁移兼容_

- [ ] 8.4 部署和监控
  - 配置数据库迁移任务
  - 配置监控告警（计费准确性、响应时间）
  - 准备灰度发布计划
  - 编写上线检查清单
  - _Requirements: 所有需求_
