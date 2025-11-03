import { useMemo } from 'react';
import { useSelector } from 'react-redux';

// material-ui
import { styled, useTheme, alpha } from '@mui/material/styles';
import { Avatar, Card, CardContent, Box, Typography, Chip, LinearProgress, Stack, Tooltip, Divider } from '@mui/material';
import User1 from 'assets/images/users/user-round.svg';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { Icon } from '@iconify/react';

const CardStyle = styled(Card)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? alpha(theme.palette.background.paper, 0.85) : theme.palette.background.paper,
  border: `1px solid ${alpha(theme.palette.divider, theme.palette.mode === 'dark' ? 0.2 : 0.12)}`,
  borderRadius: 12,
  boxShadow: 'none',
  marginBottom: theme.spacing(3),
  overflow: 'hidden'
}));

const ProgressBarWrapper = styled(Box)(({ theme }) => ({
  position: 'relative',
  width: '100%',
  height: 6,
  borderRadius: 6,
  overflow: 'hidden',
  backgroundColor: alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.12 : 0.06),
  '& .MuiLinearProgress-root': {
    height: '100%',
    borderRadius: 6,
    backgroundColor: 'transparent',
    '& .MuiLinearProgress-bar': {
      borderRadius: 6
    }
  }
}));

const InfoChip = styled(Chip)(() => ({
  height: '18px',
  fontSize: '0.65rem',
  fontWeight: 600,
  borderRadius: '4px',
  '& .MuiChip-label': {
    padding: '0 6px'
  }
}));

// ==============================|| SIDEBAR MENU Card ||============================== //

const MenuCard = () => {
  const theme = useTheme();
  const { user, userGroup } = useSelector((state) => state.account);
  const navigate = useNavigate();
  const { t } = useTranslation();

  const quotaStats = useMemo(() => {
    const quotaPerUnit = Number(localStorage.getItem('quota_per_unit')) || 500000;
    const balanceValue = user?.quota ? Number(((user.quota || 0) / quotaPerUnit).toFixed(2)) : 0;
    const usedValue = user?.used_quota ? Number(((user.used_quota || 0) / quotaPerUnit).toFixed(2)) : 0;
    const total = balanceValue + usedValue;
    const percent = total > 0 ? Math.min(100, Math.round((usedValue / total) * 100)) : 0;

    return {
      balanceValue,
      usedValue,
      total,
      percent,
      requestCount: user?.request_count || 0
    };
  }, [user]);

  const usageColor = theme.palette.primary.main;
  const usageBg = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.3 : 0.15);

  return (
    <CardStyle>
      <CardContent sx={{ p: 2 }}>
        <Stack spacing={2}>
          <Stack direction="row" spacing={1.5} alignItems="center">
            <Avatar
              src={user?.avatar_url || User1}
              alt={user?.display_name || 'User Avatar'}
              sx={{
                width: 40,
                height: 40,
                bgcolor: theme.palette.mode === 'dark' ? theme.palette.background.default : '#fff',
                border: `1px solid ${alpha(theme.palette.divider, theme.palette.mode === 'dark' ? 0.3 : 0.12)}`,
                cursor: 'pointer'
              }}
              onClick={() => navigate('/panel/profile')}
            />
            <Box sx={{ flex: 1 }}>
              <Typography variant="subtitle1" sx={{ fontWeight: 600, lineHeight: 1.2, mb: 0.25 }}>
                {user ? user.display_name || 'Loading…' : 'Loading…'}
              </Typography>
              {user && userGroup && userGroup[user.group] && (
                <InfoChip
                  label={
                    <Stack direction="row" spacing={0.5} alignItems="center">
                      <Typography variant="caption" sx={{ fontSize: '0.65rem', fontWeight: 500 }}>
                        {userGroup[user.group].name} · RPM {userGroup[user.group].api_rate}
                      </Typography>
                    </Stack>
                  }
                  size="small"
                  variant="outlined"
                  color="primary"
                />
              )}
            </Box>
          </Stack>

          <Box>
            <Stack direction="row" alignItems="center" justifyContent="space-between" sx={{ mb: 0.75 }}>
              <Typography variant="body2" color="text.secondary" sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
                <Icon icon="solar:wallet-money-linear" width={14} />
                {t('sidebar.remainingBalance')}:
              </Typography>
              <Typography variant="body2" sx={{ fontWeight: 600 }}>
                ${quotaStats.balanceValue.toFixed(2)}
              </Typography>
            </Stack>
            <ProgressBarWrapper>
              <LinearProgress
                variant="determinate"
                value={quotaStats.percent}
                sx={{
                  '& .MuiLinearProgress-bar': {
                    backgroundColor: usageColor
                  }
                }}
              />
            </ProgressBarWrapper>
            <Stack direction="row" justifyContent="space-between" sx={{ mt: 0.75 }}>
              <Typography variant="caption" color="text.secondary">
                {t('dashboard_index.used')}: ${quotaStats.usedValue.toFixed(2)}
              </Typography>
              <Typography variant="caption" sx={{ color: usageColor, fontWeight: 600 }}>
                {quotaStats.percent}%
              </Typography>
            </Stack>
          </Box>

          <Divider sx={{ borderColor: alpha(theme.palette.divider, 0.4) }} />

          <Stack spacing={1}>
            <Stack direction="row" alignItems="center" justifyContent="space-between">
              <Stack direction="row" spacing={0.75} alignItems="center">
                <Box
                  sx={{
                    width: 24,
                    height: 24,
                    borderRadius: 8,
                    bgcolor: usageBg,
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center'
                  }}
                >
                  <Icon icon="solar:call-linear" width={14} color={usageColor} />
                </Box>
                <Typography variant="body2" color="text.secondary">
                  {t('dashboard_index.calls')}
                </Typography>
              </Stack>
              <Tooltip title={t('dashboard_index.calls')}>
                <Typography variant="body2" sx={{ fontWeight: 500 }}>
                  {new Intl.NumberFormat().format(quotaStats.requestCount)}
                </Typography>
              </Tooltip>
            </Stack>
          </Stack>
        </Stack>
      </CardContent>
    </CardStyle>
  );
};

export default MenuCard;
