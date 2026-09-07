package baseenumer

type (
	// Utf32Enumer defines the interface for UTF-32 / rune-backed enumerations.
	Utf32Enumer interface {
		BaseEnumer
		Rune() rune
		ValueRune() rune
		Int32() int32
	}

	// Utf32Enum is an alias for Utf32Enumer.
	Utf32Enum = Utf32Enumer

	// UTF32Enumer is an alias for backwards compatibility.
	UTF32Enumer = Utf32Enumer

	// UTF32Enum is an alias for backwards compatibility.
	UTF32Enum = Utf32Enum

	// RuneEnumer defines the interface for rune-backed enumerations.
	RuneEnumer interface {
		Utf32Enumer
	}

	// RuneEnum is an alias for RuneEnumer.
	RuneEnum = RuneEnumer
)
