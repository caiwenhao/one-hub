import { Box, Typography, Button, Card, keyframes } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import { useTranslation } from 'react-i18next';
import { Icon } from '@iconify/react';

// 动画定义
const gradientShift = keyframes`
  0% { 
    background-position: 0% 50%; 
  }
  50% { 
    background-position: 100% 50%; 
  }
  100% { 
    background-position: 0% 50%; 
  }
`;

const FreeTrialSection = () => {
  const { t } = useTranslation();
  const theme = useTheme();

  const handleRegisterClick = () => {
    // 跳转到注册页面
    window.location.href = '/register';
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
      {/* 背景渐变 - 与明星模型荟萃区块保持一致 */}
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          background: 'linear-gradient(to bottom right, rgba(249, 250, 251, 0.5), rgba(59, 130, 246, 0.03))'
        }}
      />

      <Box sx={{ maxWidth: '1200px', mx: 'auto', px: { xs: 3, md: 6, lg: 12 }, position: 'relative', zIndex: 1 }}>
        <Card
          sx={{
            position: 'relative',
            p: { xs: 6, md: 8 },
            borderRadius: '24px',
            background: 'linear-gradient(135deg, rgba(255, 255, 255, 0.9) 0%, rgba(249, 250, 251, 0.9) 100%)',
            border: '1px solid rgba(0, 0, 0, 0.05)',
            boxShadow: '0 25px 50px rgba(0,0,0,0.1)',
            overflow: 'hidden',
            transition: 'all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275)',
            '&:hover': {
              transform: 'translateY(-8px) scale(1.02)',
              boxShadow: '0 32px 64px rgba(0,0,0,0.15)',
              borderColor: 'rgba(66, 153, 225, 0.3)'
            },
            ...(theme.palette.mode === 'dark' && {
              background: 'linear-gradient(135deg, rgba(26, 26, 26, 0.9) 0%, rgba(31, 31, 31, 0.9) 100%)',
              borderColor: 'rgba(255, 255, 255, 0.1)'
            })
          }}
        >
          {/* 背景渐变层 - 调整为与明星模型荟萃一致的风格 */}
          <Box
            sx={{
              position: 'absolute',
              inset: 0,
              background: 'linear-gradient(to bottom right, rgba(249, 250, 251, 0.3), rgba(59, 130, 246, 0.02))',
              ...(theme.palette.mode === 'dark' && {
                background: 'linear-gradient(to bottom right, rgba(31, 31, 31, 0.3), rgba(59, 130, 246, 0.05))'
              })
            }}
          />

          {/* 主要内容 */}
          <Box sx={{ position: 'relative', zIndex: 10, textAlign: 'center' }}>
            {/* 限时福利标签 - 调整为与明星模型荟萃一致的风格 */}
            <Box
              sx={{
                display: 'inline-flex',
                alignItems: 'center',
                background: 'linear-gradient(135deg, rgba(66, 153, 225, 0.1), rgba(139, 92, 246, 0.1))',
                px: 3,
                py: 1.5,
                borderRadius: '50px',
                mb: 3,
                border: 'none',
                ...(theme.palette.mode === 'dark' && {
                  background: 'linear-gradient(135deg, rgba(66, 153, 225, 0.15), rgba(139, 92, 246, 0.15))'
                })
              }}
            >
              <Icon
                icon="solar:gift-bold-duotone"
                width={20}
                height={20}
                style={{
                  color: '#0EA5FF', // 统一为AI科技主色
                  marginRight: '12px'
                }}
              />
              <Typography
                variant="body1"
                sx={{
                  color: '#0EA5FF', // 统一为AI科技主色
                  fontWeight: 600,
                  fontSize: '1rem'
                }}
              >
                限时福利
              </Typography>
            </Box>

            {/* 主标题 */}
            <Typography
              variant="h2"
              sx={{
                fontSize: { xs: '2rem', sm: '2.5rem', md: '3rem', lg: '3.5rem' },
                fontWeight: 700,
                color: theme.palette.text.primary,
                mb: 3,
                textShadow: 'none'
              }}
            >
              {t('modelpricePage.freeTrialTitle')}
            </Typography>

            {/* 描述文字 */}
            <Typography
              variant="h5"
              sx={{
                fontSize: { xs: '1.1rem', sm: '1.25rem', md: '1.5rem' },
                color: theme.palette.text.secondary,
                mb: 4,
                maxWidth: '600px',
                mx: 'auto',
                lineHeight: 1.6
              }}
            >
              每一位新注册用户，均可立即获得{' '}
              <Box
                component="span"
                sx={{
                  fontSize: { xs: '1.5rem', sm: '1.75rem', md: '2rem' },
                  fontWeight: 700,
                  background: 'linear-gradient(45deg, #22D3EE 0%, #0EA5FF 50%, #8B5CF6 100%)',
                  WebkitBackgroundClip: 'text',
                  WebkitTextFillColor: 'transparent',
                  backgroundClip: 'text'
                }}
              >
                ¥10.00
              </Box>{' '}
              免费试用额度
            </Typography>

            {/* CTA按钮区域 */}
            <Box
              sx={{
                display: 'flex',
                flexDirection: { xs: 'column', sm: 'row' },
                gap: 2,
                justifyContent: 'center',
                alignItems: 'center'
              }}
            >
              <Button
                onClick={handleRegisterClick}
                sx={{
                  position: 'relative',
                  background: `linear-gradient(-45deg, ${theme.palette.primary.main}, #22D3EE, #8B5CF6)`,
                  backgroundSize: '400% 400%',
                  color: 'white',
                  px: 5,
                  py: 2,
                  borderRadius: '50px',
                  fontWeight: 700,
                  fontSize: '1.125rem',
                  textTransform: 'none',
                  animation: `${gradientShift} 4s ease infinite`,
                  transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
                  overflow: 'hidden',
                  '&:hover': {
                    transform: 'scale(1.05)',
                    boxShadow: 'none'
                  }
                }}
              >
                立即注册，领取额度 →
              </Button>

              <Typography
                variant="body2"
                sx={{
                  color: theme.palette.text.secondary,
                  fontSize: '0.875rem'
                }}
              >
                无需信用卡 · 即开即用
              </Typography>
            </Box>
          </Box>
        </Card>
      </Box>
    </Box>
  );
};

export default FreeTrialSection;
