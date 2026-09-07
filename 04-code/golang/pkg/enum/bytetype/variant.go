package bytetype

import (
	"encoding/json"
	"strconv"
	"strings"

	"coding-guidelines/common/pkg/baseenumer"
)

type (
	Variant byte

	VariantPredicate func(v Variant) bool
)

const (
	Zero    Variant = 0
	Min     Variant = 0
	One     Variant = 1
	Two     Variant = 2
	Three   Variant = 3
	Max     Variant = 255
	Invalid Variant = Zero
	Unknown Variant = Zero
)

var (
	_ baseenumer.ByteEnumer   = Variant(0)
	_ baseenumer.NumberEnumer = Variant(0)
	_ json.Marshaler          = Variant(0)
	_ json.Unmarshaler        = (*Variant)(nil)
)

func (v Variant) Byte() byte {
	return byte(v)
}

func (v Variant) ValueByte() byte {
	return byte(v)
}

func (v Variant) Value() byte {
	return byte(v)
}

func (v Variant) Bytes() []byte {
	return []byte{byte(v)}
}

func (v Variant) Int() int {
	return int(v)
}

func (v Variant) ValueInt() int {
	return int(v)
}

func (v Variant) ValueInt8() int8 {
	return int8(v)
}

func (v Variant) ValueInt16() int16 {
	return int16(v)
}

func (v Variant) ValueInt32() int32 {
	return int32(v)
}

func (v Variant) ValueUInt16() uint16 {
	return uint16(v)
}

func (v Variant) Code() uint16 {
	return uint16(v)
}

func (v Variant) ToPtr() *Variant {
	return &v
}

func (v Variant) Name() string {
	switch v {
	case Zero:
		return "Zero"
	case One:
		return "One"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Max:
		return "Max"
	default:
		return baseenumer.FormatNameValue("Byte", byte(v))
	}
}

func (v Variant) Label() string {
	return v.Name()
}

func (v Variant) String() string {
	return v.Name()
}

func (v Variant) ValueString() string {
	return strconv.Itoa(int(v))
}

func (v Variant) StringValue() string {
	return strconv.Itoa(int(v))
}

func (v Variant) ToNumberString() string {
	return strconv.Itoa(int(v))
}

func (v Variant) NameValue() string {
	return baseenumer.FormatNameValue(v.Name(), byte(v))
}

func (v Variant) JsonString() string {
	return "\"" + v.Name() + "\""
}

func (v Variant) IsValid() bool {
	return v != Zero
}

func (v Variant) IsInvalid() bool {
	return v == Zero
}

func (v Variant) IsEnum() bool {
	return v.IsValid()
}

func (v Variant) IsZero() bool {
	return v == Zero
}

func (v Variant) IsMin() bool {
	return v == Min
}

func (v Variant) IsOne() bool {
	return v == One
}

func (v Variant) IsTwo() bool {
	return v == Two
}

func (v Variant) IsThree() bool {
	return v == Three
}

func (v Variant) IsMax() bool {
	return v == Max
}

func (v Variant) Is(n Variant) bool {
	return v == n
}

func (v Variant) IsEqual(n byte) bool {
	return byte(v) == n
}

func (v Variant) IsEqualInt(n int) bool {
	return int(v) == n
}

func (v Variant) IsGreater(n byte) bool {
	return byte(v) > n
}

func (v Variant) IsGreaterInt(n int) bool {
	return int(v) > n
}

func (v Variant) IsGreaterEqual(n byte) bool {
	return byte(v) >= n
}

func (v Variant) IsGreaterEqualInt(n int) bool {
	return int(v) >= n
}

func (v Variant) IsLess(n byte) bool {
	return byte(v) < n
}

func (v Variant) IsLessInt(n int) bool {
	return int(v) < n
}

func (v Variant) IsLessEqual(n byte) bool {
	return byte(v) <= n
}

func (v Variant) IsLessEqualInt(n int) bool {
	return int(v) <= n
}

func (v Variant) IsBetween(start, end byte) bool {
	return baseenumer.IsBetween(byte(v), start, end)
}

func (v Variant) IsBetweenInt(start, end int) bool {
	return baseenumer.IsBetween(int(v), start, end)
}

func (v Variant) IsValueEqual(val byte) bool {
	return byte(v) == val
}

func (v Variant) IsNameEqual(name string) bool {
	return strings.EqualFold(v.Name(), name)
}

func (v Variant) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if v.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (v Variant) Add(n byte) Variant {
	return Variant(byte(v) + n)
}

func (v Variant) Subtract(n byte) Variant {
	return Variant(byte(v) - n)
}

func (v Variant) HasIndexInStrings(sliceOfStrings ...string) (string, bool) {
	if len(sliceOfStrings) == 0 {
		return "", false
	}

	idx := int(v)
	if idx >= 0 && idx < len(sliceOfStrings) {
		return sliceOfStrings[idx], true
	}

	return "", false
}

func New(input byte) Variant {
	return Variant(input)
}

func GetSet(isCondition bool, trueValue, falseValue Variant) Variant {
	if isCondition {
		return trueValue
	}

	return falseValue
}

func GetSetVariant(isCondition bool, trueValue, falseValue byte) Variant {
	if isCondition {
		return Variant(trueValue)
	}

	return Variant(falseValue)
}

func String(rawBytes []byte) string {
	if len(rawBytes) == 0 {
		return ""
	}

	return string(rawBytes)
}

func (v Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Name())
}

func (v *Variant) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 || trimmed == "null" {
		*v = Zero

		return nil
	}

	return v.unmarshalData(data)
}

func (v *Variant) unmarshalData(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		return v.unmarshalString(str)
	}

	var raw byte
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*v = Variant(raw)

	return nil
}

func (v *Variant) unmarshalString(str string) error {
	res := Parse(str)
	if res.IsSuccess() {
		*v = res.Data()

		return nil
	}

	return res.Fault()
}
