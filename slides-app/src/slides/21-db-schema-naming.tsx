import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 23: DB schema naming.
 *
 * Sourced from 02-spec/17/31 §"Data and Schema Rules" (lines 84-92):
 *   - Tables/entities: PascalCase
 *   - Fields/columns:  camelCase
 *   - JSON keys:       PascalCase
 *   - Primary key:     `{TableName}Id`, integer auto-increment, no UUIDs
 */

const BEFORE = `CREATE TABLE users (
  id            UUID PRIMARY KEY,
  Email_Address TEXT,
  CreatedAt     TIMESTAMP
);`;

const AFTER = `CREATE TABLE User (
  UserId       INTEGER PRIMARY KEY AUTOINCREMENT,
  emailAddress TEXT,
  createdAt    TIMESTAMP
);`;

export default function DbSchemaNamingSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 23 · Data & Schema"
      title="PascalCase entities, camelCase fields, `{Table}Id` PKs"
      subtitle="One predictable shape across every table so joins, ORMs and JSON payloads line up without translation layers."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <CodeDiff
          language="sql"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ UUIDs, mixed casing, plural table"
          afterLabel="✅ Singular PascalCase table, `UserId` PK, camelCase columns"
        />
        <ActionPanel
          slideId="21-db-schema-naming"
          symptom="A migration lands `users.id UUID`, `Email_Address` and `CreatedAt` on the same table. The ORM now needs per-column aliases and joins read as guesswork."
          rule="Tables and entities PascalCase (singular). Columns camelCase. Primary key `{TableName}Id`, INTEGER PRIMARY KEY AUTOINCREMENT. No UUIDs. JSON keys PascalCase per 02-spec/17/31 §Data and Schema."
          doThis="Rename the table to singular PascalCase, add a `{Table}Id` integer PK, and lowercase the first letter of every column. Fix as you touch each migration, do not bulk-rewrite."
        />
      </div>
    </SlideLayout>
  );
}
