import PropTypes from 'prop-types';
import { Stack } from '@mui/material';
import { useLocation } from 'react-router-dom';
import I18nButton from 'ui-component/i18nButton';
import Profile from 'layout/MainLayout/Header/Profile';

// 统一右侧操作区（通知/主题/语言/用户抽屉）
export default function HeaderActions({
  showI18n = true,
  showProfileWhenPanel = false,
  toggleProfileDrawer
}) {
  const location = useLocation();
  const isConsoleRoute = location.pathname.startsWith('/panel');

  return (
    <Stack direction="row" spacing={{ xs: 1, md: 1.25 }} alignItems="center" sx={{ '& > *': { flexShrink: 0 } }}>
      {showI18n && <I18nButton />}
      {showProfileWhenPanel && isConsoleRoute && <Profile toggleProfileDrawer={toggleProfileDrawer} />}
    </Stack>
  );
}

HeaderActions.propTypes = {
  showI18n: PropTypes.bool,
  showProfileWhenPanel: PropTypes.bool,
  toggleProfileDrawer: PropTypes.func
};
