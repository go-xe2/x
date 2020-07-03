package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/anno"
	"github.com/go-xe2/x/xf/xfboot"
)

type tFieldValidAnn struct {
}

const FieldValidAnnName = "valid"

var _ anno.Annotation = (*tFieldValidAnn)(nil)

func (ann *tFieldValidAnn) AnnotationName() string {
	return FieldValidAnnName
}

func (ann *tFieldValidAnn) AnnCreate(caller interface{}, annParams map[string]interface{}, callParams ...interface{}) interface{} {
	inst := NewFieldValid(t.String(annParams["rule"]), t.String(annParams["msg"]), t.String(annParams["op"]), t.String(annParams["cate"]))
	return inst
}

func (ann *tFieldValidAnn) Instance() interface{} {
	return ann
}

var _ = xfboot.RegisterEntryPoint(anno.AnnotationEntryName, fieldValidAnnEntryPoint)

// 元注解初始化入口，注册外联字段类型元注解
func fieldValidAnnEntryPoint(entry xfboot.BootEntry) {
	point := entry.(anno.AnnotationEntry)
	point.Register(&tFieldValidAnn{})
}
