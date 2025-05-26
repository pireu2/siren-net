import path from "path"
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(),tailwindcss(),],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  server: {
    proxy: {
      "/auth": {
        target: "http://localhost:8080", 
        changeOrigin: true,
        secure: false,
      },
      "/protected": {
        target: "http://localhost:8080", 
        changeOrigin: true,
        secure: false,
      },
       "/client": {
        target: "http://localhost:8080", 
        changeOrigin: true,
        secure: false,
      },
      "/agents": {
        target: "http://localhost:8080", 
        changeOrigin: true,
        secure: false,
      },
      "/messages": {
        target: "http://localhost:8080", 
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
