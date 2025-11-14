# 客户级模型分组计费设计方案（基于 dev 分支）

> 状态：进行中  
> 目标：在保持现有计费架构稳定的前提下，引入「客户 × 模型 × 分组」的客户价能力，并支持三种计费模式（按 tokens、按次、按秒）。

---

## 1. 背景与问题

当前 One Hub 的计费模型核心特点：

- **统一积分制 `quota`**：所有调用最终折算为整数积分，记录在用户、Token、日志与账单中。
- **全局模型官方价**：`model/price.go` 中的 `Price` 表示每个模型的官方基准价（Rate(K)），通过 `PricingInstance.GetPrice` 使用。
- **用户组倍率**：`UserGroup.ratio` 提供了一层按用户组的统一折扣/加价。
- **按秒计费的雏形**：Sora/Veo 等视频模型已通过“秒 × 1000 token”在 `Quota` 层做折算，但没有一等的 `seconds` 计费类型。

现状不足：

- 无法为**单个客户（User）**配置独立价格。
- 无法对同一模型的不同质量档（分组）单独定价。
- 按秒计费没有变成一等公民，在价表与展示层难以表达“￥/秒”。
- 渠道路由与“产品档位（分组）”之间没有显式绑定关系。

本设计在不破坏上述核心机制的前提下，引入客户级价格与模型分组，并为未来的多租户与精细化路由做好铺垫。

---

## 2. 目标与边界

### 2.1 本次要实现的能力

1. **客户级价格**
   - 计费主体：当前阶段 **一个 `User` = 一个客户**。
   - 支持为「客户 × 模型 × 分组」配置客户价。

2. **模型分组（质量档）**
   - 每个模型有一个**默认分组**，有一份**官方基准价**（现有 `Price` 表）。
   - 同一个模型可以定义多个**非默认分组**（high_quality / enterprise 等）。
   - 分组与渠道绑定：一个分组可以绑定多个 Channel，用于后续路由。

3. **客户价覆盖规则**
   - 对**默认分组**可以配置客户价 → 覆盖官方基准价。
   - 对**非默认分组**：
     - 只有在「客户价 + 授权」都存在时，该分组才对该客户可用；
     - 否则视为该分组对该客户不存在。

4. **三种计费模式**
   - `tokens`：按输入/输出 tokens 计费（现有）。
   - `times`：按次计费（现有）。
   - `seconds`：用于少数视频模型，对外呈现“￥/秒”，内部仍折算为 tokens 等价后计费。

### 2.2 本次不做（只做兼容设计）

- 不引入独立 `Customer/Tenant` 实体，多账号共享额度属于后续版本。
- 不全面重构渠道路由，只预留“分组-渠道绑定”结构，允许后续逐步接入。
- 不修改日志与账单表结构，保证新的价格逻辑仍写入统一的 `quota` 与 tokens 统计。

---

## 3. 统一业务模型说明

### 3.1 模型与分组

- **默认分组**
  - 每个模型有一个默认分组（例如 `default` / `standard`），用于兼容现有逻辑。
  - 默认分组拥有一份全局**官方价**：
    - 由现有 `Price` 表与 `GetDefaultPrice()` 提供。
    - 在未设置客户价时，所有客户对默认分组都使用官方价。

- **非默认分组**
  - 用于表达不同质量档 / SLA / 特殊产品位（如 `hq` / `enterprise`）。
  - 非默认分组本身**没有全局价**；只有在为某客户配置客户价后才有效。

### 3.2 官方价与客户价关系

- 官方价（全局）：
  - 定义在「模型 × 默认分组」上。
  - 始终存在，用于兜底。

- 客户价（按客户）：
  - 定义在「客户(User) × 模型 × 分组」维度。
  - 对**默认分组**：
    - 若存在客户价 → 覆盖官方价。
    - 若不存在 → 使用官方价。
  - 对**非默认分组**：
    - 必须存在客户价 + 授权，才视为该分组可用；
    - 没客户价或没授权 → 该分组对该客户不可用。

### 3.3 授权规则

