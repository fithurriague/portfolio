import { globSync } from 'glob';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

export default {
  root: './', // specify the project root if needed
  build: {
    outDir: 'dist',
    rollupOptions: {
      input: Object.fromEntries(
        globSync('resources/**/*.js').map((file) => [
          path.relative(
            'resources',
            file.slice(0, file.length - path.extname(file).length)
          ),
          fileURLToPath(new URL(file, import.meta.url)),
        ])
      ),
      output: {
        entryFileNames: '[name].js',
        assetFileNames: (assetInfo) => {
          // Process CSS files correctly
          if (assetInfo.name && assetInfo.name.endsWith('.css')) {
            return `css/[name][extname]`;
          }
          // Default for other assets
          return `assets/[name][extname]`;
        },
      },
    },
  },
};
