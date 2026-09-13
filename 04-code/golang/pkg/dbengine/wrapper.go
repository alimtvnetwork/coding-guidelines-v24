package dbengine

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// DbWrapper encapsulates a SQL database connection with dialect query compilation.
type DbWrapper struct {
	conn     *sql.DB
	dialect  DatabaseDialectType
	compiler DialectCompiler
}

// TxWrapper encapsulates an active SQL transaction.
type TxWrapper struct {
	tx       *sql.Tx
	compiler DialectCompiler
}

var sqlitePragmas = []string{
	"PRAGMA busy_timeout = 5000",
	"PRAGMA journal_mode = WAL",
	"PRAGMA synchronous = NORMAL",
	"PRAGMA foreign_keys = ON",
}

func applySQLitePragmas(conn *sql.DB) *appfault.AppError {
	for _, pragma := range sqlitePragmas {
		if _, err := conn.Exec(pragma); err != nil {
			_ = conn.Close()

			return appfault.NewAppBuilder(errtype.Database, "configure sqlite pragma: "+pragma).SetCause(err).Build()
		}
	}

	return nil
}

func configureConnForDialect(dialect DatabaseDialectType, conn *sql.DB) *appfault.AppError {
	if dialect != DbSQLite {
		return nil
	}

	conn.SetMaxOpenConns(1)

	return applySQLitePragmas(conn)
}

// OpenDb opens a database connection for the specified dialect.
func OpenDb(dialect DatabaseDialectType, dsn string) (*DbWrapper, *appfault.AppError) {
	conn, err := sql.Open(string(dialect), dsn)
	if err != nil {
		return nil, appfault.NewAppBuilder(errtype.Database, fmt.Sprintf("open database for dialect %s", dialect)).SetCause(err).Build()
	}

	if appErr := configureConnForDialect(dialect, conn); appErr != nil {
		return nil, appErr
	}

	return WrapDb(conn, dialect)
}

// WrapDb wraps an existing sql.DB connection with dialect compilation.
func WrapDb(conn *sql.DB, dialect DatabaseDialectType) (*DbWrapper, *appfault.AppError) {
	compiler, appErr := ResolveCompiler(dialect)
	if appErr != nil {
		return nil, appErr
	}

	return &DbWrapper{
		conn:     conn,
		dialect:  dialect,
		compiler: compiler,
	}, nil
}

// Close closes the underlying connection.
func (w *DbWrapper) Close() *appfault.AppError {
	if w.conn == nil {
		return nil
	}

	err := w.conn.Close()
	if err != nil {
		return appfault.NewAppBuilder(errtype.Database, "close database connection").SetCause(err).Build()
	}

	return nil
}

// QueryRow executes a query that is expected to return at most one row.
func (w *DbWrapper) QueryRow(ctx context.Context, query string, args ...any) (*sql.Row, *appfault.AppError) {
	row := w.conn.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, appfault.NewAppBuilder(errtype.Database, "execute query row: "+query).SetCause(row.Err()).Build()
	}

	return row, nil
}

// Query executes a query that returns rows.
func (w *DbWrapper) Query(ctx context.Context, query string, args ...any) (*sql.Rows, *appfault.AppError) {
	rows, err := w.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, appfault.NewAppBuilder(errtype.Database, "execute query: "+query).SetCause(err).Build()
	}

	return rows, nil
}

// Exec executes a query without returning rows.
func (w *DbWrapper) Exec(ctx context.Context, query string, args ...any) (sql.Result, *appfault.AppError) {
	result, err := w.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, appfault.NewAppBuilder(errtype.Database, "execute exec: "+query).SetCause(err).Build()
	}

	return result, nil
}

// Prepare creates a prepared statement on the database connection.
func (w *DbWrapper) Prepare(ctx context.Context, query string) (*sql.Stmt, *appfault.AppError) {
	stmt, err := w.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, appfault.NewAppBuilder(errtype.Database, "prepare statement: "+query).SetCause(err).Build()
	}

	return stmt, nil
}

