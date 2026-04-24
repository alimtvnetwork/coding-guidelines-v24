import { test, expect } from "../../playwright-fixture";

// B2 — Smoke test for the landing page. Verifies the hero renders,
// the install one-liners are visible, and no console errors fire.

test.describe("Landing page", () => {
  test("renders hero and install commands", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (e) => errors.push(e.message));
    page.on("console", (m) => {
      if (m.type() === "error") errors.push(m.text());
    });

    await page.goto("/");

    // Hero heading is visible
    await expect(page.locator("h1").first()).toBeVisible();

    // Both install commands are rendered somewhere on the page
    await expect(page.getByText(/install\.ps1/i).first()).toBeVisible();
    await expect(page.getByText(/install\.sh/i).first()).toBeVisible();

    // No JS runtime errors
    expect(errors, `console errors: ${errors.join(" | ")}`).toEqual([]);
  });
});
