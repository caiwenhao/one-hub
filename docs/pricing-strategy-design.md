# 大模型API聚合平台 - 多层级定价策略产品设计方案

## 一、问题分析与概念澄清

### 1.1 核心业务场景

作为一个2B的大模型API聚合平台，您的业务模式是：

**上游供应商侧（采购端）：**
- 对接多个渠道供应商（如供应商A、B、C）
- 每个供应商提供同一模型的不同**折扣组**：
  - **号池**：成本最低，但稳定性较差，可能有封号风险
  - **官转**：官方转售，成本中等，稳定性较好
  - **逆向**：逆向工程接口，成本低，但合规风险高
  - **优质官转**：官方优质渠道，成本最高，稳定性最好，有SLA保障
- 同一模型+同一折扣组，不同供应商价格也可能不同

**下游客户侧（销售端）：**
- 不同客户对同一模型有不同的定价需求
- 客户对服务质量有不同要求（但不应知道真实的折扣组名称）
- 需要将上游的"折扣组"概念转化为客户友好的"服务等级"

**核心挑战：**
1. **成本管理复杂**：同一模型有多个供应商×多个折扣组的成本矩阵
2. **质量映射问题**：如何将上游的折扣组（号池、官转等）映射为客户侧的服务等级
3. **定价策略灵活性**：不同客户对同一模型+同一服务等级的价格可能不同
4. **路由决策复杂**：需要在成本、质量、利润之间做平衡

### 1.2 当前系统局限

根据代码审查，现有系统存在以下问题：
1. **单一价格维度**：每个模型只有一个全局价格（`model/price.go`），无法区分折扣组
2. **渠道选择简单**：基于权重和优先级的负载均衡（`model/balancer.go`），未考虑折扣组和成本
3. **用户分组粗糙**：仅支持用户组倍率（`model/user_group.go`），无法实现客户级定价
4. **缺乏质量抽象**：没有将上游折扣组抽象为客户侧服务等级的机制

## 二、产品设计方案

### 2.1 核心概念定义

#### 2.1.1 折扣组（Discount Group）- 上游真实分类

**定义：** 上游供应商提供的真实渠道类型，代表不同的成本、质量和风险等级

**典型折扣组：**

| 折扣组名称 | 成本水平 | 稳定性 | 风险等级 | 适用场景 |
|----------|---------|--------|---------|---------|
| 号池 (Pool) | 极低 | 差 | 高（封号风险） | 测试、低价值场景 |
| 逆向 (Reverse) | 低 | 中 | 高（合规风险） | 成本敏感场景 |
| 官转 (Official) | 中 | 好 | 低 | 标准生产环境 |
| 优质官转 (Premium) | 高 | 优秀 | 极低 | 关键业务、SLA保障 |

**管理维度：**
- 每个渠道可以提供一个或多个折扣组
- 同一折扣组在不同渠道的价格可能不同
- 折扣组是内部管理概念，**不对客户暴露**

#### 2.1.2 服务等级（Service Tier）- 客户侧抽象

**定义：** 面向客户的服务质量分级，隐藏了底层折扣组的技术细节

**典型服务等级：**

| 服务等级 | 客户视角描述 | 映射的折扣组 | 定价策略 |
|---------|------------|------------|---------|
| 经济版 (Economy) | 性价比之选，适合测试和开发 | 号池、逆向 | 低价 |
| 标准版 (Standard) | 稳定可靠，适合生产环境 | 官转 | 中价 |
| 专业版 (Professional) | 高可用性，适合关键业务 | 优质官转 | 高价 |
| 企业版 (Enterprise) | 最高SLA，专属支持 | 优质官转 + 专属通道 | 定制价 |

**映射规则：**
- 一个服务等级可以映射到一个或多个折扣组
- 映射关系可配置，支持动态调整
- 客户只看到服务等级，不知道底层折扣组

### 2.2 整体架构设计

```
┌─────────────────────────────────────────────────────────────────┐
│                        定价与路由引擎                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────┐  │
│  │ 上游成本管理      │  │ 服务等级映射      │  │ 客户定价管理  │  │
│  │                  │  │                  │  │              │  │
│  │ • 渠道×折扣组成本 │  │ • 折扣组→服务等级 │  │ • 客户价格表  │  │
│  │ • 供应商管理      │  │ • 映射规则配置    │  │ • 定价策略    │  │
│  │ • 成本监控        │  │ • 质量指标定义    │  │ • 加价规则    │  │
│  └──────────────────┘  └──────────────────┘  └──────────────┘  │
│           ↓                      ↓                     ↓         │
│  ┌────────────────────────────────────────────────────────────┐ │
│  │                    智能路由决策引擎                          │ │
│  │                                                              │ │
│  │  输入：客户ID + 模型 + 服务等级                              │ │
│  │  决策：选择最优的 渠道×折扣组 组合                           │ │
│  │  目标：成本最低 / 利润最高 / 质量保障                        │ │
│  └────────────────────────────────────────────────────────────┘ │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘
         ↓                           ↓                      ↓
┌─────────────────┐      ┌─────────────────┐      ┌─────────────┐
│ 上游供应商       │      │ 内部管理         │      │ 下游客户     │
│                 │      │                 │      │             │
│ 渠道A-号池       │      │ 折扣组管理       │      │ 看到：标准版  │
│ 渠道A-官转       │  →   │ 服务等级定义     │  →   │ 不知道：官转  │
│ 渠道B-优质官转   │      │ 映射规则         │      │             │
└─────────────────┘      └─────────────────┘      └─────────────┘
```

### 2.3 核心数据模型设计

#### 2.3.1 渠道折扣组成本表（channel_discount_group）

**用途：** 管理上游供应商的真实折扣组和成本信息（内部管理，不对客户暴露）

```sql
CREATE TABLE channel_discount_group (
    id INT PRIMARY KEY AUTO_INCREMENT,
    channel_id INT NOT NULL,                    -- 渠道ID
    model_name VARCHAR(100) NOT NULL,           -- 模型名称
    discount_group VARCHAR(50) NOT NULL,        -- 折扣组：pool/reverse/official/premium_official
    
    -- 成本信息（上游供应商报价）
    input_cost DECIMAL(20,10) NOT NULL,         -- 输入成本（每1k tokens）
    output_cost DECIMAL(20,10) NOT NULL,        -- 输出成本（每1k tokens）
    
    -- 供应商信息
    supplier_name VARCHAR(100),                 -- 供应商名称
    supplier_contact VARCHAR(200),              -- 供应商联系方式
    contract_id VARCHAR(100),                   -- 合同编号
    
    -- 质量指标（实际监控数据）
    avg_response_time INT,                      -- 平均响应时间（ms）
    success_rate DECIMAL(5,2),                  -- 成功率（%）
    availability DECIMAL(5,2),                  -- 可用性（%）
    
    -- 风险评估
    risk_level VARCHAR(20),                     -- 风险等级：low/medium/high
    risk_notes TEXT,                            -- 风险说明
    
    -- 额外配置
    extra_ratios JSON,                          -- 额外价格倍率
    priority INT DEFAULT 0,                     -- 优先级（同折扣组内的优先级）
    weight INT DEFAULT 1,                       -- 权重（负载均衡用）
    enabled BOOLEAN DEFAULT TRUE,               -- 是否启用
    
    -- 时间戳
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    UNIQUE KEY uk_channel_model_group (channel_id, model_name, discount_group),
    INDEX idx_model_group (model_name, discount_group),
    INDEX idx_channel (channel_id),
    INDEX idx_discount_group (discount_group)
);
```

**字段说明：**
- `discount_group`：真实的折扣组名称（pool/reverse/official/premium_official）
- `input_cost/output_cost`：上游供应商的真实成本价
- `risk_level`：风险评估，用于路由决策
- `priority/weight`：同一折扣组内多个渠道的选择策略

#### 2.3.2 服务等级定义表（service_tier_definition）

**用途：** 定义面向客户的服务等级，隐藏底层折扣组细节

