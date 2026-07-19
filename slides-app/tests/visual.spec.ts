/**
 * Playwright visual regression for the standalone slides deck.
 *
 * Baseline PNGs live under `visual.spec.ts-snapshots/` and are the
 * source of truth. CI fails when a slide renders differently from its
 * baseline beyond the `maxDiffPixelRatio` in `playwright.config.ts`.
 *
 * Regenerating baselines (intentional visual change):
 *   bunx playwright test slides-app/tests/visual.spec.ts --update-snapshots
 *   Or in CI: run `.github/workflows/slides-visual.yml` with the
 *   `update_baselines=true` input, then download the produced artifact
 *   and commit its `visual.spec.ts-snapshots/` folder.
 *
 * Animations (framer-motion) are silenced via `MotionGlobalConfig` and
 * a `prefers-reduced-motion` media override so the snapshot is stable.
 */
import { test, expect, type Page } from "@playwright/test";
import { pathToFileURL } from "node:url";
import { existsSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";
import { DECK } from "../src/deck/registry";

const HERE = dirname(fileURLToPath(import.meta.url));
const DIST_INDEX = resolve(HERE, "..", "dist", "index.html");
const DIST_URL = pathToFileURL(DIST_INDEX).href;

// Derive from the canonical deck so new slides are auto-covered and stale
// hardcoded counts (previously 17 while the deck held 70) cannot recur.
const SLIDE_COUNT = DECK.length;
const SETTLE_MS = 1200;

test.beforeAll(() => {
  if (!existsSync(DIST_INDEX)) {
    throw new Error(
      `slides-app dist missing at ${DIST_INDEX}. Run \`cd slides-app && bun run build\` first.`,
    );
  }
});

async function silenceAnimations(page: Page): Promise<void> {
  await page.emulateMedia({ reducedMotion: "reduce" });
  await page.addStyleTag({
    content: `
      *, *::before, *::after {
        animation-duration: 0s !important;
        animation-delay: 0s !important;
        transition-duration: 0s !important;
        transition-delay: 0s !important;
      }
    `,
  });
}

async function gotoSlide(page: Page, index: number): Promise<void> {
  await page.evaluate((n) => {
    window.location.hash = `/${n}`;
  }, index);
  await expect(page.locator(".slide-content")).toHaveCount(1, { timeout: 5_000 });
  await page.waitForTimeout(SETTLE_MS);
}

test.describe("visual regression: every slide matches baseline", () => {
  for (let i = 0; i < SLIDE_COUNT; i++) {
    test(`slide ${String(i).padStart(2, "0")}`, async ({ page }) => {
      await page.goto(`${DIST_URL}#/${i}`);
      await silenceAnimations(page);
      await expect(page.locator(".slide-content").first()).toBeVisible({
        timeout: 10_000,
      });
      await gotoSlide(page, i);
      await expect(page.locator(".slide-content").first()).toHaveScreenshot(
        `slide-${String(i).padStart(2, "0")}.png`,
      );
    });
  }
});
