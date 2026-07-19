import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";

/**
 * SS-02 A11y floor: chapter opener for Section H (Accessibility).
 *
 * Establishes WCAG 2.2 AA as the non-negotiable floor for every shipped UI.
 */

export default function A11yFloorSlide() {
  return (
    <SlideLayout
      eyebrow="Chapter H · Accessibility · WCAG 2.2 AA is the floor, not the goal"
      title="Every UI ships accessible by default. Keyboard-reachable, screen-reader-labelled, contrast-passing, motion-safe."
      subtitle="Accessibility is a correctness property, not a polish pass. Broken keyboard nav, missing labels, and low contrast are bugs, tracked and fixed with the same urgency as a null pointer. Automated checks (axe-core in CI) plus a short manual keyboard sweep on every PR that touches UI."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <ActionPanel
          slideId="59-a11y-floor"
          symptom="Three shipped features had to be rebuilt after launch: (1) a modal captured focus but never returned it, so keyboard users landed at the top of the page on every close; (2) a custom dropdown used `div onClick` with no `role`, no `aria-expanded`, and no arrow-key handlers, so screen readers announced 'clickable' and stopped; (3) primary-action buttons used brand purple on brand purple-tint (2.1:1 contrast), unreadable in daylight on a laptop screen. Root cause across all three: accessibility was treated as a polish pass that never happened, not as a shipping requirement."
          rule="WCAG 2.2 AA is the shipping floor for every UI in this codebase. Non-negotiable minimums: (1) every interactive element reachable by Tab in a logical order, with a visible focus ring (never `outline: none` without a replacement); (2) every form control has a programmatic label (`<label htmlFor>`, `aria-label`, or `aria-labelledby`); (3) text contrast >= 4.5:1 (normal), >= 3:1 (large or UI component boundaries); (4) no information conveyed by colour alone (icon plus colour, text plus colour); (5) motion respects `prefers-reduced-motion`; (6) modals trap focus while open and restore focus on close; (7) live regions (`aria-live`) for async status updates; (8) images have `alt` (decorative gets `alt=\"\"`, informative gets a description); (9) heading hierarchy is sequential (no h1 -> h3 skip); (10) hit targets >= 24x24 CSS px (WCAG 2.2 SC 2.5.8). Chapter H expands each into its own hard-rule slide (A11Y-002 through A11Y-010)."
          doThis="Enforce mechanically: (1) `axe-core` runs in CI on every PR via `slides-app/tests/smoke.spec.ts` and every new Playwright test (existing pattern from v5.54.0); zero violations is required to merge; (2) `eslint-plugin-jsx-a11y` at `recommended` with `strict` on `alt-text`, `no-noninteractive-element-interactions`, `click-events-have-key-events`, `label-has-associated-control`; (3) PR template checkbox: 'I tabbed through this feature with no mouse'; unchecked means the reviewer must do it; (4) design tokens enforce contrast at the token level (`--fg-on-primary` computed against `--primary` in CI, fails if < 4.5:1); (5) Storybook a11y addon on every story, `errors` mode; (6) reduced-motion CSS media query in `src/styles/motion.css`, every animation gated. Downstream slides (A11Y-002+) each pick one WCAG SC and turn it into an enforceable rule with an ESLint rule or CI gate."
        />
      </div>
    </SlideLayout>
  );
}
