package apperror

import (
	"fmt"
	"strings"
)

// Error implements the standard Go error interface.
func (f *Fault) Error() string {
	if f == nil {
		return ""
	}

	return fmt.Sprintf("[%s] %s", f.code, f.message)
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
func (f *Fault) FullString() string {
	if f == nil {
		return ""
	}

	var b strings.Builder
	appendHeader(&b, f.code, f.message, f.cause)
	appendContext(&b, f.context)
	if !f.stackTrace.IsEmpty() {
		b.WriteString("STACK TRACE:\n" + f.stackTrace.String())
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
func (f *Fault) ToClipboard() string {
	if f == nil {
		return ""
	}

	var b strings.Builder
	appendMarkdownHeader(&b, f.code, f.message, f.cause)
	appendMarkdownContext(&b, f.context)
	if !f.stackTrace.IsEmpty() {
		b.WriteString("\n```\n" + f.stackTrace.String() + "```\n")
	}

	return b.String()
}

// DisplayError prints a human-readable banner representation to stdout.
func (f *Fault) DisplayError() {
	if f != nil {
		fmt.Printf("❌ [%s] %s (at %s)\n", f.code, f.message, f.stackTrace.CallerLine())
	}
}
