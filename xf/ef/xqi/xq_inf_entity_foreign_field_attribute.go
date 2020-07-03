package xqi

type EntityForeignFieldAttribute interface {
	XqAttribute
	// 外键名称
	ForeignKey() string
	// 字段别名
	FieldAlias() string
	Rule() string
}
