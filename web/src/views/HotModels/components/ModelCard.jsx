import React from 'react';
import { Box, Card, Typography, Chip, LinearProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { colors, gradients, animationStyles, tagStyles } from '../styles/theme';

const ModelCard = ({ model, showPerformance = false, variant = 'default' }) => {
  const navigate = useNavigate();

  const handleCardClick = () => {
    // 统一跳转到控制台
    navigate('/panel');
  };

  const getTagStyle = (tagType) => {
    return tagStyles[tagType] || tagStyles.recommended;
  };

  const renderPerformanceBar = (label, value, color = colors.accent) => (
    <Box sx={{ mb: 2 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 1 }}>
        <Typography variant="body2" sx={{ color: colors.secondary, fontSize: '0.875rem' }}>
          {label}
        </Typography>
        <Typography variant="body2" sx={{ color: colors.accent, fontWeight: 600, fontSize: '0.875rem' }}>
          {value}%
        </Typography>
      </Box>
      <Box sx={{ width: '100%', backgroundColor: '#E2E8F0', borderRadius: '3px', height: '6px' }}>
        <Box
          sx={{
            width: `${value}%`,
            height: '100%',
            background: `linear-gradient(90deg, ${color}, ${color}dd)`,
            borderRadius: '3px',
            transition: 'width 0.8s ease-out'
          }}
        />
      </Box>
    </Box>
  );

  return (
    <Card
      onClick={handleCardClick}
      sx={{
        height: '100%',
        background: gradients.card,
        border: '1px solid rgba(0, 0, 0, 0.05)',
        borderRadius: '24px',
        p: { xs: 2.5, md: 3 },
        position: 'relative',
        overflow: 'hidden',
        cursor: 'pointer',
        ...animationStyles.hoverLift,
        transition: 'all 0.5s cubic-bezier(0.4, 0, 0.2, 1)',
        '&:hover': {
          borderColor: 'rgba(66, 153, 225, 0.3)',
          transform: 'translateY(-8px) scale(1.02)',
          boxShadow: '0 20px 40px rgba(0,0,0,0.12)',
          '& .model-overlay': {
            opacity: 1
          },
          '& .model-icon': {
            transform: 'scale(1.1)'
          }
        }
      }}
    >
      {/* 悬停覆盖层 */}
      <Box
        className="model-overlay"
        sx={{
          position: 'absolute',
          inset: 0,
          background: `linear-gradient(to bottom right, ${model.iconColor}05, transparent)`,
          opacity: 0,
          transition: 'opacity 0.5s ease'
        }}
      />

      {/* 标签 */}
      {model.tag && (
        <Box sx={{ position: 'absolute', top: 16, right: 16, zIndex: 2 }}>
          <Chip
            label={model.tag.label}
            sx={{
              ...getTagStyle(model.tag.type),
              fontSize: '0.75rem',
              fontWeight: 'bold',
              px: 2,
              py: 1,
              borderRadius: '20px',
              boxShadow: '0 2px 8px rgba(0,0,0,0.15)',
              ...(model.tag.type === 'hot' && animationStyles.pulseGlow)
            }}
          />
        </Box>
      )}

      <Box sx={{ position: 'relative', zIndex: 1 }}>
        {/* 模型图标 */}
        <Box
          className="model-icon"
          sx={{
            width: { xs: 56, sm: 64, md: 80 },
            height: { xs: 56, sm: 64, md: 80 },
            background: `linear-gradient(135deg, ${model.iconColor}, ${model.iconColor}dd)`,
            borderRadius: { xs: '16px', md: '24px' },
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            mb: { xs: 2, md: 3 },
            color: 'white',
            fontWeight: 'bold',
            fontSize: { xs: '0.875rem', sm: '1rem', md: '1.25rem' },
            transition: 'transform 0.3s ease',
            boxShadow: `0 8px 32px ${model.iconColor}40`,
            ...animationStyles.shimmer
          }}
        >
          {model.icon}
        </Box>

        {/* 模型名称 */}
        <Typography
          variant="h6"
          sx={{
            fontSize: { xs: '1.125rem', sm: '1.25rem', md: '1.5rem' },
            fontWeight: 'bold',
            color: colors.primary,
            mb: { xs: 1.5, md: 2 },
            textShadow: '0 2px 4px rgba(0,0,0,0.1)',
            lineHeight: 1.2
          }}
        >
          {model.name}
        </Typography>

        {/* 模型描述 */}
        <Typography
          variant="body2"
          sx={{
            color: colors.secondary,
            lineHeight: 1.6,
            mb: showPerformance ? { xs: 2, md: 3 } : { xs: 3, md: 4 },
            fontSize: { xs: '0.8125rem', md: '0.875rem' },
            fontWeight: 300,
            display: '-webkit-box',
            WebkitLineClamp: { xs: 3, md: 4 },
            WebkitBoxOrient: 'vertical',
            overflow: 'hidden'
          }}
        >
          {model.description}
        </Typography>

        {/* 性能指标 */}
        {showPerformance && model.performance && (
          <Box sx={{ mb: 3 }}>
            {Object.entries(model.performance).map(([key, value]) => {
              const labels = {
                intelligence: '智能水平',
                speed: '响应速度',
                costEfficiency: '性价比',
                safety: '安全性'
              };
              return renderPerformanceBar(labels[key] || key, value, model.iconColor);
            })}
          </Box>
        )}

        {/* 价格信息 */}
        {model.pricing && (
          <Box sx={{ mb: 3 }}>
            {Object.entries(model.pricing).map(([key, value]) => (
              <Box key={key} sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 1 }}>
                <Typography variant="body2" sx={{ color: colors.secondary, fontSize: '0.875rem' }}>
                  {key === 'input' ? '输入成本' : key === 'output' ? '输出成本' : key === 'standard' ? '标准生成' : key === 'hd' ? '高清生成' : key}
                </Typography>
                <Typography variant="body2" sx={{ color: colors.accent, fontWeight: 600, fontSize: '0.875rem' }}>
                  {value}
                </Typography>
              </Box>
            ))}
          </Box>
        )}

        {/* 底部信息 */}
        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
          <Typography variant="body2" sx={{ color: colors.secondary, fontSize: '0.875rem' }}>
            {model.provider} 出品
          </Typography>
          <Typography
            variant="body2"
            sx={{
              color: colors.accent,
              fontWeight: 600,
              fontSize: '0.875rem',
              transition: 'color 0.3s ease',
              '&:hover': {
                color: colors.purple
              }
            }}
          >
            查看详情 →
          </Typography>
        </Box>
      </Box>
    </Card>
  );
};

export default ModelCard;
