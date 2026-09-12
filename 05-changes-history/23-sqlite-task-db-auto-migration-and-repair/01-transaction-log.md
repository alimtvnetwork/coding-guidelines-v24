# Task 23: SQLite Task Database Auto-Migration and Repair

## 1. Header & Metadata

- **Date:** 2026-09-09
- **Author/Agent:** Antigravity Master Orchestrator
- **Status:** Completed
- **Affected Packages:**
  - `coding-guidelines/common/pkg/applogger/sqlitelogger`

---

## 2. Context & Goals

The user requested: *"when system is working iwth task db it should fix or mirgate that task db as well"*.

To satisfy this requirement:
1. **Automatic Schema Audit & Migration on Access:** Whenever the system accesses or opens a task database (via `GetTaskDb`, `WriteTask`, `QueryTaskLogs`, or `GetTaskSummary`) or the global logs database, it automatically validates, audits, migrates, and repairs the database schema.
2. **Resilience to Legacy / Outdated Schemas:** If a task database was created with an earlier or incomplete schema (e.g. missing columns such as `fields_json`, `stack_trace`, `duration_ms`, or `status`), the system inspects existing table columns via `PRAGMA table_info(logs)` and automatically executes `ALTER TABLE logs ADD COLUMN <col> <type>` without data loss.
3. **Index Reconstruction & Performance Tuning:** Automatically ensures query indexes (`idx_logs_task_id`, `idx_logs_timestamp`, `idx_logs_level`, `idx_logs_status`) are created idempotently.
4. **Integrity & Repair Engine:** Executes `PRAGMA quick_check` / integrity validation, guarantees presence of `schema_migrations` and `logs` tables, repairs missing structures, and tracks schema versions in `schema_migrations`.
5. **Explicit Management APIs:** Exposes first-class methods on `SplitDBManager`:
   - `MigrateTaskDb(taskId string) *appfault.AppError`
   - `RepairTaskDb(taskId string) *appfault.AppError`
   - `MigrateAllTaskDbs() *appfault.AppError`
   - `RepairAllTaskDbs() *appfault.AppError`
   - `MigrateMainDb() *appfault.AppError`
   - `RepairMainDb() *appfault.AppError`
   - `IsManualMigrationOnly() bool` and `SetManualMigrationOnly(isManual bool)`

---

## 3. Files Changed / Created

### New Source Files

- `04-code/golang/pkg/applogger/sqlitelogger/migration.go`:
  - Implements `EnsureBaseSchema`, `GetCurrentSchemaVersion`, `RecordMigrationVersion`, `ApplyIndexes`, `QueryTableColumns`, `AddMissingColumn`, `AuditAndRepairColumns`, `CheckIntegrity`, `MigrateDatabase`, `RepairDatabase`, and `MigrateAndRepairDatabase`.
  - All function bodies <= 15 lines conforming to repository guidelines.

### Modified Files

- `04-code/golang/pkg/applogger/sqlitelogger/models.go`:
  - Added `IsManualMigrationOnly bool` to `SplitDBConfig`.
  - Added `SchemaMigration` struct for migration records.
- `04-code/golang/pkg/applogger/sqlitelogger/manager.go`:
  - Integrated `MigrateAndRepairDatabase` into `openDbInternal` with modular helper functions (`openConnection`, `applyAutoMigration`).
  - Added explicit migration methods (`MigrateTaskDb`, `RepairTaskDb`, `MigrateAllTaskDbs`, `RepairAllTaskDbs`, `MigrateMainDb`, `RepairMainDb`).
  - Added `IsManualMigrationOnly` accessor and setter.
- `04-code/golang/pkg/applogger/sqlitelogger/sqlitelogger_test.go`:
  - Expanded test suite to 20 comprehensive unit tests covering 88.5% of statements.
  - Added tests for concurrent task writes, limits, incremental migrations, simulated execution/query errors, corrupt database detection, and validation errors.
- `05-changes-history/01-index.md`:
  - Registered Task 23 transaction log.

---

## 4. Architectural Decisions & Rationale

1. **Idempotent Column Audit via PRAGMA table_info:**
   Instead of destructive table recreation, `AuditAndRepairColumns` queries the live schema of the `logs` table and adds only missing columns via `ALTER TABLE`. This guarantees zero data loss when existing task databases are updated.
2. **Default-Enabled Auto-Migration with Manual Override:**
   `SplitDBConfig` uses `IsManualMigrationOnly bool` (positive naming, zero value `false`), meaning auto-migration and repair occur automatically on first access unless explicitly configured for manual control.
3. **Strict Bounded Function Lengths (<= 15 Lines):**
   Every function in `migration.go` and `manager.go` has a body <= 15 lines, with mandatory blank lines before returns and after closing braces.
4. **Defensive Validation Guards:**
   All functions in `migration.go` validate `db != nil` before issuing commands, preventing nil-pointer panics in edge cases.

---

## 5. Verification & Quality Gate Results

- **Go Tests:** 28/28 Go packages and examples passed (`go test ./pkg/... ./examples/... -count=1`).
- **Unit Test Coverage:** 88.5% statement coverage across `sqlitelogger` (20/20 test cases passing).
- **Code Formatter:** Passed cleanly (`python 03-ai-scripts/26-go-code-formatter.py`).
- **Go Preflight CI:** Passed 2/2 gates (`python 03-ai-scripts/28-go-preflight-ci.py`).
- **CI/CD Quality Gates:** Passed 36/36 gates (`python 03-ai-scripts/06-cicd-local-runner.py`).
- **Newline Styling:** Passed cleanly (`node linter-scripts/check-newline-styling.mjs`).

---

## 6. Next Steps / Hand-off Context

- The SQLite task database subsystem is fully self-repairing, self-migrating, and extensible for future schema versions.
