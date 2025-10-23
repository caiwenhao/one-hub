# 前端与控制台功能更新汇总

> 记录近期完成的 UI 与使用体验相关改动，便于快速了解已交付能力与回归重点。
## 开发者中心（Developer Center）改版
- **界面升级**：新增渐变背景、装饰元素、悬浮动画，一键复制 CodeBlock，完整响应式优化。
- **语法高亮**：引入 `modern-code.css`，按语言定制配色并支持 Python/JS/Bash 等别名，修复切换标签重新渲染问题。
- **代码位置**：`web/src/views/DeveloperCenter/index.jsx`、`web/src/ui-component/highlight.js`、`web/src/ui-component/CodeBlock.jsx`、`web/src/assets/css/modern-code.css`。
- **验证建议**：分别检查多语言代码块复制、移动端显示、动画性能与 API base URL 是否统一为 `https://models.kapon.cloud/v1`。
## 应用体验页登录跳转
- **问题**：未登录访问 `/playground` 时 API 返回 401，但页面未重定向。
- **方案**：在 `web/src/views/Playground/index.jsx` 检测登录状态并统一跳转到 `/login?redirect=/playground`，登录钩子 `useLogin.js`、`AuthWrapper.jsx`、`onWebAuthnClicked` 均支持读取重定向参数。
- **验证建议**：覆盖未登录访问、令牌过期、带 `redirect` 参数重复登录、WebAuthn/OIDC 等多种登录方式。
## 价格方案页（/price）重构
- **视觉与交互**：Hero 渐变背景、浮动装饰、CTA 卡片、15 种动画效果；表格统一展示所有模型，支持供应商/用户组筛选与搜索。
- **国际化**：`web/src/i18n/locales/zh_CN.json` 与 `en_US.json` 新增 `modelpricePage` 区块，实现实时切换。
- **主要文件**：`web/src/views/ModelPrice/index.jsx`、`components/`、`styles/animations.js`。
- **验证建议**：断点适配（xs-xl）、表格筛选与 API `/api/available_model`、`/api/user_group_map` 返回值兼容性。
## 营销前台导航统一
- **范围**：`web/src/views/Home/ModernHomePage/components/Header/index.jsx` 与 `web/src/layout/MinimalLayout/Header/index.jsx`。
- **改动**：补充“首页/热门模型/价格方案/开发者中心/应用体验/联系我们”统一导航，修正路由映射并强化活跃态样式。
- **验证建议**：桌面与移动端导航一致性、Logo 与菜单跳转、`/models` 与 `/developer` 等路由是否连通。
## SDK 栏位与代码示例更新
- **内容**：将开发者中心“SDK 与工具库”改为全部使用 OpenAI 官方 SDK；新增 Ruby，更新 Go 为 `github.com/openai/openai-go@v2.1.1`。
- **界面**：更新卡片图标颜色、安装命令展示、说明文案统一强调“更换 Base URL 即可接入”。
- **示例**：Python/Node.js/Ruby/Go 代码分别位于 `web/src/views/DeveloperCenter/components/` 相关模块。
- **验证建议**：检查命令区展示、复制按钮、国际化文本与 `openai` 依赖版本。
---

> 若需查看更细粒度的 UI 基线截图、DataGrid 调整等资料，请参考 `docs/ui-baseline/` 与组件目录下的 README 文件。
