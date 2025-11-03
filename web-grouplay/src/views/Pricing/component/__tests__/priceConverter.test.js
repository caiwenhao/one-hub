import { convertPrice, convertPriceData, isValidUnitType, getUnitLabel, EXCHANGE_RATES } from '../priceConverter';

describe('priceConverter', () => {
  describe('convertPrice', () => {
    // 基础转换测试
    describe('基础转换', () => {
      test('相同单位不应该改变价格', () => {
        expect(convertPrice(100, 'rate', 'K', 'rate', 'K')).toBe(100);
        expect(convertPrice(50, 'USD', 'M', 'USD', 'M')).toBe(50);
        expect(convertPrice(200, 'RMB', 'K', 'RMB', 'K')).toBe(200);
      });

      test('Rate 到 USD 转换', () => {
        // Rate(K) → USD(K): × 0.002
        expect(convertPrice(1000, 'rate', 'K', 'USD', 'K')).toBe(2);
        expect(convertPrice(500, 'rate', 'K', 'USD', 'K')).toBe(1);
      });

      test('Rate 到 RMB 转换', () => {
        // Rate(K) → RMB(K): × 0.014
        expect(convertPrice(1000, 'rate', 'K', 'RMB', 'K')).toBe(14);
        expect(convertPrice(100, 'rate', 'K', 'RMB', 'K')).toBe(1.4);
      });

      test('USD 到 Rate 转换', () => {
        // USD(K) → Rate(K): ÷ 0.002
        expect(convertPrice(2, 'USD', 'K', 'rate', 'K')).toBe(1000);
        expect(convertPrice(1, 'USD', 'K', 'rate', 'K')).toBe(500);
      });

      test('RMB 到 Rate 转换', () => {
        // RMB(K) → Rate(K): ÷ 0.014
        expect(convertPrice(14, 'RMB', 'K', 'rate', 'K')).toBe(1000);
        expect(convertPrice(1.4, 'RMB', 'K', 'rate', 'K')).toBe(100);
      });
    });

    // 单位转换测试
    describe('数量单位转换', () => {
      test('K 到 M 转换', () => {
        // Rate(K) → Rate(M): ÷ 1000
        expect(convertPrice(1000, 'rate', 'K', 'rate', 'M')).toBe(1);
        expect(convertPrice(500, 'rate', 'K', 'rate', 'M')).toBe(0.5);
      });

      test('M 到 K 转换', () => {
        // Rate(M) → Rate(K): × 1000
        expect(convertPrice(1, 'rate', 'M', 'rate', 'K')).toBe(1000);
        expect(convertPrice(0.5, 'rate', 'M', 'rate', 'K')).toBe(500);
      });

      test('复合单位转换', () => {
        // USD(M) → RMB(K): M→K (×1000), USD→Rate (÷0.002), Rate→RMB (×0.014)
        const result = convertPrice(1, 'USD', 'M', 'RMB', 'K');
        const expected = (1 * 1000) / EXCHANGE_RATES.USD * EXCHANGE_RATES.RMB;
        expect(result).toBeCloseTo(expected, 4);
      });
    });

    // 边界值测试
    describe('边界值处理', () => {
      test('零值处理', () => {
        expect(convertPrice(0, 'rate', 'K', 'USD', 'K')).toBe(0);
        expect(convertPrice(0, 'USD', 'M', 'RMB', 'K')).toBe(0);
      });

      test('空值处理', () => {
        expect(convertPrice('', 'rate', 'K', 'USD', 'K')).toBe(0);
        expect(convertPrice(null, 'rate', 'K', 'USD', 'K')).toBe(0);
        expect(convertPrice(undefined, 'rate', 'K', 'USD', 'K')).toBe(0);
      });

      test('非数字值处理', () => {
        expect(convertPrice('abc', 'rate', 'K', 'USD', 'K')).toBe(0);
        expect(convertPrice(NaN, 'rate', 'K', 'USD', 'K')).toBe(0);
      });

      test('负数处理', () => {
        expect(convertPrice(-100, 'rate', 'K', 'USD', 'K')).toBe(-0.2);
      });

      test('小数精度', () => {
        const result = convertPrice(1, 'rate', 'K', 'USD', 'K');
        expect(result).toBe(0.002);
        
        const result2 = convertPrice(0.001, 'USD', 'K', 'rate', 'K');
        expect(result2).toBe(0.5);
      });
    });

    // 可逆性测试
    describe('转换可逆性', () => {
      test('Rate ↔ USD 可逆性', () => {
        const original = 1000;
        const converted = convertPrice(original, 'rate', 'K', 'USD', 'K');
        const reverted = convertPrice(converted, 'USD', 'K', 'rate', 'K');
        expect(reverted).toBeCloseTo(original, 4);
      });

      test('Rate ↔ RMB 可逆性', () => {
        const original = 500;
        const converted = convertPrice(original, 'rate', 'K', 'RMB', 'K');
        const reverted = convertPrice(converted, 'RMB', 'K', 'rate', 'K');
        expect(reverted).toBeCloseTo(original, 4);
      });

      test('K ↔ M 可逆性', () => {
        const original = 2000;
        const converted = convertPrice(original, 'rate', 'K', 'rate', 'M');
        const reverted = convertPrice(converted, 'rate', 'M', 'rate', 'K');
        expect(reverted).toBeCloseTo(original, 4);
      });

      test('复合转换可逆性', () => {
        const original = 1.5;
        const converted = convertPrice(original, 'USD', 'M', 'RMB', 'K');
        const reverted = convertPrice(converted, 'RMB', 'K', 'USD', 'M');
        expect(reverted).toBeCloseTo(original, 4);
      });
    });

    // 错误处理测试
    describe('错误处理', () => {
      test('无效货币类型', () => {
        const result = convertPrice(100, 'INVALID', 'K', 'USD', 'K');
        expect(result).toBe(100); // 应该返回原值
      });

      test('无效单位类型', () => {
        const result = convertPrice(100, 'rate', 'INVALID', 'USD', 'K');
        expect(result).toBe(100); // 应该返回原值
      });
    });
  });

  describe('convertPriceData', () => {
    test('批量转换价格对象', () => {
      const priceData = {
        input: 1000,
        output: 2000,
        model: 'test-model'
      };

      const result = convertPriceData(priceData, 'rate', 'K', 'USD', 'K');
      
      expect(result.input).toBe(2); // 1000 * 0.002
      expect(result.output).toBe(4); // 2000 * 0.002
      expect(result.model).toBe('test-model'); // 其他字段保持不变
    });
  });

  describe('isValidUnitType', () => {
    test('有效的单位类型', () => {
      expect(isValidUnitType('rate', 'K')).toBe(true);
      expect(isValidUnitType('USD', 'M')).toBe(true);
      expect(isValidUnitType('RMB', 'K')).toBe(true);
    });

    test('无效的单位类型', () => {
      expect(isValidUnitType('INVALID', 'K')).toBe(false);
      expect(isValidUnitType('rate', 'INVALID')).toBe(false);
      expect(isValidUnitType('INVALID', 'INVALID')).toBe(false);
    });
  });

  describe('getUnitLabel', () => {
    test('生成正确的单位标签', () => {
      expect(getUnitLabel('rate', 'K')).toBe('Rate(K)');
      expect(getUnitLabel('USD', 'M')).toBe('USD(M)');
      expect(getUnitLabel('RMB', 'K')).toBe('RMB(K)');
    });
  });

  // 性能测试
  describe('性能测试', () => {
    test('大量转换操作性能', () => {
      const start = performance.now();
      
      for (let i = 0; i < 1000; i++) {
        convertPrice(Math.random() * 1000, 'rate', 'K', 'USD', 'M');
      }
      
      const end = performance.now();
      const duration = end - start;
      
      // 1000次转换应该在100ms内完成
      expect(duration).toBeLessThan(100);
    });
  });
});
