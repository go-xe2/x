package xclass

import "reflect"

type classInstTree interface {
	Extends(parent *classInstTreeNode) classInstTree
	Instance() interface{}
}

// class类型实例继承关系材料节点
type classInstTreeNode struct {
	// 继承类所在的字段序号
	fieldIndex []int
	// 字段所属struct实例
	owner interface{}
	// 字段实例
	inst    interface{}
	value   reflect.Value
	field   *classVTExtendField
	extends []*classInstTreeNode
}

type classInstTreeRoot struct {
	// 顶层节点
	inst    interface{}
	value   reflect.Value
	extends []*classInstTreeNode
}
