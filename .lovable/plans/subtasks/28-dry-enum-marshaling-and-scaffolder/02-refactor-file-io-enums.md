# Subtask 28.2: Refactor File I/O Enums to DRY Marshaling

## Context
Refactor all file I/O enum packages to delegate JSON serialization to `baseenumer.MarshalJSON` and `baseenumer.UnmarshalIntegerJSON`, removing repetitive `unmarshalData` and `unmarshalString` helper functions.

## Target Files
1. `04-code/golang/pkg/enum/fileoptype/variant.go`
   - Use `baseenumer.MarshalJSON(v.Name())`
   - Use `baseenumer.UnmarshalIntegerJSON(data, v, "fileoptype", variantMap, len(variantLabels)-1, Invalid)`
   - Delete `unmarshalData` and `unmarshalString`
2. `04-code/golang/pkg/enum/filewritemodetype/variant.go`
   - Use `baseenumer.MarshalJSON(v.Name())`
   - Use `baseenumer.UnmarshalIntegerJSON(data, v, "filewritemodetype", variantMap, len(variantLabels)-1, Invalid)`
   - Delete `unmarshalData` and `unmarshalString`
3. `04-code/golang/pkg/enum/filepermtype/variant.go`
   - Use `baseenumer.MarshalJSON(p.OctalString())`
   - Use `baseenumer.UnmarshalIntegerJSON(data, p, "filepermtype", nil, 07777, Standard)`
   - Delete `unmarshalData` and `unmarshalString`
4. `04-code/golang/pkg/enum/openfiletype/variant.go`
   - Use `baseenumer.MarshalJSON(v.Name())`
   - Use `baseenumer.UnmarshalIntegerJSON(data, v, "openfiletype", variantMap, len(variantLabels)-1, Invalid)`
   - Replace inlined unmarshaling logic with one-line call

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/enum/fileoptype/... ./pkg/enum/filewritemodetype/... ./pkg/enum/filepermtype/... ./pkg/enum/openfiletype/...`
