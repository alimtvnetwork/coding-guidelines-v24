package dbengine

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

type mockRowScanner struct {
	values []any
	err    error
}

func (m *mockRowScanner) Scan(dest ...any) error {
	if m.err != nil {
		return m.err
	}

	for i, d := range dest {
		if i >= len(m.values) {
			break
		}

		switch target := d.(type) {
		case *uint64:
			*target = ScanUint64(m.values[i])
		case *string:
			*target = ScanString(m.values[i])
		case *bool:
			*target = ScanBool(m.values[i])
		case *int64:
			*target = ScanInt64(m.values[i])
		case *int:
			*target = ScanInt(m.values[i])
		default:
			return fmt.Errorf("unsupported destination type in mockRowScanner")
		}
	}

	return nil
}

type mockSqlExecutor struct {
	lastQuery string
	lastArgs  []any
	compiler  DialectCompiler
}

func (m *mockSqlExecutor) Compiler() DialectCompiler {
	if m.compiler != nil {
		return m.compiler
	}

	return &SQLiteCompiler{}
}

func (m *mockSqlExecutor) QueryRow(ctx context.Context, query string, args ...any) (*sql.Row, *appfault.AppError) {
	m.lastQuery = query
	m.lastArgs = args

	return nil, appfault.New(errtype.Database, "mock query row intercepted")
}

func (m *mockSqlExecutor) Query(ctx context.Context, query string, args ...any) (*sql.Rows, *appfault.AppError) {
	m.lastQuery = query
	m.lastArgs = args

	return nil, appfault.New(errtype.Database, "mock query intercepted")
}

func (m *mockSqlExecutor) Exec(ctx context.Context, query string, args ...any) (sql.Result, *appfault.AppError) {
	m.lastQuery = query
	m.lastArgs = args

	return mockSqlResult{rowsAffected: 1}, nil
}

func (m *mockSqlExecutor) ExecRowsAffected(ctx context.Context, query string, args ...any) RowsAffectedResult {
	m.lastQuery = query
	m.lastArgs = args

	return SuccessRowsAffected(1)
}

type mockSqlResult struct {
	rowsAffected int64
}

func (r mockSqlResult) LastInsertId() (int64, error) {
	return 1, nil
}

func (r mockSqlResult) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}

type TestItem struct {
	ItemId   uint64
	ItemName string
	Category string
	IsActive bool
}

type TestItemFieldType string

func (e TestItemFieldType) Name() string   { return string(e) }
func (e TestItemFieldType) String() string { return string(e) }
func (e TestItemFieldType) Value() string  { return string(e) }

