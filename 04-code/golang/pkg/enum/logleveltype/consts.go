package logleveltype

// Enum constants for logleveltype conforming to aukgo and global standards.
const (
	// Invalid represents an uninitialized or invalid log level (iota zero-value).
	Invalid Variant = iota

	// Debug represents verbose diagnostic messages.
	Debug

	// Info represents normal operational events.
	Info

	// Warn represents non-fatal warnings or unexpected conditions.
	Warn

	// Error represents operation failures requiring attention.
	Error

	// Fatal represents critical errors causing immediate shutdown.
	Fatal
)

// Unknown is an alias to Invalid for backward compatibility.
const Unknown = Invalid

// VariantPredicate defines a filter or condition check over a Variant.
type VariantPredicate func(v Variant) bool
