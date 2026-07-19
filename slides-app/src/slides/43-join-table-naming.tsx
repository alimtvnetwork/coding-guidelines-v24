import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 45: Join-table naming and composite PK.
 *
 * Source: spec/17/31 line 91 (Type/Status/Category via N-M join tables) and
 * `mem://architecture/database-schema` (join tables exempt from Description/Notes).
 */

const BEFORE = `-- Surrogate PK on a pure join table. Duplicates allowed. Reads need DISTINCT.
CREATE TABLE user_roles (
  user_role_id  INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id       INTEGER NOT NULL,
  role_id       INTEGER NOT NULL,
  Description   TEXT NULL,   -- join table should not carry narrative columns
  Notes         TEXT NULL
);

-- Name gives no hint which two entities it joins.
-- Grep for 'User' misses it. Grep for 'Role' misses it.`;

const AFTER = `-- Join tables are {A}{B}. Composite PK is ({A}Id, {B}Id). No surrogate.
CREATE TABLE UserRole (
  UserId  INTEGER NOT NULL REFERENCES User(UserId),
  RoleId  INTEGER NOT NULL REFERENCES Role(RoleId),
  PRIMARY KEY (UserId, RoleId)
);

-- Type/Status/Category resolve through a registered enum + join table,
-- never a free-form string column on the parent (spec/17/31 line 91).
CREATE TABLE OrderStatus (
  OrderId          INTEGER NOT NULL REFERENCES "Order"(OrderId),
  StatusEnumId     INTEGER NOT NULL REFERENCES StatusEnum(StatusEnumId),
  EffectiveAt      TEXT NOT NULL,
  PRIMARY KEY (OrderId, StatusEnumId, EffectiveAt)
);`;

export default function JoinTableNamingSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 45 · Data & Schema · join tables"
      title="Join tables are `{A}{B}`. Composite PK is `({A}Id, {B}Id)`. No surrogate PK. No `Description`, no `Notes`."
      subtitle="A surrogate PK on a pure join lets the same (UserId, RoleId) pair be inserted twice. Every read then needs DISTINCT, every uniqueness invariant lives in application code, and the table cannot answer 'does user X have role Y?' with a single indexed lookup. Name it after both entities and let the PK enforce the invariant."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="sql"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ `user_roles` with a surrogate `user_role_id`, snake_case, and narrative columns. Duplicates are legal at the DB layer, the name hides both parents from grep, and the join table carries fields it is exempt from carrying."
          afterLabel="✅ `UserRole` with composite PK `(UserId, RoleId)`. Uniqueness is a schema invariant, both parents show up in every grep, and the table stays pure. Type/Status/Category resolve through registered enum + join, never a free-form string column."
        />
        <ActionPanel
          slideId="43-join-table-naming"
          symptom="A permissions bug: user Ada shows two 'admin' rows in the audit console. Support blames the UI. The UI blames the API. The API blames the DB. Nobody catches that `user_roles` allows duplicate `(user_id, role_id)` pairs because the PK is a surrogate `user_role_id`. Meanwhile `Order.Status = 'shipping'` (typo of 'shipped') passes every check because the column is a free-form string, and the analytics dashboard silently drops the row."
          rule="Pure many-to-many join tables are named `{A}{B}` in PascalCase (alphabetical or domain-natural order, pick one and stick with it) and have a composite primary key `({A}Id, {B}Id)`, both columns `NOT NULL` and FK-referencing their parent's `{Table}Id`. No surrogate PK. No `Description`, no `Notes`, no `Comments` (join tables are exempt per `mem://architecture/database-schema`). `Type`, `Status`, `Category`, and `Kind` on any entity resolve through a registered enum table plus a join table, never a free-form string column on the parent (spec/17/31 line 91). When the relationship carries history (status over time), the PK extends with the time key: `(OrderId, StatusEnumId, EffectiveAt)`."
          doThis="Enforce in the migration linter: (1) reject any table whose name matches two known entity names concatenated but has an extra surrogate PK; (2) reject any join table that declares `Description`, `Notes`, or `Comments`; (3) reject any column named `Type`, `Status`, `Category`, or `Kind` typed `TEXT` on a non-enum table, and require the join-table + enum pattern instead. On violation, log `schema.join.violation` with the table, offending column, and the rule number, and block the PR. Add a PR checklist item: 'Join table is `{A}{B}`, composite PK, no narrative columns; Status/Type/Category use enum + join'."
        />
      </div>
    </SlideLayout>
  );
}
