import PropTypes from 'prop-types';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { useTheme } from '@mui/material/styles';
import {
  Avatar,
  Box,
  Button,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
  SwipeableDrawer,
  IconButton,
  Divider,
  Stack
} from '@mui/material';
import { Icon } from '@iconify/react';

import User1 from 'assets/images/users/user-round.svg';
import useLogin from 'hooks/useLogin';
import { useTranslation } from 'react-i18next';
import { calculateQuota } from 'utils/common';

const ProfileDrawer = ({ open, onClose }) => {
  const { t } = useTranslation();
  const theme = useTheme();
  const navigate = useNavigate();
  const { user, userGroup } = useSelector((state) => state.account);
  const { logout } = useLogin();

  const handleLogout = async () => {
    logout();
    if (onClose) onClose();
  };

  const handleNavigate = (path) => {
    navigate(path);
    if (onClose) onClose();
  };

  const formatQuota = (value) => (value ? `$${calculateQuota(value)}` : t('dashboard_index.unknown'));

  return (
    <SwipeableDrawer
      anchor="right"
      open={open}
      onClose={onClose}
      onOpen={() => {}}
      PaperProps={{
        sx: {
          width: { xs: '85%', sm: 360 },
          maxWidth: 360,
          backgroundColor: theme.palette.background.paper,
          boxShadow: theme.shadows[8]
        }
      }}
      ModalProps={{
        keepMounted: true
      }}
      sx={{ zIndex: 1280 }}
    >
      <Box sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
        <Box sx={{ p: 2, display: 'flex', justifyContent: 'flex-end' }}>
          <IconButton onClick={onClose} edge="end">
            <Icon icon="material-symbols:close" />
          </IconButton>
        </Box>

        <Box sx={{ px: 3, pb: 2 }}>
          <Stack spacing={2.5} alignItems="center">
            <Stack spacing={1.5} alignItems="center">
              <Avatar
                src={user?.avatar_url || User1}
                sx={{
                  width: 72,
                  height: 72,
                  bgcolor: theme.palette.mode === 'dark' ? theme.palette.background.default : '#ffffff',
                  border: `1px solid ${theme.palette.divider}`
                }}
              />
              <Stack spacing={0.5} alignItems="center">
                <Typography variant="h5" sx={{ fontWeight: 600 }}>
                  {user?.display_name || user?.username || 'Unknown'}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {user?.email || 'Unknown'}
                </Typography>
              </Stack>
              <Box
                sx={{
                  px: 1.75,
                  py: 0.5,
                  borderRadius: '999px',
                  backgroundColor: theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.06)' : 'rgba(25,118,210,0.06)',
                  border: `1px solid ${theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.12)' : 'rgba(25,118,210,0.16)'}`
                }}
              >
                <Typography variant="caption" color="primary">
                  {t('userPage.group')}: {userGroup?.[user?.group]?.name || user?.group}
                </Typography>
              </Box>
            </Stack>

            <Stack
              spacing={1.5}
              sx={{
                width: '100%',
                p: 2,
                borderRadius: 2,
                backgroundColor: theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.04)' : 'rgba(145,158,171,0.06)',
                border: `1px solid ${theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.08)' : 'rgba(145,158,171,0.2)'}`
              }}
            >
              <Stack direction="row" justifyContent="space-between" alignItems="center">
                <Typography variant="body2" color="text.secondary">
                  {t('modelpricePage.rate')}
                </Typography>
                <Typography variant="body2" sx={{ fontWeight: 600 }}>
                  {userGroup?.[user?.group]?.ratio || t('dashboard_index.unknown')}
                </Typography>
              </Stack>
              <Stack direction="row" justifyContent="space-between" alignItems="center">
                <Typography variant="body2" color="text.secondary">
                  {t('modelpricePage.RPM')}
                </Typography>
                <Typography variant="body2" sx={{ fontWeight: 600 }}>
                  {userGroup?.[user?.group]?.api_rate || t('dashboard_index.unknown')}
                </Typography>
              </Stack>
            </Stack>
          </Stack>
        </Box>

        <Divider />

        <Box sx={{ px: 3, py: 3 }}>
          <Stack spacing={2.5}>
            <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
              <StatItem label={t('dashboard_index.balance')} value={formatQuota(user?.quota)} />
              <StatItem label={t('dashboard_index.used')} value={formatQuota(user?.used_quota)} />
            </Stack>

            <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
              <StatItem label={t('dashboard_index.calls')} value={user?.request_count || t('dashboard_index.unknown')} />
              <StatItem label={t('invite_count')} value={user?.aff_count || t('dashboard_index.unknown')} />
            </Stack>

            <StatItem label={t('invite_reward')} value={formatQuota(user?.aff_quota)} />

            <Button
              variant="contained"
              color="primary"
              fullWidth
              sx={{ borderRadius: 2, py: 1.2, textTransform: 'none', fontWeight: 600 }}
              onClick={() => handleNavigate('/panel/topup')}
            >
              {t('topup')}
            </Button>
          </Stack>

          <Divider sx={{ my: 3 }} />

          <Box>
            <Typography variant="subtitle2" color="text.secondary" sx={{ mb: 1 }}>
              {t('dashboard_index.quickEntry', { defaultValue: '快捷入口' })}
            </Typography>
            <List disablePadding>
              <ProfileNavItem icon="mingcute:home-3-line" label={t('menu.home')} onClick={() => handleNavigate('/')} />
              <ProfileNavItem icon="mingcute:dashboard-line" label={t('dashboard')} onClick={() => handleNavigate('/panel/dashboard')} />
              <ProfileNavItem icon="mingcute:user-4-line" label={t('profile')} onClick={() => handleNavigate('/panel/profile')} />
            </List>
          </Box>
        </Box>

        <Box sx={{ mt: 'auto', p: 3, borderTop: `1px solid ${theme.palette.divider}`, bgcolor: theme.palette.background.paper }}>
          <Button
            fullWidth
            variant="outlined"
            color="error"
            sx={{
              borderRadius: 2,
              py: 1.1,
              textTransform: 'none',
              fontWeight: 600
            }}
            onClick={handleLogout}
          >
            {t('menu.signout')}
          </Button>
        </Box>
      </Box>
    </SwipeableDrawer>
  );
};

