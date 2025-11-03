# 价格单位自动转换功能

## 功能概述

修复了后台运营模型价格编辑功能中，切换USD/RMB和K/M单位时价格不会自动调整的问题。

## 主要改进

### 1. 新增价格转换工具模块 (`priceConverter.js`)

- **双向转换支持**：支持 Rate ↔ USD ↔ RMB 之间的相互转换
- **单位转换**：支持 K（千）↔ M（百万）之间的转换
- **高精度计算**：使用 Decimal.js 确保计算精度
- **错误处理**：完善的边界值和异常情况处理

### 2. 增强的用户界面 (`EditModal.jsx`)

- **实时转换**：单位切换时自动转换输入框中的价格
- **视觉反馈**：转换过程中的加载状态和提示信息
- **仅单一模式**：转换功能仅在单一模式下生效，避免多选模式的复杂性

## 转换规则

### 汇率常量
```javascript
USD: 0.002  // Rate → USD
RMB: 0.014  // Rate → RMB
```

### 数量单位
```javascript
K: 1        // 千（基准）
M: 1000     // 百万 = 1000 * 千
```

### 转换示例

| 原始值 | 原始单位 | 目标单位 | 转换后值 | 说明 |
|--------|----------|----------|----------|------|
| 1000 | Rate(K) | USD(K) | 2 | 1000 × 0.002 |
| 1000 | Rate(K) | RMB(K) | 14 | 1000 × 0.014 |
| 2 | USD(K) | Rate(K) | 1000 | 2 ÷ 0.002 |
| 1000 | Rate(K) | Rate(M) | 1 | 1000 ÷ 1000 |
| 1 | USD(M) | RMB(K) | 7000 | (1×1000)÷0.002×0.014 |

## 使用方法

### 在价格编辑界面

1. 打开单个模型的价格编辑对话框
2. 输入价格值（如：1000）
3. 切换货币单位（Rate → USD）
4. 观察输入框中的价格自动调整为对应值（如：2）
5. 切换数量单位（K → M）
6. 观察价格再次自动调整

### 编程接口

```javascript
import { convertPrice, convertPriceData } from './priceConverter';

// 单个价格转换
const newPrice = convertPrice(1000, 'rate', 'K', 'USD', 'K');
// 结果: 2

// 批量转换价格对象
const priceData = { input: 1000, output: 2000 };
const converted = convertPriceData(priceData, 'rate', 'K', 'USD', 'K');
// 结果: { input: 2, output: 4 }
```

## 技术特性

### 可逆性保证
所有转换都是可逆的，确保 A→B→A 的转换结果一致：
```javascript
const original = 1000;
const toUSD = convertPrice(original, 'rate', 'K', 'USD', 'K');
const backToRate = convertPrice(toUSD, 'USD', 'K', 'rate', 'K');
// original === backToRate (在精度范围内)
```

### 精度控制
- 使用 Decimal.js 避免浮点数精度问题
- 结果保留4位小数
- 支持极小值和极大值的转换

### 错误处理
- 空值、null、undefined 自动转换为 0
- 非数字值返回 0
- 无效单位类型返回原值
- 异常情况下的降级处理

## 测试覆盖

包含全面的单元测试：
- 基础转换测试
- 单位转换测试
- 边界值测试
- 可逆性测试
- 错误处理测试
- 性能测试

## 兼容性

- 与现有后端 API 完全兼容
- 保持原有的提交逻辑不变
- 仅在前端界面增加实时转换功能
- 不影响多选模式的现有功能

## 注意事项

1. **仅单一模式生效**：转换功能只在编辑单个模型价格时生效
2. **汇率同步**：汇率常量需要与后端 `model/price.go` 保持一致
3. **精度限制**：转换结果保留4位小数，可能存在微小的精度损失
4. **性能考虑**：转换过程使用了100ms的延迟来提供更好的用户体验

## 维护指南

### 更新汇率
如需更新汇率，请同时修改：
1. `web/src/views/Pricing/component/priceConverter.js` 中的 `EXCHANGE_RATES`
2. `model/price.go` 中的 `DollarRate` 和 `RMBRate`

### 添加新货币
1. 在 `EXCHANGE_RATES` 中添加新汇率
2. 在 `isValidUnitType` 函数中添加新货币类型
3. 更新相关的转换逻辑和测试用例
