package typecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrDestinationNotPointer = errors.New("destination must be a pointer")
	ErrInvalidNullPointer    = errors.New("destination pointer is null")
	ErrInvalidValueType      = errors.New("source 'from' cannot be nil when destination is non-nil")
	ErrTypeMismatch          = errors.New("types are not compatible for reflection set")
)

// ReflectSetTo precisely and dynamically transfers the value of 'from' into 'toPointer'
// using Go's reflection API. It supports pointer-to-pointer, value-to-pointer, and 
// automatic JSON marshaling/unmarshaling for byte slices.
func ReflectSetTo(from, toPointer any) error {
	// 1. Null check bypass (both nil)
	if from == nil && toPointer == nil {
		return nil
	}

	// 2. Validate destination is a non-nil pointer
	if toPointer == nil {
		return ErrInvalidNullPointer
	}

	rightRfType := reflect.TypeOf(toPointer)
	if rightRfType.Kind() != reflect.Ptr {
		return ErrTypeMismatch
	}

	// 3. Validate source is not nil (since dest is not nil)
	if from == nil {
		return ErrInvalidValueType
	}

	leftRfType := reflect.TypeOf(from)
	leftRv := reflect.ValueOf(from)
	rightRv := reflect.ValueOf(toPointer)

	// 4. Same pointer types — direct set
	if leftRfType == rightRfType {
		rightRv.Elem().Set(leftRv.Elem())
		return nil
	}

	// 5. Non-pointer source, pointer destination of same base type
	if leftRfType.Kind() != reflect.Ptr && leftRfType == rightRfType.Elem() {
		rightRv.Elem().Set(leftRv)
		return nil
	}

	// 6. Byte-slice based marshaling/unmarshaling
	var emptyBytes []byte
	emptyBytesType := reflect.TypeOf(emptyBytes)
	emptyBytesPointerType := reflect.TypeOf(&emptyBytes)

	isLeftBytes := leftRfType == emptyBytesType
	isRightBytesPointer := rightRfType == emptyBytesPointerType

	if !(leftRfType == rightRfType || isLeftBytes || isRightBytesPointer) {
		return ErrTypeMismatch
	}

	// Case: []byte → other type (unmarshal)
	if isLeftBytes {
		return json.Unmarshal(from.([]byte), toPointer)
	}

	// Case: other type → *[]byte (marshal)
	if isRightBytesPointer {
		rawBytes, err := json.Marshal(from)
		if err != nil {
			return fmt.Errorf("failed to marshal source to bytes: %w", err)
		}

		bytesPtr := toPointer.(*[]byte)
		*bytesPtr = rawBytes
		return nil
	}

	return ErrTypeMismatch
}

// ReflectTo precisely casts a generic payload into the target type T using ReflectSetTo.
func ReflectTo[T any](payload any) (T, error) {
	var target T
	err := ReflectSetTo(payload, &target)
	return target, err
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
