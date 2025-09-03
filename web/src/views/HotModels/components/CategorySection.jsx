import React, { useState } from 'react';
import { Box, Container, Typography, Grid, Button } from '@mui/material';
import { Icon } from '@iconify/react';
import ModelCard from './ModelCard';
import { mockModels, categories } from '../data/mockData';
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
              textShadow: '0 2px 4px rgba(0,0,0,0.1)',
              '& .gradient-text': {
                fontWeight: 'bold',
                ...createGradientText('linear-gradient(45deg, #4299E1 30%, #8B5CF6 90%)')
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
                        boxShadow: '0 8px 32px rgba(66, 153, 225, 0.3)',
                        ...animationStyles.pulseGlow,
                        '&:hover': {
                          background: gradients.primary,
                          transform: 'translateY(-2px)',
                          boxShadow: '0 12px 40px rgba(66, 153, 225, 0.4)'
                        }
                      }
                    : {
                        backgroundColor: '#F1F5F9',
                        color: colors.secondary,
                        '&:hover': {
                          backgroundColor: '#E2E8F0',
                          transform: 'translateY(-2px)',
                          boxShadow: '0 4px 20px rgba(0,0,0,0.1)'
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
                  <ModelCard model={model} showPerformance={false} />
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
