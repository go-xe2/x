package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/anno"
	"github.com/go-xe2/x/xf/xfboot"
)

type entityDbTypeAttributeAnn struct {
}

const DbTypeAnnName = "dbType"

var _ anno.Annotation = (*entityDbTypeAttributeAnn)(nil)

func (ann *entityDbTypeAttributeAnn) AnnotationName() string {
	return DbTypeAnnName
}

func (ann *entityDbTypeAttributeAnn) AnnCreate(caller interface{}, params map[string]interface{}, callParams ...interface{}) interface{} {
	inst := &tEntityDbTypeAttribute{}
	inst.dataType.Parse(t.String(params["type"]))
	inst.dataSize = t.Int(params["size"])
	inst.decimal = t.Int(params["decimal"])
	inst.allowNull = t.Bool(params["allowNull"])
	inst.increment = t.Bool(params["increment"])
	inst.defaultValue = t.String(params["default"])
	return inst
}

func (ann *entityDbTypeAttributeAnn) Instance() interface{} {
	return ann
}

// 注册元注解类型
func dbTypeAnnEntryPoint(entry xfboot.BootEntry) {
	point := entry.(anno.AnnotationEntry)
	point.Register(&entityDbTypeAttributeAnn{})
}

var _ = xfboot.RegisterEntryPoint(anno.AnnotationEntryName, dbTypeAnnEntryPoint)
