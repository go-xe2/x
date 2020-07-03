package xqi

type SqlFun interface {
	SqlCompiler
	FunId() SqlFunId
}
