package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (stf *TSqlTableField) Substr(from, l int) SqlField {
	return SqlFunSubstring(stf.This(), from, l)
}

func (stf *TSqlTableField) Concat(val interface{}, more ...interface{}) SqlField {
	return SqlFunStrConcat(stf.This(), val, more...)
}

func (stf *TSqlTableField) FromBase64() SqlField {
	return SqlFunFromBase64(stf.This())
}

func (stf *TSqlTableField) ToBase64() SqlField {
	return SqlFunToBase64(stf.This())
}
