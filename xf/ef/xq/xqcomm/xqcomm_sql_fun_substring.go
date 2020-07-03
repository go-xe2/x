package xqcomm

import . "github.com/go-xe2/x/xf/ef/xdriveri"
import . "github.com/go-xe2/x/xf/ef/xqi"

type TSqlFunSubString struct {
	*TSqlFun
	exp  interface{}
	from int
	l    int
}

var _ SqlFun = &TSqlFunSubString{}

func NewSqlFunSubString(exp interface{}, from int, l ...int) SqlFun {
	n := 0
	if len(l) > 0 {
		n = l[0]
	}
	inst := &TSqlFunSubString{
		exp:  exp,
		from: from,
		l:    n,
	}
	base := newSqlFun(SFSubString, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunSubString) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.exp == nil {
		return EmptySqlToken
	}
	field := ""
	result := NewSqlToken("", SqlFunExpressTokenType)
	if cv, ok := sf.exp.(SqlCompiler); ok {
		cxt.PushState(SCPQrSelectFunParamState)
		tk := cv.Compile(builder, cxt, unPrepare...)
		cxt.PopState()
		field = tk.Val()
	} else {
		prepare := true
		if len(unPrepare) > 0 {
			prepare = !unPrepare[0]
		}
		if prepare {
			vn := cxt.MakeParamId()
			cxt.AddParam(vn, sf.exp)
			result.AddParam(vn, sf.exp)
			field = builder.PlaceHolder(vn)
		} else {
			field = builder.MakeRealValue(sf.exp)
		}
	}
	result.SetVal(builder.Substring(field, sf.from, sf.l))
	return result
}
