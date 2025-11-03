import { Stack, Button } from '@mui/material';
import LoadingButton from '@mui/lab/LoadingButton';

export default function ButtonsDoc() {
  return (
    <Stack spacing={1.5}>
      <Stack direction="row" spacing={1}>
        <Button variant="contained">主按钮</Button>
        <Button variant="outlined">次按钮</Button>
        <Button className="btn-ghost">幽灵按钮</Button>
        <Button color="error" variant="contained">危险</Button>
      </Stack>
      <Stack direction="row" spacing={1}>
        <LoadingButton loading variant="contained">加载中</LoadingButton>
        <LoadingButton loading variant="outlined">加载中</LoadingButton>
      </Stack>
    </Stack>
  );
}
