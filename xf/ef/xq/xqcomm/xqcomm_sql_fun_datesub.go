package xqcomm

import . "github.com/go-xe2/x/xf/ef/xdriveri"
import . "github.com/go-xe2/x/xf/ef/xqi"

type TSqlFunDateSub struct {
	*TSqlFun
	val      interface{}
	interval int
	datePart DatePart
}

var _ SqlFun = &TSqlFunDateSub{}

func NewSqlFunDateSub(val interface{}, interval int, datePart DatePart) SqlFun {
	inst := &TSqlFunDateSub{
		val:      val,
		interval: interval,
		datePart: datePart,
	}
	base := newSqlFun(SFDateSub, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunDateSub) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.val == nil {
		return EmptySqlToken
	}
	p1 := ""
	prepare := true
	if len(unPrepare) > 0 {
		prepare = !unPrepare[0]
	}
	result := NewSqlToken("", SqlFunExpressTokenType)

	if cv1, ok := sf.val.(SqlCompiler); ok {
		cxt.PushState(SCPQrSelectFunParamState)
		if tk := cv1.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			p1 = tk.Val()
		}
		cxt.PopState()
	} else {
		if prepare {
			vn := cxt.MakeParamId()
			cxt.AddParam(vn, sf.val)
			result.AddParam(vn, sf.val)
			p1 = builder.PlaceHolder(vn)
		} else {
			p1 = builder.MakeRealValue(sf.val)
		}
	}
	result.SetVal(builder.DateSub(p1, sf.interval, sf.datePart))
	return result
}