func handleRollback(tx *sql.Tx, txErr *appfault.AppError) *appfault.AppError {
	rbErr := tx.Rollback()
	if rbErr == nil {
		return txErr
	}

	return appfault.NewAppBuilder(errtype.Database, "rollback transaction").SetCause(rbErr).SetContext("cause_error", txErr.Error()).Build()
}

var (
	immediateTxOptions = &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	}

	readOnlyTxOptions = &sql.TxOptions{
		ReadOnly: true,
	}

	exclusiveTxOptions = &sql.TxOptions{
		Isolation: sql.LevelLinearizable,
	}
)

func checkRollbackError(p any, rbErr error) {
	if !errors.Is(rbErr, sql.ErrTxDone) {
		panic(fmt.Sprintf("%v (rollback failed: %v)", p, rbErr)) // lint-allow: panic
	}
}

func recoverTxPanic(tx *sql.Tx) {
	p := recover()
	if p == nil {
		return
	}

	rbErr := tx.Rollback()
	if rbErr != nil {
		checkRollbackError(p, rbErr)
	}

	panic(p) // lint-allow: panic
}

func (w *DbWrapper) executeTx(
	ctx context.Context,
	opts *sql.TxOptions,
	fn func(tx *TxWrapper) *appfault.AppError,
) *appfault.AppError {
	tx, err := w.conn.BeginTx(ctx, opts)
	if err != nil {
		return appfault.NewAppBuilder(errtype.Database, "begin transaction").SetCause(err).Build()
	}

	defer recoverTxPanic(tx)

	return w.runTxFunc(tx, fn)
}

func (w *DbWrapper) runTxFunc(tx *sql.Tx, fn func(tx *TxWrapper) *appfault.AppError) *appfault.AppError {
	txWrap := &TxWrapper{tx: tx, compiler: w.compiler}
	if txErr := fn(txWrap); txErr != nil {
		return handleRollback(tx, txErr)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return appfault.NewAppBuilder(errtype.Database, "commit transaction").SetCause(commitErr).Build()
	}

	return nil
}

// WithTransaction runs a function within a database transaction.
func (w *DbWrapper) WithTransaction(ctx context.Context, fn func(tx *TxWrapper) *appfault.AppError) *appfault.AppError {
	return w.executeTx(ctx, nil, fn)
}

// WithImmediateTransaction runs a function within an immediate database transaction.
func (w *DbWrapper) WithImmediateTransaction(ctx context.Context, fn func(tx *TxWrapper) *appfault.AppError) *appfault.AppError {
	return w.executeTx(ctx, immediateTxOptions, fn)
}

// WithTxOptions runs a function within a database transaction using the provided options.
func (w *DbWrapper) WithTxOptions(
	ctx context.Context,
	opts *sql.TxOptions,
	fn func(tx *TxWrapper) *appfault.AppError,
) *appfault.AppError {
	return w.executeTx(ctx, opts, fn)
}

// WithReadOnlyTransaction runs a function within a read-only database transaction.
func (w *DbWrapper) WithReadOnlyTransaction(
	ctx context.Context,
	fn func(tx *TxWrapper) *appfault.AppError,
) *appfault.AppError {
	return w.executeTx(ctx, readOnlyTxOptions, fn)
}

// WithExclusiveTransaction runs a function within an exclusive database transaction.
func (w *DbWrapper) WithExclusiveTransaction(
	ctx context.Context,
	fn func(tx *TxWrapper) *appfault.AppError,
) *appfault.AppError {
	return w.executeTx(ctx, exclusiveTxOptions, fn)
}

// Compiler returns the active DialectCompiler.
func (w *DbWrapper) Compiler() DialectCompiler {
	return w.compiler
}

// Dialect returns the active DatabaseDialectType.
func (w *DbWrapper) Dialect() DatabaseDialectType {
	return w.dialect
}

