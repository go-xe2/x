package xqi

type SqlLogic interface {
	SqlCompiler
	// 左表达式
	LExp() interface{}
	// 右表达式
	RExp() interface{}
	// 左右表达式之间关联关系
	Logic() SqlConditionLogic
}

type SqlConditionItem interface {
	SqlLogic
	// 比较符号
	Comparer() SqlCompareType
}

type SqlCondition interface {
	SqlLogic
	Items() []SqlLogic
	And(exp ...SqlLogic) SqlCondition
	Or(exp ...SqlLogic) SqlCondition
	Xor(exp ...SqlLogic) SqlCondition
}
