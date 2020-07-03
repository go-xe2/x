package t

import (
	"github.com/go-xe2/x/type/xconv"
)

// 使用字典map设置struct字段值
// params 设置struct字段值的数据字典, 字段名不区分大小写
// mapping 字典键名key与struct字段名映射关系
func Struct(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return xconv.Struct(params, pointer, mapping...)
}


// 使用字典map设置struct字段值包括内嵌struct
// params 设置struct字段值的数据字典, 字段名不区分大小写
// mapping 字典键名key与struct字段名映射关系
func StructDeep(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return xconv.StructDeep(params, pointer, mapping...)
}