func scanTestItem(s RowScanner) (*TestItem, error) {
	var item TestItem
	err := s.Scan(&item.ItemId, &item.ItemName, &item.Category, &item.IsActive)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func TestResolveCompiler(t *testing.T) {
	dialects := []DatabaseDialectType{
		DbSQLite,
		DbPostgreSQL,
		DbMySQL,
		DbMariaDB,
		DbMSSQL,
		DbOracle,
		DbMongoDB,
	}

	for _, d := range dialects {
		compiler, err := ResolveCompiler(d)
		if err != nil {
			t.Fatalf("expected compiler for %s, got err: %v", d, err)
		}

		if compiler == nil {
			t.Fatalf("compiler for %s is nil", d)
		}
	}

	_, err := ResolveCompiler("unsupported")
	if err == nil {
		t.Fatal("expected error for unsupported dialect, got nil")
	}
}

func TestCompilerSyntaxes(t *testing.T) {
	sqliteComp := &SQLiteCompiler{}
	if sqliteComp.Placeholder(1) != "?" {
		t.Errorf("expected ?, got %s", sqliteComp.Placeholder(1))
	}

	searchSqlite := sqliteComp.CompileSearch("User", []string{"UserId", "Email"}, 5)
	expectedSqlite := `SELECT * FROM "User" WHERE "UserId" = ? AND "Email" = ? LIMIT 5;`
	if searchSqlite != expectedSqlite {
		t.Errorf("sqlite search mismatch:\ngot:  %s\nwant: %s", searchSqlite, expectedSqlite)
	}

	pgComp := &PostgresCompiler{}
	if pgComp.Placeholder(1) != "$1" {
		t.Errorf("pg placeholder mismatch: %s", pgComp.Placeholder(1))
	}
	if pgComp.Placeholder(2) != "$2" {
		t.Errorf("pg placeholder mismatch: %s", pgComp.Placeholder(2))
	}

	searchPg := pgComp.CompileSearch("User", []string{"UserId"}, 1)
	expectedPg := `SELECT * FROM "User" WHERE "UserId" = $1 LIMIT 1;`
	if searchPg != expectedPg {
		t.Errorf("pg search mismatch:\ngot:  %s\nwant: %s", searchPg, expectedPg)
	}

	mssqlComp := &MSSQLCompiler{}
	if mssqlComp.Placeholder(1) != "@p1" {
		t.Errorf("mssql placeholder mismatch: %s", mssqlComp.Placeholder(1))
	}

	searchMssql := mssqlComp.CompileSearch("User", []string{"UserId"}, 1)
	expectedMssql := `SELECT * FROM [User] WHERE [UserId] = @p1 ORDER BY (SELECT NULL) OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY;`
	if searchMssql != expectedMssql {
		t.Errorf("mssql search mismatch:\ngot:  %s\nwant: %s", searchMssql, expectedMssql)
	}
}

func TestSqlOperator(t *testing.T) {
	opEq := SqlOpEqual
	if opEq.String() != "=" {
		t.Errorf("expected =, got %s", opEq.String())
	}

	if !opEq.IsEqual() {
		t.Errorf("expected Eq to be equal")
	}

	if !opEq.IsEnum() {
		t.Errorf("expected Eq to be valid enum")
	}

	opGt := SqlOpGreaterThan
	if !opGt.IsGreaterThan() {
		t.Errorf("expected Gt to be greater than")
	}

	b, err := json.Marshal(opEq)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var opUnmarshaled SqlOperator
	if unmarshalErr := json.Unmarshal(b, &opUnmarshaled); unmarshalErr != nil {
		t.Fatalf("unmarshal error: %v", unmarshalErr)
	}

	if opUnmarshaled != opEq {
		t.Errorf("expected %s, got %s", opEq, opUnmarshaled)
	}

	var badOp SqlOperator
	if badErr := json.Unmarshal([]byte(`"INVALID_OP"`), &badOp); badErr == nil {
		t.Errorf("expected error for invalid operator")
	}
}

func TestQueryCache(t *testing.T) {
	cache := NewCompiledQueryCache()
	if cache.Size() != 0 {
		t.Errorf("expected size 0, got %d", cache.Size())
	}

	cache.Put("k1", "SELECT 1")
	if cache.Size() != 1 {
		t.Errorf("expected size 1, got %d", cache.Size())
	}

	val, found := cache.Get("k1")
	if !found {
		t.Errorf("expected found")
	}
	if val != "SELECT 1" {
		t.Errorf("expected 'SELECT 1', got %s", val)
	}

	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", cache.Size())
	}
}

func TestResultTypes(t *testing.T) {
	uRes := SuccessUint64(42)
	if !uRes.IsSuccess() {
		t.Errorf("expected success")
	}
	if uRes.Value() != 42 {
		t.Errorf("expected 42, got %d", uRes.Value())
	}

	failErr := appfault.New(errtype.Database, "db failure")
	uFail := FailureUint64(failErr)
	if !uFail.IsFailed() {
		t.Errorf("expected failed")
	}
	if uFail.AppError() == nil {
		t.Errorf("expected app error")
	}

	bRes := SuccessBool(true)
	if !bRes.Value() {
		t.Errorf("expected true")
	}

	strRes := SuccessString("hello")
	if strRes.Value() != "hello" {
		t.Errorf("expected 'hello', got %s", strRes.Value())
	}

	intRes := SuccessInt64(100)
	if intRes.Value() != 100 {
		t.Errorf("expected 100, got %d", intRes.Value())
	}

	rowsRes := SuccessRowsAffected(5)
	if rowsRes.Value() != 5 {
		t.Errorf("expected 5, got %d", rowsRes.Value())
	}

	entRes := SuccessEntity(&TestItem{ItemId: 1, ItemName: "Test"})
	if entRes.Value().ItemName != "Test" {
		t.Errorf("expected 'Test', got %s", entRes.Value().ItemName)
	}

	listRes := SuccessList([]TestItem{{ItemId: 1}, {ItemId: 2}})
	if len(listRes.Value()) != 2 {
		t.Errorf("expected 2 items, got %d", len(listRes.Value()))
	}
}

