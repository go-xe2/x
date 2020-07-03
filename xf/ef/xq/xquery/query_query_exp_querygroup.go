package xquery

import . "github.com/go-xe2/x/xf/ef/xqi"

func (qe *tQueryExp) Group(fields func(tables SqlTables) []SqlField) Query {
	qe.checkMainTableSet()
	qe.com.Group(fields)
	return qe
}
