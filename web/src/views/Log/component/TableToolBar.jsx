import PropTypes from 'prop-types';
import { useState, useEffect, useMemo, useCallback, useRef } from 'react';
import { useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import {
  InputAdornment,
  OutlinedInput,
  Stack,
  FormControl,
  InputLabel,
  Grid,
  Button
} from '@mui/material';
import Autocomplete from '@mui/material/Autocomplete';
import TextField from '@mui/material/TextField';
import CircularProgress from '@mui/material/CircularProgress';
import { LocalizationProvider, DatePicker } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import dayjs from 'dayjs';
import { useTranslation } from 'react-i18next';
import 'dayjs/locale/zh-cn';
import DateRangePicker from 'ui-component/DateRangePicker';
import { API } from 'utils/api';

// ----------------------------------------------------------------------

export default function TableToolBar({ filterName, handleFilterName, userIsAdmin, actions }) {
  const { t } = useTranslation();
  const theme = useTheme();
  const grey500 = theme.palette.grey[500];
  const [openMonthPicker, setOpenMonthPicker] = useState(false);
  const handleFilterNameRef = useRef(handleFilterName);

  useEffect(() => {
    handleFilterNameRef.current = handleFilterName;
  }, [handleFilterName]);

  const allOption = useMemo(
    () => ({ label: t('common.all', { defaultValue: '全部' }), value: 'ALL' }),
    [t]
  );

  const [modelOptions, setModelOptions] = useState([allOption]);
  const [channelOptions, setChannelOptions] = useState([allOption]);
  const [userOptions, setUserOptions] = useState([allOption]);

  const [loadingModels, setLoadingModels] = useState(false);
  const [loadingChannels, setLoadingChannels] = useState(false);
  const [loadingUsers, setLoadingUsers] = useState(false);

  const [selectedModels, setSelectedModels] = useState([]);
  const [selectedChannels, setSelectedChannels] = useState([]);
  const [selectedUsers, setSelectedUsers] = useState([]);

  const formatOptions = useCallback(
    (list = []) => {
      const filtered = list.filter((item) => item && item.value !== 'ALL');
      return [allOption, ...filtered];
    },
    [allOption]
  );

  const fetchOptions = useCallback(
    async (field, query, setter, setLoading) => {
      if (!field) {
        return;
      }
      if (!userIsAdmin && field === 'username') {
        setter([allOption]);
        return;
      }
      setLoading?.(true);
      try {
        const base = userIsAdmin ? '/api/log/options' : '/api/log/self/options';
        const { data } = await API.get(base, {
          params: {
            field,
            q: (query || '').trim(),
            limit: 20
          }
        });
        if (data?.success) {
          setter(formatOptions(data.data || []));
        }
      } catch (error) {
        setter((prev) => (Array.isArray(prev) && prev.length ? prev : [allOption]));
      } finally {
        setLoading?.(false);
      }
    },
    [userIsAdmin, formatOptions, allOption]
  );

  const syncMultiToFilter = useCallback((field, arr) => {
    handleFilterNameRef.current({ target: { name: field, value: arr } });
  }, []);

  useEffect(() => {
    fetchOptions('model_name', '', setModelOptions, setLoadingModels);
    fetchOptions('channel_id', '', setChannelOptions, setLoadingChannels);
    if (userIsAdmin) {
      fetchOptions('username', '', setUserOptions, setLoadingUsers);
    } else {
      setUserOptions([allOption]);
      setSelectedUsers([]);
      syncMultiToFilter('usernames', []);
      handleFilterNameRef.current({ target: { name: 'username', value: '' } });
    }
  }, [userIsAdmin, fetchOptions, allOption, syncMultiToFilter]);

  const rangeValue = useMemo(
    () => ({
      start: filterName.start_timestamp === 0 ? null : dayjs.unix(filterName.start_timestamp),
      end: filterName.end_timestamp === 0 ? null : dayjs.unix(filterName.end_timestamp)
    }),
    [filterName.start_timestamp, filterName.end_timestamp]
  );

  const onRangeChange = (value) => {
    if (!value) return;
    const startTs = value.start ? value.start.unix() : 0;
    const endTs = value.end ? value.end.unix() : 0;
    handleFilterNameRef.current({ target: { name: 'start_timestamp', value: startTs } });
    handleFilterNameRef.current({ target: { name: 'end_timestamp', value: endTs } });
  };

  const syncSelectedValues = useCallback((values, options, currentSelected, setSelected, fallbackBuilder) => {
    if (!Array.isArray(values) || values.length === 0) {
      if (currentSelected.length) {
        setSelected([]);
      }
      return;
    }
    const mapped = values
      .map((value) => options.find((opt) => opt.value === value) || fallbackBuilder(value))
      .filter(Boolean);
    const sameLength = mapped.length === currentSelected.length;
    const same =
      sameLength &&
      mapped.every((item) => currentSelected.some((selected) => selected.value === item.value));
    if (!same) {
      setSelected(mapped);
    }
  }, []);

  useEffect(() => {
    syncSelectedValues(
      filterName.model_names,
      modelOptions,
      selectedModels,
      setSelectedModels,
      (value) => ({ label: value, value })
    );
  }, [filterName.model_names, modelOptions, selectedModels, syncSelectedValues]);

  useEffect(() => {
    syncSelectedValues(
      filterName.channel_ids,
      channelOptions,
      selectedChannels,
      setSelectedChannels,
      (value) => ({ label: String(value), value: Number(value) })
    );
  }, [filterName.channel_ids, channelOptions, selectedChannels, syncSelectedValues]);

  useEffect(() => {
    syncSelectedValues(
      filterName.usernames,
      userOptions,
      selectedUsers,
      setSelectedUsers,
      (value) => ({ label: value, value })
    );
  }, [filterName.usernames, userOptions, selectedUsers, syncSelectedValues]);

  const renderAutocompleteInput = (params, label, placeholder, loading) => (
    <TextField
      {...params}
      size="small"
      label={label}
      placeholder={placeholder}
      InputProps={{
        ...params.InputProps,
        endAdornment: (
          <>
            {loading ? <CircularProgress color="inherit" size={16} /> : null}
            {params.InputProps.endAdornment}
          </>
        )
      }}
    />
  );

  return (
    <>
      {/* 第一行：时间区间 + 月份快捷按钮 + 令牌与来源IP */}
      <Grid container spacing={1.5} padding={'16px'} paddingBottom={'0px'} alignItems="center">
        <Grid item xs={12} md={6}>
          <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1} alignItems={{ xs: 'stretch', sm: 'center' }}>
            <DateRangePicker
              defaultValue={rangeValue}
              onChange={onRangeChange}
              localeText={{ start: t('tableToolBar.startTime'), end: t('tableToolBar.endTime') }}
              size="small"
            />
            <LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale={'zh-cn'}>
              <Button
                size="small"
                variant="outlined"
                color="primary"
                startIcon={<Icon icon="solar:calendar-bold-duotone" width={18} />}
                onClick={() => setOpenMonthPicker(true)}
                sx={{ whiteSpace: 'nowrap' }}
              >
                {t('common.pickMonth', { defaultValue: '选择月份' })}
              </Button>
              <DatePicker
                open={openMonthPicker}
                onClose={() => setOpenMonthPicker(false)}
                views={['year', 'month']}
                onChange={(v) => {
                  if (v) {
                    const start = dayjs(v).startOf('month').startOf('day');
                    const end = dayjs(v).endOf('month').endOf('day');
                    onRangeChange({ start, end });
                    setOpenMonthPicker(false);
                  }
                }}
                slotProps={{ textField: { style: { display: 'none' } } }}
              />
            </LocalizationProvider>
          </Stack>
        </Grid>
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
                  <Icon icon="solar:key-bold-duotone" width={20} color={grey500} />
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
                  <Icon icon="solar:user-bold-duotone" width={20} color={grey500} />
                </InputAdornment>
              }
            />
          </FormControl>
        </Grid>
      </Grid>

      {/* 第二行：多选筛选 + 操作区 */}
      <Grid container spacing={1.5} padding={'16px'}>
        {userIsAdmin && (
          <>
            <Grid item xs={12} sm={6} md={4}>
              <Autocomplete
                multiple
                options={channelOptions}
                value={selectedChannels}
                filterOptions={(options) => options}
                isOptionEqualToValue={(option, value) => option?.value === value?.value}
                getOptionLabel={(option) => option?.label || ''}
                onOpen={() => {
                  if (channelOptions.length <= 1) {
                    fetchOptions('channel_id', '', setChannelOptions, setLoadingChannels);
                  }
                }}
                onInputChange={(_, input, reason) => {
                  if (reason === 'input') {
                    fetchOptions('channel_id', input, setChannelOptions, setLoadingChannels);
                  } else if (reason === 'clear') {
                    fetchOptions('channel_id', '', setChannelOptions, setLoadingChannels);
                  }
                }}
                onChange={(_, val) => {
                  const hasAll = val?.some((item) => item.value === 'ALL');
                  const next = hasAll ? [] : val;
                  setSelectedChannels(next);
                  const numericValues = next
                    .map((item) => Number(item.value))
                    .filter((num) => !Number.isNaN(num));
                  syncMultiToFilter('channel_ids', numericValues);
                  handleFilterName({ target: { name: 'channel_id', value: '' } });
                }}
                loading={loadingChannels}
                loadingText={t('common.loading')}
                noOptionsText={t('common.noData')}
                renderInput={(params) =>
                  renderAutocompleteInput(params, t('tableToolBar.channelId'), t('tableToolBar.channelId'), loadingChannels)
                }
                disableCloseOnSelect
                fullWidth
              />
            </Grid>

            <Grid item xs={12} sm={6} md={4}>
              <Autocomplete
                multiple
                options={userOptions}
                value={selectedUsers}
                filterOptions={(options) => options}
                isOptionEqualToValue={(option, value) => option?.value === value?.value}
                getOptionLabel={(option) => option?.label || ''}
                onOpen={() => {
                  if (userOptions.length <= 1) {
                    fetchOptions('username', '', setUserOptions, setLoadingUsers);
                  }
                }}
                onInputChange={(_, input, reason) => {
                  if (reason === 'input') {
                    fetchOptions('username', input, setUserOptions, setLoadingUsers);
                  } else if (reason === 'clear') {
                    fetchOptions('username', '', setUserOptions, setLoadingUsers);
                  }
                }}
                onChange={(_, val) => {
                  const hasAll = val?.some((item) => item.value === 'ALL');
                  const next = hasAll ? [] : val;
                  setSelectedUsers(next);
                  const values = next.map((item) => item.value).filter(Boolean);
                  syncMultiToFilter('usernames', values);
                  handleFilterName({ target: { name: 'username', value: '' } });
                }}
                loading={loadingUsers}
                loadingText={t('common.loading')}
                noOptionsText={t('common.noData')}
                renderInput={(params) =>
                  renderAutocompleteInput(params, t('tableToolBar.username'), t('tableToolBar.username'), loadingUsers)
                }
                disableCloseOnSelect
                fullWidth
              />
            </Grid>
          </>
        )}

        <Grid item xs={12} sm={6} md={userIsAdmin ? 4 : 6}>
          <Autocomplete
            multiple
            options={modelOptions}
            value={selectedModels}
            filterOptions={(options) => options}
            isOptionEqualToValue={(option, value) => option?.value === value?.value}
            getOptionLabel={(option) => option?.label || ''}
            onOpen={() => {
              if (modelOptions.length <= 1) {
                fetchOptions('model_name', '', setModelOptions, setLoadingModels);
              }
            }}
            onInputChange={(_, input, reason) => {
              if (reason === 'input') {
                fetchOptions('model_name', input, setModelOptions, setLoadingModels);
              } else if (reason === 'clear') {
                fetchOptions('model_name', '', setModelOptions, setLoadingModels);
              }
            }}
            onChange={(_, val) => {
              const hasAll = val?.some((item) => item.value === 'ALL');
              const next = hasAll ? [] : val;
              setSelectedModels(next);
              const values = next.map((item) => item.value).filter(Boolean);
              syncMultiToFilter('model_names', values);
              handleFilterName({ target: { name: 'model_name', value: '' } });
            }}
            loading={loadingModels}
            loadingText={t('common.loading')}
            noOptionsText={t('common.noData')}
            renderInput={(params) =>
              renderAutocompleteInput(params, t('tableToolBar.modelName'), t('tableToolBar.modelName'), loadingModels)
            }
            disableCloseOnSelect
            fullWidth
          />
        </Grid>

        {actions && (
          <Grid
            item
            xs={12}
            md={userIsAdmin ? 12 : 12}
            sx={{ display: 'flex', justifyContent: { xs: 'flex-start', md: 'flex-end' }, alignItems: 'center' }}
          >
            <Stack direction="row" spacing={1.5} alignItems="center" sx={{ width: '100%', justifyContent: { xs: 'space-between', md: 'flex-end' } }}>
              {actions}
            </Stack>
          </Grid>
        )}
      </Grid>

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
