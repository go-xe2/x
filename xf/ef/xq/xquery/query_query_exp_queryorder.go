package xquery

import . "github.com/go-xe2/x/xf/ef/xqi"

func (qe *tQueryExp) Order(fields func(tables SqlTables) []SqlOrderField) Query {
	qe.checkMainTableSet()
	qe.com.Order(fields)
	return qe
}
