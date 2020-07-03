package xclass

import (
	"fmt"
	"reflect"
)

func newClassVTField(name string, index []int, fdType reflect.Type) *classVTField {
	return &classVTField{
		index:     index,
		name:      name,
		fieldType: fdType,
	}
}

func (cf *classVTField) SetTag(tag ClassTag) *classVTField {
	cf.fieldTag = tag
	return cf
}

func (cf *classVTField) SetCanSet(canSet bool) *classVTField {
	cf.canSet = canSet
	return cf
}

func (cf *classVTField) SetRawValue(value interface{}) *classVTField {
	cf.rawValue = value
	return cf
}

func (cf *classVTField) SetOwner(owner *classVT) *classVTField {
	cf.owner = owner
	return cf
}

func (cf *classVTField) SetOffset(offset uintptr) *classVTField {
	cf.offset = offset
	return cf
}

func (cf *classVTField) BindSetter(setter classFieldSetter) *classVTField {
	cf.setter = setter
	return cf
}

func (cf *classVTField) BindGetter(getter classFieldGetter) *classVTField {
	cf.getter = getter
	return cf
}

func (cf *classVTField) Print(leave int) string {
	spaces := repeatString("\t", leave)
	szOwner := ""
	if cf.owner != nil {
		szOwner = cf.owner.clsType.String()
	}
	return fmt.Sprintf("%s%v %s.%s %s", spaces, cf.index, szOwner, cf.name, cf.fieldType)
}

func (cf *classVTField) String() string {
	return cf.Print(0)
}
