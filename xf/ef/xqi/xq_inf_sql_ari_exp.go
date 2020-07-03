package xqi

// 算术表达式 Arithmetic expression
type SqlAriExp interface {
	SqlCompiler
	// 运算符
	Operator() SqlAriType
	// 运算符左则表达式
	GetLExp() interface{}
	// 运算符右则表达式
	GetRExp() interface{}
}
