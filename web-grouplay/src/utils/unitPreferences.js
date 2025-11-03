/**
 * 单位偏好管理工具
 * 
 * 提供全局单位偏好的存储、获取和管理功能
 * 支持不同页面的单位设置持久化
 * 
 * @author AI Assistant
 * @version 1.0.0
 */

// 存储键常量
const UNIT_PREFERENCE_KEY = 'user_unit_preferences';

// 默认单位配置
const DEFAULT_UNIT_CONFIG = {
  // 全局默认单位 - 从K改为M
  global: 'M',
  // 页面特定单位设置
  pages: {
    modelPrice: 'M',           // 模型价格页面
    unifiedPricing: 'M',       // 统一价格表
    pricingSingle: 'M',        // 单一价格管理
    pricingMultiple: 'M'       // 批量价格管理
  }
};

/**
 * 获取单位偏好设置
 * @param {string} pageKey - 页面标识符（可选）
 * @returns {string} 单位偏好 ('K' 或 'M')
 */
export const getUnitPreference = (pageKey = null) => {
  try {
    const stored = localStorage.getItem(UNIT_PREFERENCE_KEY);
    let preferences = stored ? JSON.parse(stored) : DEFAULT_UNIT_CONFIG;
    
    // 确保配置结构完整
    if (!preferences.pages) {
      preferences = { ...DEFAULT_UNIT_CONFIG };
    }
    
    // 返回页面特定设置或全局设置
    if (pageKey && preferences.pages[pageKey]) {
      return preferences.pages[pageKey];
    }
    
    return preferences.global || DEFAULT_UNIT_CONFIG.global;
  } catch (error) {
    console.error('Error reading unit preferences:', error);
    return DEFAULT_UNIT_CONFIG.global;
  }
};

/**
 * 保存单位偏好设置
 * @param {string} unit - 单位值 ('K' 或 'M')
 * @param {string} pageKey - 页面标识符（可选）
 */
export const setUnitPreference = (unit, pageKey = null) => {
  try {
    const stored = localStorage.getItem(UNIT_PREFERENCE_KEY);
    let preferences = stored ? JSON.parse(stored) : DEFAULT_UNIT_CONFIG;
    
    // 确保配置结构完整
    if (!preferences.pages) {
      preferences = { ...DEFAULT_UNIT_CONFIG };
    }
    
    if (pageKey) {
      // 设置页面特定偏好
      preferences.pages[pageKey] = unit;
    } else {
      // 设置全局偏好
      preferences.global = unit;
    }
    
    localStorage.setItem(UNIT_PREFERENCE_KEY, JSON.stringify(preferences));
  } catch (error) {
    console.error('Error saving unit preferences:', error);
  }
};

/**
 * 重置单位偏好为默认值
 */
export const resetUnitPreferences = () => {
  try {
    localStorage.setItem(UNIT_PREFERENCE_KEY, JSON.stringify(DEFAULT_UNIT_CONFIG));
  } catch (error) {
    console.error('Error resetting unit preferences:', error);
  }
};

/**
 * 获取所有单位偏好设置
 * @returns {object} 完整的偏好配置对象
 */
export const getAllUnitPreferences = () => {
  try {
    const stored = localStorage.getItem(UNIT_PREFERENCE_KEY);
    return stored ? JSON.parse(stored) : DEFAULT_UNIT_CONFIG;
  } catch (error) {
    console.error('Error reading all unit preferences:', error);
    return DEFAULT_UNIT_CONFIG;
  }
};

/**
 * 页面键常量 - 便于其他组件引用
 */
export const PAGE_KEYS = {
  MODEL_PRICE: 'modelPrice',
  UNIFIED_PRICING: 'unifiedPricing', 
  PRICING_SINGLE: 'pricingSingle',
  PRICING_MULTIPLE: 'pricingMultiple'
};

/**
 * 单位选项常量
 */
export const UNIT_OPTIONS = [
  { value: 'K', label: 'K' },
  { value: 'M', label: 'M' }
];

/**
 * 创建带有偏好存储的单位状态Hook
 * @param {string} pageKey - 页面标识符
 * @returns {[string, function]} [当前单位, 设置单位函数]
 */
export const createUnitState = (pageKey) => {
  const initialUnit = getUnitPreference(pageKey);
  
  const setUnitWithPreference = (unit) => {
    setUnitPreference(unit, pageKey);
    return unit;
  };
  
  return [initialUnit, setUnitWithPreference];
};
