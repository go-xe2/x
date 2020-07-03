package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

// 子字符串
func (sf *TSqlField) Substr(from, l int) SqlField {
	return SqlFunSubstring(sf.This(), from, l)
}

// 字符串连接
func (sf *TSqlField) Concat(val interface{}) SqlField {
	return SqlFunStrConcat(sf.This(), val)
}

// = Substr(from,l).Eq(eqVal)
func (sf *TSqlField) SubEq(from, l int, eqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Substr(from, l), eqVal, SqlCompareEQType)
}

// = Substr(from,l).Neq(neqVal)
func (sf *TSqlField) SubNeq(from, l int, neqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Substr(from, l), neqVal, SqlCompareNEQType)
}

// = Substr(from,l).In(arr)
func (sf *TSqlField) SubIn(from, l int, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Substr(from, l), arr, SqlCompareINType)
}

// = Substr(from,l).NotIn(arr)
func (sf *TSqlField) SubNotIn(from, l int, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Substr(from, l), arr, SqlCompareNINType)
}

// = Substr(from, l).Like(val)
func (sf *TSqlField) SubLike(from, l int, val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Substr(from, l), val, SqlCompareLKType)
}

// = Substr(from,l).NotLike(val)
func (sf *TSqlField) SubNotLike(from, l int, val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Substr(from, l), val, SqlCompareNLKType)
}

// = not field
func (sf *TSqlField) Not() SqlConditionItem {
	return NewSqlConditionItem(sf.This(), nil, SqlCompareNTType)
}

// = Concat(catVal).Eq(eqVal)
func (sf *TSqlField) ConcatEq(catVal interface{}, eqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Concat(catVal), eqVal, SqlCompareEQType)
}

// = Concat(catVal).Neq(neqVal)
func (sf *TSqlField) ConcatNeq(catVal interface{}, neqVal interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Concat(catVal), neqVal, SqlCompareNEQType)
}

// = Concat(catVal).In(arr)
func (sf *TSqlField) ConcatIn(catVal interface{}, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Concat(catVal), arr, SqlCompareINType)
}

// = Concat(catVal).NotIn(arr)
func (sf *TSqlField) ConcatNotIn(catVal interface{}, arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Concat(catVal), arr, SqlCompareNINType)
}

// = Concat(catVal).Like(arr)
func (sf *TSqlField) ConcatLike(catVal interface{}, val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Concat(catVal), val, SqlCompareLKType)
}

// = Concat(catVal).NotLike(arr)
func (sf *TSqlField) ConcatNotLike(catVal interface{}, val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.Concat(catVal), val, SqlCompareNLKType)
}
