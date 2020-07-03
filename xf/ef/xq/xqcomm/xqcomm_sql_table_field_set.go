package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (stf *TSqlTableField) Set(val interface{}) FieldValue {
	return NewFieldValue(stf.This().(SqlTableField), val)
}
