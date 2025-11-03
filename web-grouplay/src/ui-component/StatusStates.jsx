import { Stack, Typography, Button } from '@mui/material';
import { Icon } from '@iconify/react';

export function LoadingState({ title = '加载中', description = '请稍候…' }) {
  return (
    <Stack alignItems="center" spacing={1.5} sx={{ py: 6 }}>
      <Icon icon="solar:refresh-linear" width={40} height={40} />
      <Typography variant="subtitle1">{title}</Typography>
      <Typography variant="body2" color="text.secondary">{description}</Typography>
    </Stack>
  );
}

export function NoPermissionState({ title = '无权限', description = '请联系管理员授予访问权限' }) {
  return (
    <Stack alignItems="center" spacing={1.5} sx={{ py: 6 }}>
      <Icon icon="solar:lock-keyhole-linear" width={40} height={40} />
      <Typography variant="subtitle1">{title}</Typography>
      <Typography variant="body2" color="text.secondary">{description}</Typography>
    </Stack>
  );
}

export function ErrorState({ title = '出错了', description = '请稍后重试或联系支持' , action }) {
  return (
    <Stack alignItems="center" spacing={1.5} sx={{ py: 6 }}>
      <Icon icon="solar:danger-triangle-linear" width={40} height={40} />
      <Typography variant="subtitle1" color="error">{title}</Typography>
      <Typography variant="body2" color="text.secondary">{description}</Typography>
      {action}
    </Stack>
  );
}
