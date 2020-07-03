package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunIf struct {
	*TSqlFun
	exp  interface{}
	val1 interface{}
	val2 interface{}
}

var _ SqlFun = &TSqlFunIf{}

func NewSqlFunIf(exp interface{}, val1 interface{}, val2 interface{}) SqlFun {
	inst := &TSqlFunIf{
		exp:  exp,
		val1: val1,
		val2: val2,
	}
	base := newSqlFun(SFConcat, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunIf) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.val1 == nil {
		return EmptySqlToken
	}
	v1 := ""
	v2 := ""
	szExp := ""
	prepare := true
	if len(unPrepare) > 0 {
		prepare = !unPrepare[0]
	}
	result := NewSqlToken("", SqlFunExpressTokenType)

	if expC, ok := sf.exp.(SqlCompiler); ok {
		if tk := expC.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			szExp = tk.Val()
		}
	} else {
		if prepare {
			vnExp := cxt.MakeParamId()
			cxt.AddParam(vnExp, sf.exp)
			result.AddParam(vnExp, sf.exp)
			szExp = builder.PlaceHolder(vnExp)
		} else {
			szExp = builder.MakeRealValue(sf.exp)
		}
	}

	if cv1, ok := sf.val1.(SqlCompiler); ok {
		if tk := cv1.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			v1 = tk.Val()
		}
	} else {
		if prepare {
			vn1 := cxt.MakeParamId()
			cxt.AddParam(vn1, sf.val1)
			result.AddParam(vn1, sf.val1)
			v1 = builder.PlaceHolder(vn1)
		} else {
			v1 = builder.MakeRealValue(sf.val1)
		}
	}

	if cv2, ok := sf.val2.(SqlCompiler); ok {
		if tk := cv2.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			v2 = tk.Val()
		}
	} else {
		if prepare {
			vn2 := cxt.MakeParamId()
			cxt.AddParam(vn2, sf.val2)
			result.AddParam(vn2, sf.val2)
			v2 = builder.PlaceHolder(vn2)
		} else {
			v2 = builder.MakeRealValue(sf.val2)
		}
	}

	result.SetVal(builder.If(szExp, v1, v2))
	return result
}
