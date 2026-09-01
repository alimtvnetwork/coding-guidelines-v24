package appfault

import (
	"errors"

	"coding-guidelines/common/pkg/errtype"
)

// AppErrorDataModel is the serializable DTO for AppError.
type AppErrorDataModel struct {
	Type       errtype.Variation `json:"Type,omitempty" yaml:"Type,omitempty"`
	Message    string            `json:"Message,omitempty" yaml:"Message,omitempty"`
	Caller     string            `json:"Caller,omitempty" yaml:"Caller,omitempty"`
	Stack      string            `json:"Stack,omitempty" yaml:"Stack,omitempty"`
	Ctx        ContextMap        `json:"Ctx,omitempty" yaml:"Ctx,omitempty"`
	Cause      string            `json:"Cause,omitempty" yaml:"Cause,omitempty"`
	StatusCode int               `json:"StatusCode,omitempty" yaml:"StatusCode,omitempty"`
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
		Type: e.Type, Message: e.Message, Caller: e.Caller,
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
		Type: m.Type, Message: m.Message, Caller: m.Caller,
		Stack: m.Stack, Ctx: m.Ctx.Clone(), Cause: buildCauseError(m.Cause),
		StatusCode: m.StatusCode,
	}
}