```sql
CREATE TABLE service_tier_definition (
    id INT PRIMARY KEY AUTO_INCREMENT,
    tier_code VARCHAR(50) NOT NULL UNIQUE,      -- 服务等级代码：economy/standard/professional/enterprise
    tier_name VARCHAR(100) NOT NULL,            -- 显示名称：经济版/标准版/专业版/企业版
    tier_name_en VARCHAR(100),                  -- 英文名称
    
    -- 客户侧描述
    description TEXT,                           -- 服务等级描述
    features JSON,                              -- 特性列表
    sla_description VARCHAR(500),               -- SLA描述（客户视角）
    
    -- 映射到折扣组
    mapped_discount_groups JSON NOT NULL,       -- 映射的折扣组列表
    
    -- 显示配置
    display_order INT DEFAULT 0,                -- 显示顺序
    is_public BOOLEAN DEFAULT TRUE,             -- 是否对客户可见
    is_default BOOLEAN DEFAULT FALSE,           -- 是否为默认等级
    
    -- 状态
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_tier_code (tier_code),
    INDEX idx_display_order (display_order)
);
```

**mapped_discount_groups JSON 结构示例：**

```json
{
  "primary": ["official"],           // 主要使用的折扣组
  "fallback": ["reverse"],           // 降级备选折扣组
  "excluded": ["pool"],              // 明确排除的折扣组
  "selection_strategy": "cost_first" // 选择策略：cost_first/quality_first
}
```

**初始化数据示例：**

```sql
INSERT INTO service_tier_definition (tier_code, tier_name, tier_name_en, description, mapped_discount_groups, display_order) VALUES
('economy', '经济版', 'Economy', '性价比之选，适合测试和开发环境', 
 '{"primary": ["pool", "reverse"], "fallback": [], "excluded": ["premium_official"], "selection_strategy": "cost_first"}', 1),

('standard', '标准版', 'Standard', '稳定可靠，适合生产环境', 
 '{"primary": ["official"], "fallback": ["reverse"], "excluded": [], "selection_strategy": "balanced"}', 2),

('professional', '专业版', 'Professional', '高可用性，适合关键业务', 
 '{"primary": ["premium_official"], "fallback": ["official"], "excluded": ["pool"], "selection_strategy": "quality_first"}', 3),

('enterprise', '企业版', 'Enterprise', '最高SLA，专属支持', 
 '{"primary": ["premium_official"], "fallback": [], "excluded": ["pool", "reverse"], "selection_strategy": "quality_first"}', 4);
```

#### 2.3.3 客户定价策略表（customer_pricing_strategy）

```sql
CREATE TABLE customer_pricing_strategy (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,                       -- 用户ID（客户）
    strategy_name VARCHAR(100) NOT NULL,        -- 策略名称
    strategy_type VARCHAR(50) NOT NULL,         -- 策略类型：model_based/tier_based/custom
    
    -- 定价配置
    pricing_config JSON NOT NULL,               -- 定价配置（详见下文）
    
    -- 质量要求
    min_quality_tier VARCHAR(50),               -- 最低质量要求
    preferred_quality_tier VARCHAR(50),         -- 偏好质量等级
    
    -- 生效条件
    effective_from TIMESTAMP,                   -- 生效开始时间
    effective_to TIMESTAMP,                     -- 生效结束时间
    priority INT DEFAULT 0,                     -- 优先级
    enabled BOOLEAN DEFAULT TRUE,               -- 是否启用
    
    -- 时间戳
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_user (user_id),
    INDEX idx_effective (effective_from, effective_to)
);
```

**pricing_config JSON 结构示例：**

```json
{
  "type": "tier_based",  // tier_based: 基于服务等级定价
  "rules": [
    {
      "model_pattern": "gpt-4*",
      "service_tier": "standard",  // 使用服务等级，而非折扣组
      "pricing": {
        "input": 15.0,
        "output": 30.0,
        "markup_type": "fixed",  // fixed: 固定价格
        "markup_value": 0
      }
    },
    {
      "model_pattern": "gpt-4*",
      "service_tier": "professional",
      "pricing": {
        "input": 18.0,
        "output": 36.0,
        "markup_type": "fixed",
        "markup_value": 0
      }
    },
    {
      "model_pattern": "gpt-3.5*",
      "service_tier": "standard",
      "pricing": {
        "input": 1.0,
        "output": 2.0,
        "markup_type": "percentage",  // percentage: 基于成本加价
        "markup_value": 30  // 30% markup
      }
    }
  ],
  "default_markup": {
    "type": "percentage",
    "value": 50  // 默认50%加价
  },
  "allowed_service_tiers": ["economy", "standard", "professional"]  // 该客户可用的服务等级
}
```

#### 2.3.4 折扣组定义表（discount_group_definition）

**用途：** 定义上游折扣组的标准分类和特征（内部管理）

```sql
CREATE TABLE discount_group_definition (
    id INT PRIMARY KEY AUTO_INCREMENT,
    group_code VARCHAR(50) NOT NULL UNIQUE,    -- 折扣组代码：pool/reverse/official/premium_official
    group_name VARCHAR(100) NOT NULL,          -- 折扣组名称
    group_name_en VARCHAR(100),                -- 英文名称
    
    -- 特征描述（内部管理用）
    description TEXT,                          -- 描述
    typical_cost_level VARCHAR(20),            -- 典型成本水平：very_low/low/medium/high
    typical_quality_level VARCHAR(20),         -- 典型质量水平：low/medium/high/very_high
    risk_level VARCHAR(20),                    -- 风险等级：low/medium/high/very_high
    
    -- 质量指标基准
    expected_success_rate DECIMAL(5,2),        -- 预期成功率
    expected_response_time INT,                -- 预期响应时间
    
    -- 管理配置
    requires_approval BOOLEAN DEFAULT FALSE,   -- 是否需要审批才能使用
    internal_notes TEXT,                       -- 内部备注
    
    -- 状态
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_group_code (group_code)
);
```

**初始化数据示例：**

```sql
INSERT INTO discount_group_definition (group_code, group_name, group_name_en, description, typical_cost_level, typical_quality_level, risk_level, expected_success_rate, expected_response_time) VALUES
('pool', '号池', 'Pool', '共享账号池，成本极低但稳定性差，有封号风险', 'very_low', 'low', 'very_high', 95.0, 3000),
('reverse', '逆向', 'Reverse', '逆向工程接口，成本低但有合规风险', 'low', 'medium', 'high', 97.0, 2000),
('official', '官转', 'Official', '官方转售渠道，稳定可靠', 'medium', 'high', 'low', 99.5, 1500),
('premium_official', '优质官转', 'Premium Official', '官方优质渠道，最高稳定性和SLA保障', 'high', 'very_high', 'very_low', 99.9, 800);
```

#### 2.3.5 客户服务等级授权表（customer_service_tier_access）

**用途：** 控制客户可以使用哪些服务等级

```sql
CREATE TABLE customer_service_tier_access (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,                      -- 用户ID
    service_tier VARCHAR(50) NOT NULL,         -- 服务等级代码
    
    -- 访问控制
    is_allowed BOOLEAN DEFAULT TRUE,           -- 是否允许使用
    is_default BOOLEAN DEFAULT FALSE,          -- 是否为该客户的默认等级
    
    -- 限制条件
    max_requests_per_day INT,                  -- 每日最大请求数（可选）
    max_tokens_per_request INT,                -- 单次请求最大tokens（可选）
    
    -- 生效时间
    effective_from TIMESTAMP,
    effective_to TIMESTAMP,
    
    -- 备注
    notes TEXT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    UNIQUE KEY uk_user_tier (user_id, service_tier),
    INDEX idx_user (user_id),
    INDEX idx_tier (service_tier)
);
```

#### 2.3.6 定价规则路由表（pricing_route_rule）

```sql
CREATE TABLE pricing_route_rule (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,                      -- 用户ID
    token_id INT,                              -- 令牌ID（可选，更细粒度）
    
    -- 路由策略
    route_strategy VARCHAR(50) NOT NULL,       -- cost_first/quality_first/balanced
    
    -- 成本控制
    max_cost_per_request DECIMAL(20,10),       -- 单次请求最大成本
    target_profit_margin DECIMAL(5,2),         -- 目标利润率（%）
    
    -- 质量控制
    quality_tier_preference VARCHAR(50),       -- 质量偏好
    allow_tier_downgrade BOOLEAN DEFAULT FALSE,-- 允许降级
    
    -- 高级配置
    config JSON,                               -- 扩展配置
    
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_user_token (user_id, token_id)
);
```

