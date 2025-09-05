import React, { useState } from 'react';
import { Box, Typography, Container, Grid, Paper, Dialog, DialogContent, IconButton } from '@mui/material';
import { keyframes } from '@mui/system';
import CodeIcon from '@mui/icons-material/Code';
import EmailIcon from '@mui/icons-material/Email';
import PhoneIcon from '@mui/icons-material/Phone';
import WechatIcon from '@mui/icons-material/Chat';
import BusinessIcon from '@mui/icons-material/Business';
import CloseIcon from '@mui/icons-material/Close';

// 定义动画
const pulseGlow = keyframes`
  0%, 100% { 
    box-shadow: 0 0 20px rgba(66, 153, 225, 0.3); 
  }
  50% { 
    box-shadow: 0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3); 
  }
`;

// 图片预览模态框组件
const ImagePreviewModal = ({ open, onClose, imageUrl, title }) => {
  return (
    <Dialog
      open={open}
      onClose={onClose}
      maxWidth="md"
      fullWidth
      PaperProps={{
        sx: {
          backgroundColor: 'rgba(255, 255, 255, 0.95)',
          backdropFilter: 'blur(10px)',
          borderRadius: '16px',
          boxShadow: '0 25px 50px -12px rgba(0, 0, 0, 0.25)'
        }
      }}
    >
      <Box sx={{ position: 'relative' }}>
        <IconButton
          onClick={onClose}
          sx={{
            position: 'absolute',
            right: 8,
            top: 8,
            zIndex: 1,
            backgroundColor: 'rgba(0, 0, 0, 0.5)',
            color: 'white',
            '&:hover': {
              backgroundColor: 'rgba(0, 0, 0, 0.7)'
            }
          }}
        >
          <CloseIcon />
        </IconButton>
        <DialogContent sx={{ p: 3, textAlign: 'center' }}>
          <Typography
            variant="h6"
            sx={{
              mb: 3,
              color: '#1A202C',
              fontWeight: 600
            }}
          >
            {title}
          </Typography>
          <img
            src={imageUrl}
            alt={title}
            style={{
              maxWidth: '100%',
              maxHeight: '500px',
              width: 'auto',
              height: 'auto',
              borderRadius: '12px',
              boxShadow: '0 10px 25px -5px rgba(0, 0, 0, 0.1)'
            }}
          />
          <Typography
            variant="body2"
            sx={{
              mt: 2,
              color: '#718096',
              fontSize: '0.875rem'
            }}
          >
            扫描二维码或长按保存到相册
          </Typography>
        </DialogContent>
      </Box>
    </Dialog>
  );
};

