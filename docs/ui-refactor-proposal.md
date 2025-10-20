# 后台 UI 重构方案（V1）

> 目标：在不重写业务逻辑的前提下，显著提升后台可用性与观感，一致化交互与视觉，降低样式维护成本，并给出可落地的迁移路径与里程碑。

## 一、现状审查（基于仓库 web/）

- 技术栈：React 18 + Vite + Redux + i18next；组件库采用 MUI v5 + Emotion，主题在 `web/src/themes` 内实现（重度自定义）。
- 主题结构：
  - `web/src/themes/index.js:1` 自定义主题入口，打包渐变/色板/排版/断点/组件覆盖；大量 `components` 级覆盖写入 `compStyleOverride.js`。
  - `web/src/themes/compStyleOverride.js:1` 单文件 30KB+ 覆盖了 Button/Paper/Card/DataGrid/Dialog 等；存在“粒度过细 + 视觉较重”的问题（阴影/渐变/玻璃拟态叠加）。
  - `web/src/themes/palette.js:1` 色板从 SCSS 变量注入，JS 与 SCSS 双源维护，存在“令牌分散”的隐患。
- 布局：
  - `web/src/layout/MainLayout/index.jsx:1` 使用 MUI AppBar/Drawer/自定义 Sidebar；断点在 `themes/index.js` 中自定义为 md=768、lg=1024、xl=1200，与 MUI 默认不一致，可能导致三方示例/文档错位。
  - `web/src/layout/MainLayout/Header/index.jsx:1` Header 逻辑与 `MinimalLayout` 存在重复，导航结构不完全统一（参见 `web/src/test-navigation.md` 已有修复记录）。
- 组件一致性：
  - Table、Form 的信息密度与操作入口不完全一致；DataGrid 覆盖较多，分页/筛选面板的视觉反馈在深浅色下差异较大。
  - `react-perfect-scrollbar`/`react-device-detect` 等增强库引入了额外样式与兼容性成本，很多场景可用原生 CSS 或 MUI 组件能力替代。
- 性能与维护：
  - Emotion 运行时样式 + 大量 styleOverrides，样式计算成本较高；主题覆盖集中在一个大文件，影响可读性与演进。

## 二、设计目标（Design Goals）

- 一致：统一间距体系、圆角、阴影强度、色彩语义与状态反馈；深浅色风格对齐。
- 轻量：弱化大面积渐变与玻璃拟态，向「克制的企业级审美」靠拢，突出信息优先。
- 可维护：令牌化（Design Tokens）统一来源，减少零散覆盖；组件抽象分层清晰。
- 可扩展：表格/表单/筛选/批量操作等后台高频模式提供规范化方案。

## 三、重构方案对比

1) 保留 MUI，系统化治理（低风险/中收益）
   - 思路：保留 MUI + Emotion，清理 `compStyleOverride.js`，引入「设计令牌」与 CSS 变量；将覆盖拆分为模块（Button/Card/Table/Dialog…）。
   - 优点：迁移成本最低，现有代码改动小；生态成熟；DataGrid 可继续沿用。
   - 风险：MUI 视觉风格与国内 B 端审美存在差异，需要在令牌与组件封装层做更多“风格化”工作。

2) 迁移 Ant Design 5 + ProComponents（中风险/高收益）
   - 思路：采用 Antd 5 Token System，使用 `@ant-design/pro-components` 提供的 ProTable/ProForm/ProLayout；替换核心基础组件（Button/Input/Form/Table/Layout）。
   - 优点：企业后台审美更贴合；表单/表格能力强，上手快；Token 模型清晰；中文社区完善。
   - 风险：组件替换面广（布局、表单、表格、抽屉/弹窗）；需要适配路由与权限；迁移时长更长。

3) Tailwind CSS + Radix UI（shadcn/ui）（中高风险/高可塑）
   - 思路：用 Tailwind 做布局与间距，Radix 作为无样式可访问性原件，借鉴 shadcn/ui 的组件实现与主题方案。
   - 优点：现代、极简、定制自由度高；原子化样式可控；暗色/品牌皮肤切换容易。
   - 风险：需自建大量组件抽象；表单/表格等后台高频组件需要额外封装；迁移量最大。

结论建议：
- 若「快速起效 + 风险低」优先：选方案 1（MUI 治理）。
- 若「统一企业风格 + 海量表单表格」优先：选方案 2（Antd 5）。
- 若追求「极简现代 + 品牌个性」：选方案 3（Tailwind + Radix）。

## 四、推荐落地路径（结合现状）

优先建议：方案 1 作为第一阶段，2~3 周见效；并在主题层预埋 Token，保留未来切到 Antd 的可能。

里程碑：
- M1（第 1 周）：
  - 建立设计令牌源（colors/spacing/radius/shadow/typography/border），统一 JS 源；通过 CSS 变量注入 `:root`（兼顾非 MUI 区域）。
  - 重构主题：拆分 `compStyleOverride.js` 为 `themes/overrides/*`，按组件模块化。
  - 布局统一：合并 Header 重复逻辑，规范 Sidebar 动效与断点；逐步移除 `react-perfect-scrollbar`。
- M2（第 2 周）：
  - 规范后台四大模式：
    1) 列表页（FilterBar + Pro Table 模式化）
    2) 表单页（区分创建/编辑，分组卡片 + 悬浮提交条）
    3) 详情页（分区信息 + 固定操作区）
    4) 看板（卡片指标 + 图表配色）
  - 深浅色统一（对比度/透明度/阴影强度感知一致）。
