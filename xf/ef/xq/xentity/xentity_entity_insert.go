package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tEntityInsert struct {
	entity *TEntity
	expr   xqi.SqlInsertExp
}

var _ xqi.EntityInsert = (*tEntityInsert)(nil)

func newEntityInsert(entity *TEntity, values ...xqi.FieldValue) xqi.EntityInsert {
	return &tEntityInsert{
		entity: entity,
		expr:   xqcomm.NewSqlInsertExp(entity).Values(values...),
	}
}

func (eni *tEntityInsert) Execute() (int, error) {

	// 如果实体实现了插入监听接口，则先通过监听接口检查
	observer, obOk := eni.entity.This().(xqi.EntityInsertObserver)
	if obOk {
		fieldList := eni.expr.GetFields()
		if !observer.BeforeInsert(fieldList) {
			return 0, nil
		}
	}
	db := xdatabase.DB(eni.entity.dbName...)
	cxt := xqcomm.NewSqlCompileContext()
	cxt.SetDatabase(db)
	builder := xdriver.GetSqlBuilderByName(cxt.Driver())
	tk := eni.expr.Compile(builder, cxt, false)
	if tk == nil || tk.TType() != xqi.SqlInsertTokenType {
		return 0, exception.NewText("表达式错误或不是插入数据表达式")
	}
	vars := builder.MakeQueryParams(cxt.Params())
	if n, err := db.Execute(tk.Val(), vars...); err != nil {
		return 0, err
	} else {
		eni.entity.lastId = db.LastInsertId()
		return int(n), nil
	}
}

func (eni *tEntityInsert) From(table xqi.SqlTable) xqi.EntityInsertExecute {
	eni.expr = eni.expr.From(table)
	return eni
}
