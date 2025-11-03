// 设计令牌（企业稳重风格）
// 说明：本文件作为全局唯一真源，后续主题与组件应尽量从此处取值。

export const tokens = {
  color: {
    brand: {
      primary: '#1677ff', // 企业蓝（接近 Antd 默认）
      primaryDark: '#1554c5',
      secondary: '#0f172a' // 深石板
    },
    // 中性色阶（符合 AA 对比度参考）
    grey: {
      50: '#FAFAFA',
      100: '#F5F5F5',
      200: '#EEEEEE',
      300: '#E0E0E0',
      500: '#9E9E9E',
      600: '#757575',
      700: '#616161',
      900: '#212121'
    },
    text: {
      primary: '#1f2937',
      secondary: '#6b7280',
      inverse: '#ffffff'
    },
    bg: {
      page: '#f7f9fc',
      card: '#ffffff',
      subtle: '#f3f4f6'
    },
    border: {
      default: '#e5e7eb',
      strong: '#d1d5db'
    },
    success: '#10B981',
    warning: '#F59E0B',
    error: '#EF4444',
    info: '#3B82F6'
  },
  radius: { xs: 6, sm: 8, md: 10, lg: 12 },
  shadow: {
    sm: '0 1px 2px rgba(0,0,0,0.06)',
    md: '0 2px 8px rgba(0,0,0,0.08)',
    lg: '0 8px 24px rgba(0,0,0,0.10)'
  },
  spacing: { xs: 4, sm: 8, md: 12, lg: 16, xl: 24, xxl: 32 },
  typography: {
    fontFamily: "'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif",
    size: { body: 14, h1: 32, h2: 24, h3: 20 }
  }
};

export default tokens;

