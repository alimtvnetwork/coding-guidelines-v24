package appfault

import (
	"fmt"
	"reflect"
)

// Type returns the reflection Type of the payload.
func (r Result[T]) Type() reflect.Type {
	if any(r.value) == nil {
		return nil
	}

	return reflect.TypeOf(r.value)
}

// TypeName returns the qualified type name of the payload.
func (r Result[T]) TypeName() string {
	t := r.Type()
	if t == nil {
		return "nil"
	}

	return fmt.Sprintf("%T", r.value)
}

// Kind returns the reflection Kind of the payload.
func (r Result[T]) Kind() reflect.Kind {
	if any(r.value) == nil {
		return reflect.Invalid
	}

	return reflect.ValueOf(r.value).Kind()
}

func dereferenceValue(rv reflect.Value) (reflect.Value, bool) {
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return rv, false
		}

		return rv.Elem(), true
	}

	return rv, true
}

func lenOfKind(rv reflect.Value) int {
	switch rv.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String, reflect.Chan:
		return rv.Len()
	default:
		return 0
	}
}

// Length returns the length of slice, array, map, string, or channel payload.
func (r Result[T]) Length() int {
	if any(r.value) == nil {
		return 0
	}

	rv, isOk := dereferenceValue(reflect.ValueOf(r.value))
	if !isOk {
		return 0
	}

	return lenOfKind(rv)
}

func isNumberKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsNumber reports whether the payload is any numeric kind.
func (r Result[T]) IsNumber() bool {
	return isNumberKind(r.Kind())
}

// IsStringType reports whether the payload is a string.
func (r Result[T]) IsStringType() bool {
	return r.Kind() == reflect.String
}

// IsSliceOrArray reports whether the payload is a slice or array.
func (r Result[T]) IsSliceOrArray() bool {
	k := r.Kind()

	return k == reflect.Slice || k == reflect.Array
}

// IsMap reports whether the payload is a map.
func (r Result[T]) IsMap() bool {
	return r.Kind() == reflect.Map
}

// IsStruct reports whether the payload is a struct.
func (r Result[T]) IsStruct() bool {
	return r.Kind() == reflect.Struct
}

// IsPointer reports whether the payload is a pointer.
func (r Result[T]) IsPointer() bool {
	return r.Kind() == reflect.Ptr
}

// IsPrimitive reports whether the payload is a primitive type.
func (r Result[T]) IsPrimitive() bool {
	k := r.Kind()
	if k == reflect.Bool || k == reflect.String {
		return true
	}

	return isNumberKind(k)
}
