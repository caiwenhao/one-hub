import React from 'react';
import { Box, Container, Typography } from '@mui/material';
import { colors, gradients, animationStyles, createGradientText } from '../styles/theme';

const HeroSection = () => {
  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        background: 'linear-gradient(to bottom right, #ffffff, rgba(59, 130, 246, 0.03), rgba(139, 92, 246, 0.02))',
        pt: { xs: 10, md: 16 },
        pb: { xs: 10, md: 16 },
        minHeight: '600px',
        display: 'flex',
        alignItems: 'center',
        overflow: 'hidden'
      }}
    >
      {/* 背景渐变层 */}
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          background: 'linear-gradient(to bottom right, rgba(66, 153, 225, 0.08) 0%, rgba(139, 92, 246, 0.05) 50%, rgba(236, 72, 153, 0.08) 100%)'
        }}
      />

      {/* 浮动装饰元素 */}
      <Box
        sx={{
          position: 'absolute',
          top: 0,
          left: 0,
          width: '100%',
          height: '100%',
          pointerEvents: 'none'
        }}
      >
        {/* 浮动圆点 */}
        <Box
          sx={{
            position: 'absolute',
            top: '20%',
            left: '10%',
            width: '12px',
            height: '12px',
            backgroundColor: 'rgba(66, 153, 225, 0.3)',
            borderRadius: '50%',
            ...animationStyles.pulseGlow,
            '@keyframes float': {
              '0%, 100%': { transform: 'translateY(0px) rotate(0deg)' },
              '50%': { transform: 'translateY(-20px) rotate(180deg)' }
            },
            animation: 'float 6s ease-in-out infinite'
          }}
        />
        <Box
          sx={{
            position: 'absolute',
            top: '40%',
            right: '15%',
            width: '8px',
            height: '8px',
            backgroundColor: 'rgba(139, 92, 246, 0.4)',
            borderRadius: '50%',
            ...animationStyles.pulseGlow,
            animation: 'float 6s ease-in-out infinite 0.3s'
          }}
        />
        <Box
          sx={{
            position: 'absolute',
            bottom: '30%',
            left: '25%',
            width: '10px',
            height: '10px',
            backgroundColor: 'rgba(236, 72, 153, 0.35)',
            borderRadius: '50%',
            ...animationStyles.pulseGlow,
            animation: 'float 6s ease-in-out infinite 0.7s'
          }}
        />
        <Box
          sx={{
            position: 'absolute',
            top: '60%',
            right: '30%',
            width: '6px',
            height: '6px',
            backgroundColor: 'rgba(245, 158, 11, 0.4)',
            borderRadius: '50%',
            ...animationStyles.pulseGlow,
            animation: 'float 6s ease-in-out infinite 0.5s'
          }}
        />
      </Box>

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 10, maxWidth: '1200px' }}>
        <Box
          sx={{
            textAlign: 'center',
            maxWidth: '900px',
            mx: 'auto',
            ...animationStyles.fadeIn
          }}
        >
          {/* 主标题 - 已删除"连接全球顶尖AI大脑"区块 */}

          {/* 副标题 */}
          <Typography
            variant="h5"
            sx={{
              fontSize: { xs: '1.125rem', sm: '1.25rem', md: '1.5rem' },
              color: colors.secondary,
              mb: { xs: 4, md: 6 },
              maxWidth: { xs: '100%', md: '700px' },
              mx: 'auto',
              lineHeight: 1.6,
              fontWeight: 300,
              textShadow: '0 2px 4px rgba(0,0,0,0.1)',
              px: { xs: 3, md: 0 }
            }}
          >
            我们实时同步全球最前沿、性能最卓越的大语言模型
            <br />
            让您的应用接入世界级AI能力，体验前所未有的智能
          </Typography>
        </Box>
      </Container>
    </Box>
  );
};

export default HeroSection;
