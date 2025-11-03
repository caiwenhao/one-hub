// 对比度校验脚本：读取 tokens.css，生成常见文本/背景/边框组合的 AA 报告
import fs from 'node:fs';
import path from 'node:path';

const root = path.resolve(process.cwd(), '..');
const tokensPath = path.resolve(process.cwd(), 'src/styles/tokens.css');
const reportPath = path.resolve(root, 'docs/ui-baseline/contrast-report.md');

function parseVars(css) {
  const map = {};
  const re = /--([a-z0-9-]+):\s*#([0-9a-fA-F]{6});/g;
  let m;
  while ((m = re.exec(css))) {
    map[`--${m[1]}`] = `#${m[2]}`;
  }
  return map;
}

function hexToRgb(hex) {
  const n = parseInt(hex.slice(1), 16);
  return { r: (n >> 16) & 255, g: (n >> 8) & 255, b: n & 255 };
}

function relLum({ r, g, b }) {
  const srgb = [r, g, b].map((v) => v / 255).map((v) => (v <= 0.03928 ? v / 12.92 : Math.pow((v + 0.055) / 1.055, 2.4)));
  return 0.2126 * srgb[0] + 0.7152 * srgb[1] + 0.0722 * srgb[2];
}

function contrastRatio(fg, bg) {
  const L1 = relLum(hexToRgb(fg));
  const L2 = relLum(hexToRgb(bg));
  const ratio = (Math.max(L1, L2) + 0.05) / (Math.min(L1, L2) + 0.05);
  return Math.round(ratio * 100) / 100;
}

function checkAA(ratio, large = false) {
  return large ? ratio >= 3 : ratio >= 4.5;
}

async function main() {
  const css = fs.readFileSync(tokensPath, 'utf-8');
  const vars = parseVars(css);
  const pairs = [
    ['--color-text-primary', '--color-bg-card'],
    ['--color-text-secondary', '--color-bg-card'],
    ['--color-text-primary', '--color-bg-page'],
    ['--color-brand-primary', '--color-bg-card'],
    ['--color-brand-secondary', '--color-bg-card'],
    ['--color-semantic-error', '--color-bg-card'],
    ['--color-semantic-warning', '--color-bg-card'],
    ['--color-semantic-success', '--color-bg-card']
  ];
  let md = `# 对比度报告 (WCAG AA)\n\n| 前景 | 背景 | 比值 | AA 正文 | AA 大字 |\n|---|---|---:|:---:|:---:|\n`;
  for (const [fgVar, bgVar] of pairs) {
    const fg = vars[fgVar];
    const bg = vars[bgVar];
    if (!fg || !bg) continue;
    const ratio = contrastRatio(fg, bg);
    const aa = checkAA(ratio);
    const aaLarge = checkAA(ratio, true);
    md += `| ${fgVar} | ${bgVar} | ${ratio} | ${aa ? '✅' : '❌'} | ${aaLarge ? '✅' : '❌'} |\n`;
  }
  fs.writeFileSync(reportPath, md, 'utf-8');
  console.log('生成报告:', reportPath);
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
