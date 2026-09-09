package applogger_test

import (
	"testing"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/applogger"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

func TestResults_DirectTypesAndConstructors(t *testing.T) {
	fileRes := applogger.FileSinkSuccess(nil)
	if fileRes.IsFailed() {
		t.Fatal("expected success FileSinkResult")
	}

	rawFail := result.WrapFailure[string](appfault.New(errtype.IO, "failed"))
	fileFail := applogger.FileSinkFailure(rawFail)
	if !fileFail.IsFailed() {
		t.Fatal("expected failed FileSinkResult")
	}

	faultFail := applogger.FileSinkFailureFault(appfault.New(errtype.IO, "fault"))
	if !faultFail.IsFailed() {
		t.Fatal("expected failed FileSinkResult from fault")
	}
}

func TestResults_SinkConstructors(t *testing.T) {
	rawFail := result.WrapFailure[int](appfault.New(errtype.Database, "db err"))

	rotSuccess := applogger.RotatingFileSinkSuccess(nil)
	rotFail := applogger.RotatingFileSinkFailure(rawFail)
	if rotSuccess.IsFailed() || !rotFail.IsFailed() {
		t.Fatal("unexpected RotatingFileSinkResult state")
	}

	sqlSuccess := applogger.SQLiteSinkSuccess(nil)
	sqlFail := applogger.SQLiteSinkFailure(rawFail)
	if sqlSuccess.IsFailed() || !sqlFail.IsFailed() {
		t.Fatal("unexpected SQLiteSinkResult state")
	}

	apiSuccess := applogger.ApiSinkSuccess(nil)
	apiFail := applogger.ApiSinkFailure(rawFail)
	if apiSuccess.IsFailed() || !apiFail.IsFailed() {
		t.Fatal("unexpected ApiSinkResult state")
	}
}

func TestResults_LoggerConstructors(t *testing.T) {
	rawFail := result.WrapFailure[bool](appfault.New(errtype.Validation, "val err"))

	logSuccess := applogger.LoggerSuccess(nil)
	logFail := applogger.LoggerFailure(rawFail)
	logFaultFail := applogger.LoggerFailureFault(appfault.New(errtype.Validation, "fault"))
	if logSuccess.IsFailed() || !logFail.IsFailed() || !logFaultFail.IsFailed() {
		t.Fatal("unexpected LoggerResult state")
	}

	sinkSuccess := applogger.LogSinkSuccess(nil)
	sinkFail := applogger.LogSinkFailure(rawFail)
	if sinkSuccess.IsFailed() || !sinkFail.IsFailed() {
		t.Fatal("unexpected LogSinkResult state")
	}
}

func TestResults_StreamerSinkConstructors(t *testing.T) {
	rawFail := result.WrapFailure[int](appfault.New(errtype.IO, "stream err"))

	streamSuccess := applogger.StreamerSinkSuccess(nil)
	if streamSuccess.IsFailed() {
		t.Fatal("expected success StreamerSinkResult")
	}

	streamFail := applogger.StreamerSinkFailure(rawFail)
	if !streamFail.IsFailed() {
		t.Fatal("expected failed StreamerSinkResult")
	}

	faultFail := applogger.StreamerSinkFailureFault(appfault.New(errtype.IO, "stream fault"))
	if !faultFail.IsFailed() {
		t.Fatal("expected failed StreamerSinkResult from fault")
	}
}
