package xclass

import (
	"fmt"
)

func newClassVTExtendField(fieldIndex []int, extend *classVT) *classVTExtendField {
	return &classVTExtendField{
		fieldIndex: fieldIndex,
		extendType: extend,
	}
}

func (ext *classVTExtendField) SetOffset(offset uintptr) *classVTExtendField {
	ext.offset = offset
	return ext
}

func (ext *classVTExtendField) Print(leave int) string {
	spaces := repeatString("\t", leave)
	return fmt.Sprintf("%s%v\n %s", spaces, ext.fieldIndex, ext.extendType.Print(leave+1))
}

func (ext *classVTExtendField) String() string {
	return ext.Print(0)
}
