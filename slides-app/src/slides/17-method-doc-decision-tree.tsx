import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";

/**
 * SS-02 task 20: Method-documentation decision tree.
 *
 * Encodes the 5-step checklist from spec/17/31 §"Method Documentation".
 * Default answer is NO DOC; a doc comment is the exceptional outcome that
 * survives only when refactor, split, and signature-restate checks all fail.
 */

type Outcome = "skip" | "delete" | "keep" | "oneliner";

interface Node {
  step: string;
  question: string;
  yes: { label: string; outcome: Outcome };
  no: string;
}

const OUTCOME_META: Record<Outcome, { label: string; accent: string; detail: string }> = {
  skip: {
    label: "Skip the doc",
    accent: "primary",
    detail: "Rename or split the method. The identifier becomes the doc.",
  },
  delete: {
    label: "Delete the doc",
    accent: "destructive",
    detail: "Restating the signature is review-blocking noise.",
  },
  keep: {
    label: "Keep · 1-2 lines",
    accent: "accent",
    detail: "WHY / cited source / short example that clarifies the contract.",
  },
  oneliner: {
    label: "One-liner on exported API",
    accent: "accent",
    detail: "Only when godoc / TypeDoc / phpDocumentor is wired in CI.",
  },
};

const NODES: readonly Node[] = [
  {
    step: "Q1",
    question: "Can I rename so the doc becomes redundant?",
    yes: { label: "Yes → rename", outcome: "skip" },
    no: "No",
  },
  {
    step: "Q2",
    question: "Can I split so each piece is trivially named?",
    yes: { label: "Yes → split", outcome: "skip" },
    no: "No",
  },
  {
    step: "Q3",
    question: "Does the draft doc restate the signature?",
    yes: { label: "Yes", outcome: "delete" },
    no: "No",
  },
  {
    step: "Q4",
    question: "Does it explain WHY, cite a source, or give a runnable example?",
    yes: { label: "Yes", outcome: "keep" },
    no: "No",
  },
  {
    step: "Q5",
    question: "Is automated doc generation wired in CI?",
    yes: { label: "Yes", outcome: "oneliner" },
    no: "No → skip",
  },
];

export default function MethodDocDecisionTreeSlide() {
  return (
    <SlideLayout
      eyebrow="Principle · Method documentation"
      title="Decision tree: when to write a doc comment"
      subtitle="Default answer is NO DOC. Docs survive only when rename, split, and signature-restate checks all fail."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <ActionPanel
          slideId="17-method-doc-decision-tree"
          symptom="Every method ships with a prose doc comment; half restate the signature, a quarter drift stale, and reviewers stop reading them."
          rule="Simple methods carry no doc. Write a comment only when refactor, split, and signature-restate checks all fail, and the surviving text is WHY / cited / example."
          doThis="On the next PR that adds a doc comment, walk Q1-Q5 in order. If Q1 or Q2 answers Yes, refactor instead. If Q3 answers Yes, delete."
        />
        <div style={{ display: "grid", gridTemplateColumns: "repeat(5, 1fr)", gap: 12 }}>
          {NODES.map((node) => {
            const outcome = OUTCOME_META[node.yes.outcome];

            return (
              <div
                key={node.step}
                style={{
                  padding: "14px 14px",
                  borderRadius: 12,
                  background: "hsl(var(--bg-raised))",
                  border: "1px solid hsl(var(--border))",
                  display: "flex",
                  flexDirection: "column",
                  gap: 10,
                  minWidth: 0,
                }}
              >
                <div
                  style={{
                    fontSize: 11,
                    fontWeight: 700,
                    letterSpacing: "0.14em",
                    textTransform: "uppercase",
                    color: "hsl(var(--fg-muted))",
                  }}
                >
                  {node.step}
                </div>
                <div style={{ fontSize: 14, lineHeight: 1.35, color: "hsl(var(--fg))", minHeight: 68 }}>
                  {node.question}
                </div>
                <div
                  style={{
                    padding: "8px 10px",
                    borderRadius: 8,
                    background: `hsl(var(--${outcome.accent}) / 0.12)`,
                    border: `1px solid hsl(var(--${outcome.accent}) / 0.5)`,
                    display: "flex",
                    flexDirection: "column",
                    gap: 4,
                  }}
                >
                  <div
                    style={{
                      fontSize: 11,
                      fontWeight: 700,
                      letterSpacing: "0.08em",
                      textTransform: "uppercase",
                      color: `hsl(var(--${outcome.accent}))`,
                    }}
                  >
                    {node.yes.label} · {outcome.label}
                  </div>
                  <div style={{ fontSize: 12, lineHeight: 1.35, color: "hsl(var(--fg))" }}>
                    {outcome.detail}
                  </div>
                </div>
                <div style={{ fontSize: 11, color: "hsl(var(--fg-muted))" }}>
                  {node.no === "No" ? "No → next question" : node.no}
                </div>
              </div>
            );
          })}
        </div>
        <div
          style={{
            padding: "12px 18px",
            borderRadius: 10,
            background: "hsl(var(--bg-raised))",
            border: "1px dashed hsl(var(--border))",
            fontSize: 13,
            color: "hsl(var(--fg-muted))",
          }}
        >
          Reference · spec/17-consolidated-guidelines/31 §Method Documentation (checklist Q1-Q5). Same rules for Go, TS, PHP, Rust, C#, PowerShell, Python; only the comment syntax changes.
        </div>
      </div>
    </SlideLayout>
  );
}