### 2.4 核心业务流程设计

#### 2.4.1 完整的请求处理流程

```
客户请求
    ↓
[1] 解析请求参数
    - 客户ID (从Token获取)
    - 模型名称 (gpt-4)
    - 服务等级 (可选，如：standard)
    ↓
[2] 确定服务等级
    - 如果客户指定了服务等级 → 验证客户是否有权限
    - 如果未指定 → 使用客户的默认服务等级
    ↓
[3] 查询服务等级映射
    - 根据服务等级查询映射的折扣组列表
    - 例如：standard → [official, reverse(fallback)]
    ↓
[4] 查询可用渠道
    - 查询所有提供该模型+折扣组的渠道
    - 过滤：状态启用、未在冷却期、权重>0
    ↓
[5] 计算成本和收益
    - 对每个候选渠道：
      * 获取上游成本 (input_cost, output_cost)
      * 获取客户价格 (根据定价策略)
      * 计算预估利润
    ↓
[6] 路由决策
    - 根据路由策略选择最优渠道：
      * cost_first: 成本最低
      * profit_first: 利润最高
      * quality_first: 质量最好
      * balanced: 综合评分
    ↓
[7] 执行请求
    - 调用选中的渠道
    - 记录实际使用情况
    ↓
[8] 计费和统计
    - 按客户价格扣费
    - 记录成本和利润
    - 更新统计数据
```

#### 2.4.2 核心数据结构设计

```go
// model/pricing_router.go
package model

// 定价路由器
type PricingRouter struct {
    // 缓存
    discountGroupCache map[string][]*ChannelDiscountGroup  // model -> discount groups
    serviceTierCache map[string]*ServiceTierDefinition     // tier_code -> definition
    customerPricingCache map[int]*CustomerPricingStrategy  // user_id -> strategy
    
    sync.RWMutex
}

// 渠道折扣组（上游成本）
type ChannelDiscountGroup struct {
    ID              int
    ChannelID       int
    Channel         *Channel
    ModelName       string
    DiscountGroup   string    // pool/reverse/official/premium_official
    InputCost       float64   // 上游成本
    OutputCost      float64
    SupplierName    string
    AvgResponseTime int
    SuccessRate     float64
    RiskLevel       string
    Priority        int
    Weight          int
    Enabled         bool
}

// 服务等级定义（客户侧抽象）
type ServiceTierDefinition struct {
    ID                     int
    TierCode               string   // economy/standard/professional/enterprise
    TierName               string   // 经济版/标准版/专业版/企业版
    Description            string
    MappedDiscountGroups   DiscountGroupMapping
    DisplayOrder           int
    IsPublic               bool
    Enabled                bool
}

// 折扣组映射配置
type DiscountGroupMapping struct {
    Primary           []string  // 主要使用的折扣组
    Fallback          []string  // 降级备选
    Excluded          []string  // 明确排除
    SelectionStrategy string    // cost_first/quality_first/balanced
}

// 路由决策结果
type RouteDecision struct {
    // 选中的渠道和折扣组
    Channel         *Channel
    DiscountGroup   *ChannelDiscountGroup
    
    // 服务等级信息
    ServiceTier     string
    
    // 成本和收益
    EstimatedInputCost  float64
    EstimatedOutputCost float64
    CustomerInputPrice  float64
    CustomerOutputPrice float64
    
    // 预估指标（基于预估tokens）
    EstimatedCost     float64
    EstimatedRevenue  float64
    EstimatedProfit   float64
    ProfitMargin      float64  // 利润率 %
    
    // 质量指标
    ExpectedSuccessRate float64
    ExpectedResponseTime int
    RiskLevel           string
}

#### 2.4.3 核心算法：智能路由选择

```go
// 智能路由选择（核心方法）
func (pr *PricingRouter) SelectOptimalChannel(
    userId int,
    modelName string,
    serviceTier string,  // 客户指定的服务等级（可选）
    estimatedInputTokens int,
    estimatedOutputTokens int,
) (*RouteDecision, error) {
    
    // 步骤1: 确定服务等级
    finalServiceTier, err := pr.determineServiceTier(userId, serviceTier)
    if err != nil {
        return nil, err
    }
    
    // 步骤2: 获取服务等级映射的折扣组
    tierDef := pr.getServiceTierDefinition(finalServiceTier)
    if tierDef == nil {
        return nil, errors.New("service tier not found")
    }
    
    // 步骤3: 查询符合条件的渠道×折扣组组合
    eligibleChannels := pr.getEligibleChannelDiscountGroups(
        modelName,
        tierDef.MappedDiscountGroups,
    )
    
    if len(eligibleChannels) == 0 {
        return nil, errors.New("no eligible channels found")
    }
    
    // 步骤4: 获取客户定价策略
    customerPricing := pr.getCustomerPricingStrategy(userId)
    
    // 步骤5: 计算每个候选渠道的成本、收益和利润
    decisions := make([]*RouteDecision, 0)
    for _, channelGroup := range eligibleChannels {
        decision := pr.calculateRouteDecision(
            channelGroup,
            customerPricing,
            finalServiceTier,
            estimatedInputTokens,
            estimatedOutputTokens,
        )
        decisions = append(decisions, decision)
    }
    
    // 步骤6: 根据路由策略选择最优渠道
    routeRule := pr.getRouteRule(userId)
    selectionStrategy := tierDef.MappedDiscountGroups.SelectionStrategy
    if routeRule != nil && routeRule.RouteStrategy != "" {
        selectionStrategy = routeRule.RouteStrategy
    }
    
    optimal := pr.selectByStrategy(decisions, selectionStrategy)
    
    return optimal, nil
}

// 确定服务等级
func (pr *PricingRouter) determineServiceTier(userId int, requestedTier string) (string, error) {
    // 如果客户指定了服务等级
    if requestedTier != "" {
        // 验证客户是否有权限使用该服务等级
        hasAccess := pr.checkServiceTierAccess(userId, requestedTier)
        if !hasAccess {
            return "", errors.New("no access to requested service tier")
        }
        return requestedTier, nil
    }
    
    // 如果未指定，使用客户的默认服务等级
    defaultTier := pr.getCustomerDefaultServiceTier(userId)
    if defaultTier == "" {
        // 如果客户没有设置默认等级，使用系统默认
        defaultTier = "standard"
    }
    
    return defaultTier, nil
}

// 获取符合条件的渠道×折扣组组合
func (pr *PricingRouter) getEligibleChannelDiscountGroups(
    modelName string,
    mapping DiscountGroupMapping,
) []*ChannelDiscountGroup {
    
    pr.RLock()
    defer pr.RUnlock()
    
    allGroups := pr.discountGroupCache[modelName]
    if allGroups == nil {
        return nil
    }
    
    eligible := make([]*ChannelDiscountGroup, 0)
    
    // 优先使用 primary 折扣组
    for _, group := range allGroups {
        if !group.Enabled {
            continue
        }
        
        // 检查是否在排除列表中
        if contains(mapping.Excluded, group.DiscountGroup) {
            continue
        }
        
        // 检查是否在主要列表中
        if contains(mapping.Primary, group.DiscountGroup) {
            eligible = append(eligible, group)
        }
    }
    
    // 如果主要折扣组没有可用渠道，使用 fallback
    if len(eligible) == 0 {
        for _, group := range allGroups {
            if !group.Enabled {
                continue
            }
            if contains(mapping.Fallback, group.DiscountGroup) {
                eligible = append(eligible, group)
            }
        }
    }
    
    return eligible
}

