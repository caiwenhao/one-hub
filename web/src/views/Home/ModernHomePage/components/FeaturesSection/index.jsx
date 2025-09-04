import React from 'react';
import {
  Box,
  Container,
  Typography,
  Grid,
  Card,
  CardContent,
  useTheme,
  useMediaQuery
} from '@mui/material';
import {
  Security,
  Calculate,
  Favorite,
  Code
} from '@mui/icons-material';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGradientText } from '../../styles/theme';

const FeaturesSection = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));

  const features = [
    {
      icon: Security,
      title: '100%官方正源',
      description: '杜绝任何逆向和共享号池。我们只提供来自官方的纯净API，确保您的应用性能和数据安全，从源头避免模型"降智"。',
      gradient: gradients.featureAccent,
      iconGradient: gradients.iconAccent,
      iconColor: colors.accent
    },
    {
      icon: Calculate,
      title: '绝对计费透明',
      description: 'Token计算方式与官方严格一致。提供极具竞争力的充值折扣，无任何隐藏费用，让您的每一分投入都清晰可见。',
      gradient: gradients.featurePurple,
      iconGradient: gradients.iconPurple,
      iconColor: colors.purple
    },
    {
      icon: Favorite,
      title: '企业级稳定性',
      description: '采用高可用服务架构和智能路由技术，承诺99.9%的SLA，并提供实时服务状态监控，为您的生产环境保驾护航。',
      gradient: gradients.featurePink,
      iconGradient: gradients.iconPink,
      iconColor: colors.pink
    },
    {
      icon: Code,
      title: '开发者优先',
      description: '统一的API端点，3分钟即可轻松接入。提供清晰的文档、代码示例和专业的技术支持，让您专注于创造。',
      gradient: gradients.featureOrange,
      iconGradient: gradients.iconOrange,
      iconColor: colors.orange
    }
  ];

  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        backgroundColor: '#ffffff',
        py: { xs: 4, md: 8 },
        overflow: 'hidden'
      }}
    >
      {/* 背景渐变 */}
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          background: 'linear-gradient(to bottom right, rgba(249, 250, 251, 0.5), rgba(59, 130, 246, 0.03))'
        }}
      />

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1 }}>
        {/* 标题区域 */}
        <Box
          sx={{
            textAlign: 'center',
            mb: { xs: 3, md: 5 },
            ...animationStyles.fadeIn
          }}
        >
          <Typography
            variant="h2"
            sx={{
              fontSize: { xs: '2.5rem', md: '3.5rem', lg: '4rem' },
              fontWeight: 200,
              color: colors.primary,
              mb: 3,
              lineHeight: 1.2,
              letterSpacing: '-0.01em',
              textShadow: '0 2px 4px rgba(0,0,0,0.1)',
              '& .gradient-text': {
                fontWeight: 'bold',
                ...createGradientText()
              }
            }}
          >
            为什么开发者选择 <span className="gradient-text">Kapon AI</span>
          </Typography>
          <Typography
            variant="h5"
            sx={{
              fontSize: { xs: '1.1rem', md: '1.25rem' },
              color: colors.secondary,
              maxWidth: '600px',
              mx: 'auto',
              fontWeight: 300,
              lineHeight: 1.6
            }}
          >
            体验前所未有的稳定性、透明度和专业服务
          </Typography>
        </Box>

        {/* 特性卡片网格 */}
        <Grid container spacing={{ xs: 3, md: 4 }}>
          {features.map((feature, index) => {
            const IconComponent = feature.icon;
            return (
              <Grid item xs={12} sm={6} lg={3} key={feature.title}>
                <Card
                  sx={{
                    height: '100%',
                    background: gradients.card,
                    border: '1px solid rgba(0, 0, 0, 0.05)',
                    borderRadius: '24px',
                    p: { xs: 2.5, md: 3 },
                    position: 'relative',
                    overflow: 'hidden',
                    ...animationStyles.hoverLift,
                    ...animationStyles.fadeIn,
                    animationDelay: `${index * 0.1}s`,
                    transition: 'all 0.5s cubic-bezier(0.4, 0, 0.2, 1)',
                    '&:hover': {
                      borderColor: 'rgba(66, 153, 225, 0.3)',
                      transform: 'translateY(-8px) scale(1.02)',
                      boxShadow: '0 20px 40px rgba(0,0,0,0.12)',
                      '& .feature-overlay': {
                        opacity: 1
                      },
                      '& .feature-icon': {
                        transform: 'scale(1.1)'
                      }
                    }
                  }}
                >
                  {/* 悬浮时的背景渐变 */}
                  <Box
                    className="feature-overlay"
                    sx={{
                      position: 'absolute',
                      inset: 0,
                      background: feature.gradient,
                      opacity: 0,
                      transition: 'opacity 0.5s ease'
                    }}
                  />

                  <CardContent sx={{ p: 0, position: 'relative', zIndex: 1 }}>
                    {/* 图标 */}
                    <Box
                      className="feature-icon"
                      sx={{
                        width: 80,
                        height: 80,
                        borderRadius: '24px',
                        background: feature.iconGradient,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        mb: 4,
                        transition: 'all 0.3s ease',
                        ...animationStyles.pulseGlow
                      }}
                    >
                      <IconComponent
                        sx={{
                          fontSize: '2rem',
                          color: feature.iconColor
                        }}
                      />
                    </Box>

                    {/* 标题 */}
                    <Typography
                      variant="h4"
                      sx={{
                        fontSize: '1.5rem',
                        fontWeight: 600,
                        color: colors.primary,
                        mb: 3,
                        textShadow: '0 2px 4px rgba(0,0,0,0.1)'
                      }}
                    >
                      {feature.title}
                    </Typography>

                    {/* 描述 */}
                    <Typography
                      variant="body1"
                      sx={{
                        color: colors.secondary,
                        lineHeight: 1.6,
                        fontSize: '1.125rem',
                        fontWeight: 300
                      }}
                    >
                      {feature.description}
                    </Typography>
                  </CardContent>
                </Card>
              </Grid>
            );
          })}
        </Grid>

        {/* 底部装饰 */}
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            mt: { xs: 4, md: 6 },
            opacity: 0.6
          }}
        >
          <Box
            sx={{
              display: 'flex',
              gap: 2,
              alignItems: 'center'
            }}
          >
            {[...Array(3)].map((_, index) => (
              <Box
                key={index}
                sx={{
                  width: 8,
                  height: 8,
                  borderRadius: '50%',
                  background: gradients.primary,
                  ...animationStyles.pulseGlow,
                  animationDelay: `${index * 0.2}s`
                }}
              />
            ))}
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

export default FeaturesSection;