// 联系卡片组件
const ContactCard = ({
  title,
  description,
  email,
  phone,
  imageUrl,
  contactType = 'email', // 'email', 'phone', 'image'
  icon: Icon,
  gradientColors,
  bgGradient,
  onImageClick
}) => {
  const handleContactClick = () => {
    if (contactType === 'email' && email) {
      window.open(`mailto:${email}`, '_blank');
    } else if (contactType === 'phone' && phone) {
      window.open(`tel:${phone}`, '_blank');
    } else if (contactType === 'image' && imageUrl && onImageClick) {
      // 对于图片类型，调用图片预览功能
      onImageClick();
    }
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

      {/* 联系方式卡片 */}
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
        onClick={handleContactClick}
      >
        {contactType === 'email' && email && (
          <Typography
            sx={{
              color: gradientColors.includes('#10B981') ? '#059669' : '#4299E1',
              fontWeight: 600,
              fontSize: '1.125rem',
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
        )}

        {contactType === 'phone' && phone && (
          <Typography
            sx={{
              color: gradientColors.includes('#10B981') ? '#059669' : '#4299E1',
              fontWeight: 600,
              fontSize: '1.125rem',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              gap: 1,
              '&:hover': {
                color: gradientColors.includes('#10B981') ? '#047857' : '#3182CE'
              }
            }}
          >
            <PhoneIcon sx={{ fontSize: '1.125rem' }} />
            {phone}
          </Typography>
        )}

        {contactType === 'image' && imageUrl && (
          <Box
            sx={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              gap: 2,
              cursor: 'pointer'
            }}
            onClick={handleContactClick}
          >
            <img
              src={imageUrl}
              alt="联系方式二维码"
              style={{
                width: '120px',
                height: '120px',
                objectFit: 'contain',
                borderRadius: '8px',
                transition: 'all 0.3s ease'
              }}
            />
            <Typography
              sx={{
                color: gradientColors.includes('#10B981') ? '#059669' : '#4299E1',
                fontWeight: 600,
                fontSize: '0.875rem',
                textAlign: 'center',
                transition: 'all 0.3s ease',
                '&:hover': {
                  transform: 'scale(1.05)'
                }
              }}
            >
              点击查看大图
            </Typography>
          </Box>
        )}
      </Box>
    </Paper>
  );
};

const ContactChannels = () => {
  const [previewModal, setPreviewModal] = useState({
    open: false,
    imageUrl: '',
    title: ''
  });

  const handleImagePreview = (imageUrl, title) => {
    setPreviewModal({
      open: true,
      imageUrl,
      title
    });
  };

  const handleClosePreview = () => {
    setPreviewModal({
      open: false,
      imageUrl: '',
      title: ''
    });
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
        {/* 第一行：邮箱和电话并排 */}
        <Grid container spacing={{ xs: 6, md: 6 }} sx={{ mb: 8 }}>
          {/* 技术支持邮箱 */}
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
              contactType="email"
              icon={CodeIcon}
              gradientColors="linear-gradient(135deg, #10B981, #059669)"
              bgGradient="linear-gradient(135deg, rgba(16, 185, 129, 0.05) 0%, white 100%)"
            />
          </Grid>

          {/* 电话咨询 */}
          <Grid item xs={12} md={6}>
            <ContactCard
              title="电话咨询"
              description={
                <>
                  需要即时沟通？想要详细了解我们的服务？
                  <br />
                  欢迎直接致电，我们的客服团队随时为您服务
                </>
              }
              phone="13226413712"
              contactType="phone"
              icon={PhoneIcon}
              gradientColors="linear-gradient(135deg, #F59E0B, #D97706)"
              bgGradient="linear-gradient(135deg, rgba(245, 158, 11, 0.05) 0%, white 100%)"
            />
          </Grid>
        </Grid>

        {/* 第二行：微信公众号和企业微信 */}
        <Grid container spacing={{ xs: 6, md: 6 }}>
          {/* 微信公众号 */}
          <Grid item xs={12} md={6}>
            <ContactCard
              title="微信公众号"
              description={
                <>
                  关注我们的微信公众号，获取最新产品动态
                  <br />
                  和技术资讯，与我们保持密切联系
                </>
              }
              imageUrl="/gzh.jpg"
              contactType="image"
              icon={WechatIcon}
              gradientColors="linear-gradient(135deg, #10B981, #059669)"
              bgGradient="linear-gradient(135deg, rgba(16, 185, 129, 0.05) 0%, white 100%)"
              onImageClick={() => handleImagePreview('/gzh.jpg', '微信公众号二维码')}
            />
          </Grid>

          {/* 企业微信 */}
          <Grid item xs={12} md={6}>
            <ContactCard
              title="企业微信"
              description={
                <>
                  添加我们的企业微信，享受更专业的
                  <br />
                  一对一服务和技术支持
                </>
              }
              imageUrl="/qywx.png"
              contactType="image"
              icon={BusinessIcon}
              gradientColors="linear-gradient(135deg, #8B5CF6, #7C3AED)"
              bgGradient="linear-gradient(135deg, rgba(139, 92, 246, 0.05) 0%, white 100%)"
              onImageClick={() => handleImagePreview('/qywx.png', '企业微信二维码')}
            />
          </Grid>
        </Grid>
      </Container>

      {/* 图片预览模态框 */}
      <ImagePreviewModal
        open={previewModal.open}
        onClose={handleClosePreview}
        imageUrl={previewModal.imageUrl}
        title={previewModal.title}
      />
    </Box>
  );
};

export default ContactChannels;
