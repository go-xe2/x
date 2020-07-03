package xclass

import (
	"fmt"
	"reflect"
	"strings"
)

func newClassVTMethod(index int, name string) *classVTMethod {
	return &classVTMethod{
		index:       index,
		name:        name,
		paramTypes:  make([]reflect.Type, 0),
		resultTypes: make([]reflect.Type, 0),
	}
}

func (cm *classVTMethod) AddInParam(param reflect.Type) *classVTMethod {
	if cm.paramTypes == nil {
		cm.paramTypes = make([]reflect.Type, 0)
	}
	cm.paramTypes = append(cm.paramTypes, param)
	return cm
}

func (cm *classVTMethod) AddOutParams(param reflect.Type) *classVTMethod {
	if cm.resultTypes == nil {
		cm.resultTypes = make([]reflect.Type, 0)
	}
	cm.resultTypes = append(cm.resultTypes, param)
	return cm
}

func (cm *classVTMethod) Print(leave int) string {
	szParams := make([]string, 0)
	n := len(cm.paramTypes)
	for i := 0; i < n; i++ {
		szParams = append(szParams, cm.paramTypes[i].String())
	}
	resultItems := make([]string, 0)
	n = len(cm.resultTypes)
	for i := 0; i < n; i++ {
		resultItems = append(resultItems, cm.resultTypes[i].String())
	}
	szResult := strings.Join(resultItems, ",")
	if n > 1 {
		szResult = "(" + szResult + ")"
	}
	spaces := repeatString("\t", leave)
	return fmt.Sprintf("%s%v %s(%s) %s", spaces, cm.index, cm.name, strings.Join(szParams, ","), szResult)
}

func (cm *classVTMethod) String() string {
	return cm.Print(0)
}