// 计算路由决策（成本、收益、利润）
func (pr *PricingRouter) calculateRouteDecision(
    channelGroup *ChannelDiscountGroup,
    customerPricing *CustomerPricingStrategy,
    serviceTier string,
    estimatedInputTokens int,
    estimatedOutputTokens int,
) *RouteDecision {
    
    // 获取客户价格
    customerInputPrice, customerOutputPrice := pr.getCustomerPrice(
        customerPricing,
        channelGroup.ModelName,
        serviceTier,
        channelGroup.InputCost,
        channelGroup.OutputCost,
    )
    
    // 计算预估成本和收益
    estimatedCost := (float64(estimatedInputTokens) / 1000.0 * channelGroup.InputCost) +
                     (float64(estimatedOutputTokens) / 1000.0 * channelGroup.OutputCost)
    
    estimatedRevenue := (float64(estimatedInputTokens) / 1000.0 * customerInputPrice) +
                        (float64(estimatedOutputTokens) / 1000.0 * customerOutputPrice)
    
    estimatedProfit := estimatedRevenue - estimatedCost
    profitMargin := 0.0
    if estimatedRevenue > 0 {
        profitMargin = (estimatedProfit / estimatedRevenue) * 100.0
    }
    
    return &RouteDecision{
        Channel:              channelGroup.Channel,
        DiscountGroup:        channelGroup,
        ServiceTier:          serviceTier,
        EstimatedInputCost:   channelGroup.InputCost,
        EstimatedOutputCost:  channelGroup.OutputCost,
        CustomerInputPrice:   customerInputPrice,
        CustomerOutputPrice:  customerOutputPrice,
        EstimatedCost:        estimatedCost,
        EstimatedRevenue:     estimatedRevenue,
        EstimatedProfit:      estimatedProfit,
        ProfitMargin:         profitMargin,
        ExpectedSuccessRate:  channelGroup.SuccessRate,
        ExpectedResponseTime: channelGroup.AvgResponseTime,
        RiskLevel:            channelGroup.RiskLevel,
    }
}

// 获取客户价格
func (pr *PricingRouter) getCustomerPrice(
    strategy *CustomerPricingStrategy,
    modelName string,
    serviceTier string,
    baseCostInput float64,
    baseCostOutput float64,
) (inputPrice float64, outputPrice float64) {
    
    // 查找匹配的定价规则
    rule := pr.findMatchingPricingRule(strategy, modelName, serviceTier)
    
    if rule != nil {
        if rule.MarkupType == "fixed" {
            // 固定价格
            return rule.InputPrice, rule.OutputPrice
        } else {
            // 百分比加价
            markup := 1.0 + (rule.MarkupValue / 100.0)
            return baseCostInput * markup, baseCostOutput * markup
        }
    }
    
    // 使用默认加价策略
    defaultMarkup := 1.0 + (strategy.DefaultMarkup / 100.0)
    return baseCostInput * defaultMarkup, baseCostOutput * defaultMarkup
}

// 根据策略选择最优方案
func (pr *PricingRouter) selectByStrategy(
    decisions []*RouteDecision,
    strategy string,
) *RouteDecision {
    
    if len(decisions) == 0 {
        return nil
    }
    
    if len(decisions) == 1 {
        return decisions[0]
    }
    
    switch strategy {
    case "cost_first":
        // 成本优先：选择成本最低的
        return pr.selectLowestCost(decisions)
        
    case "quality_first":
        // 质量优先：选择质量最高的（成功率最高、响应时间最短）
        return pr.selectHighestQuality(decisions)
        
    case "profit_first":
        // 利润优先：选择利润率最高的
        return pr.selectHighestProfit(decisions)
        
    case "balanced":
        // 平衡策略：综合评分
        return pr.selectBalanced(decisions)
        
    default:
        return pr.selectBalanced(decisions)
    }
}

// 成本优先选择
func (pr *PricingRouter) selectLowestCost(decisions []*RouteDecision) *RouteDecision {
    lowest := decisions[0]
    for _, d := range decisions[1:] {
        if d.EstimatedCost < lowest.EstimatedCost {
            lowest = d
        }
    }
    return lowest
}

// 质量优先选择
func (pr *PricingRouter) selectHighestQuality(decisions []*RouteDecision) *RouteDecision {
    best := decisions[0]
    for _, d := range decisions[1:] {
        // 综合考虑成功率和响应时间
        currentScore := d.ExpectedSuccessRate - float64(d.ExpectedResponseTime)/10000.0
        bestScore := best.ExpectedSuccessRate - float64(best.ExpectedResponseTime)/10000.0
        if currentScore > bestScore {
            best = d
        }
    }
    return best
}

// 利润优先选择
func (pr *PricingRouter) selectHighestProfit(decisions []*RouteDecision) *RouteDecision {
    highest := decisions[0]
    for _, d := range decisions[1:] {
        if d.ProfitMargin > highest.ProfitMargin {
            highest = d
        }
    }
    return highest
}

