package xqi

// 字段赋值表达式
type FieldValue interface {
	// 字段
	Field() SqlTableField
	// 给字段设置的值,可以是sql表达式或实际值
	Value() interface{}
	// 是否是主键
	IsPrimaryKey() bool
	String() string
}
