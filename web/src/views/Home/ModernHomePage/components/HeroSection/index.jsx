import React from 'react';
import {
  Box,
  Container,
  Typography,
  Button,
  Stack,
  useTheme,
  useMediaQuery
} from '@mui/material';
import { Rocket, ArrowForward } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGradientText, createFloatingElement } from '../../styles/theme';

const HeroSection = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const navigate = useNavigate();

  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        background: gradients.heroBackground,
        pt: { xs: 8, md: 12 },
        pb: { xs: 8, md: 12 },
        minHeight: '700px',
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
          ...createFloatingElement('12px', colors.accentAlpha[30]),
          top: '20%',
          left: '10%',
          ...animationStyles.floating,
          ...animationStyles.pulseGlow
        }}
      />
      <Box
        sx={{
          ...createFloatingElement('8px', colors.purpleAlpha[40]),
          top: '40%',
          right: '20%',
          ...animationStyles.floating,
          animationDelay: '0.3s',
          ...animationStyles.pulseGlow
        }}
      />
      <Box
        sx={{
          ...createFloatingElement('10px', colors.pinkAlpha[35]),
          bottom: '32%',
          left: '25%',
          ...animationStyles.floating,
          animationDelay: '0.7s',
          ...animationStyles.pulseGlow
        }}
      />
      <Box
        sx={{
          ...createFloatingElement('6px', colors.orangeAlpha[40]),
          top: '60%',
          right: '33%',
          ...animationStyles.floating,
          animationDelay: '0.5s',
          ...animationStyles.pulseGlow
        }}
      />
      <Box
        sx={{
          ...createFloatingElement('14px', colors.accentAlpha[25]),
          bottom: '40%',
          right: '10%',
          ...animationStyles.floating,
          animationDelay: '1s',
          ...animationStyles.pulseGlow
        }}
      />

      {/* 大型模糊背景元素 */}
      <Box
        sx={{
          position: 'absolute',
          top: '25%',
          left: '25%',
          width: '384px',
          height: '384px',
          background: gradients.floatingBlur1,
          borderRadius: '50%',
          filter: 'blur(60px)',
          ...animationStyles.floating
        }}
      />
      <Box
        sx={{
          position: 'absolute',
          bottom: '25%',
          right: '25%',
          width: '320px',
          height: '320px',
          background: gradients.floatingBlur2,
          borderRadius: '50%',
          filter: 'blur(60px)',
          ...animationStyles.floating,
          animationDelay: '0.5s'
        }}
      />

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 10, maxWidth: '1200px' }}>
        <Box
          sx={{
            textAlign: 'center',
            maxWidth: '900px',
            mx: 'auto',
            ...animationStyles.fadeIn
          }}
        >
          {/* 主标题 */}
          <Typography
            variant="h1"
            sx={{
              fontSize: { xs: '2.5rem', md: '4rem', lg: '4.5rem' },
              fontWeight: 200,
              color: colors.primary,
              mb: 3,
              lineHeight: 1.2,
              letterSpacing: '-0.02em',
              textShadow: '0 2px 4px rgba(0,0,0,0.1)',
              '& .gradient-text': {
                fontWeight: 'bold',
                ...createGradientText(),
                ...animationStyles.gradientShift
              }
            }}
          >
            稳定，是AI应用的
            <br />
            <span className="gradient-text">唯一标准</span>
          </Typography>

          {/* 副标题 */}
          <Typography
            variant="h4"
            sx={{
              fontSize: { xs: '1.2rem', md: '1.4rem', lg: '1.75rem' },
              color: colors.secondary,
              mb: 5,
              maxWidth: '850px',
              mx: 'auto',
              lineHeight: 1.4,
              fontWeight: 300,
              textShadow: '0 2px 4px rgba(0,0,0,0.1)'
            }}
          >
            Kapon AI 为中小企业与开发者提供100%官方正源大模型API
            <br />
            告别逆向和号池，享受企业级稳定、计费透明的AI调用服务
          </Typography>

          {/* CTA按钮组 */}
          <Stack
            direction={{ xs: 'column', sm: 'row' }}
            spacing={4}
            justifyContent="center"
            alignItems="center"
            sx={{ mb: 4 }}
          >
            <Button
              variant="contained"
              size="large"
              startIcon={<Rocket />}
              onClick={() => navigate('/register')}
              sx={{
                fontSize: { xs: '1.1rem', md: '1.25rem' },
                px: { xs: 4, md: 6 },
                py: { xs: 1.5, md: 2 },
                borderRadius: '50px',
                background: gradients.primary,
                fontWeight: 600,
                boxShadow: '0 8px 32px rgba(66, 153, 225, 0.3)',
                ...animationStyles.pulseGlow,
                transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
                '&:hover': {
                  background: gradients.primary,
                  transform: 'translateY(-4px) scale(1.05)',
                  boxShadow: '0 12px 40px rgba(66, 153, 225, 0.4)'
                }
              }}
            >
              注册，获取免费试用额度
            </Button>

            <Button
              variant="text"
              endIcon={<ArrowForward />}
              onClick={() => navigate('/about')}
              sx={{
                fontSize: { xs: '1rem', md: '1.125rem' },
                color: colors.primary,
                fontWeight: 500,
                px: 2,
                py: 1,
                borderRadius: '25px',
                transition: 'all 0.3s ease',
                '&:hover': {
                  color: colors.accent,
                  backgroundColor: colors.accentAlpha[5],
                  '& .MuiSvgIcon-root': {
                    transform: 'translateX(4px)'
                  }
                },
                '& .MuiSvgIcon-root': {
                  transition: 'transform 0.3s ease'
                }
              }}
            >
              企业客户？联系我们
            </Button>
          </Stack>

          {/* 信任指标 */}
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
              flexWrap: 'wrap',
              gap: { xs: 2, md: 4 },
              mt: 6,
              opacity: 0.8
            }}
          >
            <Typography
              variant="body2"
              sx={{
                color: colors.secondary,
                fontSize: '0.9rem',
                fontWeight: 500,
                display: 'flex',
                alignItems: 'center',
                gap: 1
              }}
            >
              <Box
                sx={{
                  width: 8,
                  height: 8,
                  borderRadius: '50%',
                  background: gradients.primary,
                  ...animationStyles.pulseGlow
                }}
              />
              99.9% 服务可用性
            </Typography>
            <Typography
              variant="body2"
              sx={{
                color: colors.secondary,
                fontSize: '0.9rem',
                fontWeight: 500,
                display: 'flex',
                alignItems: 'center',
                gap: 1
              }}
            >
              <Box
                sx={{
                  width: 8,
                  height: 8,
                  borderRadius: '50%',
                  background: gradients.primary,
                  ...animationStyles.pulseGlow
                }}
              />
              10K+ 开发者信赖
            </Typography>
            <Typography
              variant="body2"
              sx={{
                color: colors.secondary,
                fontSize: '0.9rem',
                fontWeight: 500,
                display: 'flex',
                alignItems: 'center',
                gap: 1
              }}
            >
              <Box
                sx={{
                  width: 8,
                  height: 8,
                  borderRadius: '50%',
                  background: gradients.primary,
                  ...animationStyles.pulseGlow
                }}
              />
              24/7 技术支持
            </Typography>
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

export default HeroSection;
