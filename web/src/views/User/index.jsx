import { useState, useEffect } from 'react';
import { showError, showSuccess, trims } from 'utils/common';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
// import PerfectScrollbar from 'react-perfect-scrollbar';
import ScrollArea from 'ui-component/ScrollArea';
import TablePagination from '@mui/material/TablePagination';
import LinearProgress from '@mui/material/LinearProgress';
import ButtonGroup from '@mui/material/ButtonGroup';
import FilterBar from 'ui-component/FilterBar';
import ActionBar from 'ui-component/ActionBar';
import { Icon } from '@iconify/react';

import {
  Button,
  Card,
  Box,
  Stack,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  FormHelperText,
  Checkbox,
  ListItemText,
  TextField,
  Chip,
  Switch,
  FormControlLabel
} from '@mui/material';
import Autocomplete from '@mui/material/Autocomplete';
import UsersTableRow from './component/TableRow';
import KeywordTableHead from 'ui-component/TableHead';
import TableToolBar from 'ui-component/TableToolBar';
import { API } from 'utils/api';
import { PAGE_SIZE_OPTIONS, getPageSize, savePageSize } from 'constants';
import EditeModal from './component/EditModal';

import { useTranslation } from 'react-i18next';
import PageHeader from 'ui-component/PageHeader';
// ----------------------------------------------------------------------
export default function Users() {
  const { t } = useTranslation();
  const [page, setPage] = useState(0);
  const [order, setOrder] = useState('desc');
  const [orderBy, setOrderBy] = useState('id');
  const [rowsPerPage, setRowsPerPage] = useState(() => getPageSize('user'));
  const [listCount, setListCount] = useState(0);
  const [searching, setSearching] = useState(false);
  const [searchKeyword, setSearchKeyword] = useState('');
  const [users, setUsers] = useState([]);
  const [refreshFlag, setRefreshFlag] = useState(false);

  const [openModal, setOpenModal] = useState(false);
  const [editUserId, setEditUserId] = useState(0);
  const [openGroupModal, setOpenGroupModal] = useState(false);
  const [groupTargetUser, setGroupTargetUser] = useState(null);
  const [selectedGroup, setSelectedGroup] = useState('');
  const [changingGroup, setChangingGroup] = useState(false);
  const [groupOptions, setGroupOptions] = useState([]);
  const [openAllowedModal, setOpenAllowedModal] = useState(false);
  const [allowedTargetUser, setAllowedTargetUser] = useState(null);
  const [allowedSelected, setAllowedSelected] = useState([]);
  const [savingAllowed, setSavingAllowed] = useState(false);
  const [groupMetaMap, setGroupMetaMap] = useState({});
  const [allowedOnlyPrivate, setAllowedOnlyPrivate] = useState(false);

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
    savePageSize('user', newRowsPerPage);
  };

  const fetchGroups = async () => {
    try {
      // 1) 基础符号列表（兼容原有使用场景）
      const res = await API.get(`/api/group/`);
      const symbols = res.data.data || [];
      setGroupOptions(symbols);
      // 2) 管理维度的分组元数据（含 name/ratio/public/enable），用于更友好的 UI 展现
      try {
        const metaRes = await API.get(`/api/user_group/`, { params: { page: 1, size: 1000 } });
        const metaList = metaRes?.data?.data?.data || [];
        const map = {};
        metaList.forEach((it) => {
          map[it.symbol] = it;
        });
        setGroupMetaMap(map);
      } catch (e) {
        // 忽略，保底仍可使用符号列表
      }
    } catch (error) {
      console.error(error);
      showError(error.message);
    }
  };

  const searchUsers = async (event) => {
    event.preventDefault();
    const formData = new FormData(event.target);
    setPage(0);
    setSearchKeyword(formData.get('keyword'));
  };

  const handleOpenGroupModal = (user) => {
    if (!user || user.role !== 1) {
      return;
    }
    if (!groupOptions.length) {
      fetchGroups();
    }
    setGroupTargetUser(user);
    setSelectedGroup(user.group || '');
    setOpenGroupModal(true);
  };

  const handleCloseGroupModal = () => {
    setOpenGroupModal(false);
    setGroupTargetUser(null);
    setSelectedGroup('');
  };

  const handleSubmitGroupChange = async () => {
    if (!groupTargetUser || !selectedGroup) {
      showError(t('userPage.selectGroup'));
      return;
    }
    setChangingGroup(true);
    try {
      const res = await manageUser(groupTargetUser.username, 'group', selectedGroup);
      const { success, message } = res;
      if (success) {
        showSuccess(t('userPage.changeGroupSuccess'));
        handleCloseGroupModal();
        handleRefresh();
      } else {
        showError(message);
      }
    } catch (error) {
      console.error(error);
      showError(error.message);
    }
    setChangingGroup(false);
  };

  const handleOpenAllowedModal = async (user) => {
    if (!user) return;
    if (!groupOptions.length) {
      await fetchGroups();
    }
    setAllowedTargetUser(user);
    try {
      const res = await API.get(`/api/user/allowed_groups/${user.id}`);
      const groups = res?.data?.data || [];
      // 默认值为用户自身分组（当后端尚未设置白名单时）
      if (!groups || groups.length === 0) {
        const def = user.group ? [user.group] : [];
        setAllowedSelected(def);
      } else {
        setAllowedSelected(groups);
      }
    } catch (e) {
      const def = user.group ? [user.group] : [];
      setAllowedSelected(def);
    }
    setOpenAllowedModal(true);
  };

  const handleCloseAllowedModal = () => {
    setOpenAllowedModal(false);
    setAllowedTargetUser(null);
    setAllowedSelected([]);
  };

  const handleSaveAllowed = async () => {
    if (!allowedTargetUser) return;
    setSavingAllowed(true);
    try {
      const res = await API.put(`/api/user/allowed_groups/${allowedTargetUser.id}`, { groups: allowedSelected });
      if (res?.data?.success) {
        showSuccess(t('userPage.saveSuccess'));
        handleCloseAllowedModal();
      } else {
        showError(res?.data?.message || '');
      }
    } catch (e) {
      showError(e.message);
    }
    setSavingAllowed(false);
  };

  const fetchData = async (page, rowsPerPage, keyword, order, orderBy) => {
    setSearching(true);
    keyword = trims(keyword);
    try {
      if (orderBy) {
        orderBy = order === 'desc' ? '-' + orderBy : orderBy;
      }
      const res = await API.get(`/api/user/`, {
        params: {
          page: page + 1,
          size: rowsPerPage,
          keyword: keyword,
          order: orderBy
        }
      });
      const { success, message, data } = res.data;
      if (success) {
        setListCount(data.total_count);
        setUsers(data.data);
      } else {
        showError(message);
      }
    } catch (error) {
      console.error(error);
    }
    setSearching(false);
  };

  // 处理刷新
  const handleRefresh = async () => {
    setOrderBy('id');
    setOrder('desc');
    setRefreshFlag(!refreshFlag);
  };

  useEffect(() => {
    fetchData(page, rowsPerPage, searchKeyword, order, orderBy);
  }, [page, rowsPerPage, searchKeyword, order, orderBy, refreshFlag]);

  useEffect(() => {
    fetchGroups();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const manageUser = async (username, action, value) => {
    let url = '/api/user/manage';
    let valueData = {};
    // let data = { username: username, action: '' };
    let res;
    switch (action) {
      case 'delete':
        valueData = { username, action: 'delete' };
        break;
      case 'status':
        valueData = { username, action: value === 1 ? 'enable' : 'disable' };
        break;
      case 'role':
        valueData = { username, action: value === true ? 'promote' : 'demote' };
        break;
      case 'quota':
        url = `/api/user/quota/${username}`;
        valueData = value;
        break;
      case 'group':
        valueData = { username, action: 'group', group: value };
        break;
    }

    try {
      res = await API.post(url, valueData);
      const { success, message } = res.data;
      if (success) {
        showSuccess(t('userPage.operationSuccess'));
        await handleRefresh();
      } else {
        showError(message);
      }

      return res.data;
    } catch (error) {
      return;
    }
  };

  const handleOpenModal = (userId) => {
    setEditUserId(userId);
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    setOpenModal(false);
    setEditUserId(0);
  };

  const handleOkModal = (status) => {
    if (status === true) {
      handleCloseModal();
      handleRefresh();
    }
  };

  const headerActions = [
    <Button
      key="create"
      variant="contained"
      color="primary"
      startIcon={<Icon icon="solar:add-circle-line-duotone" />}
      onClick={() => handleOpenModal(0)}
    >
      {t('userPage.createUser')}
    </Button>
  ];

  return (
    <>
      <PageHeader title={t('userPage.users')} subtitle="User" actions={headerActions} />
      <Card>
        <FilterBar>
          <Box component="form" onSubmit={searchUsers} noValidate>
            <TableToolBar placeholder={t('userPage.searchPlaceholder')} />
          </Box>
        </FilterBar>
        <ActionBar>
          <ButtonGroup variant="outlined" aria-label="outlined small primary button group">
            <Button onClick={handleRefresh} startIcon={<Icon icon="solar:refresh-bold-duotone" width={18} />}>
              {t('userPage.refresh')}
            </Button>
          </ButtonGroup>
        </ActionBar>
        {searching && <LinearProgress />}
        <ScrollArea>
          <TableContainer sx={{ overflow: 'auto' }}>
            <Table sx={{ minWidth: 800 }}>
              <KeywordTableHead
                order={order}
                orderBy={orderBy}
                onRequestSort={handleSort}
                headLabel={[
                  { id: 'id', label: t('userPage.id'), disableSort: false },
                  { id: 'username', label: t('userPage.username'), disableSort: false },
                  { id: 'group', label: t('userPage.group'), disableSort: true },
                  { id: 'stats', label: t('userPage.statistics'), disableSort: true },
                  { id: 'role', label: t('userPage.userRole'), disableSort: false },
                  { id: 'bind', label: t('userPage.bind'), disableSort: true },
                  { id: 'created_time', label: t('userPage.creationTime'), disableSort: false },
                  { id: 'status', label: t('userPage.status'), disableSort: false },
                  { id: 'action', label: t('userPage.action'), disableSort: true }
                ]}
              />
              <TableBody>
                {users.map((row) => (
                  <UsersTableRow
                    item={row}
                    manageUser={manageUser}
                    key={row.id}
                    handleOpenModal={handleOpenModal}
                    setModalUserId={setEditUserId}
                    handleOpenGroupModal={handleOpenGroupModal}
                    handleOpenAllowedModal={handleOpenAllowedModal}
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
      <EditeModal open={openModal} onCancel={handleCloseModal} onOk={handleOkModal} userId={editUserId} />
      <Dialog open={openGroupModal} onClose={handleCloseGroupModal} fullWidth maxWidth={'sm'}>
        <DialogTitle>{t('userPage.changeGroup')}</DialogTitle>
        <DialogContent>
          <FormControl fullWidth sx={{ mt: 1 }}>
            <InputLabel id="change-group-label">{t('userPage.group')}</InputLabel>
            <Select
              labelId="change-group-label"
              id="change-group"
              label={t('userPage.group')}
              value={selectedGroup}
              onChange={(event) => setSelectedGroup(event.target.value)}
              disabled={changingGroup}
              MenuProps={{
                PaperProps: {
                  style: {
                    maxHeight: 260
                  }
                }
              }}
            >
              {groupOptions.map((option) => (
                <MenuItem key={option} value={option}>
                  {option}
                </MenuItem>
              ))}
            </Select>
            {!selectedGroup && <FormHelperText error>{t('userPage.selectGroup')}</FormHelperText>}
            <FormHelperText>{t('userPage.changeGroupDesc')}</FormHelperText>
          </FormControl>
          <FormHelperText sx={{ mt: 2 }}>{t('userPage.changeGroupOnlyCommon')}</FormHelperText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseGroupModal}>{t('userPage.cancel')}</Button>
          <Button variant="contained" onClick={handleSubmitGroupChange} disabled={changingGroup || !selectedGroup}>
            {t('userPage.submit')}
          </Button>
        </DialogActions>
      </Dialog>
      <Dialog open={openAllowedModal} onClose={handleCloseAllowedModal} fullWidth maxWidth={'sm'}>
        <DialogTitle>{t('userPage.allowedGroups')}</DialogTitle>
        <DialogContent>
          <Stack direction="row" alignItems="center" justifyContent="space-between" sx={{ mt: 1, mb: 1 }}>
            <FormControlLabel
              control={<Switch checked={allowedOnlyPrivate} onChange={(e) => setAllowedOnlyPrivate(e.target.checked)} />}
              label={t('userPage.privateOnly')}
            />
            <Stack direction="row" spacing={1}>
              <Button
                size="small"
                onClick={() => setAllowedSelected(groupOptions.slice())}
                disabled={groupOptions.length === 0 || savingAllowed}
              >
                {t('userPage.selectAll')}
              </Button>
              <Button size="small" onClick={() => setAllowedSelected([])} disabled={savingAllowed}>
                {t('userPage.clearAll')}
              </Button>
            </Stack>
          </Stack>

          <Autocomplete
            multiple
            disableCloseOnSelect
            options={(() => {
              const syms = groupOptions || [];
              const filtered = allowedOnlyPrivate
                ? syms.filter((s) => groupMetaMap[s]?.public === false)
                : syms;
              return filtered.map((s) => ({
                symbol: s,
                label:
                  (groupMetaMap[s]?.name ? `${groupMetaMap[s].name}` : s) +
                  (groupMetaMap[s]?.ratio ? `（×${groupMetaMap[s].ratio}）` : '') +
                  (groupMetaMap[s]?.public ? ' · 公共' : ''),
              }));
            })()}
            getOptionLabel={(opt) => opt.label}
            renderOption={(props, option, { selected }) => (
              <li {...props} key={option.symbol}>
                <Checkbox checked={selected} sx={{ mr: 1 }} />
                <ListItemText primary={option.label} secondary={option.symbol} />
              </li>
            )}
            value={(() => {
              const set = new Set(allowedSelected);
              return (groupOptions || [])
                .filter((s) => set.has(s))
                .map((s) => ({
                  symbol: s,
                  label:
                    (groupMetaMap[s]?.name ? `${groupMetaMap[s].name}` : s) +
                    (groupMetaMap[s]?.ratio ? `（×${groupMetaMap[s].ratio}）` : '') +
                    (groupMetaMap[s]?.public ? ' · 公共' : ''),
                }));
            })()}
            onChange={(_, newValue) => {
              setAllowedSelected(newValue.map((v) => v.symbol));
            }}
            renderTags={(value, getTagProps) =>
              value.map((option, index) => (
                <Chip variant="outlined" label={option.symbol} {...getTagProps({ index })} key={option.symbol} />
              ))
            }
            renderInput={(params) => <TextField {...params} label={t('userPage.group')} placeholder={t('common.search') || '搜索'} />}
          />
          <FormHelperText sx={{ mt: 1 }}>
            {t('userPage.selectedCount', { count: allowedSelected.length, total: groupOptions.length })}
          </FormHelperText>
          <FormHelperText>{t('userPage.allowedGroupsDesc')}</FormHelperText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseAllowedModal}>{t('userPage.cancel')}</Button>
          <Button variant="contained" onClick={handleSaveAllowed} disabled={savingAllowed}>
            {t('userPage.submit')}
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
