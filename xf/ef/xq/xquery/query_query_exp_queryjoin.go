package xquery

import (
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

func (qe *tQueryExp) Join(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query {
	qe.checkMainTableSet()
	qe.com.Join(table, func(joinTable SqlTable, tables SqlTables) SqlCondition {
		onLc := xqcomm.NewSqlCondition()
		return on(joinTable, tables, onLc)
	})
	return qe
}

func (qe *tQueryExp) LeftJoin(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query {
	qe.checkMainTableSet()
	var fn = func(joinTable SqlTable, tables SqlTables) SqlCondition {
		onLc := xqcomm.NewSqlCondition()
		return on(joinTable, tables, onLc)
	}
	qe.com.LeftJoin(table, fn)
	return qe
}

func (qe *tQueryExp) RightJoin(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query {
	qe.checkMainTableSet()
	qe.com.RightJoin(table, func(joinTable SqlTable, tables SqlTables) SqlCondition {
		onLc := xqcomm.NewSqlCondition()
		return on(joinTable, tables, onLc)
	})
	return qe
}

func (qe *tQueryExp) CrossJoin(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query {
	qe.checkMainTableSet()
	qe.com.CrossJoin(table, func(joinTable SqlTable, tables SqlTables) SqlCondition {
		onLc := xqcomm.NewSqlCondition()
		return on(joinTable, tables, onLc)
	})
	return qe
}
