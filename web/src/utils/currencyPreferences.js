/**
 * 币种偏好与格式化工具
 * - 支持页面级与全局级偏好
 * - 默认币种：CNY（人民币）
 */

export const CURRENCY_STORAGE_KEY = 'user_currency_preferences';

const DEFAULT_CONFIG = {
  global: 'CNY',
  pages: {
    modelPrice: 'CNY',
    unifiedPricing: 'CNY',
    log: 'CNY'
  }
};

export const CURRENCY_OPTIONS = [
  { value: 'CNY', label: '￥ CNY' },
  { value: 'USD', label: '$ USD' }
];

export function getCurrencyPreference(pageKey = null) {
  try {
    const stored = localStorage.getItem(CURRENCY_STORAGE_KEY);
    let pref = stored ? JSON.parse(stored) : DEFAULT_CONFIG;
    if (!pref.pages) pref = { ...DEFAULT_CONFIG };
    if (pageKey && pref.pages[pageKey]) return pref.pages[pageKey];
    return pref.global || DEFAULT_CONFIG.global;
  } catch (e) {
    console.error('Error reading currency preferences:', e);
    return DEFAULT_CONFIG.global;
  }
}

export function setCurrencyPreference(currency, pageKey = null) {
  try {
    const stored = localStorage.getItem(CURRENCY_STORAGE_KEY);
    let pref = stored ? JSON.parse(stored) : DEFAULT_CONFIG;
    if (!pref.pages) pref = { ...DEFAULT_CONFIG };
    if (pageKey) {
      pref.pages[pageKey] = currency;
    } else {
      pref.global = currency;
    }
    localStorage.setItem(CURRENCY_STORAGE_KEY, JSON.stringify(pref));
  } catch (e) {
    console.error('Error saving currency preferences:', e);
  }
}

export const CURRENCY_PAGE_KEYS = {
  MODEL_PRICE: 'modelPrice',
  UNIFIED_PRICING: 'unifiedPricing',
  LOG: 'log'
};

// 汇率与单位换算（与后端/priceConverter 保持一致）
export const RATE_PER_K_TO_USD = 0.002; // 1 Rate(K) → $0.002
export const RATE_PER_K_TO_CNY = 0.014; // 1 Rate(K) → ￥0.014

export function usdToCnyRate() {
  return RATE_PER_K_TO_CNY / RATE_PER_K_TO_USD; // 0.014/0.002 = 7
}

/**
 * 根据币种格式化“倍率/Rate(K)”价格
 * @param {number} value - 价格倍率（例如 15 表示 $0.03/1K）
 * @param {string} currency - 'CNY' | 'USD'
 * @param {boolean} unitMillion - 是否按每百万显示（M→K*1000）
 */
export function formatPriceByCurrency(value, currency = 'CNY', unitMillion = false) {
  if (value == null) return '';
  if (value === 0) return 'Free';
  const mult = unitMillion ? 1000 : 1;
  const v = Number(value) * mult;
  let amount = 0;
  if (currency === 'USD') {
    amount = v * RATE_PER_K_TO_USD;
    return `$${trimTrailingZeros(amount.toFixed(6))}`;
  }
  // default CNY
  amount = v * RATE_PER_K_TO_CNY;
  return `￥${trimTrailingZeros(amount.toFixed(6))}`;
}

function trimTrailingZeros(s) {
  return s.replace(/(\.\d*?[1-9])0+$|\.0*$/, '$1');
}
