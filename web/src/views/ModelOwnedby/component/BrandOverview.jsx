import PropTypes from 'prop-types';
import { useEffect, useMemo, useState } from 'react';
import {
  Box,
  Button,
  ButtonGroup,
  Card,
  Checkbox,
  Chip,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  FormControl,
  Grid,
  IconButton,
  InputAdornment,
  InputLabel,
  LinearProgress,
  MenuItem,
  Select,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TablePagination,
  TableRow,
  TextField,
  Typography
} from '@mui/material';
import PerfectScrollbar from 'react-perfect-scrollbar';
import { Icon } from '@iconify/react';
import { useTranslation } from 'react-i18next';
import { ValueFormatter } from 'utils/common';

const MAX_VISIBLE_CHANNELS = 3;

const normalizeArray = (value) => {
  if (!Array.isArray(value)) {
    return [];
  }
  return value;
};

const formatNumber = (value) => {
  if (value === null || value === undefined) {
    return '-';
  }
  const numeric = Number(value);
  if (!Number.isFinite(numeric)) {
    return value;
  }
  if (Math.abs(numeric) >= 1) {
    return numeric.toLocaleString(undefined, { maximumFractionDigits: 2 });
  }
  return numeric.toLocaleString(undefined, { maximumFractionDigits: 6 });
};

// 以 1M 价格显示（tokens类型），times 类型直接按次计费
const formatPrice = (price, t) => {
  if (!price) return '-';
  if (price.type === 'times') {
    // 按次计费：仅显示美元价（不做单位换算）
    const v = ValueFormatter(price.input, true, false);
    return t('modelOwnedby.overview.timesPrice', { value: v });
  }
  // tokens 按 1M 显示
  const input = `${ValueFormatter(price.input, true, true)} / 1M`;
  const output = `${ValueFormatter(price.output, true, true)} / 1M`;
  return t('modelOwnedby.overview.tokensPrice', { input, output });
};

const renderChannels = (channels, t) => {
  const list = normalizeArray(channels);

  if (list.length === 0) {
    return (
      <Typography variant="body2" color="text.secondary">
        {t('modelOwnedby.overview.noChannels')}
      </Typography>
    );
  }

  return (
    <Stack direction="row" flexWrap="wrap" gap={0.5}>
      {list.slice(0, MAX_VISIBLE_CHANNELS).map((channel) => (
        <Chip
          key={`${channel.id}-${channel.group}-${channel.priority}`}
          size="small"
          variant="outlined"
          label={`${channel.name} (${channel.channel_type_name || channel.channel_type})`}
        />
      ))}
      {list.length > MAX_VISIBLE_CHANNELS && (
        <Chip
          size="small"
          color="primary"
          label={t('modelOwnedby.overview.moreChannels', { count: list.length - MAX_VISIBLE_CHANNELS })}
        />
      )}
    </Stack>
  );
};

const renderGroups = (groups, t) => {
  const list = normalizeArray(groups);

  if (list.length === 0) {
    return (
      <Typography variant="body2" color="text.secondary">
        {t('modelOwnedby.overview.noGroups')}
      </Typography>
    );
  }

  return (
    <Stack direction="row" flexWrap="wrap" gap={0.5}>
      {list.map((group) => (
        <Chip key={group} label={group} size="small" variant="outlined" />
      ))}
    </Stack>
  );
};

const renderIssues = (issues, issueLabels, t) => {
  const list = normalizeArray(issues);

  if (list.length === 0) {
    return (
      <Chip size="small" color="success" variant="outlined" label={t('modelOwnedby.overview.noIssues')} />
    );
  }

  return (
    <Stack direction="row" flexWrap="wrap" gap={0.5}>
      {list.map((issue) => (
        <Chip key={issue} size="small" color="warning" variant="outlined" label={issueLabels[issue] || issue} />
      ))}
    </Stack>
  );
};

