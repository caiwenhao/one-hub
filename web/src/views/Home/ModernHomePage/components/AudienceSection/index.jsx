import React from 'react';
import {
  Box,
  Container,
  Typography,
  Grid,
  Card,
  CardContent,
  Button,
  useTheme,
  useMediaQuery
} from '@mui/material';
import { Business, Code, ArrowForward } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGradientText, createFloatingElement } from '../../styles/theme';

const AudienceSection = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const navigate = useNavigate();

  const audiences = [
    {
      icon: Business,
      title: '面向中小企业',
      description: '为您的产品快速、低成本地集成强大的AI能力。我们提供正规合同与发票，让您合作无忧，专注业务增长。',
      gradient: gradients.featureAccent,
      iconGradient: gradients.iconAccent,
      iconColor: colors.accent,
      features: [
        '正规合同与发票',
        '企业级技术支持',
        '定制化解决方案',
        '批量折扣优惠'
      ]
    },
    {
      icon: Code,
      title: '面向个人开发者/创业者',
      description: '将您的创意快速落地。我们提供灵活的按需付费和免费试用额度，是您启动AI项目的最佳伙伴。',
      gradient: gradients.featurePurple,
      iconGradient: gradients.iconPurple,
      iconColor: colors.purple,
      features: [
        '免费试用额度',
        '按需付费模式',
        '完整开发文档',
        '社区技术支持'
      ]
    }
  ];

  return (
    <Box
      component="section"
      sx={{
        position: 'relative',
        background: 'linear-gradient(to bottom right, rgba(249, 250, 251, 0.5) 0%, rgba(59, 130, 246, 0.03) 30%, rgba(139, 92, 246, 0.03) 70%, rgba(249, 250, 251, 0.5) 100%)',
        py: { xs: 4, md: 8 },
        overflow: 'hidden'
      }}
    >
      {/* 浮动背景元素 */}
      <Box
        sx={{
          position: 'absolute',
          bottom: '25%',
          left: '25%',
          width: '320px',
          height: '320px',
          background: 'linear-gradient(to right, rgba(139, 92, 246, 0.08), rgba(236, 72, 153, 0.08))',
          borderRadius: '50%',
          filter: 'blur(60px)',
          ...animationStyles.floating
        }}
      />

      <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1 }}>
        {/* 标题 */}
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
            专为追求 <span className="gradient-text">卓越</span> 的你而设计
          </Typography>
        </Box>

        {/* 用户群体卡片 */}
        <Grid container spacing={{ xs: 4, md: 6 }}>
          {audiences.map((audience, index) => {
            const IconComponent = audience.icon;
            return (
              <Grid item xs={12} lg={6} key={audience.title}>
                <Card
                  sx={{
                    height: '100%',
                    background: gradients.card,
                    border: '1px solid rgba(0, 0, 0, 0.05)',
                    borderRadius: '24px',
                    p: { xs: 3, md: 4 },
                    position: 'relative',
                    overflow: 'hidden',
                    transition: 'border-color 0.2s ease',
                    '&:hover': {
                      borderColor: 'rgba(66, 153, 225, 0.3)'
                    }
                  }}
                >


                  <CardContent sx={{ p: 0, position: 'relative', zIndex: 1 }}>
                    {/* 图标 */}
                    <Box
                      className="audience-icon"
                      sx={{
                        width: 80,
                        height: 80,
                        borderRadius: '24px',
                        background: audience.iconGradient,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        mb: 4
                      }}
                    >
                      <IconComponent
                        sx={{
                          fontSize: '2rem',
                          color: audience.iconColor
                        }}
                      />
                    </Box>

                    {/* 标题 */}
                    <Typography
                      variant="h3"
                      sx={{
                        fontSize: '1.875rem',
                        fontWeight: 'bold',
                        color: colors.primary,
                        mb: 3
                      }}
                    >
                      {audience.title}
                    </Typography>

                    {/* 描述 */}
                    <Typography
                      variant="body1"
                      sx={{
                        color: colors.secondary,
                        lineHeight: 1.6,
                        fontSize: '1.25rem',
                        fontWeight: 300,
                        mb: 4
                      }}
                    >
                      {audience.description}
                    </Typography>

                    {/* 特性列表 */}
                    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                      {audience.features.map((feature, featureIndex) => (
                        <Box
                          key={feature}
                          sx={{
                            display: 'flex',
                            alignItems: 'center',
                            gap: 2,
                            ...animationStyles.fadeIn,
                            animationDelay: `${index * 0.3 + featureIndex * 0.1 + 0.5}s`
                          }}
                        >
                          <Box
                            sx={{
                              width: 8,
                              height: 8,
                              borderRadius: '50%',
                              background: gradients.primary,
                              ...animationStyles.pulseGlow,
                              flexShrink: 0
                            }}
                          />
                          <Typography
                            variant="body2"
                            sx={{
                              color: colors.primary,
                              fontSize: '1rem',
                              fontWeight: 500
                            }}
                          >
                            {feature}
                          </Typography>
                        </Box>
                      ))}
                    </Box>

                    {/* 企业客户联系我们按钮 */}
                    {index === 0 && (
                      <Box sx={{ mt: 4 }}>
                        <Button
                          variant="contained"
                          endIcon={<ArrowForward />}
                          onClick={() => navigate('/contact')}
                          sx={{
                            background: gradients.primary,
                            color: '#ffffff',
                            px: 4,
                            py: 1.5,
                            borderRadius: '25px',
                            fontSize: '1rem',
                            fontWeight: 600,
                            textTransform: 'none',
                            boxShadow: '0 8px 20px rgba(66, 153, 225, 0.3)',
                            transition: 'all 0.3s ease',
                            '&:hover': {
                              background: gradients.primary,
                              transform: 'translateY(-2px)',
                              boxShadow: '0 12px 30px rgba(66, 153, 225, 0.4)'
                            }
                          }}
                        >
                          联系我们
                        </Button>
                      </Box>
                    )}
                  </CardContent>
                </Card>
              </Grid>
            );
          })}
        </Grid>

        {/* 底部信任指标 */}
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
            '已服务 500+ 企业客户',
            '支持 10,000+ 个人开发者',
            '累计处理 100M+ API 调用',
            '客户满意度 99.5%'
          ].map((stat, index) => (
            <Box
              key={stat}
              sx={{
                display: 'flex',
                alignItems: 'center',
                gap: 1,
                px: 3,
                py: 1.5,
                borderRadius: '25px',
                background: 'rgba(255, 255, 255, 0.8)',
                backdropFilter: 'blur(10px)',
                border: '1px solid rgba(0, 0, 0, 0.05)',
                transition: 'background-color 0.2s ease',
                '&:hover': {
                  background: 'rgba(255, 255, 255, 0.9)'
                }
              }}
            >
              <Box
                sx={{
                  width: 6,
                  height: 6,
                  borderRadius: '50%',
                  background: gradients.primary,
                  ...animationStyles.pulseGlow
                }}
              />
              <Typography
                variant="body2"
                sx={{
                  color: colors.primary,
                  fontSize: '0.9rem',
                  fontWeight: 500
                }}
              >
                {stat}
              </Typography>
            </Box>
          ))}
        </Box>
      </Container>
    </Box>
  );
};

export default AudienceSection;
