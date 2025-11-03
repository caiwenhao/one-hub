import Chart from 'react-apexcharts';
import { useMemo } from 'react';
import { useApexTheme } from 'ui-component/charts/apexTheme';

export default function ApexChartDemo() {
  const apexTheme = useApexTheme();
  const series = useMemo(() => [{ name: '访问量', data: [10, 22, 18, 30, 26, 35] }], []);
  const options = useMemo(() => ({
    ...apexTheme,
    chart: { ...apexTheme.chart, type: 'area', height: 260 },
    xaxis: { ...apexTheme.xaxis, categories: ['1月', '2月', '3月', '4月', '5月', '6月'] }
  }), [apexTheme]);
  return <Chart series={series} options={options} height={260} />;
}
