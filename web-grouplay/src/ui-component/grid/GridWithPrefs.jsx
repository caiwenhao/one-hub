import { useMemo } from 'react';
import { DataGrid } from '@mui/x-data-grid';
import { getInitialStateFromPrefs, saveGridPrefs } from './useGridPrefs';
import { getDensity } from 'design/density';

// 带本地偏好持久化的 DataGrid 包装
export default function GridWithPrefs({ gridId, initialState, density, onStateChange, ...rest }) {
  const { initialState: saved, prefs } = useMemo(() => getInitialStateFromPrefs(gridId), [gridId]);
  const mergedInitial = useMemo(() => ({ ...(initialState || {}), ...saved }), [initialState, saved]);
  const finalDensity = density || prefs?.density || getDensity() || 'standard';

  const handleStateChange = (state, event) => {
    saveGridPrefs(gridId, state);
    onStateChange?.(state, event);
  };

  return <DataGrid initialState={mergedInitial} density={finalDensity} onStateChange={handleStateChange} {...rest} />;
}
