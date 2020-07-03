package xproto

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/xtime"
	"reflect"
	"time"
)

type ProtoDataType uint8

var (
	ProtoDatasetType   = reflect.TypeOf([]ProtoDataset{}).Elem()
	ProtoEnumType      = reflect.TypeOf([]ProtoEnum{}).Elem()
	ProtoStructType    = reflect.TypeOf([]ProtoStruct{}).Elem()
	ProtoUnionType     = reflect.TypeOf([]ProtoUnion{}).Elem()
	ProtoClassType     = reflect.TypeOf([]ProtoClass{}).Elem()
	ProtoExceptionType = reflect.TypeOf([]ProtoException{}).Elem()
	ProtoDynamicType   = reflect.TypeOf([]interface{}{}).Elem()
)

const (
	PDTUnknown ProtoDataType = iota
	// 空类型
	PDTNull
	// 可变类型
	PDTDynamic
	// string
	PDTString
	// int8
	PDTInt8
	// int16
	PDTInt16
	// int32
	PDTInt32
	// int64
	PDTInt64
	// bool类型
	PDTBool
	// 浮点类型
	PDTDouble
	// 日期类型
	PDTDatetime

	// 字典类型
	PDTMap
	// 列表类型
	PDTList
	// 表记录类型
	PDTDataset

	// 枚举类型
	PDTEnum
	// struct数据结构体
	PDTStruct
	// union结构体
	PDTUnion
	// 其他自定义实现BinSerializer接口的类型
	PDTClass
	// 异常类结构体
	PDTException
)

func (pdt ProtoDataType) String() string {
	switch pdt {
	case PDTUnknown:
		return "unknownType"
	case PDTNull:
		return "nullType"
	case PDTDynamic:
		return "dynamicType"
	case PDTString:
		return "stringType"
	case PDTInt8:
		return "int8Type"
	case PDTInt16:
		return "int16Type"
	case PDTInt32:
		return "int32Type"
	case PDTInt64:
		return "int64Type"
	case PDTDouble:
		return "doubleType"
	case PDTBool:
		return "boolType"
	case PDTDatetime:
		return "datetimeType"
	case PDTMap:
		return "mapType"
	case PDTList:
		return "listType"
	case PDTDataset:
		return "datasetType"
	case PDTEnum:
		return "enumType"
	case PDTStruct:
		return "structType"
	case PDTUnion:
		return "unionType"
	case PDTClass:
		return "classType"
	case PDTException:
		return "exceptionType"
	}
	return "unknownType"
}

func (pdt ProtoDataType) IsBasicType() bool {
	return pdt >= PDTString && pdt <= PDTDataset
}

func (pdt ProtoDataType) IsUserType() bool {
	return pdt >= PDTEnum && pdt <= PDTException
}

var kind2ProtoDataTypes = map[reflect.Kind]ProtoDataType{
	reflect.String:  PDTString,
	reflect.Int:     PDTInt64,
	reflect.Int8:    PDTInt8,
	reflect.Int16:   PDTInt16,
	reflect.Int32:   PDTInt32,
	reflect.Int64:   PDTInt64,
	reflect.Uint:    PDTInt64,
	reflect.Uint8:   PDTInt8,
	reflect.Uint16:  PDTInt16,
	reflect.Uint32:  PDTInt32,
	reflect.Uint64:  PDTInt64,
	reflect.Bool:    PDTBool,
	reflect.Float32: PDTDouble,
	reflect.Float64: PDTDouble,
	reflect.Slice:   PDTList,
	reflect.Array:   PDTList,
	reflect.Map:     PDTMap,
}

func GetProtoDataType(rv reflect.Type) ProtoDataType {
	if t, ok := kind2ProtoDataTypes[rv.Kind()]; ok {
		return t
	} else if rv.Kind() == reflect.Interface && rv.Name() == "" {
		return PDTDynamic
	}
	t1 := time.Time{}
	t2 := xtime.Time{}
	if rv == reflect.TypeOf(t1) ||
		rv == reflect.TypeOf(&t1) ||
		rv == reflect.TypeOf(t2) ||
		rv == reflect.TypeOf(&t2) {
		return PDTDatetime
	} else if rv == ProtoDatasetType {
		return PDTDataset
	} else if rv == ProtoEnumType {
		return PDTEnum
	} else if rv == ProtoStructType {
		return PDTStruct
	} else if rv == ProtoUnionType {
		return PDTUnion
	} else if rv == ProtoClassType {
		return PDTClass
	} else if rv == ProtoExceptionType {
		return PDTException
	} else {
		return PDTUnknown
	}
}

