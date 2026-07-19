import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 WF-003: test pyramid. Unit-heavy, integration-thin, e2e-minimal.
 */

const BEFORE = `// tests/checkout.e2e.spec.ts  --  47 Playwright specs, 22 minutes wall time.
// Every price rule, every currency, every discount coded as a full browser flow.
test("EU VAT applies at 20 percent on physical goods over 50 EUR", async ({ page }) => {
  await page.goto("/");
  await login(page, "eu-user@example.com");
  await addToCart(page, "SKU-PHYS-001", 3);
  await page.click("text=Checkout");
  await expect(page.getByTestId("total")).toHaveText("EUR 216.00");
});
// ... 46 more like this.`;

const AFTER = `// src/pricing/vat.ts  --  pure function, zero I/O.
export function computeVat(input: VatInput): VatBreakdown { ... }

// src/pricing/vat.test.ts  --  200 unit specs, 400 ms wall time.
describe("computeVat", () => {
  test.each(VAT_CASES)("$region $goodsKind $amount -> $expected", ({ input, expected }) => {
    expect(computeVat(input)).toEqual(expected);
  });
});

// src/pricing/checkout.integration.test.ts  --  8 specs, 6 s wall time.
// One per pricing-service seam: DB fixture + real pricing module + mocked payment gateway.
test("checkout persists a VatBreakdown row per line item", async () => { ... });

// tests/checkout.e2e.spec.ts  --  3 Playwright specs, 90 s wall time.
// Happy path, one auth-required path, one payment-declined path. That's it.
test("guest completes a card checkout", async ({ page }) => { ... });`;

export default function TestPyramidSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 61 · Workflow · test pyramid, unit-heavy, integration-thin, e2e-minimal"
      title="Every rule gets tested at the cheapest layer that can catch it. Push tests down, not up."
      subtitle="Target shape: 80 percent unit (pure functions, milliseconds), 15 percent integration (one seam per test, seconds), 5 percent e2e (browser, tens of seconds). Inverted pyramids are the single biggest cause of slow CI, flaky merges, and untested edge cases."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="ts"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. 47 e2e specs to cover pricing math because the math was never extracted from the checkout page. CI takes 22 minutes; developers stop running it locally; a flaky login step (2 percent failure rate) means 60 percent of PRs need a re-run; every new VAT rule requires a browser session; edge cases (negative amounts, zero-rated goods, cross-border) are impossible to exercise because they need a matching test user, product, and shipping address in the fixture DB."
          afterLabel="Good. Extract pure `computeVat` and test it with 200 parameterised unit cases at the millisecond layer. Add 8 integration tests, one per real seam (DB write, pricing module, gateway boundary), running in seconds against fixtures. Keep e2e down to three specs that prove the browser wires the pieces together (happy path, auth, decline). Total CI shrinks from 22 min to under 3 min; flakiness drops to near zero; a new VAT rule is a one-line addition to `VAT_CASES`."
        />
        <ActionPanel
          slideId="60-test-pyramid"
          symptom="Q1 velocity report: 41 percent of engineering time went to 'test infra' (re-running flakes, waiting for CI, debugging Playwright timeouts). Two production incidents in the same quarter shipped because the failing case (a zero-quantity line item, and a negative-discount coupon) could not be added to the e2e fixture DB without breaking six other specs; nobody wrote a unit test because the logic lived inside a 400-line React component. Root cause: an inverted pyramid, most tests at the most expensive layer, the cheap layer nearly empty because the code was not shaped for it."
          rule="Every test lives at the cheapest layer that can catch the class of bug it targets. Non-negotiable targets: (1) 80 percent unit, tests of pure functions and small modules with zero I/O, milliseconds each, run on every save; (2) 15 percent integration, one real seam per test (DB + module, module + gateway, hook + reducer), seconds each, run on every PR; (3) 5 percent e2e, browser-level flows that prove wiring, tens of seconds each, kept to happy path plus one representative failure per critical journey. Coverage of business rules belongs at the unit layer, not the e2e layer. If a rule cannot be unit-tested, extract it until it can (pairs with REACT-007 file size, REACT-012 name-or-split). E2e is for wiring proofs and for user-visible regressions; it is never the primary defence."
          doThis="Enforce mechanically: (1) CI job matrix with three separate steps `test:unit` (must finish under 60 s), `test:integration` (under 5 min), `test:e2e` (under 5 min), each with independent pass/fail; (2) coverage gate `vitest --coverage` requires 90 percent line coverage on `src/` excluding UI files, with 100 percent on `src/pricing`, `src/auth`, `src/schema`; (3) test-shape report script `scripts/report-test-shape.mjs` counts specs per layer and fails CI if e2e exceeds 10 percent of total; (4) PR template checkbox (see WF-004) 'tests added at the correct pyramid layer'; (5) code review reflex: any new Playwright spec for pure logic is rejected with 'push this down to a unit test'; (6) extract-first workflow: when a bug is caught in e2e, the fix is a new unit test on the extracted function AND a note in the retro (RCA-002) if it happened twice. Rule of thumb: if you can compute it with a pure function, you must test it as one."
        />
      </div>
    </SlideLayout>
  );
}