// 平衡策略评分算法
func (pr *PricingRouter) selectBalanced(decisions []*RouteDecision) *RouteDecision {
    type scoredDecision struct {
        decision *RouteDecision
        score    float64
    }
    
    scored := make([]*scoredDecision, 0)
    
    // 计算归一化所需的最大最小值
    maxCost, minCost := pr.getMaxMinCost(decisions)
    maxProfit, minProfit := pr.getMaxMinProfit(decisions)
    maxQuality, minQuality := pr.getMaxMinQuality(decisions)
    
    for _, d := range decisions {
        // 成本分数（成本越低分数越高）
        costScore := 0.0
        if maxCost > minCost {
            costScore = 1.0 - (d.EstimatedCost-minCost)/(maxCost-minCost)
        }
        
        // 利润分数（利润率越高分数越高）
        profitScore := 0.0
        if maxProfit > minProfit {
            profitScore = (d.ProfitMargin-minProfit)/(maxProfit-minProfit)
        }
        
        // 质量分数（成功率越高、响应时间越短分数越高）
        qualityValue := d.ExpectedSuccessRate - float64(d.ExpectedResponseTime)/10000.0
        qualityScore := 0.0
        if maxQuality > minQuality {
            qualityScore = (qualityValue-minQuality)/(maxQuality-minQuality)
        }
        
        // 综合评分（权重可配置）
        // 默认权重：利润40%、质量40%、成本20%
        score := 0.4*profitScore + 0.4*qualityScore + 0.2*costScore
        
        scored = append(scored, &scoredDecision{
            decision: d,
            score:    score,
        })
    }
    
    // 选择得分最高的
    sort.Slice(scored, func(i, j int) bool {
        return scored[i].score > scored[j].score
    })
    
    return scored[0].decision
}
```

### 2.5 管理后台功能设计

#### 2.5.1 上游成本管理模块

**页面：渠道折扣组管理**

**功能列表：**

1. **折扣组配置**
   - 为每个渠道配置支持的折扣组（号池、逆向、官转、优质官转）
   - 设置每个折扣组的成本价格（input_cost, output_cost）
   - 配置供应商信息和合同编号
   - 批量导入功能（Excel/CSV）

2. **成本监控面板**
   - 实时成本对比：同一模型不同渠道×折扣组的成本对比
   - 成本趋势图：历史成本变化趋势
   - 成本预警：成本异常上涨自动告警

3. **质量监控**
   - 实时监控各渠道×折扣组的成功率、响应时间
   - 质量评级：自动根据监控数据评级
   - 风险预警：低质量渠道自动标记

4. **供应商管理**
   - 供应商档案：联系方式、合同信息
   - 供应商评分：基于成本、质量、稳定性的综合评分
   - 供应商对比：多供应商横向对比

**界面示例：**

```
┌─────────────────────────────────────────────────────────────┐
│ 渠道折扣组管理                                    [+ 新增配置] │
├─────────────────────────────────────────────────────────────┤
│ 筛选：渠道 [全部▼] 模型 [gpt-4▼] 折扣组 [全部▼]  [搜索]      │
├─────────────────────────────────────────────────────────────┤
│ 渠道  │ 模型   │ 折扣组    │ 输入成本 │ 输出成本 │ 成功率 │ 操作│
├─────────────────────────────────────────────────────────────┤
│ 渠道A │ gpt-4  │ 号池      │ 5.0     │ 10.0    │ 95.2% │ 编辑│
│ 渠道A │ gpt-4  │ 官转      │ 12.0    │ 24.0    │ 99.5% │ 编辑│
│ 渠道B │ gpt-4  │ 优质官转  │ 14.0    │ 28.0    │ 99.9% │ 编辑│
│ 渠道C │ gpt-4  │ 逆向      │ 8.0     │ 16.0    │ 97.3% │ 编辑│
└─────────────────────────────────────────────────────────────┘
```

#### 2.5.2 服务等级管理模块

**页面：服务等级配置**

**功能列表：**

1. **服务等级定义**
   - 创建/编辑服务等级（经济版、标准版、专业版、企业版）
   - 设置客户侧显示名称和描述
   - 配置SLA描述（客户视角）

2. **折扣组映射**
   - 为每个服务等级配置映射的折扣组
   - 设置主要折扣组（primary）和降级备选（fallback）
   - 设置排除的折扣组（excluded）
   - 配置选择策略（cost_first/quality_first/balanced）

3. **服务等级预览**
   - 客户视角预览：查看客户看到的服务等级信息
   - 内部视角：查看实际映射的折扣组

**界面示例：**

```
┌─────────────────────────────────────────────────────────────┐
│ 服务等级配置                                      [+ 新增等级] │
├─────────────────────────────────────────────────────────────┤
│ 等级代码：standard                                           │
│ 显示名称：标准版                                             │
│ 英文名称：Standard                                           │
│ 客户描述：稳定可靠，适合生产环境                              │
│                                                              │
│ 折扣组映射：                                                 │
│   主要使用：☑ 官转  ☐ 优质官转  ☐ 逆向  ☐ 号池              │
│   降级备选：☐ 官转  ☐ 优质官转  ☑ 逆向  ☐ 号池              │
│   明确排除：☐ 官转  ☐ 优质官转  ☐ 逆向  ☑ 号池              │
│                                                              │
│ 选择策略：◉ 平衡策略  ○ 成本优先  ○ 质量优先                │
│                                                              │
│ [保存] [取消]                                                │
└─────────────────────────────────────────────────────────────┘
```

#### 2.5.3 客户定价管理模块

**页面：客户定价策略**

**功能列表：**

1. **客户定价配置**
   - 为客户创建专属定价策略
   - 支持按模型×服务等级配置价格
   - 支持固定价格和百分比加价两种模式
   - 批量配置工具

2. **定价模板**
   - 预设定价模板（VIP客户、标准客户、试用客户）
   - 一键应用模板到客户
   - 模板克隆和修改

3. **服务等级授权**
   - 配置客户可使用的服务等级
   - 设置默认服务等级
   - 设置使用限制（每日请求数、单次tokens等）

4. **价格审批**
   - 特殊定价申请和审批流程
   - 价格变更历史记录
   - 审批日志

**界面示例：**

```
┌─────────────────────────────────────────────────────────────┐
│ 客户定价策略 - 客户A                          [应用模板▼]     │
├─────────────────────────────────────────────────────────────┤
│ 基础配置：                                                   │
│   策略名称：客户A专属定价                                    │
│   默认加价：30 %                                             │
│                                                              │
│ 服务等级授权：                                               │
│   ☑ 经济版  ☑ 标准版 (默认)  ☑ 专业版  ☐ 企业版             │
│                                                              │
│ 定价规则：                                        [+ 添加规则]│
├─────────────────────────────────────────────────────────────┤
│ 模型      │ 服务等级 │ 定价方式 │ 输入价格 │ 输出价格 │ 操作 │
├─────────────────────────────────────────────────────────────┤
│ gpt-4*    │ 标准版   │ 固定价格 │ 15.0    │ 30.0    │ 编辑 │
│ gpt-4*    │ 专业版   │ 固定价格 │ 18.0    │ 36.0    │ 编辑 │
│ gpt-3.5*  │ 标准版   │ 加价30%  │ -       │ -       │ 编辑 │
│ *         │ *        │ 加价30%  │ -       │ -       │ 默认 │
└─────────────────────────────────────────────────────────────┘
```

#### 2.5.4 路由策略配置模块

**页面：路由策略管理**

**功能列表：**

1. **全局路由策略**
   - 设置默认路由算法（成本优先/质量优先/利润优先/平衡）
   - 配置平衡策略的权重（成本、质量、利润）
   - 设置降级策略

2. **客户路由策略**
   - 为特定客户配置专属路由策略
   - 覆盖全局策略
   - 设置成本上限和质量下限

3. **路由模拟**
   - 输入模型和服务等级，模拟路由决策
   - 查看候选渠道和评分
   - 对比不同策略的结果

#### 2.5.5 数据分析模块

**页面：成本利润分析**

**功能列表：**

1. **成本分析**
   - 总成本趋势
   - 按渠道、折扣组、模型的成本分布
   - 成本节省统计（智能路由带来的节省）

2. **收入分析**
   - 总收入趋势
   - 按客户、模型、服务等级的收入分布
   - 客户价值分析

3. **利润分析**
   - 利润率趋势
   - 按客户、模型的利润分析
   - 高利润/低利润客户识别

4. **质量报告**
   - 各服务等级的质量达成率
   - 渠道质量排名
   - 质量问题统计

5. **路由效果分析**
   - 路由决策分布（各策略使用频率）
   - 路由优化效果（成本节省、利润提升）
   - 渠道利用率分析

### 2.6 API接口设计

#### 2.6.1 客户侧API（向后兼容）

**标准调用（不指定服务等级）**

```http
POST /v1/chat/completions
Headers:
  Authorization: Bearer sk-xxx
  
Body: {
  "model": "gpt-4",
  "messages": [...]
}

Response: {
  "id": "chatcmpl-xxx",
  "model": "gpt-4",
  "choices": [...],
  "usage": {
    "prompt_tokens": 100,
    "completion_tokens": 50,
    "total_tokens": 150
  }
}
```
> 系统自动使用客户的默认服务等级（如：标准版）

**指定服务等级调用**

```http
POST /v1/chat/completions
Headers:
  Authorization: Bearer sk-xxx
  X-Service-Tier: professional  // 新增：指定服务等级
  
Body: {
  "model": "gpt-4",
  "messages": [...]
}

Response: {
  "id": "chatcmpl-xxx",
  "model": "gpt-4",
  "choices": [...],
  "usage": {
    "prompt_tokens": 100,
    "completion_tokens": 50,
    "total_tokens": 150
  },
  "x_service_tier": "professional"  // 返回实际使用的服务等级
}
```

**查询可用服务等级**

```http
GET /v1/service-tiers
Headers:
  Authorization: Bearer sk-xxx

Response: {
  "available_tiers": [
    {
      "tier_code": "economy",
      "tier_name": "经济版",
      "description": "性价比之选，适合测试和开发环境",
      "is_default": false
    },
    {
      "tier_code": "standard",
      "tier_name": "标准版",
      "description": "稳定可靠，适合生产环境",
      "is_default": true
    },
    {
      "tier_code": "professional",
      "tier_name": "专业版",
      "description": "高可用性，适合关键业务",
      "is_default": false
    }
  ]
}
```

**查询定价信息（可选功能）**

```http
GET /v1/pricing?model=gpt-4&service_tier=standard
Headers:
  Authorization: Bearer sk-xxx

Response: {
  "model": "gpt-4",
  "service_tier": "standard",
  "pricing": {
    "input_price": 15.0,
    "output_price": 30.0,
    "unit": "per 1K tokens",
    "currency": "CNY"
  }
}
```

#### 2.6.2 管理API

**1. 渠道折扣组管理**

```http
# 创建渠道折扣组配置
POST /api/admin/channel-discount-groups
Body: {
  "channel_id": 1,
  "model_name": "gpt-4",
  "discount_group": "official",
  "input_cost": 12.0,
  "output_cost": 24.0,
  "supplier_name": "供应商A",
  "priority": 10,
  "weight": 5
}

# 查询渠道折扣组
GET /api/admin/channel-discount-groups?channel_id=1&model=gpt-4

# 更新
PUT /api/admin/channel-discount-groups/:id

# 删除
DELETE /api/admin/channel-discount-groups/:id

# 批量导入
POST /api/admin/channel-discount-groups/batch-import
Body: FormData (Excel/CSV file)
```

**2. 服务等级管理**

```http
# 创建服务等级
POST /api/admin/service-tiers
Body: {
  "tier_code": "standard",
  "tier_name": "标准版",
  "description": "稳定可靠，适合生产环境",
  "mapped_discount_groups": {
    "primary": ["official"],
    "fallback": ["reverse"],
    "excluded": ["pool"],
    "selection_strategy": "balanced"
  }
}

