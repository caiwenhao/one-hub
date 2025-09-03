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
import { ArrowForward } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGradientText, createFloatingElement } from '../../styles/theme';

const CTASection = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const navigate = useNavigate();

  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        background: gradients.finalCTA,
        py: { xs: 8, md: 16 },
        overflow: 'hidden'
      }}
    >
      {/* 背景渐变层 */}
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          background: gradients.finalCTAOverlay
        }}
      />

      {/* 浮动装饰元素 */}
      <Box
        sx={{
          ...createFloatingElement('12px', colors.accentAlpha[40]),
          top: '20%',
          left: '25%',
          ...animationStyles.floating,
          ...animationStyles.pulseGlow
        }}
      />
      <Box
        sx={{
          ...createFloatingElement('8px', colors.purpleAlpha[40]),
          bottom: '32%',
          right: '33%',
          ...animationStyles.floating,
          animationDelay: '0.5s',
          ...animationStyles.pulseGlow
        }}
      />
      <Box
        sx={{
          ...createFloatingElement('10px', colors.pinkAlpha[35]),
          top: '50%',
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
          top: '33%',
          left: '33%',
          width: '384px',
          height: '384px',
          background: 'linear-gradient(to right, rgba(66, 153, 225, 0.1), rgba(139, 92, 246, 0.1))',
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
          background: 'linear-gradient(to right, rgba(236, 72, 153, 0.1), rgba(245, 158, 11, 0.1))',
          borderRadius: '50%',
          filter: 'blur(60px)',
          ...animationStyles.floating,
          animationDelay: '0.7s'
        }}
      />

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 10, maxWidth: '1200px' }}>
        <Box
          sx={{
            textAlign: 'center',
            maxWidth: '800px',
            mx: 'auto',
            ...animationStyles.fadeIn
          }}
        >
          {/* 主标题 */}
          <Typography
            variant="h1"
            sx={{
              fontSize: { xs: '3rem', md: '4.5rem', lg: '5.5rem' },
              fontWeight: 200,
              color: '#ffffff',
              mb: 4,
              lineHeight: 1.1,
              letterSpacing: '-0.02em',
              textShadow: '0 4px 8px rgba(0,0,0,0.3)',
              '& .gradient-text': {
                fontWeight: 'bold',
                background: 'linear-gradient(to right, #63B3ED, #8B5CF6)',
                WebkitBackgroundClip: 'text',
                WebkitTextFillColor: 'transparent',
                backgroundClip: 'text'
              }
            }}
          >
            准备好构建更<span className="gradient-text">可靠</span>的AI应用了吗？
          </Typography>

          {/* 副标题 */}
          <Typography
            variant="h4"
            sx={{
              fontSize: { xs: '1.25rem', md: '1.5rem', lg: '2rem' },
              color: 'rgba(255, 255, 255, 0.9)',
              mb: 8,
              maxWidth: '600px',
              mx: 'auto',
              lineHeight: 1.5,
              fontWeight: 300,
              textShadow: '0 2px 4px rgba(0,0,0,0.2)'
            }}
          >
            立即注册，体验Kapon AI带来的极致稳定与便捷
          </Typography>

          {/* CTA按钮组 */}
          <Stack
            direction={{ xs: 'column', sm: 'row' }}
            spacing={4}
            justifyContent="center"
            alignItems="center"
            sx={{ mb: 6 }}
          >
            <Button
              variant="contained"
              size="large"
              onClick={() => navigate('/register')}
              sx={{
                fontSize: { xs: '1.25rem', md: '1.5rem' },
                px: { xs: 6, md: 8 },
                py: { xs: 2, md: 2.5 },
                borderRadius: '50px',
                background: gradients.primary,
                fontWeight: 'bold',
                boxShadow: '0 12px 40px rgba(66, 153, 225, 0.4)',
                ...animationStyles.pulseGlow,
                transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
                '&:hover': {
                  background: gradients.primary,
                  transform: 'translateY(-4px) scale(1.05)',
                  boxShadow: '0 16px 50px rgba(66, 153, 225, 0.5)'
                }
              }}
            >
              立即开始免费试用
            </Button>

            <Button
              variant="text"
              endIcon={<ArrowForward />}
              onClick={() => navigate('/about')}
              sx={{
                fontSize: { xs: '1rem', md: '1.125rem' },
                color: 'rgba(255, 255, 255, 0.9)',
                fontWeight: 500,
                px: 2,
                py: 1,
                borderRadius: '25px',
                transition: 'all 0.3s ease',
                '&:hover': {
                  color: colors.accentLight,
                  backgroundColor: 'rgba(255, 255, 255, 0.1)',
                  '& .MuiSvgIcon-root': {
                    transform: 'translateX(4px)'
                  }
                },
                '& .MuiSvgIcon-root': {
                  transition: 'transform 0.3s ease'
                }
              }}
            >
              或了解企业方案
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
              opacity: 0.8
            }}
          >
            <Typography
              variant="body2"
              sx={{
                color: 'rgba(255, 255, 255, 0.8)',
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
              无需信用卡
            </Typography>
            <Typography
              variant="body2"
              sx={{
                color: 'rgba(255, 255, 255, 0.8)',
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
              30秒快速注册
            </Typography>
            <Typography
              variant="body2"
              sx={{
                color: 'rgba(255, 255, 255, 0.8)',
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
              立即获得试用额度
            </Typography>
          </Box>

          {/* 底部装饰 */}
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center',
              mt: 8,
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
              {[...Array(5)].map((_, index) => (
                <Box
                  key={index}
                  sx={{
                    width: index === 2 ? 12 : 8,
                    height: index === 2 ? 12 : 8,
                    borderRadius: '50%',
                    background: 'rgba(255, 255, 255, 0.3)',
                    ...animationStyles.pulseGlow,
                    animationDelay: `${index * 0.2}s`
                  }}
                />
              ))}
            </Box>
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

export default CTASection;
