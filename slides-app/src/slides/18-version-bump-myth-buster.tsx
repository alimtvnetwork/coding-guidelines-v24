import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CalloutQuote } from "@/components/CalloutQuote";

/**
 * SS-02 task 21: Version-bump myth-buster.
 *
 * Kills the misread of the v1.2 update ("workflow simplified") as
 * "version bumps are optional now". They are not. Spec 17/31 §Must Follow
 * (v1.4.0, line 16) still lists "bump the version, update changelog and
 * release notes" as a non-negotiable, same tier as RULE 0.
 */

interface MythRow {
  claim: string;
  reality: string;
}

const ROWS: readonly MythRow[] = [
  {
    claim: "\"v1.2 dropped the mandatory version bump.\"",
    reality: "v1.2 simplified the workflow wording. §Must Follow (v1.4.0) still lists the bump as non-negotiable.",
  },
  {
    claim: "\"Docs-only or comment-only PRs skip the bump.\"",
    reality: "Any merged change ships. Consumers pin versions; a silent change breaks reproducibility. Bump patch, at minimum.",
  },
  {
    claim: "\"Bumping is CI's job; I just push.\"",
    reality: "The bump encodes intent (patch / minor / major). CI cannot infer intent from a diff. Author bumps, CI verifies.",
  },
  {
    claim: "\"Changelog can be filled in later.\"",
    reality: "Later never comes. Same PR, same commit. Reviewer rejects a bump without a changelog entry.",
  },
];

export default function VersionBumpMythBusterSlide() {
  return (
    <SlideLayout
      eyebrow="Principle · Release hygiene"
      title="Version bump: still required. Every PR. No exceptions."
      subtitle="The v1.2 update simplified the workflow. It did not delete the bump. §Must Follow in spec/17/31 (v1.4.0) still lists it as auto-reject on violation."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <ActionPanel
          slideId="18-version-bump-myth-buster"
          symptom="PRs land without a version bump or changelog entry, citing 'the v1.2 update made it optional'. Downstream pins break silently because two different code states share one version string."
          rule="Every merged PR bumps package.json (patch / minor / major per SemVer intent) AND adds a CHANGELOG.md entry in the same commit. Non-negotiable per spec/17/31 §Must Follow (v1.4.0, line 16)."
          doThis="Add the bump + changelog to your PR template checklist. Reviewer rejects on missing bump the same way they reject on a failing test. Run `node scripts/sync-version.mjs` after editing package.json."
        />
        <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 12 }}>
          {ROWS.map((row) => (
            <div
              key={row.claim}
              style={{
                padding: "14px 18px",
                borderRadius: 12,
                background: "hsl(var(--bg-raised))",
                border: "1px solid hsl(var(--border))",
                display: "grid",
                gridTemplateColumns: "auto 1fr",
                columnGap: 14,
                rowGap: 6,
                alignItems: "start",
              }}
            >
              <div
                style={{
                  fontSize: 11,
                  fontWeight: 700,
                  letterSpacing: "0.12em",
                  textTransform: "uppercase",
                  color: "hsl(var(--destructive))",
                }}
              >
                Myth
              </div>
              <div style={{ fontSize: 14, lineHeight: 1.4, color: "hsl(var(--fg))" }}>
                {row.claim}
              </div>
              <div
                style={{
                  fontSize: 11,
                  fontWeight: 700,
                  letterSpacing: "0.12em",
                  textTransform: "uppercase",
                  color: "hsl(var(--primary))",
                }}
              >
                Reality
              </div>
              <div style={{ fontSize: 14, lineHeight: 1.4, color: "hsl(var(--fg))" }}>
                {row.reality}
              </div>
            </div>
          ))}
        </div>
        <CalloutQuote
          quote="If it shipped, it has a version. If it has a version, it has a changelog line."
          attribution="spec/17/31 §Must Follow"
          accent="primary"
        />
      </div>
    </SlideLayout>
  );
}
