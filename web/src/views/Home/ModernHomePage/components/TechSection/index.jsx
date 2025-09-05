import React from 'react';
import {
  Box,
  Container,
  Typography,
  Grid,
  Button,
  useTheme,
  useMediaQuery
} from '@mui/material';
import { ArrowForward, Launch } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGradientText, createFloatingElement } from '../../styles/theme';

const TechSection = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const navigate = useNavigate();

  // 带语法高亮的代码示例
  const renderCodeWithHighlight = () => {
    return (
      <>
        <span style={{ color: '#ff7b72' }}>from</span>{' '}
        <span style={{ color: '#79c0ff' }}>openai</span>{' '}
        <span style={{ color: '#ff7b72' }}>import</span>{' '}
        <span style={{ color: '#79c0ff' }}>OpenAI</span>
        {'\n\n'}
        <span style={{ color: '#f85149' }}>client</span>{' '}
        <span style={{ color: '#ff7b72' }}>=</span>{' '}
        <span style={{ color: '#79c0ff' }}>OpenAI</span>
        <span style={{ color: '#e6edf3' }}>(</span>
        {'\n    '}
        <span style={{ color: '#79c0ff' }}>api_key</span>
        <span style={{ color: '#ff7b72' }}>=</span>
        <span style={{ color: '#a5d6ff' }}>"your-kapon-api-key"</span>
        <span style={{ color: '#e6edf3' }}>,</span>
        {'\n    '}
        <span style={{ color: '#79c0ff' }}>base_url</span>
        <span style={{ color: '#ff7b72' }}>=</span>
        <span style={{ color: '#a5d6ff' }}>"https://models.kapon.cloud/v1"</span>
        {'\n'}
        <span style={{ color: '#e6edf3' }}>)</span>
        {'\n\n'}
        <span style={{ color: '#f85149' }}>response</span>{' '}
        <span style={{ color: '#ff7b72' }}>=</span>{' '}
        <span style={{ color: '#f85149' }}>client</span>
        <span style={{ color: '#e6edf3' }}>.</span>
        <span style={{ color: '#f85149' }}>chat</span>
        <span style={{ color: '#e6edf3' }}>.</span>
        <span style={{ color: '#f85149' }}>completions</span>
        <span style={{ color: '#e6edf3' }}>.</span>
        <span style={{ color: '#d2a8ff' }}>create</span>
        <span style={{ color: '#e6edf3' }}>(</span>
        {'\n    '}
        <span style={{ color: '#79c0ff' }}>model</span>
        <span style={{ color: '#ff7b72' }}>=</span>
        <span style={{ color: '#a5d6ff' }}>"gpt-4o"</span>
        <span style={{ color: '#e6edf3' }}>,</span>
        {'\n    '}
        <span style={{ color: '#79c0ff' }}>messages</span>
        <span style={{ color: '#ff7b72' }}>=</span>
        <span style={{ color: '#e6edf3' }}>[</span>
        {'\n        '}
        <span style={{ color: '#e6edf3' }}>{"{"}</span>
        <span style={{ color: '#a5d6ff' }}>"role"</span>
        <span style={{ color: '#e6edf3' }}>:</span>{' '}
        <span style={{ color: '#a5d6ff' }}>"user"</span>
        <span style={{ color: '#e6edf3' }}>,</span>{' '}
        <span style={{ color: '#a5d6ff' }}>"content"</span>
        <span style={{ color: '#e6edf3' }}>:</span>{' '}
        <span style={{ color: '#a5d6ff' }}>"Hello!"</span>
        <span style={{ color: '#e6edf3' }}>{"}"}</span>
        {'\n    '}
        <span style={{ color: '#e6edf3' }}>]</span>
        {'\n'}
        <span style={{ color: '#e6edf3' }}>)</span>
        {'\n\n'}
        <span style={{ color: '#d2a8ff' }}>print</span>
        <span style={{ color: '#e6edf3' }}>(</span>
        <span style={{ color: '#f85149' }}>response</span>
        <span style={{ color: '#e6edf3' }}>.</span>
        <span style={{ color: '#f85149' }}>choices</span>
        <span style={{ color: '#e6edf3' }}>[</span>
        <span style={{ color: '#79c0ff' }}>0</span>
        <span style={{ color: '#e6edf3' }}>].</span>
        <span style={{ color: '#f85149' }}>message</span>
        <span style={{ color: '#e6edf3' }}>.</span>
        <span style={{ color: '#f85149' }}>content</span>
        <span style={{ color: '#e6edf3' }}>)</span>
      </>
    );
  };

  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        background: 'linear-gradient(to bottom right, #ffffff 0%, rgba(59, 130, 246, 0.02) 30%, rgba(139, 92, 246, 0.02) 70%, #ffffff 100%)',
        py: { xs: 4, md: 8 },
        overflow: 'hidden'
      }}
    >


      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1, maxWidth: '1200px' }}>
        <Grid container spacing={{ xs: 4, md: 8 }} alignItems="center">
          {/* 代码示例 */}
          <Grid item xs={12} lg={6}>
            <Box
              sx={{
                background: '#1a202c',
                borderRadius: '24px',
                p: { xs: 3, md: 5 },
                color: '#ffffff',
                border: '1px solid #2d3748',
                position: 'relative',
                overflow: 'hidden'
              }}
            >


              {/* 终端头部 */}
              <Box
                sx={{
                  display: 'flex',
                  alignItems: 'center',
                  mb: 4,
                  position: 'relative',
                  zIndex: 1
                }}
              >
                <Box
                  sx={{
                    width: 16,
                    height: 16,
                    borderRadius: '50%',
                    backgroundColor: '#ef4444',
                    mr: 1.5,
                    ...animationStyles.pulseGlow
                  }}
                />
                <Box
                  sx={{
                    width: 16,
                    height: 16,
                    borderRadius: '50%',
                    backgroundColor: '#f59e0b',
                    mr: 1.5,
                    ...animationStyles.pulseGlow
                  }}
                />
                <Box
                  sx={{
                    width: 16,
                    height: 16,
                    borderRadius: '50%',
                    backgroundColor: '#10b981',
                    mr: 4,
                    ...animationStyles.pulseGlow
                  }}
                />
                <Typography
                  sx={{
                    color: 'rgba(255, 255, 255, 0.6)',
                    fontSize: '1.125rem',
                    fontWeight: 500
                  }}
                >
                  main.py
                </Typography>
              </Box>

              {/* 代码内容 */}
              <Box
                component="pre"
                sx={{
                  fontFamily: '"Fira Code", "Monaco", "Consolas", monospace',
                  fontSize: { xs: '0.85rem', md: '1rem' },
                  lineHeight: 1.6,
                  overflow: 'auto',
                  margin: 0,
                  position: 'relative',
                  zIndex: 1,
                  color: '#e2e8f0'
                }}
              >
                <code>
                  {renderCodeWithHighlight()}
                </code>
              </Box>
            </Box>
          </Grid>

          {/* 说明文字 */}
          <Grid item xs={12} lg={6}>
            <Box sx={{ ...animationStyles.fadeIn }}>
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
                <span className="gradient-text">3分钟</span>，完成您的首次调用
              </Typography>

              <Typography
                variant="h5"
                sx={{
                  fontSize: { xs: '1.1rem', md: '1.25rem' },
                  color: colors.secondary,
                  mb: 6,
                  lineHeight: 1.6,
                  fontWeight: 300
                }}
              >
                无需复杂的配置。只需替换API Key和Base URL，即可在您现有的代码中无缝切换至Kapon AI。
              </Typography>

              <Box
                sx={{
                  display: 'flex',
                  flexDirection: { xs: 'column', sm: 'row' },
                  gap: 3,
                  alignItems: { xs: 'stretch', sm: 'center' }
                }}
              >
                <Button
                  variant="text"
                  endIcon={<ArrowForward />}
                  onClick={() => navigate('/developer')}
                  sx={{
                    fontSize: '1.125rem',
                    color: colors.accent,
                    fontWeight: 600,
                    px: 0,
                    py: 1,
                    justifyContent: 'flex-start',
                    transition: 'all 0.3s ease',
                    '&:hover': {
                      color: colors.purple,
                      backgroundColor: 'transparent',
                      '& .MuiSvgIcon-root': {
                        transform: 'translateX(4px)'
                      }
                    },
                    '& .MuiSvgIcon-root': {
                      transition: 'transform 0.3s ease'
                    }
                  }}
                >
                  查看开发者文档
                </Button>

                <Button
                  variant="text"
                  endIcon={<Launch />}
                  onClick={() => window.open('https://status.onehub.ai', '_blank')}
                  sx={{
                    fontSize: '1.125rem',
                    color: colors.primary,
                    fontWeight: 500,
                    px: 0,
                    py: 1,
                    justifyContent: 'flex-start',
                    transition: 'all 0.3s ease',
                    '&:hover': {
                      color: colors.accent,
                      backgroundColor: 'transparent',
                      '& .MuiSvgIcon-root': {
                        transform: 'translateX(2px)'
                      }
                    },
                    '& .MuiSvgIcon-root': {
                      transition: 'transform 0.3s ease'
                    }
                  }}
                >
                  实时服务状态
                </Button>
              </Box>
            </Box>
          </Grid>
        </Grid>
      </Container>
    </Box>
  );
};

export default TechSection;
