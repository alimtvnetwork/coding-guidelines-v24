package appfault_test

import (
	"fmt"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

func TestAppErrorCreationAndNilSafety(t *testing.T) {
	var nilErr *appfault.AppError
	if nilErr.HasError() || !nilErr.IsSuccess() || appfault.New(errtype.None, "") != nil {
		t.Fatal("expected nil AppError and None constructor to be IsSuccess")
	}

	appErr := appfault.New(errtype.NotFound, "record not found").WithOp("repo.find").WithSeverity(appfault.SeverityWarn)
	if !appErr.Is(errtype.NotFound) || !appErr.HasValidError() {
		t.Fatalf("expected NotFound error, got %v", appErr.GetType())
	}
}

func TestContextMapOperations(t *testing.T) {
	cm := appfault.NewContextMap().Set("siteId", 101).Set("slug", "test-plugin")
	if !cm.Has("siteId") || cm.GetString("siteId") != "101" || cm.Count() != 2 {
		t.Fatalf("expected siteId=101 and count 2, got %s (count=%d)", cm.GetString("siteId"), cm.Count())
	}

	cm.Remove("slug")
	if cm.Has("slug") {
		t.Fatal("expected slug to be removed")
	}
}

func TestStackFrameAndCaller(t *testing.T) {
	frame := appfault.NewStackFrame("main.run", "main.go", 42)
	if frame.Function != "main.run" || frame.File != "main.go" || frame.Line != 42 {
		t.Fatalf("unexpected frame: %+v", frame)
	}

	caller := appfault.CaptureCaller(0)
	if len(caller) == 0 {
		t.Fatal("expected non-empty caller")
	}
}

func TestCallerAndStackTraceObjects(t *testing.T) {
	appErr := appfault.New(errtype.Database, "db err")
	caller := appErr.Caller()
	if caller.IsEmpty() || caller.Line == 0 {
		t.Fatalf("expected caller info to have line and file: %+v", caller)
	}

	stack := appErr.StackTrace()
	if len(stack) == 0 {
		t.Fatal("expected non-empty stack trace")
	}
}

func TestResultMonadicOperations(t *testing.T) {
	res := result.SuccessResult("data-payload")
	if !res.IsSuccess() || res.IsFailed() || res.Data() != "data-payload" {
		t.Fatal("expected success result")
	}

	failRes := result.NewFailureWithType[string](errtype.Validation, "bad input", "validator")
	if !failRes.IsFailed() || failRes.UnwrapOr("fallback") != "fallback" {
		t.Fatal("expected failed result with fallback")
	}
}

func TestResultSliceAndMap(t *testing.T) {
	sliceRes := appfault.OkSlice([]string{"alpha", "beta"})
	if sliceRes.Count() != 2 || !sliceRes.HasItems() {
		t.Fatalf("expected count 2, got %d", sliceRes.Count())
	}

	mapRes := appfault.OkMap(map[string]int{"key1": 100})
	if !mapRes.Has("key1") || mapRes.Count() != 1 {
		t.Fatal("expected key1 to exist")
	}
}

func TestAppErrorPathAndVariableEmbedding(t *testing.T) {
	targetPath := "/var/log/app/trace.json"

	// 1. NewWithPath
	err1 := appfault.NewWithPath(errtype.IO, "cannot open log file", targetPath)
	if err1 == nil {
		t.Fatal("expected non-nil error")
	}
	ctx1 := err1.Context()
	valPath, isFound := ctx1.Get("Path")
	if !isFound || valPath != targetPath {
		t.Fatalf("expected Path to be %s, got %v", targetPath, valPath)
	}
	valFilePath, isFound := ctx1.Get("FilePath")
	if !isFound || valFilePath != targetPath {
		t.Fatalf("expected FilePath to be %s, got %v", targetPath, valFilePath)
	}

	// 2. WrapWithPath
	rootErr := fmt.Errorf("permission denied: root level")
	err2 := appfault.WrapWithPath(errtype.Forbidden, rootErr, "access violation", targetPath)
	if err2 == nil || err2.Cause() != rootErr {
		t.Fatal("expected wrapped cause error")
	}
	ctx2 := err2.Context()
	if p, ok := ctx2.Get("Path"); !ok || p != targetPath {
		t.Fatalf("expected Path in wrapped error to be %s", targetPath)
	}

	// 3. WithVar and WithVars
	err3 := appfault.NewWithVar(errtype.Validation, "invalid batch size", "batchSize", 500).
		WithVar("minSize", 1).
		WithVar("maxSize", 100)
	ctx3 := err3.Context()
	if v, ok := ctx3.Get("batchSize"); !ok || v != 500 {
		t.Fatalf("expected batchSize 500, got %v", v)
	}
	varsRaw, isFound := ctx3.Get("Variables")
	if !isFound {
		t.Fatal("expected Variables map to be present")
	}
	varsMap, isMap := varsRaw.(map[string]any)
	if !isMap || varsMap["batchSize"] != 500 || varsMap["minSize"] != 1 || varsMap["maxSize"] != 100 {
		t.Fatalf("unexpected Variables map: %v", varsMap)
	}

	// 4. Multiple Paths
	srcPath := "/tmp/source.csv"
	dstPath := "/data/destination.csv"
	err4 := appfault.New(errtype.IO, "failed to move file").WithPaths(srcPath, dstPath)
	ctx4 := err4.Context()
	pathsRaw, isFound := ctx4.Get("Paths")
	if !isFound {
		t.Fatal("expected Paths slice in context")
	}
	paths, isSlice := pathsRaw.([]string)
	if !isSlice || len(paths) != 2 || paths[0] != srcPath || paths[1] != dstPath {
		t.Fatalf("unexpected Paths slice: %v", pathsRaw)
	}

	// 5. Builder Path & Variable integration
	builder := appfault.NewBuilder(errtype.Execution, "lock file conflict").
		WithPath(targetPath).
		WithVar("lockHolderPid", 12345)
	builtErr := builder.Build()
	if builtErr == nil {
		t.Fatal("expected built error to be non-nil")
	}
	builtCtx := builtErr.Context()
	if p, ok := builtCtx.Get("Path"); !ok || p != targetPath {
		t.Fatalf("expected Path in builder: %v", p)
	}
	if pid, ok := builtCtx.Get("lockHolderPid"); !ok || pid != 12345 {
		t.Fatalf("expected lockHolderPid in builder: %v", pid)
	}
}
