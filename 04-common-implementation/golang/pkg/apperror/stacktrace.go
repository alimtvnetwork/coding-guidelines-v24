package apperror

import (
	"fmt"
	"runtime"
	"strings"
)

// StackFrame holds structured metadata for a single call frame.
type StackFrame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

// StackTrace is a collection of structured call frames.
type StackTrace []StackFrame

// parseFrames iterates over runtime frames and extracts non-runtime frames.
func parseFrames(frames *runtime.Frames) StackTrace {
	var trace StackTrace
	for {
		f, more := frames.Next()
		if isAppFrame(f.File) {
			trace = append(trace, StackFrame{Function: f.Function, File: f.File, Line: f.Line})
		}

		if !more {
			break
		}
	}

	return trace
}

// captureStackTrace captures up to 32 caller frames starting from skip offset.
func captureStackTrace(skip int) StackTrace {
	pc := make([]uintptr, 32)
	n := runtime.Callers(skip+1, pc)
	if n == 0 {
		return StackTrace{}
	}

	return parseFrames(runtime.CallersFrames(pc[:n]))
}

// isAppFrame filters out runtime and test runner boilerplate.
func isAppFrame(file string) bool {
	return !strings.Contains(file, "runtime/")
}

// Depth returns the number of captured frames.
func (st StackTrace) Depth() int {
	return len(st)
}

// IsEmpty returns true if no frames were captured.
func (st StackTrace) IsEmpty() bool {
	return len(st) == 0
}

// CallerLine returns a compact "file.go:line" representation of the top frame.
func (st StackTrace) CallerLine() string {
	if len(st) == 0 {
		return "unknown:0"
	}

	return fmt.Sprintf("%s:%d", st[0].File, st[0].Line)
}

// String formats the full multi-line stack trace.
func (st StackTrace) String() string {
	var builder strings.Builder
	for idx, frame := range st {
		builder.WriteString(fmt.Sprintf("#%d %s\n   %s:%d\n", idx, frame.Function, frame.File, frame.Line))
	}

	return builder.String()
}
