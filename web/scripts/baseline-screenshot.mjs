// 基线截图脚本：构建 -> 预览 -> 打开 /panel/styleguide?baseline=1 -> 滚动到 Before/After -> 截图
import { spawn } from 'node:child_process';
import http from 'node:http';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { chromium } from 'playwright';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

function run(cmd, args, opts = {}) {
  return new Promise((resolve, reject) => {
    const p = spawn(cmd, args, { stdio: 'inherit', ...opts });
    p.on('close', (code) => (code === 0 ? resolve() : reject(new Error(`${cmd} ${args.join(' ')} exit ${code}`))));
    p.on('error', reject);
  });
}

function waitForUrl(url, timeout = 30000) {
  const started = Date.now();
  return new Promise((resolve, reject) => {
    const tryOnce = () => {
      const req = http.get(url, (res) => {
        if (res.statusCode && res.statusCode >= 200 && res.statusCode < 500) {
          res.resume();
          resolve(true);
        } else {
          res.resume();
          retry();
        }
      });
      req.on('error', retry);
    };
    const retry = () => {
      if (Date.now() - started > timeout) return reject(new Error(`Timeout waiting for ${url}`));
      setTimeout(tryOnce, 500);
    };
    tryOnce();
  });
}

async function main() {
  const cwd = path.resolve(__dirname, '..');
  // 开发服务器（DEV 环境确保 AuthGuard 绕过生效）
  const preview = spawn('npx', ['vite', '--port', '5174', '--strictPort'], { cwd, stdio: 'inherit' });
  try {
    await waitForUrl('http://127.0.0.1:5174/');
    const url = 'http://127.0.0.1:5174/panel/styleguide?baseline=1';
    const browser = await chromium.launch();
    const page = await browser.newPage({ viewport: { width: 1440, height: 1000 } });
    await page.goto(url, { waitUntil: 'networkidle' });
    // 滚动到 Before/After 区块
    await page.locator('h3:has-text("Before / After")').scrollIntoViewIfNeeded();
    await page.waitForTimeout(500);
    const targetPath = path.resolve(__dirname, '../../docs/ui-baseline/styleguide-before-after.png');
    await page.screenshot({ path: targetPath, fullPage: false });
    await browser.close();
    console.log(`Saved screenshot: ${targetPath}`);
  } finally {
    preview.kill('SIGTERM');
  }
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
