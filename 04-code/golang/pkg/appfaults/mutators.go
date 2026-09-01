package appfaults

import "coding-guidelines/common/pkg/appfault"

// Add appends an AppError if it is non-nil and has an active error.
func (c *Collection) Add(err *appfault.AppError) *Collection {
	if c == nil || err == nil || !err.HasError() {
		return c
	}

	c.items = append(c.items, err)

	return c
}

// AddError wraps a standard error and appends it.
func (c *Collection) AddError(err error) *Collection {
	if c == nil || err == nil {
		return c
	}

	return c.Add(appfault.WrapSimple(err))
}

// AddAll appends multiple AppErrors in order.
func (c *Collection) AddAll(faults ...*appfault.AppError) *Collection {
	for _, f := range faults {
		c.Add(f)
	}

	return c
}

// Merge copies all items from another collection into c.
func (c *Collection) Merge(other *Collection) *Collection {
	if c == nil || other == nil {
		return c
	}

	return c.AddAll(other.items...)
}

// Clear removes all items from the collection.
func (c *Collection) Clear() *Collection {
	if c != nil {
		c.items = c.items[:0]
	}

	return c
}
