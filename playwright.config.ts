import { defineConfig } from "@playwright/test";

/**
 * Root Playwright config — used by CI smoke tests that don't need the Vite
 * dev server (the slides deck loads from a pre-built `dist/index.html` via
 * `file://`). Keep this minimal; per-suite configs can override.
 */
export default defineConfig({
  testDir: "./slides-app/tests",
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 1 : 0,
  reporter: process.env.CI ? "github" : "list",
  use: {
    headless: true,
    viewport: { width: 1280, height: 800 },
    trace: "retain-on-failure",
  },
});