# 查询服务等级列表
GET /api/admin/service-tiers

# 更新
PUT /api/admin/service-tiers/:id

# 删除
DELETE /api/admin/service-tiers/:id
```

**3. 客户定价策略管理**

```http
# 创建客户定价策略
POST /api/admin/customer-pricing-strategies
Body: {
  "user_id": 123,
  "strategy_name": "客户A专属定价",
  "pricing_config": {
    "type": "tier_based",
    "rules": [...],
    "default_markup": {"type": "percentage", "value": 30}
  }
}

# 查询客户定价策略
GET /api/admin/customer-pricing-strategies/:userId

# 更新
PUT /api/admin/customer-pricing-strategies/:id

# 应用定价模板
POST /api/admin/customer-pricing-strategies/apply-template
Body: {
  "user_id": 123,
  "template_id": 1
}
```

**4. 客户服务等级授权**

```http
# 授权客户使用服务等级
POST /api/admin/customer-service-tier-access
Body: {
  "user_id": 123,
  "service_tier": "professional",
  "is_allowed": true,
  "is_default": false
}

# 查询客户授权
GET /api/admin/customer-service-tier-access/:userId

# 批量授权
POST /api/admin/customer-service-tier-access/batch
Body: {
  "user_ids": [123, 456, 789],
  "service_tier": "standard",
  "is_allowed": true
}
```

**5. 路由策略管理**

```http
# 创建路由规则
POST /api/admin/route-rules
Body: {
  "user_id": 123,
  "route_strategy": "balanced",
  "target_profit_margin": 30.0,
  "quality_tier_preference": "standard"
}

# 查询路由规则
GET /api/admin/route-rules/:userId

# 路由模拟
POST /api/admin/route-rules/simulate
Body: {
  "user_id": 123,
  "model": "gpt-4",
  "service_tier": "standard",
  "estimated_input_tokens": 1000,
  "estimated_output_tokens": 500
}

Response: {
  "selected_channel": {...},
  "discount_group": "official",
  "estimated_cost": 18.0,
  "customer_price": 22.5,
  "profit_margin": 25.0,
  "alternatives": [...]  // 其他候选方案
}
```

**6. 数据分析API**

```http
# 成本分析
GET /api/admin/analytics/cost-analysis?start_date=2024-01-01&end_date=2024-01-31

# 利润分析
GET /api/admin/analytics/profit-analysis?group_by=customer

# 质量报告
GET /api/admin/analytics/quality-report?service_tier=standard

# 路由效果分析
GET /api/admin/analytics/routing-effectiveness
```

## 三、典型使用场景

### 3.1 场景一：新客户接入

**业务流程：**

1. **销售签约阶段**
   - 销售与客户谈判，确定服务等级和价格
   - 例如：客户B购买"标准版"服务，gpt-4 定价为 15元/1K input tokens

2. **管理员配置**
   ```
   步骤1: 创建客户账号
   步骤2: 配置服务等级授权
          - 授权"经济版"、"标准版"
          - 设置"标准版"为默认
   步骤3: 配置定价策略
          - gpt-4 + 标准版 = 15元/1K input, 30元/1K output
          - 其他模型使用默认加价30%
   ```

3. **客户使用**
   - 客户调用API时不指定服务等级，自动使用"标准版"
   - 系统自动路由到"官转"折扣组的渠道
   - 按15元/1K tokens计费

**系统自动处理：**
- 根据"标准版"映射，选择"官转"折扣组
- 在多个"官转"渠道中，选择成本最低的
- 确保利润率达到目标（如30%）

### 3.2 场景二：VIP客户升级

**业务需求：**
- 客户C是VIP客户，要求最高质量和SLA保障
- 愿意支付更高价格

**配置方案：**

1. **授权专业版服务**
   ```
   服务等级授权：
   - 经济版：✓
   - 标准版：✓
   - 专业版：✓ (新增)
   - 企业版：✗
   ```

2. **配置专业版定价**
   ```
   定价规则：
   - gpt-4 + 专业版 = 18元/1K input, 36元/1K output
   - gpt-3.5 + 专业版 = 1.5元/1K input, 3元/1K output
   ```

3. **客户使用**
   ```http
   POST /v1/chat/completions
   Headers:
     X-Service-Tier: professional
   Body:
     {"model": "gpt-4", ...}
   ```

**系统自动处理：**
- 根据"专业版"映射，优先选择"优质官转"折扣组
- 确保成功率>99.9%，响应时间<1秒
- 按18元/1K tokens计费

### 3.3 场景三：成本优化

**业务场景：**
- 某客户D的测试环境，对质量要求不高
- 希望降低成本

**配置方案：**

1. **授权经济版服务**
   ```
   服务等级授权：
   - 经济版：✓ (设为默认)
   - 标准版：✓
   ```

2. **配置经济版定价**
   ```
   定价规则：
   - gpt-4 + 经济版 = 8元/1K input, 16元/1K output
   - 使用成本优先路由策略
   ```

**系统自动处理：**
- 根据"经济版"映射，选择"号池"或"逆向"折扣组
- 优先选择成本最低的渠道
- 按8元/1K tokens计费
- 利润率可能较低，但客户成本也低

### 3.4 场景四：供应商价格调整

**业务场景：**
- 供应商A的"官转"折扣组价格上涨10%
- 需要快速调整

**处理流程：**

1. **更新上游成本**
   ```
   渠道折扣组管理：
   - 渠道A + gpt-4 + 官转
   - 输入成本：12.0 → 13.2 (+10%)
   - 输出成本：24.0 → 26.4 (+10%)
   ```

2. **系统自动调整**
   - 路由引擎自动重新计算
   - 如果渠道A成本过高，自动切换到渠道B的"官转"
   - 客户价格不变，利润率自动调整

3. **管理员决策**
   - 查看利润率报告
   - 如果利润率过低，考虑：
     * 调整客户价格
     * 更换供应商
     * 调整服务等级映射

### 3.5 场景五：质量问题处理

**业务场景：**
- 渠道C的"逆向"折扣组出现质量问题
- 成功率从97%降到90%

**系统自动处理：**

1. **质量监控告警**
   ```
   告警：渠道C + gpt-4 + 逆向
   - 成功率：90% (低于预期97%)
   - 建议：暂时禁用或降低权重
   ```

2. **自动降级**
   - 如果配置了自动降级策略
   - 系统自动将该渠道的权重降为0
   - 请求自动路由到其他渠道

3. **管理员处理**
   ```
   选项1: 暂时禁用该渠道×折扣组
   选项2: 联系供应商解决问题
   选项3: 调整服务等级映射，排除该折扣组
   ```

## 四、产品价值分析

### 4.1 对平台方的价值

**1. 成本控制精细化**
- **多维度成本管理**：渠道×折扣组×模型的三维成本矩阵
- **实时成本对比**：自动选择成本最优的渠道
- **预期收益**：成本降低15-25%

**2. 定价策略灵活化**
- **客户级定价**：不同客户不同价格，实现价格歧视
- **服务分级定价**：同一客户可选择不同服务等级
- **预期收益**：收入提升20-30%

**3. 利润优化自动化**
- **智能路由**：自动在成本、质量、利润之间平衡
- **动态调整**：供应商价格变化时自动优化
- **预期收益**：利润率提升5-10个百分点

**4. 风险管理体系化**
- **质量监控**：实时监控各渠道质量
- **风险隔离**：高风险折扣组（号池、逆向）与客户隔离
- **降级机制**：质量问题自动降级

**5. 运营效率提升**
- **自动化定价**：减少人工定价工作量80%
- **批量管理**：支持批量配置和模板
- **数据驱动**：完整的数据分析和决策支持

### 4.2 对客户的价值

**1. 服务质量透明化**
- 清晰的服务等级定义（经济版、标准版、专业版）
- 明确的SLA承诺
- 不暴露底层技术细节（折扣组）

**2. 灵活的服务选择**
- 根据业务场景选择不同服务等级
- 测试环境用经济版，生产环境用标准版，关键业务用专业版
- 成本与质量的自主平衡

**3. 价格可预期**
- 明确的定价标准
- 价格锁定机制
- 无隐藏费用

**4. 服务质量保障**
- 高服务等级有SLA保障
- 质量问题自动切换渠道
- 客户无感知的高可用

### 4.3 竞争优势

**1. 相比传统聚合平台**
- **传统平台**：单一价格，无服务分级
- **本方案**：多服务等级，客户可选择

**2. 相比直接对接供应商**
- **直接对接**：需要管理多个供应商，价格不透明
- **本方案**：统一接口，透明定价，自动优选

**3. 相比其他聚合平台**
- **其他平台**：简单的负载均衡，不考虑成本
- **本方案**：智能路由，成本、质量、利润综合优化

## 五、实施路线图

### 5.1 第一阶段：数据模型和基础功能（2-3周）

**目标：** 建立多层级定价的数据基础

**核心任务：**

1. **数据库设计**
   - 创建6张核心表（折扣组、服务等级、渠道折扣组、客户定价、服务等级授权、路由规则）
   - 设计索引和约束
   - 编写数据库迁移脚本

2. **数据模型开发**
   - Go struct定义
   - GORM模型映射
   - 基础CRUD方法

3. **数据迁移**
   - 现有价格数据迁移到新结构
   - 为现有渠道创建默认折扣组配置
   - 为现有用户创建默认定价策略

4. **基础API**
   - 折扣组管理API
   - 服务等级管理API
   - 基础查询API

**交付物：**
- 数据库迁移脚本
- 数据模型代码（model/）
- 数据迁移工具
- API文档（Swagger）

**验收标准：**
- 数据库表创建成功
- 现有数据迁移无误
- 基础API可正常调用

### 5.2 第二阶段：核心路由引擎（3-4周）

**目标：** 实现智能路由和定价计算

**核心任务：**

1. **定价路由器开发**
   - PricingRouter核心类
   - 服务等级映射逻辑
   - 折扣组选择逻辑

2. **路由算法实现**
   - 成本优先算法
   - 质量优先算法
   - 利润优先算法
   - 平衡策略算法

3. **定价计算**
   - 客户价格计算
   - 成本收益分析
   - 利润率计算

4. **集成现有系统**
   - 与ChannelGroup集成
   - 与relay模块集成
   - 保持向后兼容

5. **缓存机制**
   - 内存缓存（sync.Map）
   - Redis缓存
   - 缓存更新策略

6. **测试**
   - 单元测试（覆盖率>80%）
   - 集成测试
   - 性能测试

**交付物：**
- 定价路由器代码（model/pricing_router.go）
- 算法文档
- 测试用例
- 性能测试报告

**验收标准：**
- 路由算法正确性验证
- 性能满足要求（<10ms）
- 测试覆盖率达标

### 5.3 第三阶段：管理后台（3-4周）

**目标：** 提供完整的管理功能

**核心任务：**

1. **上游成本管理**
   - 渠道折扣组配置界面
   - 批量导入功能
   - 成本监控面板

2. **服务等级管理**
   - 服务等级配置界面
   - 折扣组映射配置
   - 客户视角预览

3. **客户定价管理**
   - 客户定价策略配置
   - 定价模板管理
   - 服务等级授权

4. **路由策略配置**
   - 全局路由策略
   - 客户路由策略
   - 路由模拟工具

5. **数据分析**
   - 成本分析报表
   - 利润分析报表
   - 质量报告
   - 路由效果分析

6. **前端开发**
   - React组件开发
   - 与后端API联调
   - UI/UX优化

**交付物：**
- 管理后台页面
- 用户操作手册
- 管理员培训文档
- 前端代码

**验收标准：**
- 所有管理功能可用
- 界面友好易用
- 操作流程顺畅

### 5.4 第四阶段：优化和上线（2-3周）

**目标：** 系统优化和生产部署

**核心任务：**

1. **性能优化**
   - 数据库查询优化
   - 缓存策略优化
   - 并发性能优化
   - 压力测试

2. **监控和告警**
   - Prometheus指标接入
   - 关键指标监控
   - 告警规则配置
   - 日志完善

3. **灰度发布**
   - 灰度发布方案设计
   - 功能开关配置
   - 小流量验证
   - 逐步放量

4. **文档完善**
   - 部署文档
   - 运维手册
   - 故障处理手册
   - API文档

5. **培训和支持**
   - 管理员培训
   - 客户沟通
   - 技术支持准备

**交付物：**
- 性能优化报告
- 监控配置
- 部署文档
- 运维手册
- 培训材料

**验收标准：**
- 性能指标达标
- 监控告警正常
- 灰度发布成功
- 文档完整

### 5.5 实施时间表

```
周次 │ 阶段              │ 关键里程碑
─────┼──────────────────┼────────────────────────
1-3  │ 第一阶段          │ 数据模型完成
     │ 数据模型和基础功能 │ 数据迁移完成
     │                  │ 基础API可用
