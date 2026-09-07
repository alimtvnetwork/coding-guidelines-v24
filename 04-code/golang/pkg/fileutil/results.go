package fileutil

import (
	"os"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

func FileSuccess(f *os.File) FileResult {
	return result.WrapSuccess(f)
}

func FileFailure(variation errtype.Variation, err error, path string, msg string) FileResult {
	return result.WrapFailure[*os.File](appfault.WrapFile(variation, err, path, msg))
}

func FileFailureMsg(variation errtype.Variation, path string, msg string) FileResult {
	return result.WrapFailure[*os.File](appfault.NewFile(variation, path, msg))
}

func FileFailureFault(fault *appfault.AppError) FileResult {
	return result.WrapFailure[*os.File](fault)
}

func BoolSuccess(isSuccess bool) BoolResult {
	return result.WrapSuccess(isSuccess)
}

func BoolFailure(variation errtype.Variation, err error, path string, msg string) BoolResult {
	return result.WrapFailure[bool](appfault.WrapFile(variation, err, path, msg))
}

func BoolFailureMsg(variation errtype.Variation, path string, msg string) BoolResult {
	return result.WrapFailure[bool](appfault.NewFile(variation, path, msg))
}

func BoolFailureFault(fault *appfault.AppError) BoolResult {
	return result.WrapFailure[bool](fault)
}

func BytesSuccess(data []byte) BytesResult {
	return result.WrapSuccess(data)
}

func BytesFailure(variation errtype.Variation, err error, path string, msg string) BytesResult {
	return result.WrapFailure[[]byte](appfault.WrapFile(variation, err, path, msg))
}

func BytesFailureMsg(variation errtype.Variation, path string, msg string) BytesResult {
	return result.WrapFailure[[]byte](appfault.NewFile(variation, path, msg))
}

func BytesFailureFault(fault *appfault.AppError) BytesResult {
	return result.WrapFailure[[]byte](fault)
}

func StringSuccess(s string) StringResult {
	return result.WrapSuccess(s)
}

func StringFailure(variation errtype.Variation, err error, path string, msg string) StringResult {
	return result.WrapFailure[string](appfault.WrapFile(variation, err, path, msg))
}

func StringFailureMsg(variation errtype.Variation, path string, msg string) StringResult {
	return result.WrapFailure[string](appfault.NewFile(variation, path, msg))
}

func StringFailureFault(fault *appfault.AppError) StringResult {
	return result.WrapFailure[string](fault)
}

func LinesSuccess(lines []string) LinesResult {
	return result.WrapSuccess(lines)
}

func LinesFailure(variation errtype.Variation, err error, path string, msg string) LinesResult {
	return result.WrapFailure[[]string](appfault.WrapFile(variation, err, path, msg))
}

func LinesFailureMsg(variation errtype.Variation, path string, msg string) LinesResult {
	return result.WrapFailure[[]string](appfault.NewFile(variation, path, msg))
}

func LinesFailureFault(fault *appfault.AppError) LinesResult {
	return result.WrapFailure[[]string](fault)
}

func FileInfoSuccess(info os.FileInfo) FileInfoResult {
	return result.WrapSuccess(info)
}

func FileInfoFailure(variation errtype.Variation, err error, path string, msg string) FileInfoResult {
	return result.WrapFailure[os.FileInfo](appfault.WrapFile(variation, err, path, msg))
}

func FileInfoFailureMsg(variation errtype.Variation, path string, msg string) FileInfoResult {
	return result.WrapFailure[os.FileInfo](appfault.NewFile(variation, path, msg))
}

func FileInfoFailureFault(fault *appfault.AppError) FileInfoResult {
	return result.WrapFailure[os.FileInfo](fault)
}

func Int64Success(n int64) Int64Result {
	return result.WrapSuccess(n)
}

func Int64Failure(variation errtype.Variation, err error, path string, msg string) Int64Result {
	return result.WrapFailure[int64](appfault.WrapFile(variation, err, path, msg))
}

func Int64FailureMsg(variation errtype.Variation, path string, msg string) Int64Result {
	return result.WrapFailure[int64](appfault.NewFile(variation, path, msg))
}

func Int64FailureFault(fault *appfault.AppError) Int64Result {
	return result.WrapFailure[int64](fault)
}
