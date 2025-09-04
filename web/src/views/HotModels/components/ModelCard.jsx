import React from 'react';
import { Box, Card, Typography, Chip } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { colors, gradients, animationStyles, tagStyles } from '../styles/theme';

const ModelCard = ({ model, variant = 'default', showPerformance = false }) => {
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

  // 根据variant决定使用哪种样式
  const isFeatured = variant === 'featured';
  const cardHeight = isFeatured ? { xs: '480px', md: '520px' } : { xs: '420px', md: '450px' };



  return (
    <Card
      onClick={handleCardClick}
      sx={{
        height: cardHeight,
        background: gradients.card,
        border: '1px solid rgba(0, 0, 0, 0.05)',
        borderRadius: '24px',
        p: { xs: 2.5, md: 3 },
        position: 'relative',
        overflow: 'hidden',
        cursor: 'pointer',
        display: 'flex',
        flexDirection: 'column',
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

      <Box sx={{ position: 'relative', zIndex: 1, display: 'flex', flexDirection: 'column', height: '100%' }}>
        {/* 模型图标 - 只在非featured模式显示 */}
        {!isFeatured && (
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
        )}

        {/* 模型名称 */}
        <Typography
          variant={isFeatured ? "h5" : "h6"}
          sx={{
            fontSize: isFeatured
              ? { xs: '1.25rem', sm: '1.375rem', md: '1.625rem' }
              : { xs: '1.125rem', sm: '1.25rem', md: '1.5rem' },
            fontWeight: isFeatured ? 800 : 'bold',
            color: colors.primary,
            mb: { xs: 1.5, md: 2 },
            mt: isFeatured ? { xs: 1, md: 2 } : 0,
            textShadow: '0 2px 4px rgba(0,0,0,0.1)',
            lineHeight: 1.2,
            letterSpacing: isFeatured ? '-0.5px' : '0'
          }}
        >
          {model.name}
        </Typography>

        {/* 模型描述 */}
        <Typography
          variant="body2"
          sx={{
            color: colors.secondary,
            lineHeight: isFeatured ? 1.8 : 1.6,
            mb: isFeatured ? { xs: 4, md: 5 } : { xs: 2, md: 3 },
            fontSize: isFeatured
              ? { xs: '0.9rem', md: '0.95rem' }
              : { xs: '0.8125rem', md: '0.875rem' },
            fontWeight: 400,
            textAlign: isFeatured ? 'justify' : 'left',
            minHeight: isFeatured ? { xs: '120px', md: '140px' } : 'auto',
            display: isFeatured ? 'block' : '-webkit-box',
            WebkitLineClamp: isFeatured ? 'none' : { xs: 3, md: 4 },
            WebkitBoxOrient: isFeatured ? 'initial' : 'vertical',
            overflow: isFeatured ? 'visible' : 'hidden'
          }}
        >
          {model.description}
        </Typography>

        {/* 性能指标 - 只在非featured模式显示 */}
        {!isFeatured && model.performance && (
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

        {/* 应用场景 - 只在featured模式显示 */}
        {isFeatured && model.useCases && (
          <Box sx={{ mb: 5 }}>
            <Typography
              variant="body2"
              sx={{
                color: colors.primary,
                fontSize: '0.95rem',
                fontWeight: 600,
                mb: 2.5,
                letterSpacing: '0.5px'
              }}
            >
              💡 优势应用场景
            </Typography>
            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1.5, minHeight: '60px' }}>
              {model.useCases.map((useCase, index) => (
                <Chip
                  key={index}
                  label={useCase}
                  size="small"
                  sx={{
                    backgroundColor: `${model.iconColor}18`,
                    color: model.iconColor,
                    fontSize: '0.85rem',
                    fontWeight: 600,
                    border: `1px solid ${model.iconColor}40`,
                    borderRadius: '16px',
                    px: 2,
                    py: 0.8,
                    height: 'auto',
                    transition: 'all 0.3s ease',
                    '&:hover': {
                      backgroundColor: `${model.iconColor}30`,
                      transform: 'translateY(-1px)',
                      boxShadow: `0 4px 12px ${model.iconColor}25`
                    }
                  }}
                />
              ))}
            </Box>
          </Box>
        )}

        {/* 价格信息 - 只在非featured模式显示 */}
        {!isFeatured && model.pricing && (
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

        {/* 弹性空间 */}
        <Box sx={{ flexGrow: 1 }} />

        {/* 底部信息 */}
        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', mt: 'auto' }}>
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
