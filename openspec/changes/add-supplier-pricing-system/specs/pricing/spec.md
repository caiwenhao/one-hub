# 定价系统规范

## ADDED Requirements

### Requirement: 供应商管理

系统 SHALL 提供供应商管理功能，支持创建、查询、更新和删除上游供应商信息。

#### Scenario: 创建供应商
- **WHEN** 管理员提供供应商名称、代码和描述
- **THEN** 系统 SHALL 创建供应商记录
- **AND** 系统 SHALL 分配唯一的供应商ID
- **AND** 系统 SHALL 设置供应商状态为启用

#### Scenario: 查询供应商列表
- **WHEN** 管理员请求查看供应商列表
- **THEN** 系统 SHALL 返回所有供应商
- **AND** 系统 SHALL 按优先级降序排序
- **AND** 系统 SHALL 显示供应商的状态和配置的模型数量

#### Scenario: 更新供应商优先级
- **WHEN** 管理员修改供应商优先级
- **THEN** 系统 SHALL 更新供应商优先级
- **AND** 系统 SHALL 影响后续的渠道路由选择

#### Scenario: 禁用供应商
- **WHEN** 管理员禁用某个供应商
- **THEN** 系统 SHALL 将供应商状态设置为禁用
- **AND** 系统 SHALL 停止使用该供应商的所有渠道

### Requirement: 供应商模型配置

系统 SHALL 支持为每个供应商配置其提供的模型、质量档位、采购折扣和渠道类型。

#### Scenario: 配置供应商模型
- **WHEN** 管理员为供应商添加模型配置
- **AND** 提供模型名称、质量档位、采购折扣和渠道类型
- **THEN** 系统 SHALL 创建供应商模型配置记录
- **AND** 系统 SHALL 验证模型名称在 Price 表中存在
- **AND** 系统 SHALL 验证采购折扣在 0 到 1 之间

#### Scenario: 同一模型多质量档位
- **WHEN** 供应商对同一模型提供多个质量档位
- **THEN** 系统 SHALL 允许创建多条配置记录
- **AND** 每条记录 SHALL 有不同的质量档位标识
- **AND** 每条记录 SHALL 有独立的采购折扣

#### Scenario: 查询供应商模型配置
- **WHEN** 管理员查询某个模型的供应商配置
- **THEN** 系统 SHALL 返回所有提供该模型的供应商
- **AND** 系统 SHALL 显示每个供应商的质量档位和采购折扣
- **AND** 系统 SHALL 按供应商优先级排序

#### Scenario: 不同模型使用不同渠道类型
- **WHEN** 供应商的不同模型使用不同的 API 标准
- **THEN** 系统 SHALL 允许为每个模型配置独立的渠道类型
- **AND** 系统 SHALL 在创建 Channel 时验证类型匹配

### Requirement: 渠道关联供应商

系统 SHALL 支持将 Channel 关联到供应商和质量档位。

#### Scenario: 创建关联供应商的渠道
- **WHEN** 管理员创建新渠道
- **AND** 指定供应商ID和质量档位
- **THEN** 系统 SHALL 创建渠道记录
- **AND** 系统 SHALL 记录供应商关联关系
- **AND** 系统 SHALL 记录质量档位标识

#### Scenario: 查询供应商的所有渠道
- **WHEN** 管理员查询某个供应商的渠道
- **THEN** 系统 SHALL 返回该供应商的所有渠道
- **AND** 系统 SHALL 显示每个渠道的质量档位
- **AND** 系统 SHALL 显示每个渠道的状态

#### Scenario: 兼容现有渠道
- **WHEN** 系统中存在未关联供应商的旧渠道
- **THEN** 系统 SHALL 继续支持这些渠道的正常使用
- **AND** 系统 SHALL 允许后续为这些渠道关联供应商

### Requirement: 采购成本计算

系统 SHALL 根据品牌官方价格和供应商采购折扣计算采购成本。

#### Scenario: 计算采购成本
- **WHEN** 系统需要计算某个模型的采购成本
- **AND** 该模型有供应商配置
- **THEN** 系统 SHALL 查询品牌官方价格（Price表）
- **AND** 系统 SHALL 查询供应商采购折扣（SupplierModel表）
- **AND** 系统 SHALL 计算采购成本 = 官方价格 × 采购折扣

#### Scenario: 无供应商配置时的成本
- **WHEN** 某个模型没有供应商配置
- **THEN** 系统 SHALL 使用品牌官方价格作为采购成本
- **AND** 系统 SHALL 记录警告日志

