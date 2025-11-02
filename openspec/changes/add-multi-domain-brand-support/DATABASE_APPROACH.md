# 数据库 + 管理界面方案说明

## 方案对比

### 原方案：配置文件
- ❌ 需要修改配置文件
- ❌ 需要重启服务
- ❌ 需要技术人员操作
- ✅ 简单直接

### 新方案：数据库 + 管理界面
- ✅ 通过管理后台配置
- ✅ 动态生效，无需重启
- ✅ 管理员可自助操作
- ✅ 更加灵活
- ⚠️ 需要数据库表
- ⚠️ 需要开发管理界面

## 核心变化

### 1. 数据存储

**数据库表**：
```sql
CREATE TABLE brands (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) UNIQUE NOT NULL,
    domains TEXT NOT NULL,  -- JSON 数组：["models.kapon.cloud", "localhost:3000"]
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

**索引**：
- `name` 字段：唯一索引
- `enabled` 字段：普通索引（查询启用的品牌）
- `is_default` 字段：普通索引（查询默认品牌）

### 2. 缓存机制

**启动时加载**：
```go
func main() {
    // ... 初始化数据库
    
    // 加载品牌配置到内存
    brandManager := NewBrandManager()
    if err := brandManager.LoadFromDatabase(); err != nil {
        log.Error("Failed to load brands:", err)
    }
    
    // ... 启动服务
}
```

**动态刷新**：
```go
// 品牌配置更新后自动刷新
func (c *BrandController) CreateBrand(ctx *gin.Context) {
    // ... 创建品牌
    
    // 刷新缓存
    brandManager.RefreshCache()
    
    // ... 返回响应
}
```

### 3. 管理界面

**品牌列表页面**：
- 显示所有品牌
- 支持搜索和筛选
- 快速启用/禁用
- 设置默认品牌
- 编辑和删除操作

**品牌表单页面**：
- 完整的表单字段
- 实时验证
- 域名标签输入（支持多个）
- 前端类型切换（embedded/external）
- 友好的提示信息

**API 接口**：
```
GET    /api/brands           - 获取品牌列表
GET    /api/brands/:id       - 获取单个品牌
POST   /api/brands           - 创建品牌
PUT    /api/brands/:id       - 更新品牌
DELETE /api/brands/:id       - 删除品牌
PATCH  /api/brands/:id/toggle      - 启用/禁用
PATCH  /api/brands/:id/set-default - 设置默认
POST   /api/brands/refresh   - 刷新缓存
```

### 4. 配置验证

**后端验证**：
- 必填字段检查
- 品牌名称格式验证（小写字母、数字、连字符）
- 域名格式验证
- 域名唯一性验证（跨品牌）
- 前端配置完整性验证

**前端验证**：
- 实时表单验证
- 友好的错误提示
- 防止重复提交

## 使用流程

### 添加新品牌

1. **准备品牌资源**
   ```bash
   # 创建品牌资源目录
   mkdir -p web/public/brands/newbrand
   
   # 放置 logo 和 favicon
   cp logo.png web/public/brands/newbrand/
   cp favicon.ico web/public/brands/newbrand/
   ```

2. **创建前端项目**
   ```bash
   # 创建前端项目目录
   mkdir web/newbrand-portal
   cd web/newbrand-portal
   
   # 初始化项目
   npm init -y
   npm install react vite
   
   # 配置构建输出到 ../public/brands/newbrand/
   ```

3. **通过管理后台配置**
   - 登录管理后台
   - 进入"品牌管理"
   - 点击"添加品牌"
   - 填写表单：
     - 品牌标识：newbrand
     - 系统名称：New Brand
     - 关联域名：new.example.com
     - Logo 路径：/brands/newbrand/logo.png
     - Favicon 路径：/brands/newbrand/favicon.ico
     - 前端类型：embedded
     - 前端资源路径：/brands/newbrand/
   - 保存

4. **验证配置**
   - 访问 http://new.example.com
   - 验证返回新品牌的前端
   - 验证 logo 和 favicon 正确显示

### 修改品牌配置

1. 登录管理后台
2. 进入"品牌管理"
3. 找到要修改的品牌，点击"编辑"
4. 修改配置
5. 保存（配置立即生效）

### 删除品牌

1. 登录管理后台
2. 进入"品牌管理"
3. 找到要删除的品牌，点击"删除"
4. 确认删除
5. 品牌配置立即失效

注意：默认品牌不能删除

## 技术实现要点

### 1. 域名存储

domains 字段使用 JSON 格式存储：
```json
["models.kapon.cloud", "localhost:3000", "kapon.example.com"]
```

**序列化/反序列化**：
```go
type Brand struct {
    // ... 其他字段
    Domains []string `gorm:"type:text;serializer:json"`
}
```

### 2. 缓存更新

**线程安全**：
```go
type BrandManager struct {
    brands      map[string]*Brand
    domainMap   map[string]*Brand
    defaultBrand *Brand
    mutex       sync.RWMutex
}

