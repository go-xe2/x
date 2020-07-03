package xqi

type SqlInsertExpInfo interface {
	GetTable() SqlTable
	GetFromTable() SqlTable
	GetFields() []FieldValue
}

type SqlInsertExp interface {
	SqlCompiler
	SqlInsertExpInfo

	Table(table SqlTable) SqlInsertExp
	From(fromTable SqlTable) SqlInsertExp
	Values(fields ...FieldValue) SqlInsertExp
}
