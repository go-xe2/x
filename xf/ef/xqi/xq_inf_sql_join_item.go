package xqi

type SqlJoinItem interface {
	// 联查方式
	JoinType() SqlJoinType
	// 连接的表
	JoinTable() SqlTable
	// 连接条件表达式
	LazyConditionFn() LazyGetJoinConditionFun
}

type SqlJoin interface {
	// 联查方式
	JoinType() SqlJoinType
	// 连接的表
	JoinTable() SqlTable
	OnCondition() SqlCondition
}
