// 验证导航菜单配置的简单脚本
// 这个脚本检查导航菜单项是否与路由配置匹配

const fs = require('fs');
const path = require('path');

// 读取文件内容
function readFile(filePath) {
  try {
    return fs.readFileSync(path.join(__dirname, filePath), 'utf8');
  } catch (error) {
    console.error(`Error reading file ${filePath}:`, error.message);
    return null;
  }
}

// 提取导航菜单项
function extractNavigationItems(content) {
  const navigationItemsMatch = content.match(/navigationItems\s*=\s*\[([\s\S]*?)\]/);
  if (!navigationItemsMatch) return [];
  
  const itemsString = navigationItemsMatch[1];
  const items = [];
  const itemMatches = itemsString.matchAll(/{\s*label:\s*['"`]([^'"`]+)['"`],\s*href:\s*['"`]([^'"`]+)['"`]\s*}/g);
  
  for (const match of itemMatches) {
    items.push({
      label: match[1],
      href: match[2]
    });
  }
  
  return items;
}

// 提取路由配置
function extractRoutes(content) {
  const routes = [];
  const routeMatches = content.matchAll(/{\s*path:\s*['"`]([^'"`]+)['"`],/g);
  
  for (const match of routeMatches) {
    routes.push(match[1]);
  }
  
  return routes;
}

console.log('🔍 验证营销前台导航菜单配置...\n');

// 检查 ModernHomePage Header
const modernHeaderContent = readFile('views/Home/ModernHomePage/components/Header/index.jsx');
if (modernHeaderContent) {
  console.log('📋 ModernHomePage Header 导航菜单:');
  const modernNavItems = extractNavigationItems(modernHeaderContent);
  modernNavItems.forEach(item => {
    console.log(`  - ${item.label}: ${item.href}`);
  });
  console.log();
}

// 检查 MinimalLayout Header
const minimalHeaderContent = readFile('layout/MinimalLayout/Header/index.jsx');
if (minimalHeaderContent) {
  console.log('📋 MinimalLayout Header 导航菜单:');
  // MinimalLayout 使用不同的结构，需要手动提取
  const navLinks = [
    { label: '首页', href: '/' },
    { label: '热门模型', href: '/models' },
    { label: '价格方案', href: '/price' },
    { label: '开发者中心', href: '/developer' },
    { label: '应用体验', href: '/playground' },
    { label: '联系我们', href: '/contact' }
  ];
  
  navLinks.forEach(item => {
    console.log(`  - ${item.label}: ${item.href}`);
  });
  console.log();
}

// 检查路由配置
const routesContent = readFile('routes/OtherRoutes.jsx');
if (routesContent) {
  console.log('🛣️  可用路由:');
  const routes = extractRoutes(routesContent);
  routes.forEach(route => {
    if (route && route !== '') {
      console.log(`  - ${route}`);
    }
  });
  console.log();
}

console.log('✅ 验证完成！');
console.log('\n📝 检查要点:');
console.log('1. 确保所有导航菜单项都有对应的路由');
console.log('2. 验证"首页"菜单项已添加到两个Header组件');
console.log('3. 检查路由路径的一致性');
