package xclass

import (
	"github.com/go-xe2/x/sync/xsafeMap"
	"reflect"
	"strings"
)

var classVTCache = xsafeMap.NewStrAnyMap()

func classNameSpace(className string) []string {
	return strings.Split(className, ".")
}

func classAliasName(fullName string) string {
	if fullName == "" {
		return ""
	}
	nameSpace := classNameSpace(fullName)
	return nameSpace[len(nameSpace)-1]
}

func loopClassToClassVT(outLays []*classVT, cls reflect.Type) *classVT {
	tmpTyp := cls
	for tmpTyp.Kind() == reflect.Ptr {
		tmpTyp = tmpTyp.Elem()
	}
	key := tmpTyp.PkgPath() + "." + tmpTyp.Name()
	if v := classVTCache.Get(key); v != nil {
		return v.(*classVT)
	}
	clsElem := cls
	for clsElem.Kind() == reflect.Ptr {
		clsElem = clsElem.Elem()
	}
	if !cls.Implements(classType) {
		panic("Create方法只接受Class的继承类")
	}
	// cls为指向struct的指针，即为*struct
	clsVT := newClassVT(cls)

	numMethod := cls.NumMethod()
	for i := 0; i < numMethod; i++ {
		method := cls.Method(i)
		methodVT := newClassVTMethod(method.Index, method.Name)
		numIn := method.Type.NumIn()
		for j := 0; j < numIn; j++ {
			methodVT.AddInParam(method.Type.In(j))
		}
		numOut := method.Type.NumOut()
		for j := 0; j < numOut; j++ {
			methodVT.AddOutParams(method.Type.Out(j))
		}
		clsVT.AddMethod(methodVT)
		if outLays != nil {
			for _, lay := range outLays {
				lay.AddMethod(methodVT)
			}
		}
	}

	numField := clsElem.NumField()
	selfLays := append(outLays, clsVT)
	layCount := len(outLays)
	for i := 0; i < numField; i++ {
		field := clsElem.Field(i)
		fieldType := field.Type
		fieldName := field.Name
		if fieldType.Kind() == reflect.Ptr {
			if classAliasName(fieldName) == classAliasName(field.Type.String()) {
				if field.Type.Implements(classType) {
					fieldClsVT := loopClassToClassVT(selfLays, fieldType)
					extendField := newClassVTExtendField(field.Index, fieldClsVT)
					extendField.SetOffset(field.Offset)
					clsVT.AddExtend(extendField)
					continue
				}
			}
		}
		// 不导出私有字段
		c := fieldName[0]

		canExport := field.PkgPath == ""
		canExport = canExport && !(c >= 'a' && c <= 'z') && c != '_'
		if !canExport {
			continue
		}
		fieldVT := newClassVTField(fieldName, field.Index, field.Type)
		fieldVT.SetRawValue(nil)
		fieldVT.SetTag(NewClassTag(string(field.Tag)))
		fieldVT.SetOffset(field.Offset)
		fieldVT.BindGetter(classFieldGetters[field.Type.Kind()])
		fieldVT.BindSetter(classFieldSetters[field.Type.Kind()])

		if layCount > 0 {
			fieldVT.SetOwner(clsVT)
		}
		clsVT.AddField(fieldVT)
		if outLays != nil {
			for _, lay := range outLays {
				lay.AddField(fieldVT)
			}
		}
	}

	classVTCache.Set(key, clsVT)
	return clsVT
}

func classToClassVT(cls reflect.Type) *classVT {
	lays := make([]*classVT, 0)
	vt := loopClassToClassVT(lays, cls)
	// 处理最外层字段列表，最外层字段列表
	return vt
}
