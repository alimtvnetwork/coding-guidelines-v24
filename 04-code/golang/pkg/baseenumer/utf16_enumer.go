package baseenumer

type (
	// UTF16Enumer defines the interface for UTF-16 code-unit-backed enumerations.
	UTF16Enumer interface {
		BaseEnumer
		UTF16() uint16
		ValueUTF16() uint16
		Code() uint16
	}

	// UTF16Enum is an alias for UTF16Enumer.
	UTF16Enum = UTF16Enumer
)
