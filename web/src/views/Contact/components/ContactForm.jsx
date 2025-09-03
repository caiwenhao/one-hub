import React, { useState } from 'react';
import {
  Box,
  Typography,
  Container,
  Paper,
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Button,
  Checkbox,
  FormControlLabel,
  Grid,
  Alert,
  Snackbar
} from '@mui/material';
import { keyframes } from '@mui/system';

// 定义动画
const gradientShift = keyframes`
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
`;

const pulseGlow = keyframes`
  0%, 100% { 
    box-shadow: 0 0 20px rgba(66, 153, 225, 0.3); 
  }
  50% { 
    box-shadow: 0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3); 
  }
`;

const ContactForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    company: '',
    email: '',
    contactType: '企业合作咨询',
    message: '',
    agreeToTerms: false
  });

  const [errors, setErrors] = useState({});
  const [showSuccess, setShowSuccess] = useState(false);

  // 通用的输入框样式
  const inputStyles = {
    '& .MuiOutlinedInput-root': {
      borderRadius: '12px',
      transition: 'all 0.3s ease',
      '& fieldset': {
        border: '2px solid',
        borderColor: 'grey.300'
      },
      '&:hover fieldset': {
        borderColor: 'grey.400'
      },
      '&.Mui-focused fieldset': {
        borderColor: 'primary.main',
        boxShadow: '0 0 0 3px rgba(66, 153, 225, 0.1)'
      }
    }
  };

  // 表单验证
  const validateForm = () => {
    const newErrors = {};

    if (!formData.name.trim()) {
      newErrors.name = '请输入您的姓名';
    }

    if (!formData.email.trim()) {
      newErrors.email = '请输入邮箱地址';
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
      newErrors.email = '请输入有效的邮箱地址';
    }

    if (!formData.message.trim()) {
      newErrors.message = '请描述您的需求';
    }

    if (!formData.agreeToTerms) {
      newErrors.agreeToTerms = '请同意隐私政策和服务条款';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  // 处理表单提交
  const handleSubmit = (e) => {
    e.preventDefault();
    
    if (validateForm()) {
      // 这里暂时只做前端展示，显示成功消息
      console.log('表单数据:', formData);
      setShowSuccess(true);
      
      // 重置表单
      setFormData({
        name: '',
        company: '',
        email: '',
        contactType: '企业合作咨询',
        message: '',
        agreeToTerms: false
      });
    }
  };

  // 处理输入变化
  const handleChange = (field) => (event) => {
    const value = field === 'agreeToTerms' ? event.target.checked : event.target.value;
    setFormData(prev => ({
      ...prev,
      [field]: value
    }));
    
    // 清除对应字段的错误
    if (errors[field]) {
      setErrors(prev => ({
        ...prev,
        [field]: ''
      }));
    }
  };

  return (
    <Box
      component="section"
      sx={{
        background: 'linear-gradient(135deg, #f9fafb 0%, white 100%)',
        py: { xs: 8, md: 10 },
        px: { xs: 3, md: 6, lg: 12 }
      }}
    >
      <Container maxWidth="md" sx={{ maxWidth: '896px' }}>
        {/* 标题区域 */}
        <Box sx={{ textAlign: 'center', mb: 6 }}>
          <Typography
            variant="h2"
            sx={{
              fontSize: '2.25rem', // text-4xl
              fontWeight: 700, // font-bold
              color: '#1A202C', // text-primary
              mb: 3,
              textShadow: '0 2px 4px rgba(0,0,0,0.1)'
            }}
          >
            发送消息给我们
          </Typography>
          <Typography
            variant="h6"
            sx={{
              fontSize: '1.125rem', // text-lg
              color: '#718096' // text-gray-600
            }}
          >
            填写下方表单，我们将尽快与您取得联系
          </Typography>
        </Box>

        {/* 表单区域 */}
        <Paper
          elevation={0}
          sx={{
            backgroundColor: 'white',
            p: 5, // p-10
            borderRadius: '24px', // rounded-3xl
            boxShadow: '0 25px 50px -12px rgba(0, 0, 0, 0.25)', // shadow-xl
            border: '1px solid rgba(229, 231, 235, 0.5)' // border-gray-100/50
          }}
        >
          <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 4 }}>
            {/* 第一行：姓名和公司 */}
            <Grid container spacing={4}>
              <Grid item xs={12} md={6}>
                <Typography
                  component="label"
                  sx={{
                    display: 'block',
                    fontSize: '0.875rem', // text-sm
                    fontWeight: 600, // font-semibold
                    color: '#1A202C', // text-primary
                    mb: 1.5
                  }}
                >
                  您的姓名 *
                </Typography>
                <TextField
                  fullWidth
                  value={formData.name}
                  onChange={handleChange('name')}
                  placeholder="请输入您的姓名"
                  error={!!errors.name}
                  helperText={errors.name}
                  sx={inputStyles}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <Typography
                  component="label"
                  sx={{
                    display: 'block',
                    fontSize: '0.875rem',
                    fontWeight: 600,
                    color: '#1A202C',
                    mb: 1.5
                  }}
                >
                  公司名称
                </Typography>
                <TextField
                  fullWidth
                  value={formData.company}
                  onChange={handleChange('company')}
                  placeholder="请输入公司名称（可选）"
                  sx={inputStyles}
                />
              </Grid>
            </Grid>

            {/* 第二行：邮箱和联系类型 */}
            <Grid container spacing={4}>
              <Grid item xs={12} md={6}>
                <Typography
                  component="label"
                  sx={{
                    display: 'block',
                    fontSize: '0.875rem',
                    fontWeight: 600,
                    color: '#1A202C',
                    mb: 1.5
                  }}
                >
                  邮箱地址 *
                </Typography>
                <TextField
                  fullWidth
                  type="email"
                  value={formData.email}
                  onChange={handleChange('email')}
                  placeholder="your@email.com"
                  error={!!errors.email}
                  helperText={errors.email}
                  sx={inputStyles}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <Typography
                  component="label"
                  sx={{
                    display: 'block',
                    fontSize: '0.875rem',
                    fontWeight: 600,
                    color: '#1A202C',
                    mb: 1.5
                  }}
                >
                  联系方式
                </Typography>
                <FormControl fullWidth>
                  <Select
                    value={formData.contactType}
                    onChange={handleChange('contactType')}
                    sx={{
                      borderRadius: '12px',
                      '& .MuiOutlinedInput-notchedOutline': {
                        border: '2px solid',
                        borderColor: 'grey.300'
                      },
                      '&:hover .MuiOutlinedInput-notchedOutline': {
                        borderColor: 'grey.400'
                      },
                      '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
                        borderColor: 'primary.main'
                      }
                    }}
                  >
                    <MenuItem value="企业合作咨询">企业合作咨询</MenuItem>
                    <MenuItem value="技术支持">技术支持</MenuItem>
                    <MenuItem value="产品建议">产品建议</MenuItem>
                    <MenuItem value="其他问题">其他问题</MenuItem>
                  </Select>
                </FormControl>
              </Grid>
            </Grid>

            {/* 需求描述 */}
            <Box>
              <Typography
                component="label"
                sx={{
                  display: 'block',
                  fontSize: '0.875rem',
                  fontWeight: 600,
                  color: '#1A202C',
                  mb: 1.5
                }}
              >
                您的需求 *
              </Typography>
              <TextField
                fullWidth
                multiline
                rows={6}
                value={formData.message}
                onChange={handleChange('message')}
                placeholder="请详细描述您的需求或问题，我们将为您提供最准确的帮助..."
                error={!!errors.message}
                helperText={errors.message}
                sx={inputStyles}
              />
            </Box>

            {/* 隐私政策同意 */}
            <Box>
              <FormControlLabel
                control={
                  <Checkbox
                    checked={formData.agreeToTerms}
                    onChange={handleChange('agreeToTerms')}
                    sx={{
                      color: 'primary.main',
                      '&.Mui-checked': {
                        color: 'primary.main'
                      }
                    }}
                  />
                }
                label={
                  <Typography sx={{ fontSize: '0.875rem', color: '#718096' }}>
                    我已阅读并同意{' '}
                    <Typography
                      component="span"
                      sx={{
                        color: '#4299E1',
                        cursor: 'pointer',
                        '&:hover': {
                          color: '#3182CE'
                        }
                      }}
                    >
                      隐私政策
                    </Typography>
                    {' '}和{' '}
                    <Typography
                      component="span"
                      sx={{
                        color: '#4299E1',
                        cursor: 'pointer',
                        '&:hover': {
                          color: '#3182CE'
                        }
                      }}
                    >
                      服务条款
                    </Typography>
                  </Typography>
                }
              />
              {errors.agreeToTerms && (
                <Typography sx={{ color: 'error.main', fontSize: '0.75rem', mt: 0.5 }}>
                  {errors.agreeToTerms}
                </Typography>
              )}
            </Box>

            {/* 提交按钮 */}
            <Box sx={{ textAlign: 'center' }}>
              <Button
                type="submit"
                variant="contained"
                sx={{
                  px: 6, // px-12
                  py: 2, // py-4
                  fontSize: '1.125rem', // text-lg
                  fontWeight: 700, // font-bold
                  textTransform: 'none',
                  borderRadius: '25px', // rounded-full
                  background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
                  backgroundSize: '400% 400%',
                  color: 'white',
                  border: 'none',
                  boxShadow: 'none',
                  animation: `${gradientShift} 4s ease infinite, ${pulseGlow} 3s ease-in-out infinite`,
                  '&:hover': {
                    background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
                    transform: 'scale(1.05)',
                    transition: 'all 0.3s ease'
                  }
                }}
              >
                发送消息 →
              </Button>
            </Box>
          </Box>
        </Paper>
      </Container>

      {/* 成功提示 */}
      <Snackbar
        open={showSuccess}
        autoHideDuration={6000}
        onClose={() => setShowSuccess(false)}
        anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
      >
        <Alert
          onClose={() => setShowSuccess(false)}
          severity="success"
          sx={{ width: '100%' }}
        >
          消息发送成功！我们将尽快与您联系。
        </Alert>
      </Snackbar>
    </Box>
  );
};

export default ContactForm;
