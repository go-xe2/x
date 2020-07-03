package xqcomm

import "github.com/go-xe2/x/xf/ef/xqi"

// 字段一元运算操作

// 字段自增
func (stf *TSqlTableField) Inc(step ...int) xqi.FieldValue {
	n := 1
	if len(step) > 0 {
		n = step[0]
	}
	return NewFieldValue(stf.This().(xqi.SqlTableField), stf.Plus(n))
}

// 字段自减
func (stf *TSqlTableField) Dec(step ...int) xqi.FieldValue {
	n := 1
	if len(step) > 0 {
		n = step[0]
	}
	return NewFieldValue(stf.This().(xqi.SqlTableField), stf.Sub(n))
}

// 字段自乘运算
func (stf *TSqlTableField) UnaryMul(val interface{}) xqi.FieldValue {
	return NewFieldValue(stf.This().(xqi.SqlTableField), stf.Mul(val))
}

// 字段自除运算
func (stf *TSqlTableField) UnaryDiv(val interface{}) xqi.FieldValue {
	return NewFieldValue(stf.This().(xqi.SqlTableField), stf.Div(val))
}
