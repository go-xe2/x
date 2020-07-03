package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (stf *TSqlTableField) Substr(from, l int) SqlField {
	return SqlFunSubstring(stf.This(), from, l)
}

func (stf *TSqlTableField) Concat(val interface{}) SqlField {
	return SqlFunStrConcat(stf.This(), val)
}