func (m *BrandManager) RefreshCache() error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    // 重新从数据库加载
    brands, err := model.GetEnabledBrands()
    if err != nil {
        return err
    }
    
    // 更新缓存
    m.brands = make(map[string]*Brand)
    m.domainMap = make(map[string]*Brand)
    
    for _, brand := range brands {
        m.brands[brand.Name] = brand
        for _, domain := range brand.Domains {
            m.domainMap[domain] = brand
        }
        if brand.IsDefault {
            m.defaultBrand = brand
        }
    }
    
    return nil
}
```

### 3. 默认品牌逻辑

**确保唯一性**：
```go
func SetDefaultBrand(id int) error {
    // 开启事务
    tx := db.Begin()
    
    // 取消所有品牌的默认状态
    tx.Model(&Brand{}).Update("is_default", false)
    
    // 设置指定品牌为默认
    tx.Model(&Brand{}).Where("id = ?", id).Update("is_default", true)
    
    // 提交事务
    return tx.Commit().Error
}
```

### 4. 域名唯一性验证

```go
func ValidateDomainUniqueness(domains []string, excludeBrandID int) error {
    for _, domain := range domains {
        var count int64
        db.Model(&Brand{}).
            Where("id != ?", excludeBrandID).
            Where("JSON_CONTAINS(domains, ?)", fmt.Sprintf(`"%s"`, domain)).
            Count(&count)
        
        if count > 0 {
            return fmt.Errorf("域名 %s 已被其他品牌使用", domain)
        }
    }
    return nil
}
```

## 向后兼容

### 无品牌配置时

如果数据库中没有任何品牌配置：
- 系统自动使用单品牌模式
- 使用全局配置（SystemName, Logo）
- 所有功能正常工作

### 迁移路径

从配置文件迁移到数据库：
1. 执行数据库迁移创建 brands 表
2. 通过管理后台添加品牌配置
3. 验证配置正确
4. 删除配置文件中的 brands 配置段（可选）

## 性能考虑

### 内存占用

假设 100 个品牌，每个品牌 5 个域名：
- 品牌对象：~100 * 1KB = 100KB
- 域名映射：~500 * 100B = 50KB
- 总计：~150KB（可忽略不计）

### 查询性能

- 域名查询：O(1) - 从内存 map 查询
- 品牌列表：O(1) - 从内存读取
- 缓存刷新：O(n) - n 为品牌数量，通常很小

### 数据库压力

- 启动时：1 次查询（加载所有品牌）
- 运行时：0 次查询（从内存读取）
- 配置更新时：1 次查询（刷新缓存）

## 安全考虑

### 权限控制

- 品牌管理 API 需要管理员权限
- 普通用户无法访问品牌管理界面
- API 需要认证和授权

### 输入验证

- 品牌名称：防止 SQL 注入
- 域名：防止 XSS 攻击
- 路径：防止路径遍历

### 审计日志

建议记录品牌配置的变更：
- 谁在什么时候做了什么操作
- 变更前后的值
- 操作结果

## 总结

数据库 + 管理界面方案相比配置文件方案：

**优势**：
- ✅ 更易用（管理员可自助操作）
- ✅ 更灵活（动态配置，无需重启）
- ✅ 更安全（权限控制，审计日志）
- ✅ 更可靠（数据持久化，事务保证）

**劣势**：
- ⚠️ 实现复杂度略高（需要数据库表和管理界面）
- ⚠️ 依赖数据库（但项目本身就依赖数据库）

**推荐使用场景**：
- ✅ 需要频繁添加/修改品牌
- ✅ 多人协作管理品牌
- ✅ 需要审计和权限控制
- ✅ 希望降低运维复杂度

总体来说，数据库 + 管理界面方案更适合生产环境使用。
