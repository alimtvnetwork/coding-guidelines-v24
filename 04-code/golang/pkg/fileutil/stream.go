package fileutil

import (
	"encoding/json"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/enum/filepermtype"
	"coding-guidelines/common/pkg/enum/openfiletype"
	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

func isDelimArrayStart(t json.Token) bool {
	delim, ok := t.(json.Delim)
	if !ok {
		return false
	}

	return delim == '['
}

func verifyArrayStart(decoder *json.Decoder, path string) *appfault.AppError {
	t, err := decoder.Token()
	if err != nil {
		return appfault.WrapFile(errtype.Serialization, err, path, "failed to read JSON array start")
	}

	if !isDelimArrayStart(t) {
		return appfault.NewFile(errtype.Serialization, path, "StreamJSON requires root element to be a JSON array")
	}

	return nil
}

func decodeJsonItems[T any](decoder *json.Decoder, path string, handler func(T) *appfault.AppError) *appfault.AppError {
	for decoder.More() {
		var item T
		if err := decoder.Decode(&item); err != nil {
			return appfault.WrapFile(errtype.Serialization, err, path, "failed to decode array element")
		}

		if err := handler(item); err != nil {
			return err
		}
	}

	if _, err := decoder.Token(); err != nil {
		return appfault.WrapFile(errtype.Serialization, err, path, "failed to read JSON array end")
	}

	return nil
}

// StreamJson sequentially decodes a massive JSON array from a file, passing each element to the handler.
// This prevents excessive RAM usage when dealing with huge datasets.
func StreamJson[T any](path string, handler func(T) *appfault.AppError) BoolResult {
	fRes := OpenFile(path, openfiletype.ReadOnly, filepermtype.Standard)
	if fRes.HasError() {
		return result.WrapFailure[bool](fRes.Fault())
	}

	defer fRes.Data().Close()

	decoder := json.NewDecoder(fRes.Data())
	if err := verifyArrayStart(decoder, path); err != nil {
		return result.WrapFailure[bool](err)
	}

	if err := decodeJsonItems(decoder, path, handler); err != nil {
		return result.WrapFailure[bool](err)
	}

	return BoolSuccess(true)
}

// StreamJSON is an alias for StreamJson.
func StreamJSON[T any](path string, handler func(T) *appfault.AppError) BoolResult {
	return StreamJson[T](path, handler)
}
