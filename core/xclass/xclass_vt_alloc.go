package xclass

import (
	"reflect"
)

func classExtendAllocTree(ownerNode classInstTree, extNode *classVTExtendField) *classInstTreeNode {
	cls := extNode.extendType
	if cls == nil {
		return nil
	}
	elemType := cls.clsType
	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}
	instValue := reflect.New(elemType)
	inst := instValue.Interface()
	instNode := newClassInstTreeNode(ownerNode.Instance(), extNode, instValue, extNode.fieldIndex, inst)
	for _, extInfo := range cls.extends {
		extNode := classExtendAllocTree(instNode, extInfo)
		instNode.Extends(extNode)
	}
	return instNode
}

// 根据类型结构分类内存
func classAllocTree(cls *classVT) *classInstTreeRoot {
	if cls == nil {
		return nil
	}
	elemType := cls.clsType
	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}
	instValue := reflect.New(elemType)
	inst := instValue.Interface()
	instTree := newClassInstTreeRoot(inst)
	instTree.value = instValue
	// 创建父类
	for _, extendInfo := range cls.extends {
		extendNode := classExtendAllocTree(instTree, extendInfo)
		instTree.Extends(extendNode)
	}
	return instTree
}

func classAllocExtends(instExtends map[reflect.Type]*classExtendField, inst interface{}, instValue reflect.Value, root reflect.Value, extends []*classInstTreeNode) reflect.Value {
	count := len(extends)
	rootElem := root
	for rootElem.Kind() == reflect.Ptr {
		rootElem = rootElem.Elem()
	}
	for i := 0; i < count; i++ {
		extInfo := extends[i]
		extendField := rootElem.FieldByIndex(extInfo.fieldIndex)
		if extendField.CanSet() {
			extValue := classAllocExtends(instExtends, inst, instValue, extInfo.value, extInfo.extends)
			extendField.Set(extValue)
			// 收集继承的父类实例
			instExtends[extValue.Type()] = newClassExtendField(extInfo.field, extInfo.value, extInfo.Instance())
		}
	}
	// 只处理最顶层的this打针指向实例, 以免多次调用bindThis方法
	if count == 0 {
		v := root.Interface()
		if obj, ok := v.(*TObject); ok {
			obj.this = inst
			obj.extends = instExtends
			obj.classValue = instValue // 指针inst的指针
			obj.classTag = GetClassTag(instValue.Type())
		}
	}
	return root
}

func classAlloc(cls *classVT, props ...interface{}) interface{} {
	clsTree := classAllocTree(cls)
	if clsTree == nil {
		return nil
	}
	root := clsTree.value
	inst := clsTree.inst
	count := len(clsTree.extends)
	rootElem := root
	instExtends := make(map[reflect.Type]*classExtendField)
	for rootElem.Kind() == reflect.Ptr {
		rootElem = rootElem.Elem()
	}
	for i := 0; i < count; i++ {
		extInfo := clsTree.extends[i]
		extendField := rootElem.FieldByIndex(extInfo.fieldIndex)
		if extendField.CanSet() {
			extValue := classAllocExtends(instExtends, inst, root, extInfo.value, extInfo.extends)
			extendField.Set(extValue)
			// 收集继承的父类实例
			instExtends[extValue.Type()] = newClassExtendField(extInfo.field, extInfo.value, extInfo.Instance())
		}
	}
	// 调用构造函数
	if cls, ok := inst.(Object); ok {
		cls.Constructor(props...)
	}
	return inst
}
