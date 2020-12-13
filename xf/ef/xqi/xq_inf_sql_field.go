package xqi

type SqlField interface {
	SqlCompiler
	// 表达式
	Exp() interface{}
	// 别名
	AliasName() string
	Alias(alias string) SqlField
	Eq(val interface{}) SqlConditionItem
	Neq(val interface{}) SqlConditionItem
	Gt(val interface{}) SqlConditionItem
	Gte(val interface{}) SqlConditionItem
	Lt(val interface{}) SqlConditionItem
	Lte(val interface{}) SqlConditionItem
	In(arr interface{}) SqlConditionItem
	NotIn(arr interface{}) SqlConditionItem
	Like(val interface{}) SqlConditionItem
	NotLike(val interface{}) SqlConditionItem

	// ===== 算术运算 =======//
	// ======= 算术运算 ======= //
	// +
	Plus(val interface{}) SqlField
	// -
	Sub(val interface{}) SqlField
	// *
	Mul(val interface{}) SqlField
	// /
	Div(val interface{}) SqlField

	// ==== case when then 函数 ==//
	Case() SqlCase
	// case when field == v then else end
	CaseEq(v interface{}) SqlCaseThenElse
	// case when field != v then else end
	CaseNeq(v interface{}) SqlCaseThenElse
	// case when field > v then else end
	CaseGt(v interface{}) SqlCaseThenElse
	// case when field >= v then else end
	CaseGte(v interface{}) SqlCaseThenElse
	// case when field < v then else end
	CaseLt(v interface{}) SqlCaseThenElse
	// case when field <= v then else end
	CaseLte(v interface{}) SqlCaseThenElse

	// 数据类型转换
	Cast(asType DbType) SqlField

	// ========== 聚合函数 ===== //
	Count() SqlField
	Max() SqlField
	Min() SqlField
	Avg() SqlField
	Sum() SqlField
	String() string

	// 子字符串
	Substr(from, l int) SqlField
	FromBase64() SqlField
	ToBase64() SqlField
	// 字符串连接
	Concat(val interface{}, more ...interface{}) SqlField
	// = Substr(from,l).Eq(eqVal)
	SubEq(from, l int, eqVal interface{}) SqlConditionItem
	// = Substr(from,l).Neq(neqVal)
	SubNeq(from, l int, neqVal interface{}) SqlConditionItem
	// = Substr(from,l).In(arr)
	SubIn(from, l int, arr interface{}) SqlConditionItem
	// = Substr(from,l).NotIn(arr)
	SubNotIn(from, l int, arr interface{}) SqlConditionItem
	// = Substr(from, l).Like(val)
	SubLike(from, l int, val interface{}) SqlConditionItem
	// = Substr(from,l).NotLike(val)
	SubNotLike(from, l int, val interface{}) SqlConditionItem
	// = not field
	Not() SqlConditionItem
	// = Concat(catVal).Eq(eqVal)
	ConcatEq(catVal interface{}, eqVal interface{}) SqlConditionItem
	// = Concat(catVal).Neq(neqVal)
	ConcatNeq(catVal interface{}, neqVal interface{}) SqlConditionItem
	// = Concat(catVal).In(arr)
	ConcatIn(catVal interface{}, arr interface{}) SqlConditionItem
	// = Concat(catVal).NotIn(arr)
	ConcatNotIn(catVal interface{}, arr interface{}) SqlConditionItem
	// = Concat(catVal).Like(val)
	ConcatLike(catVal interface{}, val interface{}) SqlConditionItem
	// = Concat(catVal).NotLike(val)
	ConcatNotLike(catVal interface{}, val interface{}) SqlConditionItem

	// 排序
	Desc() SqlOrderField
	Asc() SqlOrderField
}
