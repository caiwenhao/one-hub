// material-ui
import { useState } from 'react';
import { useTheme } from '@mui/material/styles';
import {
  Box,
  Button,
  IconButton,
  Typography,
  useMediaQuery
} from '@mui/material';
import LogoSection from 'layout/MainLayout/LogoSection';
import { Link } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { IconMenu2 } from '@tabler/icons-react';
import { useTranslation } from 'react-i18next';

// ==============================|| MINIMAL NAVBAR / HEADER ||============================== //

const Header = () => {
  const theme = useTheme();
  const { pathname } = useLocation();
  const account = useSelector((state) => state.account);
  const [open, setOpen] = useState(null);
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const { t } = useTranslation();

  const handleOpenMenu = (event) => {
    setOpen(open ? null : event.currentTarget);
  };

  const handleCloseMenu = () => {
    setOpen(null);
  };

  return (
    <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', width: '100%' }}>
      {/* Logo区域 - 完全对齐设计稿 */}
      <Box sx={{ display: 'flex', alignItems: 'center' }}>
        <LogoSection />
      </Box>

      {/* 中间导航菜单 - 完全对齐设计稿 */}
      <Box
        component="nav"
        sx={{
          display: { xs: 'none', md: 'flex' },
          alignItems: 'center',
          gap: 5 // space-x-10 = 40px
        }}
      >
        <Typography
          component={Link}
          to="/"
          sx={{
            color: pathname === '/' ? '#4299E1' : '#718096', // text-gray-700
            fontWeight: pathname === '/' ? 600 : 500, // font-medium
            textDecoration: 'none',
            fontSize: '1rem',
            cursor: 'pointer',
            position: 'relative',
            transition: 'all 0.3s ease', // transition-all duration-300
            '&:hover': {
              color: '#4299E1' // hover:text-accent
            },
            '&::after': {
              content: '""',
              position: 'absolute',
              left: 0,
              right: 0,
              bottom: 0,
              height: '2px', // h-0.5
              background: 'linear-gradient(to right, #4299E1, #3182CE)', // from-accent to-accent-dark
              transform: pathname === '/' ? 'scaleX(1)' : 'scaleX(0)',
              transformOrigin: 'center',
              transition: 'transform 0.3s ease' // transition-transform duration-300
            },
            '&:hover::after': {
              transform: 'scaleX(1)'
            }
          }}
        >
          首页
        </Typography>

        <Typography
          component={Link}
          to="/models"
          sx={{
            color: pathname === '/models' ? '#4299E1' : '#718096', // text-gray-700
            fontWeight: pathname === '/models' ? 600 : 500, // font-medium
            textDecoration: 'none',
            fontSize: '1rem',
            cursor: 'pointer',
            position: 'relative',
            transition: 'all 0.3s ease', // transition-all duration-300
            '&:hover': {
              color: '#4299E1' // hover:text-accent
            },
            '&::after': {
              content: '""',
              position: 'absolute',
              left: 0,
              right: 0,
              bottom: 0,
              height: '2px', // h-0.5
              background: 'linear-gradient(to right, #4299E1, #3182CE)', // from-accent to-accent-dark
              transform: pathname === '/models' ? 'scaleX(1)' : 'scaleX(0)',
              transformOrigin: 'center',
              transition: 'transform 0.3s ease' // transition-transform duration-300
            },
            '&:hover::after': {
              transform: 'scaleX(1)'
            }
          }}
        >
          热门模型
        </Typography>

        <Typography
          component={Link}
          to="/price"
          sx={{
            color: pathname === '/price' ? '#4299E1' : '#718096', // text-gray-700
            fontWeight: pathname === '/price' ? 600 : 500, // font-medium
            textDecoration: 'none',
            fontSize: '1rem',
            cursor: 'pointer',
            position: 'relative',
            transition: 'all 0.3s ease', // transition-all duration-300
            '&:hover': {
              color: '#4299E1' // hover:text-accent
            },
            '&::after': {
              content: '""',
              position: 'absolute',
              left: 0,
              right: 0,
              bottom: 0,
              height: '2px', // h-0.5
              background: 'linear-gradient(to right, #4299E1, #3182CE)', // from-accent to-accent-dark
              transform: pathname === '/price' ? 'scaleX(1)' : 'scaleX(0)',
              transformOrigin: 'center',
              transition: 'transform 0.3s ease' // transition-transform duration-300
            },
            '&:hover::after': {
              transform: 'scaleX(1)'
            }
          }}
        >
          价格方案
        </Typography>

        <Typography
          component={Link}
          to="/developer"
          sx={{
            color: pathname.startsWith('/developer') ? '#4299E1' : '#718096',
            fontWeight: pathname.startsWith('/developer') ? 600 : 500,
            textDecoration: 'none',
            fontSize: '1rem',
            cursor: 'pointer',
            position: 'relative',
            transition: 'all 0.3s ease',
            '&:hover': {
              color: '#4299E1'
            },
            '&::after': {
              content: '""',
              position: 'absolute',
              left: 0,
              right: 0,
              bottom: 0,
              height: '2px',
              background: 'linear-gradient(to right, #4299E1, #3182CE)',
              transform: pathname.startsWith('/developer') ? 'scaleX(1)' : 'scaleX(0)',
              transformOrigin: 'center',
              transition: 'transform 0.3s ease'
            },
            '&:hover::after': {
              transform: 'scaleX(1)'
            }
          }}
        >
          开发者中心
        </Typography>

        <Typography
          component={Link}
          to="/playground"
          sx={{
            color: pathname === '/playground' ? '#4299E1' : '#718096',
            fontWeight: pathname === '/playground' ? 600 : 500,
            textDecoration: 'none',
            fontSize: '1rem',
            cursor: 'pointer',
            position: 'relative',
            transition: 'all 0.3s ease',
            '&:hover': {
              color: '#4299E1'
            },
            '&::after': {
              content: '""',
              position: 'absolute',
              left: 0,
              right: 0,
              bottom: 0,
              height: '2px',
              background: 'linear-gradient(to right, #4299E1, #3182CE)',
              transform: pathname === '/playground' ? 'scaleX(1)' : 'scaleX(0)',
              transformOrigin: 'center',
              transition: 'transform 0.3s ease'
            },
            '&:hover::after': {
              transform: 'scaleX(1)'
            }
          }}
        >
          应用体验
        </Typography>

        <Typography
          component={Link}
          to="/contact"
          sx={{
            color: pathname === '/contact' ? '#4299E1' : '#718096', // text-gray-700
            fontWeight: pathname === '/contact' ? 600 : 500, // font-medium
            textDecoration: 'none',
            fontSize: '1rem',
            cursor: 'pointer',
            position: 'relative',
            transition: 'all 0.3s ease', // transition-all duration-300
            '&:hover': {
              color: '#4299E1' // hover:text-accent
            },
            '&::after': {
              content: '""',
              position: 'absolute',
              left: 0,
              right: 0,
              bottom: 0,
              height: '2px', // h-0.5
              background: 'linear-gradient(to right, #4299E1, #3182CE)', // from-accent to-accent-dark
              transform: pathname === '/contact' ? 'scaleX(1)' : 'scaleX(0)',
              transformOrigin: 'center',
              transition: 'transform 0.3s ease' // transition-transform duration-300
            },
            '&:hover::after': {
              transform: 'scaleX(1)'
            }
          }}
        >
          联系我们
        </Typography>
      </Box>

      {/* 右侧按钮区域 - 完全对齐设计稿 */}
      <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}> {/* space-x-4 */}
        {isMobile ? (
          <IconButton
            onClick={handleOpenMenu}
            sx={{
              color: theme.palette.text.primary,
              borderRadius: '12px',
              padding: '8px',
              backgroundColor: theme.palette.mode === 'dark' ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.04)',
              '&:hover': {
                backgroundColor: theme.palette.mode === 'dark' ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.08)'
              }
            }}
          >
            <IconMenu2 stroke={1.5} size="1.3rem" />
          </IconButton>
        ) : (
          <>
            {account.user ? (
              <Button
                component={Link}
                variant="contained"
                to="/panel"
                sx={{
                  px: 3, // px-6
                  py: 1.25, // py-2.5
                  fontSize: '1rem',
                  fontWeight: 500, // font-medium
                  textTransform: 'none',
                  borderRadius: '25px', // rounded-full
                  background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)', // animated-gradient
                  backgroundSize: '400% 400%',
                  color: 'white', // text-white
                  border: 'none',
                  boxShadow: 'none',
                  animation: 'gradient-shift 4s ease infinite, pulse-glow 3s ease-in-out infinite', // animated-gradient + glow-effect
                  '@keyframes gradient-shift': {
                    '0%': { backgroundPosition: '0% 50%' },
                    '50%': { backgroundPosition: '100% 50%' },
                    '100%': { backgroundPosition: '0% 50%' }
                  },
                  '@keyframes pulse-glow': {
                    '0%, 100%': {
                      boxShadow: '0 0 20px rgba(66, 153, 225, 0.3)'
                    },
                    '50%': {
                      boxShadow: '0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3)'
                    }
                  },
                  '&:hover': {
                    background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
                    transform: 'scale(1.05)', // hover:scale-105
                    transition: 'all 0.3s ease'
                  }
                }}
              >
                {t('menu.console')}
              </Button>
            ) : (
              <>
                {/* 登录按钮 - 完全对齐设计稿 */}
                <Button
                  component={Link}
                  variant="outlined"
                  to="/login"
                  sx={{
                    px: 2.5, // px-5
                    py: 1.25, // py-2.5
                    fontSize: '1rem',
                    fontWeight: 500, // font-medium
                    textTransform: 'none',
                    borderRadius: '25px', // rounded-full
                    border: '1px solid #d1d5db', // border border-gray-300
                    color: '#718096', // text-gray-700
                    backgroundColor: 'transparent',
                    transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)', // hover-lift transition
                    '&:hover': {
                      borderColor: '#4299E1', // hover:border-accent
                      color: '#4299E1', // hover:text-accent
                      backgroundColor: 'transparent',
                      transform: 'translateY(-8px) scale(1.02)', // hover-lift effect
                      boxShadow: '0 20px 40px rgba(0,0,0,0.1)'
                    }
                  }}
                >
                  登录
                </Button>

                {/* 注册按钮 - 完全对齐设计稿 */}
                <Button
                  component={Link}
                  variant="contained"
                  to="/register"
                  sx={{
                    px: 3, // px-6
                    py: 1.25, // py-2.5
                    fontSize: '1rem',
                    fontWeight: 500, // font-medium
                    textTransform: 'none',
                    borderRadius: '25px', // rounded-full
                    background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)', // animated-gradient
                    backgroundSize: '400% 400%',
                    color: 'white', // text-white
                    border: 'none',
                    boxShadow: 'none',
                    animation: 'gradient-shift 4s ease infinite, pulse-glow 3s ease-in-out infinite', // animated-gradient + glow-effect
                    '@keyframes gradient-shift': {
                      '0%': { backgroundPosition: '0% 50%' },
                      '50%': { backgroundPosition: '100% 50%' },
                      '100%': { backgroundPosition: '0% 50%' }
                    },
                    '@keyframes pulse-glow': {
                      '0%, 100%': {
                        boxShadow: '0 0 20px rgba(66, 153, 225, 0.3)'
                      },
                      '50%': {
                        boxShadow: '0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3)'
                      }
                    },
                    '&:hover': {
                      background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
                      transform: 'scale(1.05)', // hover:scale-105
                      transition: 'all 0.3s ease'
                    }
                  }}
                >
                  注册免费试用
                </Button>
              </>
            )}
          </>
        )}
      </Box>
    </Box>
  );
};

export default Header;