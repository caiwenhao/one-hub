# 多品牌多前端支持 - 方案总结

## 更新概览

本次更新将原有的"多域名多品牌支持"提案扩展为完整的多前端架构，支持每个品牌使用完全独立的前端项目。

## 核心需求

### 1. 前端门户
- ✅ **完全独立**：每个品牌使用独立的前端项目
- ✅ **UI 完全不同**：支持不同的设计、布局、文案、甚至技术栈
- ✅ **独立开发**：各品牌前端可以独立开发、构建和部署

### 2. 管理后台
- ✅ **共享 UI**：所有品牌使用相同的管理后台
- ✅ **品牌标识差异化**：只有 logo 和系统名称根据品牌动态显示
- ✅ **统一访问**：通过各域名的 `/panel` 路径访问

### 3. 后端服务
- ✅ **完全共享**：所有品牌共享同一个后端服务
- ✅ **数据共享**：所有品牌的管理员看到相同的数据
- ✅ **域名识别**：完全依赖域名识别，前端无需额外标识品牌

### 4. 部署灵活性
- ✅ **统一部署**：支持所有前端和后端部署在同一服务器
- ✅ **独立部署**：支持前端独立部署到 CDN（可选）

## 架构设计

```
用户访问
    ↓
models.kapon.cloud → Kapon 前端（独立项目）
model.grouplay.cn  → Grouplay 前端（独立项目）
*.*/panel          → 管理后台（共享项目，显示对应品牌 logo）
    ↓
后端服务（共享）
    ├─ 品牌识别中间件（基于域名）
    ├─ 前端资源路由（返回对应品牌的前端）
    ├─ API 服务（所有品牌共享）
    └─ 数据库（所有品牌共享）
```

## 前端项目结构

```
web/
├── kapon-portal/          # Kapon 独立前端（可用 React）
│   ├── package.json
│   ├── src/
│   └── dist/ → ../public/brands/kapon/
│
├── grouplay-portal/       # Grouplay 独立前端（可用 Vue）
│   ├── package.json
│   ├── src/
│   └── dist/ → ../public/brands/grouplay/
│
├── admin/                 # 管理后台（共享）
│   ├── package.json
│   ├── src/
│   └── dist/ → ../public/admin/
│
└── public/
    ├── brands/
    │   ├── kapon/         # Kapon 前端构建产物
    │   └── grouplay/      # Grouplay 前端构建产物
    └── admin/             # 管理后台构建产物
```

## 配置方式

### 通过管理后台配置（推荐）

1. 登录管理后台
2. 进入"品牌管理"页面
3. 点击"添加品牌"按钮
4. 填写品牌信息：
   - 品牌标识：kapon
   - 系统名称：Kapon AI
   - 关联域名：models.kapon.cloud, localhost:3000
   - Logo 路径：/brands/kapon/logo.png
   - Favicon 路径：/brands/kapon/favicon.ico
   - 前端类型：embedded
   - 前端资源路径：/brands/kapon/
5. 保存后配置立即生效

### 数据库表结构

