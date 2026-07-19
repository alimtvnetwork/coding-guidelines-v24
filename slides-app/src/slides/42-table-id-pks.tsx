import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 44: `{TableName}Id` integer PKs. No UUIDs.
 *
 * Source: spec/17/31 line 90 and `mem://architecture/database-schema`.
 */

const BEFORE = `-- Every table invents its own PK name and type.
CREATE TABLE UserProfile (
  id           TEXT PRIMARY KEY,           -- UUID string
  firstName    TEXT
);

CREATE TABLE OrderItem (
  order_item_uuid  TEXT PRIMARY KEY,       -- different name, same idea
  user_id          TEXT REFERENCES UserProfile(id)
);

-- Joins read as noise. Indexes are 4x larger. Debug logs are unreadable.
SELECT o.order_item_uuid, u.id
FROM OrderItem o JOIN UserProfile u ON u.id = o.user_id;`;

const AFTER = `-- One PK rule: {TableName}Id, INTEGER, AUTOINCREMENT.
CREATE TABLE UserProfile (
  UserProfileId  INTEGER PRIMARY KEY AUTOINCREMENT,
  FirstName      TEXT
);

CREATE TABLE OrderItem (
  OrderItemId    INTEGER PRIMARY KEY AUTOINCREMENT,
  UserProfileId  INTEGER NOT NULL REFERENCES UserProfile(UserProfileId)
);

-- Joins are self-documenting. PK name equals FK name. Grep finds every use.
SELECT o.OrderItemId, o.UserProfileId
FROM OrderItem o JOIN UserProfile u ON u.UserProfileId = o.UserProfileId;`;

export default function TableIdPksSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 44 · Data & Schema · primary keys"
      title="Primary keys are {TableName}Id, INTEGER PRIMARY KEY AUTOINCREMENT. No UUIDs. No `id`. No mixed types."
      subtitle="A PK named `id` forces every join to alias, every FK to invent a new name, and every log line to ask 'id of what?'. A UUID PK quadruples the index, kills range scans, and turns debugging into copy-paste roulette. Name the PK after the table, keep it an integer, and the schema explains itself."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="sql"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Mixed PK names (`id`, `order_item_uuid`), TEXT UUID storage, FK column renamed on the child side. Joins are ambiguous, indexes are bloated, and no grep finds all references to one entity's PK."
          afterLabel="✅ Every PK is `{TableName}Id INTEGER PRIMARY KEY AUTOINCREMENT`. FKs reuse the exact PK name. Joins read like English, indexes are compact, and grep for `UserProfileId` returns every producer and consumer."
        />
        <ActionPanel
          slideId="42-table-id-pks"
          symptom="A support engineer opens `OrderItem` in the admin console and sees `order_item_uuid: 'a7f3...'` and `user_id: 'b91c...'`. They cannot tell which user without running a second query, and the UUIDs are impossible to read out loud on a call. Meanwhile the query planner picks a full scan because the UUID index no longer fits in memory, and p95 latency doubles. The engineer opens a ticket that says 'DB is slow', which lands on a queue with 40 others just like it, and nobody connects it to the PK type."
          rule="Every table's primary key is named `{TableName}Id` and typed `INTEGER PRIMARY KEY AUTOINCREMENT` (or the platform equivalent, e.g. `BIGSERIAL` on Postgres). Foreign keys reuse the exact PK name on the child side, so `OrderItem.UserProfileId` references `UserProfile.UserProfileId`. No UUIDs, no `id`, no `uuid`, no `order_item_uuid`. Per spec/17/31 line 90 and `mem://architecture/database-schema`. Public identifiers that must be unguessable (share links, invite tokens) live in a separate `PublicId TEXT UNIQUE` column, never as the PK."
          doThis="Enforce the rule in three places so nothing slips: (1) a migration linter that rejects any `CREATE TABLE` whose PK column is not `{TableName}Id` INTEGER, and logs `schema.pk.violation` with the table and offending column; (2) an FK-name linter that fails when the child column does not match the referenced PK name exactly; (3) a code-review checklist item on every schema PR: 'PK is `{Table}Id INTEGER`, FKs reuse the name, no UUID PKs'. When a legacy table cannot be renamed today, add a waiver line in `spec/17/31` with the ticket ID and the migration date. Never grant a silent exemption."
        />
      </div>
    </SlideLayout>
  );
}
