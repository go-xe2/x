package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (sf *TSqlField) Cast(asType DbType) SqlField {
	return SqlFunCast(sf.This(), asType)
}
