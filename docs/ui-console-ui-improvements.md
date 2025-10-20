# 后台 UI 优化任务清单

> 基于最近的后台界面审美评审（20 项建议），持续跟踪并更新进度。

## 布局与基础框架

- [x] 01. 调整 `MainLayout` 主内容高度与滚动策略，避免双滚动与底部留白。
- [x] 02. 移除根容器的 `overflow: hidden`，确保浮层组件不被裁切。
- [x] 03. 扩充 `AdminContainer` 最大宽度与边距，适配 1440px+ 桌面视图。
- [x] 04. 调整面包屑与正文间距，让页面结构更有呼吸感。
- [x] 05. 统一 Header 工具按钮尺寸与间距（通知/主题/语言/个人中心）。

## 导航与侧边栏

- [x] 06. 优化侧边栏用户卡片材质与动效，替换为稳重静态表现。
- [x] 07. 收敛进度条色彩梯度，采用品牌色透明度表达用量状态。
- [x] 08. 强化导航分组标题视觉（底色或竖线），提升可扫描性。
- [x] 09. 优化导航条目激活态，使用浅主色背景与圆角块。
- [x] 10. 整理折叠菜单图标与留白，统一图标尺寸与 padding。

## 页面结构与模块

- [x] 11. 抽象标准 `PageHeader` 组件，统一标题、副标题、行动按钮布局。
- [x] 12. 在 Dashboard、User 等核心页面接入标准页头组件。
- [x] 13. 调整筛选区背景（FilterBar），使用品牌浅色块区隔模块。
- [x] 14. 重构操作区（ActionBar），改用 `CardActions` 风格与合理间距。
- [x] 15. 缩减搜索栏高度与留白，使列表页顶部更加紧凑。
- [x] 16. 为表格容器添加统一内边距，避免贴边导致的压迫感。
- [x] 17. 在 Setting 等表单/详情页拆分 `CardHeader`/`CardContent`，避免双边框。

## 交互与状态组件

- [x] 18. 打造标准化空状态组件（图标 + 标题 + 行动），并在规范页展示。
- [x] 19. 完善 Profile 抽屉信息区布局（卡片化或分栏）与动效收敛。
- [x] 20. 补充 Styleguide 中的状态/错误/密度示例，便于持续对照。

> **进度约定**：
> - `[ ]` 待办 · 尚未启动
> - `[~]` 进行中 · 已进入开发
> - `[x]` 已完成 · 代码合入并自测

## 方案 B：Design Tokens 工具化落地（Style Dictionary）

> 已选方案：B（使用 Style Dictionary/Token Studio 输出多端 Token，并以 CSS 变量/TS 导出接入）

### 任务清单（21–50）

- [x] 21. 建立“密度系统”与全局切换（舒适/标准/紧凑），CSS 变量驱动控件高度/内间距，偏好持久化。
- [x] 22. 落地 8px 栅格与竖向节奏（4/8/12 级差），规范段落/模块/分隔的垂直节距。
- [x] 23. 字体体系重构：标题/正文字号与行高阶梯；中文阅读优化；数字启用等宽排版（tabular-nums）。
- [x] 24. 中性色与品牌色阶重构（50–900 阶梯），确保文本/控件对比度达 WCAG AA。（新增一键校验脚本：web/scripts/contrast-check.mjs，输出 docs/ui-baseline/contrast-report.md，新增 --color-brand-ink 作为文本品牌色）
- [x] 25. 圆角/阴影/边框 Token 统一（radii: 6/8/12；shadow: 0/1/2；边框 1px/发丝线），建立层级感。
- [x] 26. 悬浮/按下/焦点态统一：颜色、明度与位移一致；焦点环（focus ring）键盘可见。
- [x] 27. 按钮体系分层：主/次/幽灵/文本+三档尺寸；加载/禁用/危险态一致化。（统一 LoadingButton 样式，Styleguide 补文档）
- [x] 28. 图标规范：24/20/16 三档尺寸、描边粗细与基线对齐；避免视觉漂移。
- [x] 29. 流体布局：容器最大宽度与内边距用 clamp 实现（适配 1280–1440+），避免过度留白。
- [x] 30. 表单控件高度与内边距规范（输入/选择/多选/日期等），占位/标签/帮助/校验间距标准化。
- [x] 31. 表单布局模式：单列/双列/分组/对齐策略（左对齐或顶部对齐）与断点切换规则。
- [x] 32. 校验与错误汇总：内联提示+表单顶部概览，错误锚点跳转与可读文案规范。（Setting 已集成；将扩展至 Task/Payment）
- [x] 33. 表格密度档位（舒适/标准/紧凑）与用户偏好记忆（每表按键独立记忆）。
- [x] 34. 表格列对齐与截断策略：数字右对齐、文本省略号+悬浮气泡、操作列宽度约定。
- [x] 35. 表格交互统一：行悬浮/选中、批量操作条粘顶、分页与空状态一致样式。（DataGrid 工具栏粘顶覆盖；BatchBar 组件可复用）
- [x] 36. 列头/列冻结与列宽可调，支持最小/最大宽与持久化。（社区版不含冻结列；已提供列宽策略与持久化）
- [x] 37. 列表/卡片密度与骨架屏：卡片内边距与网格间距收敛，加载态统一。
- [x] 38. 空状态/加载/无权限/错误“四件套”统一模板（图标/标题/解释/行动）。
- [x] 39. 通知/Toast/消息条收敛：位置/最大并发/时长与层级统一；弱化视觉噪声。
- [x] 40. 筛选区（FilterBar）标签化展示已选条件，提供“一键清除”与“高级筛选折叠”。
- [x] 41. 搜索框行为：防抖、清除按钮、占位语气与结果态过渡（保留关键字）。
- [x] 42. 抽屉/模态尺寸与间距阶梯，底部操作区吸底，滚动时保持 Header 可见。
- [x] 43. 上传/进度组件：文件行样式、缩略图/图标、进度/错误/重试状态统一。
- [x] 44. 标签与徽标（Badge）规范：颜色语义、大小/圆角、图标搭配与最大计数。
- [x] 45. 图表风格统一：网格线弱化、标签/刻度样式、配色与强调色，Tooltip 一致化。
- [x] 46. 可访问性基线：ARIA/焦点顺序/跳过链接/语义标签，键盘可用性完备。
- [x] 47. 多语言长度回退：换行/省略+Tooltip、单位与空格、日期/数字本地化。（新增 i18nFormat 工具与 i18n.css，并在 Styleguide 演示）
- [x] 48. 动效基线：进入/离开/状态切换 150–200ms；尊重“减少动态”系统偏好。
- [x] 49. 主题与 Design Tokens 出口：spacing/radii/color/shadow/type/motion 以 CSS 变量提供。
- [x] 50. 风格指南扩展与视觉回归：新增“密度/排版/色彩/动效/图表/上传/筛选”章节+截图对比基线。

