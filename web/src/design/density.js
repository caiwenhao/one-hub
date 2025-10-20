// 密度系统：紧凑/标准/舒适
const KEY = 'ui-density';
export function applyDensity(next) {
  const root = document.documentElement;
  const val = next || localStorage.getItem(KEY) || 'standard';
  if (val === 'compact' || val === 'comfortable' || val === 'standard') {
    if (val === 'standard') root.removeAttribute('data-density');
    else root.setAttribute('data-density', val);
    localStorage.setItem(KEY, val);
  }
}
export function getDensity() {
  return document.documentElement.getAttribute('data-density') || 'standard';
}
export function toggleDensity() {
  const order = ['compact', 'standard', 'comfortable'];
  const cur = getDensity();
  const idx = order.indexOf(cur);
  const next = order[(idx + 1) % order.length];
  applyDensity(next);
}
// 可选：绑定到 window 便于调试
if (typeof window !== 'undefined') {
  window.__density__ = { applyDensity, getDensity, toggleDensity };
}
