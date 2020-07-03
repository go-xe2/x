package structs

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"reflect"
)

type TField struct {
	self     interface{}
	owner    *TStruct
	name     string
	tag      *FieldTag
	fType    reflect.Type
	fValue   reflect.Value
	offset   uintptr
	index    []int
	setFuncV FieldSetFuncV
	setFunc  FieldSetFunc
	getFuncV FieldGetFuncV
	getFunc  FieldGetFunc
}

type TSliceField struct {
	TField
	itemType reflect.Type
}

type TMapField struct {
	TField
	keyType   reflect.Type
	valueType reflect.Type
}

type TPtrField struct {
	TField
	originType reflect.Type
}

type TStructField struct {
	TField
	structType reflect.Type
}

func newFieldOwner(owner *TStruct, name string, tag string, typ reflect.Type, val reflect.Value, index []int, offset uintptr, extends ...reflect.Type) (IField, error) {
	var field IField
	switch typ.Kind() {
	case reflect.Slice, reflect.Array:
		if len(extends) < 1 {
			return nil, exception.New("extends需要转入slice项类型")
		}
		obj := &TSliceField{
			TField: TField{
				owner:  owner,
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			itemType: extends[0],
		}
		field = obj.Init(obj)
		break
	case reflect.Map:
		if len(extends) < 2 {
			return nil, exception.New("extends需要传入map的key及value的2个类型参数")
		}
		keyTyp := extends[0]
		valTyp := extends[1]

		obj := &TMapField{
			TField: TField{
				owner:  owner,
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			keyType:   keyTyp,
			valueType: valTyp,
		}
		field = obj.Init(obj)
		break
	case reflect.Struct:
		if len(extends) < 1 {
			return nil, exception.New("extends需要传入struct的类型")
		}
		obj := &TStructField{
			TField: TField{
				owner:  owner,
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			structType: extends[0],
		}
		field = obj.Init(obj)
		break
	case reflect.Ptr:
		if len(extends) < 1 {
			return nil, exception.New("extends需要转入ptr指向的类型")
		}
		obj := &TPtrField{
			TField: TField{
				owner:  owner,
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			originType: extends[0],
		}
		field = obj.Init(obj)
		break
	default:
		obj := &TField{
			owner:  owner,
			name:   name,
			tag:    NewFieldTag(tag),
			fType:  typ,
			fValue: val,
			index:  index,
			offset: offset,
		}
		field = obj.Init(obj)
	}
	return field, nil
}

// 创建struct字段信息记录
func NewField(name string, tag string, typ reflect.Type, val reflect.Value, index []int, offset uintptr, extends ...reflect.Type) (IField, error) {
	var field IField
	switch typ.Kind() {
	case reflect.Slice, reflect.Array:
		if len(extends) < 1 {
			return nil, exception.New("extends需要转入slice项类型")
		}
		obj := &TSliceField{
			TField: TField{
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			itemType: extends[0],
		}
		field = obj.Init(obj)
		break
	case reflect.Map:
		if len(extends) < 2 {
			return nil, exception.New("extends需要传入map的key及value的2个类型参数")
		}
		keyTyp := extends[0]
		valTyp := extends[1]

		obj := &TMapField{
			TField: TField{
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			keyType:   keyTyp,
			valueType: valTyp,
		}
		field = obj.Init(obj)
		break
	case reflect.Struct:
		if len(extends) < 1 {
			return nil, exception.New("extends需要传入struct的类型")
		}
		obj := &TStructField{
			TField: TField{
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			structType: extends[0],
		}
		field = obj.Init(obj)
		break
	case reflect.Ptr:
		if len(extends) < 1 {
			return nil, exception.New("extends需要转入ptr指向的类型")
		}
		obj := &TPtrField{
			TField: TField{
				name:   name,
				tag:    NewFieldTag(tag),
				fType:  typ,
				fValue: val,
				index:  index,
				offset: offset,
			},
			originType: extends[0],
		}
		field = obj.Init(obj)
		break
	default:
		obj := &TField{
			name:   name,
			tag:    NewFieldTag(tag),
			fType:  typ,
			fValue: val,
			index:  index,
			offset: offset,
		}
		field = obj.Init(obj)
	}
	return field, nil
}

func (f *TField) Init(inst interface{}) IField {
	f.self = inst
	switch f.Type().Kind() {
	case reflect.String:
		f.setFuncV = FieldSetStringFuncV
		f.setFunc = FieldSetStringFunc
		f.getFunc = FieldGetStringFunc
		f.getFuncV = FieldGetStringFuncV
		break
	case reflect.Int:
		f.setFunc = FieldSetIntFunc
		f.setFuncV = FieldSetIntFuncV
		f.getFunc = FieldGetIntFunc
		f.getFuncV = FieldGetIntFuncV
		break
	case reflect.Int8:
		f.setFunc = FieldSetInt8Func
		f.setFuncV = FieldSetInt8FuncV
		f.getFunc = FieldGetInt8Func
		f.getFuncV = FieldGetInt8FuncV
		break
	case reflect.Int16:
		f.setFunc = FieldSetInt16Func
		f.setFuncV = FieldSetInt16FuncV
		f.getFunc = FieldGetInt16Func
		f.getFuncV = FieldGetInt16FuncV
		break
	case reflect.Int32:
		f.setFunc = FieldSetInt32Func
		f.setFuncV = FieldSetInt32FuncV
		f.getFunc = FieldGetInt32Func
		f.getFuncV = FieldGetInt32FuncV
		break
	case reflect.Int64:
		f.setFunc = FieldSetInt64Func
		f.setFuncV = FieldSetInt64FuncV
		f.getFunc = FieldGetInt64Func
		f.getFuncV = FieldGetInt64FuncV
		break
	case reflect.Uint:
		f.setFunc = FieldSetUintFunc
		f.setFuncV = FieldSetUintFuncV
		f.getFunc = FieldGetUintFunc
		f.getFuncV = FieldGetUintFuncV
		break
	case reflect.Uint8:
		f.setFunc = FieldSetUint8Func
		f.setFuncV = FieldSetUint8FuncV
		f.getFunc = FieldGetUint8Func
		f.getFuncV = FieldGetUint8FuncV
		break
	case reflect.Uint16:
		f.setFunc = FieldSetInt16Func
		f.setFuncV = FieldSetInt16FuncV
		f.getFunc = FieldGetUint16Func
		f.getFuncV = FieldGetUint16FuncV
		break
	case reflect.Uint32:
		f.setFunc = FieldSetUint32Func
		f.setFuncV = FieldSetUint32FuncV
		f.getFunc = FieldGetUint32Func
		f.getFuncV = FieldGetUint32FuncV
		break
	case reflect.Uint64:
		f.setFunc = FieldSetUint64Func
		f.setFuncV = FieldSetUint64FuncV
		f.getFunc = FieldGetUint64Func
		f.getFuncV = FieldGetUint64FuncV
		break
	case reflect.Bool:
		f.setFunc = FieldSetBoolFunc
		f.setFuncV = FieldSetBoolFuncV
		f.getFunc = FieldGetBoolFunc
		f.getFuncV = FieldGetBoolFuncV
		break
	case reflect.Float32:
		f.setFunc = FieldSetFloat32Func
		f.setFuncV = FieldSetFloat32FuncV
		f.getFunc = FieldGetFloat32Func
		f.getFuncV = FieldGetFloat32FuncV
		break
	case reflect.Float64:
		f.setFunc = FieldSetFloat64Func
		f.setFuncV = FieldSetFloat64FuncV
		f.getFunc = FieldGetFloat64Func
		f.getFuncV = FieldGetFloat64FuncV
		break
	case reflect.Complex64:
		f.setFunc = FieldSetComplex64Func
		f.setFuncV = FieldSetComplex64FuncV
		f.getFunc = FieldGetComplex64Func
		f.getFuncV = FieldGetComplex64FuncV
		break
	case reflect.Complex128:
		f.setFunc = FieldSetComplex64Func
		f.setFuncV = FieldSetComplex128FuncV
		f.getFunc = FieldGetComplex128Func
		f.getFuncV = FieldGetComplex128FuncV
		break
	case reflect.Chan:
		f.setFunc = FieldSetChanFunc
		f.setFuncV = FieldSetChanFuncV
		f.getFunc = FieldGetChanFunc
		f.getFuncV = FieldGetChanFuncV
		break
	case reflect.Slice, reflect.Array:
		f.setFunc = FieldSetSliceFunc
		f.setFuncV = FieldSetSliceFuncV
		f.getFunc = FieldGetSliceFunc
		f.getFuncV = FieldGetSliceFuncV
		break
	case reflect.Ptr:
		f.setFunc = FieldSetPtrFunc
		f.setFuncV = FieldSetPtrFuncV
		f.getFunc = FieldGetInterfaceFunc
		f.getFuncV = FieldGetInterfaceFuncV
		break
	case reflect.Map:
		f.setFunc = FieldSetMapFunc
		f.setFuncV = FieldSetMapFuncV
		f.getFunc = FieldGetMapFunc
		f.getFuncV = FieldGetMapFuncV
		break
	case reflect.Struct:
		f.setFunc = FieldSetStructFunc
		f.setFuncV = FieldSetStructFuncV
		f.getFunc = FieldGetInterfaceFunc
		f.getFuncV = FieldGetInterfaceFuncV
		break
	case reflect.Uintptr:
		f.setFunc = FieldSetUIntPtrFunc
		f.setFuncV = FieldSetUIntPtrFuncV
		f.getFunc = FieldGetInterfaceFunc
		f.getFuncV = FieldGetInterfaceFuncV
		break
	}
	return f
}

func (f *TField) Name() string {
	return f.name
}

func (f *TField) Value() reflect.Value {
	return f.fValue
}

func (f *TField) Tag() *FieldTag {
	return f.tag
}

func (f *TField) Type() reflect.Type {
	return f.fType
}

func (f *TField) Offset() uintptr {
	return f.offset
}

func (f *TField) Index() []int {
	return f.index
}

// 设置字段值，异常时panic
func (f *TField) Get(inst ...interface{}) interface{} {
	var instance interface{}
	if len(inst) == 0 && f.owner == nil {
		panic(exception.New("未传入字段所属对象实例"))
	}
	if len(inst) > 0 {
		instance = inst[0]
	} else {
		instance = f.owner.Instance()
	}
	if f.getFunc == nil {
		panic(exception.New("未绑定字段设置参数，请调用先调用Init方法初始化反射字段信息"))
	}
	if v, err := f.getFunc(instance, f); err != nil {
		panic(exception.Wrapf(err, "读取字段%s值出错", f.name))
	} else {
		return v
	}
}

// 设置字段值，异常时panic
func (f *TField) GetV(instValue ...reflect.Value) interface{} {
	var instance reflect.Value
	if len(instValue) == 0 && f.owner == nil {
		panic(exception.New("未传入字段所属对象实例"))
	}
	if len(instValue) > 0 {
		instance = instValue[0]
	} else {
		instance = reflect.ValueOf(f.owner.Instance())
	}
	if f.getFuncV == nil {
		panic(exception.New("未绑定字段设置参数，请调用先调用Init方法初始化反射字段信息"))
	}
	if v, err := f.getFuncV(instance, f); err != nil {
		panic(exception.Wrapf(err, "读取字段%s值出错", f.name))
	} else {
		return v
	}
}

func (f *TField) Set(val interface{}, inst ...interface{}) {
	var instance interface{}
	if len(inst) == 0 && f.owner == nil {
		panic(exception.New("未传入字段所属对象实例"))
	}
	if len(inst) > 0 {
		instance = inst[0]
	} else {
		instance = f.owner.Instance()
	}
	if f.setFunc == nil {
		panic(exception.New("未绑定字段设置参数，请调用先调用Init方法初始化反射字段信息"))
	}
	if err := f.setFunc(instance, f, val); err != nil {
		panic(exception.Wrapf(err, "设置字段%s值出错", f.name))
	}
}

func (f *TField) SetV(val interface{}, instValue ...reflect.Value) {
	var instance reflect.Value
	if len(instValue) == 0 && f.owner == nil {
		panic(exception.New("未传入字段所属对象实例"))
	}
	if len(instValue) > 0 {
		instance = instValue[0]
	} else {
		instance = reflect.ValueOf(f.owner.Instance())
	}
	if f.setFunc == nil {
		panic(exception.New("未绑定字段设置参数，请调用先调用Init方法初始化反射字段信息"))
	}
	if err := f.setFunc(instance, f, val); err != nil {
		panic(exception.Wrapf(err, "设置字段%s值出错", f.name))
	}
}

func (f *TField) Self() interface{} {
	return f.self
}

func (f *TField) String() string {
	return fmt.Sprintf("%s %s %d", f.name, f.fType.Kind().String(), f.offset)
}

/*
sliceField methods
*/

func (sf *TSliceField) ItemType() reflect.Type {
	return sf.itemType
}

func (sf *TSliceField) String() string {
	return fmt.Sprintf("%s []%s %d", sf.name, sf.itemType.Name(), sf.offset)
}

/*
mapField methods
*/

func (mf *TMapField) KeyType() reflect.Type {
	return mf.keyType
}

func (mf *TMapField) ValueType() reflect.Type {
	return mf.valueType
}

func (mf *TMapField) String() string {
	return fmt.Sprintf("%s map[%s]%s %d", mf.name, mf.keyType.Name(), mf.valueType.Name(), mf.offset)
}

/*
ptrField methods
*/

func (pf *TPtrField) OriginType() reflect.Type {
	return pf.originType
}

func (pf *TPtrField) String() string {
	return fmt.Sprintf("%s *%s %d", pf.name, pf.originType.Name(), pf.offset)
}

/*
TstructField methods
*/

func (sf *TStructField) StructType() reflect.Type {
	return sf.structType
}

func (sf *TStructField) String() string {
	return fmt.Sprintf("%s %s %d", sf.name, sf.structType.Name(), sf.offset)
}