- M3（第 3 周，可选）：
  - DataGrid 统一样式与交互（空状态、Loading、筛选面板、列设置、行操作区）。
  - 通知/消息/抽屉/弹窗统一规范。

## 五、设计令牌（Draft）

示例（JS）：
```js
// design/tokens.js（唯一真源）
export const tokens = {
  color: {
    brand: { primary: '#0EA5FF', secondary: '#8B5CF6' },
    text: { primary: '#1F2937', secondary: '#6B7280', inverse: '#FFFFFF' },
    bg: { page: '#F7F9FC', card: '#FFFFFF', subtle: '#F3F4F6' },
    border: { default: '#E5E7EB', strong: '#D1D5DB' },
    success: '#10B981', warning: '#F59E0B', error: '#EF4444', info: '#3B82F6'
  },
  radius: { xs: 6, sm: 10, md: 14, lg: 18 },
  shadow: {
    sm: '0 1px 2px rgba(0,0,0,0.06)',
    md: '0 4px 10px rgba(0,0,0,0.10)',
    lg: '0 10px 24px rgba(0,0,0,0.12)'
  },
  spacing: { xs: 4, sm: 8, md: 12, lg: 16, xl: 24, xxl: 32 },
  typography: { fontFamily: 'Inter, Segoe UI, Roboto, system-ui', size: { body: 14, h1: 32, h2: 24, h3: 20 } }
};
```

注入 CSS 变量（供全局与非 MUI 部分复用）：
```js
// design/applyTokens.ts
import { tokens } from './tokens';
export const applyTokens = () => {
  const r = document.documentElement.style;
  r.setProperty('--c-text', tokens.color.text.primary);
  r.setProperty('--c-border', tokens.color.border.default);
  r.setProperty('--radius-md', `${tokens.radius.md}px`);
  r.setProperty('--shadow-md', tokens.shadow.md);
  // ...其余变量
};
```

在 MUI 主题中消费令牌（示意）：
```js
import { createTheme } from '@mui/material/styles';
import { tokens } from '../design/tokens';

export const theme = createTheme({
  shape: { borderRadius: tokens.radius.md },
  palette: {
    mode: 'light',
    primary: { main: tokens.color.brand.primary },
    background: { default: tokens.color.bg.page }
  },
  typography: {
    fontFamily: tokens.typography.fontFamily,
    fontSize: tokens.typography.size.body
  }
  // components 覆盖尽量聚焦变量化，不做重装饰
});
```

## 六、组件分层与规范

- 层级：
  - 基础层（Base）：对 MUI 进行「薄封装」并默认接入 tokens（如 `BaseButton`、`BaseCard`）。
  - 复合层（Compound）：后台模式化组件（FilterBar、SearchForm、ActionBar、ConfirmModal）。
  - 业务层（Biz）：面向具体页面的组合（如「渠道管理-查询头 + 列表 + 批量操作」）。
- 规范要点：
  - 按照「内容 > 组件 > 装饰」原则弱化炫彩渐变；色彩与状态严格来源于 tokens。
  - `sx` 使用统一：只允许使用 spacing/radius/color 令牌，禁止直接写像素值与随机色。
  - 空状态、加载、错误提示提供标准化反馈与插画占位。

## 七、关键页面模式设计

- 列表页（高频）
  - 顶部：`FilterBar`（折叠/展开、重置、关键筛选项显性化）。
  - 中部：`DataGrid`（列设置、密度、行操作浮层统一）。
  - 底部：操作区与分页对齐，批量操作显性化。
- 表单页
  - 左内容右侧操作（吸顶保存/返回/校验），分组卡片化，错误锚点定位。
- 详情页
  - 信息区块化 + 时间线/状态流；操作始终可达（固定在可视区）。

## 八、性能与可访问性

- 减少运行时样式生成：优先 `sx` + tokens；减少 `styled` 与大规模 styleOverrides。
- 拆分覆盖文件，利于按需加载与代码审查；降低重排与重绘的阴影/模糊强度。
- 无障碍：对比度 AA、键盘可达、焦点可见、图标按钮需文本提示。

## 九、实施计划（甘特粒度）

- W1：引入 tokens 与 CSS 变量；重构主题骨架；布局合并；建 UI 规范样例页（目录：`/styleguide`）。
- W2：列表页/表单页模式化改造；DataGrid 统一；通知/弹窗规范化。
- W3：全站组件替换收口；暗色对齐；补齐可访问性与 E2E 验收用例。

## 十、迁移到 Antd 的预案（可选并行 PoC）

- 新建分支 `feat/antd-poc`：
  - 引入 `antd@^5` 与 `@ant-design/pro-components`；根据 tokens 生成 Antd `ConfigProvider` 主题；实现 1 个典型列表页与 1 个表单页 PoC。
  - 对比开发效率、包体体积、UI 一致性与用户反馈，再决定是否切换主线。

---

附：涉及文件参考

- web/src/themes/index.js:1 — 主题入口（需要拆分/令牌化）
- web/src/themes/compStyleOverride.js:1 — 组件覆盖（建议按模块拆分）
- web/src/layout/MainLayout/index.jsx:1 — 主布局（断点与栅格/交互规范）
- web/src/layout/MainLayout/Header/index.jsx:1 — 顶部导航（合并与精简）
- web/src/views/* — 业务页面（套用统一页面模板）

