import { useState, useEffect } from 'react';
import {
  Box,
  Card,
  Stack,
  Typography,
  TextField,
  Button,
  MenuItem,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  TableContainer,
  Paper,
  FormControl,
  InputLabel,
  Select,
  FormControlLabel,
  Switch
} from '@mui/material';
import { Icon } from '@iconify/react';
import { useTranslation } from 'react-i18next';
import { API } from 'utils/api';
import { showError, showSuccess } from 'utils/common';
import ActionBar from 'ui-component/ActionBar';
import AdminContainer from 'ui-component/AdminContainer';

// 模型分组管理页面（管理员）
// - 通过 /api/customer_pricing/model_groups 管理某个模型的分组与计费类型（含 seconds）

const BILLING_TYPES = [
  { value: 'tokens', label: 'tokens' },
  { value: 'times', label: 'times' },
  { value: 'seconds', label: 'seconds' }
];

const ModelGroupAdmin = () => {
  const { t } = useTranslation();
  const [modelList, setModelList] = useState([]);
  const [model, setModel] = useState('');
  const [groups, setGroups] = useState([]);
  const [loading, setLoading] = useState(false);
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    const fetchModelList = async () => {
      try {
        const res = await API.get('/api/prices/model_list');
        const { success, message, data } = res.data || {};
        if (success) {
          setModelList(data || []);
        } else {
          showError(message || '加载模型列表失败');
        }
      } catch (e) {
        console.error(e);
      }
    };
    fetchModelList();
  }, []);

  const handleLoad = async () => {
    if (!model) {
      showError('请先选择模型');
      return;
    }
    setLoading(true);
    try {
      const res = await API.get('/api/customer_pricing/model_groups', { params: { model } });
      const { success, message, data } = res.data || {};
      if (!success) {
        showError(message || '加载分组失败');
        setGroups([]);
        return;
      }
      const list = Array.isArray(data) ? data : [];
      const mapped = list.map((it) => ({
        ...it,
        is_new: false
      }));
      setGroups(mapped);
    } catch (e) {
      console.error(e);
      showError(e.message || '加载分组失败');
      setGroups([]);
    }
    setLoading(false);
  };

  const handleAddGroup = () => {
    if (!model) {
      showError('请先选择模型');
      return;
    }
    setGroups((prev) => [
      ...prev,
      {
        id: undefined,
        model,
        group_code: '',
        display_name: '',
        description: '',
        billing_type: '',
        is_default: false,
        is_new: true
      }
    ]);
  };

  const handleFieldChange = (index, field) => (event) => {
    const value = field === 'is_default' ? event.target.checked : event.target.value;
    setGroups((prev) => {
      const next = [...prev];
      const row = { ...next[index], [field]: value };
      next[index] = row;
      return next;
    });
  };

  const handleSaveAll = async () => {
    if (!model) {
      showError('请先选择模型');
      return;
    }
    if (!groups.length) {
      showError('当前模型暂无分组需要保存');
      return;
    }
    setSaving(true);
    try {
      for (const g of groups) {
        if (!g.group_code) {
          // 跳过未填写编码的行
          continue;
        }
        const payload = {
          model,
          group_code: g.group_code,
          display_name: g.display_name,
          description: g.description,
          billing_type: g.billing_type || '',
          is_default: !!g.is_default
        };
        await API.post('/api/customer_pricing/model_groups', payload);
      }
      showSuccess('保存成功');
      handleLoad();
    } catch (e) {
      console.error(e);
      showError(e.message || '保存失败');
    }
    setSaving(false);
  };

  return (
    <AdminContainer>
      <Stack spacing={3}>
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={1}>
          <Stack direction="column" spacing={1}>
            <Typography variant="h2">模型分组管理</Typography>
            <Typography variant="subtitle1" color="text.secondary">
              为模型定义默认分组与非默认分组，并配置计费类型（tokens / times / seconds）。
            </Typography>
          </Stack>
          <ActionBar sx={{ py: 0 }}>
            <Button
              color="primary"
              variant="contained"
              onClick={handleSaveAll}
              disabled={saving || !model}
              startIcon={<Icon icon="solar:floppy-disk-bold-duotone" />}
            >
              {saving ? '保存中...' : '保存全部'}
            </Button>
          </ActionBar>
        </Stack>

        <Card sx={{ p: 2 }}>
          <Stack spacing={2} direction={{ xs: 'column', md: 'row' }}>
            <FormControl size="small" sx={{ minWidth: 220 }}>
              <InputLabel id="model-group-model-label">模型</InputLabel>
              <Select
                labelId="model-group-model-label"
                label="模型"
                value={model}
                onChange={(e) => setModel(e.target.value)}
              >
                {modelList.map((m) => (
                  <MenuItem key={m} value={m}>
                    {m}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
            <Button
              variant="outlined"
              onClick={handleLoad}
              disabled={loading || !model}
              startIcon={<Icon icon="solar:refresh-bold-duotone" />}
            >
              {loading ? '加载中...' : '加载当前模型分组'}
            </Button>
            <Button
              variant="outlined"
              color="secondary"
              onClick={handleAddGroup}
              disabled={!model}
              startIcon={<Icon icon="solar:add-circle-line-duotone" />}
            >
              新增分组
            </Button>
          </Stack>
        </Card>

        <Card>
          <TableContainer component={Paper}>
            <Table size="small">
              <TableHead>
                <TableRow>
                  <TableCell>分组编码</TableCell>
                  <TableCell>名称</TableCell>
                  <TableCell>描述</TableCell>
                  <TableCell>计费类型</TableCell>
                  <TableCell>默认分组</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {groups.map((row, idx) => (
                  <TableRow key={`${row.group_code || 'new'}-${idx}`} hover>
                    <TableCell>
                      <TextField
                        size="small"
                        value={row.group_code || ''}
                        onChange={handleFieldChange(idx, 'group_code')}
                        placeholder="例如：default / hq / enterprise"
                      />
                    </TableCell>
                    <TableCell>
                      <TextField
                        size="small"
                        value={row.display_name || ''}
                        onChange={handleFieldChange(idx, 'display_name')}
                        placeholder="展示名称"
                      />
                    </TableCell>
                    <TableCell>
                      <TextField
                        size="small"
                        value={row.description || ''}
                        onChange={handleFieldChange(idx, 'description')}
                        placeholder="说明"
                      />
                    </TableCell>
                    <TableCell>
                      <FormControl size="small" sx={{ minWidth: 140 }}>
                        <Select
                          value={row.billing_type || ''}
                          onChange={handleFieldChange(idx, 'billing_type')}
                          displayEmpty
                        >
                          <MenuItem value="">
                            <em>未设置</em>
                          </MenuItem>
                          {BILLING_TYPES.map((bt) => (
                            <MenuItem key={bt.value} value={bt.value}>
                              {bt.label}
                            </MenuItem>
                          ))}
                        </Select>
                      </FormControl>
                    </TableCell>
                    <TableCell>
                      <FormControlLabel
                        control={
                          <Switch
                            size="small"
                            checked={!!row.is_default}
                            onChange={handleFieldChange(idx, 'is_default')}
                          />
                        }
                        label=""
                      />
                    </TableCell>
                  </TableRow>
                ))}
                {!groups.length && (
                  <TableRow>
                    <TableCell colSpan={5} align="center">
                      {loading ? '正在加载分组...' : '当前模型暂无分组，请点击“新增分组”创建'}
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </TableContainer>
        </Card>
      </Stack>
    </AdminContainer>
  );
};

export default ModelGroupAdmin;

