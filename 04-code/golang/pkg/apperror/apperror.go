package apperror

import "coding-guidelines/common/pkg/appfault"

// AppError aliases appfault.AppError for backward compatibility.
type AppError = appfault.AppError

// Fault aliases appfault.Fault for backward compatibility.
type Fault = appfault.Fault

// Result aliases appfault.Result for backward compatibility.
type Result[T any] = appfault.Result[T]

// ResultSlice aliases appfault.ResultSlice for backward compatibility.
type ResultSlice[T any] = appfault.ResultSlice[T]

// ResultMap aliases appfault.ResultMap for backward compatibility.
type ResultMap[K comparable, V any] = appfault.ResultMap[K, V]

// ErrorType aliases appfault.ErrorType.
type ErrorType = appfault.ErrorType

// SeverityType aliases appfault.SeverityType.
type SeverityType = appfault.SeverityType

// ErrorCodeType aliases appfault.ErrorCodeType.
type ErrorCodeType = appfault.ErrorCodeType

// Forward common constructors
var (
	NewSimple          = appfault.NewSimple
	New                = appfault.New
	NewWithDetails     = appfault.NewWithDetails
	NewValidationError = appfault.NewValidationError
	WrapSimple         = appfault.WrapSimple
	Wrap               = appfault.Wrap
	WrapWithDetails    = appfault.WrapWithDetails
)
