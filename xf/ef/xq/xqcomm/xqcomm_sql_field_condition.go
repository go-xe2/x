package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (sf *TSqlField) Eq(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareEQType)
}

func (sf *TSqlField) Neq(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareNEQType)
}

func (sf *TSqlField) Gt(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareGTType)
}

func (sf *TSqlField) Gte(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareGTEType)
}

func (sf *TSqlField) Lt(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareLTType)
}

func (sf *TSqlField) Lte(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareLTEType)
}

func (sf *TSqlField) In(arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), arr, SqlCompareINType)
}

func (sf *TSqlField) NotIn(arr interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), arr, SqlCompareNINType)
}

func (sf *TSqlField) Like(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareLKType)
}

func (sf *TSqlField) NotLike(val interface{}) SqlConditionItem {
	return NewSqlConditionItem(sf.This(), val, SqlCompareNLKType)
}
