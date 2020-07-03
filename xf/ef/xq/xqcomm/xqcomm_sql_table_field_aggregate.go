package xqcomm

import "github.com/go-xe2/x/xf/ef/xqi"

func (stf *TSqlTableField) Count() xqi.SqlField {
	return SqlFunCount(stf.This().(xqi.SqlTableField))
}

func (stf *TSqlTableField) Max() xqi.SqlField {
	return SqlFunMax(stf.This().(xqi.SqlTableField))
}

func (stf *TSqlTableField) Min() xqi.SqlField {
	return SqlFunMin(stf.This().(xqi.SqlTableField))
}

func (stf *TSqlTableField) Avg() xqi.SqlField {
	return SqlFunAvg(stf.This().(xqi.SqlTableField))
}

func (stf *TSqlTableField) Sum() xqi.SqlField {
	return SqlFunSum(stf.This().(xqi.SqlTableField))
}
