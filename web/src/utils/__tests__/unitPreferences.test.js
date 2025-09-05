/**
 * 单位偏好管理工具测试
 * 
 * 测试单位偏好的存储、获取和管理功能
 * 
 * @author AI Assistant
 * @version 1.0.0
 */

import {
  getUnitPreference,
  setUnitPreference,
  resetUnitPreferences,
  getAllUnitPreferences,
  PAGE_KEYS,
  UNIT_OPTIONS
} from '../unitPreferences';

// Mock localStorage
const localStorageMock = (() => {
  let store = {};
  return {
    getItem: (key) => store[key] || null,
    setItem: (key, value) => {
      store[key] = value.toString();
    },
    removeItem: (key) => {
      delete store[key];
    },
    clear: () => {
      store = {};
    }
  };
})();

Object.defineProperty(window, 'localStorage', {
  value: localStorageMock
});

describe('unitPreferences', () => {
  beforeEach(() => {
    localStorage.clear();
  });

  describe('getUnitPreference', () => {
    test('应该返回默认的M单位（全局）', () => {
      const unit = getUnitPreference();
      expect(unit).toBe('M');
    });

    test('应该返回页面特定的M单位', () => {
      const unit = getUnitPreference(PAGE_KEYS.MODEL_PRICE);
      expect(unit).toBe('M');
    });

    test('应该处理损坏的localStorage数据', () => {
      localStorage.setItem('user_unit_preferences', 'invalid json');
      const unit = getUnitPreference();
      expect(unit).toBe('M');
    });
  });

  describe('setUnitPreference', () => {
    test('应该设置全局单位偏好', () => {
      setUnitPreference('K');
      const unit = getUnitPreference();
      expect(unit).toBe('K');
    });

    test('应该设置页面特定单位偏好', () => {
      setUnitPreference('K', PAGE_KEYS.MODEL_PRICE);
      const unit = getUnitPreference(PAGE_KEYS.MODEL_PRICE);
      expect(unit).toBe('K');
    });

    test('应该保持其他页面设置不变', () => {
      setUnitPreference('K', PAGE_KEYS.MODEL_PRICE);
      const modelPriceUnit = getUnitPreference(PAGE_KEYS.MODEL_PRICE);
      const unifiedPricingUnit = getUnitPreference(PAGE_KEYS.UNIFIED_PRICING);
      
      expect(modelPriceUnit).toBe('K');
      expect(unifiedPricingUnit).toBe('M'); // 应该保持默认值
    });
  });

  describe('resetUnitPreferences', () => {
    test('应该重置所有偏好为默认值', () => {
      setUnitPreference('K');
      setUnitPreference('K', PAGE_KEYS.MODEL_PRICE);
      
      resetUnitPreferences();
      
      const globalUnit = getUnitPreference();
      const pageUnit = getUnitPreference(PAGE_KEYS.MODEL_PRICE);
      
      expect(globalUnit).toBe('M');
      expect(pageUnit).toBe('M');
    });
  });

  describe('getAllUnitPreferences', () => {
    test('应该返回完整的偏好配置', () => {
      const preferences = getAllUnitPreferences();
      
      expect(preferences).toHaveProperty('global', 'M');
      expect(preferences).toHaveProperty('pages');
      expect(preferences.pages).toHaveProperty('modelPrice', 'M');
      expect(preferences.pages).toHaveProperty('unifiedPricing', 'M');
      expect(preferences.pages).toHaveProperty('pricingSingle', 'M');
      expect(preferences.pages).toHaveProperty('pricingMultiple', 'M');
    });
  });

  describe('PAGE_KEYS', () => {
    test('应该包含所有必要的页面键', () => {
      expect(PAGE_KEYS).toHaveProperty('MODEL_PRICE', 'modelPrice');
      expect(PAGE_KEYS).toHaveProperty('UNIFIED_PRICING', 'unifiedPricing');
      expect(PAGE_KEYS).toHaveProperty('PRICING_SINGLE', 'pricingSingle');
      expect(PAGE_KEYS).toHaveProperty('PRICING_MULTIPLE', 'pricingMultiple');
    });
  });

  describe('UNIT_OPTIONS', () => {
    test('应该包含K和M选项', () => {
      expect(UNIT_OPTIONS).toEqual([
        { value: 'K', label: 'K' },
        { value: 'M', label: 'M' }
      ]);
    });
  });

  describe('错误处理', () => {
    test('应该处理localStorage写入错误', () => {
      // Mock localStorage.setItem to throw error
      const originalSetItem = localStorage.setItem;
      localStorage.setItem = jest.fn(() => {
        throw new Error('Storage quota exceeded');
      });

      // 应该不抛出错误
      expect(() => setUnitPreference('K')).not.toThrow();

      // 恢复原始方法
      localStorage.setItem = originalSetItem;
    });

    test('应该处理localStorage读取错误', () => {
      // Mock localStorage.getItem to throw error
      const originalGetItem = localStorage.getItem;
      localStorage.getItem = jest.fn(() => {
        throw new Error('Storage access denied');
      });

      // 应该返回默认值
      const unit = getUnitPreference();
      expect(unit).toBe('M');

      // 恢复原始方法
      localStorage.getItem = originalGetItem;
    });
  });
});
