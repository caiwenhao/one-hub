# 后台 UI 改造清单（企业稳重风格）

> 注：本清单随进度持续更新；已完成项标注“已完成”，其余为“未完成”。

## 一、基础与主题
- 已完成：统一设计令牌源与注入
  - `web/src/design/tokens.js:1`
  - `web/src/design/applyTokens.js:1`
  - `web/src/index.jsx:1` 注入 `applyTokens()`
- 已完成：主题基线（稳重风格）
  - 断点回归 MUI 默认、圆角收敛、渐变弱化
  - `web/src/themes/index.js:1`
- 已完成：组件覆盖第一批精简（稳重风格）
  - Button/Paper/Card/Dialog/DataGrid/Table
  - `web/src/themes/compStyleOverride.js:60`, `:125`, `:148`, `:764`, `:874`, `:488` 等
- 已完成：修复 Palette 未定义导致的崩溃
  - `web/src/themes/compStyleOverride.js:133`
- 已完成：以 tokens 生成 colors，替代 SCSS 作为 JS 主题来源
  - `web/src/themes/buildColorsFromTokens.js:1`，`web/src/themes/index.js:1`
- 已完成：在 `palette.js`、`typography.js` 中通过 `theme.colors` 间接消费 tokens
  - `web/src/themes/palette.js:1`, `web/src/themes/typography.js:1`
- 已完成：拆分 overrides 为模块化目录（保留原文件为基底，模块化覆盖关键组件）
  - `web/src/themes/overrides/*`, `web/src/themes/index.js:1`

## 二、规范与风格基线
- 已完成：UI 规范页与入口
  - 视图：`web/src/views/Styleguide/index.jsx:1`
  - 路由：`web/src/routes/MainRoutes.jsx:1`（`/panel/styleguide`）
  - 菜单：`web/src/menu-items/dashboard.jsx:1`（UI规范）
- 已完成：在规范页补充空状态与表单片段样例
  - `web/src/views/Styleguide/index.jsx:1`
- 已完成：补充“错误/警示/密度”更多示例
  - `web/src/views/Styleguide/index.jsx:1`

## 三、通用复用组件
- 已完成：筛选区容器 `FilterBar`、操作区容器 `ActionBar`
  - `web/src/ui-component/FilterBar.jsx:1`
  - `web/src/ui-component/ActionBar.jsx:1`
- 已完成：统一滚动容器 `ScrollArea`（替代 PerfectScrollbar）
  - `web/src/ui-component/ScrollArea.jsx:1`
- 进行中：标准化 `ConfirmModal`、`Skeleton`、表单校验提示组件（`EmptyState` 已完成）
  - `web/src/ui-component/EmptyState.jsx:1`

## 四、页面改造（列表页模式：FilterBar + ActionBar + 表格统一）
- 已完成：渠道管理 `Channel`
  - `web/src/views/Channel/index.jsx:1`
- 已完成：用户管理 `User`
  - `web/src/views/User/index.jsx:1`
  
- 已完成：Token `web/src/views/Token/index.jsx:1`
- 已完成：Payment（网关/订单） `web/src/views/Payment/Gateway.jsx:1`, `web/src/views/Payment/Order.jsx:1`
- 已完成：Log `web/src/views/Log/index.jsx:1`
- 已完成：Telegram `web/src/views/Telegram/index.jsx:1`
- 已完成：Pricing `web/src/views/Pricing/index.jsx:1`（操作区统一）
- 已完成：ModelPrice `web/src/views/ModelPrice/index.jsx:1`（筛选/操作区统一）
- 已完成：Task `web/src/views/Task/index.jsx:1`
- 已完成：UserGroup `web/src/views/UserGroup/index.jsx:1`
- 已完成：ModelOwnedby `web/src/views/ModelOwnedby/index.jsx:1`
- 已完成：Invoice `web/src/views/Invoice/index.jsx:1`
- 已完成：Midjourney `web/src/views/Midjourney/index.jsx:1`
- 已完成：Pricing/multiple `web/src/views/Pricing/multiple.jsx:1`
- 说明：`Topup/Analytics/SystemInfo/Playground` 属信息/表单页，无检索筛选需求，已统一滚动/表格样式，无需接入 FilterBar

## 五、表单与详情页规范
- 进行中：吸顶操作区（保存/取消）、卡片分组、错误锚点
  - 已完成示例：`Profile` 吸底操作区（StickyActions） web/src/views/Profile/index.jsx:1
  - 已完成示例：`Profile` 表单错误锚点（滚动聚焦首个错误字段）
  - 待办：`Setting`、渠道/用户编辑相关 `Modal/页`
- 未完成：详情页信息分区 + 固定操作区
  - 优先：`Invoice/detail`、`Profile`

## 六、布局与导航
- 进行中：统一 `MainLayout` 与 `MinimalLayout` Header 重复逻辑
  - 已完成：右侧操作区抽取 `HeaderActions` 并接入两处
    - `web/src/layout/common/HeaderActions.jsx:1`
    - `web/src/layout/MainLayout/Header/index.jsx:1`
    - `web/src/layout/MinimalLayout/Header/index.jsx:1`
  - 已完成：Minimal 导航项样式从渐变改为纯色下划线，配色对齐 tokens.brand
  - 已完成：导航配置化（navConfig + NavLinkItem）并接入 MinimalLayout Header
    - `web/src/layout/MinimalLayout/Header/navConfig.js:1`
    - `web/src/layout/MinimalLayout/Header/NavLinkItem.jsx:1`
- 已完成：移除 `react-perfect-scrollbar`，改用原生/MUI 滚动
  - 已替换所有页面使用，保留为注释；删除样式引用 `web/src/assets/scss/style.scss:5`
- 未完成：Sidebar 动效/密度与分组标签规范（已切换为原生滚动）
- 未完成：断点调整后的全站响应式复核（`md=900` 等）

## 七、性能与可访问性
- 进行中：继续削减运行时样式与过度覆盖，优先 `sx` + tokens
- 未完成：可访问性（焦点可见、键盘可达、对比度 AA、图标按钮辅助文本）
- 未完成：表格/筛选区在移动端的交互可用性优化

## 八、工程化与文档
- 已完成：重构方案文档
  - `docs/ui-refactor-proposal.md:1`
- 已完成：UI 规范与编码规约（`sx` 使用、间距/色彩令牌约束）
  - `docs/ui-coding-guidelines.md:1`
- 未完成：验收清单与 E2E/手测用例（关键路径）
- 未完成：CI 检查样式与 ESLint 规则补充（避免直写像素/随机色）

## 九、可选预研（并行）
- 未完成：Ant Design 5 + ProComponents PoC（新分支 `feat/antd-poc`）
  - 以 tokens 生成主题，落地 1 个列表 + 1 个表单对照评估

---

建议推进顺序：
1) 批量接入列表页（Token、Payment、Log、Pricing、Telegram）
2) 表单与详情规范（Setting、Profile、渠道/用户编辑）
3) SCSS → tokens 收敛、overrides 模块化、移除 perfect-scrollbar
4) 可访问性与规范文档/E2E 收口
