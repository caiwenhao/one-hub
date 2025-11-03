import { Box, Card, CardContent, CardHeader, Grid, Stack, Button, TextField, Typography } from '@mui/material';
import { DataGrid } from '@mui/x-data-grid';
import GridWithPrefs from 'ui-component/grid/GridWithPrefs';

const rows = [
  { id: 1, name: '一段很长很长很长的名称示例，用于测试省略号效果', status: '启用' },
  { id: 2, name: '较短名称', status: '禁用' }
];
const cols = [
  { field: 'id', headerName: 'ID', type: 'number', headerAlign: 'right', align: 'right', flex: 1 },
  {
    field: 'name',
    headerName: '名称',
    flex: 2,
    renderCell: (params) => (
      <span className="table-text-ellipsis" title={params.value} style={{ display: 'inline-block', maxWidth: '100%' }}>
        {params.value}
      </span>
    )
  },
  { field: 'status', headerName: '状态', flex: 1 }
];

export default function BeforeAfter() {
  return (
    <Box sx={{ mt: 2 }}>
      <Typography variant="h3" sx={{ mb: 1.5 }}>Before / After（密度对照）</Typography>
      <Grid container spacing={2}>
        <Grid item xs={12} md={6}>
          <Card data-density="standard">
            <CardHeader title="标准 Standard" subheader="默认密度" />
            <CardContent>
              <Stack direction="row" spacing={1} sx={{ mb: 1 }}>
                <Button variant="contained">保存</Button>
                <Button variant="outlined">取消</Button>
              </Stack>
              <Stack spacing={1.5} sx={{ mb: 1 }}>
                <TextField label="名称" placeholder="请输入" fullWidth />
                <TextField label="描述" placeholder="可选" fullWidth />
              </Stack>
              <Box sx={{ height: 240 }}>
                <DataGrid rows={rows} columns={cols} density="standard" hideFooter />
              </Box>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={6}>
          <Card data-density="compact">
            <CardHeader title="紧凑 Compact" subheader="更高信息密度" />
            <CardContent>
              <Stack direction="row" spacing={1} sx={{ mb: 1 }}>
                <Button variant="contained">保存</Button>
                <Button variant="outlined">取消</Button>
              </Stack>
              <Stack spacing={1.5} sx={{ mb: 1 }}>
                <TextField label="名称" placeholder="请输入" fullWidth />
                <TextField label="描述" placeholder="可选" fullWidth />
              </Stack>
              <Box sx={{ height: 240 }}>
                <GridWithPrefs gridId="styleguide-compact" rows={rows} columns={cols} density="compact" hideFooter />
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Box>
  );
}