const BrandOverview = ({
  data,
  loading,
  onRefresh,
  page,
  pageSize,
  total,
  onPageChange,
  filters,
  onFiltersChange,
  ownedByOptions,
  issueStats,
  groups,
  selected,
  onSelectAll,
  onSelectOne,
  onBulkUpdate,
  onBulkUpdateChannel,
  bulkUpdating
}) => {
  const { t } = useTranslation();
  const [bulkDialogOpen, setBulkDialogOpen] = useState(false);
  const [bulkOwnedBy, setBulkOwnedBy] = useState('');
  // 批量调整价格渠道
  const [bulkChannelDialogOpen, setBulkChannelDialogOpen] = useState(false);
  const [bulkChannelType, setBulkChannelType] = useState('');
  const [keywordInput, setKeywordInput] = useState(filters.keyword || '');

  useEffect(() => {
    setKeywordInput(filters.keyword || '');
  }, [filters.keyword]);

  useEffect(() => {
    if (bulkDialogOpen && bulkOwnedBy === '' && ownedByOptions.length > 0) {
      setBulkOwnedBy(String(ownedByOptions[0].id));
    }
  }, [bulkDialogOpen, bulkOwnedBy, ownedByOptions]);

  useEffect(() => {
    if (bulkChannelDialogOpen && bulkChannelType === '' && ownedByOptions.length > 0) {
      setBulkChannelType(String(ownedByOptions[0].id));
    }
  }, [bulkChannelDialogOpen, bulkChannelType, ownedByOptions]);

  const currentModels = useMemo(() => data.map((item) => item.model), [data]);
  const selectedSet = useMemo(() => new Set(selected), [selected]);
  const allSelected = currentModels.length > 0 && currentModels.every((model) => selectedSet.has(model));
  const indeterminate = selectedSet.size > 0 && !allSelected;

  const handleSelectAllChange = (event) => {
    onSelectAll(event.target.checked, currentModels);
  };

  const handleSelectRow = (model) => (event) => {
    onSelectOne(model, event.target.checked);
  };

  const handleKeywordSearch = () => {
    onFiltersChange({ keyword: keywordInput });
  };

  const handleResetFilters = () => {
    setKeywordInput('');
    onFiltersChange({
      keyword: '',
      owned_by_type: '',
      channel_type: '',
      issue: '',
      group: ''
    });
  };

  const handlePageChange = (_event, newPage) => {
    onPageChange(newPage + 1, pageSize);
  };

  const handleRowsPerPageChange = (event) => {
    const newSize = parseInt(event.target.value, 10);
    onPageChange(1, newSize);
  };

  const handleBulkConfirm = async () => {
    if (!bulkOwnedBy) {
      return;
    }
    await onBulkUpdate(parseInt(bulkOwnedBy, 10));
    setBulkDialogOpen(false);
  };

  const handleBulkChannelConfirm = async () => {
    if (!bulkChannelType) {
      return;
    }
    await onBulkUpdateChannel(parseInt(bulkChannelType, 10));
    setBulkChannelDialogOpen(false);
  };

  const issueOptions = useMemo(() => Object.keys(issueStats || {}), [issueStats]);

  const headLabel = useMemo(
    () => [
      { id: 'model', label: t('modelOwnedby.overview.columns.model') },
      { id: 'ownedBy', label: t('modelOwnedby.overview.columns.ownedBy') },
      { id: 'channelType', label: t('modelOwnedby.overview.columns.channelType') },
      { id: 'groups', label: t('modelOwnedby.overview.columns.groups') },
      { id: 'channels', label: t('modelOwnedby.overview.columns.channels') },
      { id: 'price', label: t('modelOwnedby.overview.columns.price') },
      { id: 'issues', label: t('modelOwnedby.overview.columns.issues') }
    ],
    [t]
  );

  const issueLabels = useMemo(
    () => ({
      OWNED_BY_MISSING: t('modelOwnedby.overview.issues.OWNED_BY_MISSING'),
      PRICE_CHANNEL_MISMATCH: t('modelOwnedby.overview.issues.PRICE_CHANNEL_MISMATCH'),
      CHANNEL_BRAND_CONFLICT: t('modelOwnedby.overview.issues.CHANNEL_BRAND_CONFLICT'),
      PRICE_CHANNEL_UNKNOWN: t('modelOwnedby.overview.issues.PRICE_CHANNEL_UNKNOWN'),
      NO_CHANNEL_BOUND: t('modelOwnedby.overview.issues.NO_CHANNEL_BOUND')
    }),
    [t]
  );

  return (
    <Card sx={{ position: 'relative', overflow: 'hidden' }}>
      {loading && <LinearProgress />}

      <Box
        sx={{
          position: 'sticky',
          top: 0,
          zIndex: (theme) => theme.zIndex.appBar - 1,
          backgroundColor: 'background.paper',
          borderBottom: 1,
          borderColor: 'divider'
        }}
      >
        <Box sx={{ p: { xs: 2, md: 2.5 } }}>
          <Grid container spacing={2}>
            <Grid item xs={12} md={4} lg={3}>
              <TextField
                fullWidth
                size="small"
                label={t('modelOwnedby.overview.keyword')}
                placeholder={t('modelOwnedby.overview.keywordPlaceholder')}
                value={keywordInput}
                onChange={(event) => setKeywordInput(event.target.value)}
                onKeyDown={(event) => {
                  if (event.key === 'Enter') {
                    handleKeywordSearch();
                  }
                }}
                InputProps={{
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton edge="end" size="small" onClick={handleKeywordSearch}>
                        <Icon icon="solar:magnifier-bold-duotone" width={18} />
                      </IconButton>
                    </InputAdornment>
                  )
                }}
              />
            </Grid>
            <Grid item xs={12} sm={6} md={4} lg={3}>
              <FormControl fullWidth size="small">
                <InputLabel>{t('modelOwnedby.overview.filterOwnedBy')}</InputLabel>
                <Select
                  label={t('modelOwnedby.overview.filterOwnedBy')}
                  value={filters.owned_by_type ?? ''}
                  onChange={(event) => onFiltersChange({ owned_by_type: event.target.value })}
                >
                  <MenuItem value="">{t('modelOwnedby.overview.optionAll')}</MenuItem>
                  <MenuItem value="0">{t('modelOwnedby.overview.optionUnknown')}</MenuItem>
                  {ownedByOptions.map((option) => (
                    <MenuItem key={option.id} value={String(option.id)}>
                      {option.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={6} md={4} lg={3}>
              <FormControl fullWidth size="small">
                <InputLabel>{t('modelOwnedby.overview.filterChannelType')}</InputLabel>
                <Select
                  label={t('modelOwnedby.overview.filterChannelType')}
                  value={filters.channel_type ?? ''}
                  onChange={(event) => onFiltersChange({ channel_type: event.target.value })}
                >
                  <MenuItem value="">{t('modelOwnedby.overview.optionAll')}</MenuItem>
                  <MenuItem value="0">{t('modelOwnedby.overview.optionUnknown')}</MenuItem>
                  {ownedByOptions.map((option) => (
                    <MenuItem key={option.id} value={String(option.id)}>
                      {option.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={6} md={4} lg={3}>
              <FormControl fullWidth size="small">
                <InputLabel>{t('modelOwnedby.overview.filterIssue')}</InputLabel>
                <Select
                  label={t('modelOwnedby.overview.filterIssue')}
                  value={filters.issue ?? ''}
                  onChange={(event) => onFiltersChange({ issue: event.target.value })}
                >
                  <MenuItem value="">{t('modelOwnedby.overview.optionAll')}</MenuItem>
                  {issueOptions.map((issue) => (
                    <MenuItem key={issue} value={issue}>
                      {issueLabels[issue] || issue} ({issueStats?.[issue] ?? 0})
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={6} md={4} lg={3}>
              <FormControl fullWidth size="small">
                <InputLabel>{t('modelOwnedby.overview.filterGroup')}</InputLabel>
                <Select
                  label={t('modelOwnedby.overview.filterGroup')}
                  value={filters.group ?? ''}
                  onChange={(event) => onFiltersChange({ group: event.target.value })}
                >
                  <MenuItem value="">{t('modelOwnedby.overview.optionAll')}</MenuItem>
                  {groups.map((group) => (
                    <MenuItem key={group} value={group}>
                      {group}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} md={8} lg={9}>
              <Stack
                direction={{ xs: 'column', sm: 'row' }}
                spacing={1}
                alignItems={{ xs: 'stretch', sm: 'center' }}
              >
                <ButtonGroup variant="outlined" sx={{ flexWrap: 'wrap' }}>
                  <Button
                    onClick={handleKeywordSearch}
                    startIcon={<Icon icon="solar:magnifier-bold-duotone" width={18} />}
                  >
                    {t('modelOwnedby.overview.search')}
                  </Button>
                  <Button
                    onClick={handleResetFilters}
                    startIcon={<Icon icon="solar:refresh-bold-duotone" width={18} />}
                  >
                    {t('modelOwnedby.overview.reset')}
                  </Button>
                  <Button onClick={onRefresh} startIcon={<Icon icon="solar:sync-circle-bold-duotone" width={18} />}>
                    {t('userPage.refresh')}
                  </Button>
                </ButtonGroup>
                <Box sx={{ flexGrow: 1 }} />
                <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
                  <Button
                    variant="contained"
                    color="primary"
                    disabled={selected.length === 0}
                    startIcon={<Icon icon="solar:pen-2-bold-duotone" width={18} />}
                    onClick={() => setBulkDialogOpen(true)}
                    sx={{ alignSelf: { xs: 'stretch', sm: 'center' } }}
                  >
                    {t('modelOwnedby.overview.bulkAction', { count: selected.length })}
                  </Button>
                  <Button
                    variant="outlined"
                    color="primary"
                    disabled={selected.length === 0}
                    startIcon={<Icon icon="solar:settings-bold-duotone" width={18} />}
                    onClick={() => setBulkChannelDialogOpen(true)}
                    sx={{ alignSelf: { xs: 'stretch', sm: 'center' } }}
                  >
                    {t('modelOwnedby.overview.bulkChannelAction', { count: selected.length })}
                  </Button>
                </Stack>
              </Stack>
            </Grid>
          </Grid>
        </Box>
      </Box>

      <PerfectScrollbar component="div">
        <TableContainer sx={{ overflow: 'unset' }}>
          <Table sx={{ minWidth: 1024 }}>
            <TableHead>
              <TableRow>
                <TableCell padding="checkbox">
                  <Checkbox
                    indeterminate={indeterminate}
                    checked={allSelected}
                    onChange={handleSelectAllChange}
                  />
                </TableCell>
                {headLabel.map((head) => (
                  <TableCell key={head.id}>{head.label}</TableCell>
                ))}
              </TableRow>
            </TableHead>
            <TableBody>
              {data.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={headLabel.length + 1}>
                    <Box sx={{ p: 3, textAlign: 'center' }}>
                      <Typography variant="body2" color="text.secondary">
                        {loading ? '...' : t('modelOwnedby.overview.noData')}
                      </Typography>
                    </Box>
                  </TableCell>
                </TableRow>
              ) : (
                data.map((item) => {
                  const channels = normalizeArray(item.channels);
                  const groupsList = normalizeArray(item.groups);
                  const issues = normalizeArray(item.issues);
                  const checked = selectedSet.has(item.model);

                  return (
                    <TableRow hover key={item.model}>
                      <TableCell padding="checkbox">
                        <Checkbox checked={checked} onChange={handleSelectRow(item.model)} />
                      </TableCell>
                      <TableCell>
                        <Stack spacing={0.5}>
                          <Typography variant="subtitle1">{item.model}</Typography>
                          <Typography variant="caption" color="text.secondary">
                            {t('modelOwnedby.overview.channelCount', { count: channels.length })}
                          </Typography>
                        </Stack>
                      </TableCell>
                      <TableCell>
                        <Stack spacing={0.5}>
                          <Typography variant="body2">{item.owned_by_name || t('modelOwnedby.overview.unknown')}</Typography>
                          <Typography variant="caption" color="text.secondary">
                            {t('modelOwnedby.overview.brandId', { value: item.owned_by_type })}
                          </Typography>
                        </Stack>
                      </TableCell>
                      <TableCell>
                        <Stack spacing={0.5}>
                          <Typography variant="body2">{item.channel_type_name || t('modelOwnedby.overview.unknown')}</Typography>
                          <Typography variant="caption" color="text.secondary">
                            {t('modelOwnedby.overview.channelTypeId', { value: item.channel_type })}
                          </Typography>
                        </Stack>
                      </TableCell>
                      <TableCell>{renderGroups(groupsList, t)}</TableCell>
                      <TableCell>{renderChannels(channels, t)}</TableCell>
                      <TableCell>
                        <Typography variant="body2">{formatPrice(item.price, t)}</Typography>
                        <Typography variant="caption" color="text.secondary">
                          {item.price?.type?.toUpperCase() || t('modelOwnedby.overview.unknown')}
                        </Typography>
                      </TableCell>
                      <TableCell>{renderIssues(issues, issueLabels, t)}</TableCell>
                    </TableRow>
                  );
                })
              )}
            </TableBody>
          </Table>
        </TableContainer>
      </PerfectScrollbar>
      <TablePagination
        component="div"
        count={total}
        page={Math.max(page - 1, 0)}
        onPageChange={handlePageChange}
        rowsPerPage={pageSize}
        onRowsPerPageChange={handleRowsPerPageChange}
        rowsPerPageOptions={[10, 20, 50, 100, 200]}
      />

      <BulkAdjustDialog
        open={bulkDialogOpen}
        onClose={() => setBulkDialogOpen(false)}
        ownedByOptions={ownedByOptions}
        ownedByValue={bulkOwnedBy}
        onOwnedByChange={setBulkOwnedBy}
        onSubmit={handleBulkConfirm}
        submitting={bulkUpdating}
        selectedCount={selected.length}
        t={t}
      />

      <BulkChannelDialog
        open={bulkChannelDialogOpen}
        onClose={() => setBulkChannelDialogOpen(false)}
        ownedByOptions={ownedByOptions}
        channelTypeValue={bulkChannelType}
        onChannelTypeChange={setBulkChannelType}
        onSubmit={handleBulkChannelConfirm}
        submitting={bulkUpdating}
        selectedCount={selected.length}
        t={t}
      />
    </Card>
  );
};

const BulkAdjustDialog = ({
  open,
  onClose,
  ownedByOptions,
  ownedByValue,
  onOwnedByChange,
  onSubmit,
  submitting,
  selectedCount,
  t
}) => (
  <Dialog maxWidth="xs" fullWidth open={open} onClose={submitting ? undefined : onClose}>
    <DialogTitle>{t('modelOwnedby.overview.bulkDialogTitle')}</DialogTitle>
    <DialogContent>
      <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
        {t('modelOwnedby.overview.bulkDialogDescription', { count: selectedCount })}
      </Typography>
      <FormControl size="small" fullWidth>
        <InputLabel>{t('modelOwnedby.overview.filterOwnedBy')}</InputLabel>
        <Select
          label={t('modelOwnedby.overview.filterOwnedBy')}
          value={ownedByValue}
          onChange={(event) => onOwnedByChange(event.target.value)}
        >
          <MenuItem value="0">{t('modelOwnedby.overview.optionUnknown')}</MenuItem>
          {ownedByOptions.map((option) => (
            <MenuItem key={option.id} value={String(option.id)}>
              {option.name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </DialogContent>
    <DialogActions>
      <Button onClick={onClose} disabled={submitting}>
        {t('modelOwnedby.overview.bulkCancel')}
      </Button>
      <Button onClick={onSubmit} variant="contained" disabled={submitting || !ownedByValue}>
        {submitting && <CircularProgress size={18} sx={{ mr: 1 }} />}
        {t('modelOwnedby.overview.bulkSubmit')}
      </Button>
    </DialogActions>
  </Dialog>
);

BrandOverview.propTypes = {
  data: PropTypes.array.isRequired,
  loading: PropTypes.bool.isRequired,
  onRefresh: PropTypes.func.isRequired,
  page: PropTypes.number.isRequired,
  pageSize: PropTypes.number.isRequired,
  total: PropTypes.number.isRequired,
  onPageChange: PropTypes.func.isRequired,
  filters: PropTypes.object.isRequired,
  onFiltersChange: PropTypes.func.isRequired,
  ownedByOptions: PropTypes.array.isRequired,
  issueStats: PropTypes.object,
  groups: PropTypes.array,
  selected: PropTypes.array.isRequired,
  onSelectAll: PropTypes.func.isRequired,
  onSelectOne: PropTypes.func.isRequired,
  onBulkUpdate: PropTypes.func.isRequired,
  onBulkUpdateChannel: PropTypes.func.isRequired,
  bulkUpdating: PropTypes.bool
};

BulkAdjustDialog.propTypes = {
  open: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  ownedByOptions: PropTypes.array.isRequired,
  ownedByValue: PropTypes.string,
  onOwnedByChange: PropTypes.func.isRequired,
  onSubmit: PropTypes.func.isRequired,
  submitting: PropTypes.bool,
  selectedCount: PropTypes.number.isRequired,
  t: PropTypes.func.isRequired
};

BrandOverview.defaultProps = {
  issueStats: {},
  groups: [],
  bulkUpdating: false
};

// 批量调整价格渠道对话框
const BulkChannelDialog = ({
  open,
  onClose,
  ownedByOptions,
  channelTypeValue,
  onChannelTypeChange,
  onSubmit,
  submitting,
  selectedCount,
  t
}) => (
  <Dialog maxWidth="xs" fullWidth open={open} onClose={submitting ? undefined : onClose}>
    <DialogTitle>{t('modelOwnedby.overview.bulkChannelDialogTitle')}</DialogTitle>
    <DialogContent>
      <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
        {t('modelOwnedby.overview.bulkChannelDialogDescription', { count: selectedCount })}
      </Typography>
      <FormControl size="small" fullWidth>
        <InputLabel>{t('modelOwnedby.overview.filterChannelType')}</InputLabel>
        <Select
          label={t('modelOwnedby.overview.filterChannelType')}
          value={channelTypeValue}
          onChange={(event) => onChannelTypeChange(event.target.value)}
        >
          <MenuItem value="0">{t('modelOwnedby.overview.optionUnknown')}</MenuItem>
          {ownedByOptions.map((option) => (
            <MenuItem key={option.id} value={String(option.id)}>
              {option.name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </DialogContent>
    <DialogActions>
      <Button onClick={onClose} disabled={submitting}>
        {t('modelOwnedby.overview.bulkCancel')}
      </Button>
      <Button onClick={onSubmit} variant="contained" disabled={submitting || !channelTypeValue}>
        {submitting && <CircularProgress size={18} sx={{ mr: 1 }} />}
        {t('modelOwnedby.overview.bulkSubmit')}
      </Button>
    </DialogActions>
  </Dialog>
);

BulkChannelDialog.propTypes = {
  open: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  ownedByOptions: PropTypes.array.isRequired,
  channelTypeValue: PropTypes.string,
  onChannelTypeChange: PropTypes.func.isRequired,
  onSubmit: PropTypes.func.isRequired,
  submitting: PropTypes.bool,
  selectedCount: PropTypes.number.isRequired,
  t: PropTypes.func.isRequired
};

export default BrandOverview;
