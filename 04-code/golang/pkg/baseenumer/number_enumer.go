package baseenumer

type (
	// NumberEnumer defines the interface for numeric-backed enumerations.
	NumberEnumer interface {
		BaseEnumer
		Int() int
		Code() uint16
	}

	// NumberEnum is an alias for NumberEnumer.
	NumberEnum = NumberEnumer

	// IntEnumer defines the interface for integer-backed enumerations.
	IntEnumer interface {
		NumberEnumer
	}

	// IntEnum is an alias for IntEnumer.
	IntEnum = IntEnumer
)