```sql
CREATE TABLE brands (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) UNIQUE NOT NULL,
    domains TEXT NOT NULL,  -- JSON 数组
    system_name VARCHAR(100),
    logo VARCHAR(255),
    favicon VARCHAR(255),
    description TEXT,
    keywords TEXT,
    author VARCHAR(100),
    frontend_type VARCHAR(20) DEFAULT 'embedded',
    frontend_path VARCHAR(255),
    frontend_url VARCHAR(255),
    is_default BOOLEAN DEFAULT FALSE,
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## 实现任务

总共 **45 个任务**，分为 11 个主要模块：

1. **数据库表和数据模型** (4 个任务) ⭐ 新增
   - 创建 brands 数据库表
   - 品牌数据模型（CRUD 操作）
   - 品牌管理器（缓存层）
   - 单元测试

2. **品牌识别中间件** (3 个任务)
   - 中间件实现
   - 路由注册
   - 单元测试

3. **品牌管理 API** (5 个任务) ⭐ 新增
   - 品牌管理控制器（增删改查）
   - 配置验证逻辑
   - 路由注册
   - 自动刷新缓存
   - 集成测试

4. **GetStatus 接口修改** (2 个任务)
   - 返回品牌信息
   - 集成测试

5. **前端资源路由系统** (3 个任务)
   - 路由处理器
   - 路由注册
   - 单元测试

6. **品牌管理界面** (7 个任务) ⭐ 新增
   - 品牌列表页面
   - 品牌表单页面
   - 域名输入组件
   - API 封装
   - 路由配置
   - 菜单项添加
   - 界面测试

7. **多前端项目架构** (4 个任务)
   - Kapon 前端项目
   - Grouplay 前端项目
   - 管理后台改造
   - 构建脚本配置

8. **管理后台品牌标识支持** (3 个任务)
   - Logo 组件修改
   - 品牌配置集成
   - 测试

9. **CORS 配置** (2 个任务，可选)
   - CORS 中间件
   - CORS 测试

10. **集成测试与验证** (6 个任务)
    - 多前端路由测试
    - API 调用测试
    - SPA 路由测试
    - 品牌管理界面测试
    - 向后兼容性测试
    - 边界情况测试

11. **文档与部署** (6 个任务)
    - 数据库迁移文档
    - 品牌管理使用文档
    - 多前端部署文档
    - README 更新
    - 前端开发指南
    - 验证脚本

## 关键技术点

### 1. 域名识别
- 后端从 HTTP Host 头自动识别品牌
- 前端无需在 API 请求中添加品牌参数
- 支持带端口的域名（如 localhost:3000）

### 2. 前端资源路由
- 根据品牌配置自动路由到对应前端
- 支持 SPA 路由（非文件路径返回 index.html）
- `/panel` 路径统一路由到管理后台

### 3. 前端独立性
- 每个品牌前端完全独立
- 可以使用不同的技术栈（React、Vue 等）
- 独立的 package.json 和构建配置

### 4. 管理后台共享
- 所有品牌共享同一个管理后台代码
- 只有 logo 和系统名称动态显示
- 数据在所有品牌间共享

## 部署模式

### 模式 1：统一部署（推荐）
- 所有前端和后端部署在同一服务器
- 通过域名区分品牌
- 配置简单，运维方便

### 模式 2：前端独立部署
- 前端部署到 CDN 或静态服务器
- 后端只提供 API 服务
- 需要配置 CORS

## 向后兼容

- ✅ 未配置 brands 时自动回退到单品牌模式
- ✅ 使用全局配置（SystemName, Logo）
- ✅ 所有功能正常工作
- ✅ 无需修改现有配置文件

## 扩展性

### 添加新品牌
1. 在配置文件中添加新品牌配置
2. 创建新的前端项目 `web/newbrand-portal/`
3. 构建并输出到 `web/public/brands/newbrand/`
4. 重启服务器

### 支持更多前端类型
未来可以扩展：
- `frontend_type: redirect` - 重定向到外部 URL
- `frontend_type: proxy` - 反向代理到其他服务
- `frontend_type: dynamic` - 动态生成前端内容

## 相关文档

- `proposal.md` - 详细的变更提案
- `architecture.md` - 架构设计文档
- `config-example.yaml` - 配置示例
- `specs/brand-management/spec.md` - 需求规格说明
- `tasks.md` - 实现任务清单

## 下一步

1. **审查方案**：确认架构设计和技术方案
2. **开始实现**：按照 tasks.md 的顺序实施
3. **测试验证**：完成后进行全面测试
4. **部署上线**：准备部署文档和迁移计划

---

**状态**: 方案设计完成，等待审查和实施
**任务数**: 45 个任务（其中 10 个为可选任务）
**预计工作量**: 中等偏大（需要前后端协同开发 + 数据库 + 管理界面）

## 关键变更（相比配置文件方案）

✅ **数据库存储** - 品牌配置存储在数据库中，支持动态管理  
✅ **管理界面** - 提供完整的品牌管理 UI，无需修改配置文件  
✅ **即时生效** - 配置更改后自动刷新缓存，无需重启服务  
✅ **更易用** - 管理员可以自助操作，降低技术门槛  
✅ **更灵活** - 支持在线添加、编辑、删除品牌
