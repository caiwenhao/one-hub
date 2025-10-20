import PropTypes from 'prop-types';
import { useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import { InputAdornment, OutlinedInput, Stack, FormControl, InputLabel, Grid } from '@mui/material';
import { LocalizationProvider, DateTimePicker } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import dayjs from 'dayjs';
import { useTranslation } from 'react-i18next';
import 'dayjs/locale/zh-cn';

// ----------------------------------------------------------------------

export default function TableToolBar({ filterName, handleFilterName, userIsAdmin, actions }) {
  const { t } = useTranslation();
  const theme = useTheme();
  const grey500 = theme.palette.grey[500];

  return (
    <>
      {/* 第一行：令牌、模型、来源IP、起始时间、结束时间 */}
      <Grid container spacing={2} padding={'24px'} paddingBottom={'0px'}>
        <Grid item xs={12} sm={6} md={3}>
          <FormControl fullWidth>
          <InputLabel htmlFor="channel-token_name-label">{t('tableToolBar.tokenName')}</InputLabel>
          <OutlinedInput
            id="token_name"
            name="token_name"
            size="small"
            label={t('tableToolBar.tokenName')}
            value={filterName.token_name}
            onChange={handleFilterName}
            placeholder={t('tableToolBar.tokenName')}
            startAdornment={
              <InputAdornment position="start">
                <Icon icon="solar:key-bold-duotone" width="20" color={grey500} />
              </InputAdornment>
            }
          />
          </FormControl>
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <FormControl fullWidth>
          <InputLabel htmlFor="channel-model_name-label">{t('tableToolBar.modelName')}</InputLabel>
          <OutlinedInput
            id="model_name"
            name="model_name"
            size="small"
            label={t('tableToolBar.modelName')}
            value={filterName.model_name}
            onChange={handleFilterName}
            placeholder={t('tableToolBar.modelName')}
            startAdornment={
              <InputAdornment position="start">
                <Icon icon="solar:box-minimalistic-bold-duotone" width="20" color={grey500} />
              </InputAdornment>
            }
          />
          </FormControl>
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <FormControl fullWidth>
          <InputLabel htmlFor="channel-source_ip-label">{t('tableToolBar.sourceIp')}</InputLabel>
          <OutlinedInput
            id="source_ip"
            name="source_ip"
            size="small"
            label={t('tableToolBar.sourceIp')}
            value={filterName.source_ip}
            onChange={handleFilterName}
            placeholder={t('tableToolBar.sourceIp')}
            startAdornment={
              <InputAdornment position="start">
                <Icon icon="solar:user-bold-duotone" width="20" color={grey500} />
              </InputAdornment>
            }
          />
          </FormControl>
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <FormControl fullWidth>
          <LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale={'zh-cn'}>
            <DateTimePicker
              label={t('tableToolBar.startTime')}
              ampm={false}
              name="start_timestamp"
              value={filterName.start_timestamp === 0 ? null : dayjs.unix(filterName.start_timestamp)}
              onChange={(value) => {
                if (value === null) {
                  handleFilterName({ target: { name: 'start_timestamp', value: 0 } });
                  return;
                }
                handleFilterName({ target: { name: 'start_timestamp', value: value.unix() } });
              }}
              slotProps={{
                actionBar: {
                  actions: ['clear', 'today', 'accept']
                },
                textField: { size: 'small', fullWidth: true }
              }}
              sx={{ width: '100%' }}
            />
          </LocalizationProvider>
          </FormControl>
        </Grid>
        <Grid item xs={12} sm={6} md={3}>
          <FormControl fullWidth>
          <LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale={'zh-cn'}>
            <DateTimePicker
              label={t('tableToolBar.endTime')}
              name="end_timestamp"
              ampm={false}
              value={filterName.end_timestamp === 0 ? null : dayjs.unix(filterName.end_timestamp)}
              onChange={(value) => {
                if (value === null) {
                  handleFilterName({ target: { name: 'end_timestamp', value: 0 } });
                  return;
                }
                handleFilterName({ target: { name: 'end_timestamp', value: value.unix() } });
              }}
              slotProps={{
                actionBar: {
                  actions: ['clear', 'today', 'accept']
                },
                textField: { size: 'small', fullWidth: true }
              }}
              sx={{ width: '100%' }}
            />
          </LocalizationProvider>
          </FormControl>
        </Grid>
      </Grid>

      {/* 第二行：管理员可见的 渠道ID / 用户名 + 右侧操作区（刷新/搜索/列设置） */}
      <Grid container spacing={2} padding={'24px'}>
        {/* 左侧输入区 */}
        {userIsAdmin && (
          <>
            <Grid item xs={12} sm={6} md={4}>
              <FormControl fullWidth>
              <InputLabel htmlFor="channel-channel_id-label">{t('tableToolBar.channelId')}</InputLabel>
              <OutlinedInput
                id="channel_id"
                name="channel_id"
                size="small"
                label={t('tableToolBar.channelId')}
                value={filterName.channel_id}
                onChange={handleFilterName}
                placeholder={t('tableToolBar.channelId')}
                startAdornment={
                  <InputAdornment position="start">
                    <Icon icon="ph:open-ai-logo-duotone" width="20" color={grey500} />
                  </InputAdornment>
                }
              />
              </FormControl>
            </Grid>

            <Grid item xs={12} sm={6} md={4}>
              <FormControl fullWidth>
              <InputLabel htmlFor="channel-username-label">{t('tableToolBar.username')}</InputLabel>
              <OutlinedInput
                id="username"
                name="username"
                size="small"
                label={t('tableToolBar.username')}
                value={filterName.username}
                onChange={handleFilterName}
                placeholder={t('tableToolBar.username')}
                startAdornment={
                  <InputAdornment position="start">
                    <Icon icon="solar:user-bold-duotone" width="20" color={grey500} />
                  </InputAdornment>
                }
              />
              </FormControl>
            </Grid>
          </>
        )}

        {/* 右侧操作区 */}
        {actions && (
          <Grid item xs={12} md={userIsAdmin ? 4 : 12} sx={{ display: 'flex', justifyContent: { xs: 'flex-start', md: 'flex-end' }, alignItems: 'center' }}>
            <Stack direction="row" spacing={1.5} alignItems="center" sx={{ width: '100%', justifyContent: { xs: 'space-between', md: 'flex-end' } }}>
              {actions}
            </Stack>
          </Grid>
        )}
      </Grid>

      {/* 非管理员：若未渲染第二行输入，也需要把 actions 放到一行（靠右） */}
      {!userIsAdmin && actions && null}
    </>
  );
}

TableToolBar.propTypes = {
  filterName: PropTypes.object,
  handleFilterName: PropTypes.func,
  userIsAdmin: PropTypes.bool,
  actions: PropTypes.oneOfType([PropTypes.node, PropTypes.arrayOf(PropTypes.node)])
};
