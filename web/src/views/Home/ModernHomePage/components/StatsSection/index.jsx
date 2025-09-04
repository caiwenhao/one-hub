import React from 'react';
import {
  Box,
  Container,
  Typography,
  Grid,
  useTheme,
  useMediaQuery
} from '@mui/material';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGradientText, createFloatingElement } from '../../styles/theme';

const StatsSection = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));

  const stats = [
    {
      value: '99.9%',
      label: '服务可用性',
      description: '企业级稳定保障'
    },
    {
      value: '50ms',
      label: '平均响应时间',
      description: '极速API调用'
    },
    {
      value: '24/7',
      label: '技术支持',
      description: '全天候服务'
    },
    {
      value: '10K+',
      label: '开发者信赖',
      description: '活跃用户社区'
    }
  ];

  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        background: gradients.statsBackground,
        py: { xs: 3, md: 6 },
        overflow: 'hidden'
      }}
    >
      {/* 背景渐变层 */}
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          background: gradients.statsOverlay
        }}
      />

      {/* 浮动装饰元素 */}
      <Box
        sx={{
          ...createFloatingElement('8px', colors.accentAlpha[40]),
          top: '10%',
          left: '25%',
          ...animationStyles.floating,
          ...animationStyles.pulseGlow
        }}
      />
      <Box
        sx={{
          ...createFloatingElement('6px', colors.purpleAlpha[40]),
          bottom: '20%',
          right: '33%',
          ...animationStyles.floating,
          animationDelay: '0.3s',
          ...animationStyles.pulseGlow
        }}
      />

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1 }}>
        {/* 标题 */}
        <Box
          sx={{
            textAlign: 'center',
            mb: { xs: 2, md: 4 },
            ...animationStyles.fadeIn
          }}
        >
          <Typography
            variant="h2"
            sx={{
              fontSize: { xs: '2.5rem', md: '3rem', lg: '3.5rem' },
              fontWeight: 'bold',
              color: '#ffffff',
              mb: 2,
              textShadow: '0 2px 4px rgba(0,0,0,0.3)'
            }}
          >
            数据说话，实力证明
          </Typography>
        </Box>

        {/* 统计数据网格 */}
        <Grid container spacing={{ xs: 4, md: 6 }}>
          {stats.map((stat, index) => (
            <Grid item xs={6} md={3} key={stat.label}>
              <Box
                sx={{
                  textAlign: 'center',
                  ...animationStyles.fadeIn,
                  animationDelay: `${index * 0.2}s`,
                  transition: 'all 0.3s ease',
                  '&:hover': {
                    transform: 'translateY(-8px)'
                  }
                }}
              >
                {/* 数值 */}
                <Typography
                  sx={{
                    fontSize: { xs: '3rem', md: '4rem', lg: '5rem' },
                    fontWeight: 'bold',
                    mb: 2,
                    lineHeight: 1,
                    ...createGradientText(),
                    ...animationStyles.gradientShift,
                    textShadow: '0 4px 8px rgba(0,0,0,0.2)'
                  }}
                >
                  {stat.value}
                </Typography>

                {/* 标签 */}
                <Typography
                  variant="h5"
                  sx={{
                    fontSize: { xs: '1.1rem', md: '1.25rem' },
                    color: 'rgba(255, 255, 255, 0.9)',
                    fontWeight: 300,
                    mb: 1,
                    textShadow: '0 2px 4px rgba(0,0,0,0.2)'
                  }}
                >
                  {stat.label}
                </Typography>

                {/* 描述 */}
                <Typography
                  variant="body2"
                  sx={{
                    fontSize: '0.9rem',
                    color: 'rgba(255, 255, 255, 0.7)',
                    fontWeight: 300
                  }}
                >
                  {stat.description}
                </Typography>

                {/* 装饰线 */}
                <Box
                  sx={{
                    width: 40,
                    height: 2,
                    background: gradients.primary,
                    mx: 'auto',
                    mt: 3,
                    borderRadius: 1,
                    ...animationStyles.pulseGlow
                  }}
                />
              </Box>
            </Grid>
          ))}
        </Grid>

        {/* 底部信任徽章 */}
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            mt: { xs: 4, md: 6 },
            gap: { xs: 2, md: 4 },
            flexWrap: 'wrap'
          }}
        >
          {[
            '企业级安全认证',
            'ISO 27001 合规',
            'SOC 2 Type II',
            '数据加密传输'
          ].map((badge, index) => (
            <Box
              key={badge}
              sx={{
                display: 'flex',
                alignItems: 'center',
                gap: 1,
                px: 3,
                py: 1.5,
                borderRadius: '25px',
                background: 'rgba(255, 255, 255, 0.1)',
                backdropFilter: 'blur(10px)',
                border: '1px solid rgba(255, 255, 255, 0.2)',
                ...animationStyles.fadeIn,
                animationDelay: `${index * 0.1 + 0.5}s`,
                transition: 'all 0.3s ease',
                '&:hover': {
                  background: 'rgba(255, 255, 255, 0.15)',
                  transform: 'translateY(-2px)'
                }
              }}
            >
              <Box
                sx={{
                  width: 6,
                  height: 6,
                  borderRadius: '50%',
                  background: gradients.primary,
                  ...animationStyles.pulseGlow
                }}
              />
              <Typography
                variant="body2"
                sx={{
                  color: 'rgba(255, 255, 255, 0.9)',
                  fontSize: '0.85rem',
                  fontWeight: 500
                }}
              >
                {badge}
              </Typography>
            </Box>
          ))}
        </Box>
      </Container>
    </Box>
  );
};

export default StatsSection;
