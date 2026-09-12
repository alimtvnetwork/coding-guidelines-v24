package sqlitelogger

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

const (
	// CurrentSchemaVersion defines the target SQLite schema version.
	CurrentSchemaVersion = 2

	createMigrationsTableSql = `CREATE TABLE IF NOT EXISTS schema_migrations (
		version INTEGER PRIMARY KEY,
		description TEXT,
		applied_at TEXT
	);`
)

// ColumnDef defines a database column name and its SQLite type.
type ColumnDef struct {
	Name string
	Type string
}

var requiredColumns = []ColumnDef{
	{Name: "id", Type: "INTEGER"},
	{Name: "task_id", Type: "TEXT"},
	{Name: "timestamp", Type: "TEXT"},
	{Name: "level", Type: "TEXT"},
	{Name: "message", Type: "TEXT"},
	{Name: "caller", Type: "TEXT"},
	{Name: "fields_json", Type: "TEXT"},
	{Name: "stack_trace", Type: "TEXT"},
	{Name: "duration_ms", Type: "INTEGER"},
	{Name: "status", Type: "TEXT"},
}

var requiredIndexes = []string{
	"CREATE INDEX IF NOT EXISTS idx_logs_task_id ON logs(task_id);",
	"CREATE INDEX IF NOT EXISTS idx_logs_timestamp ON logs(timestamp);",
	"CREATE INDEX IF NOT EXISTS idx_logs_level ON logs(level);",
	"CREATE INDEX IF NOT EXISTS idx_logs_status ON logs(status);",
}

