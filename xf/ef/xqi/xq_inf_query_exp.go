package xqi

import (
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/encoding/xxml"
)

type QueryInfo interface {
	// 获取查询的表
	GetTable() SqlTable
	// 获取查询的字段列表
	GetSelectFields() []SqlField
	// 获取where条件表达式
	GetWhere() SqlCondition
	// 获取having条件表达式
	GetHaving() SqlCondition
	// 获取排序字段列表
	GetOrders() []SqlOrderField
	// 获取分组字段列表
	GetGroups() []SqlField
	// 获取联查字段列表
	GetJoins() []SqlJoin
	// 获取查询表达式中所用到的表，包括join中的表
	UseTables() []SqlTable
}

type Query interface {
	SqlCompiler
	QueryJoin
	QueryWhere
	QueryGroup
	QueryOrder
	QueryHaving
	QueryLimit
	Info() QueryInfo
	// 查询的数据库
	DB() Database
	// 表达式取别名转换成table查询语句以供外层查询
	Alias(alias string) SqlTable
	Xml() (xxml.XmlStr, error)               // 获取xml结果
	Json() (xjson.JsonStr, error)            // 获取json结果
	Rows() ([]map[string]interface{}, error) // 获取map
	Slices() ([][]interface{}, error)
	Dataset() (Dataset, error)
	Visitor(visitor func(row int, values ...interface{}) (interface{}, bool)) ([]interface{}, error)
	Bind(binder DbQueryBinder) (interface{}, error)
	Sql() (sql string, vars []interface{})
}

type QuerySelect interface {
	// 查询指定字段
	Fields(fields func(tables SqlTables) []SqlField) QueryFrom
	// 查询所有字段
	All() QueryFrom
}

type QueryFrom interface {
	From(table SqlTable) Query
}

type QueryWhere interface {
	Where(exp func(where SqlCondition, tables SqlTables) SqlCondition) Query
}

type QueryJoin interface {
	Join(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query
	LeftJoin(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query
	RightJoin(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query
	CrossJoin(table SqlTable, on func(joinTable SqlTable, otherTables SqlTables, on SqlCondition) SqlCondition) Query
}

type QueryGroup interface {
	Group(fields func(tables SqlTables) []SqlField) Query
}

type QueryHaving interface {
	Having(exp func(having SqlCondition, tables SqlTables) SqlCondition) Query
}

type QueryOrder interface {
	Order(fields func(tables SqlTables) []SqlOrderField) Query
}

type QueryLimit interface {
	Limit(rows int, offset ...int) Query
	Page(size int, index ...int) Query
}
