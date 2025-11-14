import { useMemo } from 'react';
import {
  TableRow,
  TableCell,
  TextField,
  Switch,
  FormControlLabel,
  Select,
  MenuItem,
  FormControl,
  Chip,
  Typography,
  Box
} from '@mui/material';
import { convertPrice } from '../../Pricing/component/priceConverter';

const PriceRow = ({
  row,
  index,
  inputMode,
  currency = 'CNY',
  defaultPrice,
  onToggleCustomerPrice,
  onTogglePermitted,
  onTypeChange,
  onDiscountChange,
  onPriceRmbChange,
  onRateChange
}) => {
  // 计算显示的折扣值
  const discountValue = useMemo(() => {
    if (!defaultPrice) return '';
    const base = row.type === 'tokens' ? defaultPrice.input || 0 : defaultPrice.output || 0;
    if (!base) return '';
    const current = row.type === 'tokens' ? row.input_rate || 0 : row.output_rate || 0;
    if (!current) return '';
    return ((current / base) * 100).toFixed(0);
  }, [defaultPrice, row.type, row.input_rate, row.output_rate]);

  const targetCurrencyType = currency === 'USD' ? 'USD' : 'RMB';
  const symbol = currency === 'USD' ? '$' : '¥';

  // 计算显示的货币价格：tokens 按每百万 tokens，times/seconds 按每次/秒
  const inputFiatValue = useMemo(() => {
    const unit = row.type === 'tokens' ? 'M' : 'K';
    const rate = row.input_rate || 0;
    return convertPrice(rate, 'rate', 'K', targetCurrencyType, unit);
  }, [row.type, row.input_rate, targetCurrencyType]);

  const outputFiatValue = useMemo(() => {
    const unit = row.type === 'tokens' ? 'M' : 'K';
    const rate = row.output_rate || 0;
    return convertPrice(rate, 'rate', 'K', targetCurrencyType, unit);
  }, [row.type, row.output_rate, targetCurrencyType]);

  // 计算全局价格显示
  const globalPriceDisplay = useMemo(() => {
    if (!defaultPrice) return null;
    
    const unit = row.type === 'tokens' ? 'M' : 'K';
    
    if (row.type === 'tokens') {
      return {
        input: `${defaultPrice.input?.toFixed(4) || '0'} (¥${convertPrice(
          defaultPrice.input || 0,
          'rate',
          'K',
          'RMB',
          'M'
        ).toFixed(2)}/M)`,
        output: `${defaultPrice.output?.toFixed(4) || '0'} (¥${convertPrice(
          defaultPrice.output || 0,
          'rate',
          'K',
          'RMB',
          'M'
        ).toFixed(2)}/M)`
      };
    } else {
      return {
        output: `${defaultPrice.output?.toFixed(4) || '0'} (¥${convertPrice(
          defaultPrice.output || 0,
          'rate',
          'K',
          'RMB',
          'K'
        ).toFixed(2)}/${row.type === 'times' ? '次' : '秒'})`
      };
    }
  }, [defaultPrice, row.type]);

  return (
    <TableRow hover>
      <TableCell>
        <Chip label={row.group_code} size="small" variant="outlined" />
      </TableCell>
      <TableCell>{row.display_name || '-'}</TableCell>
      <TableCell>
        {row.is_default && <Chip label="默认" size="small" color="primary" />}
      </TableCell>
      <TableCell>
        <Chip label={row.billing_type || '-'} size="small" />
      </TableCell>
      <TableCell>
        <FormControlLabel
          control={
            <Switch
              size="small"
              checked={!!row.has_customer_price}
              onChange={onToggleCustomerPrice(index)}
              aria-label={`启用 ${row.display_name} 客户价`}
            />
          }
          label=""
        />
      </TableCell>
      <TableCell>
        <FormControl size="small" sx={{ minWidth: 100 }}>
          <Select
            value={row.type || 'tokens'}
            onChange={onTypeChange(index)}
            disabled={!row.has_customer_price}
            aria-label="价格类型"
          >
            <MenuItem value="tokens">tokens</MenuItem>
            <MenuItem value="times">times</MenuItem>
            <MenuItem value="seconds">seconds</MenuItem>
          </Select>
        </FormControl>
      </TableCell>
      
      {/* 全局价格列 */}
      <TableCell>
        {globalPriceDisplay ? (
          <Box>
            {row.type === 'tokens' && (
              <>
                <Typography variant="caption" color="text.secondary" display="block">
                  输入: {globalPriceDisplay.input}
                </Typography>
                <Typography variant="caption" color="text.secondary" display="block">
                  输出: {globalPriceDisplay.output}
                </Typography>
              </>
            )}
            {(row.type === 'times' || row.type === 'seconds') && (
              <Typography variant="caption" color="text.secondary">
                {globalPriceDisplay.output}
              </Typography>
            )}
          </Box>
        ) : (
          <Typography variant="caption" color="text.disabled">
            无全局价格
          </Typography>
        )}
      </TableCell>

      {/* 根据输入模式显示不同的输入字段 */}
      {inputMode === 'discount' && (
        <>
          <TableCell align="right">
            <TextField
              type="number"
              size="small"
              disabled={!row.has_customer_price || !defaultPrice}
              value={discountValue}
              onChange={onDiscountChange(index)}
              inputProps={{ min: 0, max: 200, step: '1' }}
              placeholder={defaultPrice ? '80' : '无默认价'}
              helperText={!defaultPrice && row.has_customer_price ? '无默认价' : ''}
              error={!defaultPrice && row.has_customer_price}
              sx={{ width: 100 }}
              aria-label="折扣百分比"
            />
          </TableCell>
        </>
      )}

      {inputMode === 'rmb' && (
        <>
          <TableCell align="right">
            <TextField
              type="number"
              size="small"
              disabled={!row.has_customer_price}
              value={inputFiatValue}
              onChange={onPriceRmbChange(index, 'input')}
              inputProps={{ min: 0, step: '0.000001' }}
              placeholder={
                row.type === 'tokens'
                  ? `${symbol}/百万 tokens`
                  : `${symbol}/${row.type === 'times' ? '次' : '秒'}`
              }
              sx={{ width: 120 }}
              aria-label="输入价格（按币种）"
            />
          </TableCell>
          <TableCell align="right">
            <TextField
              type="number"
              size="small"
              disabled={!row.has_customer_price}
              value={outputFiatValue}
              onChange={onPriceRmbChange(index, 'output')}
              inputProps={{ min: 0, step: '0.000001' }}
              placeholder={
                row.type === 'tokens'
                  ? `${symbol}/百万 tokens`
                  : `${symbol}/${row.type === 'times' ? '次' : '秒'}`
              }
              sx={{ width: 120 }}
              aria-label="输出价格（按币种）"
            />
          </TableCell>
        </>
      )}

      <TableCell>
        <FormControlLabel
          control={
            <Switch
              size="small"
              checked={row.is_default ? true : !!row.permitted}
              onChange={onTogglePermitted(index)}
              disabled={row.is_default}
              aria-label={`授权 ${row.display_name}`}
            />
          }
          label=""
        />
      </TableCell>
    </TableRow>
  );
};

export default PriceRow;