func TestQueryBuilder_Compilation(t *testing.T) {
	mockExec := &mockSqlExecutor{compiler: &SQLiteCompiler{}}
	repo := NewRepository[TestItem, TestItemFieldType](mockExec, "TestItem", scanTestItem)

	t.Run("Basic Select", func(t *testing.T) {
		qb := repo.Query().
			Select("ItemId", "ItemName").
			WhereEq("Category", "Tool").
			OrderBy("ItemId", "ASC").
			Limit(10).
			Offset(5)

		sqlStr, args := qb.BuildSelect()
		expected := `SELECT "ItemId", "ItemName" FROM "TestItem" WHERE "TestItem"."Category" = ? ORDER BY "TestItem"."ItemId" ASC LIMIT 10 OFFSET 5;`
		if sqlStr != expected {
			t.Errorf("sql mismatch:\ngot:  %s\nwant: %s", sqlStr, expected)
		}

		if len(args) != 1 || args[0] != "Tool" {
			t.Errorf("args mismatch: %v", args)
		}
	})

	t.Run("Count Query", func(t *testing.T) {
		qb := repo.Query().WhereEq("IsActive", 1)
		sqlStr, args := qb.BuildCount()
		expected := `SELECT COUNT(*) FROM "TestItem" WHERE "TestItem"."IsActive" = ?;`
		if sqlStr != expected {
			t.Errorf("count sql mismatch:\ngot:  %s\nwant: %s", sqlStr, expected)
		}
		if len(args) != 1 || args[0] != 1 {
			t.Errorf("args mismatch: %v", args)
		}
	})

	t.Run("Delete Query", func(t *testing.T) {
		qb := repo.Query().WhereEq("Category", "Old")
		sqlStr, args := qb.BuildDelete()
		expected := `DELETE FROM "TestItem" WHERE "TestItem"."Category" = ?;`
		if sqlStr != expected {
			t.Errorf("delete sql mismatch:\ngot:  %s\nwant: %s", sqlStr, expected)
		}
		if len(args) != 1 || args[0] != "Old" {
			t.Errorf("args mismatch: %v", args)
		}
	})

	t.Run("Join Query", func(t *testing.T) {
		qb := repo.Query().
			Select("ItemId").
			Join("CategoryTable").
			Select("CategoryName").
			OnField("Category", SqlOperators.Equal, "CategoryId")

		sqlStr, _ := qb.BuildSelect()
		if !strings.Contains(sqlStr, `INNER JOIN "CategoryTable" ON "TestItem"."Category" = "CategoryTable"."CategoryId"`) {
			t.Errorf("join clause missing or invalid: %s", sqlStr)
		}
	})

	t.Run("Group By and Having", func(t *testing.T) {
		qb := repo.Query().
			Select("Category").
			GroupBy("Category").
			HavingCount(SqlOperators.GreaterThan, 5)

		sqlStr, args := qb.BuildSelect()
		if !strings.Contains(sqlStr, `GROUP BY "TestItem"."Category"`) {
			t.Errorf("group by missing: %s", sqlStr)
		}
		if !strings.Contains(sqlStr, `HAVING COUNT(*) > ?`) {
			t.Errorf("having count missing: %s", sqlStr)
		}
		if len(args) != 1 || args[0] != int64(5) {
			t.Errorf("having args mismatch: %v", args)
		}
	})

	t.Run("CTE WithView", func(t *testing.T) {
		qb := repo.Query().
			WithView("ActiveItems", "SELECT * FROM RawItems WHERE Active = 1").
			WhereEq("Category", "General")

		sqlStr, _ := qb.BuildSelect()
		if !strings.HasPrefix(sqlStr, `WITH "ActiveItems" AS (SELECT * FROM RawItems WHERE Active = 1)`) {
			t.Errorf("CTE prefix missing: %s", sqlStr)
		}
	})

	t.Run("Locate Filter", func(t *testing.T) {
		qb := repo.Query().Locate("ItemName", "tool")
		sqlStr, args := qb.BuildSelect()
		if !strings.Contains(sqlStr, `WHERE INSTR("TestItem"."ItemName", ?) > 0`) {
			t.Errorf("locate filter mismatch: %s", sqlStr)
		}
		if len(args) != 1 || args[0] != "tool" {
			t.Errorf("locate args mismatch: %v", args)
		}
	})

	t.Run("Deterministic QueryHash", func(t *testing.T) {
		qb1 := repo.Query().WhereEq("Category", "A")
		qb2 := repo.Query().WhereEq("Category", "A")

		hash1 := qb1.QueryHash()
		hash2 := qb2.QueryHash()

		if hash1 != hash2 {
			t.Errorf("expected deterministic hash: %s vs %s", hash1, hash2)
		}
		if len(hash1) != 64 {
			t.Errorf("expected sha256 64 chars, got %d", len(hash1))
		}
	})

	t.Run("Compile and GlobalQueryCache", func(t *testing.T) {
		qb := repo.Query().WhereEq("ItemName", "Hammer")
		res1 := qb.Compile()
		if res1.IsFailed() {
			t.Fatalf("compile failed: %v", res1.AppError())
		}

		res2 := qb.Compile()
		if res2.IsFailed() {
			t.Fatalf("cached compile failed: %v", res2.AppError())
		}

		if res1.Value().SQL != res2.Value().SQL {
			t.Errorf("expected identical compiled SQL")
		}
	})
}

