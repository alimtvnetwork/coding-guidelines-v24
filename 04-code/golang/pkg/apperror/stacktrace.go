package apperror

import (
	"fmt"
	"runtime"
	"strings"
)

// StackFrame holds metadata for a single caller frame.
type StackFrame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

// StackTrace is a collection of structured call frames.
type StackTrace []StackFrame

// appendFrameIfApp appends frame if not in runtime.
func appendFrameIfApp(trace StackTrace, f runtime.Frame) StackTrace {
	if isAppFrame(f.File) {
		return append(trace, StackFrame{Function: f.Function, File: f.File, Line: f.Line})
	}

	return trace
}

// parseFrames extracts non-runtime frames from runtime.Frames.
func parseFrames(frames *runtime.Frames) StackTrace {
	var trace StackTrace
	for {
		f, more := frames.Next()
		trace = appendFrameIfApp(trace, f)
		if !more {
			break
		}
	}

	return trace
}

// captureStackTrace captures caller frames starting at skip offset.
func captureStackTrace(skip int) StackTrace {
	pc := make([]uintptr, 32)
	n := runtime.Callers(skip+1, pc)
	if n == 0 {
		return StackTrace{}
	}

	return parseFrames(runtime.CallersFrames(pc[:n]))
}

// isAppFrame filters out runtime frames.
func isAppFrame(file string) bool {
	return !strings.Contains(file, "runtime/")
}

// CallerLine returns a compact "file:line" string of the top frame.
func (st StackTrace) CallerLine() string {
	if len(st) == 0 {
		return "unknown:0"
	}

	return fmt.Sprintf("%s:%d", st[0].File, st[0].Line)
}

// String formats the multi-line stack trace.
func (st StackTrace) String() string {
	var builder strings.Builder
	for idx, frame := range st {
		builder.WriteString(fmt.Sprintf("#%d %s\n   %s:%d\n", idx, frame.Function, frame.File, frame.Line))
	}

	return builder.String()
}
