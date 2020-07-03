package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (stf *TSqlTableField) Eq(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareEQType)
}

func (stf *TSqlTableField) Neq(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareNEQType)
}

func (stf *TSqlTableField) Gt(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareGTType)
}

func (stf *TSqlTableField) Gte(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareGTEType)
}

func (stf *TSqlTableField) Lt(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareLTType)
}

func (stf *TSqlTableField) Lte(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareLTEType)
}

func (stf *TSqlTableField) In(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareINType)
}

func (stf *TSqlTableField) NotIn(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareNINType)
}

func (stf *TSqlTableField) Like(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareLKType)
}

func (stf *TSqlTableField) NotLike(v interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.This(), v, SqlCompareNLKType)
}

func (stf *TSqlTableField) SubEq(from, l int, eqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Substr(from, l), eqVal, SqlCompareEQType)
}

func (stf *TSqlTableField) SubNeq(from, l int, neqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Substr(from, l), neqVal, SqlCompareNEQType)
}

func (stf *TSqlTableField) SubIn(from, l int, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Substr(from, l), arr, SqlCompareINType)
}

func (stf *TSqlTableField) SubNotIn(from, l int, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Substr(from, l), arr, SqlCompareNINType)
}

func (stf *TSqlTableField) SubLike(from, l int, val interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Substr(from, l), val, SqlCompareLKType)
}

func (stf *TSqlTableField) SubNotLike(from, l int, val interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Substr(from, l), val, SqlCompareNLKType)
}

func (stf *TSqlTableField) Not() SqlConditionItem {
	return NewSqlConditionItem(stf.This(), nil, SqlCompareNTType)
}

// 字符串字段拼接val后等于eqVal比较
func (stf *TSqlTableField) ConcatEq(catVal interface{}, eqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Concat(catVal), eqVal, SqlCompareEQType)
}

func (stf *TSqlTableField) ConcatNeq(catVal interface{}, neqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Concat(catVal), neqVal, SqlCompareNEQType)
}

func (stf *TSqlTableField) ConcatIn(catVal interface{}, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Concat(catVal), arr, SqlCompareINType)
}

func (stf *TSqlTableField) ConcatNotIn(catVal interface{}, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Concat(catVal), arr, SqlCompareNINType)
}

func (stf *TSqlTableField) ConcatLike(catVal interface{}, likeVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Concat(catVal), likeVal, SqlCompareLKType)
}

func (stf *TSqlTableField) ConcatNotLike(catVal interface{}, notLikeVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(stf.Concat(catVal), notLikeVal, SqlCompareNLKType)
}
