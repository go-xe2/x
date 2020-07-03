package xqi

type SqlStaticExpr interface {
	SqlCompiler
	Val() interface{}
}

type SqlVarExpr interface {
	SqlCompiler
	Val() interface{}
	SetVal(v interface{}) SqlVarExpr
}
