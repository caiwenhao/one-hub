# 供应商定价系统设计文档

## Context

当前系统作为大模型 API 聚合平台，需要对接多个上游供应商，每个供应商可能提供不同品牌的模型，且对同一模型可能提供不同质量档位。现有的定价系统只有全局的品牌官方价格和用户组倍率，无法满足精细化的成本管理和定价需求。

### 业务背景
- **2B 业务模式**：面向企业客户，需要差异化定价
- **多供应商采购**：同一模型可能从多个供应商采购，成本不同
- **质量分级**：供应商提供不同质量档位（高质量、标准、经济型）
- **利润核算**：需要清晰计算采购成本和销售利润

## Goals / Non-Goals

### Goals
1. 建立供应商管理体系，记录供应商信息和采购折扣
2. 支持质量档位分级，满足不同客户需求
3. 实现基于官方价格的采购成本和客户售价计算
4. 保持向后兼容，不影响现有系统运行
5. 支持利润分析和统计

### Non-Goals
1. 不实现动态定价（根据市场实时调整价格）
2. 不实现合同定价（复杂的长期协议价格）
3. 不实现阶梯定价（根据用量自动调整）
4. 不修改现有的 Price 表结构（保持品牌官方价格不变）

## Decisions

### 1. 数据模型设计

#### 核心表结构

**Supplier（供应商表）**
```go
type Supplier struct {
    Id          int    `json:"id"`
    Name        string `json:"name" gorm:"type:varchar(100);uniqueIndex"`
    Code        string `json:"code" gorm:"type:varchar(50);uniqueIndex"`
    Description string `json:"description" gorm:"type:text"`
    Status      int    `json:"status" gorm:"default:1"` // 1启用 2禁用
    Priority    int    `json:"priority" gorm:"default:0"` // 优先级
    CreatedTime int64  `json:"created_time" gorm:"bigint"`
    UpdatedTime int64  `json:"updated_time" gorm:"bigint"`
}
```

**SupplierModel（供应商模型配置表）**
```go
type SupplierModel struct {
    Id               int     `json:"id"`
    SupplierId       int     `json:"supplier_id" gorm:"index:idx_supplier_model"`
    Model            string  `json:"model" gorm:"type:varchar(100);index:idx_supplier_model"`
    QualityTier      string  `json:"quality_tier" gorm:"type:varchar(50);default:'standard'"`
    PurchaseDiscount float64 `json:"purchase_discount" gorm:"type:decimal(10,4);default:1"`
    ChannelType      int     `json:"channel_type" gorm:"default:0"`
    Status           int     `json:"status" gorm:"default:1"`
    Remark           string  `json:"remark" gorm:"type:text"`
    CreatedTime      int64   `json:"created_time" gorm:"bigint"`
    UpdatedTime      int64   `json:"updated_time" gorm:"bigint"`
}
```

**CustomerPricing（客户定价策略表，可选）**
```go
type CustomerPricing struct {
    Id            int     `json:"id"`
    UserGroupId   int     `json:"user_group_id" gorm:"index:idx_customer_pricing"`
    Model         string  `json:"model" gorm:"type:varchar(100);index:idx_customer_pricing"`
    QualityTier   string  `json:"quality_tier" gorm:"type:varchar(50)"`
    PricingType   string  `json:"pricing_type" gorm:"type:varchar(20);default:'ratio'"`
    InputValue    float64 `json:"input_value" gorm:"type:decimal(10,4);default:1"`
    OutputValue   float64 `json:"output_value" gorm:"type:decimal(10,4);default:1"`
    Priority      int     `json:"priority" gorm:"default:0"`
    Status        int     `json:"status" gorm:"default:1"`
    CreatedTime   int64   `json:"created_time" gorm:"bigint"`
    UpdatedTime   int64   `json:"updated_time" gorm:"bigint"`
}
```

**Channel 表扩展**
```go
// 在现有 Channel 结构上新增字段
type Channel struct {
    // ... 现有字段保持不变 ...
    
    SupplierId  *int    `json:"supplier_id" gorm:"index"`
    QualityTier *string `json:"quality_tier" gorm:"type:varchar(50)"`
}
```

