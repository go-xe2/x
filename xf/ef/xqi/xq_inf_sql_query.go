package xqi

type SqlQuery interface {
	SqlCompiler
	GetFromTable() SqlTable
	GetQueryFields() []SqlField
	GetJoins() []SqlJoin
	GetWhere() SqlCondition
	GetTables() SqlTables
	GetGroupFields() []SqlField
	GetHaving() SqlCondition
	GetOrderFields() []SqlOrderField
	GetLimitRows() int
	GetLimitOffset() int

	Fields(fields func(tables SqlTables) []SqlField) SqlQuery
	Where(condition func(tables SqlTables) SqlCondition) SqlQuery
	Join(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery
	LeftJoin(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery
	RightJoin(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery
	CrossJoin(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery
	Group(fields func(tables SqlTables) []SqlField) SqlQuery
	Having(condition func(tables SqlTables) SqlCondition) SqlQuery
	Order(fields func(tables SqlTables) []SqlOrderField) SqlQuery
	Limit(rows, offset int) SqlQuery
	Page(size, index int) SqlQuery

	Alias(alias string) SqlTable
}
