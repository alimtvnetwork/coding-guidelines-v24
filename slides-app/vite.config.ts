import { defineConfig, type Plugin } from "vite";
import react from "@vitejs/plugin-react";
import path from "node:path";

// Rewrites the built index.html so it loads from file:// — strips
// `type="module"` and `crossorigin` (both blocked by browser CORS over
// file://) and adds a defer attribute so the IIFE runs after DOM parse.
function fileProtocolHtmlPlugin(): Plugin {
  return {
    name: "file-protocol-html",
    apply: "build",
    transformIndexHtml: {
      order: "post",
      handler(html) {
        return html
          .replace(/<script\s+type="module"\s+crossorigin\s+/g, "<script defer ")
          .replace(/<script\s+type="module"\s+/g, "<script defer ")
          .replace(/\s+crossorigin(?=[\s>])/g, "");
      },
    },
  };
}

export default defineConfig({
  base: "./",
  plugins: [react(), fileProtocolHtmlPlugin()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
  build: {
    outDir: "dist",
    assetsInlineLimit: 0,
    // IIFE format so the bundle runs from file:// (ES modules are blocked
    // by browser CORS over file://, leaving a blank page).
    rollupOptions: {
      input: path.resolve(__dirname, "index.html"),
      output: {
        manualChunks: undefined,
        format: "iife",
        inlineDynamicImports: true,
        entryFileNames: "assets/index-[hash].js",
      },
    },
    modulePreload: { polyfill: false },
  },
});
