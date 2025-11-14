import { useState, useEffect, useCallback, useMemo } from 'react';
import { Card, Stack, Typography, Button, LinearProgress, Box, Divider, Tabs, Tab, TextField } from '@mui/material';
import { Icon } from '@iconify/react';
import { API } from 'utils/api';
import { showError, showSuccess } from 'utils/common';
import ActionBar from 'ui-component/ActionBar';
import AdminContainer from 'ui-component/AdminContainer';
import UserSelector from './components/UserSelector';
import ModelSelector from './components/ModelSelector';
import DefaultPriceCard from './components/DefaultPriceCard';
import PriceTable from './components/PriceTable';
import QuickStats from './components/QuickStats';
import ConfiguredPricesList from './components/ConfiguredPricesList';
import { convertPrice } from '../Pricing/component/priceConverter';
import ToggleButtonGroup from 'ui-component/ToggleButton';
import { getCurrencyPreference, setCurrencyPreference, CURRENCY_OPTIONS, CURRENCY_PAGE_KEYS } from 'utils/currencyPreferences';

// 客户价配置页面（管理员）
// - 基于 /api/customer_pricing/model_groups 与 /api/customer_pricing/users/:id/model_pricing
// - 支持按「用户 + 模型」维度查看与编辑分组客户价与授权信息

