package apperror

import (
	"fmt"
	"strings"
)

// formatBasicError returns formatted code, type, op, and message.
func formatBasicError(f *Fault) string {
	return fmt.Sprintf("[%s:%s] %s: %s", f.Code, f.Type, f.Op, f.Message)
}

// appendCallerAndCause appends caller site and cause to formatted string.
func appendCallerAndCause(base, caller string, cause error) string {
	if len(caller) > 0 {
		base += fmt.Sprintf(" (at=%s)", caller)
	}

	if cause != nil {
		base += fmt.Sprintf(" (cause=%v)", cause)
	}

	return base
}

// Error implements the standard Go error interface.
func (f *Fault) Error() string {
	if f == nil {
		return ""
	}

	return appendCallerAndCause(formatBasicError(f), f.Caller, f.Cause)
}

// appendHeader writes diagnostic header info.
func appendHeader(b *strings.Builder, f *Fault) {
	b.WriteString(fmt.Sprintf("ERROR: [%s:%s] %s: %s\n", f.Code, f.Type, f.Op, f.Message))
	if len(f.Caller) > 0 {
		b.WriteString(fmt.Sprintf("CALLER: %s\n", f.Caller))
	}

	if f.Cause != nil {
		b.WriteString(fmt.Sprintf("CAUSE: %v\n", f.Cause))
	}
}

// appendContextAndStack writes context map and stack trace.
func appendContextAndStack(b *strings.Builder, ctx map[string]any, stack string) {
	if len(ctx) > 0 {
		b.WriteString(fmt.Sprintf("CONTEXT: %v\n", ctx))
	}

	if len(stack) > 0 {
		b.WriteString("STACK TRACE:\n" + stack)
	}
}

// FullString returns a comprehensive diagnostic dump of the Fault.
func (f *Fault) FullString() string {
	if f == nil {
		return ""
	}

	var b strings.Builder
	appendHeader(&b, f)
	appendContextAndStack(&b, f.Ctx, f.Stack)

	return b.String()
}

// appendMarkdownCauseAndStack writes cause and codeblock stack trace.
func appendMarkdownCauseAndStack(b *strings.Builder, cause error, stack string) {
	if cause != nil {
		b.WriteString(fmt.Sprintf("- **Cause:** `%v`\n", cause))
	}

	if len(stack) > 0 {
		b.WriteString("\n```\n" + stack + "```\n")
	}
}

// ToClipboard returns a Markdown formatted report for AI analysis.
func (f *Fault) ToClipboard() string {
	if f == nil {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("### Error Report\n\n- **Code:** `%s`\n- **Type:** `%s`\n- **Op:** `%s`\n- **Message:** %s\n", f.Code, f.Type, f.Op, f.Message))
	appendMarkdownCauseAndStack(&b, f.Cause, f.Stack)

	return b.String()
}

// DisplayError prints a terminal banner representation.
func (f *Fault) DisplayError() {
	if f != nil {
		fmt.Printf("❌ [%s:%s] %s: %s (at %s)\n", f.Code, f.Type, f.Op, f.Message, f.Caller)
	}
}
