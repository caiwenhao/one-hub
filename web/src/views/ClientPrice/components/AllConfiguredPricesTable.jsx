import {
  Card,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  TableContainer,
  Chip,
  Typography,
  Box,
  Tooltip,
  Stack
} from '@mui/material';
import { Icon } from '@iconify/react';
import { convertPrice } from '../../Pricing/component/priceConverter';

// 客户价总览表：展示某个用户下所有已配置客户价（模型 × 分组）
const AllConfiguredPricesTable = ({ rows, loading }) => {
  if (loading) {
    return (
      <Card sx={{ p: 2 }}>
        <Typography variant="body2" color="text.secondary">
          正在加载已配置的客户价...
        </Typography>
      </Card>
    );
  }

  if (!rows || rows.length === 0) {
    return (
      <Card sx={{ p: 2 }}>
        <Typography variant="body2" color="text.secondary">
          暂无已配置的客户价记录。
        </Typography>
      </Card>
    );
  }

  const getDiscountChipColor = (discount) => {
    if (discount === null || discount === undefined) return 'default';
    if (discount < 100) return 'success';
    if (discount === 100) return 'default';
    if (discount > 100 && discount <= 150) return 'warning';
    return 'error';
  };

  const formatPriceDisplay = (type, rate) => {
    if (!rate) return '-';
    if (type === 'tokens') {
      const rmb = convertPrice(rate, 'rate', 'K', 'RMB', 'M');
      return `¥${rmb.toFixed(4)}/百万 tokens`;
    }
    const rmb = convertPrice(rate, 'rate', 'K', 'RMB', 'K');
    const unit = type === 'times' ? '次' : '秒';
    return `¥${rmb.toFixed(4)}/${unit}`;
  };

  return (
    <Card>
      <TableContainer>
        <Table size="small">
          <TableHead>
            <TableRow>
              <TableCell>模型</TableCell>
              <TableCell>分组</TableCell>
              <TableCell>计费类型</TableCell>
              <TableCell>原价</TableCell>
              <TableCell>客户价</TableCell>
              <TableCell>折扣</TableCell>
              <TableCell>状态</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row, idx) => {
              const discount =
                row.default_rate && row.custom_rate
                  ? (row.custom_rate / row.default_rate) * 100
                  : null;

              return (
                <TableRow key={`${row.model}-${row.group_code}-${idx}`} hover>
                  <TableCell>
                    <Stack direction="row" spacing={1} alignItems="center">
                      <Icon icon="solar:cpu-bolt-bold-duotone" width={18} />
                      <Box>
                        <Typography variant="body2" fontWeight={600}>
                          {row.model}
                        </Typography>
                      </Box>
                    </Stack>
                  </TableCell>
                  <TableCell>
                    <Stack direction="row" spacing={1} alignItems="center">
                      <Chip label={row.group_code} size="small" variant="outlined" />
                      {row.is_default && <Chip label="默认" size="small" color="primary" />}
                    </Stack>
                  </TableCell>
                  <TableCell>
                    <Chip label={row.type || '-'} size="small" />
                  </TableCell>
                  <TableCell>
                    <Tooltip
                      title={
                        <Box>
                          <Typography variant="caption" display="block">
                            Rate(K): {row.default_rate?.toFixed(6) || '-'}
                          </Typography>
                        </Box>
                      }
                    >
                      <Typography variant="body2" color="text.secondary">
                        {formatPriceDisplay(row.type, row.default_rate)}
                      </Typography>
                    </Tooltip>
                  </TableCell>
                  <TableCell>
                    <Tooltip
                      title={
                        <Box>
                          <Typography variant="caption" display="block">
                            Rate(K): {row.custom_rate?.toFixed(6) || '-'}
                          </Typography>
                        </Box>
                      }
                    >
                      <Typography variant="body2">
                        {formatPriceDisplay(row.type, row.custom_rate)}
                      </Typography>
                    </Tooltip>
                  </TableCell>
                  <TableCell>
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
                  <TableCell>
                    {row.is_default ? (
                      <Chip label="默认分组" size="small" variant="outlined" />
                    ) : row.permitted ? (
                      <Chip label="已授权" size="small" color="success" variant="outlined" />
                    ) : (
                      <Chip label="未授权" size="small" color="warning" variant="outlined" />
                    )}
                  </TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </Card>
  );
};

export default AllConfiguredPricesTable;