ProfileDrawer.propTypes = {
  open: PropTypes.bool,
  onClose: PropTypes.func
};

export default ProfileDrawer;

const StatItem = ({ label, value }) => {
  const theme = useTheme();

  return (
    <Box
      sx={{
        flex: 1,
        minWidth: 0,
        px: 2,
        py: 1.5,
        borderRadius: 2,
        border: `1px solid ${theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.1)' : 'rgba(145,158,171,0.2)'}`,
        backgroundColor: theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.03)' : theme.palette.background.default
      }}
    >
      <Typography variant="body2" color="text.secondary">
        {label}
      </Typography>
      <Typography variant="subtitle1" sx={{ fontWeight: 600, mt: 0.75 }}>
        {value}
      </Typography>
    </Box>
  );
};

StatItem.propTypes = {
  label: PropTypes.node,
  value: PropTypes.node
};

const ProfileNavItem = ({ icon, label, onClick }) => {
  const theme = useTheme();

  return (
    <ListItemButton
      onClick={onClick}
      sx={{
        borderRadius: 2,
        py: 1,
        px: 1.5,
        mb: 0.5,
        '&:hover': {
          backgroundColor: theme.palette.action.hover
        }
      }}
    >
      <ListItemIcon sx={{ minWidth: 40 }}>
        <Icon icon={icon} width="1.1rem" color={theme.palette.text.secondary} />
      </ListItemIcon>
      <ListItemText primary={label} />
    </ListItemButton>
  );
};

ProfileNavItem.propTypes = {
  icon: PropTypes.string.isRequired,
  label: PropTypes.node.isRequired,
  onClick: PropTypes.func
};
