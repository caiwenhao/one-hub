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
  Chip,
  Badge
} from '@mui/material';
import { DataGrid } from '@mui/x-data-grid';
import Chart from 'react-apexcharts';
import { useApexTheme } from 'ui-component/charts/apexTheme';
import EmptyState from 'ui-component/EmptyState';
import UploadList from 'ui-component/upload/UploadList';
import BeforeAfter from './BeforeAfter';
import ApexChartDemo from './ApexChartDemo';
import { Icon } from '@iconify/react';
import ButtonsDoc from './ButtonsDoc';
import I18nDemo from './I18nDemo';
import FilterBar from 'ui-component/FilterBar';
import FilterChips from 'ui-component/FilterChips';
import SearchInput from 'ui-component/SearchInput';

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

      <Grid container spacing={2} className="motion-fade-in">
        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="颜色" subheader="品牌与中性色阶" />
            <CardContent>
              <Stack direction="row" spacing={1} flexWrap="wrap">
                {['--color-brand-primary','--color-brand-secondary','--color-semantic-success','--color-semantic-warning','--color-semantic-error','--color-semantic-info'].map((v) => (
                  <Box key={v} sx={{ width: 80, height: 48, borderRadius: 1, boxShadow: 1, background: `var(${v})` }} title={v} />
                ))}
              </Stack>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="动效" subheader="统一进入/离开/状态切换" />
            <CardContent>
              <Stack direction="row" spacing={2}>
                <Box className="motion-fade-in" sx={{ width: 80, height: 48, borderRadius: 1, bgcolor: 'primary.main' }} />
                <Box className="motion-pop-in" sx={{ width: 80, height: 48, borderRadius: 1, bgcolor: 'secondary.main' }} />
              </Stack>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="多语言长度回退" subheader="换行/省略+Tooltip、单位空格、日期/数字本地化" />
            <CardContent>
              <I18nDemo />
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="按钮" subheader="语义色/尺寸/幽灵/危险/加载态" />
            <CardContent>
              <ButtonsDoc />
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

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="筛选与搜索" subheader="标签化与可清除、搜索防抖" />
            <CardContent>
              <FilterBar>
                <Stack spacing={1.5}>
                  <SearchInput placeholder="搜索关键词（支持清除与防抖）" onChange={() => {}} />
                  <FilterChips items={[{ key: '状态', value: '启用' }, { key: '类型', value: 'A 类' }]} onClear={() => {}} onRemove={() => {}} />
                </Stack>
              </FilterBar>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12}>
          <Card>
            <CardHeader title="图表（ApexCharts）" subheader="网格弱化、标签与 Tooltip 风格一致" />
            <CardContent>
              <ApexChartDemo />
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12}>
          <Card>
            <CardHeader title="表格（DataGrid）" subheader="表头/行高/悬停/选中态 + 粘顶工具栏" />
            <CardContent>
              <Box sx={{ height: 280 }}>
                <DataGrid rows={rows} columns={cols} disableRowSelectionOnClick hideFooter density="compact" slots={{ toolbar: () => <Box sx={{ p: 1 }}>工具栏（示例）</Box> }} />
              </Box>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="状态与警示" subheader="信息/成功/警告/错误 + 四件套" />
            <CardContent>
              <Stack spacing={1.5}>
                <Alert severity="info">这是一条信息提示</Alert>
                <Alert severity="success">操作成功的提示</Alert>
                <Alert severity="warning">需要注意的警告</Alert>
                <Alert severity="error">错误或失败提示</Alert>
              </Stack>
              <Divider sx={{ my: 2 }} />
              <Stack spacing={2}>
                {/* 四件套：加载/无权限/错误/空状态 */}
                <Alert severity="info">四件套（示例）：</Alert>
                {/* 直接复用已有 EmptyState；其它见 StatusStates 组件 */}
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
            <CardHeader title="空状态与上传" subheader="统一图标/文案/动作 + 进度行样式" />
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
                <Divider />
                {/* 上传进度演示 */}
                <Typography variant="subtitle2">上传组件</Typography>
                <UploadList files={[{ id: 1, file: { name: 'report.pdf' }, progress: 62, status: 'uploading' }]} />
              </Stack>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card>
            <CardHeader title="表单与徽标" subheader="聚焦/错误/禁用 + Badge 样式" />
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
                <Stack direction="row" spacing={2} alignItems="center">
                  <Chip label="标签" color="primary" variant="outlined" />
                  <Chip label="禁用" disabled />
                  <Button variant="contained" className="btn-ghost">幽灵按钮</Button>
                </Stack>
                <Stack direction="row" spacing={2} alignItems="center">
                  <Badge color="primary" badgeContent={9}>
                    <Box sx={{ width: 36, height: 36, borderRadius: '50%', bgcolor: 'primary.main' }} />
                  </Badge>
                  <Badge color="error" badgeContent={120} max={99}>
                    <Box sx={{ width: 36, height: 36, borderRadius: '50%', bgcolor: 'secondary.main' }} />
                  </Badge>
                </Stack>
              </Stack>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      {/* Before / After 区块：用于截图基线 */}
      <BeforeAfter />
    </Box>
  );
}
