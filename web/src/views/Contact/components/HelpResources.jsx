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
  // VIDU 官方文档跳转
  const handleViduDocsClick = () => {
    window.open('https://platform.vidu.cn/docs/introduction', '_blank');
  };

  // Kling 官方文档跳转
  const handleKlingDocsClick = () => {
    window.open('https://app.klingai.com/cn/dev/document-api/quickStart/productIntroduction/overview', '_blank');
  };

  // 火山方舟官方文档跳转
  const handleVolcDocsClick = () => {
    window.open('https://www.volcengine.com/docs/82379/1494384', '_blank');
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
            更多资源
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
          {/* VIDU 文档 */}
          <Grid item xs={12} md={4}>
            <ResourceCard
              title="VIDU API 文档"
              description="VIDU 平台官方 API 文档与使用指南"
              buttonText="查看文档"
              buttonColor="linear-gradient(135deg, #0EA5FF, #8B5CF6)"
              icon={MenuBookIcon}
              iconGradient="linear-gradient(135deg, #0EA5FF, #8B5CF6)"
              onClick={handleViduDocsClick}
            />
          </Grid>

          {/* Kling 文档 */}
          <Grid item xs={12} md={4}>
            <ResourceCard
              title="Kling API 文档"
              description="Kling 平台官方 API 文档与快速开始"
              buttonText="查看文档"
              buttonColor="linear-gradient(135deg, #F59E0B, #EF4444)"
              icon={HelpOutlineIcon}
              iconGradient="linear-gradient(135deg, #F59E0B, #EF4444)"
              onClick={handleKlingDocsClick}
            />
          </Grid>

          {/* 火山方舟文档 */}
          <Grid item xs={12} md={4}>
            <ResourceCard
              title="火山方舟 API 文档"
              description="火山引擎方舟大模型服务官方 API 文档"
              buttonText="查看文档"
              buttonColor="linear-gradient(135deg, #22C55E, #16A34A)"
              icon={FavoriteIcon}
              iconGradient="linear-gradient(135deg, #22C55E, #16A34A)"
              onClick={handleVolcDocsClick}
            />
          </Grid>
        </Grid>
      </Container>
    </Box>
  );
};

export default HelpResources;
