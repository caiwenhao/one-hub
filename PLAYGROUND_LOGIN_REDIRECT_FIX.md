# 应用体验页面登录重定向修复

## 问题描述

在未登录状态下访问应用体验页面（`/playground`）时，页面会显示错误而不是跳转到登录页面。这是因为：

1. Playground 页面可以正常加载（没有路由级别的认证保护）
2. 但调用 `/api/token/playground` API 时会因为未认证而失败
3. 用户看到错误信息而不是被引导去登录

## 修复方案

### 1. Playground 页面修改 (`web/src/views/Playground/index.jsx`)

- 添加了登录状态检查逻辑
- 未登录时自动跳转到登录页面，并携带重定向参数 `?redirect=/playground`
- API 调用失败时也会跳转到登录页面
- 在用户信息加载完成前显示加载状态

### 2. 登录 Hook 修改 (`web/src/hooks/useLogin.js`)

- 添加了 `useLocation` 来获取 URL 参数
- 创建了 `getRedirectUrl()` 辅助函数来解析重定向参数
- 修改了所有登录方法（普通登录、GitHub、OIDC、Lark、WeChat）在成功后跳转到重定向URL

### 3. 认证包装器修改 (`web/src/views/Authentication/AuthWrapper.jsx`)

- 已登录用户访问登录页面时，也会检查重定向参数并跳转到目标页面

### 4. WebAuthn 登录修改 (`web/src/utils/common.jsx`)

- 修改了 `onWebAuthnClicked` 函数，登录成功后检查重定向参数并跳转

## 使用流程

1. **未登录用户访问 `/playground`**
   - 页面检测到未登录状态
   - 自动跳转到 `/login?redirect=/playground`

2. **用户完成登录**
   - 系统检测到 `redirect` 参数
   - 登录成功后跳转到 `/playground` 而不是默认的 `/panel`

3. **已登录用户访问带重定向的登录页面**
   - 直接跳转到重定向目标页面

## 支持的登录方式

- ✅ 用户名密码登录
- ✅ GitHub OAuth 登录
- ✅ OIDC 登录  
- ✅ Lark 登录
- ✅ WeChat 登录
- ✅ WebAuthn 登录

## 测试场景

1. **基本重定向测试**
   - 未登录访问 `/playground`
   - 应该跳转到 `/login?redirect=/playground`
   - 登录后应该回到 `/playground`

2. **API 错误处理测试**
   - 登录状态过期时访问 `/playground`
   - API 返回 401 错误
   - 应该跳转到登录页面并携带重定向参数

3. **已登录用户测试**
   - 已登录用户直接访问 `/login?redirect=/playground`
   - 应该直接跳转到 `/playground`

## 技术细节

- 使用 URL 查询参数 `redirect` 来传递重定向目标
- 默认重定向目标为 `/panel`（保持向后兼容）
- 所有登录方式都统一支持重定向功能
- 保持了原有的用户体验和错误处理逻辑
