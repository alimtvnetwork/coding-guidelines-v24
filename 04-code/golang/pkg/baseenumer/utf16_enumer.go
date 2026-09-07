package baseenumer

type (
	// Utf16Enumer defines the interface for UTF-16 code-unit-backed enumerations.
	Utf16Enumer interface {
		BaseEnumer
		Utf16() uint16
		ValueUtf16() uint16
		Code() uint16
	}

	// Utf16Enum is an alias for Utf16Enumer.
	Utf16Enum = Utf16Enumer

	// UTF16Enumer is an alias for backwards compatibility.
	UTF16Enumer = Utf16Enumer

	// UTF16Enum is an alias for backwards compatibility.
	UTF16Enum = Utf16Enum
)
