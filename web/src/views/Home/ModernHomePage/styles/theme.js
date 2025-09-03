import { createTheme } from '@mui/material/styles';
import { colors, gradients } from './gradients';

// 扩展MUI主题以支持现代化设计
export const modernTheme = createTheme({
  palette: {
    primary: {
      main: colors.accent,
      light: colors.accentLight,
      dark: colors.accentDark,
    },
    secondary: {
      main: colors.purple,
    },
    background: {
      default: colors.background,
      paper: '#ffffff',
    },
    text: {
      primary: colors.primary,
      secondary: colors.secondary,
    },
  },
  typography: {
    fontFamily: '"Inter", "Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      fontSize: '4rem',
      fontWeight: 200,
      lineHeight: 1.2,
      letterSpacing: '-0.02em',
      '@media (max-width:768px)': {
        fontSize: '2.5rem',
      },
    },
    h2: {
      fontSize: '3rem',
      fontWeight: 200,
      lineHeight: 1.3,
      letterSpacing: '-0.01em',
      '@media (max-width:768px)': {
        fontSize: '2rem',
      },
    },
    h3: {
      fontSize: '2rem',
      fontWeight: 600,
      lineHeight: 1.4,
    },
    h4: {
      fontSize: '1.5rem',
      fontWeight: 600,
      lineHeight: 1.4,
    },
    body1: {
      fontSize: '1.125rem',
      fontWeight: 300,
      lineHeight: 1.6,
    },
    body2: {
      fontSize: '1rem',
      fontWeight: 300,
      lineHeight: 1.6,
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: '50px',
          textTransform: 'none',
          fontWeight: 600,
          padding: '12px 32px',
          fontSize: '1rem',
          transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
          '&:hover': {
            transform: 'translateY(-2px)',
            boxShadow: '0 8px 25px rgba(0,0,0,0.15)',
          },
        },
        contained: {
          background: gradients.primary,
          color: '#ffffff',
          boxShadow: '0 4px 15px rgba(66, 153, 225, 0.3)',
          '&:hover': {
            background: gradients.primary,
            boxShadow: '0 8px 25px rgba(66, 153, 225, 0.4)',
          },
        },
        outlined: {
          borderColor: colors.accent,
          color: colors.accent,
          '&:hover': {
            borderColor: colors.accentDark,
            color: colors.accentDark,
            backgroundColor: colors.accentAlpha[5],
          },
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          borderRadius: '24px',
          boxShadow: '0 4px 20px rgba(0,0,0,0.08)',
          border: '1px solid rgba(0,0,0,0.05)',
          transition: 'all 0.3s cubic-bezier(0.4, 0, 0.2, 1)',
          '&:hover': {
            transform: 'translateY(-8px)',
            boxShadow: '0 20px 40px rgba(0,0,0,0.12)',
          },
        },
      },
    },
    MuiContainer: {
      styleOverrides: {
        root: {
          '@media (min-width: 1200px)': {
            maxWidth: '1200px',
          },
        },
      },
    },
  },
  breakpoints: {
    values: {
      xs: 0,
      sm: 600,
      md: 768,
      lg: 1024,
      xl: 1200,
    },
  },
});

// 自定义样式工具函数
export const createGradientText = (gradient = gradients.textGradient) => ({
  background: gradient,
  WebkitBackgroundClip: 'text',
  WebkitTextFillColor: 'transparent',
  backgroundClip: 'text',
});

export const createGlassEffect = (opacity = 0.1) => ({
  background: `rgba(255, 255, 255, ${opacity})`,
  backdropFilter: 'blur(10px)',
  border: '1px solid rgba(255, 255, 255, 0.2)',
});

export const createFloatingElement = (size = '12px', color = colors.accentAlpha[30]) => ({
  width: size,
  height: size,
  backgroundColor: color,
  borderRadius: '50%',
  position: 'absolute',
});

export default modernTheme;
