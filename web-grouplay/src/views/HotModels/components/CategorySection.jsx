import React, { useState } from 'react';
import { Box, Container, Typography, Grid, Button } from '@mui/material';
import { Icon } from '@iconify/react';
import ModelCard from './ModelCard';
// 方案B：使用静态内置数据渲染“分类模型”，无外部依赖
const categories = [
  { id: 'text', name: '文本生成', icon: 'solar:document-text-bold' },
  { id: 'image', name: '图像生成', icon: 'solar:camera-bold' },
  { id: 'video', name: '视频生成', icon: 'solar:video-frame-play-vertical-bold' },
  { id: 'audio', name: '音频处理', icon: 'solar:microphone-bold' },
  { id: 'embedding', name: '向量嵌入', icon: 'solar:code-bold' }
];

const mockModels = {
  textModels: [
    {
      id: 'deepseek-v3-1',
      name: 'deepseek-v3.1',
      provider: 'DeepSeek',
      description: '支持百万级上下文，多步推理增强，覆盖 100+ 语言。',
      iconColor: '#0EA5E9',
      performance: { intelligence: 90, speed: 85, costEfficiency: 92, safety: 88 }
    },
    {
      id: 'kimi-k2',
      name: 'kimi-k2',
      provider: 'Moonshot AI',
      description: 'MoE 架构，面向工具调用与自主任务执行，性价比高。',
      iconColor: '#8B5CF6',
      performance: { intelligence: 86, speed: 88, costEfficiency: 90, safety: 86 }
    },
    {
      id: 'gpt-4o',
      name: 'GPT-4o',
      provider: 'OpenAI',
      description: '多模态旗舰，实时音视频/文本，通用能力强。',
      iconColor: '#16a34a',
      performance: { intelligence: 92, speed: 84, costEfficiency: 80, safety: 90 }
    }
  ],
  imageModels: [
    {
      id: 'gpt-image-1',
      name: 'GPT-Image-1',
      provider: 'OpenAI',
      description: '自回归图像生成，文本渲染与细粒度编辑强。',
      iconColor: '#22c55e',
      performance: { intelligence: 88, speed: 86, costEfficiency: 82, safety: 88 }
    },
    {
      id: 'midjourney-v7',
      name: 'Midjourney V7',
      provider: 'Midjourney',
      description: '艺术风格与细节表现出色，创意设计优选。',
      iconColor: '#ef4444',
      performance: { intelligence: 86, speed: 90, costEfficiency: 84, safety: 86 }
    },
    {
      id: 'doubao-seedream-4-0',
      name: 'Doubao-SeeDream-4.0',
      provider: 'Doubao',
      description: '文生图与编辑评测领先，2K/4K 输出稳定。',
      iconColor: '#0ea5e9',
      performance: { intelligence: 85, speed: 88, costEfficiency: 86, safety: 87 }
    }
  ],
  videoModels: [
    {
      id: 'sora-2',
      name: 'Sora-2',
      provider: 'OpenAI',
      description: '长镜头与物理一致性强，电影级视频生成。',
      iconColor: '#3B82F6',
      performance: { intelligence: 92, speed: 80, costEfficiency: 78, safety: 90 }
    },
    {
      id: 'veo-3-1',
      name: 'Veo 3.1',
      provider: 'Google',
      description: '1–3 张参考图保证角色一致，多风格快速渲染。',
      iconColor: '#0EA5E9',
      performance: { intelligence: 88, speed: 86, costEfficiency: 82, safety: 88 }
    },
    {
      id: 'kling',
      name: 'Kling',
      provider: 'Kling',
      description: '文本/图像/视频多模态输入，支持对象级编辑。',
      iconColor: '#22c55e',
      performance: { intelligence: 86, speed: 84, costEfficiency: 84, safety: 86 }
    }
  ],
  audioModels: [
    {
      id: 'gpt-audio',
      name: 'gpt-audio',
      provider: 'OpenAI',
      description: '兼顾识别、合成与指令理解，音频多模态。',
      iconColor: '#22d3ee',
      performance: { intelligence: 85, speed: 90, costEfficiency: 84, safety: 88 }
    },
    {
      id: 'speech-2-5-hd-preview',
      name: 'speech-2.5-hd-preview',
      provider: 'OpenAI',
      description: '高清语音合成预览，更自然音色与更低噪声。',
      iconColor: '#10b981',
      performance: { intelligence: 84, speed: 88, costEfficiency: 82, safety: 88 }
    }
  ],
  embeddingModels: [
    {
      id: 'text-embedding-3-small',
      name: 'text-embedding-3-small',
      provider: 'OpenAI',
      description: '小尺寸嵌入，性价比高，适合大规模检索。',
      iconColor: '#94a3b8',
      performance: { intelligence: 80, speed: 92, costEfficiency: 92, safety: 88 }
    },
    {
      id: 'doubao-embedding',
      name: 'doubao-embedding',
      provider: 'Doubao',
      description: '中文表现优秀，适合检索与匹配。',
      iconColor: '#10b981',
      performance: { intelligence: 82, speed: 90, costEfficiency: 88, safety: 88 }
    }
  ]
};
import { colors, gradients, animationStyles, createGradientText } from '../styles/theme';

