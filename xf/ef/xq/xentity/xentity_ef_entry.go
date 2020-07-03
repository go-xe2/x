package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xqi"
	"github.com/go-xe2/x/xf/xfboot"
	"reflect"
)

type EFTypeConstructorEntry interface {
	xfboot.BootEntry
	Register(typ reflect.Type, fieldDataType xqi.FieldDataType, construct xqi.FieldConstructor)
}

type userDefineFieldType struct {
	fieldDataType uint8
	gType         reflect.Type
	constructor   xqi.FieldConstructor
}

type tEFTypeConstructorEntry struct {
	types          map[xqi.FieldDataType]xqi.FieldConstructor
	userTypes      map[reflect.Type]uint8
	userFieldTypes map[uint8]*userDefineFieldType
}

var _ xfboot.BootEntry = (*tEFTypeConstructorEntry)(nil)

const EFTypeConstructorEntryName = "efTypeConstructorInit"

var efTypeConstructorEntry = xfboot.GetEntryOrRegister(EFTypeConstructorEntryName, func() xfboot.BootEntry {
	inst := &tEFTypeConstructorEntry{
		types:          make(map[xqi.FieldDataType]xqi.FieldConstructor),
		userTypes:      make(map[reflect.Type]uint8),
		userFieldTypes: make(map[uint8]*userDefineFieldType),
	}
	// 注册内置字段类型
	registerInternalType(inst)
	return inst
}).(*tEFTypeConstructorEntry)

func (fe *tEFTypeConstructorEntry) EntryName() string {
	return EFTypeConstructorEntryName
}

func (fe *tEFTypeConstructorEntry) Entry() interface{} {
	return fe
}

// 注册内部类型
func (fe *tEFTypeConstructorEntry) register(typ xqi.FieldDataType, constructor xqi.FieldConstructor) {
	if _, ok := fe.types[typ]; ok {
		return
	}
	fe.types[typ] = constructor
}

// 用户定义实体类型
func (fe *tEFTypeConstructorEntry) Register(typ reflect.Type, fieldDataType uint8, constructor xqi.FieldConstructor) {
	if fieldDataType < uint8(xqi.FDTDefineBase) {
		panic(exception.Newf("自定实体字段类型，请从FDTDefineBase编号开始"))
	}
	if _, ok := fe.userTypes[typ]; ok {
		panic(exception.Newf("字段类型%s已经注册过", typ.String()))
	}
	if _, ok := fe.userFieldTypes[fieldDataType]; ok {
		panic(exception.Newf("字段类型%s已经注册过", typ.String()))
	}
	item := &userDefineFieldType{
		fieldDataType: fieldDataType,
		gType:         typ,
		constructor:   constructor,
	}
	fe.userTypes[typ] = fieldDataType
	fe.userFieldTypes[fieldDataType] = item
}

func (fe *tEFTypeConstructorEntry) IsInit() bool {
	return xfboot.IsEntryInit(fe.EntryName())
}

func (fe *tEFTypeConstructorEntry) Init() {
	if fe.IsInit() {
		return
	}
	xfboot.InitEntry(fe.EntryName())
}

func (fe *tEFTypeConstructorEntry) IsUserTypeDefine(fieldDataType uint8) bool {
	if _, ok := fe.userFieldTypes[fieldDataType]; ok {
		return true
	}
	return false
}

func (fe *tEFTypeConstructorEntry) IsUserGTypeDefine(gType reflect.Type) bool {
	if _, ok := fe.userTypes[gType]; ok {
		return true
	}
	return false
}

func (fe *tEFTypeConstructorEntry) GetUserFieldType(gType reflect.Type) uint8 {
	if v, ok := fe.userTypes[gType]; ok {
		return v
	}
	return uint8(xqi.FDTUnknown)
}

func (fe *tEFTypeConstructorEntry) UserFieldTypeConstructor(fieldDataType uint8) (constructor xqi.FieldConstructor) {
	if fieldDataType == uint8(xqi.FDTUnknown) {
		return nil
	} else if fieldDataType < uint8(xqi.FDTDefineBase) {
		return fe.TypeConstructor(xqi.FieldDataType(fieldDataType))
	}
	if v, ok := fe.userFieldTypes[fieldDataType]; ok {
		return v.constructor
	}
	return nil
}

func (fe *tEFTypeConstructorEntry) UserTypeConstructor(gType reflect.Type) (constructor xqi.FieldConstructor) {
	n := fe.GetUserFieldType(gType)
	return fe.UserFieldTypeConstructor(n)
}

func (fe *tEFTypeConstructorEntry) IsDefine(typ xqi.FieldDataType) bool {
	if _, ok := fe.types[typ]; ok {
		return true
	}
	return false
}

func (fe *tEFTypeConstructorEntry) TypeConstructor(typ xqi.FieldDataType) (constructor xqi.FieldConstructor) {
	if v, ok := fe.types[typ]; ok {
		return v
	}
	return nil
}
