package xqi

type SqlOrderField interface {
	SqlCompiler
	Field() SqlField
	OrderDir() SqlOrderDirect
}
