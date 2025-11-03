import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Outlet } from 'react-router-dom';
import AuthGuard from 'utils/route-guard/AuthGuard';

// material-ui
import { styled, useTheme } from '@mui/material/styles';
import { AppBar, Box, CssBaseline, Toolbar, useMediaQuery } from '@mui/material';
import AdminContainer from 'ui-component/AdminContainer';

// project imports
import Breadcrumbs from 'ui-component/extended/Breadcrumbs';
import Header from './Header';
import Sidebar from './Sidebar';
import navigation from 'menu-items';
import { drawerWidth } from 'store/constant';
import { SET_MENU } from 'store/actions';

// assets
import { Icon } from '@iconify/react';
import ProfileDrawer from './ProfileDrawer';

// styles
export const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })(({ theme, open }) => ({
  ...theme.typography.mainContent,
  borderRadius: 0,
  backgroundColor: theme.palette.background.default,
  transition: theme.transitions.create(
    ['margin', 'width'],
    open
      ? {
          easing: theme.transitions.easing.easeOut,
          duration: theme.transitions.duration.enteringScreen
        }
      : {
          easing: theme.transitions.easing.sharp,
          duration: theme.transitions.duration.leavingScreen
        }
  ),
  flexGrow: 1,
  display: 'flex',
  flexDirection: 'column',
  paddingBottom: theme.spacing(4),
  marginTop: '64px',
  position: 'relative',
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
  },
  [theme.breakpoints.up('md')]: {
    marginLeft: 0,
    width: open ? `calc(100% - ${drawerWidth}px)` : '100%',
    paddingLeft: theme.spacing(3),
    paddingRight: theme.spacing(3)
  },
  [theme.breakpoints.down('md')]: {
    marginLeft: '0',
    width: '100%',
    padding: '16px',
    marginTop: '64px'
  },
  [theme.breakpoints.down('sm')]: {
    marginLeft: '0',
    width: '100%',
    padding: '16px',
    marginRight: '0',
    marginTop: '56px'
  }
}));

// ==============================|| MAIN LAYOUT ||============================== //

const MainLayout = () => {
  const theme = useTheme();
  const matchDownMd = useMediaQuery(theme.breakpoints.down('md'));
  // Handle left drawer
  const leftDrawerOpened = useSelector((state) => state.customization.opened);
  const dispatch = useDispatch();
  const handleLeftDrawerToggle = () => {
    dispatch({ type: SET_MENU, opened: !leftDrawerOpened });
  };

  // Profile drawer state
  const [profileDrawerOpen, setProfileDrawerOpen] = useState(false);

  const openProfileDrawer = () => {
    setProfileDrawerOpen(true);
  };

  const closeProfileDrawer = () => {
    setProfileDrawerOpen(false);
  };

  useEffect(() => {
    // 断点切换时自动收起移动端侧边栏，桌面端恢复常驻
    dispatch({ type: SET_MENU, opened: !matchDownMd });
  }, [dispatch, matchDownMd]);

  return (
    <Box
      sx={{
        display: 'flex',
        width: '100%',
        minHeight: '100vh',
        position: 'relative',
        backgroundColor: theme.palette.background.default
      }}
    >
      {/* 跳到主内容（可访问性） */}
      <Box component="a" href="#main-content" className="skip-link">跳到主内容</Box>
      <CssBaseline />
      {/* header */}
      <AppBar
        enableColorOnDark
        position="fixed"
        color="inherit"
        elevation={0}
        sx={{
          bgcolor: theme.palette.background.default,
          boxShadow: 'none',
          borderBottom: `1px solid ${theme.palette.divider}`,
          transition: leftDrawerOpened ? theme.transitions.create('width') : 'none',
          zIndex: theme.zIndex.drawer + 1,
          width: '100%',
          borderRadius: 0
        }}
      >
        <Toolbar sx={{ px: { xs: 1.5, sm: 2, md: 3 }, minHeight: { xs: 56, sm: 64 }, height: { xs: 56, sm: 64 } }}>
          <Header handleLeftDrawerToggle={handleLeftDrawerToggle} toggleProfileDrawer={openProfileDrawer} />
        </Toolbar>
      </AppBar>

      {/* drawer */}
      <Sidebar drawerOpen={leftDrawerOpened} drawerToggle={handleLeftDrawerToggle} />

      {/* main content */}
      <Main theme={theme} open={leftDrawerOpened} id="main-content">
        {/* breadcrumb */}
        <Breadcrumbs separator={<Icon icon="solar:arrow-right-linear" width="16" />} navigation={navigation} icon card={false} />
        <AuthGuard>
          <AdminContainer>
            <Outlet />
          </AdminContainer>
        </AuthGuard>
      </Main>

      {/* 用户信息抽屉 */}
      <ProfileDrawer open={profileDrawerOpen} onClose={closeProfileDrawer} />
    </Box>
  );
};

export default MainLayout;
