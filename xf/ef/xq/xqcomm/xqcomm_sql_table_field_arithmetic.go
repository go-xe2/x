package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

// 字段算术运算表达式

// 加法运算 field + val
func (stf *TSqlTableField) Plus(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(stf.This(), val, SqlAriPlusType), "")
}

// 减法运算 field - val
func (stf *TSqlTableField) Sub(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(stf.This(), val, SqlAriSubType), "")
}

// 乘法运算 field * val
func (stf *TSqlTableField) Mul(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(stf.This(), val, SqlAriMulType), "")
}

// 除法运算 field / val
func (stf *TSqlTableField) Div(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(stf.This(), val, SqlAriDivType), "")
}
