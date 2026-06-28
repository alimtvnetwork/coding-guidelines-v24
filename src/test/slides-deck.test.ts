import { describe, it, expect } from "vitest";
import { readFileSync, readdirSync } from "node:fs";
import { resolve } from "node:path";

// Smoke test for slides-app deck integrity. Static (no Playwright runtime
// required) — verifies every slide module is referenced and exports a
// default. Catches the common regressions: orphaned slide file, missing
// import in deck.ts, DECK length drift vs file count.

const SLIDES_DIR = resolve(__dirname, "../../slides-app/src/slides");
const DECK_FILE = resolve(__dirname, "../../slides-app/src/deck.ts");

describe("slides-app deck", () => {
  const slideFiles = readdirSync(SLIDES_DIR).filter((f) => f.endsWith(".tsx"));
  const deckSrc = readFileSync(DECK_FILE, "utf8");

  it("has at least 12 slides registered", () => {
    expect(slideFiles.length).toBeGreaterThanOrEqual(12);
  });

  it("imports every slide file in deck.ts", () => {
    for (const f of slideFiles) {
      const base = f.replace(/\.tsx$/, "");
      expect(deckSrc, `deck.ts missing import for ${base}`).toContain(`/slides/${base}`);
    }
  });

  it("every slide file exports a default", () => {
    for (const f of slideFiles) {
      const src = readFileSync(resolve(SLIDES_DIR, f), "utf8");
      expect(src, `${f} has no default export`).toMatch(/export\s+default\s+/);
    }
  });

  it("DECK entry count matches imported slide count", () => {
    const importCount = (deckSrc.match(/from\s+"\.\/slides\//g) ?? []).length;
    const entryCount = (deckSrc.match(/component:\s*\w+/g) ?? []).length;
    expect(entryCount).toBe(importCount);
  });
});
