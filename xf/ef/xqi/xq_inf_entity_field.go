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
}

type EFInt interface {
	EntField
	Int() int
}

type EFInt8 interface {
	EntField
	Int8() int8
}

type EFInt16 interface {
	EntField
	Int16() int16
}

type EFInt32 interface {
	EntField
	Int32() int32
}

type EFInt64 interface {
	EntField
	Int64() int64
}

type EFUint interface {
	EntField
	Uint() uint
}

type EFUint8 interface {
	EntField
	Uint8() uint8
}

type EFUint16 interface {
	EntField
	Uint16() uint16
}

type EFUint32 interface {
	EntField
	Uint32() uint32
}

type EFUint64 interface {
	EntField
	Uint64() uint64
}

type EFFloat interface {
	EntField
	Float() float32
}

type EFDouble interface {
	EntField
	Double() float64
}

type EFBool interface {
	EntField
	Bool() bool
}

type EFByte interface {
	EntField
	Byte() byte
}

type EFDate interface {
	EntField
	Date() time.Time
}

type EFBinary interface {
	EntField
	Bytes() []byte
}
