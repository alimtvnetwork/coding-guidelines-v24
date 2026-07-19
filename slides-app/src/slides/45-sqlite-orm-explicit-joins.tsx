import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 47: SQLite default + ORM + explicit joins/PK/FK.
 *
 * Source: spec/17/31 line 93.
 */

const BEFORE = `// Raw SQL string built by concatenation. No FK. No PK type. Implicit join.
const rows = await db.query(
  "SELECT * FROM \\"order\\" o, customer c " +
  "WHERE o.customer_id = c.id AND c.name LIKE '%" + q + "%'"
);
// - Implicit CROSS JOIN + WHERE (easy to forget a predicate = full Cartesian).
// - String concat = SQL injection.
// - No schema-side FK, so orphaned OrderRows exist and nobody knows.
// - Postgres in dev, MySQL in staging, SQLite nowhere. Every env drifts.`;

const AFTER = `// SQLite by default. ORM model with explicit PK, FK, and JOIN.
// (Drizzle shown; GORM / Prisma / SQLAlchemy follow the same shape.)
export const Customer = sqliteTable("Customer", {
  CustomerId: integer("CustomerId").primaryKey({ autoIncrement: true }),
  Name:       text("Name").notNull(),
});

export const Order = sqliteTable("Order", {
  OrderId:    integer("OrderId").primaryKey({ autoIncrement: true }),
  CustomerId: integer("CustomerId").notNull().references(() => Customer.CustomerId),
  Notes:      text("Notes"),
  Comments:   text("Comments"),
});

const rows = await db
  .select({ OrderId: Order.OrderId, Name: Customer.Name })
  .from(Order)
  .innerJoin(Customer, eq(Customer.CustomerId, Order.CustomerId))
  .where(like(Customer.Name, \`%\${q}%\`)); // parameterized`;

export default function SqliteOrmExplicitJoinsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 47 · Data & Schema · engine + ORM"
      title="Default DB is SQLite. Prefer an ORM. Define PKs, FKs, and joins explicitly. No raw string SQL, no implicit joins."
      subtitle="One engine across dev, CI, and prod removes the 'works on my Postgres' bug class. An ORM turns the schema into typed code that grep and the compiler can read. Explicit `JOIN ... ON` and declared `references()` make orphans impossible and reviews boring in the good way."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="ts"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Three engines across three environments, raw SQL with string concatenation, implicit `FROM a, b WHERE ...` join, no FK. One missing predicate becomes a Cartesian product in prod; one missing sanitizer becomes an injection; one orphan row corrupts a report forever."
          afterLabel="✅ SQLite by default. ORM tables with `primaryKey({ autoIncrement: true })` and `.references(() => Parent.PK)`. Query uses `.innerJoin(...).where(...)` with parameter binding. Compiler catches missing columns, migrations catch missing FKs, one engine catches env drift."
        />
        <ActionPanel
          slideId="45-sqlite-orm-explicit-joins"
          symptom="Search returns every customer for every order (Cartesian product) after a WHERE clause is accidentally removed during a refactor. Nobody catches it in code review because the raw SQL string spans four lines of concatenation. Meanwhile a security audit flags `LIKE '%' + q + '%'` as an injection, staging (MySQL) accepts a query that dev (Postgres) rejects, and the analytics team finds 312 `Order` rows whose `customer_id` points at deleted customers because there is no `FOREIGN KEY`."
          rule="Default engine is SQLite for dev, CI, and prod unless a documented requirement forces otherwise (concurrent writers, replication, geo). Use an ORM (Drizzle, GORM, Prisma, SQLAlchemy, ActiveRecord); define every primary key, foreign key, and join explicitly in the model. Queries use the ORM query builder or parameterized statements: never string concatenation, never `FROM a, b WHERE a.x = b.y` implicit joins, never a query without an `ON` clause. Cross-table reads use `.innerJoin` / `.leftJoin` with an explicit predicate. Per spec/17/31 line 93. FKs are declared at migration time so orphans are a constraint error, not a support ticket."
          doThis="Enforce with three gates: (1) a migration linter that fails when a `CREATE TABLE` declares a column named `{Other}Id` without a matching `REFERENCES {Other}({Other}Id)`, and logs `schema.fk.missing`; (2) a code linter that bans raw SQL strings outside a whitelisted `db/raw/*.sql` folder (each file requires a comment citing the ticket + reviewer); (3) a CI job that runs the full test suite against SQLite and refuses any env-specific SQL dialect. PR checklist: 'ORM model owns the schema, every `{Other}Id` has `.references(...)`, every cross-table query uses explicit `.innerJoin(..., on)`'. On violation, block the merge."
        />
      </div>
    </SlideLayout>
  );
}
