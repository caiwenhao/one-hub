import { Skeleton, Stack } from '@mui/material';

export default function SkeletonList({ rows = 5 }) {
  return (
    <Stack spacing={1.5} sx={{ py: 1 }}>
      {Array.from({ length: rows }).map((_, i) => (
        <Stack key={i} direction="row" spacing={2} alignItems="center">
          <Skeleton variant="circular" width={28} height={28} />
          <Skeleton variant="text" width="30%" height={18} />
          <Skeleton variant="text" width="15%" height={18} />
          <Skeleton variant="rectangular" width="40%" height={12} sx={{ borderRadius: 1 }} />
        </Stack>
      ))}
    </Stack>
  );
}
