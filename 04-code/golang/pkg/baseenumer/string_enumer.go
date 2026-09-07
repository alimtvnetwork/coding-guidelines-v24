package baseenumer

type (
	// StringEnumer defines the interface for string-backed enumerations.
	StringEnumer interface {
		BaseEnumer
	}

	// StringEnum is an alias for StringEnumer.
	StringEnum = StringEnumer
)
