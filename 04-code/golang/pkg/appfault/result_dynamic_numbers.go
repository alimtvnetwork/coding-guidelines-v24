package appfault

import (
	"strconv"
)

func parseIntString(s string) (int, bool) {
	parsedInt, err := strconv.Atoi(s)
	if err == nil {
		return parsedInt, true
	}

	parsedFloat, floatErr := strconv.ParseFloat(s, 64)
	if floatErr == nil {
		return int(parsedFloat), true
	}

	return 0, false
}

func toIntSignedSmall(val any) (int, bool) {
	switch v := val.(type) {
	case int:
		return v, true
	case int8:
		return int(v), true
	case int16:
		return int(v), true
	default:
		return 0, false
	}
}

func toIntSigned(val any) (int, bool) {
	if v, isOk := toIntSignedSmall(val); isOk {
		return v, true
	}

	switch v := val.(type) {
	case int32:
		return int(v), true
	case int64:
		return int(v), true
	default:
		return 0, false
	}
}

func toIntUnsignedSmall(val any) (int, bool) {
	switch v := val.(type) {
	case uint:
		return int(v), true
	case uint8:
		return int(v), true
	case uint16:
		return int(v), true
	default:
		return 0, false
	}
}

func toIntUnsigned(val any) (int, bool) {
	if v, isOk := toIntUnsignedSmall(val); isOk {
		return v, true
	}

	switch v := val.(type) {
	case uint32:
		return int(v), true
	case uint64:
		return int(v), true
	default:
		return 0, false
	}
}

func toIntFloat(val any) (int, bool) {
	switch v := val.(type) {
	case float32:
		return int(v), true
	case float64:
		return int(v), true
	default:
		return 0, false
	}
}

func toIntBoolOrString(val any) (int, bool) {
	switch v := val.(type) {
	case bool:
		if v {
			return 1, true
		}

		return 0, true
	case string:
		return parseIntString(v)
	default:
		return 0, false
	}
}

func toIntFloatOrBool(val any) (int, bool) {
	if v, isOk := toIntFloat(val); isOk {
		return v, true
	}

	return toIntBoolOrString(val)
}

func toInt(val any) (int, bool) {
	if val == nil {
		return 0, false
	}

	if v, isOk := toIntSigned(val); isOk {
		return v, true
	}

	if v, isOk := toIntUnsigned(val); isOk {
		return v, true
	}

	return toIntFloatOrBool(val)
}

func parseInt64String(s string) (int64, bool) {
	parsedInt, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return parsedInt, true
	}

	parsedFloat, floatErr := strconv.ParseFloat(s, 64)
	if floatErr == nil {
		return int64(parsedFloat), true
	}

	return 0, false
}

func toInt64Direct(val any) (int64, bool) {
	switch v := val.(type) {
	case int64:
		return v, true
	case int:
		return int64(v), true
	case uint64:
		return int64(v), true
	case float64:
		return int64(v), true
	default:
		return 0, false
	}
}

func toInt64Fallback(val any) (int64, bool) {
	asInt, isOk := toInt(val)

	return int64(asInt), isOk
}

func toInt64(val any) (int64, bool) {
	if val == nil {
		return 0, false
	}

	if s, isStr := val.(string); isStr {
		return parseInt64String(s)
	}

	if v, isOk := toInt64Direct(val); isOk {
		return v, true
	}

	return toInt64Fallback(val)
}

func toFloat64Direct(val any) (float64, bool) {
	switch v := val.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	default:
		return 0, false
	}
}

func parseFloatString(s string) (float64, bool) {
	f, err := strconv.ParseFloat(s, 64)

	return f, err == nil
}

func toFloat64Fallback(val any) (float64, bool) {
	asInt, isOk := toInt(val)

	return float64(asInt), isOk
}

func toFloat64(val any) (float64, bool) {
	if val == nil {
		return 0, false
	}

	if v, isOk := toFloat64Direct(val); isOk {
		return v, true
	}

	if str, isStr := val.(string); isStr {
		return parseFloatString(str)
	}

	return toFloat64Fallback(val)
}

func toByte(val any) (byte, bool) {
	if val == nil {
		return 0, false
	}

	asInt, isOk := toInt(val)
	if isOk && asInt >= 0 && asInt <= 255 {
		return byte(asInt), true
	}

	return 0, false
}

// Int converts the payload value to an int.
func (r Result[T]) Int() (int, bool) {
	if r.IsFailed() {
		return 0, false
	}

	return toInt(r.value)
}

// IntDefault returns the payload as an int, or defaultVal on failure.
func (r Result[T]) IntDefault(defaultVal int) int {
	val, isOk := r.Int()
	if isOk {
		return val
	}

	return defaultVal
}

// Int64 converts the payload value to an int64.
func (r Result[T]) Int64() (int64, bool) {
	if r.IsFailed() {
		return 0, false
	}

	return toInt64(r.value)
}

// Int64Default returns the payload as an int64, or defaultVal on failure.
func (r Result[T]) Int64Default(defaultVal int64) int64 {
	val, isOk := r.Int64()
	if isOk {
		return val
	}

	return defaultVal
}

// Float64 converts the payload value to a float64.
func (r Result[T]) Float64() (float64, bool) {
	if r.IsFailed() {
		return 0, false
	}

	return toFloat64(r.value)
}

// Float64Default returns the payload as a float64, or defaultVal on failure.
func (r Result[T]) Float64Default(defaultVal float64) float64 {
	val, isOk := r.Float64()
	if isOk {
		return val
	}

	return defaultVal
}

// Double is an alias for Float64.
func (r Result[T]) Double() (float64, bool) {
	return r.Float64()
}

// DoubleDefault is an alias for Float64Default.
func (r Result[T]) DoubleDefault(defaultVal float64) float64 {
	return r.Float64Default(defaultVal)
}

// Byte converts the payload value to a byte.
func (r Result[T]) Byte() (byte, bool) {
	if r.IsFailed() {
		return 0, false
	}

	return toByte(r.value)
}

// ByteDefault returns the payload as a byte, or defaultVal on failure.
func (r Result[T]) ByteDefault(defaultVal byte) byte {
	val, isOk := r.Byte()
	if isOk {
		return val
	}

	return defaultVal
}
