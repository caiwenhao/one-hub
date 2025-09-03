import React from 'react';
import { Box, Typography, Container, Grid, Paper } from '@mui/material';
import { keyframes } from '@mui/system';
import HandshakeIcon from '@mui/icons-material/Handshake';
import CodeIcon from '@mui/icons-material/Code';
import EmailIcon from '@mui/icons-material/Email';

// 定义动画
const pulseGlow = keyframes`
  0%, 100% { 
    box-shadow: 0 0 20px rgba(66, 153, 225, 0.3); 
  }
  50% { 
    box-shadow: 0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3); 
  }
`;

// 联系卡片组件
const ContactCard = ({
  title,
  description,
  email,
  icon: Icon,
  gradientColors,
  bgGradient
}) => {
  const handleEmailClick = () => {
    window.open(`mailto:${email}`, '_blank');
  };

  return (
    <Paper
      elevation={0}
      sx={{
        background: bgGradient,
        p: 5, // p-10
        borderRadius: '24px', // rounded-3xl
        boxShadow: '0 25px 50px -12px rgba(0, 0, 0, 0.25)', // shadow-xl
        border: '1px solid rgba(229, 231, 235, 0.5)', // border-gray-100/50
        textAlign: 'center',
        transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
        position: 'relative',
        overflow: 'hidden',
        '&:hover': {
          transform: 'translateY(-4px)',
          boxShadow: '0 35px 70px rgba(0, 0, 0, 0.1)'
        }
      }}
    >
      {/* 图标 */}
      <Box
        sx={{
          width: '80px',
          height: '80px',
          borderRadius: '50%',
          background: gradientColors,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: 'white',
          fontSize: '32px',
          mb: 3, // mb-6
          mx: 'auto',
          animation: `${pulseGlow} 3s ease-in-out infinite`
        }}
      >
        <Icon sx={{ fontSize: '32px' }} />
      </Box>

      {/* 标题 */}
      <Typography
        variant="h4"
        sx={{
          fontSize: '1.5rem', // text-2xl
          fontWeight: 700, // font-bold
          color: '#1A202C', // text-primary
          mb: 2 // mb-4
        }}
      >
        {title}
      </Typography>

      {/* 描述 */}
      <Typography
        variant="body1"
        sx={{
          color: '#718096', // text-gray-600
          mb: 3, // mb-6
          lineHeight: 1.6 // leading-relaxed
        }}
      >
        {description}
      </Typography>

      {/* 邮箱卡片 */}
      <Box
        sx={{
          backgroundColor: 'white',
          p: 2, // p-4
          borderRadius: '12px', // rounded-xl
          boxShadow: 'inset 0 2px 4px 0 rgba(0, 0, 0, 0.06)', // shadow-inner
          mb: 3, // mb-6
          cursor: 'pointer',
          transition: 'all 0.3s ease',
          '&:hover': {
            transform: 'scale(1.02)',
            boxShadow: 'inset 0 2px 4px 0 rgba(0, 0, 0, 0.1)'
          }
        }}
        onClick={handleEmailClick}
      >
        <Typography
          sx={{
            color: gradientColors.includes('#10B981') ? '#059669' : '#4299E1', // 根据卡片类型调整颜色
            fontWeight: 600, // font-semibold
            fontSize: '1.125rem', // text-lg
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            gap: 1,
            '&:hover': {
              color: gradientColors.includes('#10B981') ? '#047857' : '#3182CE'
            }
          }}
        >
          <EmailIcon sx={{ fontSize: '1.125rem' }} />
          {email}
        </Typography>
      </Box>
    </Paper>
  );
};

const ContactChannels = () => {
  return (
    <Box
      component="section"
      sx={{
        backgroundColor: 'white',
        py: { xs: 8, md: 10 }, // py-20
        px: { xs: 3, md: 6, lg: 12 } // px-6 md:px-12 lg:px-24
      }}
    >
      <Container maxWidth="lg" sx={{ maxWidth: '1200px' }}>
        <Grid container spacing={{ xs: 6, md: 6 }} sx={{ mb: 8 }}>
          {/* 企业合作卡片 */}
          <Grid item xs={12} md={6}>
            <ContactCard
              title="企业合作与售前咨询"
              description={
                <>
                  寻求企业级AI解决方案？想了解批量折扣或定制服务？
                  <br />
                  我们的商务团队将为您提供专业建议
                </>
              }
              email="sales@kapon.ai"
              icon={HandshakeIcon}
              gradientColors="linear-gradient(135deg, #4299E1, #3182CE)"
              bgGradient="linear-gradient(135deg, rgba(59, 130, 246, 0.05) 0%, white 100%)"
            />
          </Grid>

          {/* 技术支持卡片 */}
          <Grid item xs={12} md={6}>
            <ContactCard
              title="技术支持与开发者帮助"
              description={
                <>
                  集成遇到困难？API调用出现问题？功能建议？
                  <br />
                  我们的技术专家随时为您解答
                </>
              }
              email="support@kapon.ai"
              icon={CodeIcon}
              gradientColors="linear-gradient(135deg, #10B981, #059669)"
              bgGradient="linear-gradient(135deg, rgba(16, 185, 129, 0.05) 0%, white 100%)"
            />
          </Grid>
        </Grid>
      </Container>
    </Box>
  );
};

export default ContactChannels;
