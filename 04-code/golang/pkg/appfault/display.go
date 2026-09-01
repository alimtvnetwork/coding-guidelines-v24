package appfault

import (
	"fmt"
	"strings"
)

// formatBasicError returns formatted code, type, op, and message.
func formatBasicError(e *AppError) string {
	return fmt.Sprintf("[%s:%s] %s: %s", e.Code, e.Type, e.Op, e.Message)
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
func (e *AppError) Error() string {
	if e == nil {
		return ""
	}

	return appendCallerAndCause(formatBasicError(e), e.Caller, e.Cause)
}

// appendHeader writes diagnostic header info.
func appendHeader(b *strings.Builder, e *AppError) {
	b.WriteString(fmt.Sprintf("ERROR: [%s:%s] %s: %s\n", e.Code, e.Type, e.Op, e.Message))
	if len(e.Caller) > 0 {
		b.WriteString(fmt.Sprintf("CALLER: %s\n", e.Caller))
	}

	if e.Cause != nil {
		b.WriteString(fmt.Sprintf("CAUSE: %v\n", e.Cause))
	}
}

// appendContextAndStack writes context map and stack trace.
func appendContextAndStack(b *strings.Builder, ctx ContextMap, stack string) {
	if len(ctx) > 0 {
		b.WriteString(fmt.Sprintf("CONTEXT: %s\n", ctx.Format()))
	}

	if len(stack) > 0 {
		b.WriteString("STACK TRACE:\n" + stack)
	}
}

// FullString returns a comprehensive diagnostic dump of the AppError.
func (e *AppError) FullString() string {
	if e == nil {
		return ""
	}

	var b strings.Builder
	appendHeader(&b, e)
	appendContextAndStack(&b, e.Ctx, e.Stack)

	return b.String()
}

// appendMarkdownCauseAndStack writes cause and codeblock stack trace.
func appendMarkdownCauseAndStack(b *strings.Builder, cause error, stack string) {
	if cause != nil {
		b.WriteString(fmt.Sprintf("- **Cause:** %v\n", cause))
	}

	if len(stack) > 0 {
		b.WriteString("\n`\n" + stack + "`\n")
	}
}

// ToClipboard returns a Markdown formatted report for AI analysis.
func (e *AppError) ToClipboard() string {
	if e == nil {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("### Error Report\n\n- **Code:** %s\n- **Type:** %s\n- **Op:** %s\n- **Message:** %s\n", e.Code, e.Type, e.Op, e.Message))
	appendMarkdownCauseAndStack(&b, e.Cause, e.Stack)

	return b.String()
}

// DisplayError prints a terminal banner representation.
func (e *AppError) DisplayError() {
	if e != nil {
		fmt.Printf("❌ [%s:%s] %s: %s (at %s)\n", e.Code, e.Type, e.Op, e.Message, e.Caller)
	}
}
