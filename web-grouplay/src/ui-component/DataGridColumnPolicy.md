# 列宽策略（DataGrid）
- 最小宽度：80px；建议：文本列 ≥ 120px，数字列 ≥ 96px
- 最大宽度：480px；说明：过宽会影响信息密度与滚动体验
- 持久化：使用 onStateChange 捕获 columnLookup.computedWidth 并存入 localStorage
- 默认对齐：数字列右对齐、文本左对齐；超长文本采用 ellipsis 并提供 Tooltip
