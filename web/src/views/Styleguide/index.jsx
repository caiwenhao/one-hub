import { useTheme } from '@mui/material/styles';
import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  Divider,
  Grid,
  Stack,
  Typography,
  Alert,
  TextField,
  Chip
} from '@mui/material';
import { DataGrid } from '@mui/x-data-grid';
import EmptyState from 'ui-component/EmptyState';
import { Icon } from '@iconify/react';

// 简易 UI 规范页（企业稳重风格预览）
const rows = [
  { id: 1, name: '示例A', status: '启用' },
  { id: 2, name: '示例B', status: '禁用' }
];
const cols = [
  { field: 'id', headerName: 'ID', flex: 1 },
  { field: 'name', headerName: '名称', flex: 2 },
  { field: 'status', headerName: '状态', flex: 1 }
];

export default function Styleguide() {
  const theme = useTheme();

  return (
    <Box sx={{ p: { xs: 1.5, md: 3 } }}>
      <Typography variant="h2" sx={{ mb: 2 }}>UI 规范 · 企业稳重风格</Typography>
      <Typography variant="body2" color="text.secondary" sx={{ mb: 3 }}>
        本页用于预览与校验设计令牌、组件样式与密度，作为重构的视觉基线。
      </Typography>

      <Grid container spacing={2}>
        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="按钮" subheader="语义色与尺寸" />
            <CardContent>
              <Stack direction="row" spacing={1} sx={{ mb: 2 }}>
                <Button variant="contained">主按钮</Button>
                <Button variant="outlined">次按钮</Button>
                <Button variant="text">文本</Button>
              </Stack>
              <Stack direction="row" spacing={1}>
                <Button size="small" variant="contained">小</Button>
                <Button variant="contained">中</Button>
                <Button size="large" variant="contained">大</Button>
              </Stack>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="排版" subheader="层级与密度" />
            <CardContent>
              <Typography variant="h1">标题 H1</Typography>
              <Typography variant="h2">标题 H2</Typography>
              <Typography variant="h3" sx={{ mb: 1 }}>标题 H3</Typography>
              <Divider sx={{ my: 1.5 }} />
              <Typography variant="body1">正文（14px/默认）</Typography>
              <Typography variant="body2" color="text.secondary">辅助文本与说明</Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12}>
          <Card>
            <CardHeader title="表格（DataGrid）" subheader="表头/行高/悬停/选中态" />
            <CardContent>
              <Box sx={{ height: 280 }}>
                <DataGrid rows={rows} columns={cols} disableRowSelectionOnClick hideFooter density="compact" />
              </Box>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="状态与警示" subheader="信息/成功/警告/错误" />
            <CardContent>
              <Stack spacing={1.5}>
                <Alert severity="info">这是一条信息提示</Alert>
                <Alert severity="success">操作成功的提示</Alert>
                <Alert severity="warning">需要注意的警告</Alert>
                <Alert severity="error">错误或失败提示</Alert>
              </Stack>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="密度与对齐" subheader="紧凑型表格与按钮间距" />
            <CardContent>
              <Stack direction="row" spacing={1} sx={{ mb: 1 }}>
                <Button size="small" variant="outlined">小</Button>
                <Button variant="outlined">中</Button>
                <Button size="large" variant="outlined">大</Button>
              </Stack>
              <Box sx={{ height: 220 }}>
                <DataGrid rows={rows} columns={cols} density="compact" hideFooter />
              </Box>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="空状态" subheader="统一图标/文案/操作按钮" />
            <CardContent>
              <Stack spacing={2}>
                <EmptyState
                  title="暂无数据"
                  description="还没有符合条件的记录，尝试调整筛选条件或创建新条目。"
                  action={<Button variant="contained">新建记录</Button>}
                  icon={<Icon icon="solar:archive-down-linear" />}
                />
                <EmptyState
                  size="small"
                  title="搜索结果为空"
                  description="换一个关键词试试，或检查是否输入有误。"
                  icon={<Icon icon="solar:magnifer-linear" />}
                />
              </Stack>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="表单状态" subheader="聚焦/错误/禁用示例" />
            <CardContent>
              <Stack spacing={2}>
                <TextField label="默认" placeholder="输入内容" fullWidth />
                <TextField
                  label="错误字段"
                  placeholder="输入内容"
                  error
                  helperText="请输入合法的数值"
                  fullWidth
                />
                <TextField label="禁用字段" disabled fullWidth value="只读信息" />
                <Stack direction="row" spacing={1}>
                  <Chip label="标签" color="primary" variant="outlined" />
                  <Chip label="禁用" disabled />
                </Stack>
              </Stack>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Box>
  );
}
