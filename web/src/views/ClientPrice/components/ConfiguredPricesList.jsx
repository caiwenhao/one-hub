import { useState } from 'react';
import {
  Box,
  Card,
  Typography,
  Chip,
  IconButton,
  Collapse,
  Table,
  TableBody,
  TableRow,
  TableCell,
  TableHead,
  Stack,
  Tooltip,
  Alert
} from '@mui/material';
import { Icon } from '@iconify/react';
import { convertPrice } from '../../Pricing/component/priceConverter';

const ConfiguredPricesList = ({ configuredPrices, onEdit, loading, allPrices = [], currency = 'CNY' }) => {
  const [expandedModels, setExpandedModels] = useState({});

  const toggleExpand = (key) => {
    setExpandedModels((prev) => ({
      ...prev,
      [key]: !prev[key]
    }));
  };

  if (loading) {
    return (
      <Card sx={{ p: 2 }}>
        <Typography variant="body2" color="text.secondary">
          正在加载已配置的客户价...
        </Typography>
      </Card>
    );
  }

  if (!configuredPrices || configuredPrices.length === 0) {
    return (
      <Alert severity="info" sx={{ mb: 2 }}>
        暂无已配置的客户价记录。
      </Alert>
    );
  }

  return (
    <Box>
      <Typography variant="h6" gutterBottom sx={{ mb: 2 }}>
        已配置的客户价（{configuredPrices.length} 个「用户 × 模型」）
      </Typography>
      <Stack spacing={1.5}>
        {configuredPrices.map((item) => {
          const key = `${item.user?.id || 'unknown'}:${item.model}`;
          const isExpanded = expandedModels[key];
          const hasCustomGroups = item.groups.filter((g) => g.has_customer_price).length;
          const basePrice = allPrices.find((p) => p.model === item.model) || {};

          return (
            <Card
              key={item.model}
              sx={{
                border: '1px solid',
                borderColor: 'divider',
                '&:hover': {
                  borderColor: 'primary.main',
                  boxShadow: 1
                }
              }}
            >
              <Box
                sx={{
                  p: 2,
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'space-between',
                  cursor: 'pointer'
                }}
                onClick={() => toggleExpand(key)}
              >
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, flex: 1 }}>
                  <Icon icon="solar:cpu-bolt-bold-duotone" width={24} />
                  <Box sx={{ flex: 1 }}>
                    <Typography variant="subtitle1" fontWeight={600}>
                      {item.model}
                    </Typography>
                    {item.user && (
                      <Typography variant="body2" color="text.secondary" sx={{ mt: 0.25 }}>
                        用户：{item.user.display_name || item.user.username}（ID: {item.user.id}）
                      </Typography>
                    )}
                    <Box sx={{ display: 'flex', gap: 1, mt: 0.5, flexWrap: 'wrap' }}>
                      <Chip
                        label={`${hasCustomGroups} 个分组已配置`}
                        size="small"
                        color="primary"
                        variant="outlined"
                      />
                      {item.groups.some((g) => g.type === 'seconds') && (
                        <Chip label="按秒计费" size="small" color="warning" variant="outlined" />
                      )}
                      {item.groups.some((g) => g.type === 'times') && (
                        <Chip label="按次计费" size="small" color="info" variant="outlined" />
                      )}
                    </Box>
                  </Box>
                </Box>
                <Box sx={{ display: 'flex', gap: 1 }}>
                  <Tooltip title="编辑配置">
                    <IconButton
                      size="small"
                      color="primary"
                      onClick={(e) => {
                        e.stopPropagation();
                        onEdit(item.model, item.user);
                      }}
                    >
                      <Icon icon="solar:pen-bold-duotone" width={20} />
                    </IconButton>
                  </Tooltip>
                  <IconButton size="small">
                    <Icon
                      icon={isExpanded ? 'solar:alt-arrow-up-bold' : 'solar:alt-arrow-down-bold'}
                      width={20}
                    />
                  </IconButton>
                </Box>
              </Box>

              <Collapse in={isExpanded}>
                <Box sx={{ px: 2, pb: 2 }}>
                  <Table size="small">
                    <TableHead>
                      <TableRow>
                        <TableCell>分组</TableCell>
                        <TableCell>计费类型</TableCell>
                        <TableCell>全局价</TableCell>
                        <TableCell>客户价</TableCell>
                        <TableCell>折扣</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {item.groups
                        .filter((g) => g.has_customer_price)
                        .map((group) => {
                          const type = group.type || group.billing_type || basePrice.type || 'tokens';
                          const symbol = currency === 'USD' ? '$' : '¥';
                          const targetCurrencyType = currency === 'USD' ? 'USD' : 'RMB';

                          const defaultInputRate = basePrice.input || 0;
                          const defaultOutputRate = basePrice.output || 0;
                          const customInputRate = group.input_rate || 0;
                          const customOutputRate = group.output_rate || 0;

                          const formatTokenPrice = (rate, isMillion = true) => {
                            if (!rate) return '-';
                            const unit = isMillion ? 'M' : 'K';
                            const price = convertPrice(rate, 'rate', 'K', targetCurrencyType, unit);
                            const per = isMillion ? '百万 tokens' : '千 tokens';
                            return `${symbol}${price.toFixed(4)}/${per}`;
                          };

                          const formatNonTokenPrice = (rate, unitLabel) => {
                            if (!rate) return '-';
                            const price = convertPrice(rate, 'rate', 'K', targetCurrencyType, 'K');
                            return `${symbol}${price.toFixed(4)}/${unitLabel}`;
                          };

                          const getPriceDisplay = (t, inputRate, outputRate) => {
                            if (t === 'tokens') {
                              return (
                                <Box>
                                  <Typography variant="body2" color="text.secondary">
                                    输入：{formatTokenPrice(inputRate)}
                                  </Typography>
                                  <Typography variant="body2" color="text.secondary">
                                    输出：{formatTokenPrice(outputRate)}
                                  </Typography>
                                </Box>
                              );
                            }
                            const unitLabel = t === 'times' ? '次' : '秒';
                            return (
                              <Typography variant="body2" color="text.secondary">
                                {formatNonTokenPrice(outputRate || inputRate, unitLabel)}
                              </Typography>
                            );
                          };

                          const discountBase = defaultOutputRate || defaultInputRate || 0;
                          const discountCustom = customOutputRate || customInputRate || 0;
                          const discount =
                            discountBase > 0 && discountCustom > 0
                              ? (discountCustom / discountBase) * 100
                              : null;

                          const getDiscountChipColor = (value) => {
                            if (value == null) return 'default';
                            if (value < 100) return 'success';
                            if (value === 100) return 'default';
                            if (value > 100 && value <= 150) return 'warning';
                            return 'error';
                          };

                          return (
                            <TableRow key={group.group_code} hover>
                              <TableCell width="25%">
                                <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                                  <Typography variant="body2" fontWeight={500}>
                                    {group.display_name || group.group_code}
                                  </Typography>
                                  {group.is_default && (
                                    <Chip label="默认" size="small" color="primary" />
                                  )}
                                </Box>
                              </TableCell>
                              <TableCell width="15%">
                                <Chip label={type} size="small" variant="outlined" />
                              </TableCell>
                              <TableCell width="25%">
                                {getPriceDisplay(type, defaultInputRate, defaultOutputRate)}
                              </TableCell>
                              <TableCell width="25%">
                                {getPriceDisplay(type, customInputRate, customOutputRate)}
                              </TableCell>
                              <TableCell width="10%">
                                {discount !== null ? (
                                  <Chip
                                    label={`${discount.toFixed(0)}%`}
                                    size="small"
                                    color={getDiscountChipColor(discount)}
                                    variant="outlined"
                                  />
                                ) : (
                                  <Typography variant="caption" color="text.disabled">
                                    -
                                  </Typography>
                                )}
                              </TableCell>
                            </TableRow>
                          );
                        })}
                    </TableBody>
                  </Table>
                </Box>
              </Collapse>
            </Card>
          );
        })}
      </Stack>
    </Box>
  );
};

export default ConfiguredPricesList;