- 默认分组：
  - 所有客户默认有权限。
  - 是否有“默认分组的客户价”只影响价格，不影响能否使用。

- 非默认分组：
  - 通过「客户 × 模型 × 分组」的授权记录控制可用性。
  - 只有授权 + 客户价同时存在时，该分组才对该客户生效。

### 3.4 价格解析统一优先级

给定一次调用（`user_id`, `model_name`, `group_code?`）：

1. **确定分组**
   - 若 `group_code` 为空：
     - 从「模型分组」表中获取默认分组 `defaultGroupCode`；
     - 若找不到（极端情况），退回无分组模式 → 直接用 `PricingInstance.GetPrice(model_name)`。
   - 若 `group_code` 非空：
     - 校验 `(model_name, group_code)` 是否存在分组定义；
     - 若不存在 → 拒绝（分组不存在）。

2. **授权检查**
   - 若分组是默认分组：
     - 默认认为所有客户有权限（MVP 阶段）。
   - 若分组为非默认：
     - 检查「客户-模型-分组授权」记录是否存在且启用；
     - 若未授权 → 拒绝。

3. **价格选择**
   - 首先查「客户价」：`(user_id, model_name, group_code)`：
     - 若存在且启用 → 使用客户价。
   - 否则：
     - 若分组是默认分组 → 使用官方价（`PricingInstance.GetPrice(model_name)`）。
     - 若分组是非默认分组 → 视为该分组未配置价格 → 拒绝。

4. **计费方式**
   - 每条价记录有 `Type`：
     - `tokens`：按输入/输出 token 数 × Rate 计费。
     - `times`：按次固定扣费。
     - `seconds`：
       - 对外：客户价配置/展示为「每秒多少钱」。
       - 内部：折算为虚拟 tokens（`seconds * 1000`），再按 token 类似逻辑计费。

---

## 4. 数据结构设计（后端）

### 4.1 模型分组定义：`ModelGroup`

> 文件建议：`model/model_group.go`

**字段（Go 结构拟稿）**

- `ID`：int
- `Model`：string —— 模型名（与 `Price.Model` 一致）
- `GroupCode`：string —— 分组编码，例如：`default` / `standard` / `hq` / `enterprise`
- `DisplayName`：string —— 展示名称
- `Description`：string —— 说明
- `IsDefault`：bool —— 是否为该模型默认分组（每模型唯一）
- `BillingType`：string —— `"tokens" | "times" | "seconds"`
- `CreatedAt`, `UpdatedAt`：int64 / `time.Time`

**约束**

- `(Model, GroupCode)` 唯一。
- 每个 `Model` 至少有一个 `IsDefault=true` 的分组（通过初始化或数据迁移保证）。

### 4.2 分组-渠道绑定：`ModelGroupChannel`（预留）

> 文件建议：`model/model_group_channel.go`

MVP 只定义结构和基础 CRUD，路由逻辑可后续逐步接入。

**字段**

- `ID`
- `Model`：string
- `GroupCode`：string
- `ChannelID`：int
- `Weight`：int —— 权重
- `Priority`：int —— 优先级（主/备）
- `Enabled`：bool
- `CreatedAt`, `UpdatedAt`

### 4.3 客户价：`UserModelGroupPrice`

> 文件建议：`model/user_model_group_price.go`

**字段**

- `ID`
- `UserID`：int —— 客户（即 `User.Id`）
- `Model`：string
- `GroupCode`：string
- `Type`：string —— `"tokens" | "times" | "seconds"`
- `InputRate`：float64 —— Rate(K)，用于输入计费
- `OutputRate`：float64 —— Rate(K)，用于输出计费
- `ExtraRatios`：`datatypes.JSONType[map[string]float64]` —— 可选，兼容现有 Extra 用量（缓存、音频、推理等）
- `Enabled`：bool
- `CreatedAt`, `UpdatedAt`

**语义**

- 默认分组：
  - 若存在对应记录 → 该客户在默认分组使用客户价。
  - 若不存在 → 使用官方价（`Price`）。
