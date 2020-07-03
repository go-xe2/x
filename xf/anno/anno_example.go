package anno

import (
	"github.com/go-xe2/x/type/t"
)

type ExampleAnn struct {
	Name string
	Sex  bool
}

type ExampleAnnoClass struct {
}

var _ Annotation = (*ExampleAnnoClass)(nil)

// 以下为注册元注解类
//var _ = xfboot.RegisterEntryPoint(AnnotationEntryName, func(point xfboot.BootEntry) {
//	entry := point.(AnnotationEntry)
//	entry.Register(&ExampleAnnoClass{})
//})

// 元注解名称
func (anc *ExampleAnnoClass) AnnotationName() string {
	return "ExampleAnno"
}

// 元注解创建初始化
func (anc *ExampleAnnoClass) AnnCreate(caller interface{}, annParams map[string]interface{}, callParams ...interface{}) interface{} {
	ann := &ExampleAnn{}
	ann.Name = t.String(annParams["name"])
	ann.Sex = t.Bool(annParams["sex"])
	return ann
}

func (anc *ExampleAnnoClass) Instance() interface{} {
	return anc
}
