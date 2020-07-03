package structs

import (
	"github.com/go-xe2/x/core/exception"
	"reflect"
)

type TStruct struct {
	instance  interface{}
	value     reflect.Value
	fieldMap  map[string]IField
	fields    []IField
	methodMap map[string]*MethodInfo
	methods   []*MethodInfo
	tagName   string
	ExtraCols []string
	result    []map[string]interface{}
}

func New(v interface{}, tagName ...string) (*TStruct, error) {
	defTagName := "json"
	if len(tagName) > 0 {
		defTagName = tagName[0]
	}
	val, err := lookValue(v)
	if err != nil {
		return nil, err
	}
	return &TStruct{
		instance: v,
		value:    val,
		tagName:  defTagName,
	}, nil
}

func (s *TStruct) Name() string {
	return s.value.Type().Name()
}

func (s *TStruct) TagName() string {
	return s.tagName
}

func (s *TStruct) initFields() error {
	fields, err := getStructFields(s.value, s.tagName, s)
	if err != nil {
		return err
	}
	s.fields = make([]IField, 0)
	s.fieldMap = make(map[string]IField)
	for _, field := range fields {
		s.fieldMap[field.Name()] = field
		s.fields = append(s.fields, field)
		tag := field.Tag().Get(s.tagName)
		if !tag.IsEmpty() {
			s.fieldMap[tag.Value()] = field
		}
		// 处理内嵌struct
		if field.Name() == field.Type().Name() {
			// 内嵌struct
			innerFields, err := getStructFields(field.Value(), s.tagName, s)
			if err != nil {
				return err
			}
			for _, innerField := range innerFields {
				if _, ok := s.fieldMap[innerField.Name()]; !ok {
					s.fields = append(s.fields, innerField)
					s.fieldMap[innerField.Name()] = innerField
				}
			}
		}
	}
	return nil
}

func (s *TStruct) Fields() []IField {
	if s.fields == nil {
		s.initFields()
	}
	return s.fields
}

func (s *TStruct) FieldMap() map[string]IField {
	if s.fieldMap == nil {
		s.initFields()
	}
	return s.fieldMap
}

func (s *TStruct) HasField(fieldName string) bool {
	fm := s.FieldMap()
	if _, ok := fm[fieldName]; ok {
		return true
	}
	return false
}

func (s *TStruct) Field(fieldName string) IField {
	fm := s.FieldMap()
	if v, ok := fm[fieldName]; ok {
		return v
	}
	return nil
}

func (s *TStruct) GetFieldVal(fieldName string) interface{} {
	fd := s.Field(fieldName)
	if fd == nil {
		return nil
	}
	return fd.Get(s.instance)
}

func (s *TStruct) SetFieldVal(fieldName string, val interface{}) {
	fd := s.Field(fieldName)
	if fd == nil {
		panic(exception.Newf("字段%s不存在", fieldName))
	}
	fd.Set(val, s.instance)
}

func (s *TStruct) Instance() interface{} {
	return s.instance
}

func (s *TStruct) initMethods() {
	if methods, err := getStructMethods(s.instance); err != nil {
		panic(err)
	} else {
		s.methods = methods
		s.methodMap = make(map[string]*MethodInfo)
		for _, method := range methods {
			methodName := method.Name()
			s.methodMap[methodName] = method
		}
	}
}

func (s *TStruct) Methods() []*MethodInfo {
	if s.methods == nil {
		s.initMethods()
	}
	return s.methods
}

func (s *TStruct) MethodMap() map[string]*MethodInfo {
	if s.methodMap == nil {
		s.initMethods()
	}
	return s.methodMap
}

func (s *TStruct) HasMethod(method string) bool {
	mm := s.MethodMap()
	if _, ok := mm[method]; ok {
		return true
	}
	return false
}

func (s *TStruct) Method(method string) *MethodInfo {
	mm := s.MethodMap()
	if m, ok := mm[method]; ok {
		return m
	}
	return nil
}

func (s *TStruct) Call(method string, params ...interface{}) (result []interface{}, err error) {
	mi := s.Method(method)
	if mi == nil {
		return nil, exception.Newf("方法%s不存在", method)
	}
	result, err = mi.Invoke(params...)
	return
}
