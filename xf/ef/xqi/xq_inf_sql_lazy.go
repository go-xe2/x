package xqi

type (
	LazyGetJoinConditionFun func(joinTable SqlTable, tables SqlTables) SqlCondition
	LazyGetConditionFun     func(tables SqlTables) SqlCondition
	// 获取字段列表
	LazyGetFieldFun func(tables SqlTables) []SqlField
	// 获取排序字段列表
	LazyGetOrderFieldFun func(tables SqlTables) []SqlOrderField
)
