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
func (dr *TDbDefaultDriver) Concat(field string, v string, more ...string) string {
	s := fmt.Sprintf("CONCAT(%s,%s", field, v)
	if len(more) > 0 {
		for _, m := range more {
			s += "," + m
		}
	}
	s += ")"
	return s
}

func (dr *TDbDefaultDriver) FromBase64(field string) string {
	return fmt.Sprintf("from_base64(%s)", field)
}

func (dr *TDbDefaultDriver) ToBase64(field string) string {
	return fmt.Sprintf("to_base64(%s)", field)
}
