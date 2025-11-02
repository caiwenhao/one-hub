# Grouplay AI 前端克隆和精简指南

## 目标

基于 Kapon AI 前端代码创建 Grouplay AI 前端，只移除管理后台（panel）相关代码，其他代码和 UI 文案保持完全一致。

## 原则

1. **完整克隆**：复制所有 Kapon AI 前端代码
2. **只移除 panel**：只删除管理后台相关代码
3. **保持一致**：其他代码、UI 文案、功能保持完全一致
4. **便于后续开发**：为后续差异化开发提供干净的基础

## 步骤

### 1. 克隆前端代码

```bash
# 假设 Kapon AI 前端在 web/default/ 目录
cd web
cp -r default grouplay

cd grouplay
```

### 2. 识别 panel 相关代码

需要识别以下类型的代码：

#### 2.1 路由配置
查找包含 `/panel` 路径的路由定义：
```javascript
// 示例：需要删除的路由
{
  path: '/panel',
  component: PanelLayout,
  children: [...]
}

{
  path: '/panel/users',
  component: UserManagement
}
```

#### 2.2 页面和组件
查找管理后台相关的页面和组件目录：
```
src/pages/Panel/
src/pages/UserManagement/
src/pages/ChannelManagement/
src/pages/TokenManagement/
src/pages/RedemptionManagement/
src/pages/SystemSettings/
src/pages/Statistics/
src/pages/Logs/
src/components/AdminLayout/
src/components/AdminSidebar/
```

#### 2.3 权限检查代码
查找管理员权限检查相关代码：
```javascript
// 示例：需要删除的权限检查
if (user.role === 'admin') {
  // 管理员功能
}

const isAdmin = checkAdminPermission();
```

#### 2.4 导航菜单
查找导航菜单中的管理后台入口：
```javascript
// 示例：需要删除的菜单项
{
  title: '管理后台',
  path: '/panel',
  icon: AdminIcon
}
```

### 3. 移除 panel 相关代码

#### 3.1 删除路由配置

**文件位置**：通常在 `src/router/` 或 `src/routes/` 目录

```javascript
// 删除前
const routes = [
  {
    path: '/',
    component: Home
  },
  {
    path: '/panel',  // 删除这个路由
    component: PanelLayout,
    children: [...]
  }
];

// 删除后
const routes = [
  {
    path: '/',
    component: Home
  }
  // panel 路由已删除
];
```

#### 3.2 删除页面和组件

```bash
# 删除管理后台相关目录
rm -rf src/pages/Panel
rm -rf src/pages/UserManagement
rm -rf src/pages/ChannelManagement
rm -rf src/pages/TokenManagement
rm -rf src/pages/RedemptionManagement
rm -rf src/pages/SystemSettings
rm -rf src/pages/Statistics
rm -rf src/pages/Logs
rm -rf src/components/AdminLayout
rm -rf src/components/AdminSidebar

# 或者根据实际项目结构调整
```

#### 3.3 删除权限检查代码

查找并删除管理员权限相关代码：

```bash
# 搜索管理员权限相关代码
grep -r "isAdmin" src/
grep -r "role === 'admin'" src/
grep -r "checkAdminPermission" src/

# 手动删除相关代码
```

#### 3.4 删除导航菜单项

**文件位置**：通常在 `src/components/Navigation/` 或 `src/layouts/`

```javascript
// 删除前
const menuItems = [
  { title: '首页', path: '/' },
  { title: '个人中心', path: '/profile' },
  { title: '管理后台', path: '/panel' }  // 删除这一项
];

// 删除后
const menuItems = [
  { title: '首页', path: '/' },
  { title: '个人中心', path: '/profile' }
];
```

### 4. 保留的功能

**重要**：以下功能必须保留，不要删除：

#### 4.1 用户功能
- 用户登录和注册
- 个人资料设置
- 密码修改
- 邮箱验证
- 忘记密码

