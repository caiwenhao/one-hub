import React, { useState, useEffect } from 'react';
import {
  AppBar,
  Toolbar,
  Box,
  Button,
  Typography,
  Container,
  IconButton,
  Drawer,
  List,
  ListItem,
  ListItemText,
  useMediaQuery,
  useTheme
} from '@mui/material';
import { Menu as MenuIcon, Close as CloseIcon } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { animationStyles } from '../../styles/animations';
import { gradients, colors } from '../../styles/gradients';
import { createGlassEffect } from '../../styles/theme';

const Header = () => {
  const [scrolled, setScrolled] = useState(false);
  const [mobileOpen, setMobileOpen] = useState(false);
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  const navigate = useNavigate();

  // 监听滚动事件
  useEffect(() => {
    const handleScroll = () => {
      const isScrolled = window.scrollY > 10;
      setScrolled(isScrolled);
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  const navigationItems = [
    { label: '热门模型', href: '/price' },
    { label: '价格方案', href: '/price' },
    { label: '开发者中心', href: '/playground' },
    { label: '联系我们', href: '/about' }
  ];

  const handleNavigation = (href) => {
    navigate(href);
    setMobileOpen(false);
  };

  const drawer = (
    <Box sx={{ width: 250, pt: 2 }}>
      <Box sx={{ display: 'flex', justifyContent: 'flex-end', px: 2, pb: 2 }}>
        <IconButton onClick={handleDrawerToggle}>
          <CloseIcon />
        </IconButton>
      </Box>
      <List>
        {navigationItems.map((item) => (
          <ListItem 
            button 
            key={item.label} 
            onClick={() => handleNavigation(item.href)}
            sx={{
              '&:hover': {
                backgroundColor: colors.accentAlpha[5]
              }
            }}
          >
            <ListItemText 
              primary={item.label} 
              sx={{ 
                '& .MuiTypography-root': { 
                  fontWeight: 500,
                  color: colors.primary
                } 
              }} 
            />
          </ListItem>
        ))}
        <ListItem sx={{ pt: 2, flexDirection: 'column', gap: 2 }}>
          <Button
            variant="outlined"
            fullWidth
            onClick={() => handleNavigation('/login')}
            sx={{
              borderRadius: '25px',
              borderColor: colors.accent,
              color: colors.accent,
              '&:hover': {
                borderColor: colors.accentDark,
                backgroundColor: colors.accentAlpha[5]
              }
            }}
          >
            登录
          </Button>
          <Button
            variant="contained"
            fullWidth
            onClick={() => handleNavigation('/register')}
            sx={{
              borderRadius: '25px',
              background: gradients.primary,
              ...animationStyles.pulseGlow,
              '&:hover': {
                background: gradients.primary,
                transform: 'scale(1.02)'
              }
            }}
          >
            注册免费试用
          </Button>
        </ListItem>
      </List>
    </Box>
  );

  return (
    <>
      <AppBar
        position="fixed"
        elevation={0}
        sx={{
          background: scrolled 
            ? 'rgba(255, 255, 255, 0.95)' 
            : 'rgba(255, 255, 255, 0.8)',
          backdropFilter: 'blur(12px)',
          borderBottom: scrolled 
            ? '1px solid rgba(0, 0, 0, 0.08)' 
            : '1px solid rgba(0, 0, 0, 0.05)',
          transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
          transform: scrolled ? 'translateY(0)' : 'translateY(0)',
        }}
      >
        <Container maxWidth="xl" sx={{ maxWidth: '1200px' }}>
          <Toolbar sx={{ py: 2.5, justifyContent: 'space-between' }}>
            {/* Logo */}
            <Box 
              sx={{ 
                display: 'flex', 
                alignItems: 'center', 
                cursor: 'pointer' 
              }}
              onClick={() => navigate('/')}
            >
              <Box
                sx={{
                  width: 44,
                  height: 44,
                  borderRadius: '16px',
                  background: gradients.primary,
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  color: 'white',
                  fontWeight: 'bold',
                  fontSize: '1.2rem',
                  mr: 2,
                  boxShadow: '0 4px 15px rgba(66, 153, 225, 0.3)',
                  ...animationStyles.pulseGlow,
                  transition: 'all 0.3s ease',
                  '&:hover': {
                    transform: 'scale(1.05)'
                  }
                }}
              >
                K
              </Box>
              <Typography
                variant="h6"
                sx={{
                  color: colors.primary,
                  fontWeight: 'bold',
                  fontSize: '1.25rem',
                  letterSpacing: '-0.5px',
                  textShadow: '0 2px 4px rgba(0,0,0,0.1)'
                }}
              >
                Kapon AI
              </Typography>
            </Box>

            {/* Desktop Navigation */}
            {!isMobile && (
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 5 }}>
                {navigationItems.map((item) => (
                  <Typography
                    key={item.label}
                    onClick={() => handleNavigation(item.href)}
                    sx={{
                      color: colors.primary,
                      fontWeight: 500,
                      cursor: 'pointer',
                      position: 'relative',
                      transition: 'all 0.3s ease',
                      '&:hover': {
                        color: colors.accent
                      },
                      '&::after': {
                        content: '""',
                        position: 'absolute',
                        bottom: '-4px',
                        left: 0,
                        right: 0,
                        height: '2px',
                        background: gradients.primary,
                        transform: 'scaleX(0)',
                        transition: 'transform 0.3s ease',
                      },
                      '&:hover::after': {
                        transform: 'scaleX(1)'
                      }
                    }}
                  >
                    {item.label}
                  </Typography>
                ))}
              </Box>
            )}

            {/* Desktop Auth Buttons */}
            {!isMobile && (
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <Typography
                  onClick={() => navigate('/login')}
                  sx={{
                    color: colors.primary,
                    fontWeight: 500,
                    cursor: 'pointer',
                    transition: 'all 0.3s ease',
                    '&:hover': {
                      color: colors.accent
                    }
                  }}
                >
                  登录
                </Typography>
                <Button
                  variant="outlined"
                  onClick={() => navigate('/login')}
                  sx={{
                    borderRadius: '25px',
                    px: 3,
                    py: 1.25,
                    borderColor: '#d1d5db',
                    color: colors.primary,
                    fontWeight: 500,
                    '&:hover': {
                      borderColor: colors.accent,
                      color: colors.accent,
                      backgroundColor: 'transparent',
                      transform: 'translateY(-1px)'
                    }
                  }}
                >
                  登录
                </Button>
                <Button
                  variant="contained"
                  onClick={() => navigate('/register')}
                  sx={{
                    borderRadius: '25px',
                    px: 3,
                    py: 1.25,
                    background: gradients.primary,
                    fontWeight: 500,
                    boxShadow: '0 4px 15px rgba(66, 153, 225, 0.3)',
                    ...animationStyles.pulseGlow,
                    '&:hover': {
                      background: gradients.primary,
                      transform: 'translateY(-2px) scale(1.02)',
                      boxShadow: '0 8px 25px rgba(66, 153, 225, 0.4)'
                    }
                  }}
                >
                  注册免费试用
                </Button>
              </Box>
            )}

            {/* Mobile Menu Button */}
            {isMobile && (
              <IconButton
                color="inherit"
                aria-label="open drawer"
                edge="start"
                onClick={handleDrawerToggle}
                sx={{ color: colors.primary }}
              >
                <MenuIcon />
              </IconButton>
            )}
          </Toolbar>
        </Container>
      </AppBar>

      {/* Mobile Drawer */}
      <Drawer
        variant="temporary"
        anchor="right"
        open={mobileOpen}
        onClose={handleDrawerToggle}
        ModalProps={{
          keepMounted: true,
        }}
        sx={{
          display: { xs: 'block', md: 'none' },
          '& .MuiDrawer-paper': {
            boxSizing: 'border-box',
            width: 250,
            background: 'rgba(255, 255, 255, 0.95)',
            backdropFilter: 'blur(12px)',
          },
        }}
      >
        {drawer}
      </Drawer>
    </>
  );
};

export default Header;
