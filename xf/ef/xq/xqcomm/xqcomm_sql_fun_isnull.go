package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunIsNull struct {
	*TSqlFun
	val interface{}
}

var _ SqlFun = &TSqlFunIsNull{}

func NewSqlFunIsNull(val interface{}) SqlFun {
	inst := &TSqlFunIsNull{
		val: val,
	}
	base := newSqlFun(SFConcat, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunIsNull) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
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
		if tk := cv1.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			v1 = tk.Val()
		}
	} else {
		if prepare {
			vn1 := cxt.MakeParamId()
			cxt.AddParam(vn1, sf.val)
			result.AddParam(vn1, sf.val)
			v1 = builder.PlaceHolder(vn1)
		} else {
			v1 = builder.MakeRealValue(sf.val)
		}
	}
	result.SetVal(builder.ISNULL(v1))
	return result
}
