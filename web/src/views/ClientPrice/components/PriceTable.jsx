import {
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  TableContainer,
  Paper,
  Box,
  Typography
} from '@mui/material';
import PriceRow from './PriceRow';
import PriceInputMode from './PriceInputMode';

const PriceTable = ({
  groups,
  inputMode,
  onInputModeChange,
  defaultPrice,
  loading,
  onToggleCustomerPrice,
  onTogglePermitted,
  onTypeChange,
  onDiscountChange,
  onPriceRmbChange,
  onRateChange,
  currency = 'CNY'
}) => {
  // 根据输入模式确定列标题
  const getPriceColumnHeaders = () => {
    switch (inputMode) {
      case 'discount':
        return <TableCell align="right">折扣（%）</TableCell>;
      case 'rmb':
        return (
          <>
            <TableCell align="right">
              输入价格（每百万 tokens / 次 / 秒，{currency === 'USD' ? 'USD' : 'RMB'}）
            </TableCell>
            <TableCell align="right">
              输出价格（每百万 tokens / 次 / 秒，{currency === 'USD' ? 'USD' : 'RMB'}）
            </TableCell>
          </>
        );
      default:
        return null;
    }
  };

  return (
    <Box>
      <Box sx={{ p: 2, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Typography variant="subtitle1">价格配置</Typography>
        <PriceInputMode value={inputMode} onChange={onInputModeChange} />
      </Box>
      <TableContainer component={Paper}>
        <Table size="small">
          <TableHead>
            <TableRow>
              <TableCell>分组编码</TableCell>
              <TableCell>名称</TableCell>
              <TableCell>默认分组</TableCell>
              <TableCell>计费类型</TableCell>
              <TableCell>启用客户价</TableCell>
              <TableCell>价格类型</TableCell>
              <TableCell>全局价格（参考）</TableCell>
              {getPriceColumnHeaders()}
              <TableCell>授权</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {groups.map((row, idx) => (
              <PriceRow
                key={row.group_code || idx}
                row={row}
                index={idx}
                inputMode={inputMode}
                 currency={currency}
                defaultPrice={defaultPrice}
                onToggleCustomerPrice={onToggleCustomerPrice}
                onTogglePermitted={onTogglePermitted}
                onTypeChange={onTypeChange}
                onDiscountChange={onDiscountChange}
                onPriceRmbChange={onPriceRmbChange}
                onRateChange={onRateChange}
              />
            ))}
            {!groups.length && (
              <TableRow>
                <TableCell colSpan={10} align="center">
                  {loading ? '正在加载配置...' : '请先选择用户与模型'}
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default PriceTable;
