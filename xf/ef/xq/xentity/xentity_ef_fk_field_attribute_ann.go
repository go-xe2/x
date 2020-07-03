package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/anno"
	"github.com/go-xe2/x/xf/xfboot"
)

type tEntityForeignFieldAttributeAnn struct {
}

var entForeignKeys = xsafeMap.NewStrAnyMap()

const ForeignAnnName = "foreign"

var _ anno.Annotation = (*tEntityForeignFieldAttributeAnn)(nil)

func (ann *tEntityForeignFieldAttributeAnn) AnnotationName() string {
	return ForeignAnnName
}

func (ann *tEntityForeignFieldAttributeAnn) AnnCreate(caller interface{}, annParams map[string]interface{}, callParams ...interface{}) interface{} {
	foreignKey := t.String(annParams["fk"])
	alias := t.String(annParams["alias"])
	rule := t.String(annParams["rule"])
	formatter := t.String(annParams["format"])
	return newEntityForeignFieldAttribute(foreignKey, alias, rule, formatter)
}

func (ann *tEntityForeignFieldAttributeAnn) Instance() interface{} {
	return ann
}

func registerFieldDefine(fieldDefine *EntForeignFieldDefine) {
	if entForeignKeys.Contains(fieldDefine.ForeignKey) {
		panic(exception.Newf("实体外键字段%s已经存在", fieldDefine.ForeignKey))
	}
	entForeignKeys.Set(fieldDefine.ForeignKey, fieldDefine)
}

func getFieldDefine(foreignKey string) *EntForeignFieldDefine {
	if v := entForeignKeys.Get(foreignKey); v != nil {
		return v.(*EntForeignFieldDefine)
	}
	return nil
}

var _ = xfboot.RegisterEntryPoint(anno.AnnotationEntryName, foreignKeyAnnEntryPoint)

// 元注解初始化入口，注册外联字段类型元注解
func foreignKeyAnnEntryPoint(entry xfboot.BootEntry) {
	point := entry.(anno.AnnotationEntry)
	point.Register(&tEntityForeignFieldAttributeAnn{})
}
