package t

import (
	"github.com/go-xe2/x/type/xconv"
)

// SliceInt is alias of Ints.
func SliceInt(i interface{}) []int {
	return Ints(i)
}

// SliceStr is alias of Strings.
func SliceStr(i interface{}) []string {
	return Strings(i)
}

// SliceAny is alias of Interfaces.
func SliceAny(i interface{}) []interface{} {
	return Interfaces(i)
}

// SliceFloat is alias of Floats.
func SliceFloat(i interface{}) []float64 {
	return Floats(i)
}

// SliceMap is alias of Maps.
func SliceMap(i interface{}) []map[string]interface{} {
	return Maps(i)
}

// SliceMapDeep is alias of MapsDeep.
func SliceMapDeep(i interface{}) []map[string]interface{} {
	return MapsDeep(i)
}

// SliceStruct is alias of Structs.
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return Structs(params, pointer, mapping...)
}

// SliceStructDeep is alias of StructsDeep.
func SliceStructDeep(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return StructsDeep(params, pointer, mapping...)
}

// Ints converts <i> to []int.
func Ints(i interface{}) []int {
	return xconv.Ints(i)
}

// Strings converts <i> to []string.
func Strings(i interface{}) []string {
	return xconv.Strings(i)
}

// Strings converts <i> to []float64.
func Floats(i interface{}) []float64 {
	return xconv.Floats(i)
}

// Interfaces converts <i> to []interface{}.
func Interfaces(i interface{}) []interface{} {
	return xconv.Interfaces(i)
}

// Maps converts <i> to []map[string]interface{}.
func Maps(value interface{}, tags ...string) []map[string]interface{} {
	return xconv.Maps(value, tags...)
}

// MapsDeep converts <i> to []map[string]interface{} recursively.
func MapsDeep(value interface{}, tags ...string) []map[string]interface{} {
	return xconv.MapsDeep(value, tags...)
}
