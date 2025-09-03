import { keyframes } from '@mui/material';

// 浮动动画
export const float = keyframes`
  0%, 100% { 
    transform: translateY(0px) rotate(0deg); 
  }
  50% { 
    transform: translateY(-20px) rotate(180deg); 
  }
`;

// 发光脉冲动画
export const pulseGlow = keyframes`
  0%, 100% { 
    box-shadow: 0 0 20px rgba(66, 153, 225, 0.3); 
  }
  50% { 
    box-shadow: 0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3); 
  }
`;

// 渐变移动动画
export const gradientShift = keyframes`
  0% { 
    background-position: 0% 50%; 
  }
  50% { 
    background-position: 100% 50%; 
  }
  100% { 
    background-position: 0% 50%; 
  }
`;

// 闪光动画
export const shimmer = keyframes`
  0% { 
    transform: translateX(-100%); 
  }
  100% { 
    transform: translateX(100%); 
  }
`;

// 弹跳动画
export const bounce = keyframes`
  0%, 20%, 53%, 100% {
    animation-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
    transform: translateZ(0);
  }
  40%, 43% {
    animation-timing-function: cubic-bezier(0.755, 0.05, 0.855, 0.06);
    transform: translate3d(0, -5px, 0);
  }
  70% {
    animation-timing-function: cubic-bezier(0.755, 0.05, 0.855, 0.06);
    transform: translate3d(0, -7px, 0);
  }
  80% {
    transition-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
    transform: translateZ(0);
  }
  90% {
    transform: translate3d(0, -2px, 0);
  }
`;

// 淡入动画
export const fadeIn = keyframes`
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
`;

// 缩放动画
export const scaleIn = keyframes`
  from {
    opacity: 0;
    transform: scale(0.8);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
`;

// 滑入动画
export const slideInUp = keyframes`
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
`;

// 滑入动画（从左）
export const slideInLeft = keyframes`
  from {
    opacity: 0;
    transform: translateX(-30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
`;

// 滑入动画（从右）
export const slideInRight = keyframes`
  from {
    opacity: 0;
    transform: translateX(30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
`;

// 旋转动画
export const rotate = keyframes`
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
`;

// 心跳动画
export const heartbeat = keyframes`
  0% {
    transform: scale(1);
  }
  14% {
    transform: scale(1.1);
  }
  28% {
    transform: scale(1);
  }
  42% {
    transform: scale(1.1);
  }
  70% {
    transform: scale(1);
  }
`;

// 摇摆动画
export const swing = keyframes`
  20% {
    transform: rotate3d(0, 0, 1, 15deg);
  }
  40% {
    transform: rotate3d(0, 0, 1, -10deg);
  }
  60% {
    transform: rotate3d(0, 0, 1, 5deg);
  }
  80% {
    transform: rotate3d(0, 0, 1, -5deg);
  }
  100% {
    transform: rotate3d(0, 0, 1, 0deg);
  }
`;

// 波浪动画
export const wave = keyframes`
  0%, 60%, 100% {
    transform: initial;
  }
  30% {
    transform: translateY(-15px);
  }
`;

// 呼吸动画
export const breathe = keyframes`
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.05);
    opacity: 0.8;
  }
`;

// 彩虹渐变动画
export const rainbow = keyframes`
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
`;

// 动画工具函数
export const createAnimationStyles = (theme) => ({
  // 浮动元素样式
  floatingElement: {
    animation: `${float} 6s ease-in-out infinite`
  },
  
  // 发光效果样式
  glowEffect: {
    animation: `${pulseGlow} 3s ease-in-out infinite`
  },
  
  // 渐变动画样式
  animatedGradient: {
    background: `linear-gradient(-45deg, ${theme.palette.primary.main}, #3182CE, #2B6CB0, #2A69AC)`,
    backgroundSize: '400% 400%',
    animation: `${gradientShift} 4s ease infinite`
  },
  
  // 闪光效果样式
  shimmerEffect: {
    position: 'relative',
    overflow: 'hidden',
    '&::before': {
      content: '""',
      position: 'absolute',
      top: 0,
      left: 0,
      width: '100%',
      height: '100%',
      background: 'linear-gradient(90deg, transparent, rgba(255,255,255,0.4), transparent)',
      transform: 'translateX(-100%)',
      animation: `${shimmer} 2s infinite`
    }
  },
  
  // 玻璃效果样式
  glassEffect: {
    background: 'rgba(255, 255, 255, 0.1)',
    backdropFilter: 'blur(10px)',
    border: '1px solid rgba(255, 255, 255, 0.2)'
  },
  
  // 悬浮提升效果
  hoverLift: {
    transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
    '&:hover': {
      transform: 'translateY(-8px) scale(1.02)',
      boxShadow: '0 20px 40px rgba(0,0,0,0.1)'
    }
  },
  
  // 价格卡片效果
  priceCard: {
    transition: 'all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275)',
    position: 'relative',
    overflow: 'hidden',
    '&:hover': {
      transform: 'translateY(-12px) scale(1.03)',
      boxShadow: '0 25px 50px rgba(0,0,0,0.15)'
    }
  },
  
  // 文字阴影效果
  textShadow: {
    textShadow: '0 2px 4px rgba(0,0,0,0.1)'
  }
});

// 延迟动画工具函数
export const createDelayedAnimation = (animationName, delay = 0) => ({
  animation: `${animationName} 0.6s ease-out ${delay}s both`
});

// 交错动画工具函数
export const createStaggeredAnimation = (animationName, index, baseDelay = 0.1) => ({
  animation: `${animationName} 0.6s ease-out ${index * baseDelay}s both`
});
