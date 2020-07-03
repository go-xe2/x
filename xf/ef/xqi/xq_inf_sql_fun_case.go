package xqi

type SqlCase interface {
	SqlCompiler
	SqlCaseThenElse
	When(when, then interface{}) SqlCase
	ElseEnd(val interface{}) SqlField
	End() SqlField
	FunId() SqlFunId
}

type SqlCaseThenElse interface {
	// when thenValue else elseValue end
	ThenElse(thenValue, elseValue interface{}) SqlField
}
