package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

// 字段算术运算

func (sf *TSqlField) Plus(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(sf.This(), val, SqlAriPlusType), "")
}

func (sf *TSqlField) Sub(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(sf.This(), val, SqlAriSubType), "")
}

func (sf *TSqlField) Mul(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(sf.This(), val, SqlAriMulType), "")
}

func (sf *TSqlField) Div(val interface{}) SqlField {
	return NewSqlField(NewSqlAriExp(sf.This(), val, SqlAriDivType), "")
}