func TestRepository_Methods(t *testing.T) {
	mockExec := &mockSqlExecutor{compiler: &SQLiteCompiler{}}
	repo := NewRepository[TestItem, TestItemFieldType](mockExec, "TestItem", scanTestItem)
	ctx := context.Background()

	_ = repo.First(ctx, "ItemId", 1)
	if !strings.Contains(mockExec.lastQuery, `WHERE "TestItem"."ItemId" = ?`) {
		t.Errorf("First query mismatch: %s", mockExec.lastQuery)
	}

	_ = repo.FindById(ctx, "ItemId", 42)
	if !strings.Contains(mockExec.lastQuery, `WHERE "TestItem"."ItemId" = ?`) {
		t.Errorf("FindById query mismatch: %s", mockExec.lastQuery)
	}

	_ = repo.FindBy(ctx, "Category", "Tools", 10)
	if !strings.Contains(mockExec.lastQuery, `LIMIT 10`) {
		t.Errorf("FindBy limit mismatch: %s", mockExec.lastQuery)
	}

	_ = repo.FindBy2(ctx, "Category", "Tools", "IsActive", 1, 20)
	if !strings.Contains(mockExec.lastQuery, `WHERE "TestItem"."Category" = ? AND "TestItem"."IsActive" = ?`) {
		t.Errorf("FindBy2 where mismatch: %s", mockExec.lastQuery)
	}

	_ = repo.Count(ctx, "Category", "Tools")
	if !strings.Contains(mockExec.lastQuery, `SELECT COUNT(*) FROM "TestItem"`) {
		t.Errorf("Count query mismatch: %s", mockExec.lastQuery)
	}

	_ = repo.CountAll(ctx)
	if !strings.Contains(mockExec.lastQuery, `SELECT COUNT(*) FROM "TestItem"`) {
		t.Errorf("CountAll query mismatch: %s", mockExec.lastQuery)
	}

	_ = repo.DeleteBy(ctx, "Category", "Deprecated")
	if !strings.Contains(mockExec.lastQuery, `DELETE FROM "TestItem" WHERE "TestItem"."Category" = ?`) {
		t.Errorf("DeleteBy query mismatch: %s", mockExec.lastQuery)
	}
}

func TestMockRowScanner(t *testing.T) {
	scanner := &mockRowScanner{
		values: []any{uint64(10), "Sample", "CategoryA", true},
	}

	item, err := scanTestItem(scanner)
	if err != nil {
		t.Fatalf("scan error: %v", err)
	}

	if item.ItemId != 10 {
		t.Errorf("expected ItemId 10, got %d", item.ItemId)
	}
	if item.ItemName != "Sample" {
		t.Errorf("expected ItemName 'Sample', got %s", item.ItemName)
	}
	if !item.IsActive {
		t.Errorf("expected IsActive true")
	}

	errScanner := &mockRowScanner{err: errors.New("simulated scan error")}
	_, scanErr := scanTestItem(errScanner)
	if scanErr == nil {
		t.Errorf("expected scan error")
	}
}
