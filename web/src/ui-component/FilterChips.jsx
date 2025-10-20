// 筛选区选择项标签化展示 + 一键清除
import { Chip, Stack, Button } from '@mui/material';

export default function FilterChips({ items = [], onRemove, onClear }) {
  const hasItems = Array.isArray(items) && items.length > 0;
  if (!hasItems) return null;
  return (
    <Stack direction="row" spacing={1} alignItems="center" flexWrap="wrap">
      {items.map((it) => (
        <Chip key={`${it.key}:${it.value}`} label={`${it.label || it.key}: ${it.value}`} onDelete={() => onRemove?.(it)} />
      ))}
      <Button size="small" variant="text" color="primary" onClick={onClear} sx={{ ml: 0.5 }}>
        清除全部
      </Button>
    </Stack>
  );
}