#### Scenario: 多供应商选择最优成本
- **WHEN** 多个供应商提供同一模型的同一质量档位
- **THEN** 系统 SHALL 选择采购折扣最低的供应商
- **OR** 系统 SHALL 根据供应商优先级选择

### Requirement: 客户定价策略

系统 SHALL 支持基于官方价格的客户销售折扣配置。

#### Scenario: 使用用户组默认倍率
- **WHEN** 客户请求使用某个模型
- **AND** 没有针对该模型的特殊定价策略
- **THEN** 系统 SHALL 使用用户组的默认倍率（UserGroup.ratio）
- **AND** 系统 SHALL 计算客户售价 = 官方价格 × 倍率

#### Scenario: 配置模型级定价策略
- **WHEN** 管理员为用户组配置特定模型的定价策略
- **THEN** 系统 SHALL 创建 CustomerPricing 记录
- **AND** 系统 SHALL 记录销售折扣值
- **AND** 系统 SHALL 在计费时优先使用该策略

#### Scenario: 配置质量档位定价策略
- **WHEN** 管理员为用户组配置特定质量档位的定价策略
- **THEN** 系统 SHALL 创建 CustomerPricing 记录
- **AND** 系统 SHALL 记录质量档位和销售折扣
- **AND** 系统 SHALL 在计费时根据质量档位匹配策略

#### Scenario: 定价策略优先级
- **WHEN** 存在多个匹配的定价策略
- **THEN** 系统 SHALL 按以下优先级选择：
  1. 精确匹配（模型 + 质量档位）
  2. 模型匹配（忽略质量档位）
  3. 质量档位匹配（忽略模型）
  4. 用户组默认倍率

### Requirement: 利润计算

系统 SHALL 计算每次请求的利润（客户售价 - 采购成本）。

#### Scenario: 计算单次请求利润
- **WHEN** 完成一次模型调用计费
- **THEN** 系统 SHALL 计算客户售价
- **AND** 系统 SHALL 计算采购成本
- **AND** 系统 SHALL 计算利润 = 客户售价 - 采购成本
- **AND** 系统 SHALL 记录利润数据用于统计分析

#### Scenario: 利润统计报表
- **WHEN** 管理员查询利润统计
- **THEN** 系统 SHALL 提供按时间段的利润汇总
- **AND** 系统 SHALL 提供按供应商的利润分析
- **AND** 系统 SHALL 提供按模型的利润分析
- **AND** 系统 SHALL 提供按用户组的利润分析

### Requirement: 质量档位路由

系统 SHALL 支持根据质量档位要求路由到合适的渠道。

#### Scenario: 客户指定质量档位
- **WHEN** 客户在请求中指定质量档位（如通过 header）
- **THEN** 系统 SHALL 筛选匹配该质量档位的渠道
- **AND** 系统 SHALL 从筛选结果中选择可用渠道
- **AND** 系统 SHALL 使用对应质量档位的定价策略

#### Scenario: 使用默认质量档位
- **WHEN** 客户未指定质量档位
- **THEN** 系统 SHALL 使用默认质量档位（standard）
- **AND** 系统 SHALL 筛选标准质量档位的渠道

#### Scenario: 质量档位不可用时降级
- **WHEN** 请求的质量档位没有可用渠道
- **THEN** 系统 SHALL 尝试降级到标准质量档位
- **OR** 系统 SHALL 返回错误提示质量档位不可用

### Requirement: 数据迁移兼容

系统 SHALL 保持与现有数据的向后兼容性。

#### Scenario: 迁移现有 Channel 数据
- **WHEN** 系统升级到新版本
- **THEN** 系统 SHALL 为 Channel 表增加 supplier_id 和 quality_tier 字段
- **AND** 新字段 SHALL 允许为 NULL
- **AND** 现有 Channel 记录 SHALL 保持不变

#### Scenario: 现有定价逻辑兼容
- **WHEN** Channel 未关联供应商
- **THEN** 系统 SHALL 使用原有的定价逻辑
- **AND** 系统 SHALL 使用 Price 表的官方价格
- **AND** 系统 SHALL 使用 UserGroup.ratio 计算客户售价

#### Scenario: 渐进式迁移
- **WHEN** 管理员逐步为 Channel 关联供应商
- **THEN** 系统 SHALL 对已关联的 Channel 使用新定价逻辑
- **AND** 系统 SHALL 对未关联的 Channel 使用旧定价逻辑
- **AND** 两种逻辑 SHALL 可以共存
