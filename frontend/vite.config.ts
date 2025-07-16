
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
import vue from '@vitejs/plugin-vue'
import * as path from 'path';
import { fileURLToPath } from 'url';

// Soporte para __dirname en módulos ES
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);


// https://vite.dev/config/
export default defineConfig({
  plugins:
    [
      tailwindcss(),
      vue()
    ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  }
})
