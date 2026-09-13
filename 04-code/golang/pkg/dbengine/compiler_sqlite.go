package dbengine

import (
	"fmt"
	"strings"
)

// SQLiteCompiler compiles queries for SQLite databases.
type SQLiteCompiler struct{}

// Dialect returns DatabaseDialectSQLite.
func (c *SQLiteCompiler) Dialect() DatabaseDialectType {
	return DatabaseDialectSQLite
}

// Placeholder returns ? for SQLite parameters.
func (c *SQLiteCompiler) Placeholder(paramIndex int) string {
	return "?"
}

// QuoteIdentifier wraps column/table names in double quotes.
func (c *SQLiteCompiler) QuoteIdentifier(name string) string {
	return "\"" + name + "\""
}

// CompilePagination produces LIMIT/OFFSET syntax.
func (c *SQLiteCompiler) CompilePagination(limit, offset int) string {
	if limit <= 0 && offset <= 0 {
		return ""
	}

	if offset <= 0 {
		return fmt.Sprintf("LIMIT %d", limit)
	}

	return fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
}

// CompileLocate builds an INSTR(field, ?) > 0 clause for substring location.
func (c *SQLiteCompiler) CompileLocate(field string) string {
	return fmt.Sprintf("INSTR(%s, ?) > 0", c.QuoteIdentifier(field))
}

// CompileSearch builds a parameterized SELECT query.
func (c *SQLiteCompiler) CompileSearch(table string, fields []string, limit int) string {
	quotedTable := c.QuoteIdentifier(table)
	if len(fields) == 0 {
		return buildSelectWithoutWhere(quotedTable, c.CompilePagination(limit, 0))
	}

	whereClauses := make([]string, 0, len(fields))
	for _, f := range fields {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", c.QuoteIdentifier(f)))
	}

	whereSql := strings.Join(whereClauses, " AND ")
	pagination := c.CompilePagination(limit, 0)

	return buildSelectWithWhere(quotedTable, whereSql, pagination)
}

// CompileCreateView builds a CREATE VIEW IF NOT EXISTS statement.
func (c *SQLiteCompiler) CompileCreateView(name string, selectSql string) string {
	cleanSql := strings.TrimRight(strings.TrimSpace(selectSql), ";")

	return fmt.Sprintf("CREATE VIEW IF NOT EXISTS %s AS %s;", c.QuoteIdentifier(name), cleanSql)
}

// CompileDropView builds a DROP VIEW IF EXISTS statement.
func (c *SQLiteCompiler) CompileDropView(name string) string {
	return fmt.Sprintf("DROP VIEW IF EXISTS %s;", c.QuoteIdentifier(name))
}

// CompileFunctionCall builds a SELECT <func>(args...) query.
func (c *SQLiteCompiler) CompileFunctionCall(name string, argCount int) string {
	if argCount <= 0 {
		return fmt.Sprintf("SELECT %s();", name)
	}

	placeholders := make([]string, argCount)
	for i := 0; i < argCount; i++ {
		placeholders[i] = "?"
	}

	return fmt.Sprintf("SELECT %s(%s);", name, strings.Join(placeholders, ", "))
}

// CompileCount builds a SELECT COUNT(*) query.
func (c *SQLiteCompiler) CompileCount(table string, field string) string {
	quotedTable := c.QuoteIdentifier(table)
	if len(field) == 0 {
		return fmt.Sprintf("SELECT COUNT(*) FROM %s;", quotedTable)
	}

	return fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?;", quotedTable, c.QuoteIdentifier(field))
}

// CompileDelete builds a DELETE query.
func (c *SQLiteCompiler) CompileDelete(table string, field string) string {
	quotedTable := c.QuoteIdentifier(table)

	return fmt.Sprintf("DELETE FROM %s WHERE %s = ?;", quotedTable, c.QuoteIdentifier(field))
}

// CompileInspectColumns returns the PRAGMA query to inspect columns of a table or view.
func (c *SQLiteCompiler) CompileInspectColumns(tableOrView string) string {
	return fmt.Sprintf("PRAGMA table_info(%s);", c.QuoteIdentifier(tableOrView))
}

// CompileInspectViewExists returns a query to check if a view exists in sqlite_master.
func (c *SQLiteCompiler) CompileInspectViewExists(viewName string) string {
	return "SELECT 1 FROM sqlite_master WHERE type = 'view' AND name = ?;"
}

func buildSelectWithoutWhere(quotedTable, pagination string) string {
	if len(pagination) == 0 {
		return fmt.Sprintf("SELECT * FROM %s;", quotedTable)
	}

	return fmt.Sprintf("SELECT * FROM %s %s;", quotedTable, pagination)
}

func buildSelectWithWhere(quotedTable, whereSql, pagination string) string {
	if len(pagination) == 0 {
		return fmt.Sprintf("SELECT * FROM %s WHERE %s;", quotedTable, whereSql)
	}

	return fmt.Sprintf("SELECT * FROM %s WHERE %s %s;", quotedTable, whereSql, pagination)
}
