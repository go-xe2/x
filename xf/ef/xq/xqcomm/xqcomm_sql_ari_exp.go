package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/type/xstring"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

// 算术表达式 Arithmetic expression
type TSqlAriExp struct {
	operator SqlAriType
	lExp     interface{}
	rExp     interface{}
	this     interface{}
}

var _ SqlAriExp = &TSqlAriExp{}

func NewSqlAriExp(lExp, rExp interface{}, operator SqlAriType) *TSqlAriExp {
	inst := &TSqlAriExp{
		operator: operator,
		lExp:     lExp,
		rExp:     rExp,
	}
	inst.this = inst
	return inst
}

func (sai *TSqlAriExp) Exp() interface{} {
	return sai
}

func (sai *TSqlAriExp) TokenType() SqlTokenType {
	return SqlArithmeticTokenType
}

func (sai *TSqlAriExp) Operator() SqlAriType {
	return sai.operator
}

func (sai *TSqlAriExp) GetLExp() interface{} {
	return sai.lExp
}

func (sai *TSqlAriExp) GetRExp() interface{} {
	return sai.rExp
}

func (sai *TSqlAriExp) This() interface{} {
	return sai.this
}

func (sai *TSqlAriExp) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	var szLeft, szRight = "", ""
	isPrepare := true
	if len(unPrepare) > 0 {
		isPrepare = !unPrepare[0]
	}
	result := NewSqlToken("", SqlExpressTokenType)
	if lc, ok := sai.lExp.(SqlCompiler); ok {
		if tk := lc.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			szLeft = tk.Val()
		}
	} else {
		if isPrepare {
			vn := cxt.MakeParamId()
			cxt.AddParam(vn, sai.lExp)
			result.AddParam(vn, sai.lExp)
			szLeft = builder.PlaceHolder(vn)
		} else {
			szLeft = builder.MakeRealValue(sai.lExp)
		}
	}

	if rc, ok := sai.rExp.(SqlCompiler); ok {
		if tk := rc.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != SqlEmptyTokenType {
			szRight = tk.Val()
		}
	} else {
		if isPrepare {
			vn := cxt.MakeParamId()
			cxt.AddParam(vn, sai.rExp)
			result.AddParam(vn, sai.rExp)
			szRight = builder.PlaceHolder(vn)
		} else {
			szRight = builder.MakeRealValue(sai.rExp)
		}
	}
	op := sai.operator.Val()
	if szLeft == "" || op == "" || szRight == "" {
		return EmptySqlToken
	}
	result.SetVal(fmt.Sprintf("%s %s %s", xstring.Trim(szLeft), xstring.Trim(op), xstring.Trim(szRight)))
	return result
}