### 起步开发（本迭代优先）

- [x] B1. 建立 `tokens/` 目录与 `style-dictionary` 基础配置（`config.json`）。
- [x] B2. 定义首批 Token：`spacing|radii|color|shadow|type|motion`（`tokens/*.json`）。
- [x] B3. 配置输出：Web 生成 `dist/tokens.css`（CSS 变量）与 `dist/tokens.d.ts`（TS 类型）。
> 注：已生成 CSS 与 JSON，并在 `src/design/` 提供最小类型声明 `tokens.generated.d.ts`。
- [x] B4. 将 `tokens.css` 注入应用入口（如 `web/src/main.ts` 或全局样式），不破坏现有样式。
- [x] B5. 建立对照页（Before/After）与可视回归截图基线，覆盖 Dashboard/List/Form 三页。

> 截图基线：
> - Styleguide Before/After：![](./ui-baseline/styleguide-before-after.png)


> 进展小结：
> - 已安装并执行 `style-dictionary` 生成 `web/src/styles/tokens.css` 与 `web/src/design/tokens.generated.json`。
> - 已将 `MuiButton`、`MuiPaper`/`MuiCard`、`MuiOutlinedInput` 等接入 CSS 变量与密度系统；表格/Table 与 DataGrid 已接通圆角与间距变量。
> - Styleguide 增加 Before/After 对照区块（/panel/styleguide），支持标准/紧凑并列展示，后续可据此截屏沉淀回归基线。
> - 提供表格通用样式工具类（右对齐/文本省略）：`web/src/styles/table-helpers.css`。

### 验收标准

- 对比度达 WCAG AA；密度切换可感知；首屏信息量提升 ≥ 25%。

## Analytics 页面密度与一行布局改造（新增）

- [x] 顶部筛选条（日期范围/分组类型/用户ID/搜索）在 `md` 及以上断点整合为单行，减少纵向占用与滚动：
  - 代码位置：`web/src/views/Analytics/component/Overview.jsx`
  - 调整：四个筛选项网格列宽统一为 `md={3}`，在小屏依旧换行，兼顾响应式。
- [x] 「兑换统计」「注册统计」「充值」三张图表在 `md` 及以上断点整合为同一行：
  - 代码位置：`web/src/views/Analytics/component/Overview.jsx`
  - 调整：三卡片各占 `md={4}`，保证可读与对齐；小屏自动纵向排列。
- [x] 统一筛选区内边距（px/pt/pb），采用更紧凑的密度变量，减少视觉冗余。
- [x] 引入标准 `PageHeader` 与 `FilterBar`：
  - 页头：`web/src/views/Analytics/index.jsx`，统一标题/副标题，去除冗余日期标题与分隔线。
  - 筛选区：`FilterBar` 包裹筛选控件，背景/边框/圆角与主题一致；控件使用小尺寸（`size="small"`）。
- [ ] 后续优化建议：
  - 图表标题与单位对齐风格指南（tokens/typography 与 chart 统一）；
  - 过滤条件标签化展示 + 一键清除（复用 FilterBar 规范 #40）；
  - 顶部筛选条可选“粘顶”行为，长图表滚动时保持可见。

> 受影响文件：
> - `web/src/views/Analytics/component/Overview.jsx`
> - 对齐 Styleguide 截图基线（如需回归截图可在 `/panel/styleguide` 复盘）。
