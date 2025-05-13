import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueDevTools from "vite-plugin-vue-devtools";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    port: 5515,
    proxy: {
      "/api": {
        target: "http://localhost:3000", // Change to your Go backend port
        changeOrigin: true,
        secure: false,
        // Allow cookies to be set by backend
        cookieDomainRewrite: "localhost",
      },
    },
  },
});
