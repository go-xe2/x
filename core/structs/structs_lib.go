package structs

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/type/xstring"
	"reflect"
)

func lookValue(v interface{}) (reflect.Value, error) {
	value, ok := v.(reflect.Value)
	if !ok {
		value = reflect.ValueOf(v)
	}
	vt := value.Type()
	if vt.Kind() == reflect.Ptr {
		if value.IsNil() {
			return value, errors.New("不能传入空指针")
		}
	}
	for vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return value, errors.New("不是struct类型")
	}
	return value, nil
}

func getStructFields(v reflect.Value, tagName string, owner ...*TStruct) ([]IField, error) {
	val := v
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	result := make([]IField, 0)
	dataType := val.Type()
	var ownerObj *TStruct
	if len(owner) > 0 {
		ownerObj = owner[0]
	}
	count := dataType.NumField()
	for i := 0; i < count; i++ {
		field := dataType.Field(i)
		fieldName := field.Name
		fieldTag := string(field.Tag)
		fieldType := field.Type.Kind()
		if xstring.IsFirstLetterLower(fieldName) {
			// 不导出私有字段
			continue
		}
		if tagName != "" {
			// 忽略字段
			if tag := field.Tag.Get(tagName); tag == "-" {
				continue
			}
		}
		fdVal := v.Field(i)
		var fd IField
		var err error
		if fieldType == reflect.Ptr {
			originType := field.Type.Elem()
			if fd, err = newFieldOwner(ownerObj, fieldName, fieldTag, field.Type, fdVal, field.Index, field.Offset, originType); err != nil {
				return nil, err
			}
		} else if fieldType == reflect.Map {
			mapValueType := field.Type.Elem()
			mapKeyType := field.Type.Key()
			if fd, err = newFieldOwner(ownerObj, fieldName, fieldTag, field.Type, fdVal, field.Index, field.Offset, mapKeyType, mapValueType); err != nil {
				return nil, err
			}
		} else if fieldType == reflect.Slice || fieldType == reflect.Array {
			itemType := field.Type.Elem()
			if fd, err = newFieldOwner(ownerObj, fieldName, fieldTag, field.Type, fdVal, field.Index, field.Offset, itemType); err != nil {
				return nil, err
			}
		} else if fieldType == reflect.Struct {
			structType := field.Type
			if fd, err = newFieldOwner(ownerObj, fieldName, fieldTag, field.Type, fdVal, field.Index, field.Offset, structType); err != nil {
				return nil, err
			}

		} else {
			fd, err = newFieldOwner(ownerObj, fieldName, fieldTag, field.Type, fdVal, field.Index, field.Offset)
			if err != nil {
				return nil, err
			}
		}
		result = append(result, fd)
	}
	return result, nil
}

func getStructMethods(instance interface{}) (result []*MethodInfo, err error) {
	vType := reflect.TypeOf(instance)
	vValue := reflect.ValueOf(instance)
	if vType.Kind() == reflect.Ptr {
		if vValue.IsNil() {
			return nil, errors.New("传入的指针为空")
		}
		for vType.Kind() == reflect.Ptr {
			vType = vType.Elem()
		}
		if vType.Kind() != reflect.Struct {
			return nil, errors.New("传入的不是struct或struct指针")
		}
	}
	methodCount := vValue.NumMethod()
	fmt.Println("method count:", methodCount)

	result = make([]*MethodInfo, 0)
	for i := 0; i < methodCount; i++ {
		method := vValue.Type().Method(i)
		methodInfo := NewMethodInfo(instance, vValue, method.Name, method.Index)
		numIn := method.Func.Type().NumIn()
		numOut := method.Func.Type().NumOut()
		inParamKinds := NewMethodParams(methodInfo, numIn)
		outParamKinds := NewMethodParams(methodInfo, numOut)
		for j := 0; j < numIn; j++ {
			it := method.Func.Type().In(j)
			inParamKinds.Set(j, it)
		}
		methodInfo.SetInParams(inParamKinds)
		for k := 0; k < numOut; k++ {
			it := method.Func.Type().Out(k)
			outParamKinds.Set(k, it)
		}
		methodInfo.SetOutParams(outParamKinds)
		result = append(result, methodInfo)
	}
	return
}
