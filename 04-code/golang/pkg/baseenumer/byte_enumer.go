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

	// Utf8Enumer defines the interface for UTF-8 byte-backed enumerations.
	Utf8Enumer interface {
		ByteEnumer
	}

	// Utf8Enum is an alias for Utf8Enumer.
	Utf8Enum = Utf8Enumer

	// UTF8Enumer is an alias for backwards compatibility.
	UTF8Enumer = Utf8Enumer

	// UTF8Enum is an alias for backwards compatibility.
	UTF8Enum = Utf8Enum
)
