package typecast

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

var (
	emptyBytesType        = reflect.TypeOf([]byte(nil))
	emptyBytesPointerType = reflect.TypeOf((*[]byte)(nil))
)

var (
	ErrDestinationNotPointer = errors.New("destination must be a pointer")
	ErrInvalidNullPointer    = errors.New("destination pointer is null")
	ErrInvalidValueType      = errors.New("source 'from' cannot be nil when destination is non-nil")
	ErrTypeMismatch          = errors.New("types are not compatible for reflection set")
)

// ReflectSetTo precisely and dynamically transfers the value of 'from' into 'toPointer'.
func ReflectSetTo(from, toPointer any) error {
	if from == nil && toPointer == nil {
		return nil
	}

	if toPointer == nil {
		return ErrInvalidNullPointer
	}

	if from == nil {
		return ErrInvalidValueType
	}

	done, err := fastPathSet(from, toPointer)
	if done {
		return err
	}

	return reflectFallbackSet(from, toPointer)
}

func fastPathSet(from, toPointer any) (bool, error) {
	done, err := fastPathPrimitive(from, toPointer)
	if done {
		return true, err
	}

	return fastPathBytesOrUnmarshal(from, toPointer)
}

func fastPathPrimitive(from, toPointer any) (bool, error) {
	switch dest := toPointer.(type) {
	case *string:
		return true, setPrimitive(from, dest)
	case *int:
		return true, setPrimitive(from, dest)
	case *int64:
		return true, setPrimitive(from, dest)
	case *bool:
		return true, setPrimitive(from, dest)
	case *float64:
		return true, setPrimitive(from, dest)
	}

	return false, nil
}

func setFromPtr[T any](fromPtr *T, dest *T) error {
	if fromPtr == nil {
		return ErrInvalidValueType
	}

	*dest = *fromPtr

	return nil
}

func setPrimitive[T any](from any, dest *T) error {
	if dest == nil {
		return ErrInvalidNullPointer
	}

	if v, ok := from.(T); ok {
		*dest = v

		return nil
	}

	if v, ok := from.(*T); ok {
		return setFromPtr(v, dest)
	}

	if b, ok := from.([]byte); ok {
		return json.Unmarshal(b, dest)
	}

	return ErrTypeMismatch
}

func fastPathBytesOrUnmarshal(from, toPointer any) (bool, error) {
	if dest, ok := toPointer.(*[]byte); ok {
		return true, setBytes(from, dest)
	}

	if b, ok := from.([]byte); ok {
		return fastPathUnmarshal(b, toPointer)
	}

	return false, nil
}

func marshalToBytes(from any, dest *[]byte) error {
	rawBytes, err := json.Marshal(from)
	if err != nil {
		return err
	}

	*dest = rawBytes

	return nil
}

func setBytes(from any, dest *[]byte) error {
	if dest == nil {
		return ErrInvalidNullPointer
	}

	if b, ok := from.([]byte); ok {
		*dest = b

		return nil
	}

	if b, ok := from.(*[]byte); ok {
		return setFromPtr(b, dest)
	}

	return marshalToBytes(from, dest)
}

func fastPathUnmarshal(b []byte, toPointer any) (bool, error) {
	rightRfType := reflect.TypeOf(toPointer)
	if rightRfType.Kind() != reflect.Ptr {
		return true, ErrDestinationNotPointer
	}

	if reflect.ValueOf(toPointer).IsNil() {
		return true, ErrInvalidNullPointer
	}

	return true, json.Unmarshal(b, toPointer)
}

func checkPointers(from, toPointer any) error {
	rightRfType := reflect.TypeOf(toPointer)
	if rightRfType.Kind() != reflect.Ptr {
		return ErrDestinationNotPointer
	}

	if reflect.ValueOf(toPointer).IsNil() {
		return ErrInvalidNullPointer
	}

	leftRv := reflect.ValueOf(from)
	if leftRv.Kind() == reflect.Ptr && leftRv.IsNil() {
		return ErrInvalidValueType
	}

	return nil
}

func reflectFallbackSet(from, toPointer any) error {
	err := checkPointers(from, toPointer)
	if err != nil {
		return err
	}

	leftRfType := reflect.TypeOf(from)
	rightRfType := reflect.TypeOf(toPointer)
	rightRv := reflect.ValueOf(toPointer)
	leftRv := reflect.ValueOf(from)

	if leftRfType == rightRfType {
		rightRv.Elem().Set(leftRv.Elem())

		return nil
	}

	if leftRfType == rightRfType.Elem() {
		rightRv.Elem().Set(leftRv)

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

// CastTo safely casts a generic payload into the target type T using fast type assertion,
// falling back to reflection (ReflectTo) if direct type assertion fails.
func CastTo[T any](payload any) (T, error) {
	var target T

	if payload == nil {
		return target, errors.New("cannot cast nil payload to target type")
	}

	v, isOk := payload.(T)
	if isOk {
		return v, nil
	}

	return ReflectTo[T](payload)
}

// ToBytes converts a payload into a byte slice, handling bytes, string, string slice, error, and JSON fallback.
func ToBytes(payload any) ([]byte, error) {
	if payload == nil {
		return []byte{}, nil
	}

	switch v := payload.(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	case []string:
		return []byte(strings.Join(v, "\n")), nil
	case error:
		return []byte(v.Error()), nil
	default:
		return json.Marshal(v)
	}
}

// ToJson formats payload as indented JSON with trailing newline.
func ToJson(payload any) ([]byte, error) {
	b, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return nil, err
	}

	b = append(b, '\n')

	return b, nil
}

// ToJsonString formats payload as indented JSON string.
func ToJsonString(payload any) (string, error) {
	b, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ToJSON is an alias for ToJson.
func ToJSON(payload any) ([]byte, error) {
	return ToJson(payload)
}

// ToJSONString is an alias for ToJsonString.
func ToJSONString(payload any) (string, error) {
	return ToJsonString(payload)
}
