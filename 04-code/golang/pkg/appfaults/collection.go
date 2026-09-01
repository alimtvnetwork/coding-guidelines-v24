package appfaults

import "coding-guidelines/common/pkg/appfault"

// Collection holds an ordered slice of AppError pointers.
type Collection struct {
	items []*appfault.AppError
}

// AppFaults is an alias for Collection for domain consistency.
type AppFaults = Collection

// New creates an empty, non-nil error collection.
func New() *Collection {
	return &Collection{
		items: make([]*appfault.AppError, 0),
	}
}

// NewWithCapacity preallocates backing slice capacity.
func NewWithCapacity(capacity int) *Collection {
	return &Collection{
		items: make([]*appfault.AppError, 0, capacity),
	}
}

// NewFromFaults constructs a collection filtering out nil errors.
func NewFromFaults(faults ...*appfault.AppError) *Collection {
	c := NewWithCapacity(len(faults))
	c.AddAll(faults...)

	return c
}

// NewFromErrors wraps raw Go errors into AppError instances.
func NewFromErrors(errs ...error) *Collection {
	c := NewWithCapacity(len(errs))
	for _, err := range errs {
		c.AddError(err)
	}

	return c
}
