package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tUpdate struct {
	entity *TEntity
	expr   xqi.SqlUpdateExp
}

var _ xqi.EntityUpdate = (*tUpdate)(nil)

func newUpdate(entity *TEntity, fields ...xqi.FieldValue) xqi.EntityUpdate {
	return &tUpdate{
		entity: entity,
		expr:   xqcomm.NewSqlUpdateExp(entity).Set(fields...),
	}
}

func (up *tUpdate) Join(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) xqi.EntityUpdate {
	up.expr = up.expr.Join(xqi.SqlInnerJoinType, table, on)
	return up
}

func (up *tUpdate) LeftJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) xqi.EntityUpdate {
	up.expr = up.expr.Join(xqi.SqlLeftJoinType, table, on)
	return up
}

func (up *tUpdate) RightJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) xqi.EntityUpdate {
	up.expr = up.expr.Join(xqi.SqlRightJoinType, table, on)
	return up
}

func (up *tUpdate) CrossJoin(table xqi.SqlTable, on func(joinEnt xqi.SqlTable, leftTables xqi.SqlTables) xqi.SqlCondition) xqi.EntityUpdate {
	up.expr = up.expr.Join(xqi.SqlCrossJoinType, table, on)
	return up
}

func (up *tUpdate) Where(where ...xqi.SqlCondition) xqi.EntityUpdateExecute {
	if len(where) > 0 {
		up.expr = up.expr.Where(where[0])
	}
	return up
}

func (up *tUpdate) Execute() (int, error) {
	upObserver, ok := up.entity.This().(xqi.EntityUpdateObserver)
	if ok {
		fieldList := up.expr.GetFields()
		if !upObserver.BeforeUpdate(fieldList) {
			return 0, nil
		}
	}

	db := xdatabase.DB(up.entity.dbName...)
	cxt := xqcomm.NewSqlCompileContext()
	cxt.SetDatabase(db)
	builder := xdriver.GetSqlBuilderByName(cxt.Driver())
	tk := up.expr.Compile(builder, cxt, false)
	if tk == nil || tk.TType() != xqi.SqlUpdateTokenType {
		return 0, exception.NewText("表达式错误或不是更新数据表达式")
	}
	vars := builder.MakeQueryParams(cxt.Params())
	if n, err := db.Execute(tk.Val(), vars...); err != nil {
		return 0, err
	} else {
		return int(n), nil
	}
}
