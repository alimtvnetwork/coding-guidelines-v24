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

	// UTF8Enumer defines the interface for UTF-8 byte-backed enumerations.
	UTF8Enumer interface {
		ByteEnumer
	}

	// UTF8Enum is an alias for UTF8Enumer.
	UTF8Enum = UTF8Enumer
)
