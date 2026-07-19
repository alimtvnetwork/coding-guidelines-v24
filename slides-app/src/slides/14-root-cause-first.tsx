import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";

/**
 * SS-02 task 17: Root-cause-first workflow slide.
 *
 * Operationalizes non-negotiable #2 from `13-must-follow`: a one-sentence
 * root cause must precede any fix. Contrasts the symptom-patch loop against
 * the RCA-first loop so readers see the payoff before the discipline cost.
 */

interface Step {
  index: string;
  label: string;
  detail: string;
}

const RCA_LOOP: readonly Step[] = [
  { index: "01", label: "Reproduce", detail: "Trigger the failing signal on demand: exact input, exact log line, exact stack." },
  { index: "02", label: "Trace end-to-end", detail: "Walk the call path from entry to failure. Name the file, function, and line." },
  { index: "03", label: "State the cause", detail: "Write ONE sentence: 'X happens because Y at file:line does Z.' No 'should' or 'maybe.'" },
  { index: "04", label: "Minimum fix", detail: "Change only what the sentence names. Delete every other line you touched." },
  { index: "05", label: "Prove it", detail: "Show the log line flipping red to green. If no log exists, add one and re-run." },
];

const SYMPTOM_LOOP = [
  "Wrap the failing call in try/catch and swallow the error.",
  "Add a fallback value so the UI stops throwing.",
  "Force a re-render or extra useEffect to 'refresh' state.",
  "Ship. Watch the same bug reopen under a new stack trace next week.",
];

export default function RootCauseFirstSlide() {
  return (
    <SlideLayout
      eyebrow="Workflow · Root cause first"
      title="One sentence before one line of code"
      subtitle="If you cannot write the cause in a single sentence, you have not read enough yet. Keep reading."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 22, marginTop: 10 }}>
        <ActionPanel
          slideId="14-root-cause-first"
          symptom="Fixes that mask the failing signal (try/catch, fallback values, extra effects) so the same bug returns next week under a fresh stack trace."
          rule="Reproduce, trace end-to-end, then write the root cause in one sentence naming file, function, and line. Only after that sentence exists, write the fix."
          doThis="On your next bug, paste the one-sentence RCA into the PR description before you push code. Reviewers reject PRs without it."
        />
        <div style={{ display: "grid", gridTemplateColumns: "1.35fr 1fr", gap: 24 }}>
          <div
            style={{
              padding: "20px 22px",
              borderRadius: 14,
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--accent) / 0.4)",
              borderTop: "4px solid hsl(var(--accent))",
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
                color: "hsl(var(--accent))",
              }}
            >
              RCA-first loop
            </div>
            <ol style={{ margin: 0, padding: 0, listStyle: "none", display: "flex", flexDirection: "column", gap: 10 }}>
              {RCA_LOOP.map((step) => (
                <li key={step.index} style={{ display: "flex", gap: 12, alignItems: "flex-start" }}>
                  <span
                    style={{
                      fontFamily: "var(--font-mono)",
                      fontSize: 16,
                      fontWeight: 700,
                      color: "hsl(var(--accent))",
                      minWidth: 28,
                    }}
                  >
                    {step.index}
                  </span>
                  <span style={{ display: "flex", flexDirection: "column", gap: 2 }}>
                    <span style={{ fontSize: 16, fontWeight: 700, color: "hsl(var(--fg))" }}>{step.label}</span>
                    <span style={{ fontSize: 14, lineHeight: 1.4, color: "hsl(var(--fg-muted))" }}>{step.detail}</span>
                  </span>
                </li>
              ))}
            </ol>
          </div>
          <div
            style={{
              padding: "20px 22px",
              borderRadius: 14,
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--destructive) / 0.4)",
              borderTop: "4px solid hsl(var(--destructive))",
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
                color: "hsl(var(--destructive))",
              }}
            >
              Symptom-patch loop (rejected)
            </div>
            <ul style={{ margin: 0, padding: 0, listStyle: "none", display: "flex", flexDirection: "column", gap: 10 }}>
              {SYMPTOM_LOOP.map((line, i) => (
                <li key={i} style={{ display: "flex", gap: 10, alignItems: "flex-start", fontSize: 15, lineHeight: 1.4, color: "hsl(var(--fg-muted))" }}>
                  <span style={{ color: "hsl(var(--destructive))", fontWeight: 700 }}>✗</span>
                  <span>{line}</span>
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
    </SlideLayout>
  );
}
