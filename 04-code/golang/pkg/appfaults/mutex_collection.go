package appfaults

import (
	"sync"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// MutexCollection provides concurrency-safe operations over a Collection.
type LockCollection struct {
	sync.RWMutex
	inner *Collection
}

// NewMutexCollection creates a thread-safe error collection.
func NewLockCollection() *LockCollection {
	return &LockCollection{
		inner: New(),
	}
}

// Add safely appends an error to the collection.
func (mc *LockCollection) Add(err *appfault.AppError) *LockCollection {
	mc.Lock()
	defer mc.Unlock()
	mc.inner.Add(err)

	return mc
}

// AddType safely adds an error by type.
func (mc *LockCollection) AddType(errType errtype.Variation) *LockCollection {
	mc.Lock()
	defer mc.Unlock()
	mc.inner.AddType(errType)

	return mc
}

// AddTypeMsg safely adds an error by type and message.
func (mc *LockCollection) AddTypeMsg(errType errtype.Variation, msg string) *LockCollection {
	mc.Lock()
	defer mc.Unlock()
	mc.inner.AddTypeMsg(errType, msg)

	return mc
}

// AddError safely wraps and appends a standard error with explicit type.
func (mc *LockCollection) AddError(errType errtype.Variation, cause error) *LockCollection {
	mc.Lock()
	defer mc.Unlock()
	mc.inner.AddError(errType, cause)

	return mc
}

// HasError safely checks if any error is stored.
func (mc *LockCollection) HasError() bool {
	mc.RLock()
	defer mc.RUnlock()

	return mc.inner.HasError()
}

// IsSuccess safely checks if collection is empty.
func (mc *LockCollection) IsSuccess() bool {
	return !mc.HasError()
}

// Count safely returns the item count.
func (mc *LockCollection) Count() int {
	mc.RLock()
	defer mc.RUnlock()

	return mc.inner.Count()
}

// Snapshot safely returns a cloned Collection.
func (mc *LockCollection) Snapshot() *Collection {
	mc.RLock()
	defer mc.RUnlock()

	return NewFromFaults(mc.inner.items...)
}

// MutexCollection is an alias for LockCollection for backward compatibility.
type MutexCollection = LockCollection

// NewMutexCollection creates a thread-safe error collection (backward compatible).
func NewMutexCollection() *LockCollection {
	return NewLockCollection()
}
