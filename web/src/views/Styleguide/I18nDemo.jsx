import { Stack, Typography, Box } from '@mui/material';
import { formatNumber, formatDate, joinWithUnit } from 'utils/i18nFormat';

export default function I18nDemo() {
  const n = 12345.678;
  const d = Date.now();
  return (
    <Stack spacing={1.5}>
      <Typography variant="subtitle2">数字/日期本地化</Typography>
      <Typography variant="body2">{formatNumber(n)} · {formatDate(d)}</Typography>
      <Typography variant="subtitle2">单位与空格（不换行空格）</Typography>
      <Typography variant="body2">{joinWithUnit(formatNumber(120), 'ms')}</Typography>
      <Typography variant="subtitle2">长度回退演示</Typography>
      <Box className="l10n-ellipsis" sx={{ maxWidth: 240 }}>一段很长很长的中文文本用于演示省略与 Tooltip</Box>
      <Box className="l10n-break" sx={{ maxWidth: 240 }}>supercalifragilisticexpialidocious-国际化-示例-超长词语换行</Box>
    </Stack>
  );
}
