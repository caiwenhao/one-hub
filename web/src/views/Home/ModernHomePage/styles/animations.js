import { keyframes } from '@mui/material/styles';

// 浮动动画
export const floatAnimation = keyframes`
  0%, 100% { 
    transform: translateY(0px) rotate(0deg); 
  }
  50% { 
    transform: translateY(-20px) rotate(180deg); 
  }
`;

// 脉冲发光动画
export const pulseGlowAnimation = keyframes`
  0%, 100% { 
    box-shadow: 0 0 20px rgba(66, 153, 225, 0.3); 
  }
  50% { 
    box-shadow: 0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3); 
  }
`;

// 渐变移动动画
export const gradientShiftAnimation = keyframes`
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

// 淡入动画
export const fadeInAnimation = keyframes`
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
`;

// 缩放悬浮动画
export const scaleHoverAnimation = keyframes`
  0% {
    transform: scale(1);
  }
  100% {
    transform: scale(1.05);
  }
`;

// 动画样式对象
export const animationStyles = {
  floating: {
    animation: `${floatAnimation} 6s ease-in-out infinite`
  },
  pulseGlow: {
    animation: `${pulseGlowAnimation} 3s ease-in-out infinite`
  },
  gradientShift: {
    background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
    backgroundSize: '400% 400%',
    animation: `${gradientShiftAnimation} 4s ease infinite`
  },
  fadeIn: {
    animation: `${fadeInAnimation} 0.8s ease-out`
  },
  hoverLift: {
    transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
    '&:hover': {
      transform: 'translateY(-8px) scale(1.02)',
      boxShadow: '0 20px 40px rgba(0,0,0,0.1)'
    }
  },
  glassEffect: {
    background: 'rgba(255, 255, 255, 0.1)',
    backdropFilter: 'blur(10px)',
    border: '1px solid rgba(255, 255, 255, 0.2)'
  }
};

export default animationStyles;
