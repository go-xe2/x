package xclass

import (
	"fmt"
	"reflect"
)

type ClassField struct {
	name   string
	tag    ClassTag
	fdType reflect.Type
}

func newClassField(name string, tag ClassTag, fdType reflect.Type) *ClassField {
	return &ClassField{
		name:   name,
		tag:    tag,
		fdType: fdType,
	}
}

func (cf *ClassField) Name() string {
	return cf.name
}

func (cf *ClassField) Tag() ClassTag {
	return cf.tag
}

func (cf *ClassField) Type() reflect.Type {
	return cf.fdType
}

func (cf *ClassField) String() string {
	return fmt.Sprintf("%s %s tag:%s", cf.name, cf.fdType.String(), cf.tag.String())
}
