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
        backgroundColor: '#ffffff', // 使用白色背景与热门模型页面一致
        pt: { xs: 8, md: 12 },
        pb: { xs: 8, md: 12 },
        minHeight: '700px',
        display: 'flex',
        alignItems: 'center',
        position: 'relative',
        overflow: 'hidden'
      }}
    >
      {/* 背景渐变 - 与热门模型页面一致的轻微渐变 */}
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          background: 'linear-gradient(to bottom right, rgba(249, 250, 251, 0.5), rgba(14, 165, 255, 0.06))'
        }}
      />
      {/* 轻量网格覆盖层（AI 科技风） */}
      <Box className="ai-grid-overlay" sx={{ opacity: 0.35 }} />



      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 10, maxWidth: '1200px' }}>
        <Box
          sx={{
            textAlign: 'center',
            maxWidth: '900px',
            mx: 'auto'
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
              '& .gradient-text': {
                fontWeight: 'bold',
                color: '#0EA5FF' // 使用纯色替代渐变文字（AI 科技主色）
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
            grouplay AI 为中小企业与开发者提供100%官方正源大模型API
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
                background: '#0EA5FF', // 使用纯色背景替代渐变（AI 科技主色）
                fontWeight: 600,
                boxShadow: 'none', // 移除阴影
                transition: 'background-color 0.2s ease',
                '&:hover': {
                  background: '#0369A1' // hover时使用稍深的颜色（主色深色）
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
              component="div"
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
              component="div"
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
              component="div"
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
