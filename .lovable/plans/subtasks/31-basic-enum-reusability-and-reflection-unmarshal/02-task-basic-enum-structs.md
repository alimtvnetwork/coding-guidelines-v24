# Subtask 31.2: BasicIntegerEnum & BasicStringEnum Structs

## Context
Implement universal generic `BasicIntegerEnum[V IntNumber]` and `BasicStringEnum[V ~string]` in `04-code/golang/pkg/baseenumer/basic_enum.go` providing complete enum operations with minimal parameters.

## Target Files
1. `04-code/golang/pkg/baseenumer/basic_enum.go` [NEW]:
   - `type BasicIntegerEnum[V IntNumber] struct`
   - `func NewBasicInteger[V IntNumber](labels []string, zero V) *BasicIntegerEnum[V]`
   - `func (b *BasicIntegerEnum[V]) UnmarshalJSON(data []byte, target *V) error` (2 parameters!)
   - `func (b *BasicIntegerEnum[V]) All() []V`
   - `func (b *BasicIntegerEnum[V]) Values() []string`
   - `func (b *BasicIntegerEnum[V]) Parse(s string) result.Wrap[V]`
   - `type BasicStringEnum[V ~string] struct`
   - `func NewBasicString[V ~string](labels []string, zero V) *BasicStringEnum[V]`
   - `func (b *BasicStringEnum[V]) UnmarshalJSON(data []byte, target *V) error` (2 parameters!)
   - `func (b *BasicStringEnum[V]) All() []V`
   - `func (b *BasicStringEnum[V]) Values() []string`
   - `func (b *BasicStringEnum[V]) Parse(s string) result.Wrap[V]`
2. `04-code/golang/pkg/baseenumer/basic_enum_test.go` [NEW]:
   - Unit tests covering `UnmarshalJSON`, `All`, `Values`, and `Parse` for both integer and string generic enums.

## Verification
- `go test -v -C 04-code/golang -count=1 ./pkg/baseenumer/...`
