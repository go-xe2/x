package xq

import (
	xqcomm2 "github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

func Eq(exp1, exp2 interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareEQType)
}

func Neq(exp1, exp2 interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareNEQType)
}

func Gt(exp1, exp2 interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareGTType)
}

func Gte(exp1, exp2 interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareGTEType)
}

func Lt(exp1, exp2 interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareLTType)
}

func Lte(exp1, exp2 interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp1, exp2, xqi.SqlCompareLTEType)
}

func In(exp, arr interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp, arr, xqi.SqlCompareINType)
}

func NotIn(exp, arr interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp, arr, xqi.SqlCompareNINType)
}

func Like(exp, val interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp, val, xqi.SqlCompareLKType)
}

func NotLike(exp, val interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp, val, xqi.SqlCompareNLKType)
}

func Not(exp interface{}) SqlConditionItem {
	return xqcomm2.NewSqlConditionItem(exp, nil, xqi.SqlCompareNTType)
}
