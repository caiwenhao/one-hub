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
  const isCategory = variant === 'category';
  const cardHeight = isFeatured
    ? { xs: '500px', md: '540px' }
    : isCategory
    ? { xs: '420px', md: '440px' }
    : { xs: '420px', md: '450px' };



  return (
    <Card
      onClick={handleCardClick}
      sx={{
        height: cardHeight,
        // 分类卡片使用更接近设计稿的蓝灰边与浅蓝底色
        background: isCategory
          ? 'linear-gradient(180deg, #ffffff 0%, #f6faff 100%)'
          : gradients.card,
        border: isCategory ? '1px solid #E5EEF9' : '1px solid rgba(0, 0, 0, 0.05)',
        borderRadius: '24px',
        p: { xs: 2.5, md: 3 },
        position: 'relative',
        overflow: 'hidden',
        cursor: 'pointer',
        display: 'flex',
        flexDirection: 'column',
        ...animationStyles.hoverLift,
        transition: 'all 0.5s cubic-bezier(0.4, 0, 0.2, 1)',
        // 顶部左侧淡淡的径向高光，贴合截图效果
        ...(isCategory && {
          '&::before': {
            content: '""',
            position: 'absolute',
            top: -40,
            left: -40,
            width: 200,
            height: 200,
            background: `radial-gradient(closest-side, ${model.iconColor}15, transparent 70%)`,
            pointerEvents: 'none'
          }
        }),
        '&:hover': {
          borderColor: isCategory ? '#d7e6fb' : 'rgba(66, 153, 225, 0.3)',
          transform: 'translateY(-8px) scale(1.02)',
          boxShadow: isCategory
            ? '0 18px 40px rgba(66,153,225,0.12)'
            : '0 20px 40px rgba(0,0,0,0.12)',
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
              boxShadow: '0 2px 8px rgba(0,0,0,0.15)'
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
              width: isCategory ? { xs: 56, sm: 64, md: 72 } : { xs: 56, sm: 64, md: 80 },
              height: isCategory ? { xs: 56, sm: 64, md: 72 } : { xs: 56, sm: 64, md: 80 },
              background: `linear-gradient(135deg, ${model.iconColor}, ${model.iconColor}dd)`,
              borderRadius: { xs: '16px', md: '20px' },
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              mb: isCategory ? { xs: 1.5, md: 2 } : { xs: 2, md: 3 },
              color: 'white',
              fontWeight: 800,
              fontSize: isCategory ? { xs: '1rem', sm: '1.125rem', md: '1.25rem' } : { xs: '0.875rem', sm: '1rem', md: '1.125rem' },
              transition: 'transform 0.3s ease',
              boxShadow: isCategory
                ? '0 12px 24px rgba(59,130,246,0.18)'
                : '0 8px 16px rgba(15, 23, 42, 0.12)',
              // 轻微上移，模拟浮出效果
              position: isCategory ? 'relative' : 'static',
              top: isCategory ? -6 : 0
            }}
          >
            {/* 使用缩写字母（优先 abbr；否则根据 provider/name 生成） */}
            {(() => {
              const text = (() => {
                if (model.abbr) return String(model.abbr).toUpperCase().slice(0, 2);
                const base = (model.provider || model.name || '').trim();
                if (!base) return '?';
                // 提取字母数字词，优先两字符；多词取首字母组合
                const parts = base.split(/[^A-Za-z0-9]+/).filter(Boolean);
                if (parts.length === 1) return parts[0].slice(0, 2).toUpperCase();
                return parts.map((p) => p[0]).join('').slice(0, 2).toUpperCase();
              })();
              return (
                <Typography component="span" sx={{ lineHeight: 1, letterSpacing: '0.5px' }}>
                  {text}
                </Typography>
              );
            })()}
          </Box>
        )}

        {/* 模型名称 */}
        <Typography
          variant={isFeatured ? "h5" : "h6"}
          sx={{
            fontSize: isFeatured
              ? { xs: '1.5rem', sm: '1.75rem', md: '2rem' }
              : { xs: '1.125rem', sm: '1.25rem', md: '1.5rem' },
            fontWeight: isFeatured ? 800 : 'bold',
            color: colors.primary,
            mb: { xs: 1.5, md: 2 },
            mt: isFeatured ? { xs: 1, md: 2 } : 0,
            textShadow: 'none',
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
            textAlign: 'left',
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
                      boxShadow: 'none'
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
