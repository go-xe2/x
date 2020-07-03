package xqi

// 实体属性
type EntityAttribute interface {
	XqAttribute
	TableName() string
	TableAlias() string
}
