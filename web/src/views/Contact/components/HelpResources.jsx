import React from 'react';
import { Box, Typography, Container, Grid, Paper, Button } from '@mui/material';
import MenuBookIcon from '@mui/icons-material/MenuBook';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import FavoriteIcon from '@mui/icons-material/Favorite';

// 资源卡片组件
const ResourceCard = ({ 
  title, 
  description, 
  buttonText, 
  buttonColor, 
  icon: Icon, 
  iconGradient,
  onClick 
}) => {
  return (
    <Paper
      elevation={0}
      sx={{
        backgroundColor: 'white',
        p: 4, // p-8
        borderRadius: '24px', // rounded-3xl
        boxShadow: '0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)', // shadow-lg
        border: '1px solid rgba(229, 231, 235, 0.5)', // border-gray-100/50
        textAlign: 'center',
        height: '100%', // 确保所有卡片高度一致
        display: 'flex',
        flexDirection: 'column',
        transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)', // hover-lift
        '&:hover': {
          transform: 'translateY(-8px) scale(1.02)',
          boxShadow: '0 20px 40px rgba(0,0,0,0.1)'
        }
      }}
    >
      {/* 图标 */}
      <Box
        sx={{
          width: '64px', // w-16
          height: '64px', // h-16
          background: iconGradient,
          borderRadius: '16px', // rounded-2xl
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: 'white',
          fontSize: '1.5rem', // text-2xl
          mx: 'auto',
          mb: 3 // mb-6
        }}
      >
        <Icon sx={{ fontSize: '1.5rem' }} />
      </Box>

      {/* 标题 */}
      <Typography
        variant="h5"
        sx={{
          fontSize: '1.25rem', // text-xl
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
          lineHeight: 1.5,
          flexGrow: 1 // 让描述区域占据剩余空间
        }}
      >
        {description}
      </Typography>

      {/* 按钮 */}
      <Button
        fullWidth
        variant="contained"
        onClick={onClick}
        sx={{
          py: 1.5, // py-3
          fontSize: '1rem',
          fontWeight: 600, // font-semibold
          textTransform: 'none',
          borderRadius: '25px', // rounded-full
          backgroundColor: buttonColor,
          color: 'white',
          border: 'none',
          boxShadow: 'none',
          transition: 'all 0.3s ease',
          '&:hover': {
            backgroundColor: buttonColor,
            transform: 'scale(1.05)',
            filter: 'brightness(0.9)'
          }
        }}
      >
        {buttonText}
      </Button>
    </Paper>
  );
};

const HelpResources = () => {
  // 处理按钮点击事件
  const handleDocumentationClick = () => {
    window.open('https://docs.kapon.cloud/api/', '_blank');
  };

  const handleFAQClick = () => {
    window.open('/about', '_blank');
  };

  const handleStatusClick = () => {
    // 这里可以链接到实际的服务状态页面
    console.log('查看服务状态');
  };

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
        {/* 标题区域 */}
        <Box sx={{ textAlign: 'center', mb: 6 }}>
          <Typography
            variant="h2"
            sx={{
              fontSize: '2.25rem', // text-4xl
              fontWeight: 700, // font-bold
              color: '#1A202C', // text-primary
              mb: 3,
              textShadow: 'none'
            }}
          >
            更多帮助资源
          </Typography>
          <Typography
            variant="h6"
            sx={{
              fontSize: '1.125rem', // text-lg
              color: '#718096' // text-gray-600
            }}
          >
            在联系我们之前，您也可以先查看这些资源
          </Typography>
        </Box>

        {/* 资源卡片网格 */}
        <Grid container spacing={{ xs: 4, md: 4 }}>
          {/* 开发者文档 */}
          <Grid item xs={12} md={4}>
            <ResourceCard
              title="开发者文档"
              description="完整的API文档、代码示例和最佳实践指南"
              buttonText="查看文档"
              buttonColor="linear-gradient(135deg, #4299E1, #3182CE)"
              icon={MenuBookIcon}
              iconGradient="linear-gradient(135deg, #4299E1, #3182CE)"
              onClick={handleDocumentationClick}
            />
          </Grid>

          {/* 常见问题 */}
          <Grid item xs={12} md={4}>
            <ResourceCard
              title="常见问题"
              description="快速找到关于服务、计费和技术的常见问题答案"
              buttonText="查看FAQ"
              buttonColor="#10B981"
              icon={HelpOutlineIcon}
              iconGradient="linear-gradient(135deg, #10B981, #059669)"
              onClick={handleFAQClick}
            />
          </Grid>

          {/* 服务状态 */}
          <Grid item xs={12} md={4}>
            <ResourceCard
              title="服务状态"
              description="实时查看所有模型和服务的运行状态"
              buttonText="查看状态"
              buttonColor="#8B5CF6"
              icon={FavoriteIcon}
              iconGradient="linear-gradient(135deg, #8B5CF6, #7C3AED)"
              onClick={handleStatusClick}
            />
          </Grid>
        </Grid>
      </Container>
    </Box>
  );
};

export default HelpResources;
