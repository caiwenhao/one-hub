// https://github.com/vitejs/vite/discussions/3448
import path from 'path';
import { fileURLToPath } from 'node:url';
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import jsconfigPaths from 'vite-jsconfig-paths';

// ----------------------------------------------------------------------

const resolveFrom = (p) => fileURLToPath(new URL(p, import.meta.url));

export default defineConfig({
  plugins: [react(), jsconfigPaths()],
  // https://github.com/jpuri/react-draft-wysiwyg/issues/1317
  //   define: {
  //     global: 'window'
  //   },
  resolve: {
    alias: {
      '@': resolveFrom('./src'),
      'src': resolveFrom('./src'),
      '~': resolveFrom('./node_modules')
    }
  },
  build: {
    outDir: 'build',
    assetsDir: 'assets',
    rollupOptions: {
      output: {
        manualChunks: undefined
      }
    }
  },
  server: {
    // this ensures that the browser opens upon server start
    open: true,
    // this sets a default port to 3000
    host: true,
    port: 3010,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:3000', // 设置代理的目标服务器
        changeOrigin: true
      }
    }
  },
  preview: {
    // this ensures that the browser opens upon preview start
    open: true,
    // this sets a default port to 3000
    port: 3010
  }
});
