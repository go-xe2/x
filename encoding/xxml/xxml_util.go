package xxml

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"reflect"
	"time"
)

const xmlQuoteLeft = "<"
const xmlQuoteRight = ">"
const xmlQuoteClose = "/"

func filterXmlStr(val string) string {
	return xstring.ReplaceByMap(val, map[string]string{
		"<":  "&lt;",
		">":  "&gt;",
		"'":  "&apos;",
		"\"": "&quot;",
		"&":  "&amp;",
	})
}

func MakeOpenKey(key string) string {
	return fmt.Sprintf("%s%s%s", xmlQuoteLeft, filterXmlStr(key), xmlQuoteRight)
}

func MakeCloseKey(key string) string {
	return fmt.Sprintf("%s%s%s%s", xmlQuoteLeft, xmlQuoteClose, filterXmlStr(key), xmlQuoteRight)
}

// 生成xml节点
func MakeNode(name string, val interface{}, indent int, props ...map[string]interface{}) XmlStr {
	key := filterXmlStr(name)
	var properties = ""
	if len(props) > 0 {
		for _, mp := range props {
			for k, v := range mp {
				if properties != "" {
					properties += " "
				}
				properties += fmt.Sprintf("%s = \"%s\"", k, XmlFormat(v))
			}
		}
	}
	if properties != "" {
		properties = " " + properties
	}
	szIndent := xstring.MakeStrAndFill(indent, "\t", "")
	if val == nil || val == "" {
		return XmlStr(fmt.Sprintf("%s%s%s%s %s%s", szIndent, xmlQuoteLeft, key, properties, xmlQuoteClose, xmlQuoteRight))
	}
	return XmlStr(fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", szIndent, xmlQuoteLeft, key, xmlQuoteRight, properties, XmlFormat(val), xmlQuoteLeft, xmlQuoteClose, key, xmlQuoteRight))
}

// 格式化基础数据类型
func XmlFormat(val interface{}) interface{} {
	if val == nil {
		return ""
	}
	var s = ""
	switch v := val.(type) {
	case string:
		s = filterXmlStr(v)
		break
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		s = t.String(v)
		break
	default:
		vt := reflect.TypeOf(v)
		if vt.Kind() == reflect.Slice {
			items := t.SliceStr(v)
			tmp := ""
			for i, item := range items {
				if tmp != "" {
					tmp += ","
				}
				tmp += fmt.Sprintf("%s%d%s%s%s%s%d%s", xmlQuoteLeft, i, xmlQuoteRight, filterXmlStr(item), xmlQuoteLeft, xmlQuoteClose, i, xmlQuoteRight)
			}
			s = fmt.Sprintf("%slist%s\n%s\n%s%slist%s", xmlQuoteLeft, xmlQuoteRight, tmp, xmlQuoteLeft, xmlQuoteClose, xmlQuoteRight)
		} else {
			if vt, ok := val.(time.Time); ok {
				s = t.XTime(vt).String()
			} else if vt, ok := val.(*time.Time); ok {
				s = t.XTime(vt).String()
			} else {
				s = t.String(v)
			}
		}
	}
	return s
}
