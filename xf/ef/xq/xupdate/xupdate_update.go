package xupdate

import (
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tUpdate struct {
	expr xqi.SqlUpdateExp
	db   xqi.Database
}

var _ SqlUpdate = (*tUpdate)(nil)

func Update(db xqi.Database, table xqi.SqlTable) SqlUpdate {
	return &tUpdate{
		db:   db,
		expr: xqcomm.NewSqlUpdateExp(table),
	}
}

func (upd *tUpdate) Set(values ...xqi.FieldValue) SqlUpdateSet {
	upd.expr = upd.expr.Set(values...)
	return upd
}

func (upd *tUpdate) Join(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin {
	upd.expr = upd.expr.Join(xqi.SqlInnerJoinType, table, on)
	return upd
}

func (upd *tUpdate) LeftJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin {
	upd.expr = upd.expr.Join(xqi.SqlLeftJoinType, table, on)
	return upd
}

func (upd *tUpdate) RightJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin {
	upd.expr = upd.expr.Join(xqi.SqlRightJoinType, table, on)
	return upd
}

func (upd *tUpdate) CrossJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) SqlUpdateJoin {
	upd.expr = upd.expr.Join(xqi.SqlCrossJoinType, table, on)
	return upd
}

func (upd *tUpdate) Where(where xqi.SqlCondition) SqlUpdateExecute {
	upd.expr = upd.expr.Where(where)
	return upd
}

func (upd *tUpdate) Execute() (int, error) {
	cxt := xqcomm.NewSqlCompileContext()
	cxt.SetDatabase(upd.db)

	builder := xdriver.GetSqlBuilderByName(cxt.Driver())
	tk := upd.expr.Compile(builder, cxt, false)
	if tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
		return 0, nil
	}
	vars := builder.MakeQueryParams(cxt.Params())
	if n, err := upd.db.Execute(tk.Val(), vars...); err != nil {
		return 0, err
	} else {
		return int(n), nil
	}
}

func (upd *tUpdate) DB() xqi.Database {
	return upd.db
}
