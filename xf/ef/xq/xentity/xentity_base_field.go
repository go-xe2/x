package xentity

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type baseField struct {
	*xqcomm.TSqlTableField
	entity xqi.Entity
	// 是否是主键
	isPrimary bool
	// 查询规则
	rule       string
	defineName string
	// 字段元注解
	annotations map[string]interface{}
	attr        xqi.EntityFieldAttribute
	// 数据绑定时使用
	index     int
	formatter string
}

var _ xqi.EntField = (*baseField)(nil)
var _ xqi.DSField = (*baseField)(nil)
var _ xqi.SqlTableField = (*baseField)(nil)

func newBaseField(entity xqi.Entity, defineName string, attrs []xqi.XqAttribute, annotations map[string]interface{}, callparams []interface{}, inherited ...interface{}) *baseField {
	inst := &baseField{
		entity:      entity,
		defineName:  defineName,
		annotations: annotations}
	for _, attr := range attrs {
		if tmpAttr, ok := attr.(xqi.EntityFieldAttribute); ok {
			inst.attr = tmpAttr
			break
		}
	}
	fieldName := defineName
	fieldAlias := ""
	if inst.attr != nil {
		fieldName = inst.attr.FieldName()
		fieldAlias = inst.attr.FieldAlias()
		inst.isPrimary = inst.attr.IsPrimary()
		inst.rule = inst.attr.Rule()
		inst.formatter = inst.attr.Formatter()
	}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.SqlTableField); ok {
			this = inherited[0]
		}
	}
	base := xqcomm.NewSqlTableField(this, entity, fieldName, fieldAlias)
	inst.TSqlTableField = base
	return inst
}

func (ef *baseField) Constructor(instance interface{}, props ...interface{}) interface{} {
	return ef.TSqlTableField.Constructor(instance, props...)
}

func (ef *baseField) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &baseField{
		entity:      ef.entity,
		isPrimary:   ef.isPrimary,
		rule:        ef.rule,
		defineName:  ef.defineName,
		annotations: ef.annotations,
		attr:        ef.attr,
		index:       ef.index,
		formatter:   ef.formatter,
	}
	var this interface{} = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(xqi.SqlTableField); ok {
			this = inherited[0]
		}
	}

	base := xqcomm.NewSqlTableField(this, ef.entity, ef.FieldName(), alias)
	inst.TSqlTableField = base
	return inst
}

func (ef *baseField) Entity() xqi.Entity {
	return ef.entity
}

func (ef *baseField) DefineName() string {
	return ef.defineName
}

func (ef *baseField) Supper() xqi.EntField {
	return nil
}

func (ef *baseField) IsPrimary() bool {
	return ef.isPrimary
}

func (ef *baseField) IsForeign() bool {
	return false
}

func (ef *baseField) Rule() string {
	return ef.rule
}

func (ef *baseField) Alias(alias string) xqi.SqlField {
	return ef.This().(xqi.EntField).NewInstance(alias)
}

func (ef *baseField) GetAnnotation(annName string) interface{} {
	if ef.annotations == nil {
		return nil
	}
	if v, ok := ef.annotations[annName]; ok {
		return v
	}
	return nil
}

func (ef *baseField) Value() interface{} {
	if ef.index < 0 {
		panic(exception.NewText("打开的数据集中未包含该字段"))
	}
	if ds, ok := ef.entity.This().(xqi.Dataset); ok {
		v := ds.FieldValue(ef.index)
		if fn := ef.entity.FieldFormat(ef.DefineName()); fn != nil {
			v = fn(v)
		}
		return v
	}
	return nil
}

func (ef *baseField) entityFieldEditor(value xqi.FieldValue) xqi.FieldValue {
	if editor, ok := ef.entity.This().(xqi.EntityEditor); ok {
		return editor.FieldEdit(ef.This().(xqi.EntField), value)
	}
	return value
}

func (ef *baseField) Set(val interface{}) xqi.FieldValue {
	v := xqcomm.NewFieldValue(ef.This().(xqi.EntField), val)
	return ef.entityFieldEditor(v)
}

// 字段自增
func (ef *baseField) Inc(step ...int) xqi.FieldValue {
	v := ef.TSqlTableField.Inc(step...)
	return ef.entityFieldEditor(v)
}

// 字段自减
func (ef *baseField) Dec(step ...int) xqi.FieldValue {
	v := ef.TSqlTableField.Dec(step...)
	return ef.entityFieldEditor(v)
}

// 字段自乘
func (ef *baseField) UnaryMul(val interface{}) xqi.FieldValue {
	v := ef.TSqlTableField.UnaryMul(val)
	return ef.entityFieldEditor(v)
}

// 字段自除
func (ef *baseField) UnaryDiv(val interface{}) xqi.FieldValue {
	v := ef.TSqlTableField.UnaryDiv(val)
	return ef.entityFieldEditor(v)
}

func (ef *baseField) Desc() xqi.SqlOrderField {
	return xqcomm.NewSqlOrderField(ef, xqi.SqlOrderDescDirect)
}

func (ef *baseField) Asc() xqi.SqlOrderField {
	return xqcomm.NewSqlOrderField(ef, xqi.SqlOrderAscDirect)
}

func (ef *baseField) String() string {
	return fmt.Sprintf("%s:%v", ef.FieldName(), ef.This().(xqi.EntField).Value())
}

func (ef *baseField) FieldType() xqi.FieldDataType {
	return xqi.FDTUnknown
}

func (ef *baseField) FieldIndex() int {
	return ef.index
}

func (ef *baseField) Formatter() string {
	return ef.formatter
}