const ClientPrice = () => {
  const [selectedUser, setSelectedUser] = useState(null);
  const [modelList, setModelList] = useState([]);
  const [model, setModel] = useState('');
  const [loading, setLoading] = useState(false);
  const [loadingModels, setLoadingModels] = useState(false);
  const [loadingConfigured, setLoadingConfigured] = useState(false);
  const [saving, setSaving] = useState(false);
  const [groups, setGroups] = useState([]);
  const [allPrices, setAllPrices] = useState([]);
  const [defaultPrice, setDefaultPrice] = useState(null);
  const [inputMode, setInputMode] = useState('discount'); // discount | rmb | rate
  const [configuredPrices, setConfiguredPrices] = useState([]);
  const [activeTab, setActiveTab] = useState(0); // 0: 已配置列表, 1: 编辑配置
  const [currency, setCurrency] = useState(() => getCurrencyPreference(CURRENCY_PAGE_KEYS.CLIENT_PRICE));
  const [listKeyword, setListKeyword] = useState('');

  // 加载所有用户的已配置客户价（概览）
  // MVP 实现：前端聚合 /api/user/ 与 /api/customer_pricing/users/:id/configured_prices
  // 后续可在后端提供聚合接口以优化性能与分页体验
  const loadConfiguredPrices = useCallback(async () => {
    setLoadingConfigured(true);
    try {
      // 1. 拉取一批用户（MVP：取前 100 个，后续可做分页或条件过滤）
      const userRes = await API.get('/api/user/', {
        params: {
          page: 1,
          size: 100
        }
      });
      const { success: userSuccess, message: userMessage, data: userData } = userRes.data || {};
      if (!userSuccess) {
        showError(userMessage || '加载用户列表失败');
        setConfiguredPrices([]);
        return;
      }
      const users = userData?.data || [];
      if (!users.length) {
        setConfiguredPrices([]);
        return;
      }

      // 2. 并发拉取每个用户的已配置客户价
      const requests = users.map((user) =>
        API.get(`/api/customer_pricing/users/${user.id}/configured_prices`)
          .then((res) => ({ user, res }))
          .catch((e) => {
            console.error(e);
            return { user, res: null };
          })
      );

      const results = await Promise.all(requests);

      const aggregated = [];
      results.forEach(({ user, res }) => {
        if (!res) return;
        const { success, data } = res.data || {};
        if (!success || !Array.isArray(data) || data.length === 0) {
          return;
        }
        data.forEach((item) => {
          aggregated.push({
            ...item,
            user: {
              id: user.id,
              username: user.username,
              display_name: user.display_name
            }
          });
        });
      });

      setConfiguredPrices(aggregated);
    } catch (e) {
      console.error(e);
      // 若部分接口尚未就绪，整体视为无数据但不阻塞页面
      setConfiguredPrices([]);
    } finally {
      setLoadingConfigured(false);
    }
  }, []);

  // 用户选择变化处理
  const handleUserChange = useCallback((user) => {
    setSelectedUser(user);
    setModel('');
    setGroups([]);
    setDefaultPrice(null);
    setActiveTab(0); // 切换到已配置列表
  }, []);

  // 拉取全局默认价格列表（用于展示与折扣换算）
  useEffect(() => {
    const fetchPrices = async () => {
      if (allPrices.length) return;
      try {
        const res = await API.get('/api/prices');
        const { success, message, data } = res.data || {};
        if (success) {
          setAllPrices(data || []);
        } else {
          showError(message || '加载全局价格失败');
        }
      } catch (e) {
        console.error(e);
        showError(e.message || '加载全局价格失败');
      }
    };
    fetchPrices();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // 页面初始化时加载所有用户的已配置客户价概览
  useEffect(() => {
    loadConfiguredPrices();
  }, [loadConfiguredPrices]);

  // 已配置列表搜索过滤（按用户 ID / 用户名 / 显示名 / 模型）
  const filteredConfiguredPrices = useMemo(() => {
    const keyword = (listKeyword || '').trim().toLowerCase();
    if (!keyword) return configuredPrices;
    return configuredPrices.filter((item) => {
      const user = item.user || {};
      const idMatch =
        user.id != null && user.id.toString().includes(keyword);
      const usernameMatch =
        user.username && user.username.toLowerCase().includes(keyword);
      const displayNameMatch =
        user.display_name && user.display_name.toLowerCase().includes(keyword);
      const modelMatch =
        item.model && item.model.toLowerCase().includes(keyword);
      return idMatch || usernameMatch || displayNameMatch || modelMatch;
    });
  }, [configuredPrices, listKeyword]);

  const handleCurrencyChange = useCallback((event, newCurrency) => {
    if (newCurrency) {
      setCurrency(newCurrency);
      setCurrencyPreference(newCurrency, CURRENCY_PAGE_KEYS.CLIENT_PRICE);
    }
  }, []);

  // 根据用户 ID 拉取"已接入渠道且对该用户分组可用"的模型列表
  useEffect(() => {
    const uid = selectedUser?.id;
    if (!uid || uid <= 0) {
      setModelList([]);
      return;
    }
    const fetchModelListForUser = async () => {
      setLoadingModels(true);
      try {
        const res = await API.get(`/api/customer_pricing/users/${uid}/available_models`);
        const { success, message, data } = res.data || {};
        if (success) {
          setModelList(data || []);
        } else {
          showError(message || '加载可用模型失败');
          setModelList([]);
        }
      } catch (e) {
        console.error(e);
        showError(e.message || '加载可用模型失败');
        setModelList([]);
      } finally {
        setLoadingModels(false);
      }
    };
    fetchModelListForUser();
  }, [selectedUser]);

  // 选择模型后自动加载配置
  useEffect(() => {
    if (model && selectedUser?.id) {
      handleLoad();
    } else {
      setGroups([]);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [model, selectedUser]);

  // 根据当前模型，计算默认价（官方价）
  useEffect(() => {
    if (!model) {
      setDefaultPrice(null);
      return;
    }
    if (!allPrices.length) {
      setDefaultPrice(null);
      return;
    }
    const p = allPrices.find((item) => item.model === model);
    setDefaultPrice(p || null);
  }, [model, allPrices]);

  const handleLoad = useCallback(async () => {
    const uid = selectedUser?.id;
    if (!uid || uid <= 0) {
      showError('请先选择用户');
      return;
    }
    if (!model) {
      showError('请先选择模型');
      return;
    }
    setLoading(true);
    try {
      const res = await API.get(`/api/customer_pricing/users/${uid}/model_pricing`, {
        params: { model }
      });
      const { success, message, data } = res.data || {};
      if (!success) {
        showError(message || '加载客户价配置失败');
        setGroups([]);
        return;
      }
      const list = Array.isArray(data) ? data : [];
      // 统一使用 type 字段，从 billing_type 或 type 中取值
      const mapped = list.map((it) => ({
        ...it,
        type: it.type || it.billing_type || 'tokens',
        has_customer_price: !!it.has_customer_price,
        permitted: !!it.permitted
      }));
      setGroups(mapped);
    } catch (e) {
      console.error(e);
      showError(e.message || '加载客户价配置失败');
      setGroups([]);
    } finally {
      setLoading(false);
    }
  }, [selectedUser, model]);

  const handleToggleCustomerPrice = useCallback((index) => (event) => {
    const checked = event.target.checked;
    setGroups((prev) => {
      const next = [...prev];
      const row = { ...next[index], has_customer_price: checked };
      // 关闭客户价时清空 rate，避免误保存旧值
      if (!checked) {
        row.input_rate = 0;
        row.output_rate = 0;
      }
      next[index] = row;
      return next;
    });
  }, []);

  const handleTogglePermitted = useCallback((index) => (event) => {
    const checked = event.target.checked;
    setGroups((prev) => {
      const next = [...prev];
      const row = { ...next[index] };
      // 默认分组始终视为可用，不允许关闭
      if (row.is_default) {
        row.permitted = true;
      } else {
        row.permitted = checked;
      }
      next[index] = row;
      return next;
    });
  }, []);

  const handleTypeChange = useCallback((index) => (event) => {
    const value = event.target.value;
    setGroups((prev) => {
      const next = [...prev];
      const row = { ...next[index], type: value };
      next[index] = row;
      return next;
    });
  }, []);

  // 折扣（%）修改：基于 defaultPrice 计算新的 Rate(K)
  const handleDiscountChange = useCallback(
    (index) => (event) => {
      const value = event.target.value;
      const discount = Number(value);
      if (Number.isNaN(discount)) {
        return;
      }
      if (!defaultPrice) {
        return;
      }
      setGroups((prev) => {
        const next = [...prev];
        const row = { ...next[index] };
        const d = discount / 100;
        if (row.type === 'tokens') {
          // tokens 类型：输入/输出均按默认 tokens 价打折
          row.input_rate = (defaultPrice.input || 0) * d;
          row.output_rate = (defaultPrice.output || 0) * d;
        } else {
          // times/seconds：按默认输出价打折
          row.input_rate = (defaultPrice.input || 0) * d;
          row.output_rate = (defaultPrice.output || 0) * d;
        }
        next[index] = row;
        return next;
      });
    },
    [defaultPrice]
  );

  // 单价修改：根据不同类型把「每百万 tokens / 次 / 秒」视角转换为内部 Rate(K)
  const handlePriceRmbChange = useCallback((index, field) => (event) => {
    const raw = event.target.value;
    const price = Number(raw);
    if (Number.isNaN(price)) {
      return;
    }
    setGroups((prev) => {
      const next = [...prev];
      const row = { ...next[index] };
      const unit = row.type === 'tokens' ? 'M' : 'K';
      const fromCurrency = currency === 'USD' ? 'USD' : 'RMB';
      const rate = convertPrice(price || 0, fromCurrency, unit, 'rate', 'K');
      if (field === 'input') {
        row.input_rate = rate;
      } else {
        row.output_rate = rate;
      }
      next[index] = row;
      return next;
    });
  }, [currency]);

  const handleRateChange = useCallback((index, field) => (event) => {
    const value = event.target.value;
    setGroups((prev) => {
      const next = [...prev];
      const row = { ...next[index], [field]: value === '' ? '' : Number(value) };
      next[index] = row;
      return next;
    });
  }, []);

  const handleSave = useCallback(async () => {
    const uid = selectedUser?.id;
    if (!uid || uid <= 0) {
      showError('请先选择用户');
      return;
    }
    if (!model) {
      showError('请先选择模型');
      return;
    }
    if (!groups.length) {
      showError('请先加载并编辑分组配置');
      return;
    }

    setSaving(true);
    try {
      const payload = {
        model,
        groups: groups.map((g) => ({
          model: g.model,
          group_code: g.group_code,
          display_name: g.display_name,
          description: g.description,
          billing_type: g.billing_type,
          is_default: g.is_default,
          has_customer_price: !!g.has_customer_price,
          type: g.type,
          input_rate: Number(g.input_rate) || 0,
          output_rate: Number(g.output_rate) || 0,
          permitted: !!g.permitted
        }))
      };
      const res = await API.post(`/api/customer_pricing/users/${uid}/model_pricing`, payload);
      const { success, message } = res.data || {};
      if (!success) {
        showError(message || '保存失败');
      } else {
        showSuccess('保存成功');
        // 重新加载当前模型配置
        handleLoad();
        // 刷新全局已配置客户价概览
        loadConfiguredPrices();
      }
    } catch (e) {
      console.error(e);
      showError(e.message || '保存失败');
    } finally {
      setSaving(false);
    }
  }, [selectedUser, model, groups, handleLoad, loadConfiguredPrices]);

  // 从已配置列表编辑某个模型
  const handleEditFromList = useCallback((modelName, user) => {
    if (user) {
      setSelectedUser(user);
    }
    setModel(modelName);
    setActiveTab(1); // 切换到编辑标签页
    // 自动加载该模型的配置
    setTimeout(() => {
      handleLoad();
    }, 100);
  }, [handleLoad]);

  return (
    <AdminContainer>
      <Stack spacing={3}>
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={1}>
          <Stack direction="column" spacing={1}>
            <Typography variant="h2">客户价配置</Typography>
            <Stack direction="row" alignItems="center" spacing={2} flexWrap="wrap">
              <Typography variant="subtitle1" color="text.secondary">
                管理所有用户的「模型 × 分组」客户价与授权（含按秒计费）。
              </Typography>
              <Stack direction="row" alignItems="center" spacing={1}>
                <Typography variant="body2" color="text.secondary">
                  价格币种：
                </Typography>
                <ToggleButtonGroup
                  value={currency}
                  onChange={handleCurrencyChange}
                  options={CURRENCY_OPTIONS}
                  aria-label="currency toggle"
                  size="small"
                  sx={{
                    '& .MuiToggleButtonGroup-grouped': {
                      borderRadius: '8px !important',
                      mx: 0.5,
                      border: 0
                    }
                  }}
                />
              </Stack>
            </Stack>
          </Stack>
        </Stack>

        {(loading || loadingConfigured) && <LinearProgress />}

        {/* 标签页切换：0 = 全局已配置列表，1 = 针对用户编辑配置 */}
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={activeTab} onChange={(_, newValue) => setActiveTab(newValue)}>
            <Tab
              label={
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                  <Icon icon="solar:list-bold-duotone" width={20} />
                  已配置列表（全部用户）
                  {configuredPrices.length > 0 && (
                    <Box
                      component="span"
                      sx={{
                        bgcolor: 'primary.main',
                        color: 'white',
                        borderRadius: '12px',
                        px: 1,
                        py: 0.25,
                        fontSize: '0.75rem',
                        fontWeight: 600
                      }}
                    >
                      {configuredPrices.length}
                    </Box>
                  )}
                </Box>
              }
            />
            <Tab
              label={
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                  <Icon icon="solar:settings-bold-duotone" width={20} />
                  编辑配置（按用户）
                </Box>
              }
            />
          </Tabs>
        </Box>

        {/* Tab 0: 全部用户的已配置客户价概览 */}
        {activeTab === 0 && (
          <Box sx={{ py: 2 }}>
            {/* 顶部统计 + 搜索过滤 */}
            <Stack
              direction={{ xs: 'column', sm: 'row' }}
              spacing={2}
              alignItems={{ xs: 'stretch', sm: 'center' }}
              justifyContent="space-between"
              sx={{ mb: 2 }}
            >
              <QuickStats configuredPrices={filteredConfiguredPrices} />
              <Box sx={{ minWidth: { xs: '100%', sm: 260 } }}>
                <TextField
                  fullWidth
                  size="small"
                  label="搜索已配置客户（用户 / 模型）"
                  placeholder="输入用户ID、用户名、昵称或模型名"
                  value={listKeyword}
                  onChange={(e) => setListKeyword(e.target.value)}
                />
              </Box>
            </Stack>

            <ConfiguredPricesList
              configuredPrices={filteredConfiguredPrices}
              onEdit={handleEditFromList}
              loading={loadingConfigured}
              allPrices={allPrices}
              currency={currency}
            />
          </Box>
        )}

        {/* Tab 1: 按用户编辑客户价配置 */}
        {activeTab === 1 && (
          <Stack spacing={3}>
            {/* 用户选择器（仅在编辑配置 Tab 内展示） */}
            <Card sx={{ p: 2 }}>
              <Stack spacing={1.5}>
                <Typography variant="subtitle2" color="text.secondary">
                  请选择需要配置客户价的用户
                </Typography>
                <UserSelector
                  value={selectedUser}
                  onChange={handleUserChange}
                  disabled={loading || saving}
                />
              </Stack>
            </Card>

            {!selectedUser && (
              <Card sx={{ p: 4, textAlign: 'center' }}>
                <Icon icon="solar:user-circle-bold-duotone" width={64} style={{ opacity: 0.3 }} />
                <Typography variant="h6" color="text.secondary" sx={{ mt: 2 }}>
                  请先在上方选择一个用户
                </Typography>
                <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                  选择用户后，可在此标签页按模型配置该用户的客户价与授权。
                </Typography>
              </Card>
            )}

            {selectedUser && (
              <Stack spacing={3}>
                <Card sx={{ p: 2 }}>
                  <Stack spacing={2}>
                    <Typography variant="subtitle2" color="text.secondary">
                      选择模型后将自动加载配置
                    </Typography>
                    <ModelSelector
                      value={model}
                      onChange={(newModel) => {
                        setModel(newModel);
                      }}
                      options={modelList}
                      disabled={loading || saving}
                      loading={loadingModels}
                    />
                  </Stack>
                </Card>

                {model && (
                  <>
                    <DefaultPriceCard defaultPrice={defaultPrice} />

                    <Card>
                      <PriceTable
                        groups={groups}
                        currency={currency}
                        inputMode={inputMode}
                        onInputModeChange={setInputMode}
                        defaultPrice={defaultPrice}
                        loading={loading}
                        onToggleCustomerPrice={handleToggleCustomerPrice}
                        onTogglePermitted={handleTogglePermitted}
                        onTypeChange={handleTypeChange}
                        onDiscountChange={handleDiscountChange}
                        onPriceRmbChange={handlePriceRmbChange}
                        onRateChange={handleRateChange}
                      />
                    </Card>

                    {/* 保存按钮放到底部，靠右对齐 */}
                    <ActionBar sx={{ py: 0, justifyContent: 'flex-end' }}>
                      <Button
                        color="primary"
                        variant="contained"
                        onClick={handleSave}
                        disabled={saving || !groups.length}
                        startIcon={<Icon icon="solar:floppy-disk-bold-duotone" />}
                      >
                        {saving ? '保存中...' : '保存配置'}
                      </Button>
                    </ActionBar>
                  </>
                )}

                {!model && (
                  <Card sx={{ p: 4, textAlign: 'center' }}>
                    <Icon icon="solar:document-text-bold-duotone" width={64} style={{ opacity: 0.3 }} />
                    <Typography variant="h6" color="text.secondary" sx={{ mt: 2 }}>
                      请选择模型
                    </Typography>
                    <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                      选择模型后将自动加载该模型的价格配置
                    </Typography>
                  </Card>
                )}
              </Stack>
            )}
          </Stack>
        )}
      </Stack>
    </AdminContainer>
  );
};

export default ClientPrice;
