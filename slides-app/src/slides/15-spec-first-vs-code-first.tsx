import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";

/**
 * SS-02 task 18: Spec-first vs code-first (before/after PR contrast).
 *
 * Operationalizes non-negotiable "write the spec before you touch the code"
 * from `13-must-follow` and `10-spec-first-workflow`. Shows two PR payloads
 * side by side so reviewers can pattern-match rejection at a glance.
 */

interface PrPayload {
  kicker: string;
  files: readonly string[];
  outcome: string;
}

const CODE_FIRST: PrPayload = {
  kicker: "PR #482 · code-first (rejected)",
  files: [
    "src/services/payments/refund.ts   +48 / -6",
    "src/hooks/useRefund.ts             +22 / -3",
    "src/pages/OrderDetail.tsx         +11 / -0",
  ],
  outcome:
    "Reviewer has no reference for correct behavior. Merges block on 'what should this do?' threads. Future AI agents cannot diff intent against implementation.",
};

const SPEC_FIRST: PrPayload = {
  kicker: "PR #483 · spec-first (accepted)",
  files: [
    "02-spec/12-payments/07-refund-flow.md   +64 / -0",
    "03-issues/482-partial-refund.md      +18 / -0",
    "src/services/payments/refund.ts      +48 / -6",
    "src/hooks/useRefund.ts               +22 / -3",
    "src/pages/OrderDetail.tsx            +11 / -0",
    "tests/refund.spec.ts                 +34 / -0",
  ],
  outcome:
    "Spec pins the contract. Reviewer diffs code against spec, not against memory. Next AI agent reads the same source of truth humans do.",
};

function PayloadCard({
  payload,
  accent,
}: {
  payload: PrPayload;
  accent: "destructive" | "accent";
}) {
  return (
    <div
      style={{
        padding: "20px 22px",
        borderRadius: 14,
        background: "hsl(var(--bg-raised))",
        border: `1px solid hsl(var(--${accent}) / 0.4)`,
        borderTop: `4px solid hsl(var(--${accent}))`,
        display: "flex",
        flexDirection: "column",
        gap: 12,
      }}
    >
      <div
        style={{
          fontSize: 14,
          fontWeight: 700,
          letterSpacing: "0.12em",
          textTransform: "uppercase",
          color: `hsl(var(--${accent}))`,
        }}
      >
        {payload.kicker}
      </div>
      <ul
        style={{
          margin: 0,
          padding: 0,
          listStyle: "none",
          display: "flex",
          flexDirection: "column",
          gap: 6,
          fontFamily: "var(--font-mono)",
          fontSize: 14,
          color: "hsl(var(--fg))",
        }}
      >
        {payload.files.map((line) => (
          <li key={line} style={{ whiteSpace: "pre" }}>
            {line}
          </li>
        ))}
      </ul>
      <div style={{ fontSize: 14, lineHeight: 1.45, color: "hsl(var(--fg-muted))" }}>
        {payload.outcome}
      </div>
    </div>
  );
}

export default function SpecFirstVsCodeFirstSlide() {
  return (
    <SlideLayout
      eyebrow="Workflow · Spec-first"
      title="Two PRs, same code diff, opposite fate"
      subtitle="The spec entry is the reviewer's yardstick. Without it, every merge negotiates behavior from scratch."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 22, marginTop: 10 }}>
        <ActionPanel
          slideId="15-spec-first-vs-code-first"
          symptom="Pull requests that touch only code force reviewers to guess intent and future AI agents to infer contracts from implementation."
          rule="Every behavior change lands with a spec file (or updates one) and an issue entry in the same PR. Code without spec is auto-reject."
          doThis="Open the spec/issue file first, write the contract, then let the code fall out of it. Paste the spec path into the PR description."
        />
        <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 24 }}>
          <PayloadCard payload={CODE_FIRST} accent="destructive" />
          <PayloadCard payload={SPEC_FIRST} accent="accent" />
        </div>
      </div>
    </SlideLayout>
  );
}
