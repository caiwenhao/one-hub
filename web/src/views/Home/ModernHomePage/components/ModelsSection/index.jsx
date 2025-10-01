import React from 'react';
import {
  Box,
  Container,
  Typography,
  Grid,
  Card,
  CardContent,
  Chip,
  Button,
  useTheme,
  useMediaQuery
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGradientText } from '../../styles/theme';

const ModelsSection = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const navigate = useNavigate();

  const models = [
    {
      name: 'GPT-5',
      provider: 'OpenAI',
      description: '最新一代旗舰模型，突破性的推理能力与创造力',
      tag: '🔥 火爆',
      tagColor: 'linear-gradient(to right, #ef4444, #ec4899)',
      iconBg: gradients.modelGPT,
      iconText: 'GPT'
    },
    {
      name: 'Gemini-2.5-Pro',
      provider: 'Google',
      description: 'Google最强大模型，卓越的多模态理解能力',
      tag: '🔥 火爆',
      tagColor: 'linear-gradient(to right, #ef4444, #ec4899)',
      iconBg: gradients.modelGemini,
      iconText: 'GEM'
    },
    {
      name: 'Claude Sonnet 4',
      provider: 'Anthropic',
      description: 'Anthropic顶级模型，安全可靠的AI助手',
      tag: '🔥 火爆',
      tagColor: 'linear-gradient(to right, #ef4444, #ec4899)',
      iconBg: gradients.modelClaude,
      iconText: 'CLU'
    }
  ];

  return (
    <Box
      component="section"
      sx={{
        backgroundColor: '#ffffff',
        py: { xs: 4, md: 8 },
        position: 'relative'
      }}
    >
      <Container maxWidth="lg">
        {/* 标题区域 */}
        <Box
          sx={{
            textAlign: 'center',
            mb: { xs: 4, md: 6 }
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
              '& .gradient-text': {
                fontWeight: 'bold',
                color: '#0EA5FF' // 使用纯色替代渐变文字（AI 科技主色）
              }
            }}
          >
            连接全球 <span className="gradient-text">顶尖AI大脑</span>
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
            实时同步全球最前沿、性能最卓越的大语言模型
          </Typography>
        </Box>

        {/* 模型卡片网格 */}
        <Grid container spacing={{ xs: 3, md: 4 }} sx={{ mb: { xs: 4, md: 6 } }}>
          {models.map((model, index) => (
            <Grid item xs={12} md={4} key={model.name}>
              <Card
                sx={{
                  height: '100%',
                  background: gradients.card,
                  border: '1px solid rgba(0, 0, 0, 0.05)',
                  borderRadius: '24px',
                  p: { xs: 3, md: 4 },
                  position: 'relative',
                  overflow: 'hidden',
                  ...animationStyles.hoverLift,
                  ...animationStyles.fadeIn,
                  animationDelay: `${index * 0.2}s`,
                  transition: 'all 0.5s cubic-bezier(0.4, 0, 0.2, 1)',
                  '&:hover': {
                    borderColor: 'rgba(66, 153, 225, 0.3)',
                    transform: 'translateY(-8px) scale(1.02)',
                    boxShadow: '0 20px 40px rgba(0,0,0,0.12)'
                  }
                }}
              >
                {/* 标签 */}
                <Box
                  sx={{
                    position: 'absolute',
                    top: 16,
                    right: 16,
                    zIndex: 2
                  }}
                >
                  <Chip
                    label={model.tag}
                    sx={{
                      background: model.tagColor,
                      color: '#ffffff',
                      fontWeight: 500,
                      fontSize: '0.75rem',
                      height: 28,
                      '& .MuiChip-label': {
                        px: 1.5
                      }
                    }}
                  />
                </Box>

                <CardContent sx={{ p: 0, position: 'relative', zIndex: 1 }}>
                  {/* 图标 */}
                  <Box
                    sx={{
                      width: 64,
                      height: 64,
                      borderRadius: '16px',
                      background: model.iconBg,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      mb: 3,
                      color: '#ffffff',
                      fontWeight: 'bold',
                      fontSize: '1.125rem',
                      boxShadow: '0 4px 15px rgba(0,0,0,0.1)'
                    }}
                  >
                    {model.iconText}
                  </Box>

                  {/* 模型名称 */}
                  <Typography
                    variant="h4"
                    sx={{
                      fontSize: '1.5rem',
                      fontWeight: 600,
                      color: colors.primary,
                      mb: 1
                    }}
                  >
                    {model.name}
                  </Typography>

                  {/* 提供商 */}
                  <Typography
                    variant="body2"
                    sx={{
                      color: colors.accent,
                      fontWeight: 500,
                      mb: 2,
                      fontSize: '0.9rem'
                    }}
                  >
                    {model.provider}
                  </Typography>

                  {/* 描述 */}
                  <Typography
                    variant="body1"
                    sx={{
                      color: colors.secondary,
                      lineHeight: 1.6,
                      fontWeight: 300
                    }}
                  >
                    {model.description}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>

        {/* 查看更多按钮 */}
        <Box
          sx={{
            textAlign: 'center'
          }}
        >
          <Button
            variant="contained"
            size="large"
            onClick={() => navigate('/price')}
            sx={{
              fontSize: '1.125rem',
              px: 5,
              py: 2,
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
            查看所有模型定价 →
          </Button>
        </Box>

        {/* 底部特性说明 */}
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
            '实时同步最新模型',
            '统一API接口',
            '透明计费标准',
            '7x24小时可用'
          ].map((feature, index) => (
            <Box
              key={feature}
              sx={{
                display: 'flex',
                alignItems: 'center',
                gap: 1,
                px: 3,
                py: 1.5,
                borderRadius: '25px',
                background: 'rgba(66, 153, 225, 0.05)',
                border: '1px solid rgba(66, 153, 225, 0.1)',
                transition: 'background-color 0.2s ease',
                '&:hover': {
                  background: 'rgba(66, 153, 225, 0.1)'
                }
              }}
            >
              <Box
                sx={{
                  width: 6,
                  height: 6,
                  borderRadius: '50%',
                  background: '#0EA5FF' // 使用纯色替代渐变（AI 科技主色）
                }}
              />
              <Typography
                variant="body2"
                sx={{
                  color: colors.accent,
                  fontSize: '0.9rem',
                  fontWeight: 500
                }}
              >
                {feature}
              </Typography>
            </Box>
          ))}
        </Box>
      </Container>
    </Box>
  );
};

export default ModelsSection;
