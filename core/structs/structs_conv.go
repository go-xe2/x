package structs

import (
	"reflect"
)

func (s *TStruct) AppendResult(arg map[string]interface{}) {
	s.result = append(s.result, arg)
}

func (s *TStruct) SetResult(arg []map[string]interface{}) {
	s.result = arg
}

func (s *TStruct) GetResult() []map[string]interface{} {
	return s.result
}

func (s *TStruct) SetExtraCols(args []string) *TStruct {
	s.ExtraCols = args
	return s
}

func (s *TStruct) StructContent2Map(data interface{}) []map[string]interface{} {
	s.SetResult(make([]map[string]interface{}, 0))
	val := reflect.Indirect(reflect.ValueOf(data))
	switch val.Kind() {
	case reflect.Struct: // struct
		s.getStructContent(val)
	case reflect.Slice: // []struct
		for i := 0; i < val.Len(); i++ {
			s.getStructContent(reflect.Indirect(val.Index(i)))
		}
	}
	return s.GetResult()
}

func inArray(needle string, arr []string) bool {
	for _, item := range arr {
		if needle == item {
			return true
		}
	}
	return false
}

func (s *TStruct) getStructContent(val reflect.Value) {
	valType := val.Type()
	var mapTmp = make(map[string]interface{}, 0)
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := valType.Field(i)
		switch valueField.Kind() {
		case reflect.Struct:
			continue
		default:
			var fieldName = typeField.Tag.Get(s.TagName())
			// 如果该字段没有被忽略, 则获取值
			if fieldName != "-" {
				// 如果tag为空, 则获取字段名字
				if fieldName == "" {
					fieldName = typeField.Name
				}
				// 如果是struct字段类型的默认值, 则不获取
				v := valueField.Interface()
				if b, ok := v.(bool); b && ok {
					mapTmp[fieldName] = valueField.Interface()
				} else {
					// 如果指定了强制获取, 则也获取
					if inArray(fieldName, s.ExtraCols) {
						mapTmp[fieldName] = valueField.Interface()
					}
				}
			}
		}
	}
	s.AppendResult(mapTmp)
}
