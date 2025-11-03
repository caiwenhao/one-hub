import React from 'react';
import { Box, Container, Typography, Grid, Chip } from '@mui/material';
import { Icon } from '@iconify/react';
import ModelCard from './ModelCard';
// 方案B：使用静态内置数据渲染“明星模型”，无外部依赖
const mockModels = {
  featured: [
    {
      id: 'gpt-5',
      name: 'GPT-5',
      provider: 'OpenAI',
      abbr: 'O',
      description:
        '新一代通用模型，多模态与复杂推理更强，适合复杂编程任务与企业级集成。',
      useCases: ['复杂编程任务', '前端原型开发', '多步推理', '企业级 API 集成'],
      iconColor: '#3B82F6',
      tag: { type: 'hot', label: '🔥 火爆' }
    },
    {
      id: 'claude-sonnet-4',
      name: 'Claude Sonnet 4.5',
      provider: 'Anthropic',
      abbr: 'A',
      description:
        '面向自主智能体与长时工具调度优化，代码能力强，支持更长输出。',
      useCases: ['全栈开发', '漏洞修补', '自动化测试', '长时运行智能体'],
      iconColor: '#F59E0B',
      tag: { type: 'hot', label: '🔥 火爆' }
    },
    {
      id: 'gemini-2-5-pro',
      name: 'Gemini 2.5 Pro',
      provider: 'Google',
      abbr: 'G',
      description:
        '原生多模态与百万级上下文，长文档与代码理解优秀，适配复杂业务场景。',
      useCases: ['Web 应用开发', '大规模代码分析', '视频理解', '企业文档处理'],
      iconColor: '#6366F1',
      tag: { type: 'recommended', label: '⭐ 推荐' }
    }
  ]
};
import { colors, gradients, animationStyles, createGradientText } from '../styles/theme';

const FeaturedModels = () => {
  return (
    <Box
      component="section"
      sx={{
        backgroundColor: '#ffffff',
        py: { xs: 8, md: 16 },
        position: 'relative',
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

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1, maxWidth: '1200px' }}>
        {/* 标题区域 */}
        <Box
          sx={{
            textAlign: 'center',
            mb: { xs: 6, md: 10 },
            ...animationStyles.fadeIn
          }}
        >
          {/* 标签 */}
          <Box sx={{ mb: 3 }}>
            <Chip
              icon={<Icon icon="solar:fire-bold" width={20} height={20} />}
              label="最受欢迎"
              sx={{
                background: 'linear-gradient(135deg, rgba(66, 153, 225, 0.1), rgba(139, 92, 246, 0.1))',
                color: colors.accent,
                fontWeight: 600,
                fontSize: '1rem',
                px: 3,
                py: 2,
                borderRadius: '50px',
                border: 'none',
                '& .MuiChip-icon': {
                  color: colors.accent
                }
              }}
            />
          </Box>

          {/* 主标题 */}
          <Typography
            variant="h2"
            sx={{
              fontSize: { xs: '2.5rem', md: '3.5rem', lg: '4rem' },
              fontWeight: 200,
              color: colors.primary,
              mb: 2,
              lineHeight: 1.1,
              letterSpacing: '-0.02em',
              textShadow: 'none',
              '& .gradient-text': {
                fontWeight: 'bold',
                ...createGradientText('linear-gradient(45deg, #22D3EE 0%, #0EA5FF 50%, #8B5CF6 100%)')
              }
            }}
          >
            <span className="gradient-text">明星模型</span> 荟萃
          </Typography>

          {/* 副标题 */}
          <Typography
            variant="h6"
            sx={{
              fontSize: { xs: '1rem', md: '1.125rem' },
              color: colors.secondary,
              maxWidth: '600px',
              mx: 'auto',
              fontWeight: 300,
              lineHeight: 1.6
            }}
          >
            开发者首选，经过市场验证的顶级AI模型
          </Typography>
        </Box>

        {/* 模型卡片网格 */}
        <Grid container spacing={{ xs: 3, md: 4 }}>
          {mockModels.featured.map((model, index) => (
            <Grid item xs={12} sm={6} lg={4} key={model.id}>
              <Box
                sx={{
                  ...animationStyles.fadeIn,
                  animationDelay: `${index * 0.2}s`
                }}
              >
                <ModelCard model={model} variant="featured" />
              </Box>
            </Grid>
          ))}
        </Grid>

        {/* 底部装饰 */}
        <Box
          sx={{
            mt: { xs: 6, md: 10 },
            textAlign: 'center'
          }}
        >
          <Typography
            variant="body1"
            sx={{
              color: colors.secondary,
              fontSize: '0.875rem',
              fontWeight: 300,
              opacity: 0.8
            }}
          >
            更多精彩模型，等你探索
          </Typography>
        </Box>
      </Container>
    </Box>
  );
};

export default FeaturedModels;
