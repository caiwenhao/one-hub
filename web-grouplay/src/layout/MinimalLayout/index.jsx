import { Outlet } from 'react-router-dom';
import { useTheme } from '@mui/material/styles';
import { AppBar, Box, CssBaseline, Toolbar, Container, useMediaQuery } from '@mui/material';
import Header from './Header';
import Footer from 'ui-component/Footer';

// ==============================|| MINIMAL LAYOUT ||============================== //

const MinimalLayout = () => {
  const theme = useTheme();
  const matchDownSm = useMediaQuery(theme.breakpoints.down('sm'));
  const matchDownMd = useMediaQuery(theme.breakpoints.down('md'));

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <CssBaseline />
      <AppBar
        enableColorOnDark
        position="fixed"
        color="inherit"
        elevation={0}
        sx={{
          bgcolor: 'rgba(255, 255, 255, 0.95)', // bg-white/95
          backdropFilter: 'blur(12px)', // backdrop-blur-md
          borderBottom: '1px solid rgba(229, 231, 235, 0.5)', // border-b border-gray-100/50
          boxShadow: 'none',
          zIndex: theme.zIndex.drawer + 1,
          width: '100%',
          borderRadius: 0,
          transition: 'all 0.3s ease' // transition-all duration-300
        }}
      >
        <Container maxWidth="xl" sx={{ maxWidth: '1200px' }}> {/* max-w-6xl */}
          <Toolbar sx={{
            px: 3, // px-6 = 24px
            py: 2.5, // py-5 = 20px
            minHeight: '80px',
            height: '80px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between'
          }}>
            <Header />
          </Toolbar>
        </Container>
      </AppBar>
      <Box
        sx={{
          flex: '1 1 auto',
          overflow: 'auto',
          marginTop: '80px', // 匹配新的Header高度
          backgroundColor: theme.palette.background.default,
          // padding: { xs: '16px', sm: '20px', md: '24px' },
          position: 'relative',
          minHeight: `calc(100vh - 80px - ${matchDownMd ? '80px' : '60px'})`,
          scrollbarWidth: 'thin',
          '&::-webkit-scrollbar': {
            width: '8px',
            height: '8px'
          },
          '&::-webkit-scrollbar-thumb': {
            background: theme.palette.mode === 'dark' ? 'rgba(255, 255, 255, 0.2)' : 'rgba(0, 0, 0, 0.15)',
            borderRadius: '4px'
          },
          '&::-webkit-scrollbar-track': {
            background: 'transparent'
          }
        }}
      >
        <Outlet />
      </Box>
      <Box sx={{ flex: 'none', position: 'relative', zIndex: 1 }}>
        <Footer />
      </Box>
    </Box>
  );
};

export default MinimalLayout;
