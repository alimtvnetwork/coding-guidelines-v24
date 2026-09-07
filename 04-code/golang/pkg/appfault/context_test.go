package appfault_test

import (
	"errors"
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

func TestAppError_WithVar(t *testing.T) {
	orig := appfault.New(errtype.Generic, "base err")
	withV := orig.WithVar("cluster", "prod-us-east")

	if withV.Context().GetString("cluster") != "prod-us-east" {
		t.Fatalf("expected cluster context to be set")
	}

	if orig.Context().Has("cluster") {
		t.Fatalf("expected immutability, original has key")
	}
}

func TestAppError_WithPathAndFilePath(t *testing.T) {
	err1 := appfault.New(errtype.IO, "io err").WithPath("/etc/config.json")
	if err1.Context().GetString("Path") != "/etc/config.json" {
		t.Fatalf("expected Path context to be set")
	}

	err2 := appfault.New(errtype.IO, "io err").WithFilePath("/var/log/syslog")
	if err2.Context().GetString("Path") != "/var/log/syslog" {
		t.Fatalf("expected Path context from WithFilePath")
	}
}

func TestAppError_WithField(t *testing.T) {
	err := appfault.New(errtype.Validation, "validation err").WithField("email", "test@example.com")
	if err.Context().GetString("email") != "test@example.com" {
		t.Fatalf("expected email context")
	}

	var nilErr *appfault.AppError
	if nilErr.WithField("k", "v") != nil {
		t.Fatalf("expected nil for nil receiver")
	}
}

func TestAppError_ContextNilReceiverSafety(t *testing.T) {
	var nilErr *appfault.AppError
	if nilErr.WithVar("k", "v") != nil {
		t.Fatalf("expected nil on WithVar")
	}

	if nilErr.WithPath("/tmp") != nil {
		t.Fatalf("expected nil on WithPath")
	}

	if nilErr.WithFilePath("/tmp") != nil {
		t.Fatalf("expected nil on WithFilePath")
	}
}

func TestBuilder_WithVarAndPath(t *testing.T) {
	b := appfault.NewBuilder(errtype.IO, "builder error")
	b.WithVar("tenant", "corp-1").WithPath("/data/file.txt")
	err := b.Build()

	if err.Context().GetString("tenant") != "corp-1" {
		t.Fatalf("expected tenant in context")
	}

	if err.Context().GetString("Path") != "/data/file.txt" {
		t.Fatalf("expected Path in context")
	}
}

func TestBuilder_WithFilePathAndField(t *testing.T) {
	b := appfault.NewBuilder(errtype.Validation, "builder error")
	b.WithFilePath("/configs/app.yaml").WithField("retries", 3)
	err := b.Build()

	if err.Context().GetString("Path") != "/configs/app.yaml" {
		t.Fatalf("expected Path from WithFilePath")
	}

	val, exists := err.Context().Get("retries")
	if !exists || val != 3 {
		t.Fatalf("expected retries=3 in context")
	}
}

func TestConstructors_NewPathAndNewFile(t *testing.T) {
	errPath := appfault.NewPath(errtype.NotFound, "/var/data.bin", "file missing")
	if errPath.Context().GetString("Path") != "/var/data.bin" {
		t.Fatalf("expected Path on NewPath")
	}

	errFile := appfault.NewFile(errtype.IO, "/var/lock", "lock error")
	if errFile.Context().GetString("Path") != "/var/lock" {
		t.Fatalf("expected Path on NewFile")
	}

	if appfault.NewPath(errtype.None, "/path", "none") != nil {
		t.Fatalf("expected nil for None errtype")
	}
}

func TestConstructors_WrapPathAndWrapFile(t *testing.T) {
	cause := errors.New("access denied")
	errPath := appfault.WrapPath(errtype.Unauthorized, cause, "/etc/shadow", "permission denied")
	if errPath.Context().GetString("Path") != "/etc/shadow" {
		t.Fatalf("expected Path on WrapPath")
	}

	if errPath.Unwrap() != cause {
		t.Fatalf("expected root cause preserved")
	}

	errFile := appfault.WrapFile(errtype.IO, cause, "/etc/hosts", "cannot read")
	if errFile.Context().GetString("Path") != "/etc/hosts" {
		t.Fatalf("expected Path on WrapFile")
	}
}

func TestConstructors_WrapNilHandling(t *testing.T) {
	if appfault.WrapPath(errtype.IO, nil, "/path", "msg") != nil {
		t.Fatalf("expected nil on nil cause")
	}

	if appfault.WrapFile(errtype.IO, nil, "/path", "msg") != nil {
		t.Fatalf("expected nil on nil cause")
	}

	cause := errors.New("err")
	if appfault.WrapPath(errtype.None, cause, "/path", "msg") != nil {
		t.Fatalf("expected nil on None errtype")
	}
}

func TestConstructors_NewVarAndWrapVar(t *testing.T) {
	errVar := appfault.NewVar(errtype.Validation, "port", 8080, "invalid port")
	if errVar.Context().GetString("port") != "8080" {
		t.Fatalf("expected port on NewVar")
	}

	cause := errors.New("underlying var error")
	wrapV := appfault.WrapVar(errtype.Database, cause, "table", "users", "query failed")
	if wrapV.Context().GetString("table") != "users" {
		t.Fatalf("expected table on WrapVar")
	}

	if wrapV.Unwrap() != cause {
		t.Fatalf("expected root cause on WrapVar")
	}
}

func TestConstructors_WrapVarNilHandling(t *testing.T) {
	if appfault.WrapVar(errtype.IO, nil, "key", "val", "msg") != nil {
		t.Fatalf("expected nil on nil cause")
	}

	cause := errors.New("err")
	if appfault.WrapVar(errtype.None, cause, "key", "val", "msg") != nil {
		t.Fatalf("expected nil on None errtype")
	}

	if appfault.NewVar(errtype.None, "key", "val", "msg") != nil {
		t.Fatalf("expected nil on None errtype")
	}
}

func TestResultConstructors_FailurePathAndFile(t *testing.T) {
	rPath := appfault.FailurePath[string](errtype.NotFound, "/missing/file", "not found")
	if rPath.IsSuccess() || rPath.Fault().Context().GetString("Path") != "/missing/file" {
		t.Fatalf("expected failure with Path context")
	}

	rFile := appfault.FailureFile[int](errtype.IO, "/dev/null", "cannot write")
	if rFile.IsSuccess() || rFile.Fault().Context().GetString("Path") != "/dev/null" {
		t.Fatalf("expected failure with File Path context")
	}

	rVar := appfault.FailureVar[bool](errtype.Validation, "flag", false, "invalid flag")
	if rVar.IsSuccess() || !rVar.Fault().Context().Has("flag") {
		t.Fatalf("expected failure with Var context")
	}
}

func TestResultConstructors_NewFailureWithFileAndPath(t *testing.T) {
	cause := errors.New("raw disk failure")
	rFile := appfault.NewFailureWithFile[string](errtype.IO, cause, "/mnt/disk1", "disk failure")
	if rFile.IsSuccess() || rFile.Fault().Context().GetString("Path") != "/mnt/disk1" {
		t.Fatalf("expected failure with Path context")
	}

	rPath := appfault.NewFailureWithPath[int](errtype.IO, cause, "/mnt/disk2", "path failure")
	if rPath.IsSuccess() || rPath.Fault().Context().GetString("Path") != "/mnt/disk2" {
		t.Fatalf("expected failure with Path context")
	}

	rVar := appfault.NewFailureWithVar[bool](errtype.Network, cause, "host", "api.internal", "net err")
	if rVar.IsSuccess() || rVar.Fault().Context().GetString("host") != "api.internal" {
		t.Fatalf("expected failure with Var context")
	}
}
