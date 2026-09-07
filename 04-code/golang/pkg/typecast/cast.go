package typecast

import (
	"errors"
	"fmt"
	"reflect"
)

// ReflectTo precisely casts a generic payload into the target type T using Go's reflection API.
// It returns a standard Go error to prevent import cycle issues with appfault.
func ReflectTo[T any](payload any) (T, error) {
	var target T

	if payload == nil {
		return target, errors.New("cannot reflect nil payload to target type")
	}

	payloadVal := reflect.ValueOf(payload)
	targetType := reflect.TypeOf(&target).Elem() // Get the actual type of T (handles interface types better)

	// If the payload is already the target type, just return it.
	if payloadVal.Type() == targetType || payloadVal.Type().AssignableTo(targetType) {
		return payload.(T), nil
	}

	if !payloadVal.Type().ConvertibleTo(targetType) {
		return target, fmt.Errorf("type %v is not convertible to %v", payloadVal.Type(), targetType)
	}

	// This panics if it cannot convert, but we checked ConvertibleTo.
	converted := payloadVal.Convert(targetType)
	if !converted.IsValid() || !converted.CanInterface() {
		return target, fmt.Errorf("reflection conversion to %v failed", targetType)
	}

	result, ok := converted.Interface().(T)
	if !ok {
		return target, fmt.Errorf("type assertion post-conversion to %v failed", targetType)
	}

	return result, nil
}

// CastTo safely casts a generic payload into the target type T using type assertion.
func CastTo[T any](payload any) (T, error) {
	var target T

	if payload == nil {
		return target, errors.New("cannot cast nil payload to target type")
	}

	v, ok := payload.(T)
	if !ok {
		return target, fmt.Errorf("type assertion to %T failed for payload of type %T", target, payload)
	}

	return v, nil
}
