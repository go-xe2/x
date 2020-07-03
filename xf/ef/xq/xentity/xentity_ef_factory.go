package xentity

import (
	"github.com/go-xe2/x/xf/ef/xqi"
)

type DataFieldConstructor = func(entity xqi.Entity, fieldName string, alias string, options ...map[string]interface{}) interface{}

// 创建指定数据类型的的字段，如果指定类型的字段类型存在则返回该类型字段实例，否则返回nil
func CreateField(fdType xqi.FieldDataType, table xqi.Entity, defineName string, attrs []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{} {
	efTypeConstructorEntry.Init()
	if fn := efTypeConstructorEntry.TypeConstructor(fdType); fn != nil {
		v := fn(table, defineName, attrs, annotations, params...)
		return v
	}
	return nil
}
