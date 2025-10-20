// ApexCharts 统一主题与网格、标签、Tooltip 风格（弱化网格，文本跟随主题）
import { useTheme } from '@mui/material/styles';

export function useApexTheme() {
  const theme = useTheme();
  return {
    chart: { toolbar: { show: false }, foreColor: theme.palette.text.secondary },
    grid: { show: true, borderColor: theme.palette.divider, strokeDashArray: 3 },
    dataLabels: { enabled: false },
    stroke: { width: 2, curve: 'smooth' },
    xaxis: { labels: { style: { colors: theme.palette.text.secondary } }, axisBorder: { show: false }, axisTicks: { show: false } },
    yaxis: { labels: { style: { colors: theme.palette.text.secondary } } },
    legend: { position: 'top', horizontalAlign: 'left', labels: { colors: theme.palette.text.secondary } },
    tooltip: { theme: theme.palette.mode === 'dark' ? 'dark' : 'light' },
    colors: [theme.palette.primary.main, theme.palette.secondary.main, theme.palette.success.main]
  };
}
