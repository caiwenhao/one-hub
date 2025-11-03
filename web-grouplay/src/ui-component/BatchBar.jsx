// 批量操作条（粘顶），在选择项>0时显示
import { Box, Stack, Typography, Button } from '@mui/material';

export default function BatchBar({ selected = 0, onClear, actions = [] }) {
  if (!selected || selected <= 0) return null;
  return (
    <Box sx={{ position: 'sticky', top: 0, zIndex: 3, bgcolor: 'background.default', borderBottom: (t) => `1px solid ${t.palette.divider}` }}>
      <Stack direction="row" spacing={1} alignItems="center" justifyContent="space-between" sx={{ p: 1 }}>
        <Typography variant="body2">已选择 {selected} 项</Typography>
        <Stack direction="row" spacing={1}>
          {actions}
          <Button size="small" onClick={onClear} color="primary" variant="text">清除选择</Button>
        </Stack>
      </Stack>
    </Box>
  );
}
