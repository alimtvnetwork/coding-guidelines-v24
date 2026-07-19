import { defineConfig } from "@playwright/test";

/**
 * Root Playwright config. Used by CI smoke, a11y, and visual-regression
 * tests that load the pre-built `slides-app/dist/index.html` via `file://`.
 *
 * Visual snapshots are pinned to Linux Chromium so the CI baselines match
 * the runner (`ubuntu-latest`). Local runs on macOS/Windows will diff
 * against Linux baselines and are expected to fail; use CI for baselines.
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
    deviceScaleFactor: 1,
  },
  expect: {
    // Small tolerance absorbs sub-pixel font-rendering jitter on Chromium
    // while still catching real layout, color, and content drift.
    toHaveScreenshot: {
      maxDiffPixelRatio: 0.01,
      animations: "disabled",
      caret: "hide",
      scale: "css",
    },
  },
});
