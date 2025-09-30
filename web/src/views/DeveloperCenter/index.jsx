import React, { useState } from 'react';
import {
  Box,
  Container,
  Typography,
  Button,
  Grid,
  Card,
  CardContent,
  Tab,
  Tabs,
  useTheme,
  useMediaQuery,
  Chip,
  Stack
} from '@mui/material';
import { styled, keyframes } from '@mui/system';
import {
  Book as BookIcon,
  Favorite as HeartIcon,
  Headset as HeadsetIcon,
  Code as CodeIcon,
  AccessTime as ClockIcon,
  CheckCircle as CheckIcon,
  PlayArrow as PlayIcon
} from '@mui/icons-material';
import CodeBlock from 'ui-component/CodeBlock';

// 动画定义
const float = keyframes`
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-20px) rotate(180deg); }
`;

const gradientShift = keyframes`
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
`;

// 样式化组件
const AnimatedGradientBox = styled(Box)(({ theme }) => ({
  background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
  backgroundSize: '400% 400%',
  animation: `${gradientShift} 4s ease infinite`,
  borderRadius: '16px',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  color: 'white',
  fontWeight: 'bold',
  fontSize: '18px',
  width: '44px',
  height: '44px',
  marginRight: '16px',
  boxShadow: '0 4px 12px rgba(0,0,0,0.15)'
}));

const FloatingElement = styled(Box)(({ delay = 0 }) => ({
  position: 'absolute',
  borderRadius: '50%',
  animation: `${float} 6s ease-in-out infinite`,
  animationDelay: `${delay}ms`
}));

const HoverLiftCard = styled(Card)(({ theme }) => ({
  transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
  borderRadius: '24px',
  border: '1px solid rgba(229, 231, 235, 0.5)',
  '&:hover': {
    transform: 'translateY(-8px) scale(1.02)',
    boxShadow: '0 20px 40px rgba(0,0,0,0.1)'
  }
}));

const StepCard = styled(Card)(({ theme, isActive }) => ({
  transition: 'all 0.4s cubic-bezier(0.4, 0, 0.2, 1)',
  borderRadius: '24px',
  border: isActive ? '2px solid #4299E1' : '1px solid rgba(229, 231, 235, 0.5)',
  position: 'relative',
  overflow: 'hidden',
  background: isActive ? 'linear-gradient(135deg, rgba(66, 153, 225, 0.05) 0%, rgba(255, 255, 255, 1) 100%)' : 'white',
  '&:hover': {
    transform: 'translateY(-8px) scale(1.02)',
    boxShadow: '0 20px 40px rgba(66, 153, 225, 0.15)',
    border: '2px solid #4299E1'
  },
  '&::before': {
    content: '""',
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    height: '4px',
    background: isActive ? 'linear-gradient(90deg, #4299E1, #34D399)' : 'transparent',
    transition: 'all 0.3s ease'
  }
}));

const StepNumber = styled(Box)(({ theme, isActive }) => ({
  background: isActive 
    ? 'linear-gradient(135deg, #4299E1, #34D399)' 
    : 'linear-gradient(135deg, #E2E8F0, #CBD5E0)',
  width: '80px',
  height: '80px',
  borderRadius: '50%',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  color: isActive ? 'white' : '#718096',
  fontWeight: 'bold',
  fontSize: '28px',
  marginBottom: '24px',
  margin: '0 auto 24px auto',
  boxShadow: isActive ? '0 8px 25px rgba(66, 153, 225, 0.3)' : '0 4px 15px rgba(0,0,0,0.1)',
  transition: 'all 0.3s ease',
  position: 'relative'
}));



const CodeTab = styled(Tab)(({ theme }) => ({
  transition: 'all 0.3s ease',
  textTransform: 'none',
  fontWeight: 600,
  '&.Mui-selected': {
    background: 'linear-gradient(135deg, #4299E1, #3182CE)',
    color: 'white',
    borderRadius: '8px 8px 0 0'
  }
}));

