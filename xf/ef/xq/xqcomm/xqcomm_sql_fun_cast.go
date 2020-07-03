package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunCast struct {
	*TSqlFun
	val      interface{}
	dataType DbType
}

var _ SqlFun = &TSqlFunCast{}

func NewSqlFunCast(val interface{}, dType DbType) SqlFun {
	inst := &TSqlFunCast{
		val:      val,
		dataType: dType,
	}
	base := newSqlFun(SFDateToUnix, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunCast) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.val == nil || sf.dataType == nil {
		return EmptySqlToken
	}
	prepare := true
	if len(unPrepare) > 0 {
		prepare = !unPrepare[0]
	}

	v1 := ""
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
	result.SetVal(builder.CAST(v1, sf.dataType.GetType(), sf.dataType.GetSize(), sf.dataType.GetDecimal()...))
	return result
}
