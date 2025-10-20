## UI 编码规约（企业稳重风格）

> 目标：统一视觉与交互风格，降低维护成本，提升可读性与一致性。

### 1. 设计令牌优先
- 颜色、间距、圆角、阴影、字号必须优先使用 tokens
  - 颜色：`tokens.color.*` 或 `theme.palette.*`，禁止硬编码随机色
  - 圆角：`theme.shape.borderRadius` 或 `tokens.radius.*`
  - 阴影：`tokens.shadow.*`，表格/卡片使用轻阴影
  - 间距：统一使用 `theme.spacing(n)` 或 tokens 的语义尺寸

### 2. sx 使用规范
- 仅在 sx 中书写简单布局与少量差异样式
- 避免在 sx 中大段定义复杂视觉（改为组件覆盖或复用组件）
- 尽量避免直接像素值，使用 `theme.spacing` 与 `theme.shape`

### 3. 组件覆盖策略
- 使用 `web/src/themes/overrides/*` 模块化覆盖常用组件
- 不再在页面中重复覆盖相同规则（如表格密度/边界/hover）

### 4. 列表页模板
- 标准结构：`FilterBar`（筛选） + `ActionBar`（工具按钮） + 表格 + 分页
- 移动端：筛选项可折叠 / 工具按钮换行

### 5. 表单与详情
- 吸顶操作区（保存/取消）+ 卡片分组 + 错误锚点
- 字段标签、占位、帮助信息与错误提示齐全

### 6. 可访问性（A11y）
- 焦点可见：交互元素需有 `:focus-visible` 样式（主题已内置 Button/Link）
- 键盘可达：Tab 顺序正确，弹窗/菜单支持 ESC/Enter
- 对比度：遵循 AA 标准，浅色下文本对比度≥4.5:1

### 7. 命名与结构
- 组件：`PascalCase`，函数/变量：`camelCase`，常量：`UPPER_SNAKE`
- 复合组件放入 `ui-component/`，业务组件放各自页面 `component/`

### 8. 禁止项
- 禁止使用内联 style 写大段视觉样式
- 禁止使用未在 tokens/主题中定义的随机色值
- 禁止在页面中重复实现已存在的通用组件

---

配套文件：
- 令牌：`web/src/design/tokens.js`、`web/src/design/applyTokens.js`
- 主题：`web/src/themes/index.js`、`web/src/themes/overrides/*`
- 通用组件：`web/src/ui-component/FilterBar.jsx`、`ActionBar.jsx`、`ScrollArea.jsx`