- 非默认分组：
  - 若存在客户价 + 授权 → 分组对该客户可用。
  - 若缺失客户价或授权 → 分组对该客户不可用。

### 4.4 分组授权：`UserModelGroupPermission`

> 文件建议：`model/user_model_group_permission.go`

**字段**

- `ID`
- `UserID`：int
- `Model`：string
- `GroupCode`：string
- `Enabled`：bool
- `CreatedAt`, `UpdatedAt`

**规则**

- 默认分组：
  - MVP 阶段可以不写入授权记录，视为所有客户默认有权限。
- 非默认分组：
  - 必须存在 `Enabled=true` 的授权 + 客户价记录。

### 4.5 AutoMigrate 接入

> 文件：`model/main.go`

在 `InitDB` 中新增：

- `db.AutoMigrate(&ModelGroup{}, &ModelGroupChannel{}, &UserModelGroupPrice{}, &UserModelGroupPermission{})`

并在 `migrationAfter` 或初始化逻辑中，为已有模型补充默认分组（以 `Price` 中存在的模型为基准，创建 `GroupCode="default"` 的分组）。

---

## 5. 价格解析逻辑设计

### 5.1 函数入口

> 文件建议：`model/pricing_customer.go`

```go
// ResolveCustomerPrice 解析客户在指定模型/分组上的有效价格
// groupCode 为空时使用默认分组。
func ResolveCustomerPrice(userID int, modelName, groupCode string) (*Price, string, error)
```

返回值：

- `*Price`：最终用于计费的价格对象（可直接复用现有 `Price` 类型）。
- `string`：实际使用的分组编码（包括默认分组编码，用于日志调试）。
- `error`：任何授权/配置问题均通过 error 表达，并映射为 OpenAI 兼容错误响应。

### 5.2 逻辑步骤

1. **确定分组**
   - 若 `groupCode == ""`：
     - 从 `ModelGroup` 表中查询 `IsDefault=true` 的记录；
     - 若不存在，退回到旧逻辑：直接 `PricingInstance.GetPrice(modelName)`，分组记为空。
   - 若 `groupCode != ""`：
     - 校验 `(modelName, groupCode)` 是否存在 `ModelGroup` 记录；
     - 不存在则返回错误（分组不存在）。

2. **授权检查**
   - 若分组为默认分组：
     - MVP 阶段直接视为授权通过（后续可引入显式控制）。
   - 若为非默认分组：
     - 查询 `UserModelGroupPermission`，要求 `Enabled=true`；
     - 否则返回错误（未授权分组）。

3. **价格选择**
   - 查询 `UserModelGroupPrice(userID, modelName, groupCode)`：
     - 若存在且 `Enabled=true`：
       - 将其转换为 `Price` 结构（填充 `Model/Type/Input/Output/ExtraRatios` 等）。
       - 返回该价格。
   - 若不存在客户价：
     - 若分组为默认：
       - 调用 `PricingInstance.GetPrice(modelName)` 获取官方价，返回。
     - 若为非默认：
       - 返回错误（该分组未配置价格）。

4. **计费类型映射**
   - 对于 `Type="tokens"` / `"times"`：
     - 直接等价于现有 `Price.Type`。
   - 对于 `Type="seconds"`：
     - 仍使用 `InputRate/OutputRate` 表示内部 Rate(K)，但在 `Quota` 中：
       - 通过统一的“秒 → 虚拟 tokens”函数，将实际秒数换算为 token 数；
       - 再走 tokens 模式的计费逻辑。

---

## 6. 计费引擎接入方案（Quota）

### 6.1 分组选择（通过令牌而非 Header）

当前设计不再使用自定义 Header（例如 `X-Model-Group`）来选择分组，以避免对 OpenAI 兼容协议造成侵入。

**分组选择原则：**

- 每次调用由一个 API 令牌（Token）发起；
- 不同价档（分组）通过为客户生成不同的「分组令牌」体现：
  - 使用默认分组价时，使用默认分组令牌；
  - 使用某个高质量/特殊分组价时，为该分组生成单独的分组令牌。
