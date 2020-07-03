package xqcomm

import "github.com/go-xe2/x/xf/ef/xqi"

var _ xqi.SqlTableFields = (*TSqlTable)(nil)

func (st *TSqlTable) AddField(field xqi.SqlField) xqi.SqlTable {
	return st
}
