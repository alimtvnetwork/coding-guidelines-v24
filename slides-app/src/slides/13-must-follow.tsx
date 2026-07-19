import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";

/**
 * SS-02 task 16: "Must Follow" opener slide.
 *
 * Distills the 5 non-negotiables from
 * `spec/17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md`
 * §"Must Follow and without negotiation" into a slide-friendly checklist.
 * Uses <ActionPanel> so the SS-02 SRA structural validator passes.
 */

interface NonNegotiable {
  index: string;
  title: string;
  body: string;
}

const NON_NEGOTIABLES: readonly NonNegotiable[] = [
  {
    index: "01",
    title: "Read first, guess never",
    body: "Open the exact files, functions, and lines involved. No skimming, no filename-only assumptions.",
  },
  {
    index: "02",
    title: "One-sentence root cause",
    body: "Write the root cause in a single sentence before touching any code. If you cannot, keep reading.",
  },
  {
    index: "03",
    title: "Minimum correct fix",
    body: "Change only what the root cause requires. No try/catch camouflage, no fallback values, no re-render hacks.",
  },
  {
    index: "04",
    title: "Verify in the logs",
    body: "Show the failing signal turning green: build output, console, network, or preview. If no log exists, add one.",
  },
  {
    index: "05",
    title: "Ship the trail",
    body: "List every remaining task, bump the version, update the changelog and release notes. Depth is the job.",
  },
];

export default function MustFollowSlide() {
  return (
    <SlideLayout
      eyebrow="Must Follow · Non-Negotiable"
      title="Five rules or auto-reject"
      subtitle="Violating any of these is treated the same as violating RULE 0. Same tier, same outcome: the change does not ship."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 24, marginTop: 12 }}>
        <ActionPanel
          slideId="13-must-follow"
          symptom="Past turns shipped wrong step counts, missing remaining-tasks lists, symptom patches sold as fixes, ignored guidelines, forgotten version bumps, and unread logs."
          rule="Read the code, find the root cause in one sentence, apply the minimum correct fix, verify it in the logs, list every remaining task, bump version, and update changelog + release notes."
          doThis="Before your next commit, walk the five-item checklist below and confirm each one out loud. Any 'no' means the change is not ready."
        />
        <ol
          style={{
            display: "grid",
            gridTemplateColumns: "1fr 1fr",
            gap: "14px 28px",
            margin: 0,
            padding: 0,
            listStyle: "none",
          }}
        >
          {NON_NEGOTIABLES.map((item) => (
            <li
              key={item.index}
              style={{
                display: "flex",
                gap: 14,
                alignItems: "flex-start",
                padding: "12px 16px",
                borderRadius: 12,
                background: "hsl(var(--bg-raised))",
                border: "1px solid hsl(var(--border))",
              }}
            >
              <span
                style={{
                  fontSize: 22,
                  fontWeight: 700,
                  color: "hsl(var(--accent))",
                  minWidth: 32,
                  fontFamily: "var(--font-mono)",
                }}
              >
                {item.index}
              </span>
              <span style={{ display: "flex", flexDirection: "column", gap: 4 }}>
                <span style={{ fontSize: 18, fontWeight: 700, color: "hsl(var(--fg))" }}>
                  {item.title}
                </span>
                <span style={{ fontSize: 15, lineHeight: 1.35, color: "hsl(var(--fg-muted))" }}>
                  {item.body}
                </span>
              </span>
            </li>
          ))}
        </ol>
      </div>
    </SlideLayout>
  );
}