─────┼──────────────────┼────────────────────────
4-7  │ 第二阶段          │ 路由引擎完成
     │ 核心路由引擎      │ 算法验证通过
     │                  │ 集成测试通过
─────┼──────────────────┼────────────────────────
8-11 │ 第三阶段          │ 管理界面完成
     │ 管理后台          │ 数据分析可用
     │                  │ 用户培训完成
─────┼──────────────────┼────────────────────────
12-14│ 第四阶段          │ 性能优化完成
     │ 优化和上线        │ 灰度发布成功
     │                  │ 正式上线
```

### 5.6 风险控制

**技术风险：**
- **风险**：路由算法复杂度高，可能影响性能
- **应对**：充分的性能测试，多级缓存，异步处理

**业务风险：**
- **风险**：定价调整可能引起客户不满
- **应对**：提前沟通，价格锁定，平滑过渡

**数据风险：**
- **风险**：数据迁移可能出错
- **应对**：完整备份，分批迁移，回滚方案

**上线风险：**
- **风险**：新功能可能影响现有业务
- **应对**：灰度发布，功能开关，快速回滚

## 四、技术实现要点

### 4.1 性能优化

1. **多级缓存策略**
```go
// 缓存层次
L1: 内存缓存（sync.Map） - 热点数据
L2: Redis缓存 - 定价策略、渠道质量
L3: 数据库 - 持久化存储

// 缓存更新策略
- 定价策略：TTL 5分钟
- 渠道质量：实时更新 + 5秒缓存
- 路由决策：请求级缓存
```

2. **数据库优化**
```sql
-- 关键索引
CREATE INDEX idx_channel_model_tier ON channel_quality_tier(channel_id, model_name, quality_tier);
CREATE INDEX idx_user_strategy ON customer_pricing_strategy(user_id, enabled);
CREATE INDEX idx_model_quality ON channel_quality_tier(model_name, quality_tier, enabled);

-- 分区表（大数据量时）
ALTER TABLE pricing_route_log PARTITION BY RANGE (YEAR(created_at)) (
    PARTITION p2024 VALUES LESS THAN (2025),
    PARTITION p2025 VALUES LESS THAN (2026)
);
```

3. **异步处理**
```go
// 非关键路径异步化
- 质量指标统计：异步队列
- 成本分析报表：定时任务
- 日志记录：批量写入
```

### 4.2 兼容性保障

1. **向后兼容**
```go
// 保持现有API不变
// 内部逻辑增强，对外透明

// 渐进式迁移
if config.EnableAdvancedPricing {
    // 使用新的定价路由
    decision := pricingRouter.SelectOptimalChannel(...)
} else {
    // 使用原有逻辑
    channel := channelGroup.Next(...)
}
```

2. **数据迁移**
```go
// 自动迁移脚本
func MigrateExistingPrices() {
    // 1. 读取现有价格表
    prices := GetAllPrices()
    
    // 2. 为每个渠道创建标准质量等级
    for _, channel := range channels {
        for _, price := range prices {
            CreateChannelQualityTier(&ChannelQualityTier{
                ChannelID: channel.Id,
                ModelName: price.Model,
                QualityTier: "standard",  // 默认标准等级
                InputPrice: price.Input,
                OutputPrice: price.Output,
            })
        }
    }
    
    // 3. 为现有用户创建默认定价策略
    for _, user := range users {
        CreateDefaultPricingStrategy(user.Id)
    }
}
```

### 4.3 监控和告警

```go
// 关键指标监控
type PricingMetrics struct {
    // 成本指标
    TotalCost float64
    AvgCostPerRequest float64
    CostByChannel map[int]float64
    
    // 收入指标
    TotalRevenue float64
    AvgRevenuePerRequest float64
    
    // 利润指标
    TotalProfit float64
    ProfitMargin float64
    
    // 质量指标
    AvgQualityScore float64
    QualityDistribution map[string]int
    
    // 路由指标
    RouteDecisionTime time.Duration
    ChannelUtilization map[int]float64
}

