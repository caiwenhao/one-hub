import { Box, Typography, keyframes } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import { useTranslation } from 'react-i18next';

// 动画定义
const float = keyframes`
  0%, 100% { 
    transform: translateY(0px) rotate(0deg); 
  }
  50% { 
    transform: translateY(-20px) rotate(180deg); 
  }
`;

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

const HeroSection = () => {
  const { t } = useTranslation();
  const theme = useTheme();

  return (
    <Box
      sx={{
        position: 'relative',
        background: 'linear-gradient(135deg, #ffffff 0%, rgba(59, 130, 246, 0.03) 25%, rgba(147, 51, 234, 0.02) 50%, #ffffff 100%)',
        pt: { xs: 8, md: 12 },
        pb: { xs: 8, md: 10 },
        px: { xs: 3, md: 6, lg: 12 },
        minHeight: '500px',
        display: 'flex',
        alignItems: 'center',
        overflow: 'hidden',
        ...(theme.palette.mode === 'dark' && {
          background: 'linear-gradient(135deg, #1a1a1a 0%, rgba(59, 130, 246, 0.05) 25%, rgba(147, 51, 234, 0.03) 50%, #1a1a1a 100%)'
        })
      }}
    >
      {/* 背景渐变层 */}
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          background: 'linear-gradient(135deg, rgba(66, 153, 225, 0.08) 0%, rgba(147, 51, 234, 0.05) 50%, rgba(236, 72, 153, 0.08) 100%)',
          ...(theme.palette.mode === 'dark' && {
            background: 'linear-gradient(135deg, rgba(66, 153, 225, 0.12) 0%, rgba(147, 51, 234, 0.08) 50%, rgba(236, 72, 153, 0.12) 100%)'
          })
        }}
      />

      {/* 浮动装饰元素 */}
      <Box
        sx={{
          position: 'absolute',
          top: '80px',
          left: '40px',
          width: '12px',
          height: '12px',
          backgroundColor: 'rgba(66, 153, 225, 0.3)',
          borderRadius: '50%',
          animation: `${float} 6s ease-in-out infinite`,
          display: { xs: 'none', md: 'block' }
        }}
      />
      <Box
        sx={{
          position: 'absolute',
          top: '160px',
          right: '80px',
          width: '8px',
          height: '8px',
          backgroundColor: 'rgba(147, 51, 234, 0.4)',
          borderRadius: '50%',
          animation: `${float} 6s ease-in-out infinite 1.2s`,
          display: { xs: 'none', md: 'block' }
        }}
      />
      <Box
        sx={{
          position: 'absolute',
          bottom: '120px',
          left: '25%',
          width: '10px',
          height: '10px',
          backgroundColor: 'rgba(236, 72, 153, 0.35)',
          borderRadius: '50%',
          animation: `${float} 6s ease-in-out infinite 2.8s`,
          display: { xs: 'none', md: 'block' }
        }}
      />
      <Box
        sx={{
          position: 'absolute',
          top: '240px',
          right: '33%',
          width: '6px',
          height: '6px',
          backgroundColor: 'rgba(245, 158, 11, 0.4)',
          borderRadius: '50%',
          animation: `${float} 6s ease-in-out infinite 2s`,
          display: { xs: 'none', md: 'block' }
        }}
      />

      {/* 主要内容 */}
      <Box
        sx={{
          position: 'relative',
          maxWidth: '1200px',
          mx: 'auto',
          textAlign: 'center',
          zIndex: 10
        }}
      >
        {/* 已删除"简单透明只为价值付费"标题区块 */}

        <Typography
          variant="h5"
          sx={{
            fontSize: { xs: '1.1rem', sm: '1.25rem', md: '1.5rem' },
            color: theme.palette.text.secondary,
            mb: 6,
            maxWidth: '800px',
            mx: 'auto',
            lineHeight: 1.6,
            fontWeight: 300,
            textShadow: 'none'
          }}
        >
          {t('modelpricePage.heroDescription')}
          <br />
          {t('modelpricePage.heroDescription2')}
        </Typography>
      </Box>
    </Box>
  );
};

export default HeroSection;