// Conn returns the underlying *sql.DB connection.
func (w *DbWrapper) Conn() *sql.DB {
	return w.conn
}

// CreateView creates a database view with the specified SELECT statement.
func (w *DbWrapper) CreateView(ctx context.Context, name string, selectSql string) BoolResult {
	query := w.compiler.CompileCreateView(name, selectSql)
	_, err := w.Exec(ctx, query)
	if err != nil {
		return FailureBool(err)
	}

	return SuccessBool(true)
}

// DropView drops a database view.
func (w *DbWrapper) DropView(ctx context.Context, name string) BoolResult {
	query := w.compiler.CompileDropView(name)
	_, err := w.Exec(ctx, query)
	if err != nil {
		return FailureBool(err)
	}

	return SuccessBool(true)
}

// CallFunction executes a database function and returns the scalar string result.
func (w *DbWrapper) CallFunction(ctx context.Context, name string, args ...any) StringResult {
	query := w.compiler.CompileFunctionCall(name, len(args))
	row, appErr := w.QueryRow(ctx, query, args...)
	if appErr != nil {
		return FailureString(appErr)
	}

	var res string
	if err := row.Scan(&res); err != nil {
		return FailureString(appfault.NewAppBuilder(errtype.Database, "scan result of function "+name).SetCause(err).Build())
	}

	return SuccessString(res)
}

// ExecRowsAffected executes a statement and returns the number of rows affected wrapped in RowsAffectedResult.
func (w *DbWrapper) ExecRowsAffected(ctx context.Context, query string, args ...any) RowsAffectedResult {
	res, appErr := w.Exec(ctx, query, args...)
	if appErr != nil {
		return FailureRowsAffected(appErr)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return FailureRowsAffected(appfault.NewAppBuilder(errtype.Database, "get rows affected").SetCause(err).Build())
	}

	return SuccessRowsAffected(affected)
}

// ViewExists checks whether a database view exists.
func (w *DbWrapper) ViewExists(ctx context.Context, name string) (bool, *appfault.AppError) {
	query := w.compiler.CompileInspectViewExists(name)
	if len(query) == 0 {
		return false, nil
	}

	row, appErr := w.QueryRow(ctx, query, name)
	if appErr != nil {
		return false, appErr
	}

	var dummy int
	err := row.Scan(&dummy)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return false, appfault.NewAppBuilder(errtype.Database, "check view existence "+name).SetCause(err).Build()
}

// GetTableColumns returns column names for a table or view.
func (w *DbWrapper) GetTableColumns(ctx context.Context, name string) ([]string, *appfault.AppError) {
	query := w.compiler.CompileInspectColumns(name)
	if len(query) == 0 {
		return nil, nil
	}

	if w.dialect == DbSQLite {
		return w.scanSqliteColumns(ctx, query)
	}

	return w.scanStandardColumns(ctx, query, name)
}

func (w *DbWrapper) scanSqliteColumns(ctx context.Context, query string) ([]string, *appfault.AppError) {
	rows, appErr := w.Query(ctx, query)
	if appErr != nil {
		return nil, appErr
	}

	defer rows.Close()

	var cols []string
	for rows.Next() {
		var (
			cid     int
			colName string
			ctype   string
			notnull int
			dflt    any
			pk      int
		)
		if err := rows.Scan(&cid, &colName, &ctype, &notnull, &dflt, &pk); err != nil {
			return nil, appfault.NewAppBuilder(errtype.Database, "scan sqlite column info").SetCause(err).Build()
		}

		cols = append(cols, colName)
	}

	return cols, nil
}

func (w *DbWrapper) scanStandardColumns(ctx context.Context, query string, name string) ([]string, *appfault.AppError) {
	rows, appErr := w.Query(ctx, query, name)
	if appErr != nil {
		return nil, appErr
	}

	defer rows.Close()

	var cols []string
	for rows.Next() {
		var colName string
		if err := rows.Scan(&colName); err != nil {
			return nil, appfault.NewAppBuilder(errtype.Database, "scan column info for "+name).SetCause(err).Build()
		}

		cols = append(cols, colName)
	}

	return cols, nil
}

