package xinsert

import (
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tInsert struct {
	db   xqi.Database
	expr xqi.SqlInsertExp
}

var _ SqlInsert = (*tInsert)(nil)

func Insert(db xqi.Database, table xqi.SqlTable) SqlInsert {
	return &tInsert{
		db:   db,
		expr: xqcomm.NewSqlInsertExp(table),
	}
}

func (ins *tInsert) Values(values ...xqi.FieldValue) SqlInsertFrom {
	ins.expr = ins.expr.Values(values...)
	return ins
}

func (ins *tInsert) From(table xqi.SqlTable) SqlInsertExecute {
	ins.expr = ins.expr.From(table)
	return ins
}

func (ins *tInsert) Execute() (int, error) {
	cxt := xqcomm.NewSqlCompileContext()
	cxt.SetDatabase(ins.db)

	builder := xdriver.GetSqlBuilderByName(cxt.Driver())
	tk := ins.expr.Compile(builder, cxt, false)
	if tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
		return 0, nil
	}
	vars := builder.MakeQueryParams(cxt.Params())
	if n, err := ins.db.Execute(tk.Val(), vars...); err != nil {
		return 0, err
	} else {
		return int(n), nil
	}
}

func (ins *tInsert) DB() xqi.Database {
	return ins.db
}
