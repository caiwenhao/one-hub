// 渐变色配置
export const gradients = {
  // 主要渐变（AI 科技风：电光蓝 → 电子紫）
  primary: 'linear-gradient(135deg, #0EA5FF 0%, #8B5CF6 100%)',
  
  // 英雄区域渐变
  hero: 'linear-gradient(135deg, #0EA5FF 0%, #8B5CF6 100%)',
  heroBackground: 'linear-gradient(to bottom right, #ffffff 0%, rgba(14,165,255,0.06) 30%, rgba(139,92,246,0.06) 70%, #ffffff 100%)',
  
  // 卡片渐变
  card: 'linear-gradient(145deg, #ffffff 0%, #f8fafc 100%)',
  
  // CTA按钮渐变
  cta: 'linear-gradient(135deg, #0EA5FF 0%, #8B5CF6 100%)',
  
  // 文字渐变
  textGradient: 'linear-gradient(to right, #22D3EE, #0EA5FF, #8B5CF6)',
  
  // 特性卡片渐变
  featureAccent: 'linear-gradient(to bottom right, rgba(14, 165, 255, 0.06), transparent)',
  featurePurple: 'linear-gradient(to bottom right, rgba(139, 92, 246, 0.06), transparent)',
  featurePink: 'linear-gradient(to bottom right, rgba(236, 72, 153, 0.05), transparent)',
  featureOrange: 'linear-gradient(to bottom right, rgba(245, 158, 11, 0.05), transparent)',
  
  // 图标背景渐变
  iconAccent: 'linear-gradient(to bottom right, rgba(14, 165, 255, 0.25), rgba(139, 92, 246, 0.25))',
  iconPurple: 'linear-gradient(to bottom right, rgba(139, 92, 246, 0.2), rgba(236, 72, 153, 0.2))',
  iconPink: 'linear-gradient(to bottom right, rgba(236, 72, 153, 0.2), rgba(245, 158, 11, 0.2))',
  iconOrange: 'linear-gradient(to bottom right, rgba(245, 158, 11, 0.2), rgba(66, 153, 225, 0.2))',
  
  // 统计区域渐变
  statsBackground: 'linear-gradient(to bottom right, #1a202c 0%, #2d3748 50%, #1a202c 100%)',
  statsOverlay: 'linear-gradient(to right, rgba(66, 153, 225, 0.1), rgba(139, 92, 246, 0.1))',
  
  // 最终CTA区域渐变
  finalCTA: 'linear-gradient(to bottom right, #1a202c 0%, #2d3748 50%, #1a202c 100%)',
  finalCTAOverlay: 'linear-gradient(to right, rgba(66, 153, 225, 0.15), rgba(139, 92, 246, 0.15))',
  
  // 浮动元素渐变
  floatingBlur1: 'linear-gradient(to right, rgba(66, 153, 225, 0.1), rgba(139, 92, 246, 0.1))',
  floatingBlur2: 'linear-gradient(to right, rgba(236, 72, 153, 0.1), rgba(245, 158, 11, 0.1))',
  
  // 模型卡片渐变
  modelGPT: 'linear-gradient(to bottom right, #10b981, #059669)',
  modelClaude: 'linear-gradient(to bottom right, #f97316, #ea580c)',
  modelGemini: 'linear-gradient(to bottom right, #3b82f6, #2563eb)'
};

// 颜色配置
export const colors = {
  primary: '#1A202C',
  accent: '#0EA5FF',
  accentLight: '#7BD3FF',
  accentDark: '#0369A1',
  background: '#F7FAFC',
  secondary: '#718096',
  purple: '#8B5CF6',
  pink: '#EC4899',
  orange: '#F59E0B',
  
  // 透明度变体
  accentAlpha: {
    5: 'rgba(14, 165, 255, 0.05)',
    10: 'rgba(14, 165, 255, 0.10)',
    20: 'rgba(14, 165, 255, 0.20)',
    30: 'rgba(14, 165, 255, 0.30)',
    40: 'rgba(14, 165, 255, 0.40)'
  },
  
  purpleAlpha: {
    5: 'rgba(139, 92, 246, 0.05)',
    10: 'rgba(139, 92, 246, 0.1)',
    20: 'rgba(139, 92, 246, 0.2)',
    30: 'rgba(139, 92, 246, 0.3)',
    40: 'rgba(139, 92, 246, 0.4)'
  },
  
  pinkAlpha: {
    5: 'rgba(236, 72, 153, 0.05)',
    10: 'rgba(236, 72, 153, 0.1)',
    20: 'rgba(236, 72, 153, 0.2)',
    30: 'rgba(236, 72, 153, 0.3)',
    35: 'rgba(236, 72, 153, 0.35)'
  },
  
  orangeAlpha: {
    5: 'rgba(245, 158, 11, 0.05)',
    10: 'rgba(245, 158, 11, 0.1)',
    20: 'rgba(245, 158, 11, 0.2)',
    40: 'rgba(245, 158, 11, 0.4)'
  }
};

export default { gradients, colors };
