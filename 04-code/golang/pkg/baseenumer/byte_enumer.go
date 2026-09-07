package baseenumer

type (
	// ByteEnumer defines the interface for byte-backed enumerations.
	ByteEnumer interface {
		BaseEnumer
		Byte() byte
		ValueByte() byte
		Bytes() []byte
	}

	// ByteEnum is an alias for ByteEnumer.
	ByteEnum = ByteEnumer
)