const DeveloperCenter = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const [codeTab, setCodeTab] = useState(0);

  const handleCodeTabChange = (event, newValue) => {
    setCodeTab(newValue);
  };

  const codeExamples = {
    0: `import openai

# 配置 Kapon AI
client = openai.OpenAI(
    api_key="kp-xxxxxxxxxxxxxxxx",
    base_url="https://models.kapon.cloud/v1"
)

# 发起聊天请求
response = client.chat.completions.create(
    model="gpt-4o",
    messages=[
        {"role": "user", "content": "Hello, Kapon AI!"}
    ]
)

print(response.choices[0].message.content)`,
    1: `const OpenAI = require('openai');

// 配置 Kapon AI
const client = new OpenAI({
    apiKey: 'kp-xxxxxxxxxxxxxxxx',
    baseURL: 'https://models.kapon.cloud/v1'
});

// 发起聊天请求
async function main() {
    const response = await client.chat.completions.create({
        model: 'gpt-4o',
        messages: [
            { role: 'user', content: 'Hello, Kapon AI!' }
        ]
    });
    
    console.log(response.choices[0].message.content);
}

main();`,
    2: `curl -X POST "https://models.kapon.cloud/v1/chat/completions" \\
     -H "Content-Type: application/json" \\
     -H "Authorization: Bearer kp-xxxxxxxxxxxxxxxx" \\
     -d '{
       "model": "gpt-4o",
       "messages": [
         {
           "role": "user",
           "content": "Hello, Kapon AI!"
         }
       ]
     }'`
  };

  return (
    <Box sx={{ overflow: 'hidden' }}>
      {/* Hero Section */}
      <Box
        sx={{
          position: 'relative',
          background: 'linear-gradient(135deg, #ffffff 0%, rgba(59, 130, 246, 0.05) 50%, #ffffff 100%)',
          pt: { xs: 8, md: 12 },
          pb: { xs: 8, md: 10 },
          px: { xs: 3, md: 6, lg: 12 },
          minHeight: '600px',
          display: 'flex',
          alignItems: 'center',
          overflow: 'hidden'
        }}
      >
        {/* 背景装饰元素 */}
        <Box sx={{ position: 'absolute', inset: 0, background: 'linear-gradient(135deg, rgba(66, 153, 225, 0.08) 0%, rgba(147, 51, 234, 0.05) 50%, rgba(34, 197, 94, 0.08) 100%)' }} />
        
        {/* 浮动装饰点 */}
        <FloatingElement
          delay={0}
          sx={{
            top: '80px',
            left: '40px',
            width: '12px',
            height: '12px',
            backgroundColor: 'rgba(66, 153, 225, 0.3)'
          }}
        />
        <FloatingElement
          delay={300}
          sx={{
            top: '160px',
            right: '80px',
            width: '8px',
            height: '8px',
            backgroundColor: 'rgba(34, 197, 94, 0.4)'
          }}
        />
        <FloatingElement
          delay={700}
          sx={{
            bottom: '128px',
            left: '25%',
            width: '10px',
            height: '10px',
            backgroundColor: 'rgba(147, 51, 234, 0.35)'
          }}
        />
        <FloatingElement
          delay={500}
          sx={{
            top: '240px',
            right: '33%',
            width: '6px',
            height: '6px',
            backgroundColor: 'rgba(249, 115, 22, 0.4)'
          }}
        />

        <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 10 }}>
          <Box sx={{ textAlign: 'center' }}>
            <Typography
              variant="h1"
              sx={{
                fontSize: { xs: '2.5rem', md: '4rem', lg: '4.5rem' },
                fontWeight: 200,
                color: '#1A202C',
                mb: 4,
                lineHeight: 1.2,
                letterSpacing: '-0.02em',
                textShadow: '0 2px 4px rgba(0,0,0,0.1)'
              }}
            >
              3分钟，为您的应用
              <br />
              <Box
                component="span"
                sx={{
                  fontWeight: 700,
                  background: 'linear-gradient(45deg, #4299E1 0%, #34D399 50%, #A855F7 100%)',
                  backgroundClip: 'text',
                  WebkitBackgroundClip: 'text',
                  color: 'transparent',
                  backgroundSize: '400% 400%',
                  animation: `${gradientShift} 4s ease infinite`
                }}
              >
                注入AI动力
              </Box>
            </Typography>
            
            <Typography
              variant="h5"
              sx={{
                fontSize: { xs: '1.1rem', md: '1.25rem' },
                color: '#718096',
                mb: 6,
                maxWidth: '600px',
                mx: 'auto',
                lineHeight: 1.6,
                fontWeight: 300,
                textShadow: '0 2px 4px rgba(0,0,0,0.1)'
              }}
            >
              欢迎来到 Kapon AI 开发者中心。这里有您需要的一切
              <br />
              让模型接入变得前所未有的简单
            </Typography>

            <Box sx={{ display: 'flex', flexDirection: { xs: 'column', sm: 'row' }, gap: 3, justifyContent: 'center', alignItems: 'center' }}>
              <Button
                variant="contained"
                size="large"
                component="a"
                href="/panel/token"
              sx={{
                px: 4,
                py: 2,
                fontSize: '1.1rem',
                fontWeight: 700,
                borderRadius: '25px',
                background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
                backgroundSize: '400% 400%',
                animation: `${gradientShift} 4s ease infinite`,
                textTransform: 'none',
                textDecoration: 'none',
                '&:hover': {
                  transform: 'scale(1.05)',
                    transition: 'all 0.3s ease'
                  }
                }}
              >
                立即开始集成 →
              </Button>
              
              <Box sx={{ display: 'flex', alignItems: 'center', color: '#718096' }}>
                <ClockIcon sx={{ mr: 1, fontSize: '1.2rem' }} />
                <Typography variant="body1">
                  平均集成时间: 3分钟
                </Typography>
              </Box>
            </Box>
          </Box>
        </Container>
      </Box>

      {/* Quick Start Guide */}
      <Box sx={{ 
        bgcolor: 'white', 
        py: { xs: 8, md: 12 }, 
        px: { xs: 3, md: 6, lg: 12 },
        position: 'relative',
        overflow: 'hidden'
      }}>
        {/* 背景装饰 */}
        <Box sx={{
          position: 'absolute',
          top: 0,
          left: 0,
          right: 0,
          height: '200px',
          background: 'linear-gradient(135deg, rgba(66, 153, 225, 0.03) 0%, rgba(52, 211, 153, 0.03) 100%)',
          borderRadius: '0 0 50% 50%',
          transform: 'scale(1.5)'
        }} />
        
        <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1 }}>
          <Box sx={{ textAlign: 'center', mb: 10 }}>
            <Chip
              label="快速上手"
              sx={{
                mb: 3,
                px: 2,
                py: 1,
                fontSize: '0.875rem',
                fontWeight: 600,
                background: 'linear-gradient(135deg, #4299E1, #34D399)',
                color: 'white',
                '& .MuiChip-label': { px: 2 }
              }}
            />
            <Typography
              variant="h2"
              sx={{
                fontSize: { xs: '2.5rem', md: '3.5rem' },
                fontWeight: 200,
                color: '#1A202C',
                mb: 4,
                letterSpacing: '-0.02em',
                textShadow: '0 2px 4px rgba(0,0,0,0.1)'
              }}
            >
              快速
              <Box
                component="span"
                sx={{
                  fontWeight: 700,
                  background: 'linear-gradient(45deg, #4299E1, #34D399)',
                  backgroundClip: 'text',
                  WebkitBackgroundClip: 'text',
                  color: 'transparent'
                }}
              >
                接入指南
              </Box>
            </Typography>
            <Typography
              variant="h6"
              sx={{
                color: '#718096',
                maxWidth: '600px',
                mx: 'auto',
                fontWeight: 300,
                lineHeight: 1.6
              }}
            >
              三个简单步骤，即可完成从传统API到Kapon AI的无缝切换
              <br />
              <Box component="span" sx={{ color: '#4299E1', fontWeight: 500 }}>
                平均集成时间仅需3分钟
              </Box>
            </Typography>
          </Box>

          <Grid container spacing={6} sx={{ mb: 8 }}>
            {[
              {
                step: 1,
                title: '获取API Key',
                description: '注册并登录您的Kapon AI账户，在控制台中生成专属API密钥',
                code: 'API_KEY = "kp-xxxxxxxxxxxxxxxx"',
                icon: <CodeIcon />,
                color: '#4299E1',
                isActive: true
              },
              {
                step: 2,
                title: '更换Base URL',
                description: '将您原有代码中的API端点替换为Kapon AI地址',
                code: 'base_url = "https://models.kapon.cloud/v1"',
                icon: <PlayIcon />,
                color: '#34D399'
              },
              {
                step: 3,
                title: 'API 文档',
                description: '查看完整的接口文档、参数定义与示例',
                code: 'docs.kapon.cloud/api',
                icon: <BookIcon />,
                color: '#4299E1',
                isSuccess: true,
                link: 'https://docs.kapon.cloud/api/',
                buttonText: '查看完整文档'
              }
            ].map((item, index) => (
              <Grid item xs={12} md={4} key={item.step}>
                <StepCard 
                  isActive={item.isActive}
                  sx={{ 
                    p: 5, 
                    textAlign: 'center', 
                    height: '100%',
                    position: 'relative'
                  }}
                >
                  <CardContent sx={{ p: 0 }}>
                    <StepNumber isActive={item.isActive}>
                      {item.step}
                    </StepNumber>
                    
                    <Stack spacing={3} alignItems="center">
                      <Typography
                        variant="h5"
                        sx={{
                          fontWeight: 700,
                          color: '#1A202C',
                          fontSize: '1.5rem'
                        }}
                      >
                        {item.title}
                      </Typography>
                      
                      <Typography
                        variant="body1"
                        sx={{
                          color: '#718096',
                          lineHeight: 1.7,
                          fontSize: '1rem'
                        }}
                      >
                        {item.description}
                      </Typography>
                      
                      {item.link && (
                        <Button
                          variant="contained"
                          component="a"
                          href={item.link}
                          target="_blank"
                          rel="noopener noreferrer"
                          sx={{
                            mt: 1,
                            borderRadius: '24px',
                            px: 3,
                            py: 1,
                            fontWeight: 700,
                            background: 'linear-gradient(135deg, #4299E1, #3182CE)',
                            textTransform: 'none'
                          }}
                        >
                          {item.buttonText || '查看文档'}
                        </Button>
                      )}
                    </Stack>
                  </CardContent>
                  
                  {/* 连接线 */}
                  {index < 2 && (
                    <Box
                      sx={{
                        position: 'absolute',
                        top: '50%',
                        right: { xs: 'auto', md: '-24px' },
                        bottom: { xs: '-24px', md: 'auto' },
                        left: { xs: '50%', md: 'auto' },
                        width: { xs: '2px', md: '48px' },
                        height: { xs: '48px', md: '2px' },
                        background: 'linear-gradient(90deg, #4299E1, #34D399)',
                        transform: { xs: 'translateX(-50%)', md: 'translateY(-50%)' },
                        display: { xs: 'none', md: 'block' },
                        '&::after': {
                          content: '""',
                          position: 'absolute',
                          right: '-6px',
                          top: '50%',
                          transform: 'translateY(-50%)',
                          width: 0,
                          height: 0,
                          borderLeft: '6px solid #34D399',
                          borderTop: '4px solid transparent',
                          borderBottom: '4px solid transparent'
                        }
                      }}
                    />
                  )}
                </StepCard>
              </Grid>
            ))}
          </Grid>
          
          {/* 底部CTA */}
          <Box sx={{ textAlign: 'center' }}>
            <Button
              variant="contained"
              size="large"
              component="a"
              href="/panel/token"
              startIcon={<PlayIcon />}
              sx={{
                px: 6,
                py: 2.5,
                fontSize: '1.1rem',
                fontWeight: 700,
                borderRadius: '30px',
                background: 'linear-gradient(-45deg, #4299E1, #34D399)',
                backgroundSize: '400% 400%',
                animation: `${gradientShift} 4s ease infinite`,
                textTransform: 'none',
                textDecoration: 'none',
                boxShadow: '0 8px 25px rgba(66, 153, 225, 0.3)',
                '&:hover': {
                  transform: 'translateY(-2px) scale(1.05)',
                  boxShadow: '0 12px 35px rgba(66, 153, 225, 0.4)',
                  transition: 'all 0.3s ease'
                }
              }}
            >
              立即开始集成
            </Button>
          </Box>
        </Container>
      </Box>

      {/* Core Resources section removed as requested */}

      {/* Code Samples */}
      <Box sx={{ 
        bgcolor: 'white', 
        py: { xs: 8, md: 12 }, 
        px: { xs: 3, md: 6, lg: 12 },
        position: 'relative'
      }}>
        <Container maxWidth="lg">
          <Box sx={{ textAlign: 'center', mb: 8 }}>
            <Chip
              label="代码示例"
              sx={{
                mb: 3,
                px: 2,
                py: 1,
                fontSize: '0.875rem',
                fontWeight: 600,
                background: 'linear-gradient(135deg, #A855F7, #EC4899)',
                color: 'white',
                '& .MuiChip-label': { px: 2 }
              }}
            />
            <Typography
              variant="h2"
              sx={{
                fontSize: { xs: '2.5rem', md: '3rem' },
                fontWeight: 700,
                color: '#1A202C',
                mb: 3,
                textShadow: '0 2px 4px rgba(0,0,0,0.1)'
              }}
            >
              代码示例
            </Typography>
            <Typography
              variant="h6"
              sx={{
                color: '#718096',
                fontWeight: 300,
                maxWidth: '600px',
                mx: 'auto',
                lineHeight: 1.6
              }}
            >
              复制粘贴即可使用的代码片段，支持多种编程语言
              <br />
              <Box component="span" sx={{ color: '#A855F7', fontWeight: 500 }}>
                语法高亮，一键复制
              </Box>
            </Typography>
          </Box>

          <Box
            sx={{
              bgcolor: 'white',
              borderRadius: '24px',
              boxShadow: '0 25px 50px rgba(0,0,0,0.1)',
              border: '1px solid rgba(229, 231, 235, 0.5)',
              overflow: 'hidden',
              position: 'relative'
            }}
          >
            {/* 顶部装饰条 */}
            <Box sx={{
              height: '4px',
              background: 'linear-gradient(90deg, #4299E1, #A855F7, #EC4899, #34D399)',
              backgroundSize: '400% 400%',
              animation: `${gradientShift} 6s ease infinite`
            }} />
            
            <Box sx={{ 
              borderBottom: '1px solid rgba(229, 231, 235, 1)',
              background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
            }}>
              <Tabs
                value={codeTab}
                onChange={handleCodeTabChange}
                sx={{
                  '& .MuiTabs-indicator': {
                    display: 'none'
                  },
                  px: 2,
                  pt: 1
                }}
              >
                <CodeTab
                  icon={<Box sx={{ 
                    fontSize: '1.2rem', 
                    mr: 1,
                    color: codeTab === 0 ? 'white' : '#3776ab'
                  }}>🐍</Box>}
                  iconPosition="start"
                  label="Python"
                  sx={{ 
                    px: 3, 
                    py: 2,
                    mx: 0.5,
                    borderRadius: '12px 12px 0 0',
                    fontWeight: 600
                  }}
                />
                <CodeTab
                  icon={<Box sx={{ 
                    fontSize: '1.2rem', 
                    mr: 1,
                    color: codeTab === 1 ? 'white' : '#339933'
                  }}>📗</Box>}
                  iconPosition="start"
                  label="Node.js"
                  sx={{ 
                    px: 3, 
                    py: 2,
                    mx: 0.5,
                    borderRadius: '12px 12px 0 0',
                    fontWeight: 600
                  }}
                />
                <CodeTab
                  icon={<Box sx={{ 
                    fontSize: '1.2rem', 
                    mr: 1,
                    color: codeTab === 2 ? 'white' : '#f89820'
                  }}>⚡</Box>}
                  iconPosition="start"
                  label="cURL"
                  sx={{ 
                    px: 3, 
                    py: 2,
                    mx: 0.5,
                    borderRadius: '12px 12px 0 0',
                    fontWeight: 600
                  }}
                />
              </Tabs>
            </Box>

            <Box sx={{ p: 0, position: 'relative' }}>
              <CodeBlock 
                language={['python', 'javascript', 'bash'][codeTab]}
                code={codeExamples[codeTab]}
              />
            </Box>
          </Box>

          <Box sx={{ textAlign: 'center', mt: 6 }}>
            <Stack direction={{ xs: 'column', sm: 'row' }} spacing={3} justifyContent="center" alignItems="center">
              <Button
                variant="contained"
                startIcon={<BookIcon />}
                sx={{
                  px: 5,
                  py: 2,
                  borderRadius: '25px',
                  background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
                  backgroundSize: '400% 400%',
                  animation: `${gradientShift} 4s ease infinite`,
                  fontWeight: 600,
                  textTransform: 'none',
                  boxShadow: '0 8px 25px rgba(66, 153, 225, 0.3)',
                  '&:hover': {
                    transform: 'translateY(-2px) scale(1.05)',
                    boxShadow: '0 12px 35px rgba(66, 153, 225, 0.4)',
                    transition: 'all 0.3s ease'
                  }
                }}
              >
                查看完整文档
              </Button>
              
              <Button
                variant="outlined"
                startIcon={<CodeIcon />}
                sx={{
                  px: 4,
                  py: 2,
                  borderRadius: '25px',
                  borderColor: '#A855F7',
                  color: '#A855F7',
                  fontWeight: 600,
                  textTransform: 'none',
                  '&:hover': {
                    borderColor: '#A855F7',
                    backgroundColor: 'rgba(168, 85, 247, 0.1)',
                    transform: 'translateY(-2px)',
                    transition: 'all 0.3s ease'
                  }
                }}
              >
                更多示例
              </Button>
            </Stack>
          </Box>
        </Container>
      </Box>

      {/* SDK Libraries */}
      <Box sx={{
        background: 'linear-gradient(135deg, #f9fafb 0%, #ffffff 100%)',
        py: { xs: 8, md: 10 },
        px: { xs: 3, md: 6, lg: 12 }
      }}>
        <Container maxWidth="lg">
          <Box sx={{ textAlign: 'center', mb: 6 }}>
            <Typography
              variant="h2"
              sx={{
                fontSize: { xs: '2rem', md: '2.5rem' },
                fontWeight: 700,
                color: '#1A202C',
                mb: 3,
                textShadow: '0 2px 4px rgba(0,0,0,0.1)'
              }}
            >
              SDK 与工具库
            </Typography>
            <Typography
              variant="h6"
              sx={{
                color: '#718096',
                fontWeight: 300,
                maxWidth: '600px',
                mx: 'auto',
                lineHeight: 1.6
              }}
            >
              本网站支持各家原生 API，直接使用各厂商官方 SDK 接入
              <br />
              <Box component="span" sx={{ color: '#4299E1', fontWeight: 500 }}>
                覆盖 OpenAI / Claude（Anthropic）/ 火山方舟 等主流生态
              </Box>
            </Typography>
          </Box>

          <Grid container spacing={3}>
            {[ 
              // 仅保留 Python SDK
              {
                icon: '🟩',
                title: 'OpenAI · Python SDK',
                command: 'pip install openai',
                color: '#10a37f',
                description: '本网站原生兼容 OpenAI API'
              },
              {
                icon: '🟣',
                title: 'Claude · Python SDK',
                command: 'pip install anthropic',
                color: '#8b5cf6',
                description: '本网站原生兼容 Claude API（Anthropic）'
              },
              {
                icon: '🔵',
                title: '火山方舟 · Python SDK',
                command: 'pip install volcengine-python-sdk',
                color: '#1e80ff',
                description: '本网站原生兼容火山方舟 Ark API'
              }
            ].map((sdk, index) => (
              <Grid item xs={12} sm={6} md={4} key={index}>
                <HoverLiftCard sx={{ p: 3, textAlign: 'center', height: '100%' }}>
                  <CardContent sx={{ p: 0 }}>
                    <Typography
                      sx={{
                        fontSize: '3rem',
                        mb: 2
                      }}
                    >
                      {sdk.icon}
                    </Typography>
                    <Typography
                      variant="h6"
                      sx={{
                        fontWeight: 700,
                        color: '#1A202C',
                        mb: 1
                      }}
                    >
                      {sdk.title}
                    </Typography>
                    <Typography
                      variant="body2"
                      sx={{
                        color: '#718096',
                        mb: 2,
                        fontSize: '0.875rem'
                      }}
                    >
                      {sdk.description}
                    </Typography>
                    <Box
                      sx={{
                        bgcolor: '#f8fafc',
                        border: '1px solid #e2e8f0',
                        borderRadius: '8px',
                        p: 2,
                        textAlign: 'left'
                      }}
                    >
                      <Typography
                        variant="body2"
                        sx={{
                          color: sdk.color,
                          fontFamily: 'Monaco, Menlo, Ubuntu Mono, monospace',
                          fontSize: '0.75rem',
                          wordBreak: 'break-all',
                          fontWeight: 600
                        }}
                      >
                        {sdk.command}
                      </Typography>
                    </Box>
                  </CardContent>
                </HoverLiftCard>
              </Grid>
            ))}
          </Grid>
        </Container>
      </Box>
    </Box>
  );
};

export default DeveloperCenter;
