import { useState, useEffect, useCallback } from 'react';
import { showError, trims } from 'utils/common';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
// import PerfectScrollbar from 'react-perfect-scrollbar';
import FilterBar from 'ui-component/FilterBar';
import ScrollArea from 'ui-component/ScrollArea';
import TablePagination from '@mui/material/TablePagination';
import LinearProgress from '@mui/material/LinearProgress';
import ButtonGroup from '@mui/material/ButtonGroup';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Divider from '@mui/material/Divider';
import { Button, Card, Stack, Container, Typography, Box, Menu, MenuItem, Checkbox, ListItemText, Tabs, Tab } from '@mui/material';
import PageHeader from 'ui-component/PageHeader';
import LogTableRow from './component/TableRow';
import KeywordTableHead from 'ui-component/TableHead';
import TableToolBar from './component/TableToolBar';
import { API } from 'utils/api';
import { useIsAdmin } from 'utils/common';
import { PAGE_SIZE_OPTIONS, getPageSize, savePageSize } from 'constants';
import { Icon } from '@iconify/react';
import dayjs from 'dayjs';
import { useTranslation } from 'react-i18next';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import { useSelector } from 'react-redux';
import { useLogType } from './type/LogType';

export default function Log() {
  const { t } = useTranslation();
  const LogType = useLogType();
  const originalKeyword = {
    p: 0,
    username: '',
    usernames: [],
    token_name: '',
    model_name: '',
    model_names: [],
    start_timestamp: 0,
    end_timestamp: dayjs().unix() + 3600,
    log_type: '0',
    channel_id: '',
    channel_ids: [],
    source_ip: ''
  };

  const [page, setPage] = useState(0);
  const [order, setOrder] = useState('desc');
  const [orderBy, setOrderBy] = useState('created_at');
  const [rowsPerPage, setRowsPerPage] = useState(() => getPageSize('log'));
  const [listCount, setListCount] = useState(0);
  const [searching, setSearching] = useState(false);
  const [toolBarValue, setToolBarValue] = useState({ ...originalKeyword });
  const [searchKeyword, setSearchKeyword] = useState({ ...originalKeyword });
  const [refreshFlag, setRefreshFlag] = useState(false);
  const { userGroup } = useSelector((state) => state.account);
  const theme = useTheme();
  const matchUpMd = useMediaQuery(theme.breakpoints.up('sm'));

  const [logs, setLogs] = useState([]);
  const userIsAdmin = useIsAdmin();

  // 添加列显示设置相关状态
  const [columnVisibility, setColumnVisibility] = useState({
    created_at: true,
    channel_id: true,
    user_id: true,
    group: true,
    token_name: true,
    type: true,
    model_name: true,
    duration: true,
    message: true,
    completion: true,
    quota: true,
    source_ip: true,
    detail: true
  });
  const [columnMenuAnchor, setColumnMenuAnchor] = useState(null);

  // 处理列显示菜单打开和关闭
  const handleColumnMenuOpen = (event) => {
    setColumnMenuAnchor(event.currentTarget);
  };

  const handleColumnMenuClose = () => {
    setColumnMenuAnchor(null);
  };

  // 处理列显示状态变更
  const handleColumnVisibilityChange = (columnId) => {
    setColumnVisibility({
      ...columnVisibility,
      [columnId]: !columnVisibility[columnId]
    });
  };

  // 处理全选/取消全选列显示
  const handleSelectAllColumns = () => {
    const allColumns = Object.keys(columnVisibility);
    const areAllVisible = allColumns.every((column) => columnVisibility[column]);

    const newColumnVisibility = {};
    allColumns.forEach((column) => {
      newColumnVisibility[column] = !areAllVisible;
    });

    setColumnVisibility(newColumnVisibility);
  };

  const handleSort = (event, id) => {
    const isAsc = orderBy === id && order === 'asc';
    if (id !== '') {
      setOrder(isAsc ? 'desc' : 'asc');
      setOrderBy(id);
    }
  };

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    const newRowsPerPage = parseInt(event.target.value, 10);
    setPage(0);
    setRowsPerPage(newRowsPerPage);
    savePageSize('log', newRowsPerPage);
  };

  const searchLogs = async () => {
    setPage(0);
    setSearchKeyword({ ...toolBarValue });
  };

  const handleToolBarValue = (event) => {
    const { name, value } = event.target;
    setToolBarValue((prev) => ({ ...prev, [name]: value }));
  };

  const handleTabsChange = async (event, newValue) => {
    const updatedToolBarValue = { ...toolBarValue, log_type: newValue };
    setToolBarValue(updatedToolBarValue);
    setPage(0);
    setSearchKeyword(updatedToolBarValue);
  };

  const fetchData = useCallback(
    async (page, rowsPerPage, keyword, order, orderBy) => {
      setSearching(true);
      keyword = trims(keyword);
      try {
        if (orderBy) {
          orderBy = order === 'desc' ? '-' + orderBy : orderBy;
        }
        const url = userIsAdmin ? '/api/log/' : '/api/log/self/';
        if (!userIsAdmin) {
          delete keyword.username;
          delete keyword.channel_id;
        }

        const res = await API.get(url, {
          params: {
            page: page + 1,
            size: rowsPerPage,
            order: orderBy,
            ...keyword
          }
        });
        const { success, message, data } = res.data;
        if (success) {
          setListCount(data.total_count);
          setLogs(data.data);
        } else {
          showError(message);
        }
      } catch (error) {
        console.error(error);
      }
      setSearching(false);
    },
    [userIsAdmin]
  );

  // 处理刷新
  const handleRefresh = async () => {
    setOrderBy('created_at');
    setOrder('desc');
    setToolBarValue({ ...originalKeyword });
    setSearchKeyword({ ...originalKeyword });
    setRefreshFlag(!refreshFlag);
  };

  // 导出 CSV（按当前筛选条件）
  const handleExport = async () => {
    try {
      // 以当前已应用的筛选条件导出（searchKeyword）
      let params = { ...searchKeyword };
      let sort = orderBy;
      if (sort) {
        params.order = order === 'desc' ? `-${sort}` : sort;
      }
      const url = userIsAdmin ? '/api/log/export' : '/api/log/self/export';
      const res = await API.get(url, {
        params,
        responseType: 'blob'
      });
      const blob = new Blob([res.data], { type: 'text/csv;charset=utf-8' });
      const link = document.createElement('a');
      const href = URL.createObjectURL(blob);
      link.href = href;
      const now = dayjs().format('YYYYMMDD_HHmmss');
      link.download = `logs_${now}.csv`;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      URL.revokeObjectURL(href);
    } catch (error) {
      console.error(error);
      showError('导出失败，请检查筛选条件与网络状态');
    }
  };

  useEffect(() => {
    fetchData(page, rowsPerPage, searchKeyword, order, orderBy);
  }, [page, rowsPerPage, searchKeyword, order, orderBy, fetchData, refreshFlag]);

  // 行为区（刷新/搜索/列设置）
  const actionsNode = matchUpMd ? (
    <ButtonGroup variant="outlined" aria-label="outlined small primary button group">
      <Button onClick={handleRefresh} size="small" startIcon={<Icon icon="solar:refresh-bold-duotone" width={18} />}>
        {t('logPage.refreshButton')}
      </Button>
      <Button onClick={searchLogs} size="small" startIcon={<Icon icon="solar:minimalistic-magnifer-line-duotone" width={18} />}>
        {t('logPage.searchButton')}
      </Button>
      <Button onClick={handleExport} size="small" startIcon={<Icon icon="solar:download-minimalistic-bold-duotone" width={18} />}>
        {t('logPage.exportButton', { defaultValue: '导出CSV' })}
      </Button>
      <Button onClick={handleColumnMenuOpen} size="small" startIcon={<Icon icon="solar:settings-bold-duotone" width={18} />}>
        {t('logPage.columnSettings')}
      </Button>
    </ButtonGroup>
  ) : (
    <Stack direction="row" spacing={1} divider={<Divider orientation="vertical" flexItem />} justifyContent="space-around" alignItems="center">
      <IconButton onClick={handleRefresh} size="small">
        <Icon icon="solar:refresh-bold-duotone" width={18} />
      </IconButton>
      <IconButton onClick={searchLogs} size="small">
        <Icon icon="solar:minimalistic-magnifer-line-duotone" width={18} />
      </IconButton>
      <IconButton onClick={handleExport} size="small">
        <Icon icon="solar:download-minimalistic-bold-duotone" width={18} />
      </IconButton>
      <IconButton onClick={handleColumnMenuOpen} size="small">
        <Icon icon="solar:settings-bold-duotone" width={18} />
      </IconButton>
    </Stack>
  );

  return (
    <>
      <PageHeader title={t('logPage.title')} subtitle="Log" />
      <Card>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs
            value={toolBarValue.log_type}
            onChange={handleTabsChange}
            aria-label="basic tabs example"
            variant="scrollable"
            scrollButtons="auto"
            allowScrollButtonsMobile
            sx={{
              '& .MuiTabs-indicator': {
                display: 'none'
              }
            }}
          >
            {Object.values(LogType).map((option) => {
              return <Tab key={option.value} label={option.text} value={option.value} />;
            })}
          </Tabs>
        </Box>
        <FilterBar>
          <Box component="form" noValidate>
            <TableToolBar
              filterName={toolBarValue}
              handleFilterName={handleToolBarValue}
              userIsAdmin={userIsAdmin}
              actions={actionsNode}
            />
          </Box>
        </FilterBar>
        <Container maxWidth="xl" disableGutters>
          <Menu
              anchorEl={columnMenuAnchor}
              open={Boolean(columnMenuAnchor)}
              onClose={handleColumnMenuClose}
              PaperProps={{
                style: {
                  maxHeight: 300,
                  width: 200
                }
              }}
            >
              <MenuItem disabled>
                <Typography variant="subtitle2">{t('logPage.selectColumns')}</Typography>
              </MenuItem>
              <MenuItem onClick={handleSelectAllColumns} dense>
                <Checkbox
                  checked={Object.values(columnVisibility).every((visible) => visible)}
                  indeterminate={
                    !Object.values(columnVisibility).every((visible) => visible) &&
                    Object.values(columnVisibility).some((visible) => visible)
                  }
                  size="small"
                />
                <ListItemText primary={t('logPage.columnSelectAll')} />
              </MenuItem>
              {[
                { id: 'created_at', label: t('logPage.timeLabel') },
                { id: 'channel_id', label: t('logPage.channelLabel'), adminOnly: true },
                { id: 'user_id', label: t('logPage.userLabel'), adminOnly: true },
                { id: 'group', label: t('logPage.groupLabel') },
                { id: 'token_name', label: t('logPage.tokenLabel') },
                { id: 'type', label: t('logPage.typeLabel') },
                { id: 'model_name', label: t('logPage.modelLabel') },
                { id: 'duration', label: t('logPage.durationLabel') },
                { id: 'message', label: t('logPage.inputLabel') },
                { id: 'completion', label: t('logPage.outputLabel') },
                { id: 'quota', label: t('logPage.quotaLabel') },
                { id: 'source_ip', label: t('logPage.sourceIp') },
                { id: 'detail', label: t('logPage.detailLabel') }
              ].map(
                (column) =>
                  (!column.adminOnly || userIsAdmin) && (
                    <MenuItem key={column.id} onClick={() => handleColumnVisibilityChange(column.id)} dense>
                      <Checkbox checked={columnVisibility[column.id] || false} size="small" />
                      <ListItemText primary={column.label} />
                    </MenuItem>
                  )
              )}
            </Menu>
          </Container>
        {searching && <LinearProgress />}
        <ScrollArea>
          <TableContainer sx={{ overflow: 'auto' }}>
            <Table sx={{ minWidth: 800 }}>
              <KeywordTableHead
                order={order}
                orderBy={orderBy}
                onRequestSort={handleSort}
                headLabel={[
                  {
                    id: 'created_at',
                    label: t('logPage.timeLabel'),
                    disableSort: false,
                    hide: !columnVisibility.created_at
                  },
                  {
                    id: 'channel_id',
                    label: t('logPage.channelLabel'),
                    disableSort: false,
                    hide: !columnVisibility.channel_id || !userIsAdmin
                  },
                  {
                    id: 'user_id',
                    label: t('logPage.userLabel'),
                    disableSort: false,
                    hide: !columnVisibility.user_id || !userIsAdmin
                  },
                  {
                    id: 'group',
                    label: t('logPage.groupLabel'),
                    disableSort: false,
                    hide: !columnVisibility.group
                  },
                  {
                    id: 'token_name',
                    label: t('logPage.tokenLabel'),
                    disableSort: false,
                    hide: !columnVisibility.token_name
                  },
                  {
                    id: 'type',
                    label: t('logPage.typeLabel'),
                    disableSort: false,
                    hide: !columnVisibility.type
                  },
                  {
                    id: 'model_name',
                    label: t('logPage.modelLabel'),
                    disableSort: false,
                    hide: !columnVisibility.model_name
                  },
                  {
                    id: 'duration',
                    label: t('logPage.durationLabel'),
                    tooltip: t('logPage.durationTooltip'),
                    disableSort: true,
                    hide: !columnVisibility.duration
                  },
                  {
                    id: 'message',
                    label: t('logPage.inputLabel'),
                    disableSort: true,
                    hide: !columnVisibility.message
                  },
                  {
                    id: 'completion',
                    label: t('logPage.outputLabel'),
                    disableSort: true,
                    hide: !columnVisibility.completion
                  },
                  {
                    id: 'quota',
                    label: t('logPage.quotaLabel'),
                    disableSort: true,
                    hide: !columnVisibility.quota
                  },
                  {
                    id: 'source_ip',
                    label: t('logPage.sourceIp'),
                    disableSort: true,
                    hide: !columnVisibility.source_ip
                  },
                  {
                    id: 'detail',
                    label: t('logPage.detailLabel'),
                    disableSort: true,
                    hide: !columnVisibility.detail
                  }
                ]}
              />
              <TableBody>
                {logs.map((row, index) => (
                  <LogTableRow
                    item={row}
                    key={`${row.id}_${index}`}
                    userIsAdmin={userIsAdmin}
                    userGroup={userGroup}
                    columnVisibility={columnVisibility}
                  />
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </ScrollArea>
        <TablePagination
          page={page}
          component="div"
          count={listCount}
          rowsPerPage={rowsPerPage}
          onPageChange={handleChangePage}
          rowsPerPageOptions={PAGE_SIZE_OPTIONS}
          onRowsPerPageChange={handleChangeRowsPerPage}
          showFirstButton
          showLastButton
        />
      </Card>
    </>
  );
}
