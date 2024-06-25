import { defineConfig } from "vite";
import viteReact from '@vitejs/plugin-react'
import { TanStackRouterVite } from '@tanstack/router-plugin/vite'
import path from "path"

export default defineConfig({
  plugins: [
      TanStackRouterVite({
          routesDirectory: './frontend/src/routes',
          generatedRouteTree: './frontend/src/routeTree.gen.ts',
          enableRouteGeneration: true,
      }),
      viteReact()
  ],
  server: {
        watch: {
            usePolling: true
        }
    },
  resolve: {
      alias: {
          '@': path.resolve(__dirname, './frontend/src'),
      }
  },
  root: "./frontend",
  publicDir: "./frontend/public"
});