func (w *DbWrapper) verifyViewColumns(ctx context.Context, name string, requiredColumns []string) (bool, *appfault.AppError) {
	existingCols, appErr := w.GetTableColumns(ctx, name)
	if appErr != nil {
		return false, appErr
	}

	colMap := make(map[string]bool, len(existingCols))
	for _, col := range existingCols {
		colMap[col] = true
	}

	for _, req := range requiredColumns {
		if !colMap[req] {
			return false, nil
		}
	}

	return true, nil
}

const sqlCreateViewMeta = `
CREATE TABLE IF NOT EXISTS __dbengine_view_meta (
    ViewName TEXT PRIMARY KEY,
    QueryHash TEXT NOT NULL,
    ViewSql TEXT NOT NULL,
    UpdatedAt TEXT NOT NULL
);`

// ComputeSqlHash computes a deterministic SHA-256 hex string for a given SQL statement.
func ComputeSqlHash(sqlStr string) string {
	sum := sha256.Sum256([]byte(strings.TrimSpace(sqlStr)))

	return hex.EncodeToString(sum[:])
}

// ValidateSql performs a dry-run syntax and reference check using EXPLAIN on the database connection.
func (w *DbWrapper) ValidateSql(ctx context.Context, sqlStr string) *appfault.AppError {
	cleanSql := strings.TrimRight(strings.TrimSpace(sqlStr), ";")
	explainQuery := fmt.Sprintf("EXPLAIN %s", cleanSql)
	rows, err := w.conn.QueryContext(ctx, explainQuery)
	if err != nil {
		return appfault.NewAppBuilder(errtype.Database, "validate sql query syntax").SetCause(err).Build()
	}

	defer rows.Close()

	return nil
}

// EnsureViewMetaTable initializes the view metadata table if it does not already exist.
func (w *DbWrapper) EnsureViewMetaTable(ctx context.Context) *appfault.AppError {
	_, err := w.Exec(ctx, sqlCreateViewMeta)
	if err != nil {
		return appfault.NewAppBuilder(errtype.Database, "ensure view metadata table").SetCause(err).Build()
	}

	return nil
}

// GetViewHash retrieves the recorded query hash for a view from __dbengine_view_meta.
func (w *DbWrapper) GetViewHash(ctx context.Context, name string) (string, *appfault.AppError) {
	query := "SELECT QueryHash FROM __dbengine_view_meta WHERE ViewName = ?"
	row, appErr := w.QueryRow(ctx, query, name)
	if appErr != nil {
		return "", appErr
	}

	var hash string
	err := row.Scan(&hash)
	if err == nil {
		return hash, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}

	return "", appfault.NewAppBuilder(errtype.Database, "scan view hash for "+name).SetCause(err).Build()
}

// SaveViewMeta records or updates the view hash and SQL in __dbengine_view_meta.
func (w *DbWrapper) SaveViewMeta(ctx context.Context, name, queryHash, viewSql string) *appfault.AppError {
	now := time.Now().UTC().Format(time.RFC3339)
	query := `
INSERT OR REPLACE INTO __dbengine_view_meta (ViewName, QueryHash, ViewSql, UpdatedAt)
VALUES (?, ?, ?, ?);`
	_, appErr := w.Exec(ctx, query, name, queryHash, viewSql, now)
	if appErr != nil {
		return appfault.NewAppBuilder(errtype.Database, "save view metadata for "+name).SetCause(appErr).Build()
	}

	return nil
}

func (w *DbWrapper) isViewCurrent(ctx context.Context, name, queryHash string) (bool, *appfault.AppError) {
	existingHash, appErr := w.GetViewHash(ctx, name)
	if appErr != nil {
		return false, appErr
	}

	if len(existingHash) == 0 {
		return false, nil
	}

	return existingHash == queryHash, nil
}

