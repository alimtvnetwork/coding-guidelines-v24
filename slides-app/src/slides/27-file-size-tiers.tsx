import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";

/**
 * SS-02 task 29: File-size tier caps.
 *
 * Source: 02-spec/17/31 line 30. Any file 300 lines max, any React .tsx 100
 * lines max, any class or struct 120 lines max. Waiver only via inline
 * lint-allow comment stating the reason and explicit max.
 */

type Tier = {
  cap: string;
  scope: string;
  trigger: string;
  extractTo: string;
};

const TIERS: Tier[] = [
  {
    cap: "100 lines",
    scope: "React .tsx component file",
    trigger: "JSX plus hooks plus handlers plus derived state in one file",
    extractTo: "child component, custom hook, or pure helper module",
  },
  {
    cap: "120 lines",
    scope: "Class or struct definition",
    trigger: "Too many responsibilities on one type",
    extractTo: "collaborator class, value object, or free function",
  },
  {
    cap: "300 lines",
    scope: "Any file (fallback cap)",
    trigger: "Multiple top-level concerns cohabiting one module",
    extractTo: "split by concern into sibling modules with clear names",
  },
];

function TierCard({ tier }: { tier: Tier }) {
  return (
    <div
      style={{
        background: "hsl(var(--card))",
        border: "2px solid hsl(var(--slide-accent))",
        borderRadius: 20,
        padding: "28px 32px",
        display: "flex",
        flexDirection: "column",
        gap: 12,
        minHeight: 320,
      }}
    >
      <div className="slide-kicker" style={{ color: "hsl(var(--slide-accent))" }}>
        {tier.scope}
      </div>
      <div className="slide-title" style={{ fontSize: 72, lineHeight: 1 }}>
        {tier.cap}
      </div>
      <div className="slide-body" style={{ opacity: 0.85 }}>
        <strong>Trigger:</strong> {tier.trigger}
      </div>
      <div className="slide-body" style={{ opacity: 0.85 }}>
        <strong>Extract to:</strong> {tier.extractTo}
      </div>
    </div>
  );
}

export default function FileSizeTiersSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 29 · File-size tiers"
      title="Three caps, not one. Pick the tightest that fits."
      subtitle="300 lines is the fallback ceiling, not a target. React components cap at 100, classes and structs at 120. When a file passes its cap, extract; do not add a waiver comment to buy space."
    >
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "repeat(3, 1fr)",
          gap: 22,
          marginTop: 16,
        }}
      >
        {TIERS.map((t) => (
          <TierCard key={t.cap} tier={t} />
        ))}
      </div>
      <ActionPanel
        slideId="27-file-size-tiers"
        symptom="A .tsx file grows to 340 lines with 6 hooks, 4 handlers, and nested subcomponents. Reviewers scroll past the diff because it no longer fits one screen; regressions land because nobody spots the state coupling."
        rule="Any file 300 lines max. Any React .tsx 100 lines max. Any class or struct 120 lines max. The tightest applicable cap wins. Waiver only via inline `lint-allow: file-length reason='...' max=N` and only for generated fixtures or one-shot data tables. Per 02-spec/17/31 line 30."
        doThis="When a .tsx passes 100 lines, extract child components into siblings, move hooks into `useXxx.ts` under a colocated `hooks/` folder, and move pure helpers into `.ts` modules. When a class passes 120, extract collaborators; do not inline more methods. When any file passes 300, split by concern before merging."
      />
    </SlideLayout>
  );
}