- 即：**调用使用哪个令牌，就隐含选择了哪个模型分组**。

在现有实现中：

- `ResolveCustomerPrice(userID, modelName, groupCode)` 的 `groupCode` 目前为空时，会自动回退到默认分组；
- 后续可以在 Token 结构（例如 `Token.Setting`）中增加「定价分组」字段，并在鉴权阶段根据 Token 写入 `gin.Context`：
  - 例如：`c.Set("model_group", pricingGroupCode)`；
- 这样即可实现：通过不同令牌区分不同分组，而不需要在调用层增加额外 Header。

### 6.2 Quota 中替换价格获取逻辑

> 文件：`relay/relay_util/quota.go`

当前逻辑：

```go
quota.price = *model.PricingInstance.GetPrice(quota.modelName)
```

改为：

```go
groupCode := c.GetString("model_group")
userID := c.GetInt("id")
price, resolvedGroup, err := model.ResolveCustomerPrice(userID, quota.modelName, groupCode)
if err != nil {
    // 转换为 OpenAI 兼容错误返回（带上错误码与提示）
}
quota.price = *price
// 可选：将 resolvedGroup 写入日志或 context 便于调试
```

其余逻辑（`group_ratio`、预扣、最终额度计算）保持不变。

### 6.3 秒计费模型处理

对 `Type="seconds"` 的价格：

- 在 `Quota` 中统一调用辅助函数，例如：

```go
func secondsToVirtualTokens(modelName string, seconds int) int
```

- 针对视频模型（如 Sora/Veo）的现有特殊处理，迁移为使用上述函数，以便统一维护。
- 逻辑仍是：
  - `virtual_tokens = seconds * 1000`（或模型特定的换算规则）
  - 再按 token 模型计算 `quota`。

---

## 7. 管理端前端与 API 设计（MVP）

> 本节为后续实施参考，本次后端实现以打通核心计费链路为主。

### 7.1 管理端页面

1. **模型分组管理**
   - 入口：管理员后台 → 模型价格 / ModelPrice → 分组管理。
   - 能力：
     - 为某模型增加/编辑分组（编码、展示名、计费类型、描述）。
     - 查看分组列表，标记默认分组。

2. **客户价配置**
   - 入口：管理员后台 → 用户详情 → 客户价（Pricing）Tab。
   - 能力：
     - 选择模型；
     - 查看该用户在该模型的所有分组：
       - 默认分组：显示官方价（只读） + 可编辑客户价（tokens/times/seconds）。
       - 非默认分组：可配置授权 + 客户价。

3. **（后续）分组-渠道绑定**
   - 入口：模型分组详情页中的“渠道绑定”区域。
   - 能力：
     - 为分组选择一个或多个 Channel，设置权重与启停。

### 7.2 API 粗略规划

以 `/api/admin/...` 为前缀，建议至少：

- 模型分组：
  - `GET /api/admin/model_groups?model=xxx`
  - `POST /api/admin/model_group`（增/改）
- 客户价与授权（可设计为聚合接口）：
  - `GET /api/admin/users/{id}/model_pricing?model=xxx`
  - `POST /api/admin/users/{id}/model_pricing`（批量提交某用户对某模型的分组价与授权）

---

## 8. 实施任务拆分与验收标准

### 阶段 0：准备 & 校验

**任务：**

- [ ] 确认 dev 分支为当前开发基线，创建 `feature/customer-pricing` 分支。
- [ ] 梳理首批需要配置客户价的模型与分组方案。

**验收标准：**

- 有一份明确的清单：`{model_name, 默认分组编码, 其它分组编码, 计费类型}`。

---

### 阶段 1：数据结构与 AutoMigrate

**任务：**

- [x] 新增 `ModelGroup`、`UserModelGroupPrice`、`UserModelGroupPermission` Go 结构与表定义（`ModelGroupChannel` 留待后续路由优化时再引入）。
- [x] 在 `model/main.go` 的 `InitDB` 中接入上述结构的 `AutoMigrate`。
- [x] 编写初始化逻辑：为现有的所有 `Price.Model` 创建默认分组（若不存在），见 `EnsureDefaultModelGroupForPrice`。

