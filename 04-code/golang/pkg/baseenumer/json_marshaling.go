package baseenumer

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// MarshalJSON serializes an enum's label or string representation to JSON.
func MarshalJSON(name string) ([]byte, error) {
	return json.Marshal(name)
}

// UnmarshalIntegerJSON unmarshals JSON data into an integer-backed enum variant with auto-resolved type name.
func UnmarshalIntegerJSON[V IntNumber](
	data []byte,
	target *V,
	variantMap map[string]V,
	maxValid int,
	zero V,
) error {
	typeName := ResolveTypeName(target)

	return UnmarshalIntegerJSONWithName(data, target, typeName, variantMap, maxValid, zero)
}

// UnmarshalIntegerJSONWithName unmarshals JSON data into an integer-backed enum with an explicit type name.
func UnmarshalIntegerJSONWithName[V IntNumber](
	data []byte,
	target *V,
	typeName string,
	variantMap map[string]V,
	maxValid int,
	zero V,
) error {
	if isNullOrEmpty(data) {
		*target = zero

		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		return unmarshalIntegerString(str, target, typeName, variantMap, maxValid, zero)
	}

	return unmarshalIntegerNumeric(data, target, typeName, maxValid)
}

// UnmarshalStringJSON unmarshals JSON data into a string-backed enum variant with auto-resolved type name.
func UnmarshalStringJSON[V ~string](
	data []byte,
	target *V,
	variantMap map[string]V,
	zero V,
) error {
	typeName := ResolveTypeName(target)

	return UnmarshalStringJSONWithName(data, target, typeName, variantMap, zero)
}

// UnmarshalStringJSONWithName unmarshals JSON data into a string-backed enum with an explicit type name.
func UnmarshalStringJSONWithName[V ~string](
	data []byte,
	target *V,
	typeName string,
	variantMap map[string]V,
	zero V,
) error {
	if isNullOrEmpty(data) {
		*target = zero

		return nil
	}

	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	v, err := parseStringVariant(raw, typeName, variantMap, zero)
	if err != nil {
		return err
	}

	*target = v

	return nil
}

func isNullOrEmpty(data []byte) bool {
	trimmed := strings.TrimSpace(string(data))

	return len(trimmed) == 0 || trimmed == "null"
}

func parseStringVariant[V any](raw, typeName string, variantMap map[string]V, zero V) (V, error) {
	cleaned := strings.TrimSpace(raw)
	if len(cleaned) == 0 || strings.EqualFold(cleaned, "null") {
		return zero, nil
	}

	if len(variantMap) > 0 {
		if v, ok := variantMap[strings.ToLower(cleaned)]; ok {
			return v, nil
		}
	}

	return zero, fmt.Errorf("unknown %s variant %q", typeName, raw)
}

func unmarshalIntegerString[V IntNumber](
	str string,
	target *V,
	typeName string,
	variantMap map[string]V,
	maxValid int,
	zero V,
) error {
	cleaned := strings.TrimSpace(str)
	if len(cleaned) == 0 || strings.EqualFold(cleaned, "null") {
		*target = zero

		return nil
	}

	if len(variantMap) > 0 {
		if v, ok := variantMap[strings.ToLower(cleaned)]; ok {
			*target = v

			return nil
		}
	}

	return parseAndAssignNumericString(cleaned, target, typeName, maxValid)
}

func parseAndAssignNumericString[V IntNumber](
	str string,
	target *V,
	typeName string,
	maxValid int,
) error {
	parsed, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		return fmt.Errorf("unknown %s variant %q", typeName, str)
	}

	if maxValid >= 0 {
		if int64(parsed) > int64(maxValid) {
			return FormatNumericRangeError(typeName, parsed, maxValid)
		}
	}

	*target = V(parsed)

	return nil
}

func unmarshalIntegerNumeric[V IntNumber](
	data []byte,
	target *V,
	typeName string,
	maxValid int,
) error {
	var num int64
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}

	if maxValid >= 0 {
		if isOutOfRange(num, maxValid) {
			return FormatNumericRangeError(typeName, num, maxValid)
		}
	}

	*target = V(num)

	return nil
}

func isOutOfRange(num int64, maxValid int) bool {
	if num < 0 {
		return true
	}

	return num > int64(maxValid)
}
