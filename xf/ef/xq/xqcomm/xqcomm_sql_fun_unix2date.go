package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunUnixToDate struct {
	*TSqlFun
	val interface{}
}

var _ SqlFun = &TSqlFunUnixToDate{}

func NewSqlFunUnixToDate(val interface{}) SqlFun {
	inst := &TSqlFunUnixToDate{
		val: val,
	}
	base := newSqlFun(SFUnixToDate, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunUnixToDate) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.val == nil {
		return EmptySqlToken
	}
	v1 := ""
	prepare := true
	if len(unPrepare) > 0 {
		prepare = !unPrepare[0]
	}

	result := NewSqlToken("", SqlFunExpressTokenType)

	if cv1, ok := sf.val.(SqlCompiler); ok {
		cxt.PushState(SCPQrSelectFunParamState)
		if tk := cv1.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			v1 = tk.Val()
		}
		cxt.PopState()
	} else {
		if prepare {
			vn := cxt.MakeParamId()
			cxt.AddParam(vn, sf.val)
			result.AddParam(vn, sf.val)
			v1 = builder.PlaceHolder(vn)
		} else {
			v1 = builder.MakeRealValue(sf.val)
		}
	}
	result.SetVal(builder.UnixToDate(v1))
	return result
}