// EnsureBaseSchema guarantees presence of logs and schema_migrations tables.
func EnsureBaseSchema(db *sql.DB) *appfault.AppError {
	if db == nil {
		return appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	if _, err := db.Exec(createTableSql); err != nil {
		return appfault.Wrap(errtype.Database, err, "failed to create logs table")
	}

	if _, err := db.Exec(createMigrationsTableSql); err != nil {
		return appfault.Wrap(errtype.Database, err, "failed to create schema_migrations table")
	}

	return nil
}

// GetCurrentSchemaVersion queries the latest version registered in schema_migrations.
func GetCurrentSchemaVersion(db *sql.DB) (int, *appfault.AppError) {
	if db == nil {
		return 0, appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	var version sql.NullInt64
	row := db.QueryRow("SELECT MAX(version) FROM schema_migrations")
	err := row.Scan(&version)
	isNoRows := err == sql.ErrNoRows
	if isNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, appfault.Wrap(errtype.Database, err, "failed to query schema version")
	}

	isValid := version.Valid
	if !isValid {
		return 0, nil
	}

	return int(version.Int64), nil
}

// RecordMigrationVersion records an executed schema upgrade in schema_migrations.
func RecordMigrationVersion(db *sql.DB, version int, description string) *appfault.AppError {
	if db == nil {
		return appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	stmt := "INSERT OR REPLACE INTO schema_migrations (version, description, applied_at) VALUES (?, ?, ?)"
	now := time.Now().UTC().Format(time.RFC3339)

	_, err := db.Exec(stmt, version, description, now)
	if err != nil {
		msg := fmt.Sprintf("failed to record migration v%d", version)

		return appfault.Wrap(errtype.Database, err, msg)
	}

	return nil
}

// ApplyIndexes creates all performance indexes for log queries and diagnostics.
func ApplyIndexes(db *sql.DB) *appfault.AppError {
	if db == nil {
		return appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	for _, idxSql := range requiredIndexes {
		if _, err := db.Exec(idxSql); err != nil {
			return appfault.Wrap(errtype.Database, err, "failed to apply index")
		}
	}

	return nil
}

// QueryTableColumns fetches discovered column names for a given SQLite table.
func QueryTableColumns(db *sql.DB, tableName string) (map[string]bool, *appfault.AppError) {
	if db == nil {
		return nil, appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	query := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "failed to query table info")
	}

	defer rows.Close()

	return scanColumnNames(rows)
}

// scanColumnNames extracts column names from PRAGMA table_info rows.
func scanColumnNames(rows *sql.Rows) (map[string]bool, *appfault.AppError) {
	cols := make(map[string]bool)
	for rows.Next() {
		var cid, notNull, pk int
		var name, colType string
		var dfltVal sql.NullString

		if err := rows.Scan(&cid, &name, &colType, &notNull, &dfltVal, &pk); err != nil {
			return nil, appfault.Wrap(errtype.Database, err, "failed to scan column info")
		}

		cols[strings.ToLower(name)] = true
	}

	return cols, nil
}

// AddMissingColumn issues an ALTER TABLE command to add a missing column.
func AddMissingColumn(db *sql.DB, tableName, colName, colType string) *appfault.AppError {
	if db == nil {
		return appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	alterSql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", tableName, colName, colType)
	if _, err := db.Exec(alterSql); err != nil {
		msg := fmt.Sprintf("failed to add column %s", colName)

		return appfault.Wrap(errtype.Database, err, msg)
	}

	return nil
}

// AuditAndRepairColumns ensures all required columns exist in the target table.
func AuditAndRepairColumns(db *sql.DB, tableName string) *appfault.AppError {
	if db == nil {
		return appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	cols, fault := QueryTableColumns(db, tableName)
	if fault != nil {
		return fault
	}

	hasCols := len(cols) > 0
	if !hasCols {
		return nil
	}

	return applyMissingColumns(db, tableName, cols)
}

// applyMissingColumns compares discovered columns against required specifications.
func applyMissingColumns(db *sql.DB, tableName string, cols map[string]bool) *appfault.AppError {
	for _, col := range requiredColumns {
		hasCol := cols[strings.ToLower(col.Name)]
		if hasCol {
			continue
		}

		if fault := AddMissingColumn(db, tableName, col.Name, col.Type); fault != nil {
			return fault
		}
	}

	return nil
}

// CheckIntegrity runs a quick integrity check on the SQLite database.
func CheckIntegrity(db *sql.DB) *appfault.AppError {
	if db == nil {
		return appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	var checkResult string
	row := db.QueryRow("PRAGMA quick_check")
	if err := row.Scan(&checkResult); err != nil {
		return handleIntegrityScanError(err)
	}

	return validateIntegrityResult(checkResult)
}

func handleIntegrityScanError(err error) *appfault.AppError {
	if err == sql.ErrNoRows {
		return nil
	}

	return appfault.Wrap(errtype.Database, err, "failed to run integrity check")
}

func validateIntegrityResult(res string) *appfault.AppError {
	isHealthy := strings.ToLower(res) == "ok"
	if !isHealthy {
		return appfault.New(errtype.Database, fmt.Sprintf("database integrity failure: %s", res))
	}

	return nil
}

// MigrateDatabase executes pending schema migrations sequentially.
func MigrateDatabase(db *sql.DB) *appfault.AppError {
	if fault := EnsureBaseSchema(db); fault != nil {
		return fault
	}

	ver, fault := GetCurrentSchemaVersion(db)
	if fault != nil {
		return fault
	}

	return applyPendingMigrations(db, ver)
}

// applyMigrationV1 records the base logs table migration.
func applyMigrationV1(db *sql.DB) *appfault.AppError {
	return RecordMigrationVersion(db, 1, "initialize base logs table")
}

// applyMigrationV2 creates indexes and records version 2.
func applyMigrationV2(db *sql.DB) *appfault.AppError {
	fault := ApplyIndexes(db)
	if fault != nil {
		return fault
	}

	return RecordMigrationVersion(db, 2, "create query indexes")
}

// applyMigrationV1ThenV2 applies migration 1 and then migration 2 sequentially.
func applyMigrationV1ThenV2(db *sql.DB) *appfault.AppError {
	fault := applyMigrationV1(db)
	if fault != nil {
		return fault
	}

	return applyMigrationV2(db)
}

// applyPendingMigrations steps through unapplied version increments.
func applyPendingMigrations(db *sql.DB, currentVersion int) *appfault.AppError {
	hasV1 := currentVersion >= 1
	if !hasV1 {
		return applyMigrationV1ThenV2(db)
	}

	hasV2 := currentVersion >= 2
	if !hasV2 {
		return applyMigrationV2(db)
	}

	return nil
}

// RepairDatabase validates integrity, recreates tables, adds missing columns and indexes.
func RepairDatabase(db *sql.DB) *appfault.AppError {
	if db == nil {
		return appfault.New(errtype.Validation, "database connection cannot be nil")
	}

	_ = CheckIntegrity(db)

	if fault := EnsureBaseSchema(db); fault != nil {
		return fault
	}

	if fault := AuditAndRepairColumns(db, "logs"); fault != nil {
		return fault
	}

	if fault := ApplyIndexes(db); fault != nil {
		return fault
	}

	return MigrateDatabase(db)
}

// MigrateAndRepairDatabase runs both repair and migration cycles to guarantee schema correctness.
func MigrateAndRepairDatabase(db *sql.DB) *appfault.AppError {
	if fault := RepairDatabase(db); fault != nil {
		return fault
	}

	return MigrateDatabase(db)
}
