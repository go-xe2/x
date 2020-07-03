package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

// 转换字段数据类型
func (stf *TSqlTableField) Cast(asType DbType) SqlField {
	return SqlFunCast(stf.This(), asType)
}
