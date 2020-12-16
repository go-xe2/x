package xq

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
	"github.com/go-xe2/x/xf/ef/xq/xdelete"
	"github.com/go-xe2/x/xf/ef/xq/xinsert"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	xquery2 "github.com/go-xe2/x/xf/ef/xq/xquery"
	"github.com/go-xe2/x/xf/ef/xq/xupdate"
	"github.com/go-xe2/x/xf/ef/xqi"
)

// 通过名称获取指定数据库
// @param dbName 配置的数据库名称
func Database(dbName ...string) xqi.Database {
	return xdatabase.DB(dbName...)
}

// 在指定数据库上查询数据
// @param dbName 为空时在默认数据库上查询数据
func Query(dbName ...string) xqi.QuerySelect {
	db := Database(dbName...)
	if db == nil {
		panic(exception.NewText("未设置默认数据库"))
	}
	return xquery2.NewQueryExp(db)
}

// 在指定数据库上查询数据
func QueryByDb(db xqi.Database) xqi.QuerySelect {
	return xquery2.NewQueryExp(db)
}

// 在指定名称的数据库上向指定表插入数据
func Insert(table xqi.SqlTable, dbName ...string) xinsert.SqlInsert {
	db := Database(dbName...)
	return xinsert.Insert(db, table)
}

// 在指定的数据库上向指定表插入数据
func InsertByDb(table xqi.SqlTable, db xqi.Database) xinsert.SqlInsert {
	return xinsert.Insert(db, table)
}

// 在指定名称的数据库上更新数据库表
func Update(table xqi.SqlTable, dbName ...string) xupdate.SqlUpdate {
	db := Database(dbName...)
	return xupdate.Update(db, table)
}

// 在指定的数据库上更新数据库表
func UpdateByDb(table xqi.SqlTable, db xqi.Database) xupdate.SqlUpdate {
	return xupdate.Update(db, table)
}

func Delete(table xqi.SqlTable, dbName ...string) xdelete.SqlDelete {
	db := Database(dbName...)
	return xdelete.Delete(table, db)
}

func DeleteByDb(table xqi.SqlTable, db xqi.Database) xdelete.SqlDelete {
	return xdelete.Delete(table, db)
}

func ExprField(sqlExpress string, alias string) xqi.SqlField {
	return xqcomm.NewExpressField(sqlExpress, alias)
}

func ExprTable(sqlExpress string, alias string) xqi.SqlTable {
	return xqcomm.NewExpressQuery(sqlExpress, alias)
}
