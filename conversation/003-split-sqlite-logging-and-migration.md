# Split SQLite Logging and Database Migration

Sequence: 003
CapturedUtc: 2026-09-08T19:45:00Z
Span: 5 user prompts
Topic: Split SQLite log architecture, configurable task paths, automatic schema migration and column repair

---

## User Instructions (verbatim)

### 1.

> D:\wp-work\riseup-asia\03-aukgo\errorwrapper\errcmd
> 
> 
> Okay, so for the, um, SQLite DB on the logging part, there could be two different things. One, that we could, um, we, we could actually have task by task logging, uh, with a separate DB. That means there could be a main DB just like the split DB concept. So, uh, the main DB about the log that could have the logs DB or that could also be changeable from the loggers. Default one would be there. So inside where it is keeping, so there should be a work directory or DB directory. So in that DB directory, uh, it should have a logs.db. Uh, but also inside that, you can have a task DB, tasks folder. Inside these folders, it can have each task by task logs DB, which can be retrieved by, uh, API methods, and that means there should be methods to retrieve all of these. Um, and also, you should look into the split DB concept, uh, so that you understand how it's going to work. And also rotation, uh, for the text-based logging, we need to have the rotation for--

### 2.

> make sure abd cnfirm if these paths are chanbale
> 
> <workDir>/tasks/<

### 3.

> when system is working iwth task db it should fix or mirgate that task db as well

### 4.

> can you please add more tests and ensure everything function is working as expected

### 5.

> can you please add more tests and ensure everything function is working as expected

---

## Assistant Actions Summary (one bullet per turn, no chain-of-thought)

- Designed and implemented `SplitDBManager` in `04-code/golang/pkg/applogger/sqlitelogger/manager.go` supporting global `logs.db` and isolated per-task `<taskId>.db` databases in `<workDir>/tasks/` with lazy connection pooling.
- Verified and implemented runtime path configurability via `SetMainDbPath`, `SetTasksDir`, `SetTaskDbPath`, and dynamic `TaskDbPathResolverFunc`.
- Built automated schema migration and repair engine in `04-code/golang/pkg/applogger/sqlitelogger/migration.go` ensuring PRAGMA integrity verification, schema version tracking, idempotent index creation, and automatic `ALTER TABLE logs ADD COLUMN` column repair without data loss.
- Expanded the `sqlitelogger` unit test suite to 20 comprehensive tests covering 88.5% statement coverage, verifying concurrent writes, limits, corrupted DB handling, and incremental migrations.

---

## Outcomes / Decisions

- Adopted split DB architecture for logging: global log database for shared lifecycle telemetry and per-task isolated databases for fine-grained task diagnostics.
- Added automatic schema migration and column repair on open/write calls with optional manual-only toggle (`IsManualMigrationOnly`).
- Documented in `05-changes-history/22-split-sqlite-logging-rotating-lazyonce-errcmd/01-transaction-log.md` and `05-changes-history/23-sqlite-task-db-auto-migration-and-repair/01-transaction-log.md`.

## Open Threads (carry-over)

- User inquired about further improvements across the codebase and `pkg/` folder.
