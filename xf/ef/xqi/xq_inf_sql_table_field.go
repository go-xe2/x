package xqi

// SqlTableField兼容SqlField
type SqlTableField interface {
	SqlField
	Table() SqlTable
	FieldName() string
	Set(val interface{}) FieldValue
	// 字段自增
	Inc(step ...int) FieldValue
	// 字段自减
	Dec(step ...int) FieldValue
	// 字段自乘
	UnaryMul(val interface{}) FieldValue
	// 字段自乘
	UnaryDiv(val interface{}) FieldValue
}
