package xentity

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type baseEntity struct {
	*xqcomm.TSqlTable
	// 字段列表
	fields    []xqi.SqlField
	fieldMaps map[string]xqi.SqlField
	// 实体关联的其他实体
	foreignTables xqi.SqlTables
	foreignFields []xqi.EFForeign
	attributes    map[string]xqi.XqAttribute
}

var _ xqi.BasicEntity = (*baseEntity)(nil)

func newBaseEntity(inherited interface{}, tableName string, alias ...string) *baseEntity {
	inst := &baseEntity{
		TSqlTable:     nil,
		fields:        make([]xqi.SqlField, 0),
		fieldMaps:     make(map[string]xqi.SqlField),
		foreignTables: xqcomm.NewSqlTables(),
		foreignFields: make([]xqi.EFForeign, 0),
		attributes:    make(map[string]xqi.XqAttribute),
	}
	var this interface{} = inst
	if inherited != nil {
		if _, ok := inherited.(xqi.SqlTable); ok {
			this = inherited
		}
	}
	base := xqcomm.NewSqlTable(this, tableName, alias...)
	inst.TSqlTable = base
	return inst
}

func (ent *baseEntity) AllField() []xqi.SqlField {
	return ent.fields
}

func (ent *baseEntity) FieldByName(fieldName string) xqi.SqlField {
	if f, ok := ent.fieldMaps[fieldName]; ok {
		return f
	}
	return nil
}

func (ent *baseEntity) FieldByNameE(fieldName string) xqi.EntField {
	if f, ok := ent.fieldMaps[fieldName].(xqi.EntField); ok {
		return f
	}
	return nil
}

func (ent *baseEntity) FieldCount() int {
	return len(ent.fields)
}

func (ent *baseEntity) Field(index int) xqi.SqlField {
	return ent.fields[index]
}

func (ent *baseEntity) AddField(field xqi.SqlField) xqi.SqlTable {
	if fd, ok := field.This().(xqi.SqlTableField); ok {
		if _, fdOk := ent.fieldMaps[fd.FieldName()]; !fdOk {
			ent.fieldMaps[fd.FieldName()] = field
			ent.fieldMaps[fd.AliasName()] = field
			ent.fields = append(ent.fields, field)
		}
	} else {
		if _, ok := ent.fieldMaps[field.AliasName()]; !ok {
			ent.fieldMaps[field.AliasName()] = field
			ent.fields = append(ent.fields, field)
		}
	}
	return ent
}

// 实体构造方法
func (ent *baseEntity) Constructor(attrs []xqi.XqAttribute, inherited ...interface{}) interface{} {
	ent.attributes = make(map[string]xqi.XqAttribute, 0)
	for _, attr := range attrs {
		if _, ok := ent.attributes[attr.AttrName()]; !ok {
			ent.attributes[attr.AttrName()] = attr
		}
	}
	return ent
}

// 初始化实体字段
func (ent *baseEntity) init() *baseEntity {
	return ent
}

func (ent *baseEntity) ForeignTables() xqi.SqlTables {
	return ent.foreignTables
}

func (ent *baseEntity) String() string {
	return fmt.Sprintf("baseEntity { sqlTable: %v }", ent.TSqlTable)
}
