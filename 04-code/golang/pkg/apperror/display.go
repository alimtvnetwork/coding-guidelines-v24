package apperror

import (
	"fmt"
	"strings"
)

// Error implements the standard Go error interface.
func (e *AppError) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

// appendContext formats context map entries into the string builder.
func appendContext(b *strings.Builder, ctx map[string]any) {
	if len(ctx) > 0 {
		b.WriteString("CONTEXT:\n")
		for k, v := range ctx {
			b.WriteString(fmt.Sprintf("  • %s: %v\n", k, v))
		}
	}
}

// appendHeader writes standard error prefix and optional cause.
func appendHeader(b *strings.Builder, code ErrorCodeType, msg string, cause error) {
	b.WriteString(fmt.Sprintf("ERROR: [%s] %s\n", code, msg))
	if cause != nil {
		b.WriteString(fmt.Sprintf("CAUSE: %v\n", cause))
	}
}

// FullString returns a comprehensive diagnostic dump of the error.
func (e *AppError) FullString() string {
	if e == nil {
		return ""
	}

	var b strings.Builder
	appendHeader(&b, e.code, e.message, e.cause)
	appendContext(&b, e.context)
	if !e.stackTrace.IsEmpty() {
		b.WriteString("STACK TRACE:\n" + e.stackTrace.String())
	}

	return b.String()
}

// appendMarkdownHeader writes Markdown formatted title and cause.
func appendMarkdownHeader(b *strings.Builder, code ErrorCodeType, msg string, cause error) {
	b.WriteString(fmt.Sprintf("### Error Report\n\n- **Code:** `%s`\n- **Message:** %s\n", code, msg))
	if cause != nil {
		b.WriteString(fmt.Sprintf("- **Cause:** `%v`\n", cause))
	}
}

// appendMarkdownContext formats context map into Markdown list.
func appendMarkdownContext(b *strings.Builder, ctx map[string]any) {
	if len(ctx) > 0 {
		b.WriteString("\n#### Context\n\n")
		for k, v := range ctx {
			b.WriteString(fmt.Sprintf("- `%s`: `%v`\n", k, v))
		}
	}
}

// ToClipboard returns a Markdown-formatted error report for AI analysis.
func (e *AppError) ToClipboard() string {
	if e == nil {
		return ""
	}

	var b strings.Builder
	appendMarkdownHeader(&b, e.code, e.message, e.cause)
	appendMarkdownContext(&b, e.context)
	if !e.stackTrace.IsEmpty() {
		b.WriteString("\n```\n" + e.stackTrace.String() + "```\n")
	}

	return b.String()
}

// DisplayError prints a human-readable banner representation to stdout.
func (e *AppError) DisplayError() {
	if e != nil {
		fmt.Printf("❌ [%s] %s (at %s)\n", e.code, e.message, e.stackTrace.CallerLine())
	}
}
