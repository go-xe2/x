package xclass

import (
	"reflect"
	"unsafe"
)

func (o *TObject) initClassVT() {
	if o.vt == nil {
		o.vt = classToClassVT(o.classValue.Type())
	}
}

func (o *TObject) getFieldValueByName(fieldName string) (fd reflect.Value, fieldOwner reflect.Value, fieldInfo *classVTField) {
	var ok = false
	fieldInfo, ok = o.vt.fieldMaps[fieldName]
	if ok {
		if fieldInfo.owner == nil || (fieldInfo.owner != nil && fieldInfo.owner.clsType.Elem() == o.classValue.Elem().Type()) {
			// 继承类的字段
			fieldOwner = o.classValue
			fd = reflect.NewAt(fieldInfo.fieldType, unsafe.Pointer(o.classValue.Elem().UnsafeAddr()+fieldInfo.offset)).Elem()
		} else {
			// 在父类中定义的字段
			parent, ok := o.extends[fieldInfo.owner.clsType]
			if ok {
				fieldOwner = parent.value
				fd = reflect.NewAt(fieldInfo.fieldType, unsafe.Pointer(parent.value.Elem().UnsafeAddr()+fieldInfo.offset)).Elem()
			}
		}
	}
	return
}
