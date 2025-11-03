import { useState, useEffect, useMemo } from 'react';
import { showError, showSuccess } from 'utils/common';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
// import PerfectScrollbar from 'react-perfect-scrollbar';
import ScrollArea from 'ui-component/ScrollArea';
import ButtonGroup from '@mui/material/ButtonGroup';
import Toolbar from '@mui/material/Toolbar';

import { Button, Card, Stack, Container, Typography, Tabs, Tab, Box } from '@mui/material';
import ModelOwnedbyTableRow from './component/TableRow';
import KeywordTableHead from 'ui-component/TableHead';
import { API } from 'utils/api';
import EditeModal from './component/EditModal';
import { Icon } from '@iconify/react';

import { useTranslation } from 'react-i18next';
import BrandOverview from './component/BrandOverview';

const TabPanel = ({ value, current, children }) => {
  if (value !== current) {
    return null;
  }

  return (
    <Box sx={{ mt: 3 }}>
      {children}
    </Box>
  );
};

export default function ModelOwnedby() {
  const { t } = useTranslation();
  const [tab, setTab] = useState('dictionary');
  const [modelOwnedby, setModelOwnedby] = useState([]);
  const [refreshFlag, setRefreshFlag] = useState(false);

  const [overviewParams, setOverviewParams] = useState({
    keyword: '',
    owned_by_type: '',
    channel_type: '',
    issue: '',
    group: '',
    page: 1,
    page_size: 20
  });
  const [overviewResult, setOverviewResult] = useState({
    list: [],
    total: 0,
    page: 1,
    page_size: 20,
    issue_stats: {},
    groups: []
  });
  const [overviewLoading, setOverviewLoading] = useState(false);
  const [overviewLoaded, setOverviewLoaded] = useState(false);
  const [selectedModels, setSelectedModels] = useState([]);
  const [bulkUpdating, setBulkUpdating] = useState(false);
  // 价格渠道批量更新使用同一loading标志，保持交互一致

  const [openModal, setOpenModal] = useState(false);
  const [editId, setEditId] = useState(0);

  const fetchDictionary = async () => {
    try {
      const res = await API.get(`/api/model_ownedby/`);
      const { success, message, data } = res.data;
      if (success) {
        setModelOwnedby(data);
      } else {
        showError(message);
      }
    } catch (error) {
      showError(error.message);
    }
  };

  const loadOverview = async (override = {}) => {
    const nextParams = {
      ...overviewParams,
      ...override
    };

    const safePage = nextParams.page > 0 ? nextParams.page : 1;
    const safePageSize = nextParams.page_size > 0 ? Math.min(nextParams.page_size, 200) : 20;

    const requestParams = {
      page: safePage,
      page_size: safePageSize
    };

    if (nextParams.keyword) {
      requestParams.keyword = nextParams.keyword;
    }
    if (nextParams.owned_by_type !== '') {
      requestParams.owned_by_type = nextParams.owned_by_type;
    }
    if (nextParams.channel_type !== '') {
      requestParams.channel_type = nextParams.channel_type;
    }
    if (nextParams.issue !== '') {
      requestParams.issue = nextParams.issue;
    }
    if (nextParams.group !== '') {
      requestParams.group = nextParams.group;
    }

    setOverviewParams({
      ...nextParams,
      page: safePage,
      page_size: safePageSize
    });
    setOverviewLoading(true);
    try {
      const res = await API.get(`/api/model_ownedby/overview`, { params: requestParams });
      const { success, message, data } = res.data;
      if (success) {
        const list = data?.list || [];
        setOverviewResult({
          list,
          total: data?.total ?? 0,
          page: data?.page ?? safePage,
          page_size: data?.page_size ?? safePageSize,
          issue_stats: data?.issue_stats || {},
          groups: data?.groups || []
        });
        setOverviewLoaded(true);
        setSelectedModels((prevSelected) => prevSelected.filter((model) => list.some((item) => item.model === model)));
      } else {
        showError(message);
      }
    } catch (error) {
      showError(error.message);
    } finally {
      setOverviewLoading(false);
    }
  };

  const handleRefresh = async () => {
    setRefreshFlag(!refreshFlag);
  };

  const handleOverviewRefresh = async () => {
    await loadOverview();
  };

  useEffect(() => {
    fetchDictionary();
  }, [refreshFlag]);

  useEffect(() => {
    if (tab === 'overview' && !overviewLoaded) {
      loadOverview();
    }
  }, [tab, overviewLoaded]);

  const manageModelOwnedBy = async (id, action) => {
    const url = '/api/model_ownedby/';
    let res;
    try {
      switch (action) {
        case 'delete':
          res = await API.delete(url + id);
          break;
        default:
          return false;
      }

      const { success, message } = res.data;
      if (success) {
        showSuccess(t('userPage.operationSuccess'));
        await handleRefresh();
      } else {
        showError(message);
      }

      return res.data;
    } catch (error) {
      showError(error.message);
      return;
    }
  };

  const handleOpenModal = (userId) => {
    setEditId(userId);
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    setOpenModal(false);
    setEditId(0);
  };

  const handleOkModal = (status) => {
    if (status === true) {
      handleCloseModal();
      handleRefresh();
    }
  };

  const handleTabChange = (_event, newValue) => {
    setTab(newValue);
  };

  const overviewTotal = useMemo(() => overviewResult.total, [overviewResult.total]);

  const handleOverviewFiltersChange = (changes) => {
    loadOverview({ ...changes, page: 1 });
  };

  const handleOverviewPageChange = (newPage, newPageSize) => {
    loadOverview({ page: newPage, page_size: newPageSize });
  };

  const handleSelectAll = (checked, models) => {
    if (checked) {
      setSelectedModels((prev) => {
        const nextSet = new Set(prev);
        models.forEach((model) => nextSet.add(model));
        return Array.from(nextSet);
      });
    } else {
      setSelectedModels((prev) => prev.filter((model) => !models.includes(model)));
    }
  };

  const handleSelectOne = (model, checked) => {
    setSelectedModels((prev) => {
      if (checked) {
        if (prev.includes(model)) {
          return prev;
        }
        return [...prev, model];
      }
      return prev.filter((item) => item !== model);
    });
  };

  const handleBulkUpdate = async (ownedByType) => {
    if (!selectedModels.length) {
      return;
    }
    setBulkUpdating(true);
    try {
      const res = await API.post('/api/model_ownedby/overview/batch', {
        models: selectedModels,
        owned_by_type: ownedByType
      });
      const { success, message } = res.data;
      if (success) {
        showSuccess(t('modelOwnedby.overview.bulkSuccess'));
        setSelectedModels([]);
        await loadOverview();
      } else {
        showError(message);
      }
    } catch (error) {
      showError(error.message);
    } finally {
      setBulkUpdating(false);
    }
  };

  // 批量调整价格渠道
  const handleBulkUpdateChannel = async (channelType) => {
    if (!selectedModels.length) {
      return;
    }
    setBulkUpdating(true);
    try {
      const res = await API.post('/api/model_ownedby/overview/batch_channel', {
        models: selectedModels,
        channel_type: channelType
      });
      const { success, message } = res.data;
      if (success) {
        showSuccess(t('modelOwnedby.overview.bulkChannelSuccess'));
        setSelectedModels([]);
        await loadOverview();
      } else {
        showError(message);
      }
    } catch (error) {
      showError(error.message);
    } finally {
      setBulkUpdating(false);
    }
  };

  return (
    <>
      <Stack direction="row" alignItems="center" justifyContent="space-between" mb={2}>
        <Stack direction="column" spacing={1}>
          <Typography variant="h2">{t('modelOwnedby.title')}</Typography>
          <Typography variant="subtitle1" color="text.secondary">
            Model Owned By
          </Typography>
        </Stack>

        {tab === 'dictionary' && (
          <Button
            variant="contained"
            color="primary"
            startIcon={<Icon icon="solar:add-circle-line-duotone" />}
            onClick={() => handleOpenModal(0)}
          >
            {t('modelOwnedby.create')}
          </Button>
        )}
      </Stack>

      <Tabs value={tab} onChange={handleTabChange} variant="scrollable" allowScrollButtonsMobile>
        <Tab value="dictionary" label={t('modelOwnedby.tabDictionary')} />
        <Tab
          value="overview"
          label={
            overviewTotal > 0
              ? `${t('modelOwnedby.tabOverview')} (${overviewTotal})`
              : t('modelOwnedby.tabOverview')
          }
        />
      </Tabs>

      <TabPanel value="dictionary" current={tab}>
        <Card>
          <Toolbar
            sx={{
              textAlign: 'right',
              height: 50,
              display: 'flex',
              justifyContent: 'space-between',
              p: (theme) => theme.spacing(0, 1, 0, 3)
            }}
          >
            <Container maxWidth="xl">
              <ButtonGroup variant="outlined" aria-label="outlined small primary button group">
                <Button onClick={handleRefresh} startIcon={<Icon icon="solar:refresh-bold-duotone" width={18} />}>
                  {t('userPage.refresh')}
                </Button>
              </ButtonGroup>
            </Container>
          </Toolbar>
          <ScrollArea>
            <TableContainer sx={{ overflow: 'unset' }}>
              <Table sx={{ minWidth: 800 }}>
                <KeywordTableHead
                  headLabel={[
                    { id: 'id', label: t('modelOwnedby.id'), disableSort: false },
                    { id: 'name', label: t('modelOwnedby.name'), disableSort: false },
                    { id: 'icon', label: t('modelOwnedby.icon'), disableSort: false },
                    { id: 'action', label: t('modelOwnedby.action'), disableSort: true }
                  ]}
                />
                <TableBody>
                  {modelOwnedby.map((row) => (
                    <ModelOwnedbyTableRow
                      item={row}
                      manageModelOwnedBy={manageModelOwnedBy}
                      key={row.id}
                      handleOpenModal={handleOpenModal}
                      setModalId={setEditId}
                    />
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </ScrollArea>
        </Card>
      </TabPanel>

      <TabPanel value="overview" current={tab}>
        <BrandOverview
          data={overviewResult.list}
          loading={overviewLoading}
          onRefresh={handleOverviewRefresh}
          page={overviewResult.page}
          pageSize={overviewResult.page_size}
          total={overviewResult.total}
          onPageChange={handleOverviewPageChange}
          filters={overviewParams}
          onFiltersChange={handleOverviewFiltersChange}
          ownedByOptions={modelOwnedby}
          issueStats={overviewResult.issue_stats}
          groups={overviewResult.groups}
          selected={selectedModels}
          onSelectAll={handleSelectAll}
          onSelectOne={handleSelectOne}
          onBulkUpdate={handleBulkUpdate}
          onBulkUpdateChannel={handleBulkUpdateChannel}
          bulkUpdating={bulkUpdating}
        />
      </TabPanel>

      <EditeModal open={openModal} onCancel={handleCloseModal} onOk={handleOkModal} Oid={editId} />
    </>
  );
}
