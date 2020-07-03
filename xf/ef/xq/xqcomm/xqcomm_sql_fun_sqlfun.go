package xqcomm

import . "github.com/go-xe2/x/xf/ef/xdriveri"
import . "github.com/go-xe2/x/xf/ef/xqi"

type TSqlFun struct {
	funId SqlFunId
	this  interface{}
}

var _ SqlFun = &TSqlFun{}

func newSqlFun(id SqlFunId, inherited ...interface{}) *TSqlFun {
	inst := &TSqlFun{
		funId: id,
	}
	inst.this = inst
	if len(inherited) > 0 {
		if _, ok := inherited[0].(SqlFun); ok {
			inst.this = inherited[0]
		}
	}
	return inst
}

func (sf *TSqlFun) Exp() interface{} {
	return sf
}

func (sf *TSqlFun) TokenType() SqlTokenType {
	return SqlFunExpressTokenType
}

func (sf *TSqlFun) FunId() SqlFunId {
	return sf.funId
}

func (sf *TSqlFun) This() interface{} {
	return sf.this
}

func (sf *TSqlFun) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if c, ok := sf.this.(SqlCompiler); ok {
		return c.Compile(builder, cxt, unPrepare...)
	}
	return nil
}
