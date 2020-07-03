package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

// 排序
func (sf *TSqlField) Desc() SqlOrderField {
	return NewSqlOrderField(sf.This().(SqlField), SqlOrderDescDirect)
}

func (sf *TSqlField) Asc() SqlOrderField {
	return NewSqlOrderField(sf.This().(SqlField), SqlOrderAscDirect)
}
