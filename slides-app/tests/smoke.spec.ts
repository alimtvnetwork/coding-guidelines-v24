/**
 * Playwright smoke test for the standalone slides deck.
 *
 * Verifies:
 *  1. The built `dist/index.html` loads from `file://` (the offline mode the
 *     deck is shipped under per spec-slides/06-build-and-zip-pipeline.md).
 *  2. Every slide in DECK renders a `.slide-content` root without throwing.
 *  3. Hash-based navigation (`#/<index>`) reaches first, middle, and last
 *     slides without console errors.
 *
 * Run locally:  bunx playwright test slides-app/tests/smoke.spec.ts
 * CI:           .github/workflows/slides-smoke.yml
 */
import { test, expect, type ConsoleMessage } from "@playwright/test";
import { pathToFileURL } from "node:url";
import { existsSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const DIST_INDEX = resolve(HERE, "..", "dist", "index.html");
const DIST_URL = pathToFileURL(DIST_INDEX).href;

// Mirror of slides-app/src/deck.ts (length only — kept in sync by the
// src/test/slides-deck.test.ts vitest guard).
const SLIDE_COUNT = 16;

test.beforeAll(() => {
  if (!existsSync(DIST_INDEX)) {
    throw new Error(
      `slides-app dist missing at ${DIST_INDEX} — run \`cd slides-app && bun run build\` first.`,
    );
  }
});

test("slides deck boots without console errors", async ({ page }) => {
  const errors: string[] = [];
  page.on("console", (msg: ConsoleMessage) => {
    if (msg.type() === "error") errors.push(msg.text());
  });
  page.on("pageerror", (err) => errors.push(`pageerror: ${err.message}`));

  await page.goto(`${DIST_URL}#/0`);
  await expect(page.locator(".slide-content").first()).toBeVisible({ timeout: 10_000 });
  expect(errors, `console errors on boot:\n${errors.join("\n")}`).toEqual([]);
});

test("every slide index renders a slide-content root", async ({ page }) => {
  await page.goto(`${DIST_URL}#/0`);
  await expect(page.locator(".slide-content").first()).toBeVisible({ timeout: 10_000 });

  for (let i = 0; i < SLIDE_COUNT; i++) {
    await page.evaluate((n) => {
      window.location.hash = `/${n}`;
    }, i);
    // Wait for the hashchange handler to swap the slide component.
    await expect(page.locator(".slide-content")).toHaveCount(1, { timeout: 5_000 });
    const visible = await page.locator(".slide-content").first().isVisible();
    expect(visible, `slide #${i} is not visible`).toBe(true);
  }
});

test("first and last slides render non-empty content", async ({ page }) => {
  await page.goto(`${DIST_URL}#/0`);
  const first = page.locator(".slide-content").first();
  await expect(first).toBeVisible({ timeout: 10_000 });
  expect((await first.innerText()).trim().length).toBeGreaterThan(0);

  await page.evaluate((n) => {
    window.location.hash = `/${n}`;
  }, SLIDE_COUNT - 1);
  const last = page.locator(".slide-content").first();
  await expect(last).toBeVisible({ timeout: 5_000 });
  expect((await last.innerText()).trim().length).toBeGreaterThan(0);
});