#### 索引设计
- `suppliers`: `name` 唯一索引, `code` 唯一索引
- `supplier_models`: 复合索引 `(supplier_id, model)`, 单独索引 `model`
- `customer_pricings`: 复合索引 `(user_group_id, model)`
- `channels`: 新增索引 `supplier_id`

### 2. 质量档位定义

采用三级质量档位：
- **premium**：高质量，响应快，稳定性高，价格高
- **standard**：标准质量，平衡性能和成本（默认）
- **economy**：经济型，成本优先，可能响应慢或限流

**决策理由**：
- 三级分类简单清晰，易于理解和管理
- 可扩展：未来可以增加更多档位
- 与行业惯例一致

### 3. 价格计算逻辑

#### 采购成本计算
```
采购成本 = 品牌官方价格 × 供应商采购折扣
```

示例：
- GPT-4 官方价：input=$0.03/1k, output=$0.06/1k
- 供应商A采购折扣：0.85（85%）
- 采购成本：input=$0.0255/1k, output=$0.051/1k

#### 客户售价计算（基于官方价）
```
客户售价 = 品牌官方价格 × 客户销售折扣
```

**优先级匹配**：
1. CustomerPricing 精确匹配（模型 + 质量档位）
2. CustomerPricing 模型匹配（忽略质量档位）
3. CustomerPricing 质量档位匹配（忽略模型）
4. UserGroup.ratio（全局默认）

示例：
- GPT-4 官方价：$0.03/1k
- VIP 用户组倍率：1.2
- 客户售价：$0.036/1k

#### 利润计算
```
利润 = 客户售价 - 采购成本
利润率 = (客户售价 - 采购成本) / 客户售价 × 100%
```

**决策理由**：
- 基于官方价定价，客户容易理解和接受
- 不同供应商成本不同，但客户看到的价格一致
- 利润率可能因供应商不同而变化，需要通过供应商选择优化

### 4. 路由策略

#### 质量档位路由
```
1. 客户请求携带质量档位参数（如 X-Quality-Tier: premium）
2. 系统筛选匹配该质量档位的 Channel
3. 从筛选结果中根据优先级、权重、响应时间选择
4. 如果没有可用 Channel，尝试降级到 standard
```

#### 供应商优先级
```
1. 查询该模型的所有 SupplierModel 配置
2. 按 Supplier.priority 降序排序
3. 筛选状态为启用的供应商
4. 获取对应的 Channel 列表
5. 应用现有的负载均衡策略
```

### 5. 数据迁移策略

#### 阶段1：表结构迁移
```sql
-- 新增表
CREATE TABLE suppliers (...);
CREATE TABLE supplier_models (...);
CREATE TABLE customer_pricings (...);

-- 扩展 channels 表
ALTER TABLE channels ADD COLUMN supplier_id INT NULL;
ALTER TABLE channels ADD COLUMN quality_tier VARCHAR(50) NULL;
ALTER TABLE channels ADD INDEX idx_supplier_id (supplier_id);
```

#### 阶段2：数据兼容
- 现有 Channel 的 `supplier_id` 和 `quality_tier` 为 NULL
- 系统检测到 NULL 时，使用原有逻辑（不查询供应商折扣）
- 管理员可以逐步为 Channel 关联供应商

#### 阶段3：渐进式迁移
- 提供管理界面，允许管理员：
  1. 创建供应商
  2. 配置供应商模型
  3. 为现有 Channel 关联供应商
- 新旧逻辑共存，不影响业务

## Alternatives Considered

### 方案A：基于采购成本定价（未采用）
```
客户售价 = 采购成本 × 加价倍率
```

**优点**：
- 利润率统一，易于控制
- 成本变化时自动调整售价

**缺点**：
- 客户看到的价格不稳定（供应商变化导致价格变化）
- 需要实时计算，复杂度高
- 客户难以理解定价逻辑

**决策**：不采用，保持基于官方价定价

