package xclass

import (
	"fmt"
	"reflect"
	"strings"
)

var _ classInstTree = (*classInstTreeRoot)(nil)
var _ classInstTree = (*classInstTreeNode)(nil)

func newClassInstTreeRoot(inst interface{}) *classInstTreeRoot {
	return &classInstTreeRoot{
		inst:    inst,
		extends: nil,
	}
}

func newClassInstTreeNode(owner interface{}, field *classVTExtendField, value reflect.Value, index []int, inst interface{}) *classInstTreeNode {
	return &classInstTreeNode{
		owner:      owner,
		fieldIndex: index,
		inst:       inst,
		value:      value,
		extends:    nil,
		field:      field,
	}
}

func (trt *classInstTreeRoot) Extends(parent *classInstTreeNode) classInstTree {
	if trt.extends == nil {
		trt.extends = make([]*classInstTreeNode, 0)
	}
	trt.extends = append(trt.extends, parent)
	return trt
}

func (trt *classInstTreeRoot) Instance() interface{} {
	return trt.inst
}

func (trt *classInstTreeRoot) String() string {
	count := len(trt.extends)
	szExtends := make([]string, 0)
	for i := 0; i < count; i++ {
		s := trt.extends[i].String()
		szExtends = append(szExtends, s)
	}
	return fmt.Sprintf("%s extends [%s]", reflect.TypeOf(trt.inst), strings.Join(szExtends, ","))
}

func (trn *classInstTreeNode) Extends(parent *classInstTreeNode) classInstTree {
	if trn.extends == nil {
		trn.extends = make([]*classInstTreeNode, 0)
	}
	trn.extends = append(trn.extends, parent)
	return trn
}

func (trn *classInstTreeNode) Instance() interface{} {
	return trn.inst
}

func (trn *classInstTreeNode) String() string {
	szExtends := make([]string, 0)
	count := len(trn.extends)
	for i := 0; i < count; i++ {
		s := trn.extends[i].String()
		szExtends = append(szExtends, s)
	}
	return fmt.Sprintf("%s[%v] extends [%s]", reflect.TypeOf(trn.inst), trn.fieldIndex, strings.Join(szExtends, ","))
}