#### 4.2 API Key 管理
- 查看自己的 API Key
- 创建新的 API Key
- 删除 API Key
- 重置 API Key

#### 4.3 使用记录
- 查看自己的使用记录
- 查看消费统计
- 查看余额

#### 4.4 核心业务功能
- 所有用户可见的业务功能
- 所有公开页面
- 所有用户交互功能

### 5. 清理依赖（可选）

检查 `package.json` 中是否有仅用于管理后台的依赖：

```bash
# 查看依赖
cat package.json

# 如果有仅用于管理后台的依赖，可以移除
# 例如：管理后台专用的图表库、表格库等
npm uninstall [package-name]
```

**注意**：如果不确定某个依赖是否只用于管理后台，建议保留。

### 6. 验证和测试

#### 6.1 检查编译错误

```bash
# 安装依赖
npm install

# 检查是否有编译错误
npm run build
```

#### 6.2 修复导入错误

如果有编译错误，通常是因为某些文件导入了已删除的组件：

```javascript
// 错误示例
import AdminLayout from '@/components/AdminLayout';  // 已删除

// 解决方案：删除这个导入语句
```

#### 6.3 测试运行

```bash
# 启动开发服务器
npm run dev

# 访问 http://localhost:3000
# 测试以下功能：
# - 用户登录
# - 个人设置
# - API Key 管理
# - 使用记录查看
# - 确认 /panel 路径不可访问（404）
```

### 7. 构建和部署

```bash
# 构建生产版本
npm run build

# 构建产物在 dist/ 目录
ls -la dist/

# 部署到服务器
scp -r dist/* user@grouplay-server:/var/www/grouplay/
```

## 常见问题

### Q1: 如何确定哪些代码是 panel 相关的？

**A**: 通常可以通过以下方式识别：
- 路径包含 `/panel`
- 文件名包含 `Admin`、`Panel`、`Management`
- 组件名包含 `Admin`、`Panel`
- 需要管理员权限才能访问的功能

### Q2: 删除代码后出现编译错误怎么办？

**A**: 
1. 检查错误信息，找到引用已删除代码的文件
2. 删除相关的导入语句
3. 删除相关的使用代码
4. 重新编译

### Q3: 如何确保没有遗漏需要删除的代码？

**A**: 
1. 搜索关键词：`panel`、`admin`、`management`
2. 检查路由配置文件
3. 检查导航菜单配置
4. 测试访问 `/panel` 路径，确认返回 404

### Q4: UI 文案需要修改吗？

**A**: 不需要。保持所有 UI 文案与 Kapon AI 完全一致。后续可以根据需要进行差异化修改。

### Q5: 需要修改 Logo 或品牌标识吗？

**A**: 不需要在这个阶段修改。品牌标识会通过后端 `/api/status` 接口动态返回。如果需要，可以在后续开发中调整。

## 检查清单

完成以下检查，确保 Grouplay AI 前端正确配置：

- [ ] 已完整克隆 Kapon AI 前端代码
- [ ] 已删除所有 panel 相关路由
- [ ] 已删除所有管理后台页面和组件
- [ ] 已删除管理员权限检查代码
- [ ] 已删除导航菜单中的管理后台入口
- [ ] 保留了所有用户功能
- [ ] 保留了所有 UI 文案
- [ ] 编译无错误
- [ ] 开发服务器可以正常运行
- [ ] 用户登录功能正常
- [ ] 个人设置功能正常
- [ ] API Key 管理功能正常
- [ ] 访问 `/panel` 返回 404
- [ ] 构建产物完整

## 后续开发

完成以上步骤后，Grouplay AI 前端已经准备就绪。后续可以基于此代码进行差异化开发：

1. 修改 UI 文案和品牌标识
2. 调整页面布局和设计
3. 添加或修改功能
4. 定制用户体验

所有修改都在 `web/grouplay/` 目录中进行，不会影响 Kapon AI 前端。
