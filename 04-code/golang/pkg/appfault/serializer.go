package appfault

import "errors"

// AppErrorDataModel is the serializable DTO for AppError.
type AppErrorDataModel struct {
	Op         string       `json:"Op,omitempty" yaml:"Op,omitempty"`
	Code       string       `json:"Code,omitempty" yaml:"Code,omitempty"`
	Type       ErrorType    `json:"Type,omitempty" yaml:"Type,omitempty"`
	Severity   SeverityType `json:"Severity,omitempty" yaml:"Severity,omitempty"`
	Creator    string       `json:"Creator,omitempty" yaml:"Creator,omitempty"`
	Message    string       `json:"Message,omitempty" yaml:"Message,omitempty"`
	Caller     string       `json:"Caller,omitempty" yaml:"Caller,omitempty"`
	Stack      string       `json:"Stack,omitempty" yaml:"Stack,omitempty"`
	Ctx        ContextMap   `json:"Ctx,omitempty" yaml:"Ctx,omitempty"`
	Cause      string       `json:"Cause,omitempty" yaml:"Cause,omitempty"`
	StatusCode int          `json:"StatusCode,omitempty" yaml:"StatusCode,omitempty"`
}

// extractCauseString safely extracts the cause error message string.
func extractCauseString(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}

// ToDataModel converts an AppError into its serializable data model.
func (e *AppError) ToDataModel() AppErrorDataModel {
	if e == nil {
		return AppErrorDataModel{}
	}

	return AppErrorDataModel{
		Op: e.Op, Code: e.Code, Type: e.Type, Severity: e.Severity,
		Creator: e.Creator, Message: e.Message, Caller: e.Caller,
		Stack: e.Stack, Ctx: e.Ctx.Clone(), Cause: extractCauseString(e.Cause),
		StatusCode: e.StatusCode,
	}
}

// buildCauseError safely constructs an error if cause is non-empty.
func buildCauseError(cause string) error {
	if len(cause) == 0 {
		return nil
	}

	return errors.New(cause)
}

// ToAppError reconstructs an AppError from the data model.
func (m AppErrorDataModel) ToAppError() *AppError {
	return &AppError{
		Op: m.Op, Code: m.Code, Type: m.Type, Severity: m.Severity,
		Creator: m.Creator, Message: m.Message, Caller: m.Caller,
		Stack: m.Stack, Ctx: m.Ctx.Clone(), Cause: buildCauseError(m.Cause),
		StatusCode: m.StatusCode,
	}
}
