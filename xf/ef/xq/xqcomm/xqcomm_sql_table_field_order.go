package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (stf *TSqlTableField) Desc() SqlOrderField {
	return NewSqlOrderField(stf.This().(SqlTableField), SqlOrderDescDirect)
}

func (stf *TSqlTableField) Asc() SqlOrderField {
	return NewSqlOrderField(stf.This().(SqlTableField), SqlOrderAscDirect)
}
