package defaultDriver

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xdbUtil"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	"reflect"
	"strings"
)

var _ DbDriverSqlBuilder = &TDbDefaultDriver{}

func (dr *TDbDefaultDriver) Driver() DbDriver {
	return dr
}

func (dr *TDbDefaultDriver) QuotesName(name string) string {
	return fmt.Sprintf("`%s`", name)
}

func (dr *TDbDefaultDriver) PlaceHolder(varName string) string {
	return "?"
}

func (dr *TDbDefaultDriver) MakeRealValue(val interface{}) string {
	if val == nil {
		return ""
	}
	var s = ""
	switch v := val.(type) {
	case string:
		s = xdbUtil.AddQuotes(v, "'")
		break
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		s = t.String(v)
		break
	default:
		vt := reflect.TypeOf(v)
		if vt.Kind() == reflect.Slice {
			items := t.SliceStr(v)
			if vt.Elem().Kind() == reflect.String {
				tmp := ""
				for _, item := range items {
					if tmp != "" {
						tmp += ","
					}
					tmp += xdbUtil.AddQuotes(item, "'")
				}
				s = "(" + tmp + ")"
			} else {
				s = "(" + strings.Join(items, ",") + ")"
			}
		} else {
			s = xdbUtil.AddQuotes(t.String(v), "'")
		}
	}
	return s
}

func (dr *TDbDefaultDriver) MakeDataDefine(ddType DbDataType, size int, decimal ...int) string {
	if fn, ok := DbDataTypeDefineFunMaps[ddType]; ok {
		return fn(size, decimal...)
	} else {
		return ""
	}
}

func (dr *TDbDefaultDriver) MakeQueryParams(vars []SqlParam) (result []interface{}) {
	for _, item := range vars {
		result = append(result, item.Val())
	}
	return
}

func (dr *TDbDefaultDriver) BuildQuery(table, fields, joins, where, group, having, order string, rows, offset int) string {
	result := "SELECT "
	if fields == "" {
		result += " * "
	} else {
		result += fields
	}
	result += " FROM " + table
	if joins != "" {
		result += " " + joins
	}
	if where != "" {
		result += " WHERE " + where
	}
	if group != "" {
		result += " GROUP BY " + group
	}
	if having != "" {
		result += " HAVING " + having
	}
	if order != "" {
		result += " ORDER BY " + order
	}
	if offset <= 0 && rows > 0 {
		result += " LIMIT " + t.String(rows)
	} else if offset > 0 && rows > 0 {
		result += " LIMIT " + t.String(offset) + "," + t.String(rows)
	}
	return result
}

func (dr *TDbDefaultDriver) BuildUpdate(table, fields, joins, where string) string {
	sql := "UPDATE " + table
	if joins != "" {
		sql += " " + joins
	}
	sql += " SET " + fields
	if where != "" {
		sql += " WHERE " + where
	}
	return sql
}

// 生成插入脚本
// Deprecated: 不再使用
func (dr *TDbDefaultDriver) BuildInsert(table, fields, values, fromTable string) string {
	sql := fmt.Sprintf("INSERT INTO %s(%s)", table, fields)
	if fromTable == "" {
		sql += " VALUES (" + values + ")"
	} else {
		sql += " SELECT " + values + " FROM " + fromTable
	}
	return sql
}
