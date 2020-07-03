package xclass

import (
	"fmt"
	"reflect"
	"strings"
)

func newClassVT(classType reflect.Type) *classVT {
	return &classVT{
		clsType:    classType,
		extends:    make([]*classVTExtendField, 0),
		fieldMaps:  make(map[string]*classVTField),
		methodMaps: make(map[string]*classVTMethod),
		classTag:   GetClassTag(classType),
	}
}

func (cr *classVT) AddField(field *classVTField) *classVT {
	if cr.fieldMaps == nil {
		cr.fieldMaps = make(map[string]*classVTField)
	}
	if _, ok := cr.fieldMaps[field.name]; !ok {
		cr.fieldMaps[field.name] = field
	}
	return cr
}

func (cr *classVT) AddMethod(method *classVTMethod) *classVT {
	if cr.methodMaps == nil {
		cr.methodMaps = make(map[string]*classVTMethod)
	}
	if _, ok := cr.methodMaps[method.name]; !ok {
		cr.methodMaps[method.name] = method
	}
	return cr
}

func (cr *classVT) AddExtend(extend *classVTExtendField) *classVT {
	if extend == nil {
		return cr
	}
	if cr.extends == nil {
		cr.extends = make([]*classVTExtendField, 0)
	}
	cr.extends = append(cr.extends, extend)
	return cr
}

func (cr *classVT) Print(leave int) string {
	extendItems := make([]string, 0)
	for _, v := range cr.extends {
		extendItems = append(extendItems, "\t"+v.Print(leave+1))
	}
	fieldItems := make([]string, 0)
	for _, v := range cr.fieldMaps {
		fieldItems = append(fieldItems, "\t"+v.Print(leave+1))
	}
	methodItems := make([]string, 0)
	for _, v := range cr.methodMaps {
		methodItems = append(methodItems, "\t"+v.Print(leave+1))
	}
	spaces := repeatString("\t", leave)
	szExtends := strings.Join(extendItems, "\n")
	if len(cr.extends) > 0 {
		szExtends = "[\n" + szExtends + "\n" + spaces + "\t]"
	} else {
		szExtends = "[]"
	}
	szFields := strings.Join(fieldItems, "\n")
	if len(cr.fieldMaps) > 0 {
		szFields = "[\n" + szFields + "\n" + spaces + "\t]"
	} else {
		szFields = "[]"
	}
	szMethods := strings.Join(methodItems, "\n")
	if len(cr.methodMaps) > 0 {
		szMethods = "[\n" + szMethods + "\n" + spaces + "\t]"
	} else {
		szMethods = "[]"
	}
	return fmt.Sprintf("%s %s {\n%s\textends: %s\n%s\tfields:%s\n%s\tmethods:%s\n%sclassTag:%s\n%s}\n",
		spaces, cr.clsType,
		spaces, szExtends,
		spaces, szFields,
		spaces, szMethods,
		spaces, cr.classTag,
		spaces)
}

func (cr *classVT) String() string {
	return cr.Print(0)
}
