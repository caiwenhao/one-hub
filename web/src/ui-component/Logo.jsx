// material-ui
import logoLight from 'assets/images/logo.svg';
import logoDark from 'assets/images/logo-white.svg';
import { useSelector } from 'react-redux';
import { useTheme } from '@mui/material/styles';
import { Box, Typography } from '@mui/material';

/**
 * if you want to use image instead of <svg> uncomment following.
 *
 * import logoDark from 'assets/images/logo-dark.svg';
 * import logo from 'assets/images/logo.svg';
 *
 */

// ==============================|| LOGO SVG ||============================== //

const Logo = () => {
  const siteInfo = useSelector((state) => state.siteInfo);
  const theme = useTheme();
  const defaultLogo = theme.palette.mode === 'light' ? logoLight : logoDark;

  if (siteInfo.isLoading) {
    return null; // 数据加载未完成时不显示 logo
  }

  // 如果系统名称是 Kapon AI，显示自定义的 K 字母 logo
  if (siteInfo.system_name === 'Kapon AI') {
    return (
      <Box sx={{ display: 'flex', alignItems: 'center' }}>
        <Box
          sx={{
            width: '44px',  // w-11 = 44px
            height: '44px', // h-11 = 44px
            borderRadius: '16px', // rounded-2xl
            background: 'linear-gradient(-45deg, #4299E1, #3182CE, #2B6CB0, #2A69AC)',
            backgroundSize: '400% 400%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: 'white',
            fontWeight: 'bold',
            fontSize: '1.125rem', // text-lg
            mr: 2, // mr-4 = 16px
            boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)', // shadow-lg
            animation: 'gradient-shift 4s ease infinite, pulse-glow 3s ease-in-out infinite',
            '@keyframes gradient-shift': {
              '0%': { backgroundPosition: '0% 50%' },
              '50%': { backgroundPosition: '100% 50%' },
              '100%': { backgroundPosition: '0% 50%' }
            },
            '@keyframes pulse-glow': {
              '0%, 100%': {
                boxShadow: '0 0 20px rgba(66, 153, 225, 0.3), 0 10px 15px -3px rgba(0, 0, 0, 0.1)'
              },
              '50%': {
                boxShadow: '0 0 40px rgba(66, 153, 225, 0.6), 0 0 60px rgba(66, 153, 225, 0.3), 0 10px 15px -3px rgba(0, 0, 0, 0.1)'
              }
            }
          }}
        >
          K
        </Box>
        <Typography
          component="span"
          sx={{
            color: '#1A202C', // text-primary
            fontSize: '1.25rem', // text-xl
            fontWeight: 'bold',
            letterSpacing: '-0.025em', // tracking-tight
            textShadow: '0 2px 4px rgba(0,0,0,0.1)' // text-shadow
          }}
        >
          Kapon AI
        </Typography>
      </Box>
    );
  }

  // 否则显示原有的 logo
  const logoToDisplay = siteInfo.logo ? siteInfo.logo : defaultLogo;
  return <img src={logoToDisplay} alt={siteInfo.system_name} height="50" />;
};

export default Logo;
