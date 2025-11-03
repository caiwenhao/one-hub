// 将设计令牌注入为 CSS 变量，便于非 MUI 区域也能复用
import { tokens } from './tokens';

export const applyTokens = () => {
  const r = document.documentElement.style;
  // 颜色
  r.setProperty('--c-brand', tokens.color.brand.primary);
  r.setProperty('--c-text', tokens.color.text.primary);
  r.setProperty('--c-text-2', tokens.color.text.secondary);
  r.setProperty('--c-bg', tokens.color.bg.page);
  r.setProperty('--c-card', tokens.color.bg.card);
  r.setProperty('--c-border', tokens.color.border.default);
  // 圆角 & 阴影
  r.setProperty('--radius-sm', `${tokens.radius.sm}px`);
  r.setProperty('--radius-md', `${tokens.radius.md}px`);
  r.setProperty('--shadow-sm', tokens.shadow.sm);
  r.setProperty('--shadow-md', tokens.shadow.md);
};

export default applyTokens;

