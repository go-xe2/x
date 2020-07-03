package anno

import (
	"bytes"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/xstring"
	"regexp"
)

func CreateAnn(name string, params map[string]interface{}) AnnotationContainer {
	annotationEntry.Init()
	if ann := annotationEntry.GetAnnotation(name); ann != nil {
		return newAnnotationContainer(ann, params)
	}
	panic(exception.Newf("未注册元注解类型%s", name))
}

var annoRegex = regexp.MustCompile(`^@(.+)\((.*)\)$`)

func IsAnnotationString(str string) bool {
	return annoRegex.MatchString(str)
}

func parseAnnParams(s string) map[string]interface{} {
	result := make(map[string]interface{})
	chars := []rune(s)
	nLen := len(chars)
	buf := bytes.NewBufferString("")
	szKey := ""
	szVal := ""
	layMode := 0 // 0:查找键,1:查找值
	vt := 0      // 0:普通值，1：单引号内字符串
	for i := 0; i < nLen; i++ {
		c := chars[i]
		if c == '=' && vt == 0 && layMode == 0 {
			szKey = buf.String()
			buf.Reset()
			layMode = 1
		} else if c == '\'' {
			// 字符串内容
			if vt == 0 {
				vt = 1
			} else if vt == 1 {
				vt = 0
			}
		} else if c == ',' && vt == 0 && layMode == 1 {
			// 参数分割
			szVal = buf.String()
			buf.Reset()
			layMode = 0
			if szKey != "" {
				result[xstring.Trim(szKey)] = xstring.Trim(szVal)
				szKey = ""
			}
		} else {
			buf.WriteRune(c)
		}
	}
	if layMode == 1 && szKey != "" {
		result[xstring.Trim(szKey)] = xstring.Trim(buf.String())
	}
	return result
}

// 创建元注解实例
// 参数:
// @param annStr 元注解字符串，格式为:@name(paramName = paramValue)
func Create(annStr string) AnnotationContainer {
	if !IsAnnotationString(annStr) {
		return nil
	}
	items := annoRegex.FindStringSubmatch(annStr)
	annName := items[1]
	szParams := items[2]
	params := parseAnnParams(szParams)
	return CreateAnn(annName, params)
}
