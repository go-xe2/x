package xqi

import "time"

type EntFieldIndex interface {
	GetIndex() int
	SetIndex(index int)
}

type EntField interface {
	SqlField
	// 字段所属的实体
	Entity() Entity
	Table() SqlTable
	FieldName() string
	Rule() string
	GetAnnotation(annName string) interface{}
	// 定义的名称
	DefineName() string
	Value() interface{}
	// 是否打开
	IsOpen() bool
	TryGetVal() interface{}
	Set(val interface{}) FieldValue
	Supper() EntField
	// 获取字段元注解
	IsPrimary() bool
	// 是否外联字段
	IsForeign() bool

	// 更新时字段运算
	// 字段自增
	Inc(step ...int) FieldValue
	// 字段自减
	Dec(step ...int) FieldValue
	// 字段自乘
	UnaryMul(val interface{}) FieldValue
	// 字段自乘
	UnaryDiv(val interface{}) FieldValue
	Formatter() string
	NewInstance(alias string, inherited ...interface{}) EntField
}

// 外联字段
type EFForeign interface {
	EntField
	// 外联关键字
	ForeignKey() string
	// 连接方式
	JoinType() SqlJoinType
	JoinTable() SqlTable
	LookField() SqlField
	On() func(on SqlCondition, joinEnt interface{}, tables SqlTables) SqlCondition
}

type EFExpr interface {
	EntField
	Expression(expr func(ent Entity) SqlField)
}

type EFString interface {
	EntField
	Str() string
	TryStr() (string, bool)
}

type EFInt interface {
	EntField
	Int() int
	TryInt() (int, bool)
}

type EFInt8 interface {
	EntField
	Int8() int8
	TryInt8() (int8, bool)
}

type EFInt16 interface {
	EntField
	Int16() int16
	TryInt16() (int16, bool)
}

type EFInt32 interface {
	EntField
	Int32() int32
	TryInt32() (int32, bool)
}

type EFInt64 interface {
	EntField
	Int64() int64
	TryInt64() (int64, bool)
}

type EFUint interface {
	EntField
	Uint() uint
	TryUint() (uint, bool)
}

type EFUint8 interface {
	EntField
	Uint8() uint8
	TryUint8() (uint8, bool)
}

type EFUint16 interface {
	EntField
	Uint16() uint16
	TryUint16() (uint16, bool)
}

type EFUint32 interface {
	EntField
	Uint32() uint32
	TryUint32() (uint32, bool)
}

type EFUint64 interface {
	EntField
	Uint64() uint64
	TryUint64() (uint64, bool)
}

type EFFloat interface {
	EntField
	Float() float32
	TryFloat() (float32, bool)
}

type EFDouble interface {
	EntField
	Double() float64
	TryDouble() (float64, bool)
}

type EFBool interface {
	EntField
	Bool() bool
	TryBool() (bool, bool)
}

type EFByte interface {
	EntField
	Byte() byte
	TryByte() (byte, bool)
}

type EFDate interface {
	EntField
	Date() time.Time
	TryDate() (time.Time, bool)
}

type EFBinary interface {
	EntField
	Bytes() []byte
	TryBytes() ([]byte, bool)
}
