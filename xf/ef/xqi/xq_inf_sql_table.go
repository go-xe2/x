package xqi

type SqlTableFields interface {
	AddField(field SqlField) SqlTable
}

type SqlTable interface {
	SqlCompiler
	TableName() string
	TableAlias() string
	Alias(name string) SqlTable
	AllField() []SqlField
	Field(index int) SqlField
	FieldByName(name string) SqlField
	// 源字段转换成查询结果中的字段
	SelField(field SqlField) SqlField
	String() string
	FieldCount() int
}
