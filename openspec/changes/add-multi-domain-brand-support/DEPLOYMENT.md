# 多品牌部署指南

## 概述

本文档说明如何部署多品牌支持功能，包括：
- Kapon AI（默认品牌）：保持原有部署
- Grouplay AI（新品牌）：独立前端部署 + 反向代理

## 架构说明

### Kapon AI（默认品牌）
```
用户 → models.kapon.cloud
  ├─ 前端静态资源（原有部署）
  ├─ /api/* → 后端服务
  └─ /panel → 管理后台
```

### Grouplay AI（新品牌）
```
用户 → model.grouplay.cn
  ├─ 前端静态资源（基于 Kapon AI 克隆）
  ├─ /api/* → 反向代理到 Kapon AI 后端
  └─ /panel → 反向代理到 Kapon AI 管理后台
```

## 部署步骤

### 1. 数据库迁移

```bash
# 执行数据库迁移，创建 brands 表
# 根据项目使用的迁移工具执行

# 迁移会自动创建 Kapon AI 默认品牌记录：
# - name: kapon
# - domains: ["models.kapon.cloud"]
# - system_name: Kapon AI
# - is_default: true
# - enabled: true
```

### 2. 后端部署

后端代码已包含品牌识别功能，无需额外配置。

**验证后端功能**：
```bash
# 测试品牌识别
curl -H "Host: models.kapon.cloud" http://localhost:3000/api/status
curl -H "Host: model.grouplay.cn" http://localhost:3000/api/status

# 应该返回不同的品牌信息
```

### 3. Kapon AI 前端（保持不变）

Kapon AI 前端保持原有部署方式，无需任何修改。

### 4. Grouplay AI 前端部署

#### 4.1 克隆前端代码

```bash
# 假设 Kapon AI 前端在 web/default/ 目录
cd web
cp -r default grouplay

cd grouplay
```

#### 4.2 移除 panel 相关代码

**只移除管理后台相关代码，其他代码保持不变：**

需要移除的内容：
- 管理后台页面和组件（通常在 `/panel` 路径下）
- 管理后台路由配置（如 `/panel/*` 路由）
- 管理员权限检查代码
- 管理功能模块：
  - 用户管理
  - 渠道管理
  - 令牌管理
  - 兑换码管理
  - 系统设置
  - 数据统计
  - 日志查看

保留的功能（与 Kapon AI 完全一致）：
- 用户登录和注册
- 个人资料设置
- 密码修改
- API Key 管理（用户自己的）
- 使用记录查看（用户自己的）
- 所有核心业务功能
- **所有 UI 文案保持不变**

#### 4.3 清理依赖（可选）

```bash
# 检查是否有仅用于管理后台的依赖
# 如果有，可以移除
npm uninstall [管理后台专用的包]

# 更新 package.json
```

#### 4.4 调整品牌标识读取（可选）

**如果需要动态显示品牌 Logo，可以修改 Logo 组件：**

```javascript
// 示例代码
const { data } = await fetch('/api/status');
const brandLogo = data.brand_logo || data.logo;
const systemName = data.system_name;
```

**如果不需要，保持原有代码不变。**

#### 4.5 构建前端

```bash
# 在 web/grouplay/ 目录
npm install
npm run build

# 构建产物输出到 dist/ 目录
```

#### 4.6 部署到服务器

```bash
# 将构建产物复制到服务器
scp -r dist/* user@grouplay-server:/var/www/grouplay/
```

### 5. Nginx 配置

参考 `nginx-example.conf` 文件配置 Nginx。

**关键配置点**：
1. Grouplay AI 静态资源服务
2. `/api/*` 反向代理到 Kapon AI 后端（传递 Host 头）
3. `/panel` 反向代理到 Kapon AI 管理后台

```bash
# 复制配置文件
sudo cp nginx-example.conf /etc/nginx/sites-available/grouplay
sudo ln -s /etc/nginx/sites-available/grouplay /etc/nginx/sites-enabled/

# 测试配置
sudo nginx -t

# 重载 Nginx
sudo nginx -s reload
```

### 6. DNS 配置

配置域名解析：
- `models.kapon.cloud` → Kapon AI 服务器 IP（保持不变）
- `model.grouplay.cn` → Grouplay AI 前端服务器 IP（新增）

### 7. 品牌配置

通过管理后台添加 Grouplay AI 品牌：

1. 访问 `models.kapon.cloud/panel`
2. 登录管理后台
3. 进入"品牌管理"页面
4. 点击"添加品牌"
5. 填写品牌信息：
   - 品牌标识：grouplay
   - 系统名称：Grouplay AI
   - 关联域名：model.grouplay.cn
   - Logo URL：https://cdn.example.com/grouplay-logo.png
   - Favicon URL：https://cdn.example.com/grouplay-favicon.ico
   - 描述、关键词、作者等
6. 保存

### 8. 验证部署

#### 8.1 验证 Kapon AI
```bash
# 访问 Kapon AI
curl https://models.kapon.cloud/api/status

# 应该返回 Kapon AI 品牌信息
```

#### 8.2 验证 Grouplay AI
```bash
# 访问 Grouplay AI
curl https://model.grouplay.cn/api/status

# 应该返回 Grouplay AI 品牌信息
```

#### 8.3 验证管理后台
```bash
# 从 Grouplay AI 访问管理后台
curl https://model.grouplay.cn/panel

# 应该能够正常访问
```

## 故障排查

### 问题 1：API 调用失败

**症状**：Grouplay AI 前端无法调用 API

**解决**：
1. 检查 Nginx 反向代理配置
2. 确认 Host 头正确传递
3. 检查后端日志

### 问题 2：品牌信息不正确

**症状**：返回的品牌信息不是预期的

**解决**：
1. 检查数据库中的品牌配置
2. 确认域名配置正确
3. 刷新品牌缓存：调用 `POST /api/brands/refresh`

### 问题 3：管理后台无法访问

**症状**：从 Grouplay AI 访问 /panel 失败

**解决**：
1. 检查 Nginx 反向代理配置
2. 确认 Kapon AI 管理后台正常运行
3. 检查 Cookie 域名配置

## 回滚方案

如果部署出现问题，可以快速回滚：

1. **Kapon AI**：无需回滚，保持原有部署
2. **Grouplay AI**：
   - 停止 Nginx 服务
   - 删除 Grouplay AI 配置
   - 从数据库删除 Grouplay AI 品牌记录

## 添加新品牌

要添加新品牌（如 Brand X），重复以下步骤：

1. 克隆 Kapon AI 前端代码
2. 移除 panel 相关代码
3. 精简和定制
4. 构建和部署
5. 配置 Nginx
6. 配置 DNS
7. 在管理后台添加品牌配置

## 注意事项

1. **Host 头传递**：确保 Nginx 正确传递 Host 头到后端
2. **Cookie 域名**：如果需要跨域访问管理后台，配置 Cookie 域名
3. **HTTPS**：生产环境建议使用 HTTPS
4. **缓存刷新**：品牌配置更新后，调用刷新缓存 API
5. **默认品牌**：Kapon AI 作为默认品牌，不能删除或禁用
