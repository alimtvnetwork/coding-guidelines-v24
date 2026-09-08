# Subtask 01: Task Database Retention Pruning & Dynamic Query Filtering

## 1. Goal
Add automated retention management (`PruneTasks`, `PruneTaskCount`) and dynamic SQL `WHERE` clause filtering (`level`, `startTime`, `endTime`, `offset`, `limit`) to `sqlitelogger`.

**Status:** ✅ Completed

## 2. Target Files
- `04-code/golang/pkg/applogger/sqlitelogger/manager.go`
- `04-code/golang/pkg/applogger/sqlitelogger/sqlitelogger_test.go`

## 3. Detailed Specifications
1. **Retention Pruning**:
   - `PruneTasks(maxAge time.Duration) (int, *appfault.AppError)`:
     - Scans `tasksDir` for all `.db` files.
     - For any file whose `FileInfo.ModTime()` is before `time.Now().Add(-maxAge)`:
       - Closes and removes cached connection from `m.taskDbs`.
       - Deletes `<file>.db`, `<file>.db-wal`, `<file>.db-shm`.
       - Increments pruned count.
   - `PruneTaskCount(maxDbs int) (int, *appfault.AppError)`:
     - If discovered `.db` count exceeds `maxDbs`, sorts files by modification time ascending (oldest first).
     - Prunes excess files until count <= `maxDbs`.
2. **Dynamic Query Filtering**:
   - In `queryLogs(db *sql.DB, filter FilterOptions)`:
     - Build dynamic query: `SELECT ... FROM logs WHERE 1=1`.
     - If `filter.Level != ""`: append ` AND level = ?`.
     - If `filter.StartTime != ""`: append ` AND timestamp >= ?`.
     - If `filter.EndTime != ""`: append ` AND timestamp <= ?`.
     - Append ` ORDER BY id DESC LIMIT ? OFFSET ?`.
     - Bind arguments in exact matching sequence.
3. **Coding Guidelines Compliance**:
   - Every function body <= 15 lines. Extract query building into `buildLogsQuery(filter FilterOptions) (string, []any)`.
   - Positive booleans, blank line before `return`.
