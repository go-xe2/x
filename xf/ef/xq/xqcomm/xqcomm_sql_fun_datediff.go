package xqcomm

import . "github.com/go-xe2/x/xf/ef/xdriveri"
import . "github.com/go-xe2/x/xf/ef/xqi"

type TSqlFunDateDiff struct {
	*TSqlFun
	val1     interface{}
	val2     interface{}
	datePart DatePart
}

var _ SqlFun = &TSqlFunDateDiff{}

func NewSqlFunDateDiff(val1 interface{}, val2 interface{}, datePart DatePart) SqlFun {
	inst := &TSqlFunDateDiff{
		val1:     val1,
		val2:     val2,
		datePart: datePart,
	}
	base := newSqlFun(SFDateDiff, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunDateDiff) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.val1 == nil || sf.val2 == nil {
		return EmptySqlToken
	}
	p1 := ""
	p2 := ""
	prepare := true
	if len(unPrepare) > 0 {
		prepare = !unPrepare[0]
	}

	result := NewSqlToken("", SqlFunExpressTokenType)

	if cv1, ok := sf.val1.(SqlCompiler); ok {
		cxt.PushState(SCPQrSelectFunParamState)
		if tk := cv1.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			p1 = tk.Val()
		}
		cxt.PopState()
	} else {
		if prepare {
			vn1 := cxt.MakeParamId()
			cxt.AddParam(vn1, sf.val1)
			result.AddParam(vn1, sf.val1)
			p1 = builder.PlaceHolder(vn1)
		} else {
			p1 = builder.MakeRealValue(sf.val1)
		}
	}

	if cv2, ok := sf.val2.(SqlCompiler); ok {
		cxt.PushState(SCPQrSelectFunParamState)
		if tk := cv2.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() == SqlEmptyTokenType {
			p2 = tk.Val()
		}
		cxt.PopState()
	} else {
		if prepare {
			vn2 := cxt.MakeParamId()
			cxt.AddParam(vn2, sf.val2)
			result.AddParam(vn2, sf.val2)
			p2 = builder.PlaceHolder(vn2)
		} else {
			p2 = builder.MakeRealValue(sf.val2)
		}
	}

	result.SetVal(builder.DateDiff(p1, p2, sf.datePart))
	return result
}
