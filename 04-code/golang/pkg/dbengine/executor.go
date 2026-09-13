package dbengine

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// SqlExecuter provides a unified interface for executing queries across database connections and transactions.
type SqlExecuter interface {
	QueryRow(ctx context.Context, query string, args ...any) (*sql.Row, *appfault.AppError)
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, *appfault.AppError)
	Exec(ctx context.Context, query string, args ...any) (sql.Result, *appfault.AppError)
	ExecRowsAffected(ctx context.Context, query string, args ...any) RowsAffectedResult
	Compiler() DialectCompiler
}

// SqlExecutor is an alias to SqlExecuter for conventional spelling.
type SqlExecutor = SqlExecuter

var (
	_ SqlExecutor = (*DbWrapper)(nil)
	_ SqlExecutor = (*TxWrapper)(nil)
)

// Tx returns the underlying *sql.Tx transaction.
func (t *TxWrapper) Tx() *sql.Tx {
	return t.tx
}

// Compiler returns the active DialectCompiler.
func (t *TxWrapper) Compiler() DialectCompiler {
	return t.compiler
}

// QueryRow executes a query within the transaction expecting at most one row.
func (t *TxWrapper) QueryRow(ctx context.Context, query string, args ...any) (*sql.Row, *appfault.AppError) {
	row := t.tx.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, appfault.Wrap(errtype.Database, row.Err(), "execute tx query row: "+query)
	}

	return row, nil
}

// Query executes a query within the transaction returning rows.
func (t *TxWrapper) Query(ctx context.Context, query string, args ...any) (*sql.Rows, *appfault.AppError) {
	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "execute tx query: "+query)
	}

	return rows, nil
}

// Exec executes a query within the transaction without returning rows.
func (t *TxWrapper) Exec(ctx context.Context, query string, args ...any) (sql.Result, *appfault.AppError) {
	result, err := t.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "execute tx exec: "+query)
	}

	return result, nil
}

// ExecRowsAffected executes a statement within the transaction and returns rows affected.
func (t *TxWrapper) ExecRowsAffected(ctx context.Context, query string, args ...any) RowsAffectedResult {
	res, appErr := t.Exec(ctx, query, args...)
	if appErr != nil {
		return FailureRowsAffected(appErr)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return FailureRowsAffected(appfault.Wrap(errtype.Database, err, "get tx rows affected"))
	}

	return SuccessRowsAffected(affected)
}

// Prepare creates a prepared statement within the transaction.
func (t *TxWrapper) Prepare(ctx context.Context, query string) (*sql.Stmt, *appfault.AppError) {
	stmt, err := t.tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, appfault.Wrap(errtype.Database, err, "prepare tx statement: "+query)
	}

	return stmt, nil
}

// ValidateSql performs a dry-run syntax and reference check using EXPLAIN within the transaction.
func (t *TxWrapper) ValidateSql(ctx context.Context, sqlStr string) *appfault.AppError {
	cleanSql := strings.TrimRight(strings.TrimSpace(sqlStr), ";")
	explainQuery := fmt.Sprintf("EXPLAIN %s", cleanSql)
	rows, appErr := t.Query(ctx, explainQuery)
	if appErr != nil {
		return appErr
	}

	defer rows.Close()

	return nil
}

// CallFunction executes a database function within the transaction and returns the scalar string result.
func (t *TxWrapper) CallFunction(ctx context.Context, name string, args ...any) StringResult {
	query := t.compiler.CompileFunctionCall(name, len(args))
	row, appErr := t.QueryRow(ctx, query, args...)
	if appErr != nil {
		return FailureString(appErr)
	}

	var res string
	if err := row.Scan(&res); err != nil {
		return FailureString(appfault.Wrap(errtype.Database, err, "scan result of tx function "+name))
	}

	return SuccessString(res)
}
