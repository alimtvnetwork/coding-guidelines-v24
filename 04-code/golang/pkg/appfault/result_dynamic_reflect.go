package appfault

import (
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/typecast"
)

// ReflectTo dynamically sets the payload into targetPointer via reflection.
func (r Result[T]) ReflectTo(targetPointer any) *AppError {
	if r.IsFailed() {
		return r.appError
	}

	err := typecast.ReflectSetTo(r.value, targetPointer)
	if err != nil {
		return Wrap(errtype.TypeMismatch, err, err.Error())
	}

	return nil
}

// Map returns the payload converted to map[string]any (alias to ToMap).
func (r Result[T]) Map() map[string]any {
	return r.ToMap()
}

// ToMapResult converts the payload to map[string]any wrapped in a ResultMap.
func (r Result[T]) ToMapResult() ResultMap[string, any] {
	if r.IsFailed() {
		return FailMap[string, any](r.appError)
	}

	return OkMap(r.ToMap())
}
