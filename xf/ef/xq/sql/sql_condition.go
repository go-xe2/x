package sql

import (
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

func Where(items ...xqi.SqlLogic) xqi.SqlCondition {
	return And(items...)
}

func On(items ...xqi.SqlLogic) xqi.SqlCondition {
	return And(items...)
}

func And(items ...xqi.SqlLogic) xqi.SqlCondition {
	return xqcomm.NewSqlCondition(xqi.SqlConditionAndLogic).And(items...)
}

func Or(items ...xqi.SqlLogic) xqi.SqlCondition {
	var result xqi.SqlCondition = xqcomm.NewSqlCondition(xqi.SqlConditionOrLogic)
	for _, v := range items {
		result = result.Or(v)
	}
	return result
	//return xqcomm.NewSqlCondition(xqi.SqlConditionOrLogic).And(items...)
}

func Xor(items ...xqi.SqlLogic) xqi.SqlCondition {
	var result xqi.SqlCondition = xqcomm.NewSqlCondition(xqi.SqlConditionOrLogic)
	for _, v := range items {
		result = result.Xor(v)
	}
	return result
	//return xqcomm.NewSqlCondition(xqi.SqlConditionXorLogic).And(items...)
}

func Eq(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareEQType)
}

func Neq(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareNEQType)
}

func Gt(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareEQType)
}

func Gte(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareGTEType)
}

func Lt(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareLTType)
}

func Lte(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareLTEType)
}

func Like(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareLKType)
}

func NotLike(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareNLKType)
}

func In(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareINType)
}

func NotIn(exp1, exp2 interface{}) xqi.SqlConditionItem {
	return xqcomm.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareNINType)
}
