package xqcomm

import (
	"github.com/go-xe2/x/xf/ef/xqi"
)

func (exp *tSqlInsertExp) Table(table xqi.SqlTable) xqi.SqlInsertExp {
	exp.insertTable = table
	return exp
}

func (exp *tSqlInsertExp) Values(fields ...xqi.FieldValue) xqi.SqlInsertExp {
	if len(fields) == 0 {
		return exp
	}
	exp.fieldExps = fields
	return exp
}

func (exp *tSqlInsertExp) From(fromTable xqi.SqlTable) xqi.SqlInsertExp {
	exp.fromTable = fromTable
	return exp
}

func (exp *tSqlInsertExp) GetTable() xqi.SqlTable {
	return exp.insertTable
}

func (exp *tSqlInsertExp) GetFromTable() xqi.SqlTable {
	return exp.fromTable
}

func (exp *tSqlInsertExp) GetFields() []xqi.FieldValue {
	return exp.fieldExps
}
