package t

import (
	"github.com/go-xe2/x/type/xconv"
)

func Map(value interface{}, tags ...string) map[string]interface{} {
	return xconv.Map(value, tags...)
}

func MapDeep(value interface{}, tags ...string) map[string]interface{} {
	return xconv.MapDeep(value, tags...)
}

func Structs(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return xconv.Structs(params, pointer, mapping...)
}

func StructsDeep(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return xconv.StructsDeep(params, pointer, mapping...)
}
