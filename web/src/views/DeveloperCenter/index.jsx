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
  IconButton
} from '@mui/material';
import { styled, keyframes } from '@mui/system';
import {
  Book as BookIcon,
  Favorite as HeartIcon,
  Headset as HeadsetIcon,
  Code as CodeIcon,
  AccessTime as ClockIcon
} from '@mui/icons-material';

// 动画定义
const float = keyframes`
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-20px) rotate(180deg); }
`;

const pulseGlow = keyframes`
  0%, 100% { box-shadow: 0 0 20px rgba(66, 153, 225, 0.3); }
  50% { box-shadow: 0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3); }
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
  boxShadow: '0 4px 12px rgba(0,0,0,0.15)',
  '&.glow-effect': {
    animation: `${gradientShift} 4s ease infinite, ${pulseGlow} 3s ease-in-out infinite`
  }
}));

const FloatingElement = styled(Box)(({ delay = 0 }) => ({
  position: 'absolute',
  borderRadius: '50%',
  animation: `${float} 6s ease-in-out infinite`,
  animationDelay: `${delay}ms`,
  '&.glow-effect': {
    animation: `${float} 6s ease-in-out infinite, ${pulseGlow} 3s ease-in-out infinite`,
    animationDelay: `${delay}ms`
  }
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

const StepCard = styled(Card)(({ theme }) => ({
  transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
  borderRadius: '24px',
  border: '1px solid rgba(229, 231, 235, 0.5)',
  position: 'relative',
  overflow: 'hidden',
  '&:hover': {
    transform: 'translateY(-4px)',
    boxShadow: '0 15px 35px rgba(0,0,0,0.1)'
  }
}));

const StepNumber = styled(Box)(({ theme }) => ({
  background: 'linear-gradient(135deg, #4299E1, #3182CE)',
  width: '60px',
  height: '60px',
  borderRadius: '50%',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  color: 'white',
  fontWeight: 'bold',
  fontSize: '24px',
  marginBottom: '20px',
  margin: '0 auto 20px auto'
}));

const CodeBlock = styled(Box)(({ theme }) => ({
  background: '#1a202c',
  borderRadius: '12px',
  padding: '24px',
  fontFamily: 'Monaco, Menlo, Ubuntu Mono, monospace',
  fontSize: '14px',
  lineHeight: 1.6,
  overflowX: 'auto',
  color: '#e2e8f0'
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
    base_url="https://api.kapon.ai/v1"
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
    baseURL: 'https://api.kapon.ai/v1'
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
    2: `curl -X POST "https://api.kapon.ai/v1/chat/completions" \\
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
          className="glow-effect"
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
          className="glow-effect"
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
          className="glow-effect"
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
          className="glow-effect"
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
                  animation: `${gradientShift} 4s ease infinite, ${pulseGlow} 3s ease-in-out infinite`,
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
      <Box sx={{ bgcolor: 'white', py: { xs: 8, md: 10 }, px: { xs: 3, md: 6, lg: 12 } }}>
        <Container maxWidth="lg">
          <Box sx={{ textAlign: 'center', mb: 8 }}>
            <Typography
              variant="h2"
              sx={{
                fontSize: { xs: '2rem', md: '3rem' },
                fontWeight: 200,
                color: '#1A202C',
                mb: 3,
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
                maxWidth: '500px',
                mx: 'auto',
                fontWeight: 300
              }}
            >
              三个简单步骤，即可完成从传统API到Kapon AI的无缝切换
            </Typography>
          </Box>

          <Grid container spacing={4}>
            {[
              {
                step: 1,
                title: '获取API Key',
                description: '注册并登录您的Kapon AI账户，在控制台中生成专属API密钥',
                code: 'API_KEY = "kp-xxxxxxxxxxxxxxxx"'
              },
              {
                step: 2,
                title: '更换Base URL',
                description: '将您原有代码中的API端点替换为Kapon AI地址',
                code: 'base_url = "https://api.kapon.ai/v1"'
              },
              {
                step: 3,
                title: '开始调用',
                description: '运行您的代码，享受稳定高效的AI服务体验',
                code: '✓ 连接成功，开始使用',
                isSuccess: true
              }
            ].map((item) => (
              <Grid item xs={12} md={4} key={item.step}>
                <StepCard sx={{ p: 4, textAlign: 'center', height: '100%' }}>
                  <CardContent>
                    <StepNumber className="glow-effect">
                      {item.step}
                    </StepNumber>
                    <Typography
                      variant="h5"
                      sx={{
                        fontWeight: 700,
                        color: '#1A202C',
                        mb: 2
                      }}
                    >
                      {item.title}
                    </Typography>
                    <Typography
                      variant="body1"
                      sx={{
                        color: '#718096',
                        mb: 3,
                        lineHeight: 1.6
                      }}
                    >
                      {item.description}
                    </Typography>
                    <Box
                      sx={{
                        bgcolor: '#f7fafc',
                        p: 2,
                        borderRadius: '12px',
                        textAlign: 'left'
                      }}
                    >
                      <Typography
                        component="code"
                        sx={{
                          fontSize: '0.875rem',
                          color: item.isSuccess ? '#22c55e' : '#4299E1',
                          fontFamily: 'Monaco, Menlo, Ubuntu Mono, monospace'
                        }}
                      >
                        {item.code}
                      </Typography>
                    </Box>
                  </CardContent>
                </StepCard>
              </Grid>
            ))}
          </Grid>
        </Container>
      </Box>

      {/* Core Resources */}
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
              核心开发资源
            </Typography>
            <Typography
              variant="h6"
              sx={{
                color: '#718096',
                fontWeight: 300
              }}
            >
              开发者最需要的工具和文档，一站式获取
            </Typography>
          </Box>

          <Grid container spacing={4}>
            {[
              {
                icon: <BookIcon sx={{ fontSize: '2rem' }} />,
                title: 'API 文档',
                description: '详尽的接口说明、参数定义和响应格式，让您快速上手',
                buttonText: '查看完整文档',
                buttonColor: 'linear-gradient(135deg, #4299E1, #3182CE)',
                link: 'https://docs.kapon.cloud/api/'
              },
              {
                icon: <HeartIcon sx={{ fontSize: '2rem' }} />,
                title: '服务状态',
                description: '实时监控所有模型和服务的可用性，透明展示系统状态',
                buttonText: '查看实时状态',
                buttonColor: 'linear-gradient(135deg, #22c55e, #16a34a)',
                link: '#'
              }
            ].map((resource, index) => (
              <Grid item xs={12} md={4} key={index}>
                <HoverLiftCard sx={{ p: 4, textAlign: 'center', height: '100%' }}>
                  <CardContent>
                    <Box
                      className="glow-effect"
                      sx={{
                        width: '64px',
                        height: '64px',
                        background: resource.buttonColor,
                        borderRadius: '16px',
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        color: 'white',
                        mx: 'auto',
                        mb: 3
                      }}
                    >
                      {resource.icon}
                    </Box>
                    <Typography
                      variant="h5"
                      sx={{
                        fontWeight: 700,
                        color: '#1A202C',
                        mb: 2
                      }}
                    >
                      {resource.title}
                    </Typography>
                    <Typography
                      variant="body1"
                      sx={{
                        color: '#718096',
                        mb: 3,
                        lineHeight: 1.6
                      }}
                    >
                      {resource.description}
                    </Typography>
                    <Button
                      variant="contained"
                      fullWidth
                      component="a"
                      href={resource.link}
                      target={resource.link.startsWith('http') ? '_blank' : '_self'}
                      rel={resource.link.startsWith('http') ? 'noopener noreferrer' : ''}
                      sx={{
                        py: 1.5,
                        borderRadius: '25px',
                        background: resource.buttonColor,
                        fontWeight: 600,
                        textTransform: 'none',
                        '&:hover': {
                          transform: 'scale(1.05)',
                          transition: 'all 0.3s ease'
                        }
                      }}
                    >
                      {resource.buttonText}
                    </Button>
                  </CardContent>
                </HoverLiftCard>
              </Grid>
            ))}
          </Grid>
        </Container>
      </Box>

      {/* Code Samples */}
      <Box sx={{ bgcolor: 'white', py: { xs: 8, md: 10 }, px: { xs: 3, md: 6, lg: 12 } }}>
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
              代码示例
            </Typography>
            <Typography
              variant="h6"
              sx={{
                color: '#718096',
                fontWeight: 300
              }}
            >
              复制粘贴即可使用的代码片段，支持多种编程语言
            </Typography>
          </Box>

          <Box
            sx={{
              bgcolor: 'white',
              borderRadius: '24px',
              boxShadow: '0 25px 50px rgba(0,0,0,0.1)',
              border: '1px solid rgba(229, 231, 235, 0.5)',
              overflow: 'hidden'
            }}
          >
            <Box sx={{ borderBottom: '1px solid rgba(229, 231, 235, 1)' }}>
              <Tabs
                value={codeTab}
                onChange={handleCodeTabChange}
                sx={{
                  '& .MuiTabs-indicator': {
                    display: 'none'
                  }
                }}
              >
                <CodeTab
                  icon={<CodeIcon sx={{ mr: 1 }} />}
                  iconPosition="start"
                  label="Python"
                  sx={{ px: 3, py: 2 }}
                />
                <CodeTab
                  icon={<CodeIcon sx={{ mr: 1 }} />}
                  iconPosition="start"
                  label="Node.js"
                  sx={{ px: 3, py: 2 }}
                />
                <CodeTab
                  icon={<CodeIcon sx={{ mr: 1 }} />}
                  iconPosition="start"
                  label="cURL"
                  sx={{ px: 3, py: 2 }}
                />
              </Tabs>
            </Box>

            <Box sx={{ p: 0 }}>
              <CodeBlock>
                <Typography
                  component="pre"
                  sx={{
                    margin: 0,
                    fontFamily: 'inherit',
                    fontSize: 'inherit',
                    lineHeight: 'inherit',
                    whiteSpace: 'pre-wrap',
                    wordBreak: 'break-word'
                  }}
                >
                  {codeExamples[codeTab]}
                </Typography>
              </CodeBlock>
            </Box>
          </Box>

          <Box sx={{ textAlign: 'center', mt: 4 }}>
            <Button
              variant="contained"
              sx={{
                px: 4,
                py: 1.5,
                borderRadius: '25px',
                background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
                backgroundSize: '400% 400%',
                animation: `${gradientShift} 4s ease infinite`,
                fontWeight: 600,
                textTransform: 'none',
                '&:hover': {
                  transform: 'scale(1.05)',
                  transition: 'all 0.3s ease'
                }
              }}
            >
              查看更多示例 →
            </Button>
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
                fontWeight: 300
              }}
            >
              为不同开发环境提供原生支持，让集成更加便捷
            </Typography>
          </Box>

          <Grid container spacing={3}>
            {[
              {
                icon: '🐍',
                title: 'Python SDK',
                command: 'pip install kapon-ai',
                color: '#3776ab'
              },
              {
                icon: '📗',
                title: 'Node.js SDK',
                command: 'npm install kapon-ai',
                color: '#339933'
              },
              {
                icon: '☕',
                title: 'Java SDK',
                command: 'maven install kapon-ai',
                color: '#f89820'
              },
              {
                icon: '🔷',
                title: 'Go SDK',
                command: 'go get kapon-ai',
                color: '#00add8'
              }
            ].map((sdk, index) => (
              <Grid item xs={6} md={3} key={index}>
                <HoverLiftCard sx={{ p: 3, textAlign: 'center', height: '100%' }}>
                  <CardContent>
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
                        fontFamily: 'Monaco, Menlo, Ubuntu Mono, monospace',
                        fontSize: '0.75rem'
                      }}
                    >
                      {sdk.command}
                    </Typography>
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
