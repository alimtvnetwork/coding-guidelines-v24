package appfault

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type (
	StackFrame struct {
		Function string `json:",omitempty" yaml:",omitempty"`
		File     string `json:",omitempty" yaml:",omitempty"`
		Line     int    `json:",omitempty" yaml:",omitempty"`
	}

	StackFrames []StackFrame

	StackTrace = StackFrames
)

func NewStackFrame(function string, file string, line int) StackFrame {
	frame := StackFrame{}
	frame.Function = function
	frame.File = file
	frame.Line = line

	return frame
}

func NewStackTrace(frames ...StackFrame) StackTrace {
	return StackTrace(frames)
}

// appendFrameIfApp appends frame if not in runtime.
func appendFrameIfApp(trace StackTrace, f runtime.Frame) StackTrace {
	if isAppFrame(f.File) {
		frame := NewStackFrame(f.Function, f.File, f.Line)

		return append(trace, frame)
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

// CaptureStackTrace captures caller frames starting at skip offset.
func CaptureStackTrace(skip int) StackTrace {
	pc := make([]uintptr, 32)
	n := runtime.Callers(skip+1, pc)
	if n == 0 {
		return StackTrace{}
	}

	return parseFrames(runtime.CallersFrames(pc[:n]))
}

// CaptureCaller captures the top caller line string, reusing CaptureStackTrace.
func CaptureCaller(skip int) string {
	trace := CaptureStackTrace(skip + 1)

	return trace.CallerLine()
}

// isAppFrame filters out runtime frames.
func isAppFrame(file string) bool {
	return !strings.Contains(file, "runtime/")
}

// IsDefined reports true if the collection contains at least one frame.
func (st StackFrames) IsDefined() bool {
	return len(st) > 0
}

// IsEmpty reports true if the collection has zero frames.
func (st StackFrames) IsEmpty() bool {
	return len(st) == 0
}

// CallerLine returns a compact "file:line" string of the top frame.
func (st StackFrames) CallerLine() string {
	if st.IsEmpty() {
		return "unknown:0"
	}

	return fmt.Sprintf("%s:%d", st[0].File, st[0].Line)
}

// Format formats each frame with the provided indentation prefix.
func (st StackFrames) Format(indent string) string {
	if st.IsEmpty() {
		return ""
	}

	var builder strings.Builder
	for idx, frame := range st {
		builder.WriteString(fmt.Sprintf("%s#%d %s\n%s   %s:%d\n", indent, idx, frame.Function, indent, frame.File, frame.Line))
	}

	return builder.String()
}

// String formats the multi-line stack trace with default indentation.
func (st StackFrames) String() string {
	return st.Format("")
}

// ToJson exports StackTrace as indented JSON bytes.
func (st StackTrace) ToJson() ([]byte, error) {
	return json.MarshalIndent(st, "", "  ")
}

// ToJsonString exports StackTrace as a JSON string.
func (st StackTrace) ToJsonString() string {
	b, err := st.ToJson()
	if err != nil {
		return "[]"
	}

	return string(b)
}

// StackTraceFromJson parses StackTrace from JSON bytes.
func StackTraceFromJson(data []byte) (StackTrace, error) {
	var st StackTrace
	if err := json.Unmarshal(data, &st); err != nil {
		return nil, err
	}

	return st, nil
}

// ToJson exports StackFrame as indented JSON bytes.
func (sf StackFrame) ToJson() ([]byte, error) {
	return json.MarshalIndent(sf, "", "  ")
}

// ToJsonString exports StackFrame as a JSON string.
func (sf StackFrame) ToJsonString() string {
	b, err := sf.ToJson()
	if err != nil {
		return "{}"
	}

	return string(b)
}
