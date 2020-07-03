package xqi

type SqlUpdateExpInfo interface {
	// 获取要更新的表
	GetTable() SqlTable
	// 获取要更新的字段列表
	GetFields() []FieldValue
	// 获取要更新的条件
	GetWhere() SqlCondition
	// 获取更新的数据来源表,如果非从其他数据源更新，返回nil
	GetJoins() []SqlJoin
}

type SqlUpdateExp interface {
	SqlCompiler
	SqlUpdateExpInfo

	// 要更新的表
	Table(table SqlTable) SqlUpdateExp
	// 更新字段
	Set(fields ...FieldValue) SqlUpdateExp
	// 更新数据条件
	Where(condition SqlCondition) SqlUpdateExp
	// 从其他数据源关联更新
	Join(joinType SqlJoinType, table SqlTable, on func(join SqlTable, others SqlTables) SqlCondition) SqlUpdateExp
	InnerJoin(table SqlTable, on func(join SqlTable, others SqlTables) SqlCondition) SqlUpdateExp
	LeftJoin(table SqlTable, on func(join SqlTable, others SqlTables) SqlCondition) SqlUpdateExp
	RightJoin(table SqlTable, on func(join SqlTable, others SqlTables) SqlCondition) SqlUpdateExp
	CrossJoin(table SqlTable, on func(join SqlTable, others SqlTables) SqlCondition) SqlUpdateExp
}
