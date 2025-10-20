// 多语言长度回退与本地化格式化工具
import i18n from 'i18n/i18n';

export function currentLocale() {
  return i18n?.language || navigator?.language || 'zh-CN';
}

export function formatNumber(value, options) {
  try {
    return new Intl.NumberFormat(currentLocale(), options).format(value);
  } catch {
    return String(value);
  }
}

export function formatDate(date, options) {
  try {
    const d = typeof date === 'number' ? new Date(date) : date;
    return new Intl.DateTimeFormat(currentLocale(), options || { year: 'numeric', month: '2-digit', day: '2-digit' }).format(d);
  } catch {
    return String(date);
  }
}

// 将数值与单位连接，默认使用不换行空格
export function joinWithUnit(value, unit, { nbsp = true } = {}) {
  const sep = nbsp ? '\u00A0' : ' ';
  return `${value}${sep}${unit}`;
}
