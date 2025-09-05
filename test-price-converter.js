// 简单的价格转换测试脚本
// 模拟 Decimal.js 的基本功能
class Decimal {
  constructor(value) {
    this.value = parseFloat(value) || 0;
  }
  
  mul(other) {
    const otherValue = typeof other === 'object' ? other.value : parseFloat(other);
    return new Decimal(this.value * otherValue);
  }
  
  div(other) {
    const otherValue = typeof other === 'object' ? other.value : parseFloat(other);
    return new Decimal(this.value / otherValue);
  }
  
  toFixed(digits) {
    return this.value.toFixed(digits);
  }
  
  toString() {
    return this.value.toString();
  }
}

// 汇率常量
const EXCHANGE_RATES = {
  USD: 0.002,
  RMB: 0.014,
};

// 数量单位转换常量
const UNIT_MULTIPLIERS = {
  K: 1,
  M: 1000,
};

// 转换函数
function convertToRateK(price, currencyType, unit) {
  if (price === '' || price == null) {
    return new Decimal(0);
  }

  let priceValue = new Decimal(price);

  // 处理数量单位转换 (M → K)
  if (unit === 'M') {
    priceValue = priceValue.mul(UNIT_MULTIPLIERS.M);
  }

  // 处理货币单位转换
  switch (currencyType) {
    case 'rate':
      break;
    case 'USD':
      priceValue = priceValue.div(EXCHANGE_RATES.USD);
      break;
    case 'RMB':
      priceValue = priceValue.div(EXCHANGE_RATES.RMB);
      break;
    default:
      throw new Error(`Unsupported currency type: ${currencyType}`);
  }

  return priceValue;
}

function convertFromRateK(rateKPrice, targetCurrencyType, targetUnit) {
  let result = new Decimal(rateKPrice);

  // 处理货币单位转换
  switch (targetCurrencyType) {
    case 'rate':
      break;
    case 'USD':
      result = result.mul(EXCHANGE_RATES.USD);
      break;
    case 'RMB':
      result = result.mul(EXCHANGE_RATES.RMB);
      break;
    default:
      throw new Error(`Unsupported target currency type: ${targetCurrencyType}`);
  }

  // 处理数量单位转换 (K → M)
  if (targetUnit === 'M') {
    result = result.div(UNIT_MULTIPLIERS.M);
  }

  return result;
}

function convertPrice(price, fromCurrencyType, fromUnit, toCurrencyType, toUnit) {
  try {
    if (price === '' || price == null || isNaN(price)) {
      return 0;
    }

    if (fromCurrencyType === toCurrencyType && fromUnit === toUnit) {
      return Number(price);
    }

    const rateKPrice = convertToRateK(price, fromCurrencyType, fromUnit);
    const targetPrice = convertFromRateK(rateKPrice, toCurrencyType, toUnit);

    return Number(targetPrice.toFixed(4));
  } catch (error) {
    console.error('Price conversion error:', error);
    return Number(price) || 0;
  }
}

// 测试用例
console.log('=== 价格转换测试 ===');

// 测试 K 到 M 的转换
console.log('\n1. K 到 M 转换测试:');
console.log('1000 Rate(K) -> Rate(M):', convertPrice(1000, 'rate', 'K', 'rate', 'M')); // 应该是 1
console.log('500 Rate(K) -> Rate(M):', convertPrice(500, 'rate', 'K', 'rate', 'M'));   // 应该是 0.5

// 测试 M 到 K 的转换
console.log('\n2. M 到 K 转换测试:');
console.log('1 Rate(M) -> Rate(K):', convertPrice(1, 'rate', 'M', 'rate', 'K'));     // 应该是 1000
console.log('0.5 Rate(M) -> Rate(K):', convertPrice(0.5, 'rate', 'M', 'rate', 'K')); // 应该是 500

// 测试货币转换
console.log('\n3. 货币转换测试:');
console.log('1000 Rate(K) -> USD(K):', convertPrice(1000, 'rate', 'K', 'USD', 'K')); // 应该是 2
console.log('2 USD(K) -> Rate(K):', convertPrice(2, 'USD', 'K', 'rate', 'K'));       // 应该是 1000

// 测试复合转换
console.log('\n4. 复合转换测试:');
console.log('1000 Rate(K) -> USD(M):', convertPrice(1000, 'rate', 'K', 'USD', 'M')); // 应该是 0.002
console.log('0.002 USD(M) -> Rate(K):', convertPrice(0.002, 'USD', 'M', 'rate', 'K')); // 应该是 1000

// 测试可逆性
console.log('\n5. 可逆性测试:');
const original = 1000;
const converted = convertPrice(original, 'rate', 'K', 'rate', 'M');
const reverted = convertPrice(converted, 'rate', 'M', 'rate', 'K');
console.log(`原值: ${original}, 转换后: ${converted}, 逆转换: ${reverted}`);
console.log('可逆性检查:', Math.abs(original - reverted) < 0.0001 ? '通过' : '失败');

console.log('\n=== 测试完成 ===');
