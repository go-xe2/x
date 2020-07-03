package xqi

type SqlTables interface {
	Table(name string) SqlTable
	HasTable(name string) bool
	All() []SqlTable
	Add(table SqlTable) SqlTables
	AddTable(name string, alias ...string) SqlTable
	Clear() SqlTables
	This() interface{}
	String() string
	Count() int
}
