import PropTypes from 'prop-types';
import { Icon } from '@iconify/react';
import { useSelector } from 'react-redux';

// material-ui
import { useTheme } from '@mui/material/styles';
import { Box, IconButton, useMediaQuery } from '@mui/material';

// project imports
import LogoSection from '../LogoSection';
import HeaderActions from 'layout/common/HeaderActions';
import { drawerWidth } from 'store/constant';

// assets
// import { Icon } from '@iconify/react';

// ==============================|| MAIN NAVBAR / HEADER ||============================== //

const Header = ({ handleLeftDrawerToggle, toggleProfileDrawer }) => {
  const theme = useTheme();
  const leftDrawerOpened = useSelector((state) => state.customization.opened);
  const isDesktop = useMediaQuery(theme.breakpoints.up('md'));
  const collapsedWidth = 72;
  const currentWidth = isDesktop ? (leftDrawerOpened ? drawerWidth : collapsedWidth) : 'auto';
  const iconColor = theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.text.primary;

  return (
    <>
      {/* logo & toggler button */}
      <Box
        sx={{
          width: currentWidth,
          display: 'flex',
          alignItems: 'center',
          [theme.breakpoints.down('md')]: {
            width: 'auto'
          }
        }}
      >
        <Box component="span" sx={{ display: { xs: 'none', md: 'block' }, flexGrow: 1 }}>
          <LogoSection />
        </Box>
        <IconButton
          size="medium"
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{
            width: 38,
            height: 38,
            borderRadius: '8px',
            backgroundColor: theme.palette.mode === 'dark' ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.03)',
            '&:hover': {
              backgroundColor: theme.palette.mode === 'dark' ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.06)'
            },
            transition: 'background-color 0.2s ease-in-out'
          }}
          onClick={handleLeftDrawerToggle}
        >
          <Icon
            icon={leftDrawerOpened ? 'tabler:layout-sidebar-right-collapse' : 'tabler:layout-sidebar-left-expand'}
            width="22px"
            height="22px"
            color={iconColor}
          />
        </IconButton>
      </Box>

      <Box sx={{ flexGrow: 1 }} />

      {/* 右侧功能按钮区（统一）：隐藏主题与语言切换 */}
      <HeaderActions showProfileWhenPanel toggleProfileDrawer={toggleProfileDrawer} showTheme={false} showI18n={false} />
    </>
  );
};

Header.propTypes = {
  handleLeftDrawerToggle: PropTypes.func,
  toggleProfileDrawer: PropTypes.func
};

export default Header;
