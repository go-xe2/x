package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunDateToUnix struct {
	*TSqlFun
	val interface{}
}

var _ SqlFun = &TSqlFunDateToUnix{}

func NewSqlFunDateToUnix(val interface{}) SqlFun {
	inst := &TSqlFunDateToUnix{
		val: val,
	}
	base := newSqlFun(SFDateToUnix, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunDateToUnix) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.val == nil {
		return EmptySqlToken
	}
	v1 := ""
	prepare := true
	result := NewSqlToken("", SqlFunExpressTokenType)

	if len(unPrepare) > 0 {
		prepare = !unPrepare[0]
	}
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
	result.SetVal(builder.DateToUnix(v1))
	return result
}
