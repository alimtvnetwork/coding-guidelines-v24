package baseenumer

type (
	// MinMaxer defines boundary value retrieval for enums.
	MinMaxer[V any] interface {
		Min() V
		Max() V
	}

	// MinMax is an alias for MinMaxer.
	MinMax[V any] = MinMaxer[V]

	// BoundedEnumer defines boundary checks for enum instances.
	BoundedEnumer[V any] interface {
		MinMaxer[V]
		IsMin() bool
		IsMax() bool
	}

	// BoundedEnum is an alias for BoundedEnumer.
	BoundedEnum[V any] = BoundedEnumer[V]

	// Bounder defines range checking for enum instances.
	Bounder[V any] interface {
		BoundedEnumer[V]
		IsInRange(min, max V) bool
	}
)
