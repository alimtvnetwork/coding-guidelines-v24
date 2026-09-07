package baseenumer

// IntNumber represents signed and unsigned integer types used for enum variants.
type IntNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// IntegerVarianter is an alias forwarder for IntNumber for non-breaking compatibility.
type IntegerVarianter = IntNumber
