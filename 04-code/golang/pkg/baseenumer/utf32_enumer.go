package baseenumer

type (
	// UTF32Enumer defines the interface for UTF-32 / rune-backed enumerations.
	UTF32Enumer interface {
		BaseEnumer
		Rune() rune
		ValueRune() rune
		Int32() int32
	}

	// UTF32Enum is an alias for UTF32Enumer.
	UTF32Enum = UTF32Enumer

	// RuneEnumer defines the interface for rune-backed enumerations.
	RuneEnumer interface {
		UTF32Enumer
	}

	// RuneEnum is an alias for RuneEnumer.
	RuneEnum = RuneEnumer
)
