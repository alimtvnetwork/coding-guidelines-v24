import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 43: Data & Schema chapter opener.
 *
 * Source: spec/17/31 lines 87-89. PascalCase for tables/types/entities and
 * JSON keys; camelCase for fields/columns.
 */

const BEFORE = `-- Mixed casing everywhere. Every join needs a mental translation.
CREATE TABLE user_profiles (
  user_profile_id  INTEGER PRIMARY KEY AUTOINCREMENT,
  first_name       TEXT,
  last_name        TEXT
);

// API returns snake_case JSON. Frontend maps by hand.
{ "user_profile_id": 42, "first_name": "Ada", "last_name": "Lovelace" }

// TypeScript type invented a third convention.
type user_profile = { user_profile_id: number; first_name: string };`;

const AFTER = `-- PascalCase tables, camelCase columns. One rule, no translation.
CREATE TABLE UserProfile (
  UserProfileId  INTEGER PRIMARY KEY AUTOINCREMENT,
  firstName      TEXT,
  lastName       TEXT
);

// PascalCase JSON keys match the entity name and PK convention.
{ "UserProfileId": 42, "FirstName": "Ada", "LastName": "Lovelace" }

// Types mirror the entity exactly. No mental remap between layers.
type UserProfile = { UserProfileId: number; FirstName: string; LastName: string };`;

export default function DataSchemaChapterOpenerSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 43 · Data & Schema · chapter opener"
      title="PascalCase tables, types, entities, and JSON keys. camelCase columns. One naming law across DB, API, and TypeScript."
      subtitle="Every layer that renames the same entity costs a translation function, a bug class, and a code review argument. Fix the naming at the boundary: PascalCase for tables, types, entities, and JSON keys; camelCase for fields and columns. The entity has ONE name from schema.sql to the React prop."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="sql"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Table is snake_case, JSON is snake_case, TypeScript invents a third name. Every join reads as translation work, every API handler ships a manual re-mapper, and every code review re-argues the naming."
          afterLabel="✅ PascalCase entity `UserProfile` flows unchanged through DB, JSON, and TypeScript. Columns/fields stay camelCase. The name is the primary key of the shared vocabulary."
        />
        <ActionPanel
          slideId="41-data-schema-chapter-opener"
          symptom="A new endpoint returns `{ user_profile_id: 42, first_name: 'Ada' }`. The React screen imports a `UserProfile` type that expects `UserProfileId` and `FirstName`. The screen renders empty and no error fires because the mapper silently returns `undefined`. Support ticket lands two days later. Nobody can find the offending field because grep for `UserProfileId` misses `user_profile_id`, and grep for `user_profile_id` misses the TypeScript side."
          rule="Tables, types, and entities are PascalCase. JSON keys are PascalCase. Fields and columns are camelCase. The entity keeps ONE name from `schema.sql` to the API response to the TypeScript type to the React prop. No layer re-cases. No manual `snake_case <-> PascalCase` mapper. The naming law is enforced at the migration linter, the API contract test, and the TypeScript type generator, so a mismatch fails CI before it fails a user. Per spec/17/31 lines 87 to 89 and the project memory `mem://architecture/database-schema`."
          doThis="Adopt this chapter as the reader's mental model before any specific rule. Point every schema PR at it. Enforce with three checks: (1) a migration linter that rejects `CREATE TABLE` names not matching `^[A-Z][A-Za-z0-9]*$` and columns not matching `^[a-z][A-Za-z0-9]*$`; (2) an API contract test that fails on any JSON key not matching `^[A-Z][A-Za-z0-9]*$`; (3) a codegen step that emits TypeScript types directly from the schema so the entity name cannot drift. On any violation, log `schema.naming.violation` with the layer, entity, and offending token, and block the PR."
        />
      </div>
    </SlideLayout>
  );
}
