package baseenumer

import (
	"reflect"
	"strings"
)

// ResolveTypeName derives a clean, human-readable type name from a pointer using reflection.
func ResolveTypeName[V any](target *V) string {
	if target == nil {
		return "Variant"
	}

	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	pkg := t.PkgPath()
	if idx := strings.LastIndex(pkg, "/"); idx >= 0 {
		pkg = pkg[idx+1:]
	}

	if pkg != "" {
		return pkg
	}

	return t.Name()
}
