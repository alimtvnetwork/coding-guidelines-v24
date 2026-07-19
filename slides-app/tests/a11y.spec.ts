/**
 * Accessibility scan for every slide in the standalone deck.
 *
 * Runs axe-core (WCAG 2.1 A/AA) against each hash-routed slide and fails on
 * any violation. Landmark/region rules are disabled: the deck root is a
 * single fullscreen `<div class="slide-content">` by design (per
 * spec-slides/03-layout-and-tokens.md), not a landmarked page.
 *
 * Run locally: bunx playwright test slides-app/tests/a11y.spec.ts
 * CI:          .github/workflows/slides-smoke.yml
 */
import { test, expect } from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";
import { pathToFileURL } from "node:url";
import { existsSync } from "node:fs";
import { resolve, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const DIST_INDEX = resolve(HERE, "..", "dist", "index.html");
const DIST_URL = pathToFileURL(DIST_INDEX).href;
const SLIDE_COUNT = 16;

const DISABLED_RULES = ["region", "landmark-one-main", "page-has-heading-one"];

test.beforeAll(() => {
  if (!existsSync(DIST_INDEX)) {
    throw new Error(
      `slides-app dist missing at ${DIST_INDEX} - run \`cd slides-app && bun run build\` first.`,
    );
  }
});

test("every slide passes axe-core WCAG 2.1 A/AA", async ({ page }) => {
  await page.goto(`${DIST_URL}#/0`);
  await expect(page.locator(".slide-content").first()).toBeVisible({ timeout: 10_000 });

  const failures: string[] = [];

  for (let i = 0; i < SLIDE_COUNT; i++) {
    await page.evaluate((n) => {
      window.location.hash = `/${n}`;
    }, i);
    await expect(page.locator(".slide-content")).toHaveCount(1, { timeout: 5_000 });

    const results = await new AxeBuilder({ page })
      .withTags(["wcag2a", "wcag2aa", "wcag21a", "wcag21aa"])
      .disableRules(DISABLED_RULES)
      .analyze();

    if (results.violations.length === 0) continue;
    const summary = results.violations
      .map((v) => `  - [${v.id}] ${v.help} (${v.nodes.length} node(s))`)
      .join("\n");
    failures.push(`slide #${i}:\n${summary}`);
  }

  expect(failures, `axe-core violations:\n${failures.join("\n\n")}`).toEqual([]);
});