**验收标准：**

- dev 环境启动成功且自动迁移通过。
- 新表可写入/读取，现有计费逻辑在未启用客户价情况下无变化。

---

### 阶段 2：价格解析与授权接入

**任务：**

- [x] 在 `model/pricing_customer.go` 实现 `ResolveCustomerPrice`，按「客户价 → 默认分组官方价兜底」优先级解析。
- [x] 在 `relay/relay_util/quota.go` 中改用 `ResolveCustomerPrice` 获取价格，失败时回退旧逻辑并记录日志，保证兼容性。
- [x] 在令牌鉴权阶段（`middleware/auth.go/tokenAuth`）读取 `TokenSetting.PricingGroup`，将分组信息写入 `gin.Context` 的 `model_group`，实现“分组令牌”选择分组（不再使用自定义 Header）。

**验收标准：**

- 未配置任何客户价、不传 `X-Model-Group` 时，所有模型扣费结果与改造前完全一致。
- 为某用户 + 某模型默认分组配置客户价后，不传分组时扣费结果按客户价计算。
- 为某用户 + 某模型 + 非默认分组配置客户价并授权后，传入该分组时扣费按客户价计算，未授权时返回明确错误。

---

### 阶段 3：秒计费模型落地（视频）

**任务：**

- [ ] 确认首批按秒计费模型列表与秒数来源字段。
- [ ] 在相关 `ModelGroup` 分组上设置 `BillingType="seconds"`。
- [ ] 在客户价配置中支持按“每秒价格”录入，并在后端映射为内部 Rate(K)。

**验收标准：**

- 对指定视频模型，给定客户价为 `X 元/秒`，调用时长为 `T 秒`，日志中的 `quota` 与等价金额在允许误差内匹配 `X*T`。

---

### 阶段 4：管理端配置与后续优化（规划）

> 本阶段可视资源分配安排在后续迭代，本设计文档做为实现参考。

**任务：**

- [ ] 管理端实现模型分组、客户价与授权的配置页面。
- [x] 后端实现对应的管理 API，支持按用户/模型批量配置（`/api/customer_pricing/model_groups` 与 `/api/customer_pricing/users/:id/model_pricing`）。

**验收标准：**

- 管理员可以仅通过 UI 完成：为某客户配置默认分组客户价、开通非默认分组并配置专属价。

---

## 9. 本次迭代实施范围说明

本次在 `dev` 分支上新开的功能迭代，优先实现：

1. 新增后端数据结构与自动迁移（`ModelGroup`、`UserModelGroupPrice`、`UserModelGroupPermission`）——已完成。
2. 实现 `ResolveCustomerPrice` 并在 Quota 计费流程中接入，打通「客户 × 模型 × 分组」价格链路，并通过 `TokenSetting.PricingGroup` 支持分组令牌 —— 已完成。
3. 为秒计费预留 `seconds` 类型及内部折算机制（类型常量已预留，统一换算函数与首批视频模型接入待后续迭代）。

管理端 UI 与完整的渠道绑定与路由优化，将在后续迭代中按本设计文档继续推进。

---

## 10. 当前进度摘要（2025-11）

- 已完成：
  - 后端数据结构与自动迁移（模型分组、客户价、授权）。
  - 默认分组初始化逻辑（基于现有 `Price` 自动补齐 `default` 分组）。
  - 客户价解析函数 `ResolveCustomerPrice` 与 Quota 接入。
  - 分组令牌机制：通过 `TokenSetting.PricingGroup` 选择计费分组。
  - 客户价与分组管理 API 骨架：`/api/customer_pricing/model_groups` 与 `/api/customer_pricing/users/:id/model_pricing`。
- 规划中 / 未完成：
  - 按秒计费模型的统一换算函数与首批视频模型接入。
  - 管理端前端页面（分组管理、客户价配置）。
  - 分组-渠道绑定（`ModelGroupChannel`）及其在路由层的实际使用。
