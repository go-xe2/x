package xquery

import (
	xqcomm2 "github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

func (qe *tQueryExp) Having(exp func(having SqlCondition, tables SqlTables) SqlCondition) Query {
	qe.checkMainTableSet()
	qe.com.Having(func(tables SqlTables) SqlCondition {
		lc := xqcomm2.NewSqlCondition()
		return exp(lc, tables)
	})
	return qe
}
