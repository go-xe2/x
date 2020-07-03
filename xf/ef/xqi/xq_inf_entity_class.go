package xqi

import (
	"github.com/go-xe2/x/xf/anno"
	"reflect"
)

type FieldConstructor = func(entity Entity, defineName string, attrs []XqAttribute, annotations map[string]interface{}, params ...interface{}) interface{}

type EntityFieldClass interface {
	FieldIndex() []int
	DefineName() string
	Annotations() map[string]anno.AnnotationContainer
	IsForeign() bool
	Constructor() FieldConstructor
	NewField(entity Entity, params ...interface{}) interface{}
}

type EntityClass interface {
	EntType() reflect.Type
	Constructor() func() Entity
	// 实体属性
	Attributes() []XqAttribute
	Annotations() map[string]map[string]anno.AnnotationContainer
	TableName() string
	TableAlias() string
	Create() interface{}
	Fields() map[string]EntityFieldClass
	UpdateValidItems() map[string]FieldValid
	InsertValidItems() map[string]FieldValid
	String() string
}
