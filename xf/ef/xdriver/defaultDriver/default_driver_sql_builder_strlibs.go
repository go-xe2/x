package defaultDriver

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

var _ DbDriverSqlBuilderStrLibs = &TDbDefaultDriver{}

func (dr *TDbDefaultDriver) Substring(field string, from int, l int) string {
	return fmt.Sprintf("Substring(%s, %d, %d)", field, from, l)
}

// 字符串连接函数
func (dr *TDbDefaultDriver) Concat(field string, v string) string {
	return fmt.Sprintf("CONCAT(%s,%s)", field, v)
}
