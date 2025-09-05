import React from 'react';
import { Box, Typography, Container } from '@mui/material';
import { keyframes } from '@mui/system';

// 定义动画
const float = keyframes`
  0%, 100% { 
    transform: translateY(0px) rotate(0deg); 
  }
  50% { 
    transform: translateY(-20px) rotate(180deg); 
  }
`;

const pulseGlow = keyframes`
  0%, 100% { 
    box-shadow: 0 0 20px rgba(66, 153, 225, 0.3); 
  }
  50% { 
    box-shadow: 0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3); 
  }
`;

const gradientShift = keyframes`
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
`;

// 浮动元素组件
const FloatingElement = ({ size, color, top, left, right, bottom, delay = 0 }) => (
  <Box
    sx={{
      position: 'absolute',
      width: size,
      height: size,
      backgroundColor: color,
      borderRadius: '50%',
      top,
      left,
      right,
      bottom,
      animation: `${float} 6s ease-in-out infinite, ${pulseGlow} 3s ease-in-out infinite`,
      animationDelay: `${delay}ms`,
      zIndex: 1
    }}
  />
);

const HeroSection = () => {
  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        backgroundColor: '#ffffff',
        pt: { xs: 8, md: 12 },
        pb: { xs: 8, md: 10 },
        px: { xs: 3, md: 6, lg: 12 },
        height: '500px',
        display: 'flex',
        alignItems: 'center',
        overflow: 'hidden',
        '&::before': {
          content: '""',
          position: 'absolute',
          inset: 0,
          background: 'linear-gradient(to bottom right, rgba(249, 250, 251, 0.5), rgba(59, 130, 246, 0.03))',
          zIndex: 0
        }
      }}
    >
      {/* 浮动动画元素 */}
      <FloatingElement 
        size="12px" 
        color="rgba(66, 153, 225, 0.3)" 
        top="80px" 
        left="40px" 
        delay={0}
      />
      <FloatingElement 
        size="8px" 
        color="rgba(34, 197, 94, 0.4)" 
        top="160px" 
        right="80px" 
        delay={300}
      />
      <FloatingElement 
        size="10px" 
        color="rgba(147, 51, 234, 0.35)" 
        bottom="128px" 
        left="25%" 
        delay={700}
      />
      <FloatingElement 
        size="6px" 
        color="rgba(249, 115, 22, 0.4)" 
        top="240px" 
        right="33%" 
        delay={500}
      />

      {/* 主要内容 */}
      <Container 
        maxWidth="lg" 
        sx={{ 
          position: 'relative', 
          zIndex: 10, 
          textAlign: 'center',
          maxWidth: '1200px'
        }}
      >
        <Typography
          variant="h1"
          sx={{
            fontSize: { xs: '2.5rem', md: '3.75rem' }, // text-5xl md:text-6xl
            fontWeight: 200, // font-extralight
            color: '#1A202C', // text-primary
            mb: 4,
            lineHeight: 1.1, // leading-tight
            letterSpacing: '-0.025em', // tracking-tight
            textShadow: '0 2px 4px rgba(0,0,0,0.1)' // text-shadow
          }}
        >
          我们随时
          <Box
            component="span"
            sx={{
              fontWeight: 700, // font-bold
              background: 'linear-gradient(45deg, #4299E1, #10B981, #8B5CF6)',
              backgroundSize: '400% 400%',
              backgroundClip: 'text',
              WebkitBackgroundClip: 'text',
              color: 'transparent',
              animation: `${gradientShift} 4s ease infinite`
            }}
          >
            倾听您的需求
          </Box>
        </Typography>
        
        <Typography
          variant="h5"
          sx={{
            fontSize: { xs: '1.125rem', md: '1.25rem' }, // text-xl
            color: '#718096', // text-gray-600
            mb: 4,
            maxWidth: '768px', // max-w-3xl
            mx: 'auto',
            lineHeight: 1.6, // leading-relaxed
            fontWeight: 300, // font-light
            textShadow: '0 2px 4px rgba(0,0,0,0.1)' // text-shadow
          }}
        >
          无论您是寻求企业合作、需要技术支持，还是有任何疑问
          <br />
          都欢迎随时联系我们。我们的团队将热情为您服务
        </Typography>
      </Container>
    </Box>
  );
};

export default HeroSection;
