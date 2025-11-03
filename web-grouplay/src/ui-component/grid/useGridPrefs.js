// DataGrid 偏好持久化（列可见性/顺序/密度）
// 说明：社区版不支持列冻结；此 Hook 持久化可见性与顺序、列宽在 onStateChange 中尽量获取
const KEY = (id) => `grid:prefs:${id}`;

export function loadGridPrefs(gridId) {
  try {
    const raw = localStorage.getItem(KEY(gridId));
    if (!raw) return {};
    const data = JSON.parse(raw);
    return data || {};
  } catch {
    return {};
  }
}

export function saveGridPrefs(gridId, state) {
  try {
    const payload = {
      density: state?.density || undefined,
      columns: {
        columnVisibilityModel: state?.columns?.columnVisibilityModel || undefined,
        orderedFields: state?.columns?.orderedFields || undefined,
        widths: state?.columns?.columnLookup
          ? Object.fromEntries(
              Object.entries(state.columns.columnLookup).map(([f, v]) => [f, v?.computedWidth || v?.width]).filter(([, w]) => !!w)
            )
          : undefined
      }
    };
    localStorage.setItem(KEY(gridId), JSON.stringify(payload));
  } catch {
    // ignore
  }
}

export function getInitialStateFromPrefs(gridId) {
  const prefs = loadGridPrefs(gridId);
  const initialState = {};
  if (prefs?.columns?.columnVisibilityModel) {
    initialState.columns = initialState.columns || {};
    initialState.columns.columnVisibilityModel = prefs.columns.columnVisibilityModel;
  }
  if (prefs?.columns?.orderedFields) {
    initialState.columns = initialState.columns || {};
    initialState.columns.orderedFields = prefs.columns.orderedFields;
  }
  return { initialState, prefs };
}
