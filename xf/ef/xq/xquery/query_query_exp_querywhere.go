package xquery

import (
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

func (qe *tQueryExp) Where(exp func(where SqlCondition, tables SqlTables) SqlCondition) Query {
	qe.checkMainTableSet()
	qe.com.Where(func(tables SqlTables) SqlCondition {
		lc := xqcomm.NewSqlCondition()
		result := exp(lc, tables)
		return result
	})
	return qe
}
