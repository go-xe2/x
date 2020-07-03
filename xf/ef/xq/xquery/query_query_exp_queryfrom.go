package xquery

import (
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

func (qe *tQueryExp) From(table SqlTable) Query {
	qe.com = xqcomm.NewSqlQuery(table)
	qe.com.Fields(qe.fields)
	return qe
}
