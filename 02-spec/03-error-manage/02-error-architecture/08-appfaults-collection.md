# Specification: `appfaults` Error Collection Architecture

**Version:** 3.3.0  
**Status:** Draft Specification (Pending Review)  
**Package:** `04-code/golang/pkg/appfaults`  
**Reference Implementations:** `https://gitlab.com/auk-go/errorwrapper` (`errwrappers.Collection`, `errwrappers.MutexCollection`)

---

## 1. Executive Summary & Purpose

In multi-step workflows, batch validations, data pipeline migrations, and concurrent task executions, a single function often encounters multiple errors rather than failing on the very first issue.

The **`appfaults`** package provides a first-class, memory-efficient collection type (`appfaults.Collection` / `appfaults.AppFaults`) for aggregating, filtering, transforming, and propagating multiple `*appfault.AppError` instances across application layers and Go context (`context.Context`).

---

## 2. Core Architecture & Interfaces

### 2.1 Collection Structs

```go
package appfaults

import (
	"sync"
	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// Collection holds an ordered slice of AppError pointers.
type Collection struct {
	items []*appfault.AppError
}

// AppFaults is an alias for Collection for domain consistency.
type AppFaults = Collection

// MutexCollection provides concurrency-safe operations over Collection.
type MutexCollection struct {
	sync.RWMutex
	inner Collection
}
```

---

## 3. Method Specifications

### 3.1 Constructors

| Constructor | Signature | Description |
| :--- | :--- | :--- |
| `New()` | `func New() *Collection` | Creates an empty, non-nil error collection. |
| `NewWithCapacity(n)` | `func NewWithCapacity(capacity int) *Collection` | Preallocates backing slice capacity. |
| `NewFromFaults(faults...)` | `func NewFromFaults(faults ...*appfault.AppError) *Collection` | Constructs collection filtering out nil errors. |
| `NewFromErrors(errs...)` | `func NewFromErrors(errs ...error) *Collection` | Wraps raw Go errors into `AppError` and stores. |
| `NewMutexCollection()` | `func NewMutexCollection() *MutexCollection` | Creates a thread-safe mutex-wrapped collection. |

### 3.2 Mutators & Aggregation

| Method | Signature | Description |
| :--- | :--- | :--- |
| `Add(err)` | `func (c *Collection) Add(err *appfault.AppError) *Collection` | Appends error if non-nil and `HasError()` is true. |
| `AddError(err)` | `func (c *Collection) AddError(err error) *Collection` | Wraps raw error and appends if non-nil. |
| `AddAll(faults...)` | `func (c *Collection) AddAll(faults ...*appfault.AppError) *Collection` | Appends multiple errors in order. |
| `Merge(other)` | `func (c *Collection) Merge(other *Collection) *Collection` | Ingests all items from another collection. |
| `Clear()` | `func (c *Collection) Clear() *Collection` | Resets the collection to empty. |

### 3.3 Status & Introspection

| Method | Signature | Semantics / Behavior |
| :--- | :--- | :--- |
| `HasError()` | `func (c *Collection) HasError() bool` | Returns `true` if collection contains $\ge 1$ active errors. Returns `false` if `c == nil` or empty. |
| `IsSuccess()` | `func (c *Collection) IsSuccess() bool` | Returns `true` if `c == nil` or `Count() == 0` (no errors). |
| `IsEmpty()` | `func (c *Collection) IsEmpty() bool` | Alias for `!c.HasError()`. |
| `Count()` | `func (c *Collection) Count() int` | Returns total number of active errors. |
| `Items()` | `func (c *Collection) Items() []*appfault.AppError` | Returns copy of backing slice. |
| `First()` | `func (c *Collection) First() *appfault.AppError` | Returns first error or `nil`. |
| `Last()` | `func (c *Collection) Last() *appfault.AppError` | Returns last error or `nil`. |

### 3.4 Filtering & Transformation

```go
// Filter returns a new Collection containing items that satisfy predicate.
func (c *Collection) Filter(predicate func(*appfault.AppError) bool) *Collection

// FilterByType returns errors matching a specific errtype.Variation.
func (c *Collection) FilterByType(errType errtype.Variation) *Collection

// ToAppError merges the collection into a single composite AppError.
func (c *Collection) ToAppError(compositeType errtype.Variation, msg string) *appfault.AppError

// Errors converts collection into a standard []error slice.
func (c *Collection) Errors() []error
```

---

## 4. Context Integration (`context.Context`)

`appfaults` attaches directly to Go `context.Context` to aggregate errors across middleware, interceptors, and nested routines without passing mutable collectors through every function signature:

```go
// Key for context error collection storage
type contextKey struct{}

// WithFaults creates a child context holding a dedicated AppFaults collector.
func WithFaults(ctx context.Context) (context.Context, *Collection)

// FromContext extracts the existing AppFaults collector from context.
func FromContext(ctx context.Context) (*Collection, bool)

// RecordContextError records an error into context collector if present.
func RecordContextError(ctx context.Context, err *appfault.AppError) bool
```

---

## 5. Serialization & Diagnostics

- **JSON Output:** Serializes collection to PascalCase array:
  ```json
  [
    {
      "Type": 2,
      "Message": "email format is invalid",
      "Caller": "validator.go:28"
    },
    {
      "Type": 2,
      "Message": "password too short",
      "Caller": "validator.go:34"
    }
  ]
  ```
- **Formatted Summary:** `Format()` provides newline-separated diagnostic logs with bullet points for user-facing and AI analysis.
