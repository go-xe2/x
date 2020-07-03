package xjson

import (
	"fmt"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"reflect"
	"strings"
	"time"
)

var jsonFormatters = xsafeMap.NewAnyAnyMap()

func RegJsonFormatter(vt reflect.Type, formatter JsonDataFormatter) {
	jsonFormatters.Set(vt, formatter)
}

func JsonFormat(val interface{}) interface{} {
	vt := reflect.TypeOf(val)
	if jsonFormatters.Contains(vt) {
		f := jsonFormatters.Get(vt).(JsonDataFormatter)
		return f.Format(val)
	}
	return baseTypeFormat(val)
}

// 由key生成符合json键名格式的字符串
func MakeJsonKey(key string) string {
	return addQuotes(key)
}

func addQuotes(val interface{}) string {
	ret := t.String(val)
	ret = strings.Replace(ret, JsonQuotes(), "\\"+mJSON_VALUE_QUOTES, -1)
	return fmt.Sprintf("%s%s%s", mJSON_VALUE_QUOTES, ret, mJSON_VALUE_QUOTES)
}

// 格式化基础数据类型
func baseTypeFormat(val interface{}) interface{} {
	if val == nil {
		return ""
	}
	var s = ""
	switch v := val.(type) {
	case string:
		if xstring.IsNumeric(v) {
			s = v
		} else {
			s = addQuotes(v)
		}
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
					tmp += addQuotes(item)
				}
				s = "[" + tmp + "]"
			} else {
				s = "[" + strings.Join(items, ",") + "]"
			}
		} else {
			if vt, ok := v.(time.Time); ok {
				s = t.String(JsonFormat(vt))
			} else if vt, ok := v.(*time.Time); ok {
				s = t.String(JsonFormat(vt))
			} else {
				s = t.String(v)
				if !xstring.IsNumeric(s) {
					s = addQuotes(s)
				}
			}
		}
	}
	return s
}
