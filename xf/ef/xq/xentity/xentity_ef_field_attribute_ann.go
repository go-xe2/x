package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/anno"
	"github.com/go-xe2/x/xf/xfboot"
)

type entityFieldAttributeAnn struct {
}

const FieldAnnName = "field"

var _ anno.Annotation = (*entityFieldAttributeAnn)(nil)

func (ann *entityFieldAttributeAnn) AnnotationName() string {
	return FieldAnnName
}

func (ann *entityFieldAttributeAnn) AnnCreate(caller interface{}, params map[string]interface{}, callParams ...interface{}) interface{} {
	return NewFieldAttribute(t.String(params["name"]), t.String(params["alias"]), t.String(params["rule"]), t.String(params["format"]), t.Bool(params["primary"]))
}

func (ann *entityFieldAttributeAnn) Instance() interface{} {
	return ann
}

// 注册元注解类型
func fieldTagAnnEntryPoint(entry xfboot.BootEntry) {
	point := entry.(anno.AnnotationEntry)
	point.Register(&entityFieldAttributeAnn{})
}

var _ = xfboot.RegisterEntryPoint(anno.AnnotationEntryName, fieldTagAnnEntryPoint)
