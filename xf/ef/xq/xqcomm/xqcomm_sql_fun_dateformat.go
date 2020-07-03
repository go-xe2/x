package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunDateFormat struct {
	*TSqlFun
	val    interface{}
	format string
}

var _ SqlFun = &TSqlFunDateFormat{}

func NewSqlFunDateFormat(val interface{}, format string) SqlFun {
	inst := &TSqlFunDateFormat{
		val:    val,
		format: format,
	}
	base := newSqlFun(SFDateFormat, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunDateFormat) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
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
	result.SetVal(builder.DateFormat(v1, sf.format))
	return result
}