func (w *DbWrapper) recordViewMetaIfPresent(ctx context.Context, name, queryHash, selectSql string) BoolResult {
	if len(queryHash) == 0 {
		return SuccessBool(true)
	}

	saveErr := w.SaveViewMeta(ctx, name, queryHash, selectSql)
	if saveErr != nil {
		return FailureBool(saveErr)
	}

	return SuccessBool(true)
}

func (w *DbWrapper) createAndRegisterView(ctx context.Context, name, selectSql, queryHash string) BoolResult {
	createRes := w.CreateView(ctx, name, selectSql)
	if createRes.IsFailed() {
		return createRes
	}

	return w.recordViewMetaIfPresent(ctx, name, queryHash, selectSql)
}

func (w *DbWrapper) dropAndRecreateView(ctx context.Context, name, selectSql, queryHash string) BoolResult {
	dropRes := w.DropView(ctx, name)
	if dropRes.IsFailed() {
		return dropRes
	}

	return w.createAndRegisterView(ctx, name, selectSql, queryHash)
}

func (w *DbWrapper) recreateAndRegisterView(ctx context.Context, name string, selectSql string, queryHash string) BoolResult {
	valErr := w.ValidateSql(ctx, selectSql)
	if valErr != nil {
		return FailureBool(valErr)
	}

	exists, appErr := w.ViewExists(ctx, name)
	if appErr != nil {
		return FailureBool(appErr)
	}

	if exists {
		return w.dropAndRecreateView(ctx, name, selectSql, queryHash)
	}

	return w.createAndRegisterView(ctx, name, selectSql, queryHash)
}

// CreateViewOrUseViewWithHash checks if a view exists and matches queryHash in __dbengine_view_meta.
func (w *DbWrapper) CreateViewOrUseViewWithHash(ctx context.Context, name string, selectSql string, queryHash string) BoolResult {
	metaErr := w.EnsureViewMetaTable(ctx)
	if metaErr != nil {
		return FailureBool(metaErr)
	}

	exists, appErr := w.ViewExists(ctx, name)
	if appErr != nil {
		return FailureBool(appErr)
	}

	if !exists {
		return w.recreateAndRegisterView(ctx, name, selectSql, queryHash)
	}

	current, currentErr := w.isViewCurrent(ctx, name, queryHash)
	if currentErr != nil {
		return FailureBool(currentErr)
	}

	if current {
		return SuccessBool(true)
	}

	return w.recreateAndRegisterView(ctx, name, selectSql, queryHash)
}

// CreateViewOrUseView inspects if a view exists and contains the required columns or matches query hash.
func (w *DbWrapper) CreateViewOrUseView(ctx context.Context, name string, selectSql string, requiredColumns ...string) BoolResult {
	hash := ComputeSqlHash(selectSql)
	if len(requiredColumns) == 0 {
		return w.CreateViewOrUseViewWithHash(ctx, name, selectSql, hash)
	}

	metaErr := w.EnsureViewMetaTable(ctx)
	if metaErr != nil {
		return FailureBool(metaErr)
	}

	exists, appErr := w.ViewExists(ctx, name)
	if appErr != nil {
		return FailureBool(appErr)
	}

	if !exists {
		return w.recreateAndRegisterView(ctx, name, selectSql, hash)
	}

	hasAll, verifyErr := w.verifyViewColumns(ctx, name, requiredColumns)
	if verifyErr != nil {
		return FailureBool(verifyErr)
	}

	if !hasAll {
		return w.recreateAndRegisterView(ctx, name, selectSql, hash)
	}

	saveRes := w.recordViewMetaIfPresent(ctx, name, hash, selectSql)
	if saveRes.IsFailed() {
		return saveRes
	}

	return SuccessBool(true)
}
