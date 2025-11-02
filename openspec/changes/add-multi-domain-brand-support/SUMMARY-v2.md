# 多品牌支持方案总结（简化版 v2）

## 核心原则

1. **Kapon AI 保持不变**：作为默认品牌，保持原有部署和代码
2. **Grouplay AI 独立部署**：基于 Kapon AI 克隆，移除 panel，独立部署
3. **统一后端和管理后台**：通过反向代理实现
4. **最小化改动**：只在后端添加品牌管理功能

## 架构设计

### Kapon AI（默认品牌）
- **域名**：models.kapon.cloud
- **部署**：保持原有部署方式
- **功能**：完整功能（用户门户 + 管理后台）
- **改动**：无（可选：支持动态品牌 Logo）

### Grouplay AI（新品牌）
- **域名**：model.grouplay.cn
- **部署**：独立前端服务器 + Nginx 反向代理
- **功能**：用户门户（移除管理后台）
- **改动**：
  - 基于 Kapon AI 前端代码完整克隆
  - 只移除 panel 相关代码（管理后台）
  - 其他代码保持完全一致（包括 UI 文案）
  - 后续可基于此代码进行差异化开发

### 后端服务
- **品牌识别**：根据 Host 头识别品牌
- **品牌管理**：提供 RESTful API 管理品牌
- **品牌配置**：通过 /api/status 返回品牌信息
- **默认品牌**：Kapon AI（models.kapon.cloud）

## 关键技术点

### 1. 品牌识别中间件
```go
// 从 Host 头识别品牌
host := c.Request.Host
brand := BrandManager.GetBrandByDomain(host)
c.Set("brand", brand)
```

### 2. Nginx 反向代理（Grouplay AI）
```nginx
# 前端静态资源
location / {
    root /var/www/grouplay;
    try_files $uri $uri/ /index.html;
}

# API 反向代理（传递 Host 头）
location /api/ {
    proxy_pass http://kapon-backend:3000/api/;
    proxy_set_header Host $host;  # model.grouplay.cn
}

# 管理后台反向代理
location /panel {
    proxy_pass http://kapon-server/panel;
    proxy_set_header Host models.kapon.cloud;
}
```

### 3. 前端品牌配置读取
```javascript
// 从 /api/status 获取品牌配置
const { data } = await fetch('/api/status');
const brandLogo = data.brand_logo;
const systemName = data.system_name;
```

## 实施步骤

### 阶段 1：后端品牌管理（核心）
1. 创建 brands 数据库表
2. 实现品牌数据模型和管理器
3. 实现品牌识别中间件
4. 实现品牌管理 API
5. 修改 /api/status 返回品牌信息
6. 初始化 Kapon AI 默认品牌

### 阶段 2：管理后台界面
1. 创建品牌管理页面
2. 实现品牌列表和表单
3. 实现品牌 API 调用
4. 添加路由和菜单

### 阶段 3：Grouplay AI 前端
1. 完整克隆 Kapon AI 前端代码
2. 只移除 panel 相关代码（管理后台）
3. 清理仅用于管理后台的依赖（可选）
4. 保持其他代码和 UI 文案完全一致
5. 构建和部署
6. 后续基于此代码进行差异化开发

### 阶段 4：部署和配置
1. 配置 Nginx 反向代理
2. 配置 DNS 解析
3. 在管理后台添加 Grouplay AI 品牌
4. 验证功能

## 数据库设计

### brands 表
```sql
CREATE TABLE brands (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) UNIQUE NOT NULL,
    domains JSON NOT NULL,
    system_name VARCHAR(100) NOT NULL,
    logo VARCHAR(255),
    favicon VARCHAR(255),
    description TEXT,
    keywords VARCHAR(255),
    author VARCHAR(100),
    is_default BOOLEAN DEFAULT FALSE,
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 初始化默认品牌
INSERT INTO brands (name, domains, system_name, is_default, enabled)
VALUES ('kapon', '["models.kapon.cloud"]', 'Kapon AI', TRUE, TRUE);
```

## API 设计

### 品牌管理 API
- `GET /api/brands` - 获取品牌列表
- `GET /api/brands/:id` - 获取单个品牌
- `POST /api/brands` - 创建品牌
- `PUT /api/brands/:id` - 更新品牌
- `DELETE /api/brands/:id` - 删除品牌
- `PATCH /api/brands/:id/toggle` - 启用/禁用品牌
- `PATCH /api/brands/:id/set-default` - 设置默认品牌
- `POST /api/brands/refresh` - 刷新缓存

### /api/status 响应（新增字段）
```json
{
  "success": true,
  "data": {
    "version": "v1.0.0",
    "brand_name": "grouplay",
    "brand_logo": "https://cdn.example.com/grouplay-logo.png",
    "brand_favicon": "https://cdn.example.com/grouplay-favicon.ico",
    "system_name": "Grouplay AI",
    "description": "Grouplay AI 企业级AI服务平台",
    "keywords": "AI API,企业AI,Grouplay AI",
    "author": "Grouplay AI"
  }
}
```

## 优势

1. **最小化改动**：Kapon AI 保持不变
2. **灵活扩展**：可以轻松添加新品牌
3. **统一管理**：所有品牌共享后端和管理后台
4. **独立定制**：每个品牌可以独立定制前端
5. **向后兼容**：完全兼容现有系统

## 风险和注意事项

1. **Host 头传递**：确保 Nginx 正确传递 Host 头
2. **Cookie 域名**：跨域访问管理后台需要配置 Cookie
3. **缓存刷新**：品牌配置更新后需要刷新缓存
4. **默认品牌**：Kapon AI 不能删除或禁用
5. **前端克隆**：需要手动移除 panel 相关代码

## 后续扩展

1. 支持更多品牌
2. 品牌级别的主题定制
3. 品牌级别的功能开关
4. 品牌级别的数据隔离（如果需要）
5. 品牌级别的独立配置
