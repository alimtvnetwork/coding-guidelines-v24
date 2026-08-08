import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 48: Any PR touching the database ships with a Mermaid ERD.
 *
 * Source: spec/17/31 line 93 (rule 8 in Data & Schema).
 */

const BEFORE = `# PR #4821: "Add referral program"
Files changed:
  + migrations/2026_07_add_referral.sql   (+62 lines)
  + src/db/schema/referral.ts             (+41 lines)
  ~ src/db/schema/customer.ts             (+3 lines)

Description:
  Adds Referral + ReferralReward tables. Small change.

Reviewers see:
  62 lines of CREATE TABLE + 3 new FKs across 3 files, no picture.
  Nobody can tell if Referral joins Customer once or twice, or whether
  ReferralReward hangs off Referral or off Customer directly.
Result: rubber-stamped in 4 minutes. Two weeks later an analytics query
double-counts rewards because the join cardinality was wrong.`;

const AFTER = `# PR #4821: "Add referral program"
Files changed:
  + migrations/2026_07_add_referral.sql   (+62 lines)
  + src/db/schema/referral.ts             (+41 lines)
  ~ src/db/schema/customer.ts             (+3 lines)
  + docs/erd/referral.mmd                 (+22 lines)   <-- REQUIRED

\`\`\`mermaid
erDiagram
  Customer        ||--o{ Referral        : "refers"
  Customer        ||--o{ Referral        : "referred by"
  Referral        ||--o{ ReferralReward  : "grants"
  Referral {
    int    ReferralId PK
    int    ReferrerCustomerId FK
    int    RefereeCustomerId  FK
    text   Notes
  }

  ReferralReward {
    int    ReferralRewardId PK
    int    ReferralId       FK
    int    AmountCents
    text   Comments
  }

\`\`\`
CI gate \`erd-required\` blocks the PR until the diagram renders and
covers every table changed in the migration.`;

export default function ErdRequiredOnDbPrsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 48 · Data & Schema · review artifact"
      title="Every PR that touches the database ships a Mermaid ERD. No diagram, no merge."
      subtitle="Schema changes are hard to review as pure DDL text. A Mermaid ERD makes cardinality, FK direction, and orphan risk visible in seconds so reviewers catch double-joins, missing FKs, and 1-vs-N confusion before they ship."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="markdown"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Migration + two schema files, no picture. Reviewers can't see that `Referral` FKs `Customer` twice (referrer + referee) or that `ReferralReward` hangs off `Referral`, not `Customer`. Approval is a shrug. The join-cardinality bug lands in prod and shows up as inflated reward totals two sprints later."
          afterLabel="✅ Same code change, plus `docs/erd/referral.mmd` covering every table the migration touches. Mermaid renders the double FK from `Referral` to `Customer` (referrer + referee) and the 1:N to `ReferralReward` so reviewers see the shape at a glance. CI check `erd-required` fails when a migration adds or alters a table without a matching `.mmd` update."
        />
        <ActionPanel
          slideId="46-erd-required-on-db-prs"
          symptom="Someone opens a PR that adds a `Referral` table with two FKs to `Customer` (referrer and referee) plus a child `ReferralReward` table, and the review is 62 lines of `CREATE TABLE` DDL across three files with no diagram. Reviewers can't tell whether `Referral` joins `Customer` once or twice, or whether `ReferralReward` hangs off `Referral` or `Customer`. It merges in four minutes. Two weeks later analytics reports doubled reward totals because a join was written against the wrong cardinality, and the fix requires a data-repair migration on 180k rows."
          rule="Any pull request whose diff touches a migration, an ORM schema file, or a raw DDL file must include a Mermaid ERD (`.mmd` file under `docs/erd/`) that covers every table added or altered. The ERD shows PKs, FKs (with direction and cardinality), and the narrative columns (`Description` / `Notes` / `Comments`). If the change modifies an existing table, the existing ERD is updated in the same commit. Per spec/17/31 line 93 (Data & Schema rule 8). Reviewers approve only after opening the rendered diagram: the diagram, not the DDL, is the review artifact."
          doThis="Enforce with a required CI check `erd-required`: (1) parse the diff for any file matching `migrations/**`, `**/db/schema/**`, or `**/*.sql`; (2) extract every table name added or altered; (3) fail the check when any of those tables is missing from `docs/erd/**.mmd`, and print the missing table names in the check output; (4) also fail when a listed `.mmd` doesn't render via `mermaid-cli` in CI. Log `erd.required.missing` with `{ pr, table, migrationFile }` on failure. PR template: 'ERD updated? path: `docs/erd/<area>.mmd`'. Reviewer checklist: 'I opened the rendered diagram and confirmed FK direction, cardinality, and narrative columns'. On violation, block the merge until the diagram lands."
        />
      </div>
    </SlideLayout>
  );
}
