package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunCase struct {
	*TSqlFun
	exp      SqlCompiler
	whenThen [][]interface{}
	elseVal  interface{}
}

var _ SqlCase = &TSqlFunCase{}

func NewSqlFunCase(exp ...SqlCompiler) *TSqlFunCase {
	var tmp SqlCompiler
	if len(exp) > 0 {
		tmp = exp[0]
	}
	inst := &TSqlFunCase{
		exp:      tmp,
		whenThen: make([][]interface{}, 0),
	}
	base := newSqlFun(SFConcat, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunCase) FunId() SqlFunId {
	return SFCase
}

func (sf *TSqlFunCase) When(when, then interface{}) SqlCase {
	sf.whenThen = append(sf.whenThen, []interface{}{when, then})
	return sf
}

func (sf *TSqlFunCase) ElseEnd(val interface{}) SqlField {
	sf.elseVal = val
	return NewSqlField(sf, "")
}

func (sf *TSqlFunCase) ThenElse(thenValue, elseValue interface{}) SqlField {
	sf.whenThen = [][]interface{}{
		{sf.exp, thenValue},
	}
	sf.exp = nil
	sf.elseVal = elseValue
	return NewSqlField(sf, "")
}

func (sf *TSqlFunCase) End() SqlField {
	return NewSqlField(sf, "")
}

func (sf *TSqlFunCase) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.exp == nil && len(sf.whenThen) == 0 {
		return EmptySqlToken
	}
	v1 := ""
	result := NewSqlToken("", SqlExpressTokenType)

	if cv1, ok := sf.exp.(SqlCompiler); ok {
		if tk := cv1.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			v1 = tk.Val()
		}
	} else {
		v1 = ""
	}
	whenThen := make([][]string, 0)
	for _, items := range sf.whenThen {
		if len(items) != 2 {
			continue
		}
		whenItem := items[0]
		thenItem := items[1]
		var szWhen, szThen = "", ""
		if wc, ok := whenItem.(SqlCompiler); ok {
			if tk := wc.Compile(builder, cxt, unPrepare...); ok {
				szWhen = tk.Val()
			}
		} else {
			szWhen = builder.MakeRealValue(whenItem)
		}
		if tc, ok := thenItem.(SqlCompiler); szWhen != "" && ok {
			if tk := tc.Compile(builder, cxt, unPrepare...); ok {
				szThen = tk.Val()
			}
		} else {
			szThen = builder.MakeRealValue(thenItem)
		}
		if szWhen != "" && szThen != "" {
			whenThen = append(whenThen, []string{szWhen, szThen})
		}
	}

	if len(whenThen) <= 0 {
		return EmptySqlToken
	}
	szElse := ""
	if ec, ok := sf.elseVal.(SqlCompiler); ok {
		if tk := ec.Compile(builder, cxt, unPrepare...); tk != nil {
			szElse = tk.Val()
		}
	} else {
		szElse = builder.MakeRealValue(sf.elseVal)
	}
	result.SetVal(builder.Case(v1, whenThen, szElse))
	return result
}
