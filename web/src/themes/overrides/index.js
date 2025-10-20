import componentStyleOverrides from '../compStyleOverride';
import button from './button';
import paperCard from './paperCard';
import dialog from './dialog';
import table from './table';
import datagrid from './datagrid';

export default function overrides(theme) {
  // 基于原有覆盖作为“全集”，再用模块化覆盖关键组件，逐步迁移剩余项
  const base = componentStyleOverrides(theme) || {};
  return {
    ...base,
    ...button(theme),
    ...paperCard(theme),
    ...dialog(theme),
    ...table(theme),
    ...datagrid(theme)
  };
}

