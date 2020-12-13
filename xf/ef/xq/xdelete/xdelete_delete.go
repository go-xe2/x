package xdelete

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tDelete struct {
	db    xqi.Database
	table xqi.SqlTable
	where []xqi.SqlCondition
}

var _ SqlDelete = (*tDelete)(nil)

func Delete(table xqi.SqlTable, db xqi.Database) SqlDelete {
	return &tDelete{
		db:    db,
		table: table,
	}
}

func (del *tDelete) Where(where ...xqi.SqlCondition) SqlDeleteExecute {
	del.where = where
	return del
}

func (del *tDelete) Execute() (int, error) {
	if del.db == nil || del.table == nil {
		return 0, exception.NewText("未设置数据库或待删除的表实体")
	}
	cxt := xqcomm.NewSqlCompileContext()
	cxt.SetDatabase(del.db)
	builder := xdriver.GetSqlBuilderByName(cxt.Driver())

	//szTable := ""
	cxt.PushState(xqi.SCPBuildDeleteTableState)
	if tk := del.table.Compile(builder, cxt, false); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
		return 0, exception.NewText("表达式中删除数据的表实体无效")
	} else {
		_ = tk.Val()
	}
	defer cxt.PopState()
	tableName := builder.QuotesName(cxt.TablePrefix() + del.table.TableName())

	tableAlias := builder.QuotesName(del.table.TableAlias())

	sql := ""
	if tableAlias != "" {
		// mysql delete语句别名处理, delete cc from tableName cc
		sql = fmt.Sprintf("DELETE %s FROM %s AS %s", tableAlias, tableName, tableAlias)
	} else {
		sql = fmt.Sprintf("DELETE FROM %s", tableName)
	}
	vars := make([]interface{}, 0)
	szWhere := ""

	if len(del.where) > 0 {
		cxt.PushState(xqi.SCPBuildDeleteWhereState)
		tk := del.where[0].Compile(builder, cxt, false)
		if tk != nil && tk.TType() != xqi.SqlEmptyTokenType {
			szWhere = tk.Val()
			vars = builder.MakeQueryParams(cxt.Params())
		}
		cxt.PopState()
	}
	if szWhere != "" {
		sql += " WHERE " + szWhere
	}
	if n, err := del.db.Execute(sql, vars...); err != nil {
		return 0, err
	} else {
		return int(n), nil
	}
}
