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

  const codeExample = `import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: "your-kapon-api-key",
  baseURL: "https://api.kapon.ai/v1"
});

const response = await client.chat.completions.create({
  model: "gpt-4o",
  messages: [
    {"role": "user", "content": "Hello!"}
  ]
});

console.log(response.choices[0].message.content);`;

  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        background: 'linear-gradient(to bottom right, #ffffff 0%, rgba(59, 130, 246, 0.02) 30%, rgba(139, 92, 246, 0.02) 70%, #ffffff 100%)',
        py: { xs: 12, md: 16 },
        overflow: 'hidden'
      }}
    >
      {/* 浮动背景元素 */}
      <Box
        sx={{
          position: 'absolute',
          top: '33%',
          right: '25%',
          width: '288px',
          height: '288px',
          background: 'linear-gradient(to right, rgba(66, 153, 225, 0.08), rgba(139, 92, 246, 0.08))',
          borderRadius: '50%',
          filter: 'blur(60px)',
          ...animationStyles.floating
        }}
      />

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1 }}>
        <Grid container spacing={{ xs: 4, md: 8 }} alignItems="center">
          {/* 代码示例 */}
          <Grid item xs={12} lg={6}>
            <Box
              sx={{
                background: '#1a202c',
                borderRadius: '24px',
                p: { xs: 3, md: 5 },
                color: '#ffffff',
                boxShadow: '0 20px 40px rgba(0,0,0,0.15)',
                border: '1px solid #2d3748',
                position: 'relative',
                overflow: 'hidden',
                ...animationStyles.hoverLift,
                ...animationStyles.pulseGlow
              }}
            >
              {/* 背景渐变 */}
              <Box
                sx={{
                  position: 'absolute',
                  inset: 0,
                  background: 'linear-gradient(to bottom right, rgba(66, 153, 225, 0.05), transparent)'
                }}
              />

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
                  main.js
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
                  '& .keyword': { color: '#8b5cf6' },
                  '& .string': { color: colors.accent, fontWeight: 600 },
                  '& .comment': { color: '#6b7280' }
                }}
              >
                <code>
                  {codeExample.split('\n').map((line, index) => (
                    <div key={index} style={{ minHeight: '1.5em' }}>
                      {line
                        .replace(/import|const|await|new/g, '<span class="keyword">$&</span>')
                        .replace(/"[^"]*"/g, '<span class="string">$&</span>')
                        .split('<span').map((part, i) => {
                          if (part.startsWith(' class="keyword"')) {
                            return <span key={i} className="keyword" dangerouslySetInnerHTML={{ __html: part.replace(' class="keyword">', '') }} />;
                          } else if (part.startsWith(' class="string"')) {
                            return <span key={i} className="string" dangerouslySetInnerHTML={{ __html: part.replace(' class="string">', '') }} />;
                          }
                          return part;
                        })}
                    </div>
                  ))}
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
                  mb: 4,
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
                  onClick={() => navigate('/playground')}
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