### 方案B：SupplierChannel 关联表（未采用）
```
创建独立的 supplier_channels 表关联 Supplier 和 Channel
```

**优点**：
- 关系更清晰
- 支持多对多关系

**缺点**：
- 增加表数量和查询复杂度
- 实际业务中一个 Channel 只属于一个供应商

**决策**：直接在 Channel 表增加 supplier_id 字段

### 方案C：质量档位作为独立表（未采用）
```
创建 quality_tiers 表存储档位定义
```

**优点**：
- 档位可配置
- 支持更多元数据

**缺点**：
- 过度设计，三个固定档位足够
- 增加查询复杂度

**决策**：使用字符串枚举（premium/standard/economy）

## Risks / Trade-offs

### 风险1：数据一致性
**风险**：Channel 关联的供应商被删除或禁用

**缓解措施**：
- 删除供应商前检查是否有关联的 Channel
- 禁用供应商时自动禁用关联的 Channel
- 提供数据一致性检查工具

### 风险2：性能影响
**风险**：价格计算需要多表关联查询，可能影响性能

**缓解措施**：
- 缓存供应商配置和定价策略
- 使用 Redis 缓存热点数据
- 定期预加载供应商和模型配置到内存

### 风险3：迁移复杂度
**风险**：现有系统数据量大，迁移可能影响业务

**缓解措施**：
- 采用渐进式迁移，新旧逻辑共存
- 提供回滚机制
- 充分测试后再上线

### Trade-off：灵活性 vs 复杂度
- **选择**：基于官方价定价，而非采购成本
- **得到**：定价逻辑简单，客户容易理解
- **失去**：利润率不统一，需要通过供应商选择优化

## Migration Plan

### 步骤1：数据库迁移（1天）
1. 创建新表：suppliers, supplier_models, customer_pricings
2. 扩展 channels 表：增加 supplier_id, quality_tier 字段
3. 创建索引
4. 执行数据一致性检查

### 步骤2：代码实现（1周）
1. 实现 Supplier 相关 CRUD（model/supplier.go）
2. 扩展 Channel 模型和查询逻辑
3. 实现价格计算逻辑（采购成本、客户售价、利润）
4. 实现质量档位路由逻辑
5. 单元测试和集成测试

### 步骤3：管理界面（3天）
1. 供应商管理页面（列表、创建、编辑、删除）
2. 供应商模型配置页面
3. 客户定价策略配置页面
4. Channel 关联供应商功能
5. 利润统计报表

### 步骤4：测试和上线（2天）
1. 功能测试
2. 性能测试
3. 数据迁移演练
4. 灰度发布
5. 监控和回滚准备

### 回滚方案
1. 保留原有定价逻辑代码
2. 通过配置开关控制新旧逻辑
3. 数据库迁移可回滚（删除新表和字段）
4. 监控关键指标（计费准确性、响应时间）

## Open Questions

1. **质量档位如何传递？**
   - 通过 HTTP Header（X-Quality-Tier）？
   - 通过请求参数（quality_tier）？
   - 通过用户组默认配置？
   
   **建议**：支持多种方式，优先级：请求参数 > Header > 用户组默认

2. **供应商优先级如何影响路由？**
   - 是否完全按优先级选择（高优先级优先）？
   - 还是结合权重、响应时间等因素？
   
   **建议**：优先级作为初筛条件，然后应用现有的负载均衡策略

3. **利润统计的时间粒度？**
   - 实时计算还是定时汇总？
   - 存储在哪里（logs 表还是独立统计表）？
   
   **建议**：实时计算并记录在 logs 表，定时汇总到统计表

4. **是否需要供应商成本预警？**
   - 当采购成本接近或超过售价时预警？
   - 当利润率低于阈值时预警？
   
   **建议**：第一版不实现，后续根据需求添加

## Success Metrics

1. **功能完整性**：所有需求场景通过测试
2. **性能指标**：价格计算延迟 < 10ms
3. **数据准确性**：计费金额与预期一致，误差 < 0.01%
4. **迁移成功率**：现有数据 100% 兼容
5. **用户满意度**：管理员能够顺利配置供应商和定价策略
