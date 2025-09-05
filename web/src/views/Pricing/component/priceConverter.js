import Decimal from 'decimal.js';

/**
 * 价格单位转换工具模块
 *
 * 支持以下转换：
 * - 货币单位：Rate ↔ USD ↔ RMB
 * - 数量单位：K ↔ M (千 ↔ 百万)
 *
 * 转换基准：所有转换都通过 Rate(K) 作为中间单位
 *
 * @author AI Assistant
 * @version 1.0.0
 */

// 汇率常量 - 与后端 model/price.go 保持一致
export const EXCHANGE_RATES = {
  USD: 0.002,  // Rate → USD 转换率
  RMB: 0.014,  // Rate → RMB 转换率
};

// 数量单位转换常量
export const UNIT_MULTIPLIERS = {
  K: 1,      // 千（基准单位）
  M: 1000,   // 百万 = 1000 * 千
};

/**
 * 将价格从源单位转换为 Rate(K) 基准单位
 * @param {number|string} price - 价格值
 * @param {string} currencyType - 货币类型: 'rate', 'USD', 'RMB'
 * @param {string} unit - 数量单位: 'K', 'M'
 * @returns {Decimal} Rate(K) 单位的价格
 */
function convertToRateK(price, currencyType, unit) {
  if (price === '' || price == null) {
    return new Decimal(0);
  }

  let priceValue = new Decimal(price);

  // 处理数量单位转换 (M → K):
  // 注意：在倍率(rate)模式下，K 与 M 不改变数值，故不做缩放
  if (currencyType !== 'rate' && unit === 'M') {
    priceValue = priceValue.div(UNIT_MULTIPLIERS.M);
  }

  // 处理货币单位转换
  switch (currencyType) {
    case 'rate':
      // Rate 单位不需要转换
      break;
    case 'USD':
      // USD → Rate: 除以汇率
      priceValue = priceValue.div(EXCHANGE_RATES.USD);
      break;
    case 'RMB':
      // RMB → Rate: 除以汇率
      priceValue = priceValue.div(EXCHANGE_RATES.RMB);
      break;
    default:
      throw new Error(`Unsupported currency type: ${currencyType}`);
  }

  return priceValue;
}

/**
 * 将 Rate(K) 基准单位转换为目标单位
 * @param {Decimal} rateKPrice - Rate(K) 单位的价格
 * @param {string} targetCurrencyType - 目标货币类型: 'rate', 'USD', 'RMB'
 * @param {string} targetUnit - 目标数量单位: 'K', 'M'
 * @returns {Decimal} 目标单位的价格
 */
function convertFromRateK(rateKPrice, targetCurrencyType, targetUnit) {
  let result = new Decimal(rateKPrice);

  // 处理货币单位转换
  switch (targetCurrencyType) {
    case 'rate':
      // Rate 单位不需要转换
      break;
    case 'USD':
      // Rate → USD: 乘以汇率
      result = result.mul(EXCHANGE_RATES.USD);
      break;
    case 'RMB':
      // Rate → RMB: 乘以汇率
      result = result.mul(EXCHANGE_RATES.RMB);
      break;
    default:
      throw new Error(`Unsupported target currency type: ${targetCurrencyType}`);
  }

  // 处理数量单位转换 (K → M):
  // 注意：在倍率(rate)模式下，K 与 M 不改变数值，故不做缩放
  if (targetCurrencyType !== 'rate' && targetUnit === 'M') {
    result = result.mul(UNIT_MULTIPLIERS.M);
  }

  return result;
}

/**
 * 价格单位转换主函数
 * @param {number|string} price - 原始价格
 * @param {string} fromCurrencyType - 源货币类型: 'rate', 'USD', 'RMB'
 * @param {string} fromUnit - 源数量单位: 'K', 'M'
 * @param {string} toCurrencyType - 目标货币类型: 'rate', 'USD', 'RMB'
 * @param {string} toUnit - 目标数量单位: 'K', 'M'
 * @returns {number} 转换后的价格
 */
export function convertPrice(price, fromCurrencyType, fromUnit, toCurrencyType, toUnit) {
  try {
    // 边界情况处理
    if (price === '' || price == null || isNaN(price)) {
      return 0;
    }

    // 如果源单位和目标单位相同，直接返回
    if (fromCurrencyType === toCurrencyType && fromUnit === toUnit) {
      return Number(price);
    }

    // 两步转换：源单位 → Rate(K) → 目标单位
    const rateKPrice = convertToRateK(price, fromCurrencyType, fromUnit);
    const targetPrice = convertFromRateK(rateKPrice, toCurrencyType, toUnit);

    // 返回保留6位小数的数值（避免在 K↔M 切换时出现小值被四舍五入为0的问题）
    return Number(targetPrice.toFixed(6));
  } catch (error) {
    console.error('Price conversion error:', error);
    return Number(price) || 0;
  }
}

/**
 * 批量转换价格对象中的 input 和 output 字段
 * @param {Object} priceData - 包含 input 和 output 的价格对象
 * @param {string} fromCurrencyType - 源货币类型
 * @param {string} fromUnit - 源数量单位
 * @param {string} toCurrencyType - 目标货币类型
 * @param {string} toUnit - 目标数量单位
 * @returns {Object} 转换后的价格对象
 */
export function convertPriceData(priceData, fromCurrencyType, fromUnit, toCurrencyType, toUnit) {
  console.log('convertPriceData called with:', {
    priceData,
    from: `${fromCurrencyType}(${fromUnit})`,
    to: `${toCurrencyType}(${toUnit})`
  });

  // 验证输入参数
  if (!priceData || typeof priceData !== 'object') {
    console.error('Invalid priceData:', priceData);
    return priceData;
  }

  // 如果单位相同，直接返回原数据
  if (fromCurrencyType === toCurrencyType && fromUnit === toUnit) {
    console.log('Units are the same, returning original data');
    return priceData;
  }

  const convertedInput = convertPrice(priceData.input, fromCurrencyType, fromUnit, toCurrencyType, toUnit);
  const convertedOutput = convertPrice(priceData.output, fromCurrencyType, fromUnit, toCurrencyType, toUnit);

  const result = {
    ...priceData,
    input: convertedInput,
    output: convertedOutput,
  };

  console.log('Conversion result:', {
    original: { input: priceData.input, output: priceData.output },
    converted: { input: convertedInput, output: convertedOutput }
  });

  return result;
}

/**
 * 验证单位类型是否有效
 * @param {string} currencyType - 货币类型
 * @param {string} unit - 数量单位
 * @returns {boolean} 是否有效
 */
export function isValidUnitType(currencyType, unit) {
  const validCurrencyTypes = ['rate', 'USD', 'RMB'];
  const validUnits = ['K', 'M'];
  
  return validCurrencyTypes.includes(currencyType) && validUnits.includes(unit);
}

/**
 * 获取单位显示标签
 * @param {string} currencyType - 货币类型
 * @param {string} unit - 数量单位
 * @returns {string} 显示标签
 */
export function getUnitLabel(currencyType, unit) {
  if (currencyType === 'rate') {
    return `Rate(${unit})`;
  }
  return `${currencyType}(${unit})`;
}