// 获取定义数据类型
func GetProtoDataTypes(typ reflect.Type) []ProtoDataType {
	result := make([]ProtoDataType, 0)
	nElem := GetProtoDataType(typ)
	result = append(result, nElem)
	if nElem == PDTList {
		nextTypes := GetProtoDataTypes(typ.Elem())
		result = append(result, nextTypes...)
	} else if nElem == PDTMap {
		// keyType, valueType
		nextTypes := GetProtoDataTypes(typ.Key())
		result = append(result, nextTypes...)
		nextTypes = GetProtoDataTypes(typ.Elem())
		result = append(result, nextTypes...)
	}
	return result
}

// 由协议数据类型生成反射数据类型
func NewTypeByProtoDataTypes(types []ProtoDataType) (reflect.Type, error) {
	nLen := len(types)
	index := nLen - 1
	pop := func() ProtoDataType {
		if index >= 0 {
			i := index
			index--
			return types[i]
		} else {
			return PDTUnknown
		}
	}

	var priorTypes = make([]reflect.Type, 0)

	for t := pop(); t != PDTUnknown; t = pop() {

		if t == PDTList || t == PDTMap {
			if t == PDTList {
				if len(priorTypes) == 0 {
					return nil, exception.NewText("数据类型定义错误，缺少列表项类型定义")
				}
				tmpType := reflect.SliceOf(priorTypes[0])
				priorTypes = []reflect.Type{tmpType}
			} else {
				if len(priorTypes) != 2 {
					return nil, exception.NewText("数据类型定义错误,缺少map的key或value类型定义")
				}
				tmpType := reflect.MapOf(priorTypes[0], priorTypes[1])
				priorTypes = []reflect.Type{tmpType}
			}
		} else {
			switch t {
			case PDTDynamic:
				priorTypes = append([]reflect.Type{ProtoDynamicType}, priorTypes...)
			case PDTString:
				priorTypes = append([]reflect.Type{reflect.TypeOf(string(""))}, priorTypes...)
				break
			case PDTInt8:
				priorTypes = append([]reflect.Type{reflect.TypeOf(int8(0))}, priorTypes...)
				break
			case PDTInt16:
				priorTypes = append([]reflect.Type{reflect.TypeOf(int16(0))}, priorTypes...)
				break
			case PDTInt32:
				priorTypes = append([]reflect.Type{reflect.TypeOf(int32(0))}, priorTypes...)
				break
			case PDTInt64:
				priorTypes = append([]reflect.Type{reflect.TypeOf(int64(0))}, priorTypes...)
				break
			case PDTDouble:
				priorTypes = append([]reflect.Type{reflect.TypeOf(float64(0))}, priorTypes...)
				break
			case PDTBool:
				priorTypes = append([]reflect.Type{reflect.TypeOf(bool(false))}, priorTypes...)
				break
			case PDTDatetime:
				priorTypes = append([]reflect.Type{reflect.TypeOf(time.Time{})}, priorTypes...)
				break
			case PDTDataset:
				priorTypes = append([]reflect.Type{ProtoDatasetType}, priorTypes...)
				break
			case PDTEnum:
				priorTypes = append([]reflect.Type{ProtoEnumType}, priorTypes...)
				break
			case PDTStruct:
				priorTypes = append([]reflect.Type{ProtoStructType}, priorTypes...)
				break
			case PDTUnion:
				priorTypes = append([]reflect.Type{ProtoUnionType}, priorTypes...)
				break
			case PDTClass:
				priorTypes = append([]reflect.Type{ProtoClassType}, priorTypes...)
				break
			case PDTException:
				priorTypes = append([]reflect.Type{ProtoExceptionType}, priorTypes...)
				break
			}
		}
	}
	typLen := len(priorTypes)
	if typLen == 0 || typLen > 1 {
		return reflect.TypeOf(nil), exception.NewText("数据类型定义错误")
	} else {
		return priorTypes[0], nil
	}
}
