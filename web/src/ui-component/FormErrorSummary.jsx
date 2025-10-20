// 表单顶部错误汇总容器（可锚点跳转）
import { Alert, Link, Stack, Typography } from '@mui/material';

export default function FormErrorSummary({ errors = {}, onJump }) {
  const keys = Object.keys(errors || {});
  if (!keys.length) return null;
  return (
    <Alert severity="error" sx={{ mb: 2 }}>
      <Stack spacing={0.5}>
        <Typography variant="subtitle2">请修正以下字段：</Typography>
        <Stack direction="row" spacing={1} flexWrap="wrap">
          {keys.map((k) => (
            <Link
              key={k}
              component="button"
              variant="body2"
              underline="hover"
              onClick={() => onJump?.(k)}
              sx={{ cursor: 'pointer' }}
            >
              {errors[k] || k}
            </Link>
          ))}
        </Stack>
      </Stack>
    </Alert>
  );
}
