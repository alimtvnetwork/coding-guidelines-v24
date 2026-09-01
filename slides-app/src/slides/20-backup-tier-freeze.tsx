import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CalloutQuote } from "@/components/CalloutQuote";

/**
 * SS-02 task 23: Backup-tier freeze / v2.0 seed-key deferral.
 *
 * Explains why `MainWorker.Backup.*` seed keys in spec/19 remain listed as
 * defaults-only (no v2.0.0 materialization). Materializing them would break
 * MAIN-900-01 (SpecContradiction) safe-fail for the frozen v2.0 Backup tier.
 */

interface Column {
  label: string;
  tone: "danger" | "ok";
  body: readonly string[];
}

const COLUMNS: readonly Column[] = [
  {
    label: "If we materialize now",
    tone: "danger",
    body: [
      "Freezes seed values before Phase 12 (Backup consolidation) closes.",
      "Removes MAIN-900-01 SpecContradiction safe-fail: implementers stop halting on gaps and start guessing.",
      "Locks tunables (`MaxKeyAgeSeconds`, `RsaKeySizeBits`) that ch. 20 §7 still marks as under review.",
      "Downstream Workers pin to values that will change, forcing a breaking v3.0 later.",
    ],
  },
  {
    label: "Defer until Phase 12 closes",
    tone: "ok",
    body: [
      "Chapter 20 ships defaults as guidance, not as a v2.0 contract.",
      "MAIN-900-01 keeps halting literal AI implementers on any Backup-tier gap.",
      "Rotation cadence, RSA size, and grace windows stay tunable during the Phase-12 threat-model pass (ch. 24).",
      "Materialization lands together as v2.1 with a single SemVer minor bump and one changelog entry.",
    ],
  },
];

export default function BackupTierFreezeSlide() {
  return (
    <SlideLayout
      eyebrow="Principle · Spec discipline"
      title="Backup-tier freeze: defaults today, contract at Phase 12"
      subtitle="`MainWorker.Backup.*` seed keys stay defaults-only. Materializing them prematurely would delete the MAIN-900-01 safe-fail that keeps mediocre AI implementers honest."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <ActionPanel
          slideId="20-backup-tier-freeze"
          symptom="A PR proposes replacing the `TBD` / defaults-only rows in spec/19 ch. 20 with concrete v2.0.0 values so 'the AI can just implement it'."
          rule="Backup tier stays frozen until Phase 12 (Backup consolidation) closes. Defaults in ch. 20 §7 are guidance. Contract-grade materialization ships as v2.1 with the ch. 24 threat model."
          doThis="Reject the PR. If a value must be pinned early, file it in a non-active proposal doc under 02-spec/19/audit and cite MAIN-900-01. Do NOT edit the authoritative chapter."
        />
        <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 12 }}>
          {COLUMNS.map((col) => (
            <div
              key={col.label}
              style={{
                padding: "16px 20px",
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
                  letterSpacing: "0.12em",
                  textTransform: "uppercase",
                  color: col.tone === "danger" ? "hsl(var(--destructive))" : "hsl(var(--primary))",
                }}
              >
                {col.label}
              </div>
              <ul style={{ margin: 0, paddingLeft: 20, display: "flex", flexDirection: "column", gap: 6 }}>
                {col.body.map((line) => (
                  <li key={line} style={{ fontSize: 13, lineHeight: 1.4, color: "hsl(var(--fg))" }}>
                    {line}
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>
        <CalloutQuote
          quote="A frozen tier with a loud safe-fail beats a pinned tier with a silent guess."
          attribution="spec/19 ch. 25 §6 (MAIN-900-01 rationale)"
          accent="primary"
        />
      </div>
    </SlideLayout>
  );
}
