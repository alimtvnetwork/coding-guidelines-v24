package appfault

import (
	"encoding/json"

	"coding-guidelines/common/pkg/errtype"
)

// SerializeToJson serializes any value to JSON bytes returning *AppError on failure.
func SerializeToJson(val any) ([]byte, *AppError) {
	data, err := json.Marshal(val)
	if err != nil {
		return nil, Wrap(errtype.Execution, err, "failed to marshal value to json")
	}

	return data, nil
}

// SerializeToJsonString serializes any value to a JSON string returning *AppError on failure.
func SerializeToJsonString(val any) (string, *AppError) {
	data, appErr := SerializeToJson(val)
	if appErr != nil {
		return "", appErr
	}

	return string(data), nil
}

// DeserializeFromJson parses JSON bytes into target value returning *AppError on failure.
func DeserializeFromJson[T any](data []byte) (T, *AppError) {
	var target T
	if err := json.Unmarshal(data, &target); err != nil {
		return target, Wrap(errtype.Validation, err, "failed to unmarshal json data")
	}

	return target, nil
}

// DeserializeFromJsonString parses JSON string into target value returning *AppError on failure.
func DeserializeFromJsonString[T any](str string) (T, *AppError) {
	return DeserializeFromJson[T]([]byte(str))
}

// SerializeToJSON is an alias for SerializeToJson.
func SerializeToJSON(val any) ([]byte, *AppError) {
	return SerializeToJson(val)
}

// SerializeToJSONString is an alias for SerializeToJsonString.
func SerializeToJSONString(val any) (string, *AppError) {
	return SerializeToJsonString(val)
}

// DeserializeFromJSON is an alias for DeserializeFromJson.
func DeserializeFromJSON[T any](data []byte) (T, *AppError) {
	return DeserializeFromJson[T](data)
}

// DeserializeFromJSONString is an alias for DeserializeFromJsonString.
func DeserializeFromJSONString[T any](str string) (T, *AppError) {
	return DeserializeFromJsonString[T](str)
}
