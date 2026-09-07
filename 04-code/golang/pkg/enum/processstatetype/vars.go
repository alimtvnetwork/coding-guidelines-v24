package processstatetype

import (
	"fmt"
	"strconv"
	"strings"

	"coding-guidelines/common/pkg/errtype"
	"coding-guidelines/common/pkg/result"
)

var (
	variantLabels = [...]string{
		Invalid:   "Unknown",
		Pending:   "Pending",
		Running:   "Running",
		Completed: "Completed",
		Failed:    "Failed",
		Cancelled: "Cancelled",
	}

	variantMap = compileVariantMap()
)

func compileVariantMap() map[string]Variant {
	m := make(map[string]Variant, (len(variantLabels)*4)+4)
	for i, label := range variantLabels {
		v := Variant(i)
		m[label] = v
		m[strings.ToLower(label)] = v
		m[strings.ToUpper(label)] = v
		m[strconv.Itoa(i)] = v
	}

	m["unknown"] = Invalid
	m["invalid"] = Invalid
	m["UNKNOWN"] = Invalid
	m["INVALID"] = Invalid

	return m
}

func All() []Variant {
	items := make([]Variant, 0, len(variantLabels)-1)
	for i := 1; i < len(variantLabels); i++ {
		items = append(items, Variant(i))
	}

	return items
}

func Values() []string {
	names := make([]string, 0, len(variantLabels)-1)

	return append(names, variantLabels[1:]...)
}

func Parse(s string) result.Wrap[Variant] {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return result.WrapFailureWithId[Variant](errtype.Validation, "cannot parse empty string as processstatetype")
	}

	if v, ok := variantMap[strings.ToLower(trimmed)]; ok {
		return result.WrapSuccess(v)
	}

	return result.WrapFailureWithId[Variant](
		errtype.NotFound,
		fmt.Sprintf("unknown processstatetype variant %q, supported variants: [%s]", s, strings.Join(Values(), ", ")),
	)
}
