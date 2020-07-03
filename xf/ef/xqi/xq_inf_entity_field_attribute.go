package xqi

type EntityFieldAttribute interface {
	XqAttribute
	IsPrimary() bool
	// 字段名称
	FieldName() string
	FieldAlias() string
	// 字段查询规则
	Rule() string
	Formatter() string
}
