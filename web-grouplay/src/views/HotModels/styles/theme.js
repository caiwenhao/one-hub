// 热门模型页面的主题样式定义
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
  green: '#10B981',
  red: '#EF4444',
  blue: '#3B82F6'
};

export const gradients = {
  primary: 'linear-gradient(135deg, #0EA5FF 0%, #8B5CF6 100%)',
  heroBackground: 'linear-gradient(135deg, #0EA5FF 0%, #8B5CF6 100%)',
  card: 'linear-gradient(145deg, #ffffff 0%, #f8fafc 100%)',
  textGradient: 'linear-gradient(45deg, #22D3EE 0%, #0EA5FF 50%, #8B5CF6 100%)',
  
  // 特性卡片渐变
  featureAccent: 'linear-gradient(135deg, rgba(14, 165, 255, 0.06), transparent)',
  featurePurple: 'linear-gradient(135deg, rgba(139, 92, 246, 0.05), transparent)',
  featurePink: 'linear-gradient(135deg, rgba(236, 72, 153, 0.05), transparent)',
  featureOrange: 'linear-gradient(135deg, rgba(245, 158, 11, 0.05), transparent)',
  
  // 图标渐变
  iconAccent: 'linear-gradient(135deg, #0EA5FF, #8B5CF6)',
  iconPurple: 'linear-gradient(135deg, #8B5CF6, #7C3AED)',
  iconPink: 'linear-gradient(135deg, #EC4899, #DB2777)',
  iconOrange: 'linear-gradient(135deg, #F59E0B, #D97706)',
  iconGreen: 'linear-gradient(135deg, #10B981, #059669)',
  iconBlue: 'linear-gradient(135deg, #3B82F6, #2563EB)'
};

// 动画样式
export const animationStyles = {
  fadeIn: {
    '@keyframes fadeIn': {
      '0%': {
        opacity: 0,
        transform: 'translateY(30px)'
      },
      '100%': {
        opacity: 1,
        transform: 'translateY(0)'
      }
    },
    animation: 'fadeIn 0.8s ease-out forwards'
  },
  
  hoverLift: {
    transition: 'all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275)',
    '&:hover': {
      transform: 'translateY(-12px) scale(1.03)',
      boxShadow: '0 25px 50px rgba(0,0,0,0.15)'
    }
  },
  
  gradientShift: {
    '@keyframes gradientShift': {
      '0%': { backgroundPosition: '0% 50%' },
      '50%': { backgroundPosition: '100% 50%' },
      '100%': { backgroundPosition: '0% 50%' }
    },
    backgroundSize: '400% 400%',
    animation: 'gradientShift 4s ease infinite'
  }
};

// 创建渐变文本的工具函数
export const createGradientText = (gradient = gradients.textGradient) => ({
  background: gradient,
  WebkitBackgroundClip: 'text',
  WebkitTextFillColor: 'transparent',
  backgroundClip: 'text'
});

// 标签样式
export const tagStyles = {
  hot: {
    background: 'linear-gradient(135deg, #FF6B6B, #FF8E8E)',
    color: 'white',
    fontWeight: 'bold'
  },
  new: {
    background: 'linear-gradient(135deg, #4ECDC4, #44A08D)',
    color: 'white',
    fontWeight: 'bold'
  },
  recommended: {
    background: 'linear-gradient(135deg, #A8E6CF, #7FCDCD)',
    color: 'white',
    fontWeight: 'bold'
  }
};
