# 多品牌支持 - 快速参考

## 核心概念

### Kapon AI（默认品牌）
- **域名**：models.kapon.cloud
- **部署**：保持原有部署，无需改动
- **功能**：完整功能（用户门户 + 管理后台）

### Grouplay AI（新品牌）
- **域名**：model.grouplay.cn
- **部署**：独立前端 + Nginx 反向代理
- **功能**：用户门户（无管理后台）
- **代码**：基于 Kapon AI 克隆，只移除 panel

## 关键文件

| 文件 | 说明 |
|------|------|
| `proposal.md` | 提案概述 |
| `tasks.md` | 实现任务清单 |
| `design.md` | 技术设计文档 |
| `DEPLOYMENT.md` | 部署指南 |
| `FRONTEND-CLONE-GUIDE.md` | 前端克隆详细指南 |
| `nginx-example.conf` | Nginx 配置示例 |
| `SUMMARY-v2.md` | 方案总结 |

## 实施顺序

### 阶段 1：后端品牌管理
1. 创建 brands 数据库表
2. 实现品牌数据模型和管理器
3. 实现品牌识别中间件
4. 实现品牌管理 API
5. 修改 /api/status 接口
6. 初始化 Kapon AI 默认品牌

### 阶段 2：管理后台界面
1. 创建品牌管理页面
2. 实现品牌列表和表单
3. 实现品牌 API 调用

### 阶段 3：Grouplay AI 前端
1. 克隆 Kapon AI 前端代码
2. 只移除 panel 相关代码
3. 保持其他代码完全一致
4. 构建和部署

### 阶段 4：部署配置
1. 配置 Nginx
2. 配置 DNS
3. 添加品牌配置
4. 验证功能

## 关键命令

### 数据库迁移
```bash
# 执行迁移（根据项目使用的工具）
# 会自动创建 Kapon AI 默认品牌
```

### 前端克隆
```bash
cd web
cp -r default grouplay
cd grouplay
# 移除 panel 相关代码（参考 FRONTEND-CLONE-GUIDE.md）
npm install
npm run build
```

### Nginx 配置
```bash
# 复制配置文件
sudo cp nginx-example.conf /etc/nginx/sites-available/grouplay
sudo ln -s /etc/nginx/sites-available/grouplay /etc/nginx/sites-enabled/
sudo nginx -t
sudo nginx -s reload
```

### 验证
```bash
# 测试 Kapon AI
curl https://models.kapon.cloud/api/status

# 测试 Grouplay AI
curl https://model.grouplay.cn/api/status

# 测试管理后台
curl https://model.grouplay.cn/panel
```

## API 端点

### 品牌管理 API
- `GET /api/brands` - 获取品牌列表
- `POST /api/brands` - 创建品牌
- `PUT /api/brands/:id` - 更新品牌
- `DELETE /api/brands/:id` - 删除品牌
- `PATCH /api/brands/:id/toggle` - 启用/禁用
- `PATCH /api/brands/:id/set-default` - 设置默认
- `POST /api/brands/refresh` - 刷新缓存

### 状态接口（新增字段）
```json
{
  "brand_name": "grouplay",
  "brand_logo": "https://...",
  "brand_favicon": "https://...",
  "system_name": "Grouplay AI",
  "description": "...",
  "keywords": "...",
  "author": "..."
}
```

## 数据库表

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
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

## Nginx 配置要点

### Grouplay AI
```nginx
# 前端静态资源
location / {
    root /var/www/grouplay;
    try_files $uri $uri/ /index.html;
}

# API 反向代理（传递 Host 头）
location /api/ {
    proxy_pass http://kapon-backend:3000/api/;
    proxy_set_header Host $host;  # 重要！
}

# 管理后台反向代理
location /panel {
    proxy_pass http://kapon-server/panel;
    proxy_set_header Host models.kapon.cloud;
}
```

## 常见问题

### Q: Kapon AI 需要修改代码吗？
**A**: 不需要。保持原有部署和代码不变。

### Q: Grouplay AI 前端需要修改 UI 文案吗？
**A**: 不需要。保持与 Kapon AI 完全一致。后续可以根据需要修改。

### Q: 如何添加新品牌？
**A**: 
1. 克隆前端代码
2. 移除 panel
3. 配置 Nginx
4. 在管理后台添加品牌配置

### Q: 品牌配置更新后需要重启吗？
**A**: 不需要。调用 `POST /api/brands/refresh` 刷新缓存即可。

### Q: 如何访问管理后台？
**A**: 
- Kapon AI: `models.kapon.cloud/panel`
- Grouplay AI: `model.grouplay.cn/panel`（反向代理到 Kapon AI）

## 注意事项

1. ✅ Kapon AI 保持不变
2. ✅ Grouplay AI 只移除 panel
3. ✅ 保持 UI 文案一致
4. ✅ Nginx 必须传递 Host 头
5. ✅ 默认品牌不能删除
6. ✅ 品牌配置更新后刷新缓存

## 下一步

1. 阅读 `tasks.md` 了解详细任务
2. 阅读 `FRONTEND-CLONE-GUIDE.md` 了解前端克隆步骤
3. 阅读 `DEPLOYMENT.md` 了解部署流程
4. 开始实施第一个任务
