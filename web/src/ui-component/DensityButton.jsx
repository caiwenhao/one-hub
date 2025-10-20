// 密度切换按钮（紧凑/标准/舒适）
// 中文：用于快速在控制台切换 UI 密度，以验证列表/表格/表单的紧凑效果
import { useEffect, useState } from 'react';
import { IconButton, Tooltip } from '@mui/material';
import DensitySmallIcon from '@mui/icons-material/DensitySmall';
import DensityMediumIcon from '@mui/icons-material/DensityMedium';
import DensityLargeIcon from '@mui/icons-material/DensityLarge';
import { applyDensity, getDensity, toggleDensity } from 'design/density';

const nextIcon = {
  compact: <DensitySmallIcon fontSize="small" />, // 下一步会变成 standard
  standard: <DensityMediumIcon fontSize="small" />, // 下一步会变成 comfortable
  comfortable: <DensityLargeIcon fontSize="small" /> // 下一步会变成 compact
};

export default function DensityButton() {
  const [density, setDensity] = useState('standard');
  useEffect(() => {
    setDensity(getDensity());
  }, []);
  const onClick = () => {
    toggleDensity();
    const cur = getDensity();
    applyDensity(cur);
    setDensity(cur);
  };
  const tipMap = {
    compact: '当前：紧凑（点击切换为 标准）',
    standard: '当前：标准（点击切换为 舒适）',
    comfortable: '当前：舒适（点击切换为 紧凑）'
  };
  return (
    <Tooltip title={tipMap[density] || '切换密度'}>
      <IconButton size="small" onClick={onClick} aria-label="切换密度">
        {nextIcon[density]}
      </IconButton>
    </Tooltip>
  );
}