const CategorySection = () => {
  const [activeCategory, setActiveCategory] = useState('text');

  const getModelsForCategory = (categoryId) => {
    switch (categoryId) {
      case 'text':
        return mockModels.textModels;
      case 'image':
        return mockModels.imageModels;
      case 'video':
        return mockModels.videoModels;
      case 'audio':
        return mockModels.audioModels;
      case 'embedding':
        return mockModels.embeddingModels;
      default:
        return mockModels.textModels;
    }
  };

  const handleCategoryChange = (categoryId) => {
    setActiveCategory(categoryId);
  };

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
      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1, maxWidth: '1200px' }}>
        {/* 标题区域 */}
        <Box
          sx={{
            textAlign: 'center',
            mb: { xs: 6, md: 10 },
            ...animationStyles.fadeIn
          }}
        >
          <Typography
            variant="h2"
            sx={{
              fontSize: { xs: '2.5rem', md: '3.5rem', lg: '4rem' },
              fontWeight: 200,
              color: colors.primary,
              mb: 4,
              lineHeight: 1.1,
              letterSpacing: '-0.02em',
              textShadow: 'none',
              '& .gradient-text': {
                fontWeight: 'bold',
                ...createGradientText('linear-gradient(45deg, #22D3EE 0%, #0EA5FF 50%, #8B5CF6 100%)')
              }
            }}
          >
            <span className="gradient-text">全面覆盖</span> 各类AI需求
          </Typography>
        </Box>

        {/* 分类标签 */}
        <Box
          sx={{
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
            gap: { xs: 1.5, md: 2 },
            mb: { xs: 6, md: 8 },
            px: { xs: 2, md: 0 }
          }}
        >
          {categories.map((category) => {
            const isActive = activeCategory === category.id;
            return (
              <Button
                key={category.id}
                onClick={() => handleCategoryChange(category.id)}
                startIcon={<Icon icon={category.icon} width={20} height={20} />}
                sx={{
                  px: { xs: 2.5, sm: 3, md: 4 },
                  py: { xs: 1.25, md: 2 },
                  borderRadius: '50px',
                  fontWeight: 600,
                  fontSize: { xs: '0.8125rem', sm: '0.875rem', md: '1rem' },
                  textTransform: 'none',
                  transition: 'all 0.3s ease',
                  minWidth: { xs: 'auto', sm: '120px' },
                  ...(isActive
                    ? {
                        background: gradients.primary,
                        color: 'white',
                        boxShadow: 'none',
                        '&:hover': {
                          background: gradients.primary,
                          transform: 'translateY(-2px)',
                          boxShadow: 'none'
                        }
                      }
                    : {
                        backgroundColor: '#F1F5F9',
                        color: colors.secondary,
                        '&:hover': {
                          backgroundColor: '#E2E8F0',
                          transform: 'translateY(-2px)',
                          boxShadow: 'none'
                        }
                      })
                }}
              >
                {category.name}
              </Button>
            );
          })}
        </Box>

        {/* 模型展示区域 */}
        <Box
          sx={{
            minHeight: '400px',
            position: 'relative'
          }}
        >
          <Grid container spacing={{ xs: 3, md: 4 }}>
            {getModelsForCategory(activeCategory).map((model, index) => (
              <Grid item xs={12} sm={6} lg={4} key={`${activeCategory}-${model.id}`}>
                <Box
                  sx={{
                    ...animationStyles.fadeIn,
                    animationDelay: `${index * 0.1}s`
                  }}
                >
                  <ModelCard model={model} variant="category" />
                </Box>
              </Grid>
            ))}
          </Grid>

          {/* 空状态 */}
          {getModelsForCategory(activeCategory).length === 0 && (
            <Box
              sx={{
                textAlign: 'center',
                py: 8,
                ...animationStyles.fadeIn
              }}
            >
              <Icon
                icon="solar:box-minimalistic-bold-duotone"
                width={64}
                height={64}
                style={{ color: colors.secondary, opacity: 0.5, marginBottom: '16px' }}
              />
              <Typography
                variant="h6"
                sx={{
                  color: colors.secondary,
                  mb: 1,
                  fontWeight: 500
                }}
              >
                暂无模型
              </Typography>
              <Typography
                variant="body2"
                sx={{
                  color: colors.secondary,
                  opacity: 0.7
                }}
              >
                该分类下的模型正在准备中，敬请期待
              </Typography>
            </Box>
          )}
        </Box>
      </Container>
    </Box>
  );
};

export default CategorySection;
