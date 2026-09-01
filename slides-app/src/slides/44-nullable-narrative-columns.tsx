import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 46: Nullable narrative columns.
 *
 * Source: 02-spec/17/31 line 92 and `mem://architecture/database-schema`
 * (DB-FREETEXT-001 + MISSING-DESC-001 linters).
 */

const BEFORE = `-- Entity table without a narrative column. Ops has nowhere to leave context.
CREATE TABLE Customer (
  CustomerId  INTEGER PRIMARY KEY AUTOINCREMENT,
  Name        TEXT NOT NULL
);

-- Transactional table with NOT NULL narrative + DEFAULT. Every insert lies.
CREATE TABLE "Order" (
  OrderId   INTEGER PRIMARY KEY AUTOINCREMENT,
  Notes     TEXT NOT NULL DEFAULT '',   -- empty string is not a note
  Comments  TEXT NOT NULL DEFAULT 'n/a' -- literal 'n/a' pollutes analytics
);`;

const AFTER = `-- Entity / reference tables carry Description TEXT NULL. Nullable. No DEFAULT.
CREATE TABLE Customer (
  CustomerId   INTEGER PRIMARY KEY AUTOINCREMENT,
  Name         TEXT NOT NULL,
  Description  TEXT NULL
);

-- Transactional tables carry BOTH Notes and Comments. Nullable. No DEFAULT.
CREATE TABLE "Order" (
  OrderId    INTEGER PRIMARY KEY AUTOINCREMENT,
  Notes      TEXT NULL,
  Comments   TEXT NULL
);

-- Join tables are exempt (Rule 45). No narrative columns on UserRole.`;

export default function NullableNarrativeColumnsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 46 · Data & Schema · narrative columns"
      title="Entity/ref tables: `Description TEXT NULL`. Transactional: `Notes TEXT NULL` + `Comments TEXT NULL`. All nullable. No `DEFAULT`. Join tables exempt."
      subtitle="A `NOT NULL DEFAULT ''` narrative column reads as 'the field is present' in every query, but every row is a lie. Analytics counts empty strings as real notes, dashboards show '100% documented', and the first real note gets lost in the noise. Make absence explicit: `NULL` means 'no context yet', a value means 'someone typed this'."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="sql"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ `Customer` has no narrative column (ops has nowhere to leave context). `Order` forces `NOT NULL DEFAULT ''` and `DEFAULT 'n/a'`, so `COUNT(Notes)` and `COUNT(DISTINCT Comments)` both return garbage and every dashboard 'has notes' turns out to mean nothing."
          afterLabel="✅ Entity/ref tables carry `Description TEXT NULL`. Transactional tables carry both `Notes TEXT NULL` and `Comments TEXT NULL`. Nullable, no `DEFAULT`. `IS NULL` cleanly separates 'nobody wrote anything' from 'somebody wrote something'. Join tables stay pure (Rule 45)."
        />
        <ActionPanel
          slideId="44-nullable-narrative-columns"
          symptom="A product manager asks 'what percent of orders have operator notes?'. The dashboard says 100%. Ops swears they never write notes. An engineer runs `SELECT DISTINCT Notes FROM Order LIMIT 5` and sees `''`, `''`, `'n/a'`, `''`, `'n/a'`. Every insert has been writing the DEFAULT. Meanwhile a customer complaint escalates and support cannot find the `Description` on `Customer` because the column does not exist, so the context lives in a Slack thread that already scrolled off."
          rule="Entity and reference tables include `Description TEXT NULL`. Transactional tables include BOTH `Notes TEXT NULL` and `Comments TEXT NULL`. Every narrative column is nullable and has no `DEFAULT`. Join tables are exempt. Per 02-spec/17/31 line 92 and `mem://architecture/database-schema`. Absence is `NULL`; a value means a human (or a well-identified system) wrote it. Never coerce empty into a sentinel string like `''`, `'n/a'`, `'-'`, or `'TBD'`; those are lies that pass every `NOT NULL` check and corrupt every aggregate."
          doThis="Enforce with two existing SQL linters (see `mem://sessions/2026-04-sql-linter-rules`): (1) `MISSING-DESC-001` fails a migration when an entity/ref table lacks `Description TEXT NULL` or a transactional table lacks `Notes`+`Comments TEXT NULL`, and rejects any `NOT NULL` or `DEFAULT` on those columns; (2) `DB-FREETEXT-001` blocks free-form `Status/Type/Category TEXT` columns (Rule 45). Waivers require a line in the migration file citing the ticket. On violation, log `schema.narrative.violation` with the table, column, and rule id. PR checklist: 'Narrative columns present, nullable, no DEFAULT; join tables carry none'."
        />
      </div>
    </SlideLayout>
  );
}
