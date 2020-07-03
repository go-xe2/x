package xqi

import (
	"reflect"
)

type TQueryColDef struct {
	// 字段名称
	ColName string
	// 扫描数据类型
	ColType FieldDataType
}

// 查询字段信息
type TQueryColInfo struct {
	TQueryColDef
	// 字段序号
	ColIndex int
	// 数据库定义类型
	DbType string
}

type TQueryColValue struct {
	*TQueryColInfo
	// 字段值
	ColValue interface{}
}

type QueryColValue = *TQueryColValue

// 数据访问方法
type QueryBinderVisit = func(row int, values ...interface{}) (interface{}, bool)

type DbQueryBinder interface {
	SetOptions(options map[string]interface{}) DbQueryBinder

	// 使用定义参数创建新实例
	NewInstance(options ...map[string]interface{}) DbQueryBinder
	// 开始绑定，返回false时结束绑定
	StartBuild(colInfos ...*TQueryColInfo) bool
	// 开始创建行,返回false忽略该行
	// rowIndex: 行序号
	// colCount: 总列数
	StartBuildRow(rowIndex int, colCount int) bool
	// 创建行
	// row:当前行号
	// colInfos行中的列信息
	// 返回:result行数据结构，exit: 返回true结束绑定
	BuildRow(row int, colInfos *[]QueryColValue) (result interface{}, exit bool)
	// 行创建完成
	// rowData为BuildRow所创建的行数据
	EndBuildRow(rowData interface{})
	// 绑定完成,返回所有行数据
	EndBuild() interface{}
	//// 字段映射名
	// @param qryName 查询字段名
	// @return 返回输出需要的字段名
	FieldName(colIndex int, qryName string) string
	// 字段数据转换
	// @param qryName 字段名
	// @param val 原数据
	// @return 返回转换后的数据
	FieldConvert(colIndex int, qryName string, val interface{}) interface{}
}

func NewQueryColDefByType(name string, gType reflect.Type, dbType ...string) *TQueryColDef {
	s := ""
	if len(dbType) > 0 {
		s = dbType[0]
	}
	ft := QueryDataType2FieldType(gType, s)
	return NewQueryColDef(name, ft)
}

func NewQueryColDef(name string, fieldType FieldDataType) *TQueryColDef {
	return &TQueryColDef{
		ColName: name,
		ColType: fieldType,
	}
}

func QueryDataType2FieldType(gType reflect.Type, dbType string) FieldDataType {
	for gType.Kind() == reflect.Ptr {
		gType = gType.Elem()
	}
	var fieldType = FDTUnknown
	switch gType.Kind() {
	case reflect.String:
		fieldType = FDTString
		break
	case reflect.Int:
		fieldType = FDTInt
		break
	case reflect.Int8:
		fieldType = FDTInt8
		break
	case reflect.Int16:
		fieldType = FDTInt16
		break
	case reflect.Int32:
		fieldType = FDTInt32
		break
	case reflect.Int64:
		fieldType = FDTInt64
		break
	case reflect.Uint:
		fieldType = FDTUint
		break
	case reflect.Uint8:
		fieldType = FDTUint8
		break
	case reflect.Uint16:
		fieldType = FDTUint16
		break
	case reflect.Uint32:
		fieldType = FDTUint32
		break
	case reflect.Uint64:
		fieldType = FDTInt64
		break
	case reflect.Float32:
		fieldType = FDTFloat
		break
	case reflect.Float64:
		fieldType = FDTDouble
		break
	case reflect.Bool:
		fieldType = FDTBool
		break
	default:
		if dbType == "DATE" || dbType == "TIME" || dbType == "DATETIME" ||
			dbType == "TIMESTAMP" {
			fieldType = FDTDatetime
		} else if dbType == "BINARY" {
			fieldType = FDTBinary
		}
	}
	return fieldType
}

func NewQueryColInfo(index int, name string, gType reflect.Type, dbType string) *TQueryColInfo {
	fieldType := QueryDataType2FieldType(gType, dbType)
	return &TQueryColInfo{
		ColIndex: index,
		DbType:   dbType,
		TQueryColDef: TQueryColDef{
			ColName: name,
			ColType: fieldType,
		},
	}
}
