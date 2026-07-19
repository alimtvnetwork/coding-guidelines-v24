import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CalloutQuote } from "@/components/CalloutQuote";

/**
 * SS-02 task 22: Trust-boundary teaser.
 *
 * Surfaces the one-way trust gradient defined in spec/19 chapters 26 and 27:
 * Main > Worker > Backup > Git. Every arrow is one-way at the trust layer.
 * Reverse channels are the bug, not a shortcut.
 */

interface TierRow {
  from: string;
  to: string;
  allowed: string;
  forbidden: string;
}

const TIERS: readonly TierRow[] = [
  {
    from: "Main",
    to: "Worker",
    allowed: "Push update instructions, mint Worker JWT, route tenants.",
    forbidden: "Worker cannot mutate Main. Heartbeat/ack/keyset only.",
  },
  {
    from: "Worker",
    to: "Backup",
    allowed: "Seal + push CDC envelopes (ch. 19/22).",
    forbidden: "Backup has NO outbound to primary. Zero contact with Main.",
  },
  {
    from: "Worker",
    to: "Git",
    allowed: "Push to its OWN repo namespace only (ch. 27).",
    forbidden: "Cannot touch other workers' repos. Git cannot call back.",
  },
];

export default function TrustBoundariesTeaserSlide() {
  return (
    <SlideLayout
      eyebrow="Principle · Trust boundaries"
      title="One-way trust: Main > Worker > Backup > Git"
      subtitle="Compromise of any lower-trust node MUST NOT escalate upward. Reverse channels are the bug, per spec/19 chs. 26 and 27."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <ActionPanel
          slideId="19-trust-boundaries-teaser"
          symptom="An 'admin override' or 'temporary reverse call' from Worker back into Main (or Backup back into Worker) shows up in a PR, framed as a shortcut for ops tooling."
          rule="Every trust arrow is one-way. Worker cannot mutate Main. Backup has NO contact with Main and NO outbound to primary Worker. Only the primary Worker holds git creds and only for its own repo namespace."
          doThis="Read spec/19 ch. 26 §2 (trust matrix) and ch. 27 (git backup targets). If your change needs a reverse call, the design is wrong: model it as a Main-initiated pull, not a lower-tier push."
        />
        <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr 1fr", gap: 12 }}>
          {TIERS.map((tier) => (
            <div
              key={`${tier.from}-${tier.to}`}
              style={{
                padding: "14px 18px",
                borderRadius: 12,
                background: "hsl(var(--bg-raised))",
                border: "1px solid hsl(var(--border))",
                display: "flex",
                flexDirection: "column",
                gap: 10,
              }}
            >
              <div
                style={{
                  fontSize: 12,
                  fontWeight: 700,
                  letterSpacing: "0.1em",
                  textTransform: "uppercase",
                  color: "hsl(var(--muted-fg))",
                }}
              >
                {tier.from} → {tier.to}
              </div>
              <div>
                <div style={{ fontSize: 10, fontWeight: 700, letterSpacing: "0.12em", textTransform: "uppercase", color: "hsl(var(--primary))" }}>
                  Allowed
                </div>
                <div style={{ fontSize: 13, lineHeight: 1.4, color: "hsl(var(--fg))", marginTop: 3 }}>
                  {tier.allowed}
                </div>
              </div>
              <div>
                <div style={{ fontSize: 10, fontWeight: 700, letterSpacing: "0.12em", textTransform: "uppercase", color: "hsl(var(--destructive))" }}>
                  Forbidden
                </div>
                <div style={{ fontSize: 13, lineHeight: 1.4, color: "hsl(var(--fg))", marginTop: 3 }}>
                  {tier.forbidden}
                </div>
              </div>
            </div>
          ))}
        </div>
        <CalloutQuote
          quote="If your change needs a reverse arrow, the design is wrong. Model it as a higher-tier pull, not a lower-tier push."
          attribution="spec/19 ch. 26 §2"
          accent="primary"
        />
      </div>
    </SlideLayout>
  );
}
