package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (sf *TSqlField) Count() SqlField {
	return SqlFunCount(sf.This().(SqlField))
}

func (sf *TSqlField) Max() SqlField {
	return SqlFunMax(sf.This().(SqlField))
}

func (sf *TSqlField) Min() SqlField {
	return SqlFunMin(sf.This().(SqlField))
}

func (sf *TSqlField) Avg() SqlField {
	return SqlFunAvg(sf.This().(SqlField))
}

func (sf *TSqlField) Sum() SqlField {
	return SqlFunSum(sf.This().(SqlField))
}
