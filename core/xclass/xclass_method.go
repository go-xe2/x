package xclass

import (
	"fmt"
	"reflect"
	"strings"
)

type ClassMethod struct {
	name        string
	paramTypes  []reflect.Type
	resultTypes []reflect.Type
}

func newClassMethod(name string, params []reflect.Type, results []reflect.Type) *ClassMethod {
	return &ClassMethod{
		name:        name,
		paramTypes:  params,
		resultTypes: results,
	}
}

func (cm *ClassMethod) Name() string {
	return cm.name
}

func (cm *ClassMethod) ParamTypes() []reflect.Type {
	return cm.paramTypes
}

func (cm *ClassMethod) ResultTypes() []reflect.Type {
	return cm.resultTypes
}

func (cm *ClassMethod) String() string {
	szParams := make([]string, 0)
	for _, t := range cm.paramTypes {
		szParams = append(szParams, t.String())
	}
	szResults := make([]string, 0)
	for _, t := range cm.resultTypes {
		szResults = append(szResults, t.String())
	}
	s1 := strings.Join(szParams, ",")
	s2 := strings.Join(szResults, ",")
	if len(szResults) > 1 {
		s2 = "(" + s2 + ")"
	}
	return fmt.Sprintf("%s(%s)%s", cm.name, s1, s2)
}