// 告警规则
- 利润率低于阈值（如10%）
- 某渠道成本异常上涨
- 质量等级降级频繁
- 路由决策耗时过长
```

## 五、预期收益

### 5.1 成本优化

- **智能路由**：根据实时成本选择最优渠道，预计降低成本 15-25%
- **批量采购**：基于质量等级的批量采购谈判，预计降低成本 10-15%
- **负载优化**：避免高成本渠道过载，预计降低成本 5-10%

### 5.2 收入增长

- **差异化定价**：针对不同客户提供定制化价格，预计提升收入 20-30%
- **质量分级**：高质量服务溢价，预计提升高端客户收入 30-50%
- **灵活定价**：支持更多定价模式，预计扩大客户群 15-20%

### 5.3 运营效率

- **自动化定价**：减少人工定价工作量 80%
- **实时分析**：快速响应市场变化，决策周期缩短 70%
- **客户满意度**：质量保障和透明定价，预计提升满意度 25%

## 六、风险和应对

### 6.1 技术风险

**风险：** 系统复杂度增加，可能影响性能
**应对：**
- 充分的性能测试和压力测试
- 多级缓存和异步处理
- 灰度发布，逐步切换

### 6.2 业务风险

**风险：** 定价策略调整可能引起客户不满
**应对：**
- 提前沟通，给予过渡期
- 保持向后兼容，不强制升级
- 提供价格锁定选项

### 6.3 数据风险

**风险：** 数据迁移可能出现错误
**应对：**
- 完整的数据备份
- 分批迁移，逐步验证
- 回滚方案准备

## 六、关键设计原则

### 6.1 概念隔离原则

**上游概念（内部管理）：**
- 折扣组：号池、逆向、官转、优质官转
- 真实成本价格
- 供应商信息
- 风险等级

**下游概念（客户视角）：**
- 服务等级：经济版、标准版、专业版、企业版
- 客户价格
- SLA承诺
- 服务描述

**隔离机制：**
- 数据库层面：不同的表结构
- API层面：客户API不暴露折扣组信息
- 界面层面：客户只看到服务等级

### 6.2 灵活映射原则

**服务等级到折扣组的映射是可配置的：**
- 一个服务等级可以映射多个折扣组
- 支持主要折扣组和降级备选
- 支持排除特定折扣组
- 映射关系可以随时调整

**示例：**
```
标准版 → 主要使用"官转"，降级备选"逆向"，排除"号池"
专业版 → 仅使用"优质官转"，不降级
经济版 → 主要使用"号池"和"逆向"，降级备选"官转"
```

### 6.3 成本收益平衡原则

**路由决策考虑三个维度：**
1. **成本**：上游供应商的真实成本
2. **质量**：成功率、响应时间、稳定性
3. **利润**：客户价格 - 成本

**平衡策略：**
- 不是单纯追求成本最低
- 不是单纯追求利润最高
- 而是在保证质量的前提下，优化成本和利润

### 6.4 向后兼容原则

**保持现有功能不变：**
- 现有API接口不变
- 现有数据结构保留
- 现有业务逻辑兼容

**渐进式升级：**
- 通过功能开关控制新旧逻辑
- 支持灰度发布
- 支持快速回滚

### 6.5 数据驱动原则

**所有决策基于数据：**
- 实时监控质量指标
- 自动计算成本收益
- 数据分析支持决策

**持续优化：**
- 根据历史数据优化路由策略
- 根据质量数据调整渠道权重
- 根据成本数据调整定价策略

## 七、总结

### 7.1 核心创新点

本方案的核心创新在于**双层抽象机制**：

1. **上游层（折扣组）**
   - 真实反映供应商的渠道类型和成本结构
   - 内部管理使用，精细化成本控制
   - 支持灵活的供应商管理

2. **下游层（服务等级）**
   - 客户友好的服务分级
   - 隐藏技术细节，提升客户体验
   - 支持灵活的定价策略

3. **智能映射层**
   - 将上游折扣组映射到下游服务等级
   - 根据成本、质量、利润智能路由
   - 自动优化，持续改进

### 7.2 核心价值

**对平台方：**
- 成本降低15-25%（智能路由优化）
- 收入提升20-30%（差异化定价）
- 利润率提升5-10个百分点
- 运营效率提升80%（自动化管理）

**对客户：**
- 服务质量透明化（清晰的服务等级）
- 灵活的服务选择（根据场景选择）
- 价格可预期（明确的定价标准）
- 服务质量保障（SLA承诺）

**竞争优势：**
- 相比传统聚合平台：多服务等级，客户可选择
- 相比直接对接供应商：统一接口，透明定价
- 相比其他聚合平台：智能路由，成本质量利润综合优化

### 7.3 实施建议

**分阶段实施：**
1. 先完成数据模型和基础功能（2-3周）
2. 再实现核心路由引擎（3-4周）
3. 然后开发管理后台（3-4周）
4. 最后优化和上线（2-3周）

**风险控制：**
- 充分的测试和验证
- 灰度发布，逐步放量
- 保持向后兼容，支持快速回滚

**持续优化：**
- 根据实际运营数据持续优化算法
- 根据客户反馈持续改进体验
- 根据市场变化持续调整策略

### 7.4 预期效果

**短期效果（3个月内）：**
- 成本管理精细化，成本可控
- 定价策略灵活化，支持差异化定价
- 管理效率提升，减少人工工作量

**中期效果（6-12个月）：**
- 成本降低15-25%
- 收入提升20-30%
- 利润率提升5-10个百分点

**长期效果（1年以上）：**
- 建立竞争壁垒，形成核心竞争力
- 客户满意度提升，客户粘性增强
- 数据积累，支持更多智能化功能

---

## 附录

### 附录A：术语表

| 术语 | 英文 | 定义 |
|-----|------|-----|
| 折扣组 | Discount Group | 上游供应商提供的真实渠道类型，如号池、逆向、官转、优质官转 |
| 服务等级 | Service Tier | 面向客户的服务质量分级，如经济版、标准版、专业版、企业版 |
| 号池 | Pool | 共享账号池，成本极低但稳定性差 |
| 逆向 | Reverse | 逆向工程接口，成本低但有合规风险 |
| 官转 | Official | 官方转售渠道，稳定可靠 |
| 优质官转 | Premium Official | 官方优质渠道，最高稳定性和SLA保障 |
| 路由决策 | Route Decision | 根据成本、质量、利润选择最优渠道的过程 |
| 定价策略 | Pricing Strategy | 针对客户的定价规则和加价方式 |

### 附录B：配置示例

**服务等级配置示例：**

```json
{
  "tier_code": "standard",
  "tier_name": "标准版",
  "description": "稳定可靠，适合生产环境",
  "mapped_discount_groups": {
    "primary": ["official"],
    "fallback": ["reverse"],
    "excluded": ["pool"],
    "selection_strategy": "balanced"
  }
}
```

**客户定价策略示例：**

```json
{
  "user_id": 123,
  "strategy_name": "VIP客户定价",
  "pricing_config": {
    "type": "tier_based",
    "rules": [
      {
        "model_pattern": "gpt-4*",
        "service_tier": "professional",
        "pricing": {
          "input": 18.0,
          "output": 36.0,
          "markup_type": "fixed"
        }
      }
    ],
    "default_markup": {
      "type": "percentage",
      "value": 30
    }
  }
}
```

---

**文档版本：** v2.0  
**创建日期：** 2025-01-XX  
**最后更新：** 2025-01-XX  
**作者：** Kiro AI Assistant  
**审核状态：** 待审核

**变更记录：**
- v2.0: 澄清折扣组概念，完善双层抽象机制
- v1.0: 初始版本
